// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides details about an EC2 Local Gateway Route Table.
//
// This data source can prove useful when a module accepts a local gateway route table id as
// an input variable and needs to, for example, find the associated Outpost or Local Gateway.
//
// ## Example Usage
//
// The following example returns a specific local gateway route table ID
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
// 		opt0 := awsEc2LocalGatewayRouteTable
// 		_, err := ec2.GetLocalGatewayRouteTable(ctx, &ec2.GetLocalGatewayRouteTableArgs{
// 			LocalGatewayRouteTableId: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetLocalGatewayRouteTable(ctx *pulumi.Context, args *GetLocalGatewayRouteTableArgs, opts ...pulumi.InvokeOption) (*GetLocalGatewayRouteTableResult, error) {
	var rv GetLocalGatewayRouteTableResult
	err := ctx.Invoke("aws:ec2/getLocalGatewayRouteTable:getLocalGatewayRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGatewayRouteTable.
type GetLocalGatewayRouteTableArgs struct {
	Filters []GetLocalGatewayRouteTableFilter `pulumi:"filters"`
	// The id of the specific local gateway route table to retrieve.
	LocalGatewayId *string `pulumi:"localGatewayId"`
	// Local Gateway Route Table Id assigned to desired local gateway route table
	LocalGatewayRouteTableId *string `pulumi:"localGatewayRouteTableId"`
	// The arn of the Outpost the local gateway route table is associated with.
	OutpostArn *string `pulumi:"outpostArn"`
	// The state of the local gateway route table.
	State *string `pulumi:"state"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired local gateway route table.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLocalGatewayRouteTable.
type GetLocalGatewayRouteTableResult struct {
	Filters []GetLocalGatewayRouteTableFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string            `pulumi:"id"`
	LocalGatewayId           string            `pulumi:"localGatewayId"`
	LocalGatewayRouteTableId string            `pulumi:"localGatewayRouteTableId"`
	OutpostArn               string            `pulumi:"outpostArn"`
	State                    string            `pulumi:"state"`
	Tags                     map[string]string `pulumi:"tags"`
}
