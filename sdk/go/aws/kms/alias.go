// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an alias for a KMS customer master key. AWS Console enforces 1-to-1 mapping between aliases & keys,
// but API (hence this provider too) allows you to create as many aliases as
// the [account limits](http://docs.aws.amazon.com/kms/latest/developerguide/limits.html) allow you.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_alias.html.markdown.
type Alias struct {
	s *pulumi.ResourceState
}

// NewAlias registers a new resource with the given unique name, arguments, and options.
func NewAlias(ctx *pulumi.Context,
	name string, args *AliasArgs, opts ...pulumi.ResourceOpt) (*Alias, error) {
	if args == nil || args.TargetKeyId == nil {
		return nil, errors.New("missing required argument 'TargetKeyId'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["namePrefix"] = nil
		inputs["targetKeyId"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["namePrefix"] = args.NamePrefix
		inputs["targetKeyId"] = args.TargetKeyId
	}
	inputs["arn"] = nil
	inputs["targetKeyArn"] = nil
	s, err := ctx.RegisterResource("aws:kms/alias:Alias", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Alias{s: s}, nil
}

// GetAlias gets an existing Alias resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlias(ctx *pulumi.Context,
	name string, id pulumi.ID, state *AliasState, opts ...pulumi.ResourceOpt) (*Alias, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["arn"] = state.Arn
		inputs["name"] = state.Name
		inputs["namePrefix"] = state.NamePrefix
		inputs["targetKeyArn"] = state.TargetKeyArn
		inputs["targetKeyId"] = state.TargetKeyId
	}
	s, err := ctx.ReadResource("aws:kms/alias:Alias", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Alias{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Alias) URN() pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Alias) ID() pulumi.IDOutput {
	return r.s.ID()
}

// The Amazon Resource Name (ARN) of the key alias.
func (r *Alias) Arn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["arn"])
}

// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
func (r *Alias) Name() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["name"])
}

// Creates an unique alias beginning with the specified prefix.
// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
func (r *Alias) NamePrefix() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["namePrefix"])
}

// The Amazon Resource Name (ARN) of the target key identifier.
func (r *Alias) TargetKeyArn() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["targetKeyArn"])
}

// Identifier for the key for which the alias is for, can be either an ARN or key_id.
func (r *Alias) TargetKeyId() pulumi.StringOutput {
	return (pulumi.StringOutput)(r.s.State["targetKeyId"])
}

// Input properties used for looking up and filtering Alias resources.
type AliasState struct {
	// The Amazon Resource Name (ARN) of the key alias.
	Arn interface{}
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name interface{}
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix interface{}
	// The Amazon Resource Name (ARN) of the target key identifier.
	TargetKeyArn interface{}
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId interface{}
}

// The set of arguments for constructing a Alias resource.
type AliasArgs struct {
	// The display name of the alias. The name must start with the word "alias" followed by a forward slash (alias/)
	Name interface{}
	// Creates an unique alias beginning with the specified prefix.
	// The name must start with the word "alias" followed by a forward slash (alias/).  Conflicts with `name`.
	NamePrefix interface{}
	// Identifier for the key for which the alias is for, can be either an ARN or key_id.
	TargetKeyId interface{}
}
