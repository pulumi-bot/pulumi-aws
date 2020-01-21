// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package listenerPolicy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches a load balancer policy to an ELB Listener.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/load_balancer_listener_policy_legacy.html.markdown.
type ListenerPolicy struct {
	pulumi.CustomResourceState

	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntOutput `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayOutput `pulumi:"policyNames"`
}

// NewListenerPolicy registers a new resource with the given unique name, arguments, and options.
func NewListenerPolicy(ctx *pulumi.Context,
	name string, args *ListenerPolicyArgs, opts ...pulumi.ResourceOption) (*ListenerPolicy, error) {
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil || args.LoadBalancerPort == nil {
		return nil, errors.New("missing required argument 'LoadBalancerPort'")
	}
	if args == nil {
		args = &ListenerPolicyArgs{}
	}
	var resource ListenerPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/listenerPolicy:ListenerPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetListenerPolicy gets an existing ListenerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetListenerPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ListenerPolicyState, opts ...pulumi.ResourceOption) (*ListenerPolicy, error) {
	var resource ListenerPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/listenerPolicy:ListenerPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ListenerPolicy resources.
type listenerPolicyState struct {
	// The load balancer to attach the policy to.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort *int `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

type ListenerPolicyState struct {
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringPtrInput
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntPtrInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (ListenerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerPolicyState)(nil)).Elem()
}

type listenerPolicyArgs struct {
	// The load balancer to attach the policy to.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort int `pulumi:"loadBalancerPort"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

// The set of arguments for constructing a ListenerPolicy resource.
type ListenerPolicyArgs struct {
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringInput
	// The load balancer listener port to apply the policy to.
	LoadBalancerPort pulumi.IntInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (ListenerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*listenerPolicyArgs)(nil)).Elem()
}

