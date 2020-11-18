// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a CloudWatch Logs destination policy resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		testDestination, err := cloudwatch.NewLogDestination(ctx, "testDestination", &cloudwatch.LogDestinationArgs{
// 			RoleArn:   pulumi.Any(aws_iam_role.Iam_for_cloudwatch.Arn),
// 			TargetArn: pulumi.Any(aws_kinesis_stream.Kinesis_for_cloudwatch.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudwatch.NewLogDestinationPolicy(ctx, "testDestinationPolicyLogDestinationPolicy", &cloudwatch.LogDestinationPolicyArgs{
// 			DestinationName: testDestination.Name,
// 			AccessPolicy: testDestinationPolicyPolicyDocument.ApplyT(func(testDestinationPolicyPolicyDocument iam.GetPolicyDocumentResult) (string, error) {
// 				return testDestinationPolicyPolicyDocument.Json, nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LogDestinationPolicy struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringOutput `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName pulumi.StringOutput `pulumi:"destinationName"`
}

// NewLogDestinationPolicy registers a new resource with the given unique name, arguments, and options.
func NewLogDestinationPolicy(ctx *pulumi.Context,
	name string, args *LogDestinationPolicyArgs, opts ...pulumi.ResourceOption) (*LogDestinationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessPolicy == nil {
		return nil, errors.New("invalid value for required argument 'AccessPolicy'")
	}
	if args.DestinationName == nil {
		return nil, errors.New("invalid value for required argument 'DestinationName'")
	}
	var resource LogDestinationPolicy
	err := ctx.RegisterResource("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogDestinationPolicy gets an existing LogDestinationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogDestinationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogDestinationPolicyState, opts ...pulumi.ResourceOption) (*LogDestinationPolicy, error) {
	var resource LogDestinationPolicy
	err := ctx.ReadResource("aws:cloudwatch/logDestinationPolicy:LogDestinationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogDestinationPolicy resources.
type logDestinationPolicyState struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy *string `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName *string `pulumi:"destinationName"`
}

type LogDestinationPolicyState struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringPtrInput
	// A name for the subscription filter
	DestinationName pulumi.StringPtrInput
}

func (LogDestinationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationPolicyState)(nil)).Elem()
}

type logDestinationPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy string `pulumi:"accessPolicy"`
	// A name for the subscription filter
	DestinationName string `pulumi:"destinationName"`
}

// The set of arguments for constructing a LogDestinationPolicy resource.
type LogDestinationPolicyArgs struct {
	// The policy document. This is a JSON formatted string.
	AccessPolicy pulumi.StringInput
	// A name for the subscription filter
	DestinationName pulumi.StringInput
}

func (LogDestinationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logDestinationPolicyArgs)(nil)).Elem()
}

type LogDestinationPolicyInput interface {
	pulumi.Input

	ToLogDestinationPolicyOutput() LogDestinationPolicyOutput
	ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput
}

func (LogDestinationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDestinationPolicy)(nil)).Elem()
}

func (i LogDestinationPolicy) ToLogDestinationPolicyOutput() LogDestinationPolicyOutput {
	return i.ToLogDestinationPolicyOutputWithContext(context.Background())
}

func (i LogDestinationPolicy) ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogDestinationPolicyOutput)
}

type LogDestinationPolicyOutput struct {
	*pulumi.OutputState
}

func (LogDestinationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogDestinationPolicyOutput)(nil)).Elem()
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyOutput() LogDestinationPolicyOutput {
	return o
}

func (o LogDestinationPolicyOutput) ToLogDestinationPolicyOutputWithContext(ctx context.Context) LogDestinationPolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LogDestinationPolicyOutput{})
}
