// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroupstaggingapi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about resource tagging.
//
// ## Example Usage
// ### Get All Resource Tag Mappings
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/resourcegroupstaggingapi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcegroupstaggingapi.GetResources(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter By Tag Key and Value
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/resourcegroupstaggingapi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcegroupstaggingapi.GetResources(ctx, &resourcegroupstaggingapi.GetResourcesArgs{
// 			TagFilters: []resourcegroupstaggingapi.GetResourcesTagFilter{
// 				resourcegroupstaggingapi.GetResourcesTagFilter{
// 					Key: "tag-key",
// 					Values: []string{
// 						"tag-value-1",
// 						"tag-value-2",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Filter By Resource Type
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/resourcegroupstaggingapi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := resourcegroupstaggingapi.GetResources(ctx, &resourcegroupstaggingapi.GetResourcesArgs{
// 			ResourceTypeFilters: []string{
// 				"ec2:instance",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetResources(ctx *pulumi.Context, args *GetResourcesArgs, opts ...pulumi.InvokeOption) (*GetResourcesResult, error) {
	var rv GetResourcesResult
	err := ctx.Invoke("aws:resourcegroupstaggingapi/getResources:getResources", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResources.
type GetResourcesArgs struct {
	// Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `includeComplianceDetails` argument is also set to `true`.
	ExcludeCompliantResources *bool `pulumi:"excludeCompliantResources"`
	// Specifies whether to include details regarding the compliance with the effective tag policy.
	IncludeComplianceDetails *bool `pulumi:"includeComplianceDetails"`
	// Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
	ResourceArnLists []string `pulumi:"resourceArnLists"`
	// The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
	ResourceTypeFilters []string `pulumi:"resourceTypeFilters"`
	// Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resourceArnList`.
	TagFilters []GetResourcesTagFilter `pulumi:"tagFilters"`
}

// A collection of values returned by getResources.
type GetResourcesResult struct {
	ExcludeCompliantResources *bool `pulumi:"excludeCompliantResources"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string   `pulumi:"id"`
	IncludeComplianceDetails *bool    `pulumi:"includeComplianceDetails"`
	ResourceArnLists         []string `pulumi:"resourceArnLists"`
	// List of objects matching the search criteria.
	ResourceTagMappingLists []GetResourcesResourceTagMappingList `pulumi:"resourceTagMappingLists"`
	ResourceTypeFilters     []string                             `pulumi:"resourceTypeFilters"`
	TagFilters              []GetResourcesTagFilter              `pulumi:"tagFilters"`
}

func GetResourcesApply(ctx *pulumi.Context, args GetResourcesApplyInput, opts ...pulumi.InvokeOption) GetResourcesResultOutput {
	return args.ToGetResourcesApplyOutput().ApplyT(func(v GetResourcesArgs) (GetResourcesResult, error) {
		r, err := GetResources(ctx, &v, opts...)
		return *r, err

	}).(GetResourcesResultOutput)
}

// GetResourcesApplyInput is an input type that accepts GetResourcesApplyArgs and GetResourcesApplyOutput values.
// You can construct a concrete instance of `GetResourcesApplyInput` via:
//
//          GetResourcesApplyArgs{...}
type GetResourcesApplyInput interface {
	pulumi.Input

	ToGetResourcesApplyOutput() GetResourcesApplyOutput
	ToGetResourcesApplyOutputWithContext(context.Context) GetResourcesApplyOutput
}

// A collection of arguments for invoking getResources.
type GetResourcesApplyArgs struct {
	// Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `includeComplianceDetails` argument is also set to `true`.
	ExcludeCompliantResources pulumi.BoolPtrInput `pulumi:"excludeCompliantResources"`
	// Specifies whether to include details regarding the compliance with the effective tag policy.
	IncludeComplianceDetails pulumi.BoolPtrInput `pulumi:"includeComplianceDetails"`
	// Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
	ResourceArnLists pulumi.StringArrayInput `pulumi:"resourceArnLists"`
	// The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
	ResourceTypeFilters pulumi.StringArrayInput `pulumi:"resourceTypeFilters"`
	// Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resourceArnList`.
	TagFilters GetResourcesTagFilterArrayInput `pulumi:"tagFilters"`
}

func (GetResourcesApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcesArgs)(nil)).Elem()
}

func (i GetResourcesApplyArgs) ToGetResourcesApplyOutput() GetResourcesApplyOutput {
	return i.ToGetResourcesApplyOutputWithContext(context.Background())
}

func (i GetResourcesApplyArgs) ToGetResourcesApplyOutputWithContext(ctx context.Context) GetResourcesApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetResourcesApplyOutput)
}

// A collection of arguments for invoking getResources.
type GetResourcesApplyOutput struct{ *pulumi.OutputState }

func (GetResourcesApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcesArgs)(nil)).Elem()
}

func (o GetResourcesApplyOutput) ToGetResourcesApplyOutput() GetResourcesApplyOutput {
	return o
}

func (o GetResourcesApplyOutput) ToGetResourcesApplyOutputWithContext(ctx context.Context) GetResourcesApplyOutput {
	return o
}

// Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `includeComplianceDetails` argument is also set to `true`.
func (o GetResourcesApplyOutput) ExcludeCompliantResources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetResourcesArgs) *bool { return v.ExcludeCompliantResources }).(pulumi.BoolPtrOutput)
}

// Specifies whether to include details regarding the compliance with the effective tag policy.
func (o GetResourcesApplyOutput) IncludeComplianceDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetResourcesArgs) *bool { return v.IncludeComplianceDetails }).(pulumi.BoolPtrOutput)
}

// Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
func (o GetResourcesApplyOutput) ResourceArnLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourcesArgs) []string { return v.ResourceArnLists }).(pulumi.StringArrayOutput)
}

// The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
func (o GetResourcesApplyOutput) ResourceTypeFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourcesArgs) []string { return v.ResourceTypeFilters }).(pulumi.StringArrayOutput)
}

// Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resourceArnList`.
func (o GetResourcesApplyOutput) TagFilters() GetResourcesTagFilterArrayOutput {
	return o.ApplyT(func(v GetResourcesArgs) []GetResourcesTagFilter { return v.TagFilters }).(GetResourcesTagFilterArrayOutput)
}

// A collection of values returned by getResources.
type GetResourcesResultOutput struct{ *pulumi.OutputState }

func (GetResourcesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourcesResult)(nil)).Elem()
}

func (o GetResourcesResultOutput) ToGetResourcesResultOutput() GetResourcesResultOutput {
	return o
}

func (o GetResourcesResultOutput) ToGetResourcesResultOutputWithContext(ctx context.Context) GetResourcesResultOutput {
	return o
}

func (o GetResourcesResultOutput) ExcludeCompliantResources() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetResourcesResult) *bool { return v.ExcludeCompliantResources }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourcesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourcesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetResourcesResultOutput) IncludeComplianceDetails() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetResourcesResult) *bool { return v.IncludeComplianceDetails }).(pulumi.BoolPtrOutput)
}

func (o GetResourcesResultOutput) ResourceArnLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourcesResult) []string { return v.ResourceArnLists }).(pulumi.StringArrayOutput)
}

// List of objects matching the search criteria.
func (o GetResourcesResultOutput) ResourceTagMappingLists() GetResourcesResourceTagMappingListArrayOutput {
	return o.ApplyT(func(v GetResourcesResult) []GetResourcesResourceTagMappingList { return v.ResourceTagMappingLists }).(GetResourcesResourceTagMappingListArrayOutput)
}

func (o GetResourcesResultOutput) ResourceTypeFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetResourcesResult) []string { return v.ResourceTypeFilters }).(pulumi.StringArrayOutput)
}

func (o GetResourcesResultOutput) TagFilters() GetResourcesTagFilterArrayOutput {
	return o.ApplyT(func(v GetResourcesResult) []GetResourcesTagFilter { return v.TagFilters }).(GetResourcesTagFilterArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourcesApplyOutput{})
	pulumi.RegisterOutputType(GetResourcesResultOutput{})
}
