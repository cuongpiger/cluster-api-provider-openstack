/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package networking

import (
	"fmt"

	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/attributestags"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/record"
)

const (
	secGroupPrefix     string = "k8s"
	controlPlaneSuffix string = "controlplane"
	workerSuffix       string = "worker"
	bastionSuffix      string = "bastion"
	allNodesSuffix     string = "allNodes"
	remoteGroupIDSelf  string = "self"
)

// ReconcileSecurityGroups reconcile the security groups.
func (s *Service) ReconcileSecurityGroups(openStackCluster *infrav1.OpenStackCluster, clusterName string) error {
	s.scope.Logger().Info("Reconciling security groups")
	if openStackCluster.Spec.ManagedSecurityGroups == nil {
		s.scope.Logger().V(4).Info("No need to reconcile security groups")
		return nil
	}

	secControlPlaneGroupName := getSecControlPlaneGroupName(clusterName)
	secWorkerGroupName := getSecWorkerGroupName(clusterName)
	secGroupNames := map[string]string{
		controlPlaneSuffix: secControlPlaneGroupName,
		workerSuffix:       secWorkerGroupName,
	}

	if openStackCluster.Spec.Bastion != nil && openStackCluster.Spec.Bastion.Enabled {
		secBastionGroupName := getSecBastionGroupName(clusterName)
		secGroupNames[bastionSuffix] = secBastionGroupName
	}

	// create security groups first, because desired rules use group ids.
	for _, v := range secGroupNames {
		if err := s.createSecurityGroupIfNotExists(openStackCluster, v); err != nil {
			return err
		}
	}
	// create desired security groups
	desiredSecGroups, err := s.generateDesiredSecGroups(openStackCluster, secGroupNames)
	if err != nil {
		return err
	}

	observedSecGroups := make(map[string]*infrav1.SecurityGroupStatus)
	for k, desiredSecGroup := range desiredSecGroups {
		var err error
		observedSecGroups[k], err = s.getSecurityGroupByName(desiredSecGroup.Name)

		if err != nil {
			return err
		}

		if observedSecGroups[k].ID != "" {
			observedSecGroup, err := s.reconcileGroupRules(desiredSecGroup, *observedSecGroups[k])
			if err != nil {
				return err
			}
			observedSecGroups[k] = &observedSecGroup
			continue
		}
	}

	openStackCluster.Status.ControlPlaneSecurityGroup = observedSecGroups[controlPlaneSuffix]
	openStackCluster.Status.WorkerSecurityGroup = observedSecGroups[workerSuffix]
	openStackCluster.Status.BastionSecurityGroup = observedSecGroups[bastionSuffix]

	return nil
}

type securityGroupSpec struct {
	Name  string
	Rules []resolvedSecurityGroupRuleSpec
}

type resolvedSecurityGroupRuleSpec struct {
	Description    string `json:"description,omitempty"`
	Direction      string `json:"direction,omitempty"`
	EtherType      string `json:"etherType,omitempty"`
	PortRangeMin   int    `json:"portRangeMin,omitempty"`
	PortRangeMax   int    `json:"portRangeMax,omitempty"`
	Protocol       string `json:"protocol,omitempty"`
	RemoteGroupID  string `json:"remoteGroupID,omitempty"`
	RemoteIPPrefix string `json:"remoteIPPrefix,omitempty"`
}

func (r resolvedSecurityGroupRuleSpec) Matches(other infrav1.SecurityGroupRuleStatus) bool {
	return r.Description == *other.Description &&
		r.Direction == other.Direction &&
		r.EtherType == *other.EtherType &&
		r.PortRangeMin == *other.PortRangeMin &&
		r.PortRangeMax == *other.PortRangeMax &&
		r.Protocol == *other.Protocol &&
		r.RemoteGroupID == *other.RemoteGroupID &&
		r.RemoteIPPrefix == *other.RemoteIPPrefix
}

func (s *Service) generateDesiredSecGroups(openStackCluster *infrav1.OpenStackCluster, secGroupNames map[string]string) (map[string]securityGroupSpec, error) {
	if openStackCluster.Spec.ManagedSecurityGroups == nil {
		return nil, nil
	}

	desiredSecGroups := make(map[string]securityGroupSpec)

	var secControlPlaneGroupID string
	var secWorkerGroupID string
	var secBastionGroupID string

	// remoteManagedGroups is a map of suffix to security group ID.
	// It will be used to fill in the RemoteGroupID field of the security group rules
	// that reference a managed security group.
	// For now, we only reference the control plane and worker security groups.
	remoteManagedGroups := make(map[string]string)

	for i, v := range secGroupNames {
		secGroup, err := s.getSecurityGroupByName(v)
		if err != nil {
			return desiredSecGroups, err
		}
		switch i {
		case controlPlaneSuffix:
			secControlPlaneGroupID = secGroup.ID
			remoteManagedGroups[controlPlaneSuffix] = secControlPlaneGroupID
		case workerSuffix:
			secWorkerGroupID = secGroup.ID
			remoteManagedGroups[workerSuffix] = secWorkerGroupID
		case bastionSuffix:
			secBastionGroupID = secGroup.ID
			remoteManagedGroups[bastionSuffix] = secBastionGroupID
		}
	}

	// Start with the default rules
	controlPlaneRules := append([]resolvedSecurityGroupRuleSpec{}, defaultRules...)
	workerRules := append([]resolvedSecurityGroupRuleSpec{}, defaultRules...)

	controlPlaneRules = append(controlPlaneRules, getSGControlPlaneHTTPS()...)
	workerRules = append(workerRules, getSGWorkerNodePort()...)

	// If we set additional ports to LB, we need create secgroup rules those ports, this apply to controlPlaneRules only
	if openStackCluster.Spec.APIServerLoadBalancer.Enabled {
		controlPlaneRules = append(controlPlaneRules, getSGControlPlaneAdditionalPorts(openStackCluster.Spec.APIServerLoadBalancer.AdditionalPorts)...)
	}

	if openStackCluster.Spec.ManagedSecurityGroups != nil && openStackCluster.Spec.ManagedSecurityGroups.AllowAllInClusterTraffic {
		// Permit all ingress from the cluster security groups
		controlPlaneRules = append(controlPlaneRules, getSGControlPlaneAllowAll(remoteGroupIDSelf, secWorkerGroupID)...)
		workerRules = append(workerRules, getSGWorkerAllowAll(remoteGroupIDSelf, secControlPlaneGroupID)...)
	} else {
		controlPlaneRules = append(controlPlaneRules, getSGControlPlaneGeneral(remoteGroupIDSelf, secWorkerGroupID)...)
		workerRules = append(workerRules, getSGWorkerGeneral(remoteGroupIDSelf, secControlPlaneGroupID)...)
	}

	// For now, we do not create a separate security group for allNodes.
	// Instead, we append the rules for allNodes to the control plane and worker security groups.
	allNodesRules, err := getAllNodesRules(remoteManagedGroups, openStackCluster.Spec.ManagedSecurityGroups.AllNodesSecurityGroupRules)
	if err != nil {
		return desiredSecGroups, err
	}
	controlPlaneRules = append(controlPlaneRules, allNodesRules...)
	workerRules = append(workerRules, allNodesRules...)

	if openStackCluster.Spec.Bastion != nil && openStackCluster.Spec.Bastion.Enabled {
		controlPlaneRules = append(controlPlaneRules, getSGControlPlaneSSH(secBastionGroupID)...)
		workerRules = append(workerRules, getSGWorkerSSH(secBastionGroupID)...)

		desiredSecGroups[bastionSuffix] = securityGroupSpec{
			Name: secGroupNames[bastionSuffix],
			Rules: append(
				[]resolvedSecurityGroupRuleSpec{
					{
						Description:  "SSH",
						Direction:    "ingress",
						EtherType:    "IPv4",
						PortRangeMin: 22,
						PortRangeMax: 22,
						Protocol:     "tcp",
					},
				},
				defaultRules...,
			),
		}
	}

	desiredSecGroups[controlPlaneSuffix] = securityGroupSpec{
		Name:  secGroupNames[controlPlaneSuffix],
		Rules: controlPlaneRules,
	}

	desiredSecGroups[workerSuffix] = securityGroupSpec{
		Name:  secGroupNames[workerSuffix],
		Rules: workerRules,
	}
	return desiredSecGroups, nil
}

// getAllNodesRules returns the rules for the allNodes security group that should be created.
func getAllNodesRules(remoteManagedGroups map[string]string, allNodesSecurityGroupRules []infrav1.SecurityGroupRuleSpec) ([]resolvedSecurityGroupRuleSpec, error) {
	rules := make([]resolvedSecurityGroupRuleSpec, 0, len(allNodesSecurityGroupRules))
	for _, rule := range allNodesSecurityGroupRules {
		if err := validateRemoteManagedGroups(remoteManagedGroups, rule.RemoteManagedGroups); err != nil {
			return nil, err
		}
		r := resolvedSecurityGroupRuleSpec{
			Direction: rule.Direction,
		}
		if rule.Description != nil {
			r.Description = *rule.Description
		}
		if rule.EtherType != nil {
			r.EtherType = *rule.EtherType
		}
		if rule.PortRangeMin != nil {
			r.PortRangeMin = *rule.PortRangeMin
		}
		if rule.PortRangeMax != nil {
			r.PortRangeMax = *rule.PortRangeMax
		}
		if rule.Protocol != nil {
			r.Protocol = *rule.Protocol
		}
		if rule.RemoteGroupID != nil {
			r.RemoteGroupID = *rule.RemoteGroupID
		}
		if rule.RemoteIPPrefix != nil {
			r.RemoteIPPrefix = *rule.RemoteIPPrefix
		}

		if len(rule.RemoteManagedGroups) > 0 {
			if rule.RemoteGroupID != nil {
				return nil, fmt.Errorf("remoteGroupID must not be set when remoteManagedGroups is set")
			}

			for _, rg := range rule.RemoteManagedGroups {
				rc := r
				rc.RemoteGroupID = remoteManagedGroups[rg.String()]
				rules = append(rules, rc)
			}
		} else {
			rules = append(rules, r)
		}
	}
	return rules, nil
}

// validateRemoteManagedGroups validates that the remoteManagedGroups target existing managed security groups.
func validateRemoteManagedGroups(remoteManagedGroups map[string]string, ruleRemoteManagedGroups []infrav1.ManagedSecurityGroupName) error {
	if len(ruleRemoteManagedGroups) == 0 {
		return fmt.Errorf("remoteManagedGroups is required")
	}

	for _, group := range ruleRemoteManagedGroups {
		if _, ok := remoteManagedGroups[group.String()]; !ok {
			return fmt.Errorf("remoteManagedGroups: %s is not a valid remote managed security group", group)
		}
	}
	return nil
}

func (s *Service) GetSecurityGroups(securityGroupParams []infrav1.SecurityGroupFilter) ([]string, error) {
	var sgIDs []string
	for _, sg := range securityGroupParams {
		// Don't validate an explicit UUID if we were given one
		if sg.ID != "" {
			if isDuplicate(sgIDs, sg.ID) {
				continue
			}
			sgIDs = append(sgIDs, sg.ID)
			continue
		}

		listOpts := sg.ToListOpt()
		if listOpts.ProjectID == "" {
			listOpts.ProjectID = s.scope.ProjectID()
		}
		SGList, err := s.client.ListSecGroup(listOpts)
		if err != nil {
			return nil, err
		}

		if len(SGList) == 0 {
			return nil, fmt.Errorf("security group %s not found", sg.Name)
		}

		for _, group := range SGList {
			if isDuplicate(sgIDs, group.ID) {
				continue
			}
			sgIDs = append(sgIDs, group.ID)
		}
	}
	return sgIDs, nil
}

func (s *Service) DeleteSecurityGroups(openStackCluster *infrav1.OpenStackCluster, clusterName string) error {
	secGroupNames := []string{
		getSecControlPlaneGroupName(clusterName),
		getSecWorkerGroupName(clusterName),
	}

	if openStackCluster.Spec.Bastion != nil && openStackCluster.Spec.Bastion.Enabled {
		secGroupNames = append(secGroupNames, getSecBastionGroupName(clusterName))
	}

	for _, secGroupName := range secGroupNames {
		if err := s.deleteSecurityGroup(openStackCluster, secGroupName); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) deleteSecurityGroup(openStackCluster *infrav1.OpenStackCluster, name string) error {
	group, err := s.getSecurityGroupByName(name)
	if err != nil {
		return err
	}
	if group.ID == "" {
		// nothing to do
		return nil
	}
	err = s.client.DeleteSecGroup(group.ID)
	if err != nil {
		record.Warnf(openStackCluster, "FailedDeleteSecurityGroup", "Failed to delete security group %s with id %s: %v", group.Name, group.ID, err)
		return err
	}

	record.Eventf(openStackCluster, "SuccessfulDeleteSecurityGroup", "Deleted security group %s with id %s", group.Name, group.ID)
	return nil
}

// reconcileGroupRules reconciles an already existing observed group by deleting rules not needed anymore and
// creating rules that are missing.
func (s *Service) reconcileGroupRules(desired securityGroupSpec, observed infrav1.SecurityGroupStatus) (infrav1.SecurityGroupStatus, error) {
	var rulesToDelete []string
	// fills rulesToDelete by calculating observed - desired
	for _, observedRule := range observed.Rules {
		deleteRule := true
		for _, desiredRule := range desired.Rules {
			r := desiredRule
			if r.RemoteGroupID == remoteGroupIDSelf {
				r.RemoteGroupID = observed.ID
			}
			if r.Matches(observedRule) {
				deleteRule = false
				break
			}
		}
		if deleteRule {
			rulesToDelete = append(rulesToDelete, observedRule.ID)
		}
	}

	rulesToCreate := []resolvedSecurityGroupRuleSpec{}
	reconciledRules := make([]infrav1.SecurityGroupRuleStatus, 0, len(desired.Rules))
	// fills rulesToCreate by calculating desired - observed
	// also adds rules which are in observed and desired to reconcileGroupRules.
	for _, desiredRule := range desired.Rules {
		r := desiredRule
		if r.RemoteGroupID == remoteGroupIDSelf {
			r.RemoteGroupID = observed.ID
		}
		createRule := true
		for _, observedRule := range observed.Rules {
			if r.Matches(observedRule) {
				// add already existing rules to reconciledRules because we won't touch them anymore
				reconciledRules = append(reconciledRules, observedRule)
				createRule = false
				break
			}
		}
		if createRule {
			rulesToCreate = append(rulesToCreate, desiredRule)
		}
	}

	s.scope.Logger().V(4).Info("Deleting rules not needed anymore for group", "name", observed.Name, "amount", len(rulesToDelete))
	for _, rule := range rulesToDelete {
		s.scope.Logger().V(6).Info("Deleting rule", "ID", rule, "name", observed.Name)
		err := s.client.DeleteSecGroupRule(rule)
		if err != nil {
			return infrav1.SecurityGroupStatus{}, err
		}
	}

	s.scope.Logger().V(4).Info("Creating new rules needed for group", "name", observed.Name, "amount", len(rulesToCreate))
	for _, rule := range rulesToCreate {
		r := rule
		if r.RemoteGroupID == remoteGroupIDSelf {
			r.RemoteGroupID = observed.ID
		}
		newRule, err := s.createRule(observed.ID, r)
		if err != nil {
			return infrav1.SecurityGroupStatus{}, err
		}
		reconciledRules = append(reconciledRules, newRule)
	}
	observed.Rules = reconciledRules

	if len(reconciledRules) == 0 {
		return infrav1.SecurityGroupStatus{}, nil
	}

	return observed, nil
}

func (s *Service) createSecurityGroupIfNotExists(openStackCluster *infrav1.OpenStackCluster, groupName string) error {
	secGroup, err := s.getSecurityGroupByName(groupName)
	if err != nil {
		return err
	}
	if secGroup == nil || secGroup.ID == "" {
		s.scope.Logger().V(6).Info("Group doesn't exist, creating it", "name", groupName)

		createOpts := groups.CreateOpts{
			Name:        groupName,
			Description: "Cluster API managed group",
		}
		s.scope.Logger().V(6).Info("Creating group", "name", groupName)

		group, err := s.client.CreateSecGroup(createOpts)
		if err != nil {
			record.Warnf(openStackCluster, "FailedCreateSecurityGroup", "Failed to create security group %s: %v", groupName, err)
			return err
		}

		if len(openStackCluster.Spec.Tags) > 0 {
			_, err = s.client.ReplaceAllAttributesTags("security-groups", group.ID, attributestags.ReplaceAllOpts{
				Tags: openStackCluster.Spec.Tags,
			})
			if err != nil {
				return err
			}
		}

		record.Eventf(openStackCluster, "SuccessfulCreateSecurityGroup", "Created security group %s with id %s", groupName, group.ID)
		return nil
	}

	sInfo := fmt.Sprintf("Reuse Existing SecurityGroup %s with %s", groupName, secGroup.ID)
	s.scope.Logger().V(6).Info(sInfo)

	return nil
}

func (s *Service) getSecurityGroupByName(name string) (*infrav1.SecurityGroupStatus, error) {
	opts := groups.ListOpts{
		Name: name,
	}

	s.scope.Logger().V(6).Info("Attempting to fetch security group with", "name", name)
	allGroups, err := s.client.ListSecGroup(opts)
	if err != nil {
		return &infrav1.SecurityGroupStatus{}, err
	}

	switch len(allGroups) {
	case 0:
		return &infrav1.SecurityGroupStatus{}, nil
	case 1:
		return convertOSSecGroupToConfigSecGroup(allGroups[0]), nil
	}

	return &infrav1.SecurityGroupStatus{}, fmt.Errorf("more than one security group found named: %s", name)
}

func (s *Service) createRule(securityGroupID string, r resolvedSecurityGroupRuleSpec) (infrav1.SecurityGroupRuleStatus, error) {
	dir := rules.RuleDirection(r.Direction)
	proto := rules.RuleProtocol(r.Protocol)
	etherType := rules.RuleEtherType(r.EtherType)

	createOpts := rules.CreateOpts{
		Description:    r.Description,
		Direction:      dir,
		PortRangeMin:   r.PortRangeMin,
		PortRangeMax:   r.PortRangeMax,
		Protocol:       proto,
		EtherType:      etherType,
		RemoteGroupID:  r.RemoteGroupID,
		RemoteIPPrefix: r.RemoteIPPrefix,
		SecGroupID:     securityGroupID,
	}
	s.scope.Logger().V(6).Info("Creating rule", "description", r.Description, "direction", dir, "portRangeMin", r.PortRangeMin, "portRangeMax", r.PortRangeMax, "proto", proto, "etherType", etherType, "remoteGroupID", r.RemoteGroupID, "remoteIPPrefix", r.RemoteIPPrefix, "securityGroupID", securityGroupID)
	rule, err := s.client.CreateSecGroupRule(createOpts)
	if err != nil {
		return infrav1.SecurityGroupRuleStatus{}, err
	}
	return convertOSSecGroupRuleToConfigSecGroupRule(*rule), nil
}

func getSecControlPlaneGroupName(clusterName string) string {
	return fmt.Sprintf("%s-cluster-%s-secgroup-%s", secGroupPrefix, clusterName, controlPlaneSuffix)
}

func getSecWorkerGroupName(clusterName string) string {
	return fmt.Sprintf("%s-cluster-%s-secgroup-%s", secGroupPrefix, clusterName, workerSuffix)
}

func getSecBastionGroupName(clusterName string) string {
	return fmt.Sprintf("%s-cluster-%s-secgroup-%s", secGroupPrefix, clusterName, bastionSuffix)
}

func convertOSSecGroupToConfigSecGroup(osSecGroup groups.SecGroup) *infrav1.SecurityGroupStatus {
	securityGroupRules := make([]infrav1.SecurityGroupRuleStatus, len(osSecGroup.Rules))
	for i, rule := range osSecGroup.Rules {
		securityGroupRules[i] = convertOSSecGroupRuleToConfigSecGroupRule(rule)
	}
	return &infrav1.SecurityGroupStatus{
		ID:    osSecGroup.ID,
		Name:  osSecGroup.Name,
		Rules: securityGroupRules,
	}
}

func convertOSSecGroupRuleToConfigSecGroupRule(osSecGroupRule rules.SecGroupRule) infrav1.SecurityGroupRuleStatus {
	return infrav1.SecurityGroupRuleStatus{
		ID:             osSecGroupRule.ID,
		Direction:      osSecGroupRule.Direction,
		Description:    &osSecGroupRule.Description,
		EtherType:      &osSecGroupRule.EtherType,
		PortRangeMin:   &osSecGroupRule.PortRangeMin,
		PortRangeMax:   &osSecGroupRule.PortRangeMax,
		Protocol:       &osSecGroupRule.Protocol,
		RemoteGroupID:  &osSecGroupRule.RemoteGroupID,
		RemoteIPPrefix: &osSecGroupRule.RemoteIPPrefix,
	}
}

func isDuplicate(list []string, name string) bool {
	if len(list) == 0 {
		return false
	}
	for _, element := range list {
		if element == name {
			return true
		}
	}
	return false
}
