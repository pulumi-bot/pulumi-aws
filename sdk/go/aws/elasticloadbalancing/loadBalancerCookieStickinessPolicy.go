// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package elasticloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_cookie_stickiness_policy.html.markdown.
type LoadBalancerCookieStickinessPolicy struct {
	pulumi.CustomResourceState

	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntPtrOutput `pulumi:"cookieExpirationPeriod"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntOutput `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringOutput `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewLoadBalancerCookieStickinessPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerCookieStickinessPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerCookieStickinessPolicyArgs, opts ...pulumi.ResourceOption) (*LoadBalancerCookieStickinessPolicy, error) {
	if args == nil || args.LbPort == nil {
		return nil, errors.New("missing required argument 'LbPort'")
	}
	if args == nil || args.LoadBalancer == nil {
		return nil, errors.New("missing required argument 'LoadBalancer'")
	}
	if args == nil {
		args = &LoadBalancerCookieStickinessPolicyArgs{}
	}
	var resource LoadBalancerCookieStickinessPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerCookieStickinessPolicy gets an existing LoadBalancerCookieStickinessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerCookieStickinessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerCookieStickinessPolicyState, opts ...pulumi.ResourceOption) (*LoadBalancerCookieStickinessPolicy, error) {
	var resource LoadBalancerCookieStickinessPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancerCookieStickinessPolicy:LoadBalancerCookieStickinessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerCookieStickinessPolicy resources.
type loadBalancerCookieStickinessPolicyState struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod *int `pulumi:"cookieExpirationPeriod"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort *int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer *string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

type LoadBalancerCookieStickinessPolicyState struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntPtrInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntPtrInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringPtrInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (LoadBalancerCookieStickinessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCookieStickinessPolicyState)(nil)).Elem()
}

type loadBalancerCookieStickinessPolicyArgs struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod *int `pulumi:"cookieExpirationPeriod"`
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort int `pulumi:"lbPort"`
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer string `pulumi:"loadBalancer"`
	// The name of the stickiness policy.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a LoadBalancerCookieStickinessPolicy resource.
type LoadBalancerCookieStickinessPolicyArgs struct {
	// The time period after which
	// the session cookie should be considered stale, expressed in seconds.
	CookieExpirationPeriod pulumi.IntPtrInput
	// The load balancer port to which the policy
	// should be applied. This must be an active listener on the load
	// balancer.
	LbPort pulumi.IntInput
	// The load balancer to which the policy
	// should be attached.
	LoadBalancer pulumi.StringInput
	// The name of the stickiness policy.
	Name pulumi.StringPtrInput
}

func (LoadBalancerCookieStickinessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerCookieStickinessPolicyArgs)(nil)).Elem()
}

