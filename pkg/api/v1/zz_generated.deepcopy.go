//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) MongoDB, Inc. 2020-present.

Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License. You may obtain
a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1/project"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasCluster) DeepCopyInto(out *AtlasCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasCluster.
func (in *AtlasCluster) DeepCopy() *AtlasCluster {
	if in == nil {
		return nil
	}
	out := new(AtlasCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasClusterList) DeepCopyInto(out *AtlasClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AtlasCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasClusterList.
func (in *AtlasClusterList) DeepCopy() *AtlasClusterList {
	if in == nil {
		return nil
	}
	out := new(AtlasClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasClusterSpec) DeepCopyInto(out *AtlasClusterSpec) {
	*out = *in
	out.Project = in.Project
	if in.ClusterSpec != nil {
		in, out := &in.ClusterSpec, &out.ClusterSpec
		*out = new(ClusterSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasClusterSpec.
func (in *AtlasClusterSpec) DeepCopy() *AtlasClusterSpec {
	if in == nil {
		return nil
	}
	out := new(AtlasClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasDatabaseUser) DeepCopyInto(out *AtlasDatabaseUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasDatabaseUser.
func (in *AtlasDatabaseUser) DeepCopy() *AtlasDatabaseUser {
	if in == nil {
		return nil
	}
	out := new(AtlasDatabaseUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasDatabaseUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasDatabaseUserList) DeepCopyInto(out *AtlasDatabaseUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AtlasDatabaseUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasDatabaseUserList.
func (in *AtlasDatabaseUserList) DeepCopy() *AtlasDatabaseUserList {
	if in == nil {
		return nil
	}
	out := new(AtlasDatabaseUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasDatabaseUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasDatabaseUserSpec) DeepCopyInto(out *AtlasDatabaseUserSpec) {
	*out = *in
	out.Project = in.Project
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LabelSpec, len(*in))
		copy(*out, *in)
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RoleSpec, len(*in))
		copy(*out, *in)
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ScopeSpec, len(*in))
		copy(*out, *in)
	}
	if in.PasswordSecret != nil {
		in, out := &in.PasswordSecret, &out.PasswordSecret
		*out = new(ResourceRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasDatabaseUserSpec.
func (in *AtlasDatabaseUserSpec) DeepCopy() *AtlasDatabaseUserSpec {
	if in == nil {
		return nil
	}
	out := new(AtlasDatabaseUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasProject) DeepCopyInto(out *AtlasProject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasProject.
func (in *AtlasProject) DeepCopy() *AtlasProject {
	if in == nil {
		return nil
	}
	out := new(AtlasProject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasProject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasProjectList) DeepCopyInto(out *AtlasProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AtlasProject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasProjectList.
func (in *AtlasProjectList) DeepCopy() *AtlasProjectList {
	if in == nil {
		return nil
	}
	out := new(AtlasProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AtlasProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AtlasProjectSpec) DeepCopyInto(out *AtlasProjectSpec) {
	*out = *in
	if in.ConnectionSecret != nil {
		in, out := &in.ConnectionSecret, &out.ConnectionSecret
		*out = new(ResourceRef)
		**out = **in
	}
	if in.ProjectIPAccessList != nil {
		in, out := &in.ProjectIPAccessList, &out.ProjectIPAccessList
		*out = make([]project.IPAccessList, len(*in))
		copy(*out, *in)
	}
	if in.PrivateEndpoints != nil {
		in, out := &in.PrivateEndpoints, &out.PrivateEndpoints
		*out = make([]project.PrivateEndpoint, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AtlasProjectSpec.
func (in *AtlasProjectSpec) DeepCopy() *AtlasProjectSpec {
	if in == nil {
		return nil
	}
	out := new(AtlasProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingSpec) DeepCopyInto(out *AutoScalingSpec) {
	*out = *in
	if in.AutoIndexingEnabled != nil {
		in, out := &in.AutoIndexingEnabled, &out.AutoIndexingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DiskGBEnabled != nil {
		in, out := &in.DiskGBEnabled, &out.DiskGBEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Compute != nil {
		in, out := &in.Compute, &out.Compute
		*out = new(ComputeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingSpec.
func (in *AutoScalingSpec) DeepCopy() *AutoScalingSpec {
	if in == nil {
		return nil
	}
	out := new(AutoScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorSpec) DeepCopyInto(out *BiConnectorSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorSpec.
func (in *BiConnectorSpec) DeepCopy() *BiConnectorSpec {
	if in == nil {
		return nil
	}
	out := new(BiConnectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BIConnector != nil {
		in, out := &in.BIConnector, &out.BIConnector
		*out = new(BiConnectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DiskSizeGB != nil {
		in, out := &in.DiskSizeGB, &out.DiskSizeGB
		*out = new(int)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]LabelSpec, len(*in))
		copy(*out, *in)
	}
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(int)
		**out = **in
	}
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.PitEnabled != nil {
		in, out := &in.PitEnabled, &out.PitEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProviderBackupEnabled != nil {
		in, out := &in.ProviderBackupEnabled, &out.ProviderBackupEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ProviderSettings != nil {
		in, out := &in.ProviderSettings, &out.ProviderSettings
		*out = new(ProviderSettingsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ReplicationSpecs != nil {
		in, out := &in.ReplicationSpecs, &out.ReplicationSpecs
		*out = make([]ReplicationSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeSpec) DeepCopyInto(out *ComputeSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ScaleDownEnabled != nil {
		in, out := &in.ScaleDownEnabled, &out.ScaleDownEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeSpec.
func (in *ComputeSpec) DeepCopy() *ComputeSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelSpec) DeepCopyInto(out *LabelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelSpec.
func (in *LabelSpec) DeepCopy() *LabelSpec {
	if in == nil {
		return nil
	}
	out := new(LabelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSettingsSpec) DeepCopyInto(out *ProviderSettingsSpec) {
	*out = *in
	if in.DiskIOPS != nil {
		in, out := &in.DiskIOPS, &out.DiskIOPS
		*out = new(int64)
		**out = **in
	}
	if in.EncryptEBSVolume != nil {
		in, out := &in.EncryptEBSVolume, &out.EncryptEBSVolume
		*out = new(bool)
		**out = **in
	}
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSettingsSpec.
func (in *ProviderSettingsSpec) DeepCopy() *ProviderSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionsConfig) DeepCopyInto(out *RegionsConfig) {
	*out = *in
	if in.AnalyticsNodes != nil {
		in, out := &in.AnalyticsNodes, &out.AnalyticsNodes
		*out = new(int64)
		**out = **in
	}
	if in.ElectableNodes != nil {
		in, out := &in.ElectableNodes, &out.ElectableNodes
		*out = new(int64)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(int64)
		**out = **in
	}
	if in.ReadOnlyNodes != nil {
		in, out := &in.ReadOnlyNodes, &out.ReadOnlyNodes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionsConfig.
func (in *RegionsConfig) DeepCopy() *RegionsConfig {
	if in == nil {
		return nil
	}
	out := new(RegionsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationSpec) DeepCopyInto(out *ReplicationSpec) {
	*out = *in
	if in.NumShards != nil {
		in, out := &in.NumShards, &out.NumShards
		*out = new(int64)
		**out = **in
	}
	if in.RegionsConfig != nil {
		in, out := &in.RegionsConfig, &out.RegionsConfig
		*out = make(map[string]RegionsConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationSpec.
func (in *ReplicationSpec) DeepCopy() *ReplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRefNamespaced) DeepCopyInto(out *ResourceRefNamespaced) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRefNamespaced.
func (in *ResourceRefNamespaced) DeepCopy() *ResourceRefNamespaced {
	if in == nil {
		return nil
	}
	out := new(ResourceRefNamespaced)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleSpec) DeepCopyInto(out *RoleSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleSpec.
func (in *RoleSpec) DeepCopy() *RoleSpec {
	if in == nil {
		return nil
	}
	out := new(RoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopeSpec) DeepCopyInto(out *ScopeSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopeSpec.
func (in *ScopeSpec) DeepCopy() *ScopeSpec {
	if in == nil {
		return nil
	}
	out := new(ScopeSpec)
	in.DeepCopyInto(out)
	return out
}
