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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReservedSpec defines a set of reserved CPUs with an optional PowerProfile.
type ReservedSpec struct {
	Cores        []uint `json:"cores"`
	PowerProfile string `json:"powerProfile,omitempty"`
}

// PowerNodeConfigSpec defines the desired state of PowerNodeConfig.
type PowerNodeConfigSpec struct {
	// PowerProfile is the name of the PowerProfile to apply to the shared CPU pool.
	// +kubebuilder:validation:MinLength=1
	PowerProfile string `json:"powerProfile"`

	// NodeSelector specifies which nodes this PowerNodeConfig applies to.
	// If not specified, the config applies to all nodes running the Power Node Agent.
	// +optional
	NodeSelector NodeSelector `json:"nodeSelector,omitempty"`

	// ReservedCPUs defines the CPUs reserved by kubelet with optional per-group PowerProfiles.
	// If not specified, reserved CPUs remain in the default reserved pool with no power profile applied.
	// +optional
	ReservedCPUs []ReservedSpec `json:"reservedCPUs,omitempty"`
}

// +kubebuilder:object:root=true

// PowerNodeConfig is the Schema for the powernodeconfigs API.
// It configures shared and reserved CPU pools on nodes matching spec.nodeSelector.
type PowerNodeConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PowerNodeConfigSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

// PowerNodeConfigList contains a list of PowerNodeConfig.
type PowerNodeConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PowerNodeConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PowerNodeConfig{}, &PowerNodeConfigList{})
}
