package operator

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	route53v1alpha1 "github.com/tuannvm/r53-operator/apis/route53/v1alpha1"
	"github.com/tuannvm/r53-operator/log"
	"github.com/tuannvm/r53-operator/service/route53"
)

// Handler  is the pod terminator handler that will handle the
// events received from kubernetes.
type handler struct {
	route53Service route53.Syncer
	logger         log.Logger
}

// newHandler returns a new handler.
func newHandler(k8sCli kubernetes.Interface, logger log.Logger) *handler {
	return &handler{
		route53Service: route53.Newroute53(k8sCli, logger),
		logger:         logger,
	}
}

// Add will ensure that the required pod terminator is running.
func (h *handler) Add(_ context.Context, obj runtime.Object) error {
	pt, ok := obj.(*route53v1alpha1.Route53)
	if !ok {
		return fmt.Errorf("%v is not a pod terminator object", obj.GetObjectKind())
	}

	return h.route53Service.EnsureRoute53(pt)
}

// Delete will ensure the reuited pod terminator is not running.
func (h *handler) Delete(_ context.Context, name string) error {
	return h.route53Service.DeleteRoute53(name)
}
