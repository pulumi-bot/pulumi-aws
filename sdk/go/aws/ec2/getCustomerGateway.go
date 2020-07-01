// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get an existing AWS Customer Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		foo, err := ec2.LookupCustomerGateway(ctx, &ec2.LookupCustomerGatewayArgs{
// 			Filters: []ec2.GetCustomerGatewayFilter{
// 				ec2.GetCustomerGatewayFilter{
// 					Name: "tag:Name",
// 					Values: []string{
// 						"foo-prod",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		main, err := ec2.NewVpnGateway(ctx, "main", &ec2.VpnGatewayArgs{
// 			AmazonSideAsn: pulumi.String("7224"),
// 			VpcId:         pulumi.Any(aws_vpc.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewVpnConnection(ctx, "transit", &ec2.VpnConnectionArgs{
// 			CustomerGatewayId: pulumi.String(foo.Id),
// 			StaticRoutesOnly:  pulumi.Bool(false),
// 			Type:              pulumi.String(foo.Type),
// 			VpnGatewayId:      main.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupCustomerGateway(ctx *pulumi.Context, args *LookupCustomerGatewayArgs, opts ...pulumi.InvokeOption) (*LookupCustomerGatewayResult, error) {
	var rv LookupCustomerGatewayResult
	err := ctx.Invoke("aws:ec2/getCustomerGateway:getCustomerGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerGateway.
type LookupCustomerGatewayArgs struct {
	// One or more [name-value pairs][dcg-filters] to filter by.
	Filters []GetCustomerGatewayFilter `pulumi:"filters"`
	// The ID of the gateway.
	Id *string `pulumi:"id"`
	// Map of key-value pairs assigned to the gateway.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCustomerGateway.
type LookupCustomerGatewayResult struct {
	// The ARN of the customer gateway.
	Arn string `pulumi:"arn"`
	// (Optional) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn  int                        `pulumi:"bgpAsn"`
	Filters []GetCustomerGatewayFilter `pulumi:"filters"`
	Id      *string                    `pulumi:"id"`
	// (Optional) The IP address of the gateway's Internet-routable external interface.
	IpAddress string `pulumi:"ipAddress"`
	// Map of key-value pairs assigned to the gateway.
	Tags map[string]string `pulumi:"tags"`
	// (Optional) The type of customer gateway. The only type AWS supports at this time is "ipsec.1".
	Type string `pulumi:"type"`
}
