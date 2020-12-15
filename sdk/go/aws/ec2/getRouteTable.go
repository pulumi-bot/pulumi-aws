// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ec2.RouteTable` provides details about a specific Route Table.
//
// This resource can prove useful when a module accepts a Subnet id as
// an input variable and needs to, for example, add a route in
// the Route Table.
//
// ## Example Usage
//
// The following example shows how one might accept a Route Table id as a variable
// and use this data source to obtain the data necessary to create a route.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		subnetId := cfg.RequireObject("subnetId")
// 		opt0 := subnetId
// 		selected, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
// 			SubnetId: _opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewRoute(ctx, "route", &ec2.RouteArgs{
// 			RouteTableId:           pulumi.String(selected.Id),
// 			DestinationCidrBlock:   pulumi.String("10.0.1.0/22"),
// 			VpcPeeringConnectionId: pulumi.String("pcx-45ff3dc1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	var rv LookupRouteTableResult
	err := ctx.Invoke("aws:ec2/getRouteTable:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableArgs struct {
	// Custom filter block as described below.
	Filters []GetRouteTableFilter `pulumi:"filters"`
	// The id of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
	GatewayId *string `pulumi:"gatewayId"`
	// The id of the specific Route Table to retrieve.
	RouteTableId *string `pulumi:"routeTableId"`
	// The id of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired Route Table.
	Tags map[string]string `pulumi:"tags"`
	// The id of the VPC that the desired Route Table belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	Associations []GetRouteTableAssociationType `pulumi:"associations"`
	Filters      []GetRouteTableFilter          `pulumi:"filters"`
	// The Gateway ID. Only set when associated with an Internet Gateway or Virtual Private Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The ID of the AWS account that owns the route table
	OwnerId string `pulumi:"ownerId"`
	// The Route Table ID.
	RouteTableId string               `pulumi:"routeTableId"`
	Routes       []GetRouteTableRoute `pulumi:"routes"`
	// The Subnet ID. Only set when associated with a Subnet.
	SubnetId string            `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	VpcId    string            `pulumi:"vpcId"`
}
