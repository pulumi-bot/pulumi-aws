// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ebs.getEbsVolumes` provides identifying information for EBS volumes matching given criteria.
//
// This data source can be useful for getting a list of volume IDs with (for example) matching tags.
func GetEbsVolumes(ctx *pulumi.Context, args *GetEbsVolumesArgs, opts ...pulumi.InvokeOption) (*GetEbsVolumesResult, error) {
	var rv GetEbsVolumesResult
	err := ctx.Invoke("aws:ebs/getEbsVolumes:getEbsVolumes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEbsVolumes.
type GetEbsVolumesArgs struct {
	// Custom filter block as described below.
	Filters []GetEbsVolumesFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired volumes.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getEbsVolumes.
type GetEbsVolumesResult struct {
	Filters []GetEbsVolumesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of all the EBS Volume IDs found. This data source will fail if
	// no volumes match the provided criteria.
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}

func GetEbsVolumesApply(ctx *pulumi.Context, args GetEbsVolumesApplyInput, opts ...pulumi.InvokeOption) GetEbsVolumesResultOutput {
	return args.ToGetEbsVolumesApplyOutput().ApplyT(func(v GetEbsVolumesArgs) (GetEbsVolumesResult, error) {
		r, err := GetEbsVolumes(ctx, &v, opts...)
		return *r, err

	}).(GetEbsVolumesResultOutput)
}

// GetEbsVolumesApplyInput is an input type that accepts GetEbsVolumesApplyArgs and GetEbsVolumesApplyOutput values.
// You can construct a concrete instance of `GetEbsVolumesApplyInput` via:
//
//          GetEbsVolumesApplyArgs{...}
type GetEbsVolumesApplyInput interface {
	pulumi.Input

	ToGetEbsVolumesApplyOutput() GetEbsVolumesApplyOutput
	ToGetEbsVolumesApplyOutputWithContext(context.Context) GetEbsVolumesApplyOutput
}

// A collection of arguments for invoking getEbsVolumes.
type GetEbsVolumesApplyArgs struct {
	// Custom filter block as described below.
	Filters GetEbsVolumesFilterArrayInput `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired volumes.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (GetEbsVolumesApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEbsVolumesArgs)(nil)).Elem()
}

func (i GetEbsVolumesApplyArgs) ToGetEbsVolumesApplyOutput() GetEbsVolumesApplyOutput {
	return i.ToGetEbsVolumesApplyOutputWithContext(context.Background())
}

func (i GetEbsVolumesApplyArgs) ToGetEbsVolumesApplyOutputWithContext(ctx context.Context) GetEbsVolumesApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetEbsVolumesApplyOutput)
}

// A collection of arguments for invoking getEbsVolumes.
type GetEbsVolumesApplyOutput struct{ *pulumi.OutputState }

func (GetEbsVolumesApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEbsVolumesArgs)(nil)).Elem()
}

func (o GetEbsVolumesApplyOutput) ToGetEbsVolumesApplyOutput() GetEbsVolumesApplyOutput {
	return o
}

func (o GetEbsVolumesApplyOutput) ToGetEbsVolumesApplyOutputWithContext(ctx context.Context) GetEbsVolumesApplyOutput {
	return o
}

// Custom filter block as described below.
func (o GetEbsVolumesApplyOutput) Filters() GetEbsVolumesFilterArrayOutput {
	return o.ApplyT(func(v GetEbsVolumesArgs) []GetEbsVolumesFilter { return v.Filters }).(GetEbsVolumesFilterArrayOutput)
}

// A map of tags, each pair of which must exactly match
// a pair on the desired volumes.
func (o GetEbsVolumesApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetEbsVolumesArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getEbsVolumes.
type GetEbsVolumesResultOutput struct{ *pulumi.OutputState }

func (GetEbsVolumesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetEbsVolumesResult)(nil)).Elem()
}

func (o GetEbsVolumesResultOutput) ToGetEbsVolumesResultOutput() GetEbsVolumesResultOutput {
	return o
}

func (o GetEbsVolumesResultOutput) ToGetEbsVolumesResultOutputWithContext(ctx context.Context) GetEbsVolumesResultOutput {
	return o
}

func (o GetEbsVolumesResultOutput) Filters() GetEbsVolumesFilterArrayOutput {
	return o.ApplyT(func(v GetEbsVolumesResult) []GetEbsVolumesFilter { return v.Filters }).(GetEbsVolumesFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetEbsVolumesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetEbsVolumesResult) string { return v.Id }).(pulumi.StringOutput)
}

// A set of all the EBS Volume IDs found. This data source will fail if
// no volumes match the provided criteria.
func (o GetEbsVolumesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetEbsVolumesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetEbsVolumesResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetEbsVolumesResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetEbsVolumesApplyOutput{})
	pulumi.RegisterOutputType(GetEbsVolumesResultOutput{})
}
