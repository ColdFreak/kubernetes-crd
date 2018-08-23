// make the new types known to the client library.
// This will allow the client to (more or less) automatically
// process your new types when communicating with the API server.

package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const GroupName = "kubernetes.idcf.jp"
const GroupVersion = "v1alpha1"

var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}

var (
	// https://godoc.org/k8s.io/apimachinery/pkg/runtime#NewSchemeBuilder
	// func NewSchemeBuilder(funcs ...func(*Scheme) error) SchemeBuilder
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	// func (sb *SchemeBuilder) AddToScheme(s *Scheme) error
	// AddToScheme applies all the stored functions to the scheme.
	// A non-nil error indicates that one function failed and the attempt was abandoned.
	AddToScheme = SchemeBuilder.AddToScheme
)

func addKnownTypes(scheme *runtime.Scheme) error {
	// https://godoc.org/k8s.io/apimachinery/pkg/runtime#Scheme.AddKnownTypes
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Project{},
		&ProjectList{},
	)
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
