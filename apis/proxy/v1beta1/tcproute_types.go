/*
Copyright 2021.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TCPRouteSpec defines the desired state of TCPRoute
type TCPRouteSpec struct {
	// IngressClassName is the name of the IngressClass cluster resource.
	// +optional
	IngressClassName *string `json:"ingressClassName,omitempty"`
	// +optional
	Streams []Stream `json:"streams,omitempty"`
}

type Stream struct {
	// +optional
	Port int32 `json:"port,omitempty"`
	// SecretName,需要和CRD在同一个namespace下
	// +optional
	TLS *TLS `json:"tls,omitempty"`
	// +optional
	ServiceName string `json:"serviceName,omitempty"`
	// +optional
	ServicePort int32 `json:"servicePort,omitempty"`
}

// TCPRouteStatus defines the observed state of TCPRoute
type TCPRouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
//+kubebuilder:printcolumn:name="IngressClassName",type="string",priority=0,JSONPath=".spec.ingressClassName",description="The IngressClassName"
//+kubebuilder:resource:scope=Namespaced,shortName=tr
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// TCPRoute is the Schema for the tcproutes API
type TCPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPRouteSpec   `json:"spec,omitempty"`
	Status TCPRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCPRouteList contains a list of TCPRoute
type TCPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPRoute{}, &TCPRouteList{})
}
