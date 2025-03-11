/*
Copyright 2025.

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

package v1alpha1

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Gateway struct {
	Name string `json:"name"`
}

type Destination struct {
	Host string `json:"host"`
}

type Route struct {
	Destination Destination `json:"destination"`
	Weight      int         `json:"weight"`
}

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type VirtualService struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

type Step struct {
	SetWeight *int32 `json:"setWeight,omitempty"`
	Pause     *Pause `json:"pause,omitempty"`
}

type Pause struct {
	Duration string `json:"duration"`
}

type Strategy struct {
	Canary CanaryStrategy `json:"canary"`
}

type Istio struct {
	Gateways       []Gateway      `json:"gateways"`
	VirtualService VirtualService `json:"virtualService"`
}

// AristiSpec defines the desired state of Aristi
type AristiSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Aristi. Edit aristi_types.go to remove/update
	Istio   Istio       `json:"istio"`
	Rollout RolloutSpec `json:"rollout"`
}

type RolloutSpec struct {
	// Template describes the pods that will be created.
	// +optional
	Template PodTemplateSpec `json:"template" protobuf:"bytes,2,opt,name=template"`
	Replicas *int32          `json:"replicas,omitempty" protobuf:"varint,1,opt,name=replicas"`
	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this rollout.
	// It must match the pod template's labels.
	// +optional
	Selector *metav1.LabelSelector `json:"selector" protobuf:"bytes,2,opt,name=selector"`
	Strategy RolloutStrategy       `json:"strategy" protobuf:"bytes,5,opt,name=strategy"`
}

// RolloutStrategy defines strategy to apply during next rollout
type RolloutStrategy struct {
	// +optional
	Canary *CanaryStrategy `json:"canary,omitempty" protobuf:"bytes,2,opt,name=canary"`
}

// CanaryStrategy defines parameters for a Replica Based Canary
type CanaryStrategy struct {

	// CanaryService holds the name of a service which selects pods with canary version and don't select any pods with stable version.
	// +optional
	CanaryService string `json:"canaryService,omitempty" protobuf:"bytes,1,opt,name=canaryService"`
	// StableService holds the name of a service which selects pods with stable version and don't select any pods with canary version.
	// +optional
	StableService string `json:"stableService,omitempty" protobuf:"bytes,2,opt,name=stableService"`
	// Steps define the order of phases to execute the canary deployment
	// +optional
	Steps []CanaryStep `json:"steps,omitempty" protobuf:"bytes,3,rep,name=steps"`
	// TrafficRouting hosts all the supported service meshes supported to enable more fine-grained traffic routing
	TrafficRouting *RolloutTrafficRouting `json:"trafficRouting,omitempty" protobuf:"bytes,4,opt,name=trafficRouting"`
}

type CanaryStep struct {
	// SetWeight sets what percentage of the newRS should receive
	SetWeight *int32 `json:"setWeight,omitempty" protobuf:"varint,1,opt,name=setWeight"`
	// Pause freezes the rollout by setting spec.Paused to true.
	// A Rollout will resume when spec.Paused is reset to false.
	// +optional
	Pause *RolloutPause `json:"pause,omitempty" protobuf:"bytes,2,opt,name=pause"`
}

// RolloutPause defines a pause stage for a rollout
type RolloutPause struct {
	// Duration the amount of time to wait before moving to the next step.
	// +optional
	Duration *intstr.IntOrString `json:"duration,omitempty" protobuf:"bytes,1,opt,name=duration"`
}

type RolloutTrafficRouting struct {
	// Istio holds Istio specific configuration to route traffic
	Istio *IstioTrafficRouting `json:"istio,omitempty" protobuf:"bytes,1,opt,name=istio"`
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Type=object
	// Plugins holds specific configuration that traffic router plugins can use for routing traffic
	Plugins map[string]json.RawMessage `json:"plugins,omitempty" protobuf:"bytes,10,opt,name=plugins"`
}

// IstioTrafficRouting configuration for Istio service mesh to enable fine grain configuration
type IstioTrafficRouting struct {
	// VirtualService references an Istio VirtualService to modify to shape traffic
	VirtualService *IstioVirtualService `json:"virtualService,omitempty" protobuf:"bytes,1,opt,name=virtualService"`
	// DestinationRule references an Istio DestinationRule to modify to shape traffic
	DestinationRule *IstioDestinationRule `json:"destinationRule,omitempty" protobuf:"bytes,2,opt,name=destinationRule"`
	// VirtualServices references a list of Istio VirtualService to modify to shape traffic
	VirtualServices []IstioVirtualService `json:"virtualServices,omitempty" protobuf:"bytes,3,opt,name=virtualServices"`
}

// IstioDestinationRule is a reference to an Istio DestinationRule to modify and shape traffic
type IstioDestinationRule struct {
	// Name holds the name of the DestinationRule
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// CanarySubsetName is the subset name to modify labels with canary ReplicaSet pod template hash value
	CanarySubsetName string `json:"canarySubsetName" protobuf:"bytes,2,opt,name=canarySubsetName"`
	// StableSubsetName is the subset name to modify labels with stable ReplicaSet pod template hash value
	StableSubsetName string `json:"stableSubsetName" protobuf:"bytes,3,opt,name=stableSubsetName"`
}

// IstioVirtualService holds information on the virtual service the rollout needs to modify
type IstioVirtualService struct {
	// Name holds the name of the VirtualService
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// A list of HTTP routes within VirtualService to edit. If omitted, VirtualService must have a single route of this type.
	Routes []string `json:"routes,omitempty" protobuf:"bytes,2,rep,name=routes"`
	// A list of TLS/HTTPS routes within VirtualService to edit. If omitted, VirtualService must have a single route of this type.
	TLSRoutes []TLSRoute `json:"tlsRoutes,omitempty" protobuf:"bytes,3,rep,name=tlsRoutes"`
	// A list of TCP routes within VirtualService to edit. If omitted, VirtualService must have a single route of this type.
	TCPRoutes []TCPRoute `json:"tcpRoutes,omitempty" protobuf:"bytes,4,rep,name=tcpRoutes"`
}

// TLSRoute holds the information on the virtual service's TLS/HTTPS routes that are desired to be matched for changing weights.
type TLSRoute struct {
	// Port number of the TLS Route desired to be matched in the given Istio VirtualService.
	Port int64 `json:"port,omitempty" protobuf:"bytes,1,opt,name=port"`
	// A list of all the SNI Hosts of the TLS Route desired to be matched in the given Istio VirtualService.
	SNIHosts []string `json:"sniHosts,omitempty" protobuf:"bytes,2,rep,name=sniHosts"`
}

// TCPRoute holds the information on the virtual service's TCP routes that are desired to be matched for changing weights.
type TCPRoute struct {
	// Port number of the TCP Route desired to be matched in the given Istio VirtualService.
	Port int64 `json:"port,omitempty" protobuf:"bytes,1,opt,name=port"`
}

type PodTemplateSpec struct {
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the pod.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec PodSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// PodSpec is a description of a pod.
type PodSpec struct {

	// List of containers belonging to the pod.
	// Containers cannot currently be added or removed.
	// There must be at least one container in a Pod.
	// Cannot be updated.
	// +patchMergeKey=name
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=name
	Containers []Container `json:"containers" patchStrategy:"merge" patchMergeKey:"name" protobuf:"bytes,2,rep,name=containers"`
}

type ObjectMeta struct {
	// Name must be unique within a namespace. Is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	// +optional
	Name string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name"`

	// Namespace defines the space within which each name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces
	// +optional
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`
}

// AristiStatus defines the observed state of Aristi
type AristiStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Aristi is the Schema for the aristis API
type Aristi struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AristiSpec   `json:"spec,omitempty"`
	Status AristiStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AristiList contains a list of Aristi
type AristiList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aristi `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Aristi{}, &AristiList{})
}
