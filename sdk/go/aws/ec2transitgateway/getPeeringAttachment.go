// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPeeringAttachment(ctx *pulumi.Context, args *LookupPeeringAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupPeeringAttachmentResult, error) {
	var rv LookupPeeringAttachmentResult
	err := ctx.Invoke("aws:ec2transitgateway/getPeeringAttachment:getPeeringAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPeeringAttachment.
type LookupPeeringAttachmentArgs struct {
	Filters []GetPeeringAttachmentFilter `pulumi:"filters"`
	Id      *string                      `pulumi:"id"`
	Tags    map[string]string            `pulumi:"tags"`
}

// A collection of values returned by getPeeringAttachment.
type LookupPeeringAttachmentResult struct {
	Filters              []GetPeeringAttachmentFilter `pulumi:"filters"`
	Id                   *string                      `pulumi:"id"`
	PeerAccountId        string                       `pulumi:"peerAccountId"`
	PeerRegion           string                       `pulumi:"peerRegion"`
	PeerTransitGatewayId string                       `pulumi:"peerTransitGatewayId"`
	Tags                 map[string]string            `pulumi:"tags"`
	TransitGatewayId     string                       `pulumi:"transitGatewayId"`
}
