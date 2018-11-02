// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package simpledb

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SimpleDB domain resource
type Domain struct {
	s *pulumi.ResourceState
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
	} else {
		inputs["name"] = args.Name
	}
	s, err := ctx.RegisterResource("aws:simpledb/domain:Domain", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DomainState, opts ...pulumi.ResourceOpt) (*Domain, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["name"] = state.Name
	}
	s, err := ctx.ReadResource("aws:simpledb/domain:Domain", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Domain{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Domain) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Domain) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The name of the SimpleDB domain
func (r *Domain) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Input properties used for looking up and filtering Domain resources.
type DomainState struct {
	// The name of the SimpleDB domain
	Name interface{}
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// The name of the SimpleDB domain
	Name interface{}
}
