// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves the summary of a WAFv2 Rule Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/wafv2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := wafv2.LookupRuleGroup(ctx, &wafv2.LookupRuleGroupArgs{
// 			Name:  "some-rule-group",
// 			Scope: "REGIONAL",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRuleGroup(ctx *pulumi.Context, args *LookupRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupRuleGroupResult, error) {
	var rv LookupRuleGroupResult
	err := ctx.Invoke("aws:wafv2/getRuleGroup:getRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRuleGroup.
type LookupRuleGroupArgs struct {
	// The name of the WAFv2 Rule Group.
	Name string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
}

// A collection of values returned by getRuleGroup.
type LookupRuleGroupResult struct {
	// The Amazon Resource Name (ARN) of the entity.
	Arn string `pulumi:"arn"`
	// The description of the rule group that helps with identification.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	Scope string `pulumi:"scope"`
}

func LookupRuleGroupOutput(ctx *pulumi.Context, args LookupRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleGroupResult, error) {
			args := v.(LookupRuleGroupArgs)
			r, err := LookupRuleGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupRuleGroupResultOutput)
}

// A collection of arguments for invoking getRuleGroup.
type LookupRuleGroupOutputArgs struct {
	// The name of the WAFv2 Rule Group.
	Name pulumi.StringInput `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope pulumi.StringInput `pulumi:"scope"`
}

func (LookupRuleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupArgs)(nil)).Elem()
}

// A collection of values returned by getRuleGroup.
type LookupRuleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRuleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupResult)(nil)).Elem()
}

func (o LookupRuleGroupResultOutput) ToLookupRuleGroupResultOutput() LookupRuleGroupResultOutput {
	return o
}

func (o LookupRuleGroupResultOutput) ToLookupRuleGroupResultOutputWithContext(ctx context.Context) LookupRuleGroupResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the entity.
func (o LookupRuleGroupResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The description of the rule group that helps with identification.
func (o LookupRuleGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRuleGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRuleGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRuleGroupResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) string { return v.Scope }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleGroupResultOutput{})
}
