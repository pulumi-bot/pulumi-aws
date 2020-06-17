// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a load balancer cookie stickiness policy, which allows an ELB to control the sticky session lifetime of the browser.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		lb, err := elb.NewLoadBalancer(ctx, "lb", &elb.LoadBalancerArgs{
// 			AvailabilityZones: pulumi.StringArray{
// 				pulumi.String("us-east-1a"),
// 			},
// 			Listeners: elb.LoadBalancerListenerArray{
// 				&elb.LoadBalancerListenerArgs{
// 					InstancePort:     pulumi.Int(8000),
// 					InstanceProtocol: pulumi.String("http"),
// 					LbPort:           pulumi.Int(80),
// 					LbProtocol:       pulumi.String("http"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancerCookieStickinessPolicy(ctx, "foo", &elb.LoadBalancerCookieStickinessPolicyArgs{
// 			CookieExpirationPeriod: pulumi.Int(600),
// 			LbPort:                 pulumi.Int(80),
// 			LoadBalancer:           lb.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Deprecated: aws.elasticloadbalancing.LoadBalancerCookieStickinessPolicy has been deprecated in favor of aws.elb.LoadBalancerCookieStickinessPolicy
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
