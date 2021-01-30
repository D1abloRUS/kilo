// +build !ignore_autogenerated

// Copyright 2021 the Kilo authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
func (in *DNSOrIP) DeepCopyInto(out *DNSOrIP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSOrIP.
func (in *DNSOrIP) DeepCopy() *DNSOrIP {
	if in == nil {
		return nil
	}
	out := new(DNSOrIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Peer) DeepCopyInto(out *Peer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Peer.
func (in *Peer) DeepCopy() *Peer {
	if in == nil {
		return nil
	}
	out := new(Peer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Peer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerEndpoint) DeepCopyInto(out *PeerEndpoint) {
	*out = *in
	out.DNSOrIP = in.DNSOrIP
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerEndpoint.
func (in *PeerEndpoint) DeepCopy() *PeerEndpoint {
	if in == nil {
		return nil
	}
	out := new(PeerEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerList) DeepCopyInto(out *PeerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Peer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerList.
func (in *PeerList) DeepCopy() *PeerList {
	if in == nil {
		return nil
	}
	out := new(PeerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PeerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerSpec) DeepCopyInto(out *PeerSpec) {
	*out = *in
	if in.AllowedIPs != nil {
		in, out := &in.AllowedIPs, &out.AllowedIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(PeerEndpoint)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerSpec.
func (in *PeerSpec) DeepCopy() *PeerSpec {
	if in == nil {
		return nil
	}
	out := new(PeerSpec)
	in.DeepCopyInto(out)
	return out
}
