// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IAM policy.
type Policy struct {
	s *pulumi.ResourceState
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOpt) (*Policy, error) {
	if args == nil || args.Policy == nil {
		return nil, errors.New("missing required argument 'Policy'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["description"] = nil
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["path"] = nil
		inputs["policy"] = nil
	} else {
		inputs["description"] = args.Description
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["path"] = args.Path
		inputs["policy"] = args.Policy
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:iam/policy:Policy", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PolicyState, opts ...pulumi.ResourceOpt) (*Policy, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["description"] = state.Description
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["path"] = state.Path
		inputs["policy"] = state.Policy
	}
	s, err := ctx.ReadResource("aws:iam/policy:Policy", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Policy{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Policy) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Policy) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN assigned by AWS to this policy.
func (r *Policy) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// Description of the IAM policy.
func (r *Policy) Description() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["description"])
}

// The name of the policy. If omitted, Terraform will assign a random, unique name.
func (r *Policy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
func (r *Policy) NamePrefix() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["namePrefix"])
}

// Path in which to create the policy.
// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
func (r *Policy) Path() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["path"])
}

// The policy document. This is a JSON formatted string. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
func (r *Policy) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// The ARN assigned by AWS to this policy.
	Arn interface{}
	// Description of the IAM policy.
	Description interface{}
	// The name of the policy. If omitted, Terraform will assign a random, unique name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path interface{}
	// The policy document. This is a JSON formatted string. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
	Policy interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Description of the IAM policy.
	Description interface{}
	// The name of the policy. If omitted, Terraform will assign a random, unique name.
	Name interface{}
	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix interface{}
	// Path in which to create the policy.
	// See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more information.
	Path interface{}
	// The policy document. This is a JSON formatted string. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html)
	Policy interface{}
}
