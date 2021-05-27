// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Outposts Outpost.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/outposts"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example"
// 		_, err := outposts.GetOutpost(ctx, &outposts.GetOutpostArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOutpost(ctx *pulumi.Context, args *GetOutpostArgs, opts ...pulumi.InvokeOption) (*GetOutpostResult, error) {
	var rv GetOutpostResult
	err := ctx.Invoke("aws:outposts/getOutpost:getOutpost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOutpost.
type GetOutpostArgs struct {
	// Amazon Resource Name (ARN).
	Arn *string `pulumi:"arn"`
	// Identifier of the Outpost.
	Id *string `pulumi:"id"`
	// Name of the Outpost.
	Name *string `pulumi:"name"`
	// AWS Account identifier of the Outpost owner.
	OwnerId *string `pulumi:"ownerId"`
}

// A collection of values returned by getOutpost.
type GetOutpostResult struct {
	Arn string `pulumi:"arn"`
	// Availability Zone name.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Availability Zone identifier.
	AvailabilityZoneId string `pulumi:"availabilityZoneId"`
	// Description.
	Description string `pulumi:"description"`
	Id          string `pulumi:"id"`
	Name        string `pulumi:"name"`
	OwnerId     string `pulumi:"ownerId"`
	// Site identifier.
	SiteId string `pulumi:"siteId"`
}

func GetOutpostOutput(ctx *pulumi.Context, args GetOutpostOutputArgs, opts ...pulumi.InvokeOption) GetOutpostResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOutpostResult, error) {
			args := v.(GetOutpostArgs)
			r, err := GetOutpost(ctx, &args, opts...)
			return *r, err
		}).(GetOutpostResultOutput)
}

// A collection of arguments for invoking getOutpost.
type GetOutpostOutputArgs struct {
	// Amazon Resource Name (ARN).
	Arn pulumi.StringPtrInput `pulumi:"arn"`
	// Identifier of the Outpost.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of the Outpost.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// AWS Account identifier of the Outpost owner.
	OwnerId pulumi.StringPtrInput `pulumi:"ownerId"`
}

func (GetOutpostOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOutpostArgs)(nil)).Elem()
}

// A collection of values returned by getOutpost.
type GetOutpostResultOutput struct{ *pulumi.OutputState }

func (GetOutpostResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOutpostResult)(nil)).Elem()
}

func (o GetOutpostResultOutput) ToGetOutpostResultOutput() GetOutpostResultOutput {
	return o
}

func (o GetOutpostResultOutput) ToGetOutpostResultOutputWithContext(ctx context.Context) GetOutpostResultOutput {
	return o
}

func (o GetOutpostResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Availability Zone name.
func (o GetOutpostResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Availability Zone identifier.
func (o GetOutpostResultOutput) AvailabilityZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.AvailabilityZoneId }).(pulumi.StringOutput)
}

// Description.
func (o GetOutpostResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetOutpostResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOutpostResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOutpostResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Site identifier.
func (o GetOutpostResultOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOutpostResult) string { return v.SiteId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOutpostResultOutput{})
}
