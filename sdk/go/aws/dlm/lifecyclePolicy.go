// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dlm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a [Data Lifecycle Manager (DLM) lifecycle policy](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-lifecycle.html) for managing snapshots.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/dlm"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		dlmLifecycleRole, err := iam.NewRole(ctx, "dlmLifecycleRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"dlm.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "dlmLifecycle", &iam.RolePolicyArgs{
// 			Role:   dlmLifecycleRole.ID(),
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "   \"Version\": \"2012-10-17\",\n", "   \"Statement\": [\n", "      {\n", "         \"Effect\": \"Allow\",\n", "         \"Action\": [\n", "            \"ec2:CreateSnapshot\",\n", "            \"ec2:DeleteSnapshot\",\n", "            \"ec2:DescribeVolumes\",\n", "            \"ec2:DescribeSnapshots\"\n", "         ],\n", "         \"Resource\": \"*\"\n", "      },\n", "      {\n", "         \"Effect\": \"Allow\",\n", "         \"Action\": [\n", "            \"ec2:CreateTags\"\n", "         ],\n", "         \"Resource\": \"arn:aws:ec2:*::snapshot/*\"\n", "      }\n", "   ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dlm.NewLifecyclePolicy(ctx, "example", &dlm.LifecyclePolicyArgs{
// 			Description:      pulumi.String("example DLM lifecycle policy"),
// 			ExecutionRoleArn: dlmLifecycleRole.Arn,
// 			State:            pulumi.String("ENABLED"),
// 			PolicyDetails: &dlm.LifecyclePolicyPolicyDetailsArgs{
// 				ResourceTypes: pulumi.StringArray{
// 					pulumi.String("VOLUME"),
// 				},
// 				Schedules: dlm.LifecyclePolicyPolicyDetailsScheduleArray{
// 					&dlm.LifecyclePolicyPolicyDetailsScheduleArgs{
// 						Name: pulumi.String("2 weeks of daily snapshots"),
// 						CreateRule: &dlm.LifecyclePolicyPolicyDetailsScheduleCreateRuleArgs{
// 							Interval:     pulumi.Int(24),
// 							IntervalUnit: pulumi.String("HOURS"),
// 							Times: pulumi.String(pulumi.String{
// 								pulumi.String("23:45"),
// 							}),
// 						},
// 						RetainRule: &dlm.LifecyclePolicyPolicyDetailsScheduleRetainRuleArgs{
// 							Count: pulumi.Int(14),
// 						},
// 						TagsToAdd: pulumi.StringMap{
// 							"SnapshotCreator": pulumi.String("DLM"),
// 						},
// 						CopyTags: pulumi.Bool(false),
// 					},
// 				},
// 				TargetTags: pulumi.StringMap{
// 					"Snapshot": pulumi.String("true"),
// 				},
// 			},
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
// DLM lifecyle policies can be imported by their policy ID
//
// ```sh
//  $ pulumi import aws:dlm/lifecyclePolicy:LifecyclePolicy example policy-abcdef12345678901
// ```
type LifecyclePolicy struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description for the DLM lifecycle policy.
	Description pulumi.StringOutput `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsOutput `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrOutput `pulumi:"state"`
	// Key-value map of resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewLifecyclePolicy registers a new resource with the given unique name, arguments, and options.
func NewLifecyclePolicy(ctx *pulumi.Context,
	name string, args *LifecyclePolicyArgs, opts ...pulumi.ResourceOption) (*LifecyclePolicy, error) {
	if args == nil || args.Description == nil {
		return nil, errors.New("missing required argument 'Description'")
	}
	if args == nil || args.ExecutionRoleArn == nil {
		return nil, errors.New("missing required argument 'ExecutionRoleArn'")
	}
	if args == nil || args.PolicyDetails == nil {
		return nil, errors.New("missing required argument 'PolicyDetails'")
	}
	if args == nil {
		args = &LifecyclePolicyArgs{}
	}
	var resource LifecyclePolicy
	err := ctx.RegisterResource("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:dlm/lifecyclePolicy:LifecyclePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LifecyclePolicy resources.
type lifecyclePolicyState struct {
	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn *string `pulumi:"arn"`
	// A description for the DLM lifecycle policy.
	Description *string `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails *LifecyclePolicyPolicyDetails `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

type LifecyclePolicyState struct {
	// Amazon Resource Name (ARN) of the DLM Lifecycle Policy.
	Arn pulumi.StringPtrInput
	// A description for the DLM lifecycle policy.
	Description pulumi.StringPtrInput
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringPtrInput
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsPtrInput
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (LifecyclePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyState)(nil)).Elem()
}

type lifecyclePolicyArgs struct {
	// A description for the DLM lifecycle policy.
	Description string `pulumi:"description"`
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetails `pulumi:"policyDetails"`
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State *string `pulumi:"state"`
	// Key-value map of resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LifecyclePolicy resource.
type LifecyclePolicyArgs struct {
	// A description for the DLM lifecycle policy.
	Description pulumi.StringInput
	// The ARN of an IAM role that is able to be assumed by the DLM service.
	ExecutionRoleArn pulumi.StringInput
	// See the `policyDetails` configuration block. Max of 1.
	PolicyDetails LifecyclePolicyPolicyDetailsInput
	// Whether the lifecycle policy should be enabled or disabled. `ENABLED` or `DISABLED` are valid values. Defaults to `ENABLED`.
	State pulumi.StringPtrInput
	// Key-value map of resource tags.
	Tags pulumi.StringMapInput
}

func (LifecyclePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lifecyclePolicyArgs)(nil)).Elem()
}

type LifecyclePolicyInput interface {
	pulumi.Input

	ToLifecyclePolicyOutput() LifecyclePolicyOutput
	ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput
}

func (LifecyclePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecyclePolicy)(nil)).Elem()
}

func (i LifecyclePolicy) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return i.ToLifecyclePolicyOutputWithContext(context.Background())
}

func (i LifecyclePolicy) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LifecyclePolicyOutput)
}

type LifecyclePolicyOutput struct {
	*pulumi.OutputState
}

func (LifecyclePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LifecyclePolicyOutput)(nil)).Elem()
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutput() LifecyclePolicyOutput {
	return o
}

func (o LifecyclePolicyOutput) ToLifecyclePolicyOutputWithContext(ctx context.Context) LifecyclePolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LifecyclePolicyOutput{})
}
