/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha5

import (
	"strings"

	conversion "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/utils/pointer"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
)

var _ ctrlconversion.Convertible = &OpenStackCluster{}

const trueString = "true"

func (r *OpenStackCluster) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackCluster)

	if err := Convert_v1alpha5_OpenStackCluster_To_v1beta1_OpenStackCluster(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackCluster{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackCluster) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackCluster)

	if err := Convert_v1beta1_OpenStackCluster_To_v1alpha5_OpenStackCluster(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackClusterList{}

func (r *OpenStackClusterList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1alpha5_OpenStackClusterList_To_v1beta1_OpenStackClusterList(r, dst, nil)
}

func (r *OpenStackClusterList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterList)

	return Convert_v1beta1_OpenStackClusterList_To_v1alpha5_OpenStackClusterList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackClusterTemplate{}

func (r *OpenStackClusterTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackClusterTemplate)

	if err := Convert_v1alpha5_OpenStackClusterTemplate_To_v1beta1_OpenStackClusterTemplate(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackClusterTemplate{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackClusterTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackClusterTemplate)

	if err := Convert_v1beta1_OpenStackClusterTemplate_To_v1alpha5_OpenStackClusterTemplate(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachine{}

func (r *OpenStackMachine) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachine)

	if err := Convert_v1alpha5_OpenStackMachine_To_v1beta1_OpenStackMachine(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackMachine{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackMachine) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachine)

	if err := Convert_v1beta1_OpenStackMachine_To_v1alpha5_OpenStackMachine(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachineList{}

func (r *OpenStackMachineList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineList)

	return Convert_v1alpha5_OpenStackMachineList_To_v1beta1_OpenStackMachineList(r, dst, nil)
}

func (r *OpenStackMachineList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineList)

	return Convert_v1beta1_OpenStackMachineList_To_v1alpha5_OpenStackMachineList(src, r, nil)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplate{}

func (r *OpenStackMachineTemplate) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplate)

	if err := Convert_v1alpha5_OpenStackMachineTemplate_To_v1beta1_OpenStackMachineTemplate(r, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1.OpenStackMachineTemplate{}
	if ok, err := utilconversion.UnmarshalData(r, restored); err != nil || !ok {
		return err
	}

	return nil
}

func (r *OpenStackMachineTemplate) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplate)

	if err := Convert_v1beta1_OpenStackMachineTemplate_To_v1alpha5_OpenStackMachineTemplate(src, r, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, r)
}

var _ ctrlconversion.Convertible = &OpenStackMachineTemplateList{}

func (r *OpenStackMachineTemplateList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*infrav1.OpenStackMachineTemplateList)

	return Convert_v1alpha5_OpenStackMachineTemplateList_To_v1beta1_OpenStackMachineTemplateList(r, dst, nil)
}

func (r *OpenStackMachineTemplateList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*infrav1.OpenStackMachineTemplateList)

	return Convert_v1beta1_OpenStackMachineTemplateList_To_v1alpha5_OpenStackMachineTemplateList(src, r, nil)
}

func Convert_v1beta1_OpenStackClusterSpec_To_v1alpha5_OpenStackClusterSpec(in *infrav1.OpenStackClusterSpec, out *OpenStackClusterSpec, s conversion.Scope) error {
	err := autoConvert_v1beta1_OpenStackClusterSpec_To_v1alpha5_OpenStackClusterSpec(in, out, s)
	if err != nil {
		return err
	}

	if in.ExternalNetwork.ID != "" {
		out.ExternalNetworkID = in.ExternalNetwork.ID
	}

	if len(in.ManagedSubnets) > 0 {
		out.NodeCIDR = in.ManagedSubnets[0].CIDR
		out.DNSNameservers = in.ManagedSubnets[0].DNSNameservers
	}

	if in.Subnets != nil {
		if len(in.Subnets) >= 1 {
			if err := Convert_v1beta1_SubnetFilter_To_v1alpha5_SubnetFilter(&in.Subnets[0], &out.Subnet, s); err != nil {
				return err
			}
		}
	}

	if in.ManagedSecurityGroups != nil {
		out.ManagedSecurityGroups = true
		out.AllowAllInClusterTraffic = in.ManagedSecurityGroups.AllowAllInClusterTraffic
	}

	out.CloudName = in.IdentityRef.CloudName
	out.IdentityRef = &OpenStackIdentityReference{Name: in.IdentityRef.Name}

	return nil
}

func Convert_v1alpha5_OpenStackClusterSpec_To_v1beta1_OpenStackClusterSpec(in *OpenStackClusterSpec, out *infrav1.OpenStackClusterSpec, s conversion.Scope) error {
	err := autoConvert_v1alpha5_OpenStackClusterSpec_To_v1beta1_OpenStackClusterSpec(in, out, s)
	if err != nil {
		return err
	}

	if in.ExternalNetworkID != "" {
		out.ExternalNetwork = infrav1.NetworkFilter{
			ID: in.ExternalNetworkID,
		}
	}

	emptySubnet := SubnetFilter{}
	if in.Subnet != emptySubnet {
		subnet := infrav1.SubnetFilter{}
		if err := Convert_v1alpha5_SubnetFilter_To_v1beta1_SubnetFilter(&in.Subnet, &subnet, s); err != nil {
			return err
		}
		out.Subnets = []infrav1.SubnetFilter{subnet}
	}

	if len(in.NodeCIDR) > 0 {
		out.ManagedSubnets = []infrav1.SubnetSpec{
			{
				CIDR:           in.NodeCIDR,
				DNSNameservers: in.DNSNameservers,
			},
		}
	}
	// We're dropping DNSNameservers even if these were set as without NodeCIDR it doesn't make sense.

	if in.ManagedSecurityGroups {
		out.ManagedSecurityGroups = &infrav1.ManagedSecurityGroups{}
		if !in.AllowAllInClusterTraffic {
			out.ManagedSecurityGroups.AllNodesSecurityGroupRules = infrav1.LegacyCalicoSecurityGroupRules()
		} else {
			out.ManagedSecurityGroups.AllowAllInClusterTraffic = true
		}
	}

	out.IdentityRef.CloudName = in.CloudName
	if in.IdentityRef != nil {
		out.IdentityRef.Name = in.IdentityRef.Name
	}

	return nil
}

func Convert_v1beta1_LoadBalancer_To_v1alpha5_LoadBalancer(in *infrav1.LoadBalancer, out *LoadBalancer, s conversion.Scope) error {
	return autoConvert_v1beta1_LoadBalancer_To_v1alpha5_LoadBalancer(in, out, s)
}

func Convert_v1beta1_PortOpts_To_v1alpha5_PortOpts(in *infrav1.PortOpts, out *PortOpts, s conversion.Scope) error {
	// value specs and propagate uplink status have been added in v1beta1 but have no equivalent in v1alpha5
	err := autoConvert_v1beta1_PortOpts_To_v1alpha5_PortOpts(in, out, s)
	if err != nil {
		return err
	}

	// The auto-generated function converts v1beta1 SecurityGroup to
	// v1alpha6 SecurityGroup, but v1alpha6 SecurityGroupFilter is more
	// appropriate. Unset them and convert to SecurityGroupFilter instead.
	out.SecurityGroups = nil
	if len(in.SecurityGroups) > 0 {
		out.SecurityGroupFilters = make([]SecurityGroupParam, len(in.SecurityGroups))
		for i := range in.SecurityGroups {
			securityGroupParam := &out.SecurityGroupFilters[i]
			if in.SecurityGroups[i].ID != "" {
				securityGroupParam.UUID = in.SecurityGroups[i].ID
			} else {
				if err := Convert_v1beta1_SecurityGroupFilter_To_v1alpha5_SecurityGroupFilter(&in.SecurityGroups[i], &securityGroupParam.Filter, s); err != nil {
					return err
				}
			}
		}
	}

	out.Profile = make(map[string]string)
	if pointer.BoolDeref(in.Profile.OVSHWOffload, false) {
		(out.Profile)["capabilities"] = "[\"switchdev\"]"
	}
	if pointer.BoolDeref(in.Profile.TrustedVF, false) {
		(out.Profile)["trusted"] = trueString
	}
	return nil
}

func Convert_v1alpha5_OpenStackMachineSpec_To_v1beta1_OpenStackMachineSpec(in *OpenStackMachineSpec, out *infrav1.OpenStackMachineSpec, s conversion.Scope) error {
	err := autoConvert_v1alpha5_OpenStackMachineSpec_To_v1beta1_OpenStackMachineSpec(in, out, s)
	if err != nil {
		return err
	}

	if in.ServerGroupID != "" {
		out.ServerGroup = &infrav1.ServerGroupFilter{ID: in.ServerGroupID}
	} else {
		out.ServerGroup = &infrav1.ServerGroupFilter{}
	}

	imageFilter := infrav1.ImageFilter{}
	if in.Image != "" {
		imageFilter.Name = in.Image
	}
	if in.ImageUUID != "" {
		imageFilter.ID = in.ImageUUID
	}
	out.Image = imageFilter

	if in.IdentityRef != nil {
		out.IdentityRef = &infrav1.OpenStackIdentityReference{Name: in.IdentityRef.Name}
	}
	if in.CloudName != "" {
		if out.IdentityRef == nil {
			out.IdentityRef = &infrav1.OpenStackIdentityReference{}
		}
		out.IdentityRef.CloudName = in.CloudName
	}

	return nil
}

func Convert_v1beta1_APIServerLoadBalancer_To_v1alpha5_APIServerLoadBalancer(in *infrav1.APIServerLoadBalancer, out *APIServerLoadBalancer, s conversion.Scope) error {
	// Provider was originally added in v1beta1, but was backported to v1alpha6, but has no equivalent in v1alpha5
	return autoConvert_v1beta1_APIServerLoadBalancer_To_v1alpha5_APIServerLoadBalancer(in, out, s)
}

func Convert_v1alpha5_PortOpts_To_v1beta1_PortOpts(in *PortOpts, out *infrav1.PortOpts, s conversion.Scope) error {
	// SecurityGroups have been removed in v1beta1.
	err := autoConvert_v1alpha5_PortOpts_To_v1beta1_PortOpts(in, out, s)
	if err != nil {
		return err
	}

	if len(in.SecurityGroups) > 0 || len(in.SecurityGroupFilters) > 0 {
		out.SecurityGroups = make([]infrav1.SecurityGroupFilter, 0, len(in.SecurityGroups)+len(in.SecurityGroupFilters))
		for i := range in.SecurityGroupFilters {
			sgParam := &in.SecurityGroupFilters[i]
			switch {
			case sgParam.UUID != "":
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{ID: sgParam.UUID})
			case sgParam.Name != "":
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{Name: sgParam.Name})
			case sgParam.Filter != (SecurityGroupFilter{}):
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{})
				outSG := &out.SecurityGroups[len(out.SecurityGroups)-1]
				if err := Convert_v1alpha5_SecurityGroupFilter_To_v1beta1_SecurityGroupFilter(&sgParam.Filter, outSG, s); err != nil {
					return err
				}
			}
		}
		for _, id := range in.SecurityGroups {
			out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{ID: id})
		}
	}

	if len(in.SecurityGroups) > 0 || len(in.SecurityGroupFilters) > 0 {
		out.SecurityGroups = make([]infrav1.SecurityGroupFilter, 0, len(in.SecurityGroups)+len(in.SecurityGroupFilters))
		for i := range in.SecurityGroupFilters {
			sgParam := &in.SecurityGroupFilters[i]
			switch {
			case sgParam.UUID != "":
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{ID: sgParam.UUID})
			case sgParam.Name != "":
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{Name: sgParam.Name})
			case sgParam.Filter != (SecurityGroupFilter{}):
				out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{})
				outSG := &out.SecurityGroups[len(out.SecurityGroups)-1]
				if err := Convert_v1alpha5_SecurityGroupFilter_To_v1beta1_SecurityGroupFilter(&sgParam.Filter, outSG, s); err != nil {
					return err
				}
			}
		}
		for _, id := range in.SecurityGroups {
			out.SecurityGroups = append(out.SecurityGroups, infrav1.SecurityGroupFilter{ID: id})
		}
	}

	// Profile is now a struct in v1beta1.
	if strings.Contains(in.Profile["capabilities"], "switchdev") {
		out.Profile.OVSHWOffload = pointer.Bool(true)
	}
	if in.Profile["trusted"] == trueString {
		out.Profile.TrustedVF = pointer.Bool(true)
	}
	return nil
}

func Convert_v1alpha5_Instance_To_v1beta1_BastionStatus(in *Instance, out *infrav1.BastionStatus, _ conversion.Scope) error {
	// BastionStatus is the same as Instance with unused fields removed
	out.ID = in.ID
	out.Name = in.Name
	out.SSHKeyName = in.SSHKeyName
	out.State = infrav1.InstanceState(in.State)
	out.IP = in.IP
	out.FloatingIP = in.FloatingIP
	out.ReferencedResources.ServerGroupID = in.ServerGroupID
	return nil
}

func Convert_v1beta1_BastionStatus_To_v1alpha5_Instance(in *infrav1.BastionStatus, out *Instance, _ conversion.Scope) error {
	// BastionStatus is the same as Instance with unused fields removed
	out.ID = in.ID
	out.Name = in.Name
	out.SSHKeyName = in.SSHKeyName
	out.State = InstanceState(in.State)
	out.IP = in.IP
	out.FloatingIP = in.FloatingIP
	out.ServerGroupID = in.ReferencedResources.ServerGroupID
	return nil
}

func Convert_v1alpha5_Network_To_v1beta1_NetworkStatusWithSubnets(in *Network, out *infrav1.NetworkStatusWithSubnets, s conversion.Scope) error {
	// PortOpts has been removed in v1beta1
	err := Convert_v1alpha5_Network_To_v1beta1_NetworkStatus(in, &out.NetworkStatus, s)
	if err != nil {
		return err
	}

	if in.Subnet != nil {
		out.Subnets = []infrav1.Subnet{infrav1.Subnet(*in.Subnet)}
	}
	return nil
}

func Convert_v1beta1_NetworkStatusWithSubnets_To_v1alpha5_Network(in *infrav1.NetworkStatusWithSubnets, out *Network, s conversion.Scope) error {
	// PortOpts has been removed in v1beta1
	err := Convert_v1beta1_NetworkStatus_To_v1alpha5_Network(&in.NetworkStatus, out, s)
	if err != nil {
		return err
	}

	// Can only down-convert a single subnet
	if len(in.Subnets) > 0 {
		out.Subnet = (*Subnet)(&in.Subnets[0])
	}
	return nil
}

func Convert_v1alpha5_Network_To_v1beta1_NetworkStatus(in *Network, out *infrav1.NetworkStatus, _ conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	out.Tags = in.Tags

	return nil
}

func Convert_v1beta1_NetworkStatus_To_v1alpha5_Network(in *infrav1.NetworkStatus, out *Network, _ conversion.Scope) error {
	out.ID = in.ID
	out.Name = in.Name
	out.Tags = in.Tags

	return nil
}

func Convert_v1alpha5_SecurityGroupParam_To_v1beta1_SecurityGroupFilter(in *SecurityGroupParam, out *infrav1.SecurityGroupFilter, s conversion.Scope) error {
	// SecurityGroupParam is replaced by its contained SecurityGroupFilter in v1beta1
	err := Convert_v1alpha5_SecurityGroupFilter_To_v1beta1_SecurityGroupFilter(&in.Filter, out, s)
	if err != nil {
		return err
	}

	if in.UUID != "" {
		out.ID = in.UUID
	}
	if in.Name != "" {
		out.Name = in.Name
	}
	return nil
}

func Convert_v1beta1_SecurityGroupFilter_To_v1alpha5_SecurityGroupParam(in *infrav1.SecurityGroupFilter, out *SecurityGroupParam, s conversion.Scope) error {
	// SecurityGroupParam is replaced by its contained SecurityGroupFilter in v1beta1
	err := Convert_v1beta1_SecurityGroupFilter_To_v1alpha5_SecurityGroupFilter(in, &out.Filter, s)
	if err != nil {
		return err
	}

	if in.ID != "" {
		out.UUID = in.ID
	}
	if in.Name != "" {
		out.Name = in.Name
	}
	return nil
}

func Convert_v1alpha5_SubnetParam_To_v1beta1_SubnetFilter(in *SubnetParam, out *infrav1.SubnetFilter, s conversion.Scope) error {
	if err := Convert_v1alpha5_SubnetFilter_To_v1beta1_SubnetFilter(&in.Filter, out, s); err != nil {
		return err
	}
	if in.UUID != "" {
		out.ID = in.UUID
	}
	return nil
}

func Convert_v1beta1_SubnetFilter_To_v1alpha5_SubnetParam(in *infrav1.SubnetFilter, out *SubnetParam, s conversion.Scope) error {
	if err := Convert_v1beta1_SubnetFilter_To_v1alpha5_SubnetFilter(in, &out.Filter, s); err != nil {
		return err
	}
	out.UUID = in.ID

	return nil
}

func Convert_Map_string_To_Interface_To_v1beta1_BindingProfile(in map[string]string, out *infrav1.BindingProfile, _ conversion.Scope) error {
	for k, v := range in {
		if k == "capabilities" {
			if strings.Contains(v, "switchdev") {
				out.OVSHWOffload = pointer.Bool(true)
			}
		}
		if k == "trusted" && v == trueString {
			out.TrustedVF = pointer.Bool(true)
		}
	}
	return nil
}

func Convert_v1beta1_BindingProfile_To_Map_string_To_Interface(in *infrav1.BindingProfile, out map[string]string, _ conversion.Scope) error {
	if pointer.BoolDeref(in.OVSHWOffload, false) {
		(out)["capabilities"] = "[\"switchdev\"]"
	}
	if pointer.BoolDeref(in.TrustedVF, false) {
		(out)["trusted"] = trueString
	}
	return nil
}

func Convert_v1beta1_OpenStackClusterStatus_To_v1alpha5_OpenStackClusterStatus(in *infrav1.OpenStackClusterStatus, out *OpenStackClusterStatus, s conversion.Scope) error {
	err := autoConvert_v1beta1_OpenStackClusterStatus_To_v1alpha5_OpenStackClusterStatus(in, out, s)
	if err != nil {
		return err
	}

	// Router and APIServerLoadBalancer have been moved out of Network in v1beta1
	if in.Router != nil || in.APIServerLoadBalancer != nil {
		if out.Network == nil {
			out.Network = &Network{}
		}

		out.Network.Router = (*Router)(in.Router)
		if in.APIServerLoadBalancer != nil {
			out.Network.APIServerLoadBalancer = &LoadBalancer{}
			err = Convert_v1beta1_LoadBalancer_To_v1alpha5_LoadBalancer(in.APIServerLoadBalancer, out.Network.APIServerLoadBalancer, s)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Convert_v1alpha5_OpenStackClusterStatus_To_v1beta1_OpenStackClusterStatus(in *OpenStackClusterStatus, out *infrav1.OpenStackClusterStatus, s conversion.Scope) error {
	err := autoConvert_v1alpha5_OpenStackClusterStatus_To_v1beta1_OpenStackClusterStatus(in, out, s)
	if err != nil {
		return err
	}

	// Router and APIServerLoadBalancer have been moved out of Network in v1beta1
	if in.Network != nil {
		out.Router = (*infrav1.Router)(in.Network.Router)
		if in.Network.APIServerLoadBalancer != nil {
			out.APIServerLoadBalancer = &infrav1.LoadBalancer{}
			err = Convert_v1alpha5_LoadBalancer_To_v1beta1_LoadBalancer(in.Network.APIServerLoadBalancer, out.APIServerLoadBalancer, s)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func Convert_v1beta1_OpenStackMachineSpec_To_v1alpha5_OpenStackMachineSpec(in *infrav1.OpenStackMachineSpec, out *OpenStackMachineSpec, s conversion.Scope) error {
	err := autoConvert_v1beta1_OpenStackMachineSpec_To_v1alpha5_OpenStackMachineSpec(in, out, s)
	if err != nil {
		return err
	}

	if in.ServerGroup != nil {
		out.ServerGroupID = in.ServerGroup.ID
	}

	if in.Image.Name != "" {
		out.Image = in.Image.Name
	}

	if in.Image.ID != "" {
		out.ImageUUID = in.Image.ID
	}

	if in.IdentityRef != nil {
		out.IdentityRef = &OpenStackIdentityReference{Name: in.IdentityRef.Name}
		out.CloudName = in.IdentityRef.CloudName
	}

	return nil
}

func Convert_v1beta1_OpenStackMachineStatus_To_v1alpha5_OpenStackMachineStatus(in *infrav1.OpenStackMachineStatus, out *OpenStackMachineStatus, s conversion.Scope) error {
	// ReferencedResources have no equivalent in v1alpha5
	return autoConvert_v1beta1_OpenStackMachineStatus_To_v1alpha5_OpenStackMachineStatus(in, out, s)
}

func Convert_v1beta1_Bastion_To_v1alpha5_Bastion(in *infrav1.Bastion, out *Bastion, s conversion.Scope) error {
	err := autoConvert_v1beta1_Bastion_To_v1alpha5_Bastion(in, out, s)
	if err != nil {
		return err
	}
	in.FloatingIP = out.Instance.FloatingIP
	return nil
}

func Convert_v1alpha5_Bastion_To_v1beta1_Bastion(in *Bastion, out *infrav1.Bastion, s conversion.Scope) error {
	err := autoConvert_v1alpha5_Bastion_To_v1beta1_Bastion(in, out, s)
	if err != nil {
		return err
	}
	in.Instance.FloatingIP = out.FloatingIP
	return nil
}

func Convert_v1beta1_SecurityGroupStatus_To_v1alpha5_SecurityGroup(in *infrav1.SecurityGroupStatus, out *SecurityGroup, s conversion.Scope) error { //nolint:revive
	out.ID = in.ID
	out.Name = in.Name
	out.Rules = make([]SecurityGroupRule, len(in.Rules))
	for i, rule := range in.Rules {
		out.Rules[i] = SecurityGroupRule{
			ID:        rule.ID,
			Direction: rule.Direction,
		}
		if rule.Description != nil {
			out.Rules[i].Description = *rule.Description
		}
		if rule.EtherType != nil {
			out.Rules[i].EtherType = *rule.EtherType
		}
		if rule.PortRangeMin != nil {
			out.Rules[i].PortRangeMin = *rule.PortRangeMin
		}
		if rule.PortRangeMax != nil {
			out.Rules[i].PortRangeMax = *rule.PortRangeMax
		}
		if rule.Protocol != nil {
			out.Rules[i].Protocol = *rule.Protocol
		}
		if rule.RemoteGroupID != nil {
			out.Rules[i].RemoteGroupID = *rule.RemoteGroupID
		}
		if rule.RemoteIPPrefix != nil {
			out.Rules[i].RemoteIPPrefix = *rule.RemoteIPPrefix
		}
	}
	return nil
}

func Convert_v1alpha5_SecurityGroup_To_v1beta1_SecurityGroupStatus(in *SecurityGroup, out *infrav1.SecurityGroupStatus, s conversion.Scope) error { //nolint:revive
	out.ID = in.ID
	out.Name = in.Name
	out.Rules = make([]infrav1.SecurityGroupRuleStatus, len(in.Rules))
	for i, rule := range in.Rules {
		out.Rules[i] = infrav1.SecurityGroupRuleStatus{
			ID:             rule.ID,
			Description:    pointer.String(rule.Description),
			Direction:      rule.Direction,
			EtherType:      pointer.String(rule.EtherType),
			PortRangeMin:   pointer.Int(rule.PortRangeMin),
			PortRangeMax:   pointer.Int(rule.PortRangeMax),
			Protocol:       pointer.String(rule.Protocol),
			RemoteGroupID:  pointer.String(rule.RemoteGroupID),
			RemoteIPPrefix: pointer.String(rule.RemoteIPPrefix),
		}
	}

	return nil
}

func Convert_v1alpha5_OpenStackIdentityReference_To_v1beta1_OpenStackIdentityReference(in *OpenStackIdentityReference, out *infrav1.OpenStackIdentityReference, s conversion.Scope) error {
	return autoConvert_v1alpha5_OpenStackIdentityReference_To_v1beta1_OpenStackIdentityReference(in, out, s)
}

func Convert_v1beta1_OpenStackIdentityReference_To_v1alpha5_OpenStackIdentityReference(in *infrav1.OpenStackIdentityReference, out *OpenStackIdentityReference, _ conversion.Scope) error {
	out.Name = in.Name
	return nil
}

func Convert_v1alpha5_SubnetFilter_To_v1beta1_SubnetFilter(in *SubnetFilter, out *infrav1.SubnetFilter, s conversion.Scope) error {
	if err := autoConvert_v1alpha5_SubnetFilter_To_v1beta1_SubnetFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsTo(in.Tags, in.TagsAny, in.NotTags, in.NotTagsAny, &out.FilterByNeutronTags)
	return nil
}

func Convert_v1beta1_SubnetFilter_To_v1alpha5_SubnetFilter(in *infrav1.SubnetFilter, out *SubnetFilter, s conversion.Scope) error {
	if err := autoConvert_v1beta1_SubnetFilter_To_v1alpha5_SubnetFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsFrom(&in.FilterByNeutronTags, &out.Tags, &out.TagsAny, &out.NotTags, &out.NotTagsAny)
	return nil
}

func Convert_v1alpha5_SecurityGroupFilter_To_v1beta1_SecurityGroupFilter(in *SecurityGroupFilter, out *infrav1.SecurityGroupFilter, s conversion.Scope) error {
	if err := autoConvert_v1alpha5_SecurityGroupFilter_To_v1beta1_SecurityGroupFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsTo(in.Tags, in.TagsAny, in.NotTags, in.NotTagsAny, &out.FilterByNeutronTags)

	// TenantID has been removed in v1beta1. Write it to ProjectID if ProjectID is not already set.
	if out.ProjectID == "" {
		out.ProjectID = in.TenantID
	}
	return nil
}

func Convert_v1beta1_SecurityGroupFilter_To_v1alpha5_SecurityGroupFilter(in *infrav1.SecurityGroupFilter, out *SecurityGroupFilter, s conversion.Scope) error {
	if err := autoConvert_v1beta1_SecurityGroupFilter_To_v1alpha5_SecurityGroupFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsFrom(&in.FilterByNeutronTags, &out.Tags, &out.TagsAny, &out.NotTags, &out.NotTagsAny)
	return nil
}

func Convert_v1alpha5_NetworkFilter_To_v1beta1_NetworkFilter(in *NetworkFilter, out *infrav1.NetworkFilter, s conversion.Scope) error {
	if err := autoConvert_v1alpha5_NetworkFilter_To_v1beta1_NetworkFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsTo(in.Tags, in.TagsAny, in.NotTags, in.NotTagsAny, &out.FilterByNeutronTags)
	return nil
}

func Convert_v1beta1_NetworkFilter_To_v1alpha5_NetworkFilter(in *infrav1.NetworkFilter, out *NetworkFilter, s conversion.Scope) error {
	if err := autoConvert_v1beta1_NetworkFilter_To_v1alpha5_NetworkFilter(in, out, s); err != nil {
		return err
	}
	infrav1.ConvertAllTagsFrom(&in.FilterByNeutronTags, &out.Tags, &out.TagsAny, &out.NotTags, &out.NotTagsAny)
	return nil
}
