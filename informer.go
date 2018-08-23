package main

import (
	"time"

	"github.com/ColdFreak/kubernetes-crd/api/types/v1alpha1"
	client_v1alpha1 "github.com/ColdFreak/kubernetes-crd/clientset/v1alpha1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet client_v1alpha1.ExampleV1Alpha1Interface) cache.Store {
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
