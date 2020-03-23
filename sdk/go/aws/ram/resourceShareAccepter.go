// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ram

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the [`ram.ResourceShare` resource](https://www.terraform.io/docs/providers/aws/r/ram_resource_share.html).
//
// > **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ram_resource_share_accepter.markdown.
type ResourceShareAccepter struct {
	pulumi.CustomResourceState

	// The ARN of the resource share invitation.
	InvitationArn pulumi.StringOutput `pulumi:"invitationArn"`
	// The account ID of the receiver account which accepts the invitation.
	ReceiverAccountId pulumi.StringOutput `pulumi:"receiverAccountId"`
	// A list of the resource ARNs shared via the resource share.
	Resources pulumi.StringArrayOutput `pulumi:"resources"`
	// The account ID of the sender account which submits the invitation.
	SenderAccountId pulumi.StringOutput `pulumi:"senderAccountId"`
	// The ARN of the resource share.
	ShareArn pulumi.StringOutput `pulumi:"shareArn"`
	// The ID of the resource share as displayed in the console.
	ShareId pulumi.StringOutput `pulumi:"shareId"`
	// The name of the resource share.
	ShareName pulumi.StringOutput `pulumi:"shareName"`
	// The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewResourceShareAccepter registers a new resource with the given unique name, arguments, and options.
func NewResourceShareAccepter(ctx *pulumi.Context,
	name string, args *ResourceShareAccepterArgs, opts ...pulumi.ResourceOption) (*ResourceShareAccepter, error) {
	if args == nil || args.ShareArn == nil {
		return nil, errors.New("missing required argument 'ShareArn'")
	}
	if args == nil {
		args = &ResourceShareAccepterArgs{}
	}
	var resource ResourceShareAccepter
	err := ctx.RegisterResource("aws:ram/resourceShareAccepter:ResourceShareAccepter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceShareAccepter gets an existing ResourceShareAccepter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceShareAccepter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceShareAccepterState, opts ...pulumi.ResourceOption) (*ResourceShareAccepter, error) {
	var resource ResourceShareAccepter
	err := ctx.ReadResource("aws:ram/resourceShareAccepter:ResourceShareAccepter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceShareAccepter resources.
type resourceShareAccepterState struct {
	// The ARN of the resource share invitation.
	InvitationArn *string `pulumi:"invitationArn"`
	// The account ID of the receiver account which accepts the invitation.
	ReceiverAccountId *string `pulumi:"receiverAccountId"`
	// A list of the resource ARNs shared via the resource share.
	Resources []string `pulumi:"resources"`
	// The account ID of the sender account which submits the invitation.
	SenderAccountId *string `pulumi:"senderAccountId"`
	// The ARN of the resource share.
	ShareArn *string `pulumi:"shareArn"`
	// The ID of the resource share as displayed in the console.
	ShareId *string `pulumi:"shareId"`
	// The name of the resource share.
	ShareName *string `pulumi:"shareName"`
	// The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
	Status *string `pulumi:"status"`
}

type ResourceShareAccepterState struct {
	// The ARN of the resource share invitation.
	InvitationArn pulumi.StringPtrInput
	// The account ID of the receiver account which accepts the invitation.
	ReceiverAccountId pulumi.StringPtrInput
	// A list of the resource ARNs shared via the resource share.
	Resources pulumi.StringArrayInput
	// The account ID of the sender account which submits the invitation.
	SenderAccountId pulumi.StringPtrInput
	// The ARN of the resource share.
	ShareArn pulumi.StringPtrInput
	// The ID of the resource share as displayed in the console.
	ShareId pulumi.StringPtrInput
	// The name of the resource share.
	ShareName pulumi.StringPtrInput
	// The status of the resource share (ACTIVE, PENDING, FAILED, DELETING, DELETED).
	Status pulumi.StringPtrInput
}

func (ResourceShareAccepterState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareAccepterState)(nil)).Elem()
}

type resourceShareAccepterArgs struct {
	// The ARN of the resource share.
	ShareArn string `pulumi:"shareArn"`
}

// The set of arguments for constructing a ResourceShareAccepter resource.
type ResourceShareAccepterArgs struct {
	// The ARN of the resource share.
	ShareArn pulumi.StringInput
}

func (ResourceShareAccepterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceShareAccepterArgs)(nil)).Elem()
}

