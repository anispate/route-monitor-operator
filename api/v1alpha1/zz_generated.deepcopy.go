//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUrlMonitor) DeepCopyInto(out *ClusterUrlMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUrlMonitor.
func (in *ClusterUrlMonitor) DeepCopy() *ClusterUrlMonitor {
	if in == nil {
		return nil
	}
	out := new(ClusterUrlMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUrlMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUrlMonitorList) DeepCopyInto(out *ClusterUrlMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterUrlMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUrlMonitorList.
func (in *ClusterUrlMonitorList) DeepCopy() *ClusterUrlMonitorList {
	if in == nil {
		return nil
	}
	out := new(ClusterUrlMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUrlMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUrlMonitorSpec) DeepCopyInto(out *ClusterUrlMonitorSpec) {
	*out = *in
	out.Slo = in.Slo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUrlMonitorSpec.
func (in *ClusterUrlMonitorSpec) DeepCopy() *ClusterUrlMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterUrlMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUrlMonitorStatus) DeepCopyInto(out *ClusterUrlMonitorStatus) {
	*out = *in
	out.ServiceMonitorRef = in.ServiceMonitorRef
	out.PrometheusRuleRef = in.PrometheusRuleRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUrlMonitorStatus.
func (in *ClusterUrlMonitorStatus) DeepCopy() *ClusterUrlMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterUrlMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedName) DeepCopyInto(out *NamespacedName) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedName.
func (in *NamespacedName) DeepCopy() *NamespacedName {
	if in == nil {
		return nil
	}
	out := new(NamespacedName)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMonitor) DeepCopyInto(out *RouteMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMonitor.
func (in *RouteMonitor) DeepCopy() *RouteMonitor {
	if in == nil {
		return nil
	}
	out := new(RouteMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMonitorList) DeepCopyInto(out *RouteMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RouteMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMonitorList.
func (in *RouteMonitorList) DeepCopy() *RouteMonitorList {
	if in == nil {
		return nil
	}
	out := new(RouteMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RouteMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMonitorRouteSpec) DeepCopyInto(out *RouteMonitorRouteSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMonitorRouteSpec.
func (in *RouteMonitorRouteSpec) DeepCopy() *RouteMonitorRouteSpec {
	if in == nil {
		return nil
	}
	out := new(RouteMonitorRouteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMonitorSpec) DeepCopyInto(out *RouteMonitorSpec) {
	*out = *in
	out.Route = in.Route
	out.Slo = in.Slo
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMonitorSpec.
func (in *RouteMonitorSpec) DeepCopy() *RouteMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(RouteMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RouteMonitorStatus) DeepCopyInto(out *RouteMonitorStatus) {
	*out = *in
	out.ServiceMonitorRef = in.ServiceMonitorRef
	out.PrometheusRuleRef = in.PrometheusRuleRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RouteMonitorStatus.
func (in *RouteMonitorStatus) DeepCopy() *RouteMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(RouteMonitorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SloSpec) DeepCopyInto(out *SloSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SloSpec.
func (in *SloSpec) DeepCopy() *SloSpec {
	if in == nil {
		return nil
	}
	out := new(SloSpec)
	in.DeepCopyInto(out)
	return out
}
