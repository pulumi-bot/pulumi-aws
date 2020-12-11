// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a load balancer policy, which can be attached to an ELB listener or backend server.
//
// Deprecated: aws.elasticloadbalancing.LoadBalancerPolicy has been deprecated in favor of aws.elb.LoadBalancerPolicy
type LoadBalancerPolicy struct {
	pulumi.CustomResourceState

	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayOutput `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName pulumi.StringOutput `pulumi:"policyTypeName"`
}

// NewLoadBalancerPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancerPolicy(ctx *pulumi.Context,
	name string, args *LoadBalancerPolicyArgs, opts ...pulumi.ResourceOption) (*LoadBalancerPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerName == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerName'")
	}
	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyTypeName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyTypeName'")
	}
	var resource LoadBalancerPolicy
	err := ctx.RegisterResource("aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancerPolicy gets an existing LoadBalancerPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancerPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerPolicyState, opts ...pulumi.ResourceOption) (*LoadBalancerPolicy, error) {
	var resource LoadBalancerPolicy
	err := ctx.ReadResource("aws:elasticloadbalancing/loadBalancerPolicy:LoadBalancerPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancerPolicy resources.
type loadBalancerPolicyState struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes []LoadBalancerPolicyPolicyAttribute `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName *string `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName *string `pulumi:"policyTypeName"`
}

type LoadBalancerPolicyState struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringPtrInput
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayInput
	// The name of the load balancer policy.
	PolicyName pulumi.StringPtrInput
	// The policy type.
	PolicyTypeName pulumi.StringPtrInput
}

func (LoadBalancerPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerPolicyState)(nil)).Elem()
}

type loadBalancerPolicyArgs struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName string `pulumi:"loadBalancerName"`
	// Policy attribute to apply to the policy.
	PolicyAttributes []LoadBalancerPolicyPolicyAttribute `pulumi:"policyAttributes"`
	// The name of the load balancer policy.
	PolicyName string `pulumi:"policyName"`
	// The policy type.
	PolicyTypeName string `pulumi:"policyTypeName"`
}

// The set of arguments for constructing a LoadBalancerPolicy resource.
type LoadBalancerPolicyArgs struct {
	// The load balancer on which the policy is defined.
	LoadBalancerName pulumi.StringInput
	// Policy attribute to apply to the policy.
	PolicyAttributes LoadBalancerPolicyPolicyAttributeArrayInput
	// The name of the load balancer policy.
	PolicyName pulumi.StringInput
	// The policy type.
	PolicyTypeName pulumi.StringInput
}

func (LoadBalancerPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerPolicyArgs)(nil)).Elem()
}

type LoadBalancerPolicyInput interface {
	pulumi.Input

	ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput
	ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput
}

type LoadBalancerPolicyPtrInput interface {
	pulumi.Input

	ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput
	ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput
}

func (LoadBalancerPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerPolicy)(nil)).Elem()
}

func (i LoadBalancerPolicy) ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput {
	return i.ToLoadBalancerPolicyOutputWithContext(context.Background())
}

func (i LoadBalancerPolicy) ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyOutput)
}

func (i LoadBalancerPolicy) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return i.ToLoadBalancerPolicyPtrOutputWithContext(context.Background())
}

func (i LoadBalancerPolicy) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerPolicyPtrOutput)
}

type LoadBalancerPolicyOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerPolicyOutput)(nil)).Elem()
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyOutput() LoadBalancerPolicyOutput {
	return o
}

func (o LoadBalancerPolicyOutput) ToLoadBalancerPolicyOutputWithContext(ctx context.Context) LoadBalancerPolicyOutput {
	return o
}

type LoadBalancerPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (LoadBalancerPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancerPolicy)(nil)).Elem()
}

func (o LoadBalancerPolicyPtrOutput) ToLoadBalancerPolicyPtrOutput() LoadBalancerPolicyPtrOutput {
	return o
}

func (o LoadBalancerPolicyPtrOutput) ToLoadBalancerPolicyPtrOutputWithContext(ctx context.Context) LoadBalancerPolicyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LoadBalancerPolicyOutput{})
	pulumi.RegisterOutputType(LoadBalancerPolicyPtrOutput{})
}
