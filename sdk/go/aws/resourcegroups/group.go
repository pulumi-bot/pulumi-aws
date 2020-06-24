// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package resourcegroups

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Resource Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/resourcegroups"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = resourcegroups.NewGroup(ctx, "test", &resourcegroups.GroupArgs{
// 			ResourceQuery: &resourcegroups.GroupResourceQueryArgs{
// 				Query: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"ResourceTypeFilters\": [\n", "    \"AWS::EC2::Instance\"\n", "  ],\n", "  \"TagFilters\": [\n", "    {\n", "      \"Key\": \"Stage\",\n", "      \"Values\": [\"Test\"]\n", "    }\n", "  ]\n", "}\n", "\n")),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Group struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this resource group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the resource group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `resourceQuery` block. Resource queries are documented below.
	ResourceQuery GroupResourceQueryOutput `pulumi:"resourceQuery"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.ResourceQuery == nil {
		return nil, errors.New("missing required argument 'ResourceQuery'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("aws:resourcegroups/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("aws:resourcegroups/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The ARN assigned by AWS for this resource group.
	Arn *string `pulumi:"arn"`
	// A description of the resource group.
	Description *string `pulumi:"description"`
	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name *string `pulumi:"name"`
	// A `resourceQuery` block. Resource queries are documented below.
	ResourceQuery *GroupResourceQuery `pulumi:"resourceQuery"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

type GroupState struct {
	// The ARN assigned by AWS for this resource group.
	Arn pulumi.StringPtrInput
	// A description of the resource group.
	Description pulumi.StringPtrInput
	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name pulumi.StringPtrInput
	// A `resourceQuery` block. Resource queries are documented below.
	ResourceQuery GroupResourceQueryPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// A description of the resource group.
	Description *string `pulumi:"description"`
	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name *string `pulumi:"name"`
	// A `resourceQuery` block. Resource queries are documented below.
	ResourceQuery GroupResourceQuery `pulumi:"resourceQuery"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// A description of the resource group.
	Description pulumi.StringPtrInput
	// The resource group's name. A resource group name can have a maximum of 127 characters, including letters, numbers, hyphens, dots, and underscores. The name cannot start with `AWS` or `aws`.
	Name pulumi.StringPtrInput
	// A `resourceQuery` block. Resource queries are documented below.
	ResourceQuery GroupResourceQueryInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
