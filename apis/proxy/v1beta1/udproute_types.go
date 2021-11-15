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

// UDPRouteSpec defines the desired state of UDPRoute
type UDPRouteSpec struct {
	// IngressClassName is the name of the IngressClass cluster resource.
	// +optional
	IngressClassName *string `json:"ingressClassName,omitempty"`
	// +optional
	Streams []Stream `json:"services,omitempty"`
}

// UDPRouteStatus defines the observed state of UDPRoute
type UDPRouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
//+kubebuilder:resource:scope=Namespaced,shortName=ur
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// UDPRoute is the Schema for the udproutes API
type UDPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UDPRouteSpec   `json:"spec,omitempty"`
	Status UDPRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// UDPRouteList contains a list of UDPRoute
type UDPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UDPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&UDPRoute{}, &UDPRouteList{})
}
