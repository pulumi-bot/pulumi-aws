// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package lambda

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an asynchronous invocation configuration for a Lambda Function or Alias. More information about asynchronous invocations and the configurable values can be found in the [Lambda Developer Guide](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html).
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_function_event_invoke_config.html.markdown.
type FunctionEventInvokeConfig struct {
	pulumi.CustomResourceState

	// Configuration block with destination configuration. See below for details.
	DestinationConfig FunctionEventInvokeConfigDestinationConfigPtrOutput `pulumi:"destinationConfig"`
	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds pulumi.IntPtrOutput `pulumi:"maximumEventAgeInSeconds"`
	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts pulumi.IntPtrOutput `pulumi:"maximumRetryAttempts"`
	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier pulumi.StringPtrOutput `pulumi:"qualifier"`
}

// NewFunctionEventInvokeConfig registers a new resource with the given unique name, arguments, and options.
func NewFunctionEventInvokeConfig(ctx *pulumi.Context,
	name string, args *FunctionEventInvokeConfigArgs, opts ...pulumi.ResourceOption) (*FunctionEventInvokeConfig, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil {
		args = &FunctionEventInvokeConfigArgs{}
	}
	var resource FunctionEventInvokeConfig
	err := ctx.RegisterResource("aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFunctionEventInvokeConfig gets an existing FunctionEventInvokeConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFunctionEventInvokeConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FunctionEventInvokeConfigState, opts ...pulumi.ResourceOption) (*FunctionEventInvokeConfig, error) {
	var resource FunctionEventInvokeConfig
	err := ctx.ReadResource("aws:lambda/functionEventInvokeConfig:FunctionEventInvokeConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FunctionEventInvokeConfig resources.
type functionEventInvokeConfigState struct {
	// Configuration block with destination configuration. See below for details.
	DestinationConfig *FunctionEventInvokeConfigDestinationConfig `pulumi:"destinationConfig"`
	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName *string `pulumi:"functionName"`
	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds *int `pulumi:"maximumEventAgeInSeconds"`
	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts *int `pulumi:"maximumRetryAttempts"`
	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier *string `pulumi:"qualifier"`
}

type FunctionEventInvokeConfigState struct {
	// Configuration block with destination configuration. See below for details.
	DestinationConfig FunctionEventInvokeConfigDestinationConfigPtrInput
	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName pulumi.StringPtrInput
	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds pulumi.IntPtrInput
	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts pulumi.IntPtrInput
	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier pulumi.StringPtrInput
}

func (FunctionEventInvokeConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*functionEventInvokeConfigState)(nil)).Elem()
}

type functionEventInvokeConfigArgs struct {
	// Configuration block with destination configuration. See below for details.
	DestinationConfig *FunctionEventInvokeConfigDestinationConfig `pulumi:"destinationConfig"`
	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName string `pulumi:"functionName"`
	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds *int `pulumi:"maximumEventAgeInSeconds"`
	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts *int `pulumi:"maximumRetryAttempts"`
	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier *string `pulumi:"qualifier"`
}

// The set of arguments for constructing a FunctionEventInvokeConfig resource.
type FunctionEventInvokeConfigArgs struct {
	// Configuration block with destination configuration. See below for details.
	DestinationConfig FunctionEventInvokeConfigDestinationConfigPtrInput
	// Name or Amazon Resource Name (ARN) of the Lambda Function, omitting any version or alias qualifier.
	FunctionName pulumi.StringInput
	// Maximum age of a request that Lambda sends to a function for processing in seconds. Valid values between 60 and 21600.
	MaximumEventAgeInSeconds pulumi.IntPtrInput
	// Maximum number of times to retry when the function returns an error. Valid values between 0 and 2. Defaults to 2.
	MaximumRetryAttempts pulumi.IntPtrInput
	// Lambda Function published version, `$LATEST`, or Lambda Alias name.
	Qualifier pulumi.StringPtrInput
}

func (FunctionEventInvokeConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*functionEventInvokeConfigArgs)(nil)).Elem()
}

