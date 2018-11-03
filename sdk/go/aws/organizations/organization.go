// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to create an organization.
type Organization struct {
	s *pulumi.ResourceState
}

// NewOrganization registers a new resource with the given unique name, arguments, and options.
func NewOrganization(ctx *pulumi.Context,
	name string, args *OrganizationArgs, opts ...pulumi.ResourceOpt) (*Organization, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["featureSet"] = nil
	} else {
		inputs["featureSet"] = args.FeatureSet
	}
	inputs["arn"] = nil
	inputs["masterAccountArn"] = nil
	inputs["masterAccountEmail"] = nil
	inputs["masterAccountId"] = nil
	s, err := ctx.RegisterResource("aws:organizations/organization:Organization", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Organization{s: s}, nil
}

// GetOrganization gets an existing Organization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganization(ctx *pulumi.Context,
	name string, id pulumi.ID, state *OrganizationState, opts ...pulumi.ResourceOpt) (*Organization, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["featureSet"] = state.FeatureSet
		inputs["masterAccountArn"] = state.MasterAccountArn
		inputs["masterAccountEmail"] = state.MasterAccountEmail
		inputs["masterAccountId"] = state.MasterAccountId
	}
	s, err := ctx.ReadResource("aws:organizations/organization:Organization", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Organization{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Organization) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Organization) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// ARN of the organization
func (r *Organization) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
func (r *Organization) FeatureSet() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["featureSet"])
}

// ARN of the master account
func (r *Organization) MasterAccountArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterAccountArn"])
}

// Email address of the master account
func (r *Organization) MasterAccountEmail() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterAccountEmail"])
}

// Identifier of the master account
func (r *Organization) MasterAccountId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["masterAccountId"])
}

// Input properties used for looking up and filtering Organization resources.
type OrganizationState struct {
	// ARN of the organization
	Arn interface{}
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet interface{}
	// ARN of the master account
	MasterAccountArn interface{}
	// Email address of the master account
	MasterAccountEmail interface{}
	// Identifier of the master account
	MasterAccountId interface{}
}

// The set of arguments for constructing a Organization resource.
type OrganizationArgs struct {
	// Specify "ALL" (default) or "CONSOLIDATED_BILLING".
	FeatureSet interface{}
}
