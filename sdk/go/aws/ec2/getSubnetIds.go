// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.getSubnetIds` provides a set of ids for a vpcId
//
// This resource can be useful for getting back a set of subnet ids for a vpc.
func GetSubnetIds(ctx *pulumi.Context, args *GetSubnetIdsArgs, opts ...pulumi.InvokeOption) (*GetSubnetIdsResult, error) {
	var rv GetSubnetIdsResult
	err := ctx.Invoke("aws:ec2/getSubnetIds:getSubnetIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsArgs struct {
	// Custom filter block as described below.
	Filters []GetSubnetIdsFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID that you want to filter from.
	VpcId string `pulumi:"vpcId"`
}

// A collection of values returned by getSubnetIds.
type GetSubnetIdsResult struct {
	Filters []GetSubnetIdsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of all the subnet ids found. This data source will fail if none are found.
	Ids   []string          `pulumi:"ids"`
	Tags  map[string]string `pulumi:"tags"`
	VpcId string            `pulumi:"vpcId"`
}

func GetSubnetIdsApply(ctx *pulumi.Context, args GetSubnetIdsApplyInput, opts ...pulumi.InvokeOption) GetSubnetIdsResultOutput {
	return args.ToGetSubnetIdsApplyOutput().ApplyT(func(v GetSubnetIdsArgs) (GetSubnetIdsResult, error) {
		r, err := GetSubnetIds(ctx, &v, opts...)
		return *r, err

	}).(GetSubnetIdsResultOutput)
}

// GetSubnetIdsApplyInput is an input type that accepts GetSubnetIdsApplyArgs and GetSubnetIdsApplyOutput values.
// You can construct a concrete instance of `GetSubnetIdsApplyInput` via:
//
//          GetSubnetIdsApplyArgs{...}
type GetSubnetIdsApplyInput interface {
	pulumi.Input

	ToGetSubnetIdsApplyOutput() GetSubnetIdsApplyOutput
	ToGetSubnetIdsApplyOutputWithContext(context.Context) GetSubnetIdsApplyOutput
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsApplyArgs struct {
	// Custom filter block as described below.
	Filters GetSubnetIdsFilterArrayInput `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The VPC ID that you want to filter from.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (GetSubnetIdsApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetIdsArgs)(nil)).Elem()
}

func (i GetSubnetIdsApplyArgs) ToGetSubnetIdsApplyOutput() GetSubnetIdsApplyOutput {
	return i.ToGetSubnetIdsApplyOutputWithContext(context.Background())
}

func (i GetSubnetIdsApplyArgs) ToGetSubnetIdsApplyOutputWithContext(ctx context.Context) GetSubnetIdsApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetSubnetIdsApplyOutput)
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsApplyOutput struct{ *pulumi.OutputState }

func (GetSubnetIdsApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetIdsArgs)(nil)).Elem()
}

func (o GetSubnetIdsApplyOutput) ToGetSubnetIdsApplyOutput() GetSubnetIdsApplyOutput {
	return o
}

func (o GetSubnetIdsApplyOutput) ToGetSubnetIdsApplyOutputWithContext(ctx context.Context) GetSubnetIdsApplyOutput {
	return o
}

// Custom filter block as described below.
func (o GetSubnetIdsApplyOutput) Filters() GetSubnetIdsFilterArrayOutput {
	return o.ApplyT(func(v GetSubnetIdsArgs) []GetSubnetIdsFilter { return v.Filters }).(GetSubnetIdsFilterArrayOutput)
}

// A map of tags, each pair of which must exactly match
// a pair on the desired subnets.
func (o GetSubnetIdsApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetIdsArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The VPC ID that you want to filter from.
func (o GetSubnetIdsApplyOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetIdsArgs) string { return v.VpcId }).(pulumi.StringOutput)
}

// A collection of values returned by getSubnetIds.
type GetSubnetIdsResultOutput struct{ *pulumi.OutputState }

func (GetSubnetIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSubnetIdsResult)(nil)).Elem()
}

func (o GetSubnetIdsResultOutput) ToGetSubnetIdsResultOutput() GetSubnetIdsResultOutput {
	return o
}

func (o GetSubnetIdsResultOutput) ToGetSubnetIdsResultOutputWithContext(ctx context.Context) GetSubnetIdsResultOutput {
	return o
}

func (o GetSubnetIdsResultOutput) Filters() GetSubnetIdsFilterArrayOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) []GetSubnetIdsFilter { return v.Filters }).(GetSubnetIdsFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSubnetIdsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of all the subnet ids found. This data source will fail if none are found.
func (o GetSubnetIdsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetSubnetIdsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GetSubnetIdsResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v GetSubnetIdsResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSubnetIdsApplyOutput{})
	pulumi.RegisterOutputType(GetSubnetIdsResultOutput{})
}
