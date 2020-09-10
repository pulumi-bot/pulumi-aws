// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LoadBalancerBackendServerPolicy struct {
	pulumi.CustomResourceState

	InstancePort     pulumi.IntOutput         `pulumi:"instancePort"`
	LoadBalancerName pulumi.StringOutput      `pulumi:"loadBalancerName"`
	PolicyNames      pulumi.StringArrayOutput `pulumi:"policyNames"`
}

// NewLoadBalancerBackendServerPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendServerPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendServerPolicyArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendServerPolicy, error) {
	if args == nil || args.InstancePort == nil {
		return nil, errors.New("missing required argument 'InstancePort'")
	}
	if args == nil || args.LoadBalancerName == nil {
		return nil, errors.New("missing required argument 'LoadBalancerName'")
	}
	if args == nil {
		args = &LoadBalancerBackendServerPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:elasticloadbalancing/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource LoadBalancerBackendServerPolicy
	err := ctx.RegisterResource("aws:elb/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerBackendServerPolicy gets an existing LoadBalancerBackendServerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerBackendServerPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerBackendServerPolicyState, opts ...pulumi.ResourceOption) (*LoadBalancerBackendServerPolicy, error) {
	var resource LoadBalancerBackendServerPolicy
	err := ctx.ReadResource("aws:elb/loadBalancerBackendServerPolicy:LoadBalancerBackendServerPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerBackendServerPolicy resources.
type loadBalancerBackendServerPolicyState struct {
	InstancePort     *int     `pulumi:"instancePort"`
	LoadBalancerName *string  `pulumi:"loadBalancerName"`
	PolicyNames      []string `pulumi:"policyNames"`
}

type LoadBalancerBackendServerPolicyState struct {
	InstancePort     pulumi.IntPtrInput
	LoadBalancerName pulumi.StringPtrInput
	PolicyNames      pulumi.StringArrayInput
}

func (LoadBalancerBackendServerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendServerPolicyState)(nil)).Elem()
}

type loadBalancerBackendServerPolicyArgs struct {
	InstancePort     int      `pulumi:"instancePort"`
	LoadBalancerName string   `pulumi:"loadBalancerName"`
	PolicyNames      []string `pulumi:"policyNames"`
}

// The set of arguments for constructing a LoadBalancerBackendServerPolicy resource.
type LoadBalancerBackendServerPolicyArgs struct {
	InstancePort     pulumi.IntInput
	LoadBalancerName pulumi.StringInput
	PolicyNames      pulumi.StringArrayInput
}

func (LoadBalancerBackendServerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendServerPolicyArgs)(nil)).Elem()
}
