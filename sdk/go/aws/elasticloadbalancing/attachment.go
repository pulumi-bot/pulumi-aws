// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package elasticloadbalancing

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: aws.elasticloadbalancing.Attachment has been deprecated in favor of aws.elb.Attachment
type Attachment struct {
	pulumi.CustomResourceState

	Elb      pulumi.StringOutput `pulumi:"elb"`
	Instance pulumi.StringOutput `pulumi:"instance"`
}

// NewAttachment registers a new resource with the given unique name, arguments, and options.
func NewAttachment(ctx *pulumi.Context,
	name string, args *AttachmentArgs, opts ...pulumi.ResourceOption) (*Attachment, error) {
	if args == nil || args.Elb == nil {
		return nil, errors.New("missing required argument 'Elb'")
	}
	if args == nil || args.Instance == nil {
		return nil, errors.New("missing required argument 'Instance'")
	}
	if args == nil {
		args = &AttachmentArgs{}
	}
	var resource Attachment
	err := ctx.RegisterResource("aws:elasticloadbalancing/attachment:Attachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachment gets an existing Attachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachmentState, opts ...pulumi.ResourceOption) (*Attachment, error) {
	var resource Attachment
	err := ctx.ReadResource("aws:elasticloadbalancing/attachment:Attachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Attachment resources.
type attachmentState struct {
	Elb      *string `pulumi:"elb"`
	Instance *string `pulumi:"instance"`
}

type AttachmentState struct {
	Elb      pulumi.StringPtrInput
	Instance pulumi.StringPtrInput
}

func (AttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentState)(nil)).Elem()
}

type attachmentArgs struct {
	Elb      string `pulumi:"elb"`
	Instance string `pulumi:"instance"`
}

// The set of arguments for constructing a Attachment resource.
type AttachmentArgs struct {
	Elb      pulumi.StringInput
	Instance pulumi.StringInput
}

func (AttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachmentArgs)(nil)).Elem()
}
