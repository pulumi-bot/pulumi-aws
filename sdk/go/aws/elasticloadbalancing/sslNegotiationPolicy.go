// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a load balancer SSL negotiation policy, which allows an ELB to control the ciphers and protocols that are supported during SSL negotiations between a client and a load balancer.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_ssl_negotiation_policy.html.markdown.
type SslNegotiationPolicy struct {
	pulumi.CustomResourceState

	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayOutput `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`
	// The name of the attribute
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewSslNegotiationPolicy registers a new resource with the given unique name, arguments, and options.
func NewSslNegotiationPolicy(ctx *pulumi.Context,
	name string, args *SslNegotiationPolicyArgs, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	if args == nil {
		args = &SslNegotiationPolicyArgs{}
	}
	var resource SslNegotiationPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSslNegotiationPolicy gets an existing SslNegotiationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSslNegotiationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SslNegotiationPolicyState, opts ...pulumi.ResourceOption) (*SslNegotiationPolicy, error) {
	var resource SslNegotiationPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/sslNegotiationPolicy:SslNegotiationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SslNegotiationPolicy resources.
type sslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes []SslNegotiationPolicyAttribute `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort *int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer *string `pulumi:"loadBalancer"`
	// The name of the attribute
	Name *string `pulumi:"name"`
}

type SslNegotiationPolicyState struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntPtrInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringPtrInput
	// The name of the attribute
	Name pulumi.StringPtrInput
}

func (SslNegotiationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*sslNegotiationPolicyState)(nil)).Elem()
}

type sslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes []SslNegotiationPolicyAttribute `pulumi:"attributes"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer string `pulumi:"loadBalancer"`
	// The name of the attribute
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SslNegotiationPolicy resource.
type SslNegotiationPolicyArgs struct {
	// An SSL Negotiation policy attribute. Each has two properties:
	Attributes SslNegotiationPolicyAttributeArrayInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput
	// The name of the attribute
	Name pulumi.StringPtrInput
}

func (SslNegotiationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sslNegotiationPolicyArgs)(nil)).Elem()
}

