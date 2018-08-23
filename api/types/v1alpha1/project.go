package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ProjectSpec struct {
	Replicas int `json:"replicas"`
}

type Project struct {
	meta_v1.TypeMeta `json:",inline"`
	// The meta_v1.ObjectMeta type contains the typical metadata properties
	// that you can find in any Kubernetes resource (like for example, the name,
	// namespace, labels and annotations).
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectSpec `json:"spec"`
}

type ProjectList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata,omitempty"`

	Items []Project `json:"items"`
}
