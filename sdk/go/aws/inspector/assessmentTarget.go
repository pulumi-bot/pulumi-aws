// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Inspector assessment target
type AssessmentTarget struct {
	s *pulumi.ResourceState
}

// NewAssessmentTarget registers a new resource with the given unique name, arguments, and options.
func NewAssessmentTarget(ctx *pulumi.Context,
	name string, args *AssessmentTargetArgs, opts ...pulumi.ResourceOpt) (*AssessmentTarget, error) {
	if args == nil || args.ResourceGroupArn == nil {
		return nil, errors.New("missing required argument 'ResourceGroupArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["resourceGroupArn"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["resourceGroupArn"] = args.ResourceGroupArn
	}
	inputs["arn"] = nil
	s, err := ctx.RegisterResource("aws:inspector/assessmentTarget:AssessmentTarget", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AssessmentTarget{s: s}, nil
}

// GetAssessmentTarget gets an existing AssessmentTarget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssessmentTarget(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AssessmentTargetState, opts ...pulumi.ResourceOpt) (*AssessmentTarget, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["resourceGroupArn"] = state.ResourceGroupArn
	}
	s, err := ctx.ReadResource("aws:inspector/assessmentTarget:AssessmentTarget", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &AssessmentTarget{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *AssessmentTarget) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *AssessmentTarget) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The target assessment ARN.
func (r *AssessmentTarget) Arn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["arn"])
}

// The name of the assessment target.
// * `resource_group_arn` (Required )- The resource group ARN stating tags for instance matching.
func (r *AssessmentTarget) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *AssessmentTarget) ResourceGroupArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupArn"])
}

// Input properties used for looking up and filtering AssessmentTarget resources.
type AssessmentTargetState struct {
	// The target assessment ARN.
	Arn interface{}
	// The name of the assessment target.
	// * `resource_group_arn` (Required )- The resource group ARN stating tags for instance matching.
	Name interface{}
	ResourceGroupArn interface{}
}

// The set of arguments for constructing a AssessmentTarget resource.
type AssessmentTargetArgs struct {
	// The name of the assessment target.
	// * `resource_group_arn` (Required )- The resource group ARN stating tags for instance matching.
	Name interface{}
	ResourceGroupArn interface{}
}
