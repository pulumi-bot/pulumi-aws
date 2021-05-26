// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// The following example below creates a CloudFront cache policy.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example-policy"
// 		_, err := cloudfront.LookupCachePolicy(ctx, &cloudfront.LookupCachePolicyArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCachePolicy(ctx *pulumi.Context, args *LookupCachePolicyArgs, opts ...pulumi.InvokeOption) (*LookupCachePolicyResult, error) {
	var rv LookupCachePolicyResult
	err := ctx.Invoke("aws:cloudfront/getCachePolicy:getCachePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCachePolicy.
type LookupCachePolicyArgs struct {
	// The identifier for the cache policy.
	Id *string `pulumi:"id"`
	// A unique name to identify the cache policy.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCachePolicy.
type LookupCachePolicyResult struct {
	// A comment to describe the cache policy.
	Comment string `pulumi:"comment"`
	// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	DefaultTtl int `pulumi:"defaultTtl"`
	// The current version of the cache policy.
	Etag string  `pulumi:"etag"`
	Id   *string `pulumi:"id"`
	// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MaxTtl int `pulumi:"maxTtl"`
	// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
	MinTtl int     `pulumi:"minTtl"`
	Name   *string `pulumi:"name"`
	// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
	ParametersInCacheKeyAndForwardedToOrigins []GetCachePolicyParametersInCacheKeyAndForwardedToOrigin `pulumi:"parametersInCacheKeyAndForwardedToOrigins"`
}

func LookupCachePolicyApply(ctx *pulumi.Context, args LookupCachePolicyApplyInput, opts ...pulumi.InvokeOption) LookupCachePolicyResultOutput {
	return args.ToLookupCachePolicyApplyOutput().ApplyT(func(v LookupCachePolicyArgs) (LookupCachePolicyResult, error) {
		r, err := LookupCachePolicy(ctx, &v, opts...)
		return *r, err

	}).(LookupCachePolicyResultOutput)
}

// LookupCachePolicyApplyInput is an input type that accepts LookupCachePolicyApplyArgs and LookupCachePolicyApplyOutput values.
// You can construct a concrete instance of `LookupCachePolicyApplyInput` via:
//
//          LookupCachePolicyApplyArgs{...}
type LookupCachePolicyApplyInput interface {
	pulumi.Input

	ToLookupCachePolicyApplyOutput() LookupCachePolicyApplyOutput
	ToLookupCachePolicyApplyOutputWithContext(context.Context) LookupCachePolicyApplyOutput
}

// A collection of arguments for invoking getCachePolicy.
type LookupCachePolicyApplyArgs struct {
	// The identifier for the cache policy.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// A unique name to identify the cache policy.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupCachePolicyApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCachePolicyArgs)(nil)).Elem()
}

func (i LookupCachePolicyApplyArgs) ToLookupCachePolicyApplyOutput() LookupCachePolicyApplyOutput {
	return i.ToLookupCachePolicyApplyOutputWithContext(context.Background())
}

func (i LookupCachePolicyApplyArgs) ToLookupCachePolicyApplyOutputWithContext(ctx context.Context) LookupCachePolicyApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupCachePolicyApplyOutput)
}

// A collection of arguments for invoking getCachePolicy.
type LookupCachePolicyApplyOutput struct{ *pulumi.OutputState }

func (LookupCachePolicyApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCachePolicyArgs)(nil)).Elem()
}

func (o LookupCachePolicyApplyOutput) ToLookupCachePolicyApplyOutput() LookupCachePolicyApplyOutput {
	return o
}

func (o LookupCachePolicyApplyOutput) ToLookupCachePolicyApplyOutputWithContext(ctx context.Context) LookupCachePolicyApplyOutput {
	return o
}

// The identifier for the cache policy.
func (o LookupCachePolicyApplyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyArgs) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// A unique name to identify the cache policy.
func (o LookupCachePolicyApplyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyArgs) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getCachePolicy.
type LookupCachePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupCachePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCachePolicyResult)(nil)).Elem()
}

func (o LookupCachePolicyResultOutput) ToLookupCachePolicyResultOutput() LookupCachePolicyResultOutput {
	return o
}

func (o LookupCachePolicyResultOutput) ToLookupCachePolicyResultOutputWithContext(ctx context.Context) LookupCachePolicyResultOutput {
	return o
}

// A comment to describe the cache policy.
func (o LookupCachePolicyResultOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) string { return v.Comment }).(pulumi.StringOutput)
}

// The default amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o LookupCachePolicyResultOutput) DefaultTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) int { return v.DefaultTtl }).(pulumi.IntOutput)
}

// The current version of the cache policy.
func (o LookupCachePolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupCachePolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The maximum amount of time, in seconds, that objects stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o LookupCachePolicyResultOutput) MaxTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) int { return v.MaxTtl }).(pulumi.IntOutput)
}

// The minimum amount of time, in seconds, that you want objects to stay in the CloudFront cache before CloudFront sends another request to the origin to see if the object has been updated.
func (o LookupCachePolicyResultOutput) MinTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) int { return v.MinTtl }).(pulumi.IntOutput)
}

func (o LookupCachePolicyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The HTTP headers, cookies, and URL query strings to include in the cache key. See Parameters In Cache Key And Forwarded To Origin for more information.
func (o LookupCachePolicyResultOutput) ParametersInCacheKeyAndForwardedToOrigins() GetCachePolicyParametersInCacheKeyAndForwardedToOriginArrayOutput {
	return o.ApplyT(func(v LookupCachePolicyResult) []GetCachePolicyParametersInCacheKeyAndForwardedToOrigin {
		return v.ParametersInCacheKeyAndForwardedToOrigins
	}).(GetCachePolicyParametersInCacheKeyAndForwardedToOriginArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCachePolicyApplyOutput{})
	pulumi.RegisterOutputType(LookupCachePolicyResultOutput{})
}
