package route53

import (
	"sync"

	"k8s.io/client-go/kubernetes"

	"github.com/aws/aws-sdk-go/aws"
	awsRoute53 "github.com/aws/aws-sdk-go/service/route53"

	route53v1alpha1 "github.com/tuannvm/r53-operator/apis/route53/v1alpha1"
	"github.com/tuannvm/r53-operator/log"
	"github.com/tuannvm/tools/pkg/awsutils"
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
	k8sCli    kubernetes.Interface
	reg       sync.Map
	logger    log.Logger
	awsClient *awsutils.Client
}

// Newroute53 returns a new route53 service.
func Newroute53(k8sCli kubernetes.Interface, logger log.Logger) *route53 {
	awsConfig := awsutils.NewConfig(&awsutils.Config{})
	awsClient := awsutils.New(awsConfig)
	return &route53{
		k8sCli:    k8sCli,
		reg:       sync.Map{},
		logger:    logger,
		awsClient: awsClient,
	}
}

// EnsureRoute53 satisfies route53Syncer interface.
func (c *route53) EnsureRoute53(pt *route53v1alpha1.Route53) error {
	zoneID, err := c.awsClient.GetHostedZoneID(pt.Spec.Domain)
	if err != nil {
		return err
	}

	records := pt.Spec.Records
	for record, weight := range records {
		err := c.awsClient.ChangeRoute53Record(zoneID, &awsRoute53.ResourceRecordSet{
			Name:          aws.String(pt.Spec.Name + "." + pt.Spec.Domain),
			TTL:           aws.Int64(pt.Spec.TTL),
			Type:          aws.String(pt.Spec.Type),
			Weight:        aws.Int64(weight),
			SetIdentifier: aws.String(record),
			ResourceRecords: []*awsRoute53.ResourceRecord{
				{
					Value: aws.String(record),
				},
			},
		}, aws.String("UPSERT"))
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteRoute53 satisfies route53Syncer interface.
func (c *route53) DeleteRoute53(name string) error {
	c.logger.Infof("456")
	return nil
}
