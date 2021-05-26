// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Parameter data source.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.LookupParameter(ctx, &ssm.LookupParameterArgs{
// 			Name: "foo",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// > **Note:** The data source is currently following the behavior of the [SSM API](https://docs.aws.amazon.com/sdk-for-go/api/service/ssm/#Parameter) to return a string value, regardless of parameter type.
func LookupParameter(ctx *pulumi.Context, args *LookupParameterArgs, opts ...pulumi.InvokeOption) (*LookupParameterResult, error) {
	var rv LookupParameterResult
	err := ctx.Invoke("aws:ssm/getParameter:getParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameter.
type LookupParameterArgs struct {
	// The name of the parameter.
	Name string `pulumi:"name"`
	// Whether to return decrypted `SecureString` value. Defaults to `true`.
	WithDecryption *bool `pulumi:"withDecryption"`
}

// A collection of values returned by getParameter.
type LookupParameterResult struct {
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Type           string `pulumi:"type"`
	Value          string `pulumi:"value"`
	Version        int    `pulumi:"version"`
	WithDecryption *bool  `pulumi:"withDecryption"`
}

func LookupParameterApply(ctx *pulumi.Context, args LookupParameterApplyInput, opts ...pulumi.InvokeOption) LookupParameterResultOutput {
	return args.ToLookupParameterApplyOutput().ApplyT(func(v LookupParameterArgs) (LookupParameterResult, error) {
		r, err := LookupParameter(ctx, &v, opts...)
		return *r, err

	}).(LookupParameterResultOutput)
}

// LookupParameterApplyInput is an input type that accepts LookupParameterApplyArgs and LookupParameterApplyOutput values.
// You can construct a concrete instance of `LookupParameterApplyInput` via:
//
//          LookupParameterApplyArgs{...}
type LookupParameterApplyInput interface {
	pulumi.Input

	ToLookupParameterApplyOutput() LookupParameterApplyOutput
	ToLookupParameterApplyOutputWithContext(context.Context) LookupParameterApplyOutput
}

// A collection of arguments for invoking getParameter.
type LookupParameterApplyArgs struct {
	// The name of the parameter.
	Name pulumi.StringInput `pulumi:"name"`
	// Whether to return decrypted `SecureString` value. Defaults to `true`.
	WithDecryption pulumi.BoolPtrInput `pulumi:"withDecryption"`
}

func (LookupParameterApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterArgs)(nil)).Elem()
}

func (i LookupParameterApplyArgs) ToLookupParameterApplyOutput() LookupParameterApplyOutput {
	return i.ToLookupParameterApplyOutputWithContext(context.Background())
}

func (i LookupParameterApplyArgs) ToLookupParameterApplyOutputWithContext(ctx context.Context) LookupParameterApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupParameterApplyOutput)
}

// A collection of arguments for invoking getParameter.
type LookupParameterApplyOutput struct{ *pulumi.OutputState }

func (LookupParameterApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterArgs)(nil)).Elem()
}

func (o LookupParameterApplyOutput) ToLookupParameterApplyOutput() LookupParameterApplyOutput {
	return o
}

func (o LookupParameterApplyOutput) ToLookupParameterApplyOutputWithContext(ctx context.Context) LookupParameterApplyOutput {
	return o
}

// The name of the parameter.
func (o LookupParameterApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterArgs) string { return v.Name }).(pulumi.StringOutput)
}

// Whether to return decrypted `SecureString` value. Defaults to `true`.
func (o LookupParameterApplyOutput) WithDecryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupParameterArgs) *bool { return v.WithDecryption }).(pulumi.BoolPtrOutput)
}

// A collection of values returned by getParameter.
type LookupParameterResultOutput struct{ *pulumi.OutputState }

func (LookupParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterResult)(nil)).Elem()
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutput() LookupParameterResultOutput {
	return o
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutputWithContext(ctx context.Context) LookupParameterResultOutput {
	return o
}

func (o LookupParameterResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupParameterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupParameterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupParameterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupParameterResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterResult) string { return v.Value }).(pulumi.StringOutput)
}

func (o LookupParameterResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupParameterResult) int { return v.Version }).(pulumi.IntOutput)
}

func (o LookupParameterResultOutput) WithDecryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *bool { return v.WithDecryption }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterApplyOutput{})
	pulumi.RegisterOutputType(LookupParameterResultOutput{})
}
