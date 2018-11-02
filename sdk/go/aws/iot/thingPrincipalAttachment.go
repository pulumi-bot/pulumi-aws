// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Attaches Principal to AWS IoT Thing.
type ThingPrincipalAttachment struct {
	s *pulumi.ResourceState
}

// NewThingPrincipalAttachment registers a new resource with the given unique name, arguments, and options.
func NewThingPrincipalAttachment(ctx *pulumi.Context,
	name string, args *ThingPrincipalAttachmentArgs, opts ...pulumi.ResourceOpt) (*ThingPrincipalAttachment, error) {
	if args == nil || args.Principal == nil {
		return nil, errors.New("missing required argument 'Principal'")
	}
	if args == nil || args.Thing == nil {
		return nil, errors.New("missing required argument 'Thing'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["principal"] = nil
		inputs["thing"] = nil
	} else {
		inputs["principal"] = args.Principal
		inputs["thing"] = args.Thing
	}
	s, err := ctx.RegisterResource("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThingPrincipalAttachment{s: s}, nil
}

// GetThingPrincipalAttachment gets an existing ThingPrincipalAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThingPrincipalAttachment(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ThingPrincipalAttachmentState, opts ...pulumi.ResourceOpt) (*ThingPrincipalAttachment, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["principal"] = state.Principal
		inputs["thing"] = state.Thing
	}
	s, err := ctx.ReadResource("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &ThingPrincipalAttachment{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *ThingPrincipalAttachment) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *ThingPrincipalAttachment) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
func (r *ThingPrincipalAttachment) Principal() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["principal"])
}

// The name of the thing.
func (r *ThingPrincipalAttachment) Thing() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["thing"])
}

// Input properties used for looking up and filtering ThingPrincipalAttachment resources.
type ThingPrincipalAttachmentState struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal interface{}
	// The name of the thing.
	Thing interface{}
}

// The set of arguments for constructing a ThingPrincipalAttachment resource.
type ThingPrincipalAttachmentArgs struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal interface{}
	// The name of the thing.
	Thing interface{}
}
