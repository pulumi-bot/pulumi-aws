// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create an organizational unit.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewOrganizationalUnit(ctx, "example", &organizations.OrganizationalUnitArgs{
// 			ParentId: pulumi.Any(aws_organizations_organization.Example.Roots[0].Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrganizationalUnit struct {
	pulumi.CustomResourceState

	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts OrganizationalUnitAccountArrayOutput `pulumi:"accounts"`
	// ARN of the organizational unit
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The name for the organizational unit
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringOutput `pulumi:"parentId"`
}

// NewOrganizationalUnit registers a new resource with the given unique name, arguments, and options.
func NewOrganizationalUnit(ctx *pulumi.Context,
	name string, args *OrganizationalUnitArgs, opts ...pulumi.ResourceOption) (*OrganizationalUnit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	var resource OrganizationalUnit
	err := ctx.RegisterResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationalUnit gets an existing OrganizationalUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationalUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationalUnitState, opts ...pulumi.ResourceOption) (*OrganizationalUnit, error) {
	var resource OrganizationalUnit
	err := ctx.ReadResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationalUnit resources.
type organizationalUnitState struct {
	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts []OrganizationalUnitAccount `pulumi:"accounts"`
	// ARN of the organizational unit
	Arn *string `pulumi:"arn"`
	// The name for the organizational unit
	Name *string `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId *string `pulumi:"parentId"`
}

type OrganizationalUnitState struct {
	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts OrganizationalUnitAccountArrayInput
	// ARN of the organizational unit
	Arn pulumi.StringPtrInput
	// The name for the organizational unit
	Name pulumi.StringPtrInput
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringPtrInput
}

func (OrganizationalUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationalUnitState)(nil)).Elem()
}

type organizationalUnitArgs struct {
	// The name for the organizational unit
	Name *string `pulumi:"name"`
	// ID of the parent organizational unit, which may be the root
	ParentId string `pulumi:"parentId"`
}

// The set of arguments for constructing a OrganizationalUnit resource.
type OrganizationalUnitArgs struct {
	// The name for the organizational unit
	Name pulumi.StringPtrInput
	// ID of the parent organizational unit, which may be the root
	ParentId pulumi.StringInput
}

func (OrganizationalUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationalUnitArgs)(nil)).Elem()
}
