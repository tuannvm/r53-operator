package operator

import (
	"github.com/spotahome/kooper/client/crd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	route53v1alpha1 "github.com/tuannvm/r53-operator/apis/route53/v1alpha1"
	route53k8scli "github.com/tuannvm/r53-operator/client/k8s/clientset/versioned"
)

// route53CRD is the crd pod terminator.
type route53CRD struct {
	crdCli     crd.Interface
	kubeCli    kubernetes.Interface
	route53Cli route53k8scli.Interface
}

func newRoute53CRD(route53Cli route53k8scli.Interface, crdCli crd.Interface, kubeCli kubernetes.Interface) *route53CRD {
	return &route53CRD{
		crdCli:     crdCli,
		route53Cli: route53Cli,
		kubeCli:    kubeCli,
	}
}

// route53CRD satisfies resource.crd interface.
func (p *route53CRD) Initialize() error {
	crd := crd.Conf{
		Kind:       route53v1alpha1.Route53Kind,
		NamePlural: route53v1alpha1.Route53NamePlural,
		Group:      route53v1alpha1.SchemeGroupVersion.Group,
		Version:    route53v1alpha1.SchemeGroupVersion.Version,
		Scope:      route53v1alpha1.Route53Scope,
		Categories: []string{"route53", "r53"},
	}

	return p.crdCli.EnsurePresent(crd)
}

// GetListerWatcher satisfies resource.crd interface (and retrieve.Retriever).
func (p *route53CRD) GetListerWatcher() cache.ListerWatcher {
	return &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			return p.route53Cli.Route53V1alpha1().Route53s().List(options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			return p.route53Cli.Route53V1alpha1().Route53s().Watch(options)
		},
	}
}

// GetObject satisfies resource.crd interface (and retrieve.Retriever).
func (p *route53CRD) GetObject() runtime.Object {
	return &route53v1alpha1.Route53{}
}
