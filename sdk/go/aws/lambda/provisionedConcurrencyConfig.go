// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Lambda Provisioned Concurrency Configuration.
//
// ## Example Usage
type ProvisionedConcurrencyConfig struct {
	pulumi.CustomResourceState

	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName pulumi.StringOutput `pulumi:"functionName"`
	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions pulumi.IntOutput `pulumi:"provisionedConcurrentExecutions"`
	// Lambda Function version or Lambda Alias name.
	Qualifier pulumi.StringOutput `pulumi:"qualifier"`
}

// NewProvisionedConcurrencyConfig registers a new resource with the given unique name, arguments, and options.
func NewProvisionedConcurrencyConfig(ctx *pulumi.Context,
	name string, args *ProvisionedConcurrencyConfigArgs, opts ...pulumi.ResourceOption) (*ProvisionedConcurrencyConfig, error) {
	if args == nil || args.FunctionName == nil {
		return nil, errors.New("missing required argument 'FunctionName'")
	}
	if args == nil || args.ProvisionedConcurrentExecutions == nil {
		return nil, errors.New("missing required argument 'ProvisionedConcurrentExecutions'")
	}
	if args == nil || args.Qualifier == nil {
		return nil, errors.New("missing required argument 'Qualifier'")
	}
	if args == nil {
		args = &ProvisionedConcurrencyConfigArgs{}
	}
	var resource ProvisionedConcurrencyConfig
	err := ctx.RegisterResource("aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProvisionedConcurrencyConfig gets an existing ProvisionedConcurrencyConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProvisionedConcurrencyConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProvisionedConcurrencyConfigState, opts ...pulumi.ResourceOption) (*ProvisionedConcurrencyConfig, error) {
	var resource ProvisionedConcurrencyConfig
	err := ctx.ReadResource("aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProvisionedConcurrencyConfig resources.
type provisionedConcurrencyConfigState struct {
	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName *string `pulumi:"functionName"`
	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions *int `pulumi:"provisionedConcurrentExecutions"`
	// Lambda Function version or Lambda Alias name.
	Qualifier *string `pulumi:"qualifier"`
}

type ProvisionedConcurrencyConfigState struct {
	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName pulumi.StringPtrInput
	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions pulumi.IntPtrInput
	// Lambda Function version or Lambda Alias name.
	Qualifier pulumi.StringPtrInput
}

func (ProvisionedConcurrencyConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedConcurrencyConfigState)(nil)).Elem()
}

type provisionedConcurrencyConfigArgs struct {
	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName string `pulumi:"functionName"`
	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions int `pulumi:"provisionedConcurrentExecutions"`
	// Lambda Function version or Lambda Alias name.
	Qualifier string `pulumi:"qualifier"`
}

// The set of arguments for constructing a ProvisionedConcurrencyConfig resource.
type ProvisionedConcurrencyConfigArgs struct {
	// Name or Amazon Resource Name (ARN) of the Lambda Function.
	FunctionName pulumi.StringInput
	// Amount of capacity to allocate. Must be greater than or equal to `1`.
	ProvisionedConcurrentExecutions pulumi.IntInput
	// Lambda Function version or Lambda Alias name.
	Qualifier pulumi.StringInput
}

func (ProvisionedConcurrencyConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*provisionedConcurrencyConfigArgs)(nil)).Elem()
}
