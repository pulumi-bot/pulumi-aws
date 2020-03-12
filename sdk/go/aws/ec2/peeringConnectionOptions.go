// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource to manage VPC peering connection options.
//
// > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
// both a standalone VPC Peering Connection Options and a VPC Peering Connection
// resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
// connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
// Doing so will cause a conflict of options and will overwrite the options.
// Using a VPC Peering Connection Options resource decouples management of the connection options from
// management of the VPC Peering Connection and allows options to be set correctly in cross-region and
// cross-account scenarios.
//
// Basic usage:
//
// {{% examples %}}
// {{% /examples %}}
//
// Basic cross-account usage:
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_peering_connection_options.html.markdown.
type PeeringConnectionOptions struct {
	pulumi.CustomResourceState

	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterOutput `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterOutput `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringOutput `pulumi:"vpcPeeringConnectionId"`
}

// NewPeeringConnectionOptions registers a new resource with the given unique name, arguments, and options.
func NewPeeringConnectionOptions(ctx *pulumi.Context,
	name string, args *PeeringConnectionOptionsArgs, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	if args == nil || args.VpcPeeringConnectionId == nil {
		return nil, errors.New("missing required argument 'VpcPeeringConnectionId'")
	}
	if args == nil {
		args = &PeeringConnectionOptionsArgs{}
	}
	var resource PeeringConnectionOptions
	err := ctx.RegisterResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeeringConnectionOptions gets an existing PeeringConnectionOptions resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeeringConnectionOptions(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringConnectionOptionsState, opts ...pulumi.ResourceOption) (*PeeringConnectionOptions, error) {
	var resource PeeringConnectionOptions
	err := ctx.ReadResource("aws:ec2/peeringConnectionOptions:PeeringConnectionOptions", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PeeringConnectionOptions resources.
type peeringConnectionOptionsState struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter *PeeringConnectionOptionsAccepter `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester *PeeringConnectionOptionsRequester `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId *string `pulumi:"vpcPeeringConnectionId"`
}

type PeeringConnectionOptionsState struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterPtrInput
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterPtrInput
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringPtrInput
}

func (PeeringConnectionOptionsState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionOptionsState)(nil)).Elem()
}

type peeringConnectionOptionsArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter *PeeringConnectionOptionsAccepter `pulumi:"accepter"`
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester *PeeringConnectionOptionsRequester `pulumi:"requester"`
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId string `pulumi:"vpcPeeringConnectionId"`
}

// The set of arguments for constructing a PeeringConnectionOptions resource.
type PeeringConnectionOptionsArgs struct {
	// An optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that accepts
	// the peering connection (a maximum of one).
	Accepter PeeringConnectionOptionsAccepterPtrInput
	// A optional configuration block that allows for [VPC Peering Connection]
	// (http://docs.aws.amazon.com/AmazonVPC/latest/PeeringGuide) options to be set for the VPC that requests
	// the peering connection (a maximum of one).
	Requester PeeringConnectionOptionsRequesterPtrInput
	// The ID of the requester VPC peering connection.
	VpcPeeringConnectionId pulumi.StringInput
}

func (PeeringConnectionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringConnectionOptionsArgs)(nil)).Elem()
}

