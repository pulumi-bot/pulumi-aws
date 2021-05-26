// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get all direct child organizational units under a parent organizational unit. This only provides immediate children, not all children.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		org, err := organizations.LookupOrganization(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.GetOrganizationalUnits(ctx, &organizations.GetOrganizationalUnitsArgs{
// 			ParentId: org.Roots[0].Id,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetOrganizationalUnits(ctx *pulumi.Context, args *GetOrganizationalUnitsArgs, opts ...pulumi.InvokeOption) (*GetOrganizationalUnitsResult, error) {
	var rv GetOrganizationalUnitsResult
	err := ctx.Invoke("aws:organizations/getOrganizationalUnits:getOrganizationalUnits", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrganizationalUnits.
type GetOrganizationalUnitsArgs struct {
	// The parent ID of the organizational unit.
	ParentId string `pulumi:"parentId"`
}

// A collection of values returned by getOrganizationalUnits.
type GetOrganizationalUnitsResult struct {
	// List of child organizational units, which have the following attributes:
	Childrens []GetOrganizationalUnitsChildren `pulumi:"childrens"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	ParentId string `pulumi:"parentId"`
}

func GetOrganizationalUnitsApply(ctx *pulumi.Context, args GetOrganizationalUnitsApplyInput, opts ...pulumi.InvokeOption) GetOrganizationalUnitsResultOutput {
	return args.ToGetOrganizationalUnitsApplyOutput().ApplyT(func(v GetOrganizationalUnitsArgs) (GetOrganizationalUnitsResult, error) {
		r, err := GetOrganizationalUnits(ctx, &v, opts...)
		return *r, err

	}).(GetOrganizationalUnitsResultOutput)
}

// GetOrganizationalUnitsApplyInput is an input type that accepts GetOrganizationalUnitsApplyArgs and GetOrganizationalUnitsApplyOutput values.
// You can construct a concrete instance of `GetOrganizationalUnitsApplyInput` via:
//
//          GetOrganizationalUnitsApplyArgs{...}
type GetOrganizationalUnitsApplyInput interface {
	pulumi.Input

	ToGetOrganizationalUnitsApplyOutput() GetOrganizationalUnitsApplyOutput
	ToGetOrganizationalUnitsApplyOutputWithContext(context.Context) GetOrganizationalUnitsApplyOutput
}

// A collection of arguments for invoking getOrganizationalUnits.
type GetOrganizationalUnitsApplyArgs struct {
	// The parent ID of the organizational unit.
	ParentId pulumi.StringInput `pulumi:"parentId"`
}

func (GetOrganizationalUnitsApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationalUnitsArgs)(nil)).Elem()
}

func (i GetOrganizationalUnitsApplyArgs) ToGetOrganizationalUnitsApplyOutput() GetOrganizationalUnitsApplyOutput {
	return i.ToGetOrganizationalUnitsApplyOutputWithContext(context.Background())
}

func (i GetOrganizationalUnitsApplyArgs) ToGetOrganizationalUnitsApplyOutputWithContext(ctx context.Context) GetOrganizationalUnitsApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetOrganizationalUnitsApplyOutput)
}

// A collection of arguments for invoking getOrganizationalUnits.
type GetOrganizationalUnitsApplyOutput struct{ *pulumi.OutputState }

func (GetOrganizationalUnitsApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationalUnitsArgs)(nil)).Elem()
}

func (o GetOrganizationalUnitsApplyOutput) ToGetOrganizationalUnitsApplyOutput() GetOrganizationalUnitsApplyOutput {
	return o
}

func (o GetOrganizationalUnitsApplyOutput) ToGetOrganizationalUnitsApplyOutputWithContext(ctx context.Context) GetOrganizationalUnitsApplyOutput {
	return o
}

// The parent ID of the organizational unit.
func (o GetOrganizationalUnitsApplyOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationalUnitsArgs) string { return v.ParentId }).(pulumi.StringOutput)
}

// A collection of values returned by getOrganizationalUnits.
type GetOrganizationalUnitsResultOutput struct{ *pulumi.OutputState }

func (GetOrganizationalUnitsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrganizationalUnitsResult)(nil)).Elem()
}

func (o GetOrganizationalUnitsResultOutput) ToGetOrganizationalUnitsResultOutput() GetOrganizationalUnitsResultOutput {
	return o
}

func (o GetOrganizationalUnitsResultOutput) ToGetOrganizationalUnitsResultOutputWithContext(ctx context.Context) GetOrganizationalUnitsResultOutput {
	return o
}

// List of child organizational units, which have the following attributes:
func (o GetOrganizationalUnitsResultOutput) Childrens() GetOrganizationalUnitsChildrenArrayOutput {
	return o.ApplyT(func(v GetOrganizationalUnitsResult) []GetOrganizationalUnitsChildren { return v.Childrens }).(GetOrganizationalUnitsChildrenArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrganizationalUnitsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationalUnitsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOrganizationalUnitsResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrganizationalUnitsResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrganizationalUnitsApplyOutput{})
	pulumi.RegisterOutputType(GetOrganizationalUnitsResultOutput{})
}
