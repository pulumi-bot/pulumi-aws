// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an ECR repository lifecycle policy.
// 
// > **NOTE:** Only one `ecr.LifecyclePolicy` resource can be used with the same ECR repository. To apply multiple rules, they must be combined in the `policy` JSON.
// 
// > **NOTE:** The AWS ECR API seems to reorder rules based on `rulePriority`. If you define multiple rules that are not sorted in ascending `rulePriority` order in the this provider code, the resource will be flagged for recreation every deployment.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecr_lifecycle_policy.html.markdown.
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	Policy pulumi.StringOutput `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	if args == nil || args.Repository == nil {
		return nil, errors.New("missing required argument 'Repository'")
	}
	if args == nil {
		args = &LifecyclePolicyArgs{}
	}
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws:ecr/lifecyclePolicy:LifecyclePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLifecyclePolicy gets an existing LifecyclePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLifecyclePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LifecyclePolicyState, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	var resource LifecyclePolicy
	err := ctx.ReadResource("aws:ecr/lifecyclePolicy:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
	Policy *string `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository *string `pulumi:"repository"`
}

type LifecyclePolicyState struct {
	Policy pulumi.StringPtrInput
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// Name of the repository to apply the policy.
	Repository pulumi.StringPtrInput
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	Policy interface{} `pulumi:"policy"`
	// Name of the repository to apply the policy.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	Policy pulumi.Input
	// Name of the repository to apply the policy.
	Repository pulumi.StringInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

