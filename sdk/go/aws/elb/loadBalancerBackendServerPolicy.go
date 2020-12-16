// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches a load balancer policy to an ELB backend server.
type LoadBalancerBackendServerPolicy struct {
	pulumi.CustomResourceState

	// The instance port to apply the policy to.
	InstancePort pulumi.IntOutput `pulumi:"instancePort"`
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayOutput `pulumi:"policyNames"`
}

// NewLoadBalancerBackendServerPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerBackendServerPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerBackendServerPolicyArgs, opts ...pulumi.ResourceOption) (*LoadBalancerBackendServerPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstancePort == nil {
		return nil, errors.New("invalid value for required argument 'InstancePort'")
	}
	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
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
	// The instance port to apply the policy to.
	InstancePort *int `pulumi:"instancePort"`
	// The load balancer to attach the policy to.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

type LoadBalancerBackendServerPolicyState struct {
	// The instance port to apply the policy to.
	InstancePort pulumi.IntPtrInput
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringPtrInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (LoadBalancerBackendServerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendServerPolicyState)(nil)).Elem()
}

type loadBalancerBackendServerPolicyArgs struct {
	// The instance port to apply the policy to.
	InstancePort int `pulumi:"instancePort"`
	// The load balancer to attach the policy to.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// List of Policy Names to apply to the backend server.
	PolicyNames []string `pulumi:"policyNames"`
}

// The set of arguments for constructing a LoadBalancerBackendServerPolicy resource.
type LoadBalancerBackendServerPolicyArgs struct {
	// The instance port to apply the policy to.
	InstancePort pulumi.IntInput
	// The load balancer to attach the policy to.
	LoadBalancerName pulumi.StringInput
	// List of Policy Names to apply to the backend server.
	PolicyNames pulumi.StringArrayInput
}

func (LoadBalancerBackendServerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerBackendServerPolicyArgs)(nil)).Elem()
}

type LoadBalancerBackendServerPolicyInput interface {
	pulumi.Input

	ToLoadBalancerBackendServerPolicyOutput() LoadBalancerBackendServerPolicyOutput
	ToLoadBalancerBackendServerPolicyOutputWithContext(ctx context.Context) LoadBalancerBackendServerPolicyOutput
}

func (*LoadBalancerBackendServerPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendServerPolicy)(nil))
}

func (i *LoadBalancerBackendServerPolicy) ToLoadBalancerBackendServerPolicyOutput() LoadBalancerBackendServerPolicyOutput {
	return i.ToLoadBalancerBackendServerPolicyOutputWithContext(context.Background())
}

func (i *LoadBalancerBackendServerPolicy) ToLoadBalancerBackendServerPolicyOutputWithContext(ctx context.Context) LoadBalancerBackendServerPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendServerPolicyOutput)
}

type LoadBalancerBackendServerPolicyOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerBackendServerPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendServerPolicy)(nil))
}

func (o LoadBalancerBackendServerPolicyOutput) ToLoadBalancerBackendServerPolicyOutput() LoadBalancerBackendServerPolicyOutput {
	return o
}

func (o LoadBalancerBackendServerPolicyOutput) ToLoadBalancerBackendServerPolicyOutputWithContext(ctx context.Context) LoadBalancerBackendServerPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerBackendServerPolicyOutput{})
}
