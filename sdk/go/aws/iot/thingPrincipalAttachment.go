// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Attaches Principal to AWS IoT Thing.
type ThingPrincipalAttachment struct {
	pulumi.CustomResourceState

	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal pulumi.StringOutput `pulumi:"principal"`
	// The name of the thing.
	Thing pulumi.StringOutput `pulumi:"thing"`
}

// NewThingPrincipalAttachment registers a new resource with the given unique name, arguments, and options.
func NewThingPrincipalAttachment(ctx *pulumi.Context,
	name string, args *ThingPrincipalAttachmentArgs, opts ...pulumi.ResourceOption) (*ThingPrincipalAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	if args.Thing == nil {
		return nil, errors.New("invalid value for required argument 'Thing'")
	}
	var resource ThingPrincipalAttachment
	err := ctx.RegisterResource("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetThingPrincipalAttachment gets an existing ThingPrincipalAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetThingPrincipalAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThingPrincipalAttachmentState, opts ...pulumi.ResourceOption) (*ThingPrincipalAttachment, error) {
	var resource ThingPrincipalAttachment
	err := ctx.ReadResource("aws:iot/thingPrincipalAttachment:ThingPrincipalAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ThingPrincipalAttachment resources.
type thingPrincipalAttachmentState struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal *string `pulumi:"principal"`
	// The name of the thing.
	Thing *string `pulumi:"thing"`
}

type ThingPrincipalAttachmentState struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal pulumi.StringPtrInput
	// The name of the thing.
	Thing pulumi.StringPtrInput
}

func (ThingPrincipalAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*thingPrincipalAttachmentState)(nil)).Elem()
}

type thingPrincipalAttachmentArgs struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal string `pulumi:"principal"`
	// The name of the thing.
	Thing string `pulumi:"thing"`
}

// The set of arguments for constructing a ThingPrincipalAttachment resource.
type ThingPrincipalAttachmentArgs struct {
	// The AWS IoT Certificate ARN or Amazon Cognito Identity ID.
	Principal pulumi.StringInput
	// The name of the thing.
	Thing pulumi.StringInput
}

func (ThingPrincipalAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*thingPrincipalAttachmentArgs)(nil)).Elem()
}

type ThingPrincipalAttachmentInput interface {
	pulumi.Input

	ToThingPrincipalAttachmentOutput() ThingPrincipalAttachmentOutput
	ToThingPrincipalAttachmentOutputWithContext(ctx context.Context) ThingPrincipalAttachmentOutput
}

func (ThingPrincipalAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingPrincipalAttachment)(nil)).Elem()
}

func (i ThingPrincipalAttachment) ToThingPrincipalAttachmentOutput() ThingPrincipalAttachmentOutput {
	return i.ToThingPrincipalAttachmentOutputWithContext(context.Background())
}

func (i ThingPrincipalAttachment) ToThingPrincipalAttachmentOutputWithContext(ctx context.Context) ThingPrincipalAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThingPrincipalAttachmentOutput)
}

type ThingPrincipalAttachmentOutput struct {
	*pulumi.OutputState
}

func (ThingPrincipalAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThingPrincipalAttachmentOutput)(nil)).Elem()
}

func (o ThingPrincipalAttachmentOutput) ToThingPrincipalAttachmentOutput() ThingPrincipalAttachmentOutput {
	return o
}

func (o ThingPrincipalAttachmentOutput) ToThingPrincipalAttachmentOutputWithContext(ctx context.Context) ThingPrincipalAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ThingPrincipalAttachmentOutput{})
}
