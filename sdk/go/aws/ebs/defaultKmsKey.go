// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage the default customer master key (CMK) that your AWS account uses to encrypt EBS volumes.
// 
// Your AWS account has an AWS-managed default CMK that is used for encrypting an EBS volume when no CMK is specified in the API call that creates the volume.
// By using the `ebs.DefaultKmsKey` resource, you can specify a customer-managed CMK to use in place of the AWS-managed default CMK.
// 
// > **NOTE:** Creating an `ebs.DefaultKmsKey` resource does not enable default EBS encryption. Use the `ebs.EncryptionByDefault` to enable default EBS encryption.
// 
// > **NOTE:** Destroying this resource will reset the default CMK to the account's AWS-managed default CMK for EBS.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ebs_default_kms_key.html.markdown.
type DefaultKmsKey struct {
	s *pulumi.ResourceState
}

// NewDefaultKmsKey registers a new resource with the given unique name, arguments, and options.
func NewDefaultKmsKey(ctx *pulumi.Context,
	name string, args *DefaultKmsKeyArgs, opts ...pulumi.ResourceOpt) (*DefaultKmsKey, error) {
	if args == nil || args.KeyArn == nil {
		return nil, errors.New("missing required argument 'KeyArn'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["keyArn"] = nil
	} else {
		inputs["keyArn"] = args.KeyArn
	}
	s, err := ctx.RegisterResource("aws:ebs/defaultKmsKey:DefaultKmsKey", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultKmsKey{s: s}, nil
}

// GetDefaultKmsKey gets an existing DefaultKmsKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultKmsKey(ctx *pulumi.Context,
	name string, id pulumi.ID, state *DefaultKmsKeyState, opts ...pulumi.ResourceOpt) (*DefaultKmsKey, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["keyArn"] = state.KeyArn
	}
	s, err := ctx.ReadResource("aws:ebs/defaultKmsKey:DefaultKmsKey", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &DefaultKmsKey{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *DefaultKmsKey) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *DefaultKmsKey) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
func (r *DefaultKmsKey) KeyArn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyArn"])
}

// Input properties used for looking up and filtering DefaultKmsKey resources.
type DefaultKmsKeyState struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn interface{}
}

// The set of arguments for constructing a DefaultKmsKey resource.
type DefaultKmsKeyArgs struct {
	// The ARN of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use to encrypt the EBS volume.
	KeyArn interface{}
}
