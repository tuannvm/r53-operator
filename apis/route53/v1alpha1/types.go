package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Route53 represents a pod terminator.
type Route53 struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the ddesired behaviour of the pod terminator.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Spec Route53Spec `json:"spec,omitempty"`
}

// Route53Spec is the spec for a Route53 resource.
type Route53Spec struct {
	// Selector is how the target will be selected.
	Selector map[string]string `json:"selector,omitempty"`
	// Type of route53 record.
	// More info: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/ResourceRecordTypes.html
	Type string `json:"type,omitempty"`
	// RoutingPolicy determines how Amazon Route 53 responds to queries
	// More info: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html
	RoutingPolicy string `json:"routing_policy,omitempty"`
	// Records is a map of domains (in CNAME)
	Records map[string]int64 `json:"records,omitempty"`
	// TTL time to live
	TTL    int64 `json:"ttl,omitempty"`
	DryRun bool  `json:"dryRun,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Route53List is a list of Route53 resources
type Route53List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Route53 `json:"items"`
}
