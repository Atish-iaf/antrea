// +build !ignore_autogenerated

// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destination) DeepCopyInto(out *Destination) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destination.
func (in *Destination) DeepCopy() *Destination {
	if in == nil {
		return nil
	}
	out := new(Destination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPEchoRequestHeader) DeepCopyInto(out *ICMPEchoRequestHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPEchoRequestHeader.
func (in *ICMPEchoRequestHeader) DeepCopy() *ICMPEchoRequestHeader {
	if in == nil {
		return nil
	}
	out := new(ICMPEchoRequestHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPHeader) DeepCopyInto(out *IPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPHeader.
func (in *IPHeader) DeepCopy() *IPHeader {
	if in == nil {
		return nil
	}
	out := new(IPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPv6Header) DeepCopyInto(out *IPv6Header) {
	*out = *in
	if in.NextHeader != nil {
		in, out := &in.NextHeader, &out.NextHeader
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPv6Header.
func (in *IPv6Header) DeepCopy() *IPv6Header {
	if in == nil {
		return nil
	}
	out := new(IPv6Header)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeResult) DeepCopyInto(out *NodeResult) {
	*out = *in
	if in.Observations != nil {
		in, out := &in.Observations, &out.Observations
		*out = make([]Observation, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeResult.
func (in *NodeResult) DeepCopy() *NodeResult {
	if in == nil {
		return nil
	}
	out := new(NodeResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observation) DeepCopyInto(out *Observation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observation.
func (in *Observation) DeepCopy() *Observation {
	if in == nil {
		return nil
	}
	out := new(Observation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Packet) DeepCopyInto(out *Packet) {
	*out = *in
	out.IPHeader = in.IPHeader
	if in.IPv6Header != nil {
		in, out := &in.IPv6Header, &out.IPv6Header
		*out = new(IPv6Header)
		(*in).DeepCopyInto(*out)
	}
	in.TransportHeader.DeepCopyInto(&out.TransportHeader)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Packet.
func (in *Packet) DeepCopy() *Packet {
	if in == nil {
		return nil
	}
	out := new(Packet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Source) DeepCopyInto(out *Source) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Source.
func (in *Source) DeepCopy() *Source {
	if in == nil {
		return nil
	}
	out := new(Source)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPHeader) DeepCopyInto(out *TCPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPHeader.
func (in *TCPHeader) DeepCopy() *TCPHeader {
	if in == nil {
		return nil
	}
	out := new(TCPHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Traceflow) DeepCopyInto(out *Traceflow) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Traceflow.
func (in *Traceflow) DeepCopy() *Traceflow {
	if in == nil {
		return nil
	}
	out := new(Traceflow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Traceflow) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowList) DeepCopyInto(out *TraceflowList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Traceflow, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowList.
func (in *TraceflowList) DeepCopy() *TraceflowList {
	if in == nil {
		return nil
	}
	out := new(TraceflowList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TraceflowList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowSpec) DeepCopyInto(out *TraceflowSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	in.Packet.DeepCopyInto(&out.Packet)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowSpec.
func (in *TraceflowSpec) DeepCopy() *TraceflowSpec {
	if in == nil {
		return nil
	}
	out := new(TraceflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TraceflowStatus) DeepCopyInto(out *TraceflowStatus) {
	*out = *in
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]NodeResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TraceflowStatus.
func (in *TraceflowStatus) DeepCopy() *TraceflowStatus {
	if in == nil {
		return nil
	}
	out := new(TraceflowStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransportHeader) DeepCopyInto(out *TransportHeader) {
	*out = *in
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = new(ICMPEchoRequestHeader)
		**out = **in
	}
	if in.UDP != nil {
		in, out := &in.UDP, &out.UDP
		*out = new(UDPHeader)
		**out = **in
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = new(TCPHeader)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransportHeader.
func (in *TransportHeader) DeepCopy() *TransportHeader {
	if in == nil {
		return nil
	}
	out := new(TransportHeader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UDPHeader) DeepCopyInto(out *UDPHeader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UDPHeader.
func (in *UDPHeader) DeepCopy() *UDPHeader {
	if in == nil {
		return nil
	}
	out := new(UDPHeader)
	in.DeepCopyInto(out)
	return out
}