package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

func (in *Project) DeepCopyInto(out *Project) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = ProjectSpec{
		Replicas: in.Spec.Replicas,
	}
}

// Each type that is being served by the Kubernetes API
// (in this case, Project and ProjectList) needs to implement
// the k8s.io/apimachinery/pkg/runtime.Object interface.
// This interface defines the two methods GetObjectKind() and DeepCopyObject().
// The first method is already provided by the embedded meta_v1.TypeMeta struct;
// the second you’ll have to implement yourself.
func (in *Project) DeepCopyObject() runtime.Object {
	out := Project{}
	in.DeepCopyInto(&out)
	return &out
}

// The DeepCopyObject method is intended to generate a deep copy of an object.
// Since this involves a lot of boilerplate code, these methods are often automatically generated.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	out := ProjectList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]Project, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
	return &out
}
