// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ECR repository lifecycle policy.
//
// > **NOTE:** Only one `ecr.LifecyclePolicy` resource can be used with the same ECR repository. To apply multiple rules, they must be combined in the `policy` JSON.
//
// > **NOTE:** The AWS ECR API seems to reorder rules based on `rulePriority`. If you define multiple rules that are not sorted in ascending `rulePriority` order in the this provider code, the resource will be flagged for recreation every deployment.
//
// ## Example Usage
// ### Policy on untagged image
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := ecr.NewRepository(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecr.NewLifecyclePolicy(ctx, "foopolicy", &ecr.LifecyclePolicyArgs{
// 			Repository: foo.Name,
// 			Policy:     pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"rules\": [\n", "        {\n", "            \"rulePriority\": 1,\n", "            \"description\": \"Expire images older than 14 days\",\n", "            \"selection\": {\n", "                \"tagStatus\": \"untagged\",\n", "                \"countType\": \"sinceImagePushed\",\n", "                \"countUnit\": \"days\",\n", "                \"countNumber\": 14\n", "            },\n", "            \"action\": {\n", "                \"type\": \"expire\"\n", "            }\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Policy on tagged image
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := ecr.NewRepository(ctx, "foo", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecr.NewLifecyclePolicy(ctx, "foopolicy", &ecr.LifecyclePolicyArgs{
// 			Repository: foo.Name,
// 			Policy:     pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"rules\": [\n", "        {\n", "            \"rulePriority\": 1,\n", "            \"description\": \"Keep last 30 images\",\n", "            \"selection\": {\n", "                \"tagStatus\": \"tagged\",\n", "                \"tagPrefixList\": [\"v\"],\n", "                \"countType\": \"imageCountMoreThan\",\n", "                \"countNumber\": 30\n", "            },\n", "            \"action\": {\n", "                \"type\": \"expire\"\n", "            }\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// ECR Lifecycle Policy can be imported using the name of the repository, e.g.
//
// ```sh
//  $ pulumi import aws:ecr/lifecyclePolicy:LifecyclePolicy example tf-example
// ```
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository pulumi.StringOutput `pulumi:"repository"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.Repository == nil {
		return nil, errors.New("invalid value for required argument 'Repository'")
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
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy interface{} `pulumi:"policy"`
	// The registry ID where the repository was created.
	RegistryId *string `pulumi:"registryId"`
	// Name of the repository to apply the policy.
	Repository *string `pulumi:"repository"`
}

type LifecyclePolicyState struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy pulumi.Input
	// The registry ID where the repository was created.
	RegistryId pulumi.StringPtrInput
	// Name of the repository to apply the policy.
	Repository pulumi.StringPtrInput
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy interface{} `pulumi:"policy"`
	// Name of the repository to apply the policy.
	Repository string `pulumi:"repository"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// The policy document. This is a JSON formatted string. See more details about [Policy Parameters](http://docs.aws.amazon.com/AmazonECR/latest/userguide/LifecyclePolicies.html#lifecycle_policy_parameters) in the official AWS docs.
	Policy pulumi.Input
	// Name of the repository to apply the policy.
	Repository pulumi.StringInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

type LifecyclePolicyInput interface {
	pulumi.Input

	ToLifecyclePolicyOutput() LifecyclePolicyOutput
	ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput
}

func (*LifecyclePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecyclePolicy)(nil))
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return i.ToLifecyclePolicyOutputWithContext(context.Background())
}

func (i *LifecyclePolicy) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyOutput)
}

func (i *LifecyclePolicy) ToLifecyclePolicyPtrOutput() LifecyclePolicyPtrOutput {
	return i.ToLifecyclePolicyPtrOutputWithContext(context.Background())
}

func (i *LifecyclePolicy) ToLifecyclePolicyPtrOutputWithContext(ctx context.Context) LifecyclePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyPtrOutput)
}

type LifecyclePolicyPtrInput interface {
	pulumi.Input

	ToLifecyclePolicyPtrOutput() LifecyclePolicyPtrOutput
	ToLifecyclePolicyPtrOutputWithContext(ctx context.Context) LifecyclePolicyPtrOutput
}

type lifecyclePolicyPtrType LifecyclePolicyArgs

func (*lifecyclePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil))
}

func (i *lifecyclePolicyPtrType) ToLifecyclePolicyPtrOutput() LifecyclePolicyPtrOutput {
	return i.ToLifecyclePolicyPtrOutputWithContext(context.Background())
}

func (i *lifecyclePolicyPtrType) ToLifecyclePolicyPtrOutputWithContext(ctx context.Context) LifecyclePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyPtrOutput)
}

// LifecyclePolicyArrayInput is an input type that accepts LifecyclePolicyArray and LifecyclePolicyArrayOutput values.
// You can construct a concrete instance of `LifecyclePolicyArrayInput` via:
//
//          LifecyclePolicyArray{ LifecyclePolicyArgs{...} }
type LifecyclePolicyArrayInput interface {
	pulumi.Input

	ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput
	ToLifecyclePolicyArrayOutputWithContext(context.Context) LifecyclePolicyArrayOutput
}

type LifecyclePolicyArray []LifecyclePolicyInput

func (LifecyclePolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LifecyclePolicy)(nil))
}

func (i LifecyclePolicyArray) ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput {
	return i.ToLifecyclePolicyArrayOutputWithContext(context.Background())
}

func (i LifecyclePolicyArray) ToLifecyclePolicyArrayOutputWithContext(ctx context.Context) LifecyclePolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyArrayOutput)
}

// LifecyclePolicyMapInput is an input type that accepts LifecyclePolicyMap and LifecyclePolicyMapOutput values.
// You can construct a concrete instance of `LifecyclePolicyMapInput` via:
//
//          LifecyclePolicyMap{ "key": LifecyclePolicyArgs{...} }
type LifecyclePolicyMapInput interface {
	pulumi.Input

	ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput
	ToLifecyclePolicyMapOutputWithContext(context.Context) LifecyclePolicyMapOutput
}

type LifecyclePolicyMap map[string]LifecyclePolicyInput

func (LifecyclePolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LifecyclePolicy)(nil))
}

func (i LifecyclePolicyMap) ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput {
	return i.ToLifecyclePolicyMapOutputWithContext(context.Background())
}

func (i LifecyclePolicyMap) ToLifecyclePolicyMapOutputWithContext(ctx context.Context) LifecyclePolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyMapOutput)
}

type LifecyclePolicyOutput struct {
	*pulumi.OutputState
}

func (LifecyclePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecyclePolicy)(nil))
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyPtrOutput() LifecyclePolicyPtrOutput {
	return o.ToLifecyclePolicyPtrOutputWithContext(context.Background())
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyPtrOutputWithContext(ctx context.Context) LifecyclePolicyPtrOutput {
	return o.ApplyT(func(v LifecyclePolicy) *LifecyclePolicy {
		return &v
	}).(LifecyclePolicyPtrOutput)
}

type LifecyclePolicyPtrOutput struct {
	*pulumi.OutputState
}

func (LifecyclePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LifecyclePolicy)(nil))
}

func (o LifecyclePolicyPtrOutput) ToLifecyclePolicyPtrOutput() LifecyclePolicyPtrOutput {
	return o
}

func (o LifecyclePolicyPtrOutput) ToLifecyclePolicyPtrOutputWithContext(ctx context.Context) LifecyclePolicyPtrOutput {
	return o
}

type LifecyclePolicyArrayOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LifecyclePolicy)(nil))
}

func (o LifecyclePolicyArrayOutput) ToLifecyclePolicyArrayOutput() LifecyclePolicyArrayOutput {
	return o
}

func (o LifecyclePolicyArrayOutput) ToLifecyclePolicyArrayOutputWithContext(ctx context.Context) LifecyclePolicyArrayOutput {
	return o
}

func (o LifecyclePolicyArrayOutput) Index(i pulumi.IntInput) LifecyclePolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LifecyclePolicy {
		return vs[0].([]LifecyclePolicy)[vs[1].(int)]
	}).(LifecyclePolicyOutput)
}

type LifecyclePolicyMapOutput struct{ *pulumi.OutputState }

func (LifecyclePolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LifecyclePolicy)(nil))
}

func (o LifecyclePolicyMapOutput) ToLifecyclePolicyMapOutput() LifecyclePolicyMapOutput {
	return o
}

func (o LifecyclePolicyMapOutput) ToLifecyclePolicyMapOutputWithContext(ctx context.Context) LifecyclePolicyMapOutput {
	return o
}

func (o LifecyclePolicyMapOutput) MapIndex(k pulumi.StringInput) LifecyclePolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LifecyclePolicy {
		return vs[0].(map[string]LifecyclePolicy)[vs[1].(string)]
	}).(LifecyclePolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
	pulumi.RegisterOutputType(LifecyclePolicyPtrOutput{})
	pulumi.RegisterOutputType(LifecyclePolicyArrayOutput{})
	pulumi.RegisterOutputType(LifecyclePolicyMapOutput{})
}
