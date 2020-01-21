// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getCustomerGateway

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/getCustomerGatewayFilter"
)

// Get an existing AWS Customer Gateway.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/customer_gateway.html.markdown.
func GetCustomerGateway(ctx *pulumi.Context, args *GetCustomerGatewayArgs, opts ...pulumi.InvokeOption) (*GetCustomerGatewayResult, error) {
	var rv GetCustomerGatewayResult
	err := ctx.Invoke("aws:ec2/getCustomerGateway:getCustomerGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCustomerGateway.
type GetCustomerGatewayArgs struct {
	// One or more [name-value pairs][dcg-filters] to filter by.
	Filters []ec2getCustomerGatewayFilter.GetCustomerGatewayFilter `pulumi:"filters"`
	// The ID of the gateway.
	Id *string `pulumi:"id"`
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getCustomerGateway.
type GetCustomerGatewayResult struct {
	// (Optional) The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn int `pulumi:"bgpAsn"`
	Filters []ec2getCustomerGatewayFilter.GetCustomerGatewayFilter `pulumi:"filters"`
	Id *string `pulumi:"id"`
	// (Optional) The IP address of the gateway's Internet-routable external interface.
	IpAddress string `pulumi:"ipAddress"`
	// Map of key-value pairs assigned to the gateway.
	Tags map[string]interface{} `pulumi:"tags"`
	// (Optional) The type of customer gateway. The only type AWS supports at this time is "ipsec.1".
	Type string `pulumi:"type"`
}

