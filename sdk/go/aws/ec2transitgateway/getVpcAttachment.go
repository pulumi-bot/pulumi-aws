// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on an EC2 Transit Gateway VPC Attachment.
//
// ## Example Usage
// ### By Filter
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ec2transitgateway.LookupVpcAttachment(ctx, &ec2transitgateway.LookupVpcAttachmentArgs{
// 			Filters: []ec2transitgateway.GetVpcAttachmentFilter{
// 				ec2transitgateway.GetVpcAttachmentFilter{
// 					Name: "vpc-id",
// 					Values: []string{
// 						"vpc-12345678",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### By Identifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "tgw-attach-12345678"
// 		_, err := ec2transitgateway.LookupVpcAttachment(ctx, &ec2transitgateway.LookupVpcAttachmentArgs{
// 			Id: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVpcAttachment(ctx *pulumi.Context, args *LookupVpcAttachmentArgs, opts ...pulumi.InvokeOption) (*LookupVpcAttachmentResult, error) {
	var rv LookupVpcAttachmentResult
	err := ctx.Invoke("aws:ec2transitgateway/getVpcAttachment:getVpcAttachment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcAttachment.
type LookupVpcAttachmentArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters []GetVpcAttachmentFilter `pulumi:"filters"`
	// Identifier of the EC2 Transit Gateway VPC Attachment.
	Id *string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway VPC Attachment
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcAttachment.
type LookupVpcAttachmentResult struct {
	// Whether DNS support is enabled.
	DnsSupport string                   `pulumi:"dnsSupport"`
	Filters    []GetVpcAttachmentFilter `pulumi:"filters"`
	// EC2 Transit Gateway VPC Attachment identifier
	Id *string `pulumi:"id"`
	// Whether IPv6 support is enabled.
	Ipv6Support string `pulumi:"ipv6Support"`
	// Identifiers of EC2 Subnets.
	SubnetIds []string `pulumi:"subnetIds"`
	// Key-value tags for the EC2 Transit Gateway VPC Attachment
	Tags map[string]string `pulumi:"tags"`
	// EC2 Transit Gateway identifier
	TransitGatewayId string `pulumi:"transitGatewayId"`
	// Identifier of EC2 VPC.
	VpcId string `pulumi:"vpcId"`
	// Identifier of the AWS account that owns the EC2 VPC.
	VpcOwnerId string `pulumi:"vpcOwnerId"`
}
