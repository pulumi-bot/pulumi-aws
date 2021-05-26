// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `waf.IpSet` Retrieves a WAF IP Set Resource Id.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/waf"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := waf.GetIpset(ctx, &waf.GetIpsetArgs{
// 			Name: "tfWAFIPSet",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetIpset(ctx *pulumi.Context, args *GetIpsetArgs, opts ...pulumi.InvokeOption) (*GetIpsetResult, error) {
	var rv GetIpsetResult
	err := ctx.Invoke("aws:waf/getIpset:getIpset", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpset.
type GetIpsetArgs struct {
	// The name of the WAF IP set.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIpset.
type GetIpsetResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func GetIpsetApply(ctx *pulumi.Context, args GetIpsetApplyInput, opts ...pulumi.InvokeOption) GetIpsetResultOutput {
	return args.ToGetIpsetApplyOutput().ApplyT(func(v GetIpsetArgs) (GetIpsetResult, error) {
		r, err := GetIpset(ctx, &v, opts...)
		return *r, err

	}).(GetIpsetResultOutput)
}

// GetIpsetApplyInput is an input type that accepts GetIpsetApplyArgs and GetIpsetApplyOutput values.
// You can construct a concrete instance of `GetIpsetApplyInput` via:
//
//          GetIpsetApplyArgs{...}
type GetIpsetApplyInput interface {
	pulumi.Input

	ToGetIpsetApplyOutput() GetIpsetApplyOutput
	ToGetIpsetApplyOutputWithContext(context.Context) GetIpsetApplyOutput
}

// A collection of arguments for invoking getIpset.
type GetIpsetApplyArgs struct {
	// The name of the WAF IP set.
	Name pulumi.StringInput `pulumi:"name"`
}

func (GetIpsetApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsetArgs)(nil)).Elem()
}

func (i GetIpsetApplyArgs) ToGetIpsetApplyOutput() GetIpsetApplyOutput {
	return i.ToGetIpsetApplyOutputWithContext(context.Background())
}

func (i GetIpsetApplyArgs) ToGetIpsetApplyOutputWithContext(ctx context.Context) GetIpsetApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetIpsetApplyOutput)
}

// A collection of arguments for invoking getIpset.
type GetIpsetApplyOutput struct{ *pulumi.OutputState }

func (GetIpsetApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsetArgs)(nil)).Elem()
}

func (o GetIpsetApplyOutput) ToGetIpsetApplyOutput() GetIpsetApplyOutput {
	return o
}

func (o GetIpsetApplyOutput) ToGetIpsetApplyOutputWithContext(ctx context.Context) GetIpsetApplyOutput {
	return o
}

// The name of the WAF IP set.
func (o GetIpsetApplyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsetArgs) string { return v.Name }).(pulumi.StringOutput)
}

// A collection of values returned by getIpset.
type GetIpsetResultOutput struct{ *pulumi.OutputState }

func (GetIpsetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpsetResult)(nil)).Elem()
}

func (o GetIpsetResultOutput) ToGetIpsetResultOutput() GetIpsetResultOutput {
	return o
}

func (o GetIpsetResultOutput) ToGetIpsetResultOutputWithContext(ctx context.Context) GetIpsetResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpsetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIpsetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpsetResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpsetApplyOutput{})
	pulumi.RegisterOutputType(GetIpsetResultOutput{})
}
