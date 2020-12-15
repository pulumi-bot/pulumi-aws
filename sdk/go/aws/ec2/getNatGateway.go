// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides details about a specific Nat Gateway.
//
// ## Example Usage
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
// 		opt0 := aws_subnet.Public.Id
// 		_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
// 			SubnetId: _opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Usage with tags:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := aws_subnet.Public.Id
// 		_, err := ec2.LookupNatGateway(ctx, &ec2.LookupNatGatewayArgs{
// 			SubnetId: _opt0,
// 			Tags: map[string]interface{}{
// 				"Name": "gw NAT",
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	var rv LookupNatGatewayResult
	err := ctx.Invoke("aws:ec2/getNatGateway:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNatGateway.
type LookupNatGatewayArgs struct {
	// Custom filter block as described below.
	Filters []GetNatGatewayFilter `pulumi:"filters"`
	// The id of the specific Nat Gateway to retrieve.
	Id *string `pulumi:"id"`
	// The state of the NAT gateway (pending | failed | available | deleting | deleted ).
	State *string `pulumi:"state"`
	// The id of subnet that the Nat Gateway resides in.
	SubnetId *string `pulumi:"subnetId"`
	// A map of tags, each pair of which must exactly match
	// a pair on the desired Nat Gateway.
	Tags map[string]string `pulumi:"tags"`
	// The id of the VPC that the Nat Gateway resides in.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getNatGateway.
type LookupNatGatewayResult struct {
	// The Id of the EIP allocated to the selected Nat Gateway.
	AllocationId string                `pulumi:"allocationId"`
	Filters      []GetNatGatewayFilter `pulumi:"filters"`
	Id           string                `pulumi:"id"`
	// The Id of the ENI allocated to the selected Nat Gateway.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The private Ip address of the selected Nat Gateway.
	PrivateIp string `pulumi:"privateIp"`
	// The public Ip (EIP) address of the selected Nat Gateway.
	PublicIp string            `pulumi:"publicIp"`
	State    string            `pulumi:"state"`
	SubnetId string            `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	VpcId    string            `pulumi:"vpcId"`
}
