// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// The KMS ciphertext resource allows you to encrypt plaintext into ciphertext
// by using an AWS KMS customer master key. The value returned by this resource
// is stable across every apply. For a changing ciphertext value each apply, see
// the [`aws_kms_ciphertext` data source](https://www.terraform.io/docs/providers/aws/d/kms_ciphertext.html).
// 
// > **Note:** All arguments including the plaintext be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/kms_ciphertext.html.markdown.
type Ciphertext struct {
	s *pulumi.ResourceState
}

// NewCiphertext registers a new resource with the given unique name, arguments, and options.
func NewCiphertext(ctx *pulumi.Context,
	name string, args *CiphertextArgs, opts ...pulumi.ResourceOpt) (*Ciphertext, error) {
	if args == nil || args.KeyId == nil {
		return nil, errors.New("missing required argument 'KeyId'")
	}
	if args == nil || args.Plaintext == nil {
		return nil, errors.New("missing required argument 'Plaintext'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["context"] = nil
		inputs["keyId"] = nil
		inputs["plaintext"] = nil
	} else {
		inputs["context"] = args.Context
		inputs["keyId"] = args.KeyId
		inputs["plaintext"] = args.Plaintext
	}
	inputs["ciphertextBlob"] = nil
	s, err := ctx.RegisterResource("aws:kms/ciphertext:Ciphertext", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Ciphertext{s: s}, nil
}

// GetCiphertext gets an existing Ciphertext resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCiphertext(ctx *pulumi.Context,
	name string, id pulumi.ID, state *CiphertextState, opts ...pulumi.ResourceOpt) (*Ciphertext, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["ciphertextBlob"] = state.CiphertextBlob
		inputs["context"] = state.Context
		inputs["keyId"] = state.KeyId
		inputs["plaintext"] = state.Plaintext
	}
	s, err := ctx.ReadResource("aws:kms/ciphertext:Ciphertext", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Ciphertext{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Ciphertext) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Ciphertext) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Base64 encoded ciphertext
func (r *Ciphertext) CiphertextBlob() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ciphertextBlob"])
}

// An optional mapping that makes up the encryption context.
func (r *Ciphertext) Context() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["context"])
}

// Globally unique key ID for the customer master key.
func (r *Ciphertext) KeyId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyId"])
}

// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
func (r *Ciphertext) Plaintext() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["plaintext"])
}

// Input properties used for looking up and filtering Ciphertext resources.
type CiphertextState struct {
	// Base64 encoded ciphertext
	CiphertextBlob interface{}
	// An optional mapping that makes up the encryption context.
	Context interface{}
	// Globally unique key ID for the customer master key.
	KeyId interface{}
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext interface{}
}

// The set of arguments for constructing a Ciphertext resource.
type CiphertextArgs struct {
	// An optional mapping that makes up the encryption context.
	Context interface{}
	// Globally unique key ID for the customer master key.
	KeyId interface{}
	// Data to be encrypted. Note that this may show up in logs, and it will be stored in the state file.
	Plaintext interface{}
}
