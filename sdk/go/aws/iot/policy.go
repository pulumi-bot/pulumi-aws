// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an IoT policy.
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
		inputs["name"] = nil
		inputs["policy"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["policy"] = args.Policy
	}
	inputs["arn"] = nil
	inputs["defaultVersionId"] = nil
	s, err := ctx.RegisterResource("aws:iot/policy:Policy", name, true, inputs, opts...)
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
		inputs["defaultVersionId"] = state.DefaultVersionId
		inputs["name"] = state.Name
		inputs["policy"] = state.Policy
	}
	s, err := ctx.ReadResource("aws:iot/policy:Policy", name, id, inputs, opts...)
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

// The default version of this policy.
func (r *Policy) DefaultVersionId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["defaultVersionId"])
}

// The name of the policy.
func (r *Policy) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The policy document. This is a JSON formatted string. Use the [IoT Developer Guide](http://docs.aws.amazon.com/iot/latest/developerguide/iot-policies.html) for more information on IoT Policies. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
func (r *Policy) Policy() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["policy"])
}

// Input properties used for looking up and filtering Policy resources.
type PolicyState struct {
	// The ARN assigned by AWS to this policy.
	Arn interface{}
	// The default version of this policy.
	DefaultVersionId interface{}
	// The name of the policy.
	Name interface{}
	// The policy document. This is a JSON formatted string. Use the [IoT Developer Guide](http://docs.aws.amazon.com/iot/latest/developerguide/iot-policies.html) for more information on IoT Policies. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// The name of the policy.
	Name interface{}
	// The policy document. This is a JSON formatted string. Use the [IoT Developer Guide](http://docs.aws.amazon.com/iot/latest/developerguide/iot-policies.html) for more information on IoT Policies. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
	Policy interface{}
}
