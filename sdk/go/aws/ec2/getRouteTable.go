// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `ec2.RouteTable` provides details about a specific Route Table.
//
// This resource can prove useful when a module accepts a Subnet ID as an input variable and needs to, for example, add a route in the Route Table.
//
// ## Example Usage
//
// The following example shows how one might accept a Route Table ID as a variable and use this data source to obtain the data necessary to create a route.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		cfg := config.New(ctx, "")
// 		subnetId := cfg.RequireObject("subnetId")
// 		opt0 := subnetId
// 		selected, err := ec2.LookupRouteTable(ctx, &ec2.LookupRouteTableArgs{
// 			SubnetId: &opt0,
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
	// Configuration block. Detailed below.
	Filters []GetRouteTableFilter `pulumi:"filters"`
	// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
	GatewayId *string `pulumi:"gatewayId"`
	// ID of the specific Route Table to retrieve.
	RouteTableId *string `pulumi:"routeTableId"`
	// ID of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
	SubnetId *string `pulumi:"subnetId"`
	// Map of tags, each pair of which must exactly match a pair on the desired Route Table.
	Tags map[string]string `pulumi:"tags"`
	// ID of the VPC that the desired Route Table belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	// ARN of the route table.
	Arn string `pulumi:"arn"`
	// List of associations with attributes detailed below.
	Associations []GetRouteTableAssociationType `pulumi:"associations"`
	Filters      []GetRouteTableFilter          `pulumi:"filters"`
	// Gateway ID. Only set when associated with an Internet Gateway or Virtual Private Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ID of the AWS account that owns the route table.
	OwnerId string `pulumi:"ownerId"`
	// Route Table ID.
	RouteTableId string `pulumi:"routeTableId"`
	// List of routes with attributes detailed below.
	Routes []GetRouteTableRoute `pulumi:"routes"`
	// Subnet ID. Only set when associated with a subnet.
	SubnetId string            `pulumi:"subnetId"`
	Tags     map[string]string `pulumi:"tags"`
	VpcId    string            `pulumi:"vpcId"`
}

func LookupRouteTableApply(ctx *pulumi.Context, args LookupRouteTableApplyInput, opts ...pulumi.InvokeOption) LookupRouteTableResultOutput {
	return args.ToLookupRouteTableApplyOutput().ApplyT(func(v LookupRouteTableArgs) (LookupRouteTableResult, error) {
		r, err := LookupRouteTable(ctx, &v, opts...)
		return *r, err

	}).(LookupRouteTableResultOutput)
}

// LookupRouteTableApplyInput is an input type that accepts LookupRouteTableApplyArgs and LookupRouteTableApplyOutput values.
// You can construct a concrete instance of `LookupRouteTableApplyInput` via:
//
//          LookupRouteTableApplyArgs{...}
type LookupRouteTableApplyInput interface {
	pulumi.Input

	ToLookupRouteTableApplyOutput() LookupRouteTableApplyOutput
	ToLookupRouteTableApplyOutputWithContext(context.Context) LookupRouteTableApplyOutput
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableApplyArgs struct {
	// Configuration block. Detailed below.
	Filters GetRouteTableFilterArrayInput `pulumi:"filters"`
	// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
	GatewayId pulumi.StringPtrInput `pulumi:"gatewayId"`
	// ID of the specific Route Table to retrieve.
	RouteTableId pulumi.StringPtrInput `pulumi:"routeTableId"`
	// ID of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Map of tags, each pair of which must exactly match a pair on the desired Route Table.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// ID of the VPC that the desired Route Table belongs to.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupRouteTableApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
}

func (i LookupRouteTableApplyArgs) ToLookupRouteTableApplyOutput() LookupRouteTableApplyOutput {
	return i.ToLookupRouteTableApplyOutputWithContext(context.Background())
}

func (i LookupRouteTableApplyArgs) ToLookupRouteTableApplyOutputWithContext(ctx context.Context) LookupRouteTableApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupRouteTableApplyOutput)
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableApplyOutput struct{ *pulumi.OutputState }

func (LookupRouteTableApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
}

func (o LookupRouteTableApplyOutput) ToLookupRouteTableApplyOutput() LookupRouteTableApplyOutput {
	return o
}

func (o LookupRouteTableApplyOutput) ToLookupRouteTableApplyOutputWithContext(ctx context.Context) LookupRouteTableApplyOutput {
	return o
}

// Configuration block. Detailed below.
func (o LookupRouteTableApplyOutput) Filters() GetRouteTableFilterArrayOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) []GetRouteTableFilter { return v.Filters }).(GetRouteTableFilterArrayOutput)
}

// ID of an Internet Gateway or Virtual Private Gateway which is connected to the Route Table (not exported if not passed as a parameter).
func (o LookupRouteTableApplyOutput) GatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) *string { return v.GatewayId }).(pulumi.StringPtrOutput)
}

// ID of the specific Route Table to retrieve.
func (o LookupRouteTableApplyOutput) RouteTableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) *string { return v.RouteTableId }).(pulumi.StringPtrOutput)
}

// ID of a Subnet which is connected to the Route Table (not exported if not passed as a parameter).
func (o LookupRouteTableApplyOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Map of tags, each pair of which must exactly match a pair on the desired Route Table.
func (o LookupRouteTableApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// ID of the VPC that the desired Route Table belongs to.
func (o LookupRouteTableApplyOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteTableArgs) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupRouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableResult)(nil)).Elem()
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutput() LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutputWithContext(ctx context.Context) LookupRouteTableResultOutput {
	return o
}

// ARN of the route table.
func (o LookupRouteTableResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Arn }).(pulumi.StringOutput)
}

// List of associations with attributes detailed below.
func (o LookupRouteTableResultOutput) Associations() GetRouteTableAssociationTypeArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []GetRouteTableAssociationType { return v.Associations }).(GetRouteTableAssociationTypeArrayOutput)
}

func (o LookupRouteTableResultOutput) Filters() GetRouteTableFilterArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []GetRouteTableFilter { return v.Filters }).(GetRouteTableFilterArrayOutput)
}

// Gateway ID. Only set when associated with an Internet Gateway or Virtual Private Gateway.
func (o LookupRouteTableResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRouteTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// ID of the AWS account that owns the route table.
func (o LookupRouteTableResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Route Table ID.
func (o LookupRouteTableResultOutput) RouteTableId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.RouteTableId }).(pulumi.StringOutput)
}

// List of routes with attributes detailed below.
func (o LookupRouteTableResultOutput) Routes() GetRouteTableRouteArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []GetRouteTableRoute { return v.Routes }).(GetRouteTableRouteArrayOutput)
}

// Subnet ID. Only set when associated with a subnet.
func (o LookupRouteTableResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupRouteTableResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTableApplyOutput{})
	pulumi.RegisterOutputType(LookupRouteTableResultOutput{})
}
