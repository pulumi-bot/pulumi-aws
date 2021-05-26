// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticbeanstalk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieve information about an Elastic Beanstalk Application.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/elasticbeanstalk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := elasticbeanstalk.LookupApplication(ctx, &elasticbeanstalk.LookupApplicationArgs{
// 			Name: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("arn", example.Arn)
// 		ctx.Export("description", example.Description)
// 		return nil
// 	})
// }
// ```
func LookupApplication(ctx *pulumi.Context, args *LookupApplicationArgs, opts ...pulumi.InvokeOption) (*LookupApplicationResult, error) {
	var rv LookupApplicationResult
	err := ctx.Invoke("aws:elasticbeanstalk/getApplication:getApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplication.
type LookupApplicationArgs struct {
	// The name of the application
	Name string `pulumi:"name"`
}

// A collection of values returned by getApplication.
type LookupApplicationResult struct {
	AppversionLifecycle GetApplicationAppversionLifecycle `pulumi:"appversionLifecycle"`
	// The Amazon Resource Name (ARN) of the application.
	Arn string `pulumi:"arn"`
	// Short description of the application
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupApplicationApply(ctx *pulumi.Context, args LookupApplicationApplyInput, opts ...pulumi.InvokeOption) LookupApplicationResultOutput {
	return args.ToLookupApplicationApplyOutput().ApplyT(func(v LookupApplicationArgs) (LookupApplicationResult, error) {
		r, err := LookupApplication(ctx, &v, opts...)
		return *r, err

	}).(LookupApplicationResultOutput)
}

// LookupApplicationApplyInput is an input type that accepts LookupApplicationApplyArgs and LookupApplicationApplyOutput values.
// You can construct a concrete instance of `LookupApplicationApplyInput` via:
//
//          LookupApplicationApplyArgs{...}
type LookupApplicationApplyInput interface {
	pulumi.Input

	ToLookupApplicationApplyOutput() LookupApplicationApplyOutput
	ToLookupApplicationApplyOutputWithContext(context.Context) LookupApplicationApplyOutput
}

// A collection of arguments for invoking getApplication.
type LookupApplicationApplyArgs struct {
	// The name of the application
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupApplicationApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

func (i LookupApplicationApplyArgs) ToLookupApplicationApplyOutput() LookupApplicationApplyOutput {
	return i.ToLookupApplicationApplyOutputWithContext(context.Background())
}

func (i LookupApplicationApplyArgs) ToLookupApplicationApplyOutputWithContext(ctx context.Context) LookupApplicationApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupApplicationApplyOutput)
}

// A collection of arguments for invoking getApplication.
type LookupApplicationApplyOutput struct{ *pulumi.OutputState }

func (LookupApplicationApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationArgs)(nil)).Elem()
}

func (o LookupApplicationApplyOutput) ToLookupApplicationApplyOutput() LookupApplicationApplyOutput {
	return o
}

func (o LookupApplicationApplyOutput) ToLookupApplicationApplyOutputWithContext(ctx context.Context) LookupApplicationApplyOutput {
	return o
}

// The name of the application
func (o LookupApplicationApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationArgs) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of values returned by getApplication.
type LookupApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationResult)(nil)).Elem()
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutput() LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) ToLookupApplicationResultOutputWithContext(ctx context.Context) LookupApplicationResultOutput {
	return o
}

func (o LookupApplicationResultOutput) AppversionLifecycle() GetApplicationAppversionLifecycleOutput {
	return o.ApplyT(func(v LookupApplicationResult) GetApplicationAppversionLifecycle { return v.AppversionLifecycle }).(GetApplicationAppversionLifecycleOutput)
}

// The Amazon Resource Name (ARN) of the application.
func (o LookupApplicationResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Short description of the application
func (o LookupApplicationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationApplyOutput{})
	pulumi.RegisterOutputType(LookupApplicationResultOutput{})
}
