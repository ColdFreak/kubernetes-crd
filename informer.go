package main

import (
	"time"

	// "github.com/ColdFreak/kubernetes-crd/api/types/v1alpha1"
	"./api/types/v1alpha1"
	client_v1alpha1 "./clientset/v1alpha1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet client_v1alpha1.ExampleV1Alpha1Interface) cache.Store {
	// The NewInformer method returns two objects:
	// The second return value, the controller controls the List() and Watch() calls
	// and fills the first return value, the store with a (more or less) recent cache of
	// the watched resource’s state on the API server (in this case, the project CRD).
	projectStore, projectController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(lo meta_v1.ListOptions) (result runtime.Object, err error) {
				return clientSet.Projects("default").List(lo)
			},
			WatchFunc: func(lo meta_v1.ListOptions) (watch.Interface, error) {
				return clientSet.Projects("default").Watch(lo)
			},
		},
		&v1alpha1.Project{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)

	go projectController.Run(wait.NeverStop)
	return projectStore
}
