// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an organizational unit.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/organizations_organizational_unit.html.markdown.
type OrganizationalUnit struct {
	s *pulumi.ResourceState
}

// NewOrganizationalUnit registers a new resource with the given unique name, arguments, and options.
func NewOrganizationalUnit(ctx *pulumi.Context,
	name string, args *OrganizationalUnitArgs, opts ...pulumi.ResourceOpt) (*OrganizationalUnit, error) {
	if args == nil || args.ParentId == nil {
		return nil, errors.New("missing required argument 'ParentId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["parentId"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["parentId"] = args.ParentId
	}
	inputs["accounts"] = nil
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationalUnit{s: s}, nil
}

// GetOrganizationalUnit gets an existing OrganizationalUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationalUnit(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrganizationalUnitState, opts ...pulumi.ResourceOpt) (*OrganizationalUnit, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["accounts"] = state.Accounts
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["parentId"] = state.ParentId
	}
	s, err := ctx.ReadResource("aws:organizations/organizationalUnit:OrganizationalUnit", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &OrganizationalUnit{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *OrganizationalUnit) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *OrganizationalUnit) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
func (r *OrganizationalUnit) Accounts() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["accounts"])
}

// ARN of the organizational unit
func (r *OrganizationalUnit) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The name for the organizational unit
func (r *OrganizationalUnit) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// ID of the parent organizational unit, which may be the root
func (r *OrganizationalUnit) ParentId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["parentId"])
}

// Input properties used for looking up and filtering OrganizationalUnit resources.
type OrganizationalUnitState struct {
	// List of child accounts for this Organizational Unit. Does not return account information for child Organizational Units. All elements have these attributes:
	Accounts interface{}
	// ARN of the organizational unit
	Arn interface{}
	// The name for the organizational unit
	Name interface{}
	// ID of the parent organizational unit, which may be the root
	ParentId interface{}
}

// The set of arguments for constructing a OrganizationalUnit resource.
type OrganizationalUnitArgs struct {
	// The name for the organizational unit
	Name interface{}
	// ID of the parent organizational unit, which may be the root
	ParentId interface{}
}
