// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1beta1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using WorkloadSelector within kubernetes types, where deepcopy-gen is used.
func (in *WorkloadSelector) DeepCopyInto(out *WorkloadSelector) {
	p := proto.Clone(in).(*WorkloadSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopy() *WorkloadSelector {
	if in == nil {
		return nil
	}
	out := new(WorkloadSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector. Required by controller-gen.
func (in *WorkloadSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PortSelector within kubernetes types, where deepcopy-gen is used.
func (in *PortSelector) DeepCopyInto(out *PortSelector) {
	p := proto.Clone(in).(*PortSelector)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopy() *PortSelector {
	if in == nil {
		return nil
	}
	out := new(PortSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PortSelector. Required by controller-gen.
func (in *PortSelector) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using PolicyTargetReference within kubernetes types, where deepcopy-gen is used.
func (in *PolicyTargetReference) DeepCopyInto(out *PolicyTargetReference) {
	p := proto.Clone(in).(*PolicyTargetReference)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyTargetReference. Required by controller-gen.
func (in *PolicyTargetReference) DeepCopy() *PolicyTargetReference {
	if in == nil {
		return nil
	}
	out := new(PolicyTargetReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new PolicyTargetReference. Required by controller-gen.
func (in *PolicyTargetReference) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
