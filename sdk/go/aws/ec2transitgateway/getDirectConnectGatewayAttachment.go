// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on an EC2 Transit Gateway's attachment to a Direct Connect Gateway.
//
// ## Example Usage
func GetDirectConnectGatewayAttachment(ctx *pulumi.Context, args *GetDirectConnectGatewayAttachmentArgs, opts ...pulumi.InvokeOption) (*GetDirectConnectGatewayAttachmentResult, error) {
	var rv GetDirectConnectGatewayAttachmentResult
	err := ctx.Invoke("aws:ec2transitgateway/getDirectConnectGatewayAttachment:getDirectConnectGatewayAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDirectConnectGatewayAttachment.
type GetDirectConnectGatewayAttachmentArgs struct {
	// Identifier of the Direct Connect Gateway.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// Configuration block(s) for filtering. Detailed below.
	Filters []GetDirectConnectGatewayAttachmentFilter `pulumi:"filters"`
	// A map of tags, each pair of which must exactly match a pair on the desired Transit Gateway Direct Connect Gateway Attachment.
	Tags map[string]string `pulumi:"tags"`
	// Identifier of the EC2 Transit Gateway.
	TransitGatewayId *string `pulumi:"transitGatewayId"`
}

// A collection of values returned by getDirectConnectGatewayAttachment.
type GetDirectConnectGatewayAttachmentResult struct {
	DxGatewayId *string                                   `pulumi:"dxGatewayId"`
	Filters     []GetDirectConnectGatewayAttachmentFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway Attachment
	Tags             map[string]string `pulumi:"tags"`
	TransitGatewayId *string           `pulumi:"transitGatewayId"`
}
