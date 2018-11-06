package route53

import (
	"sync"

	"k8s.io/client-go/kubernetes"

	route53v1alpha1 "github.com/tuannvm/r53-operator/apis/route53/v1alpha1"
	"github.com/tuannvm/r53-operator/log"
)

// Syncer is the interface that every r53 service implementation
// needs to implement.
type Syncer interface {
	// EnsureRoute53 will ensure that the pod terminator is running and working.
	EnsureRoute53(pt *route53v1alpha1.Route53) error
	// DeleteRoute53 will stop and delete the pod terminator.
	DeleteRoute53(name string) error
}

// route53 is the service that will ensure that the desired pod terminator CRDs are met.
// route53 will have running instances of PodDestroyers.
type route53 struct {
	k8sCli kubernetes.Interface
	reg    sync.Map
	logger log.Logger
}

// Newroute53 returns a new route53 service.
func Newroute53(k8sCli kubernetes.Interface, logger log.Logger) *route53 {
	return &route53{
		k8sCli: k8sCli,
		reg:    sync.Map{},
		logger: logger,
	}
}

// EnsureRoute53 satisfies route53Syncer interface.
func (c *route53) EnsureRoute53(pt *route53v1alpha1.Route53) error {
	c.logger.Infof("123")
	return nil
}

// DeleteRoute53 satisfies route53Syncer interface.
func (c *route53) DeleteRoute53(name string) error {
	c.logger.Infof("456")
	return nil
}
