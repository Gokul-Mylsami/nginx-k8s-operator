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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NginxRoutesSpec defines the desired state of NginxRoutes
type NginxRoutesSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ServerName       string            `json:"serverName,omitempty"`
	ServerPort       int32             `json:"serverPort"`
	TemplateFile     string            `json:"templateFile"`
	TLSCertificate   TLSCertificate    `json:"tlsCertificate,omitempty"`
	CustomLocations  []CustomLocations `json:"customLocations,omitempty"`
	CustomDirectives []string          `json:"customDirectives,omitempty"`
}

// NginxRoutesStatus defines the observed state of NginxRoutes
type NginxRoutesStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

type TLSCertificate struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type CustomLocations struct {
	Location   string `json:"location"`
	Definition string `json:"definition"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// NginxRoutes is the Schema for the nginxroutes API
type NginxRoutes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NginxRoutesSpec   `json:"spec,omitempty"`
	Status NginxRoutesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NginxRoutesList contains a list of NginxRoutes
type NginxRoutesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NginxRoutes `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NginxRoutes{}, &NginxRoutesList{})
}
