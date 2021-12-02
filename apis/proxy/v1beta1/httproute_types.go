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

// HTTPRouteSpec defines the desired state of HTTPRoute
type HTTPRouteSpec struct {
	// IngressClassName is the name of the IngressClass cluster resource.
	// +optional
	IngressClassName *string `json:"ingressClassName,omitempty"`
	// A list of host rules used to configure the Ingress. If unspecified, or
	// no rule matches, all traffic is sent to the default backend.
	// +listType=atomic
	Routes []Route `json:"routes,omitempty"`
}

// HTTPRouteStatus defines the observed state of HTTPRoute
type HTTPRouteStatus struct {
	// +optional
	Hostname string `json:"hostname,omitempty"`
}

type ProtocolType string
type SecretName string

const (
	HTTPProtocolType ProtocolType = "HTTP"
	// Accepts HTTP/1.1 or HTTP/2 sessions over TLS.
	HTTPSProtocolType ProtocolType = "HTTPS"
	// Accepts GRPC packets.
	GRPCProtocolType ProtocolType = "GRPC"
)

type Route struct {
	// +optional
	Host string `json:"host,omitempty"`
	// Protocol specifies the network protocol this listener expects to receive.
	// +optional
	// default HTTP
	Protocol ProtocolType `json:"protocol,omitempty"`

	// SecretName,需要和CRD在同一个namespace下
	// +optional
	TLS *TLS `json:"tls,omitempty"`
	// +optional
	Rules []HTTPRouteRule `json:"rules,omitempty"`
	// +optional
	Proxy Proxy `json:"proxy,omitempty"`
	// UserCustomBalancer
	// 自定义nginx特殊配置,此处填写nginx配置
	// +optional
	Options map[string]string `json:"options,omitempty"`
}

// TLS defines TLS configuration for a VirtualServer.
type TLS struct {
	Secret string `json:"secret"`
}

type PathType string

const (
	PathExact                      PathType = "exact"
	PathPrefix                     PathType = "prefix"
	PathRegularExpression          PathType = "regex"
	PathTypeImplementationSpecific PathType = "ImplementationSpecific"
)

type HTTPRouteRule struct {
	// Path specifies a HTTP request path matcher. If this field is not
	// specified, a default prefix match on the "/" path is provided.
	Path     string    `json:"path"`
	PathType *PathType `json:"pathType"`
	// +optional
	Rewrite string `json:"rewrite,omitempty"`
	// +optional
	Proxy Proxy `json:"proxy,omitempty"`
	// +optional
	RateLimit RateLimit `json:"rateLimit,omitempty"`
	// UserCustomBalancer
	// 自定义nginx特殊配置,此处填写nginx配置
	// +optional
	Options map[string]string `json:"options,omitempty"`
	// service
	// +optional
	Backends []Backend `json:"backends,omitempty"`
	// 当default backend 没有配置默认404
	// +optional
	DefaultBackend *DefaultBackend `json:"defaultBackend,omitempty"`
}

type Proxy struct {
	// +optional
	ProxyConnectTimeout string `json:"connectTimeout,omitempty"`
	// +optional
	ProxyReadTimeout string `json:"readTimeout,omitempty"`
	// +optional
	ProxySendTimeout string `json:"sendTimeout,omitempty"`
	// +optional
	ProxyNextUpstream string `json:"nextUpstream,omitempty"`
	// +optional
	ProxyNextUpstreamTimeout string `json:"nextUpstreamTimeout,omitempty"`
	// +optional
	ProxyNextUpstreamTries int `json:"nextUpstreamTries,omitempty"`
	// +optional
	ProxyBuffering *bool `json:"buffering,omitempty"`
	// +optional
	ProxyBuffers *UpstreamBuffers `json:"buffers,omitempty"`
	// +optional
	ProxyBufferSize string `json:"bufferSize,omitempty"`
	// +optional
	ClientMaxBodySize string `json:"clientMaxBodySize,omitempty"`
}

// UpstreamBuffers defines Buffer Configuration for an Upstream.
type UpstreamBuffers struct {
	// +optional
	Number int `json:"number,omitempty"`
	// +optional
	Size string `json:"size,omitempty"`
}

type RateLimit struct {
	// +optional
	Connections int `json:"connections,omitempty"`
	// +optional
	RPM int `json:"rpm,omitempty"`
	// +optional
	RPS int `json:"rps,omitempty"`
}

type OperatorType string

const (
	OperatorExact             OperatorType = "exact"
	OperatorRegularExpression OperatorType = "regex"
)

type HTTPMatch struct {
	// header or cookie
	Type string `json:"type"`
	// 同组的and，不同组的or
	GroupId  int          `json:"groupId"`
	Key      string       `json:"key"`
	Operator OperatorType `json:"operator"`
	// Value is the value of HTTP Header to be matched.
	Value string `json:"value"`
}

type Backend struct {
	// Name is the referenced service. The service must exist in
	// the same namespace as the Ingress object.
	// +optional
	Name string `json:"name,omitempty"`

	// Port of the referenced service.
	// is required for a IngressServiceBackend.
	// +optional
	Port *int32 `json:"port,omitempty"`

	// 灰度
	// +optional canary weight
	Weight *int32 `json:"weight,omitempty"`

	// +optional
	Matches []HTTPMatch `json:"matches,omitempty"`

	// 会话保持
	// +optional 可以选择按照 IP hash 或 insert cookie 来达到会话保持的效果。如果省略则默认为 round-robin
	Strategy string `json:"strategy,omitempty"`

	// +optional
	ChangeCookieOnFailure bool `json:"changeCookieOnFailure,omitempty"` // whether to set a new cookie when request failed
	// +optional
	FailTimeOutSeconds *int `json:"failTimeOut,omitempty"`
	// +optional
	MaxFails *int `json:"maxFails,omitempty"`
	// +optional
	MaxConns *int `json:"maxConns,omitempty"`
	// +optional
	Keepalive *int `json:"keepalive,omitempty"`

	// +optional
	Options map[string]string `json:"options,omitempty"`
}

type DefaultBackend struct {
	// 与response code 冲突
	// +optional
	Service *DefaultService `json:"service,omitempty"`
	// 自定义错误code,
	// +optional
	ResponseCode *int32 `json:"responseCode,omitempty"`
}

type DefaultService struct {
	// Name is the referenced service. The service must exist in
	// the same namespace as the Ingress object.
	// +optional
	Name string `json:"name,omitempty"`
	// Port of the referenced service.
	// is required for a IngressServiceBackend.
	// +optional
	Port *int32 `json:"port,omitempty"`
}

// +genclient
//+kubebuilder:printcolumn:name="IngressClassName",type="string",priority=0,JSONPath=".spec.ingressClassName",description="The IngressClassName"
//+kubebuilder:resource:scope=Namespaced,shortName=hr
//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HTTPRoute is the Schema for the httproutes API
type HTTPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPRouteSpec   `json:"spec,omitempty"`
	Status HTTPRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HTTPRouteList contains a list of HTTPRoute
type HTTPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HTTPRoute{}, &HTTPRouteList{})
}
