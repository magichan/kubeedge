/*
Copyright The KubeEdge Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	reliablesyncsv1alpha1 "github.com/kubeedge/api/reliablesyncs/v1alpha1"
	versioned "github.com/kubeedge/kubeedge/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeedge/kubeedge/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubeedge/kubeedge/pkg/client/listers/reliablesyncs/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterObjectSyncInformer provides access to a shared informer and lister for
// ClusterObjectSyncs.
type ClusterObjectSyncInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterObjectSyncLister
}

type clusterObjectSyncInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterObjectSyncInformer constructs a new informer for ClusterObjectSync type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterObjectSyncInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterObjectSyncInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterObjectSyncInformer constructs a new informer for ClusterObjectSync type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterObjectSyncInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReliablesyncsV1alpha1().ClusterObjectSyncs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReliablesyncsV1alpha1().ClusterObjectSyncs().Watch(context.TODO(), options)
			},
		},
		&reliablesyncsv1alpha1.ClusterObjectSync{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterObjectSyncInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterObjectSyncInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterObjectSyncInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&reliablesyncsv1alpha1.ClusterObjectSync{}, f.defaultInformer)
}

func (f *clusterObjectSyncInformer) Lister() v1alpha1.ClusterObjectSyncLister {
	return v1alpha1.NewClusterObjectSyncLister(f.Informer().GetIndexer())
}
