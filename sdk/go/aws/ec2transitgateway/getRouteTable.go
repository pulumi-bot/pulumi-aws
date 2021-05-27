// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on an EC2 Transit Gateway Route Table.
//
// ## Example Usage
// ### By Identifier
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "tgw-rtb-12345678"
// 		_, err := ec2transitgateway.LookupRouteTable(ctx, &ec2transitgateway.LookupRouteTableArgs{
// 			Id: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	var rv LookupRouteTableResult
	err := ctx.Invoke("aws:ec2transitgateway/getRouteTable:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters []GetRouteTableFilter `pulumi:"filters"`
	// Identifier of the EC2 Transit Gateway Route Table.
	Id *string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway Route Table
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	// EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
	Arn string `pulumi:"arn"`
	// Boolean whether this is the default association route table for the EC2 Transit Gateway
	DefaultAssociationRouteTable bool `pulumi:"defaultAssociationRouteTable"`
	// Boolean whether this is the default propagation route table for the EC2 Transit Gateway
	DefaultPropagationRouteTable bool                  `pulumi:"defaultPropagationRouteTable"`
	Filters                      []GetRouteTableFilter `pulumi:"filters"`
	// EC2 Transit Gateway Route Table identifier
	Id string `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway Route Table
	Tags map[string]string `pulumi:"tags"`
	// EC2 Transit Gateway identifier
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

func LookupRouteTableOutput(ctx *pulumi.Context, args LookupRouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupRouteTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteTableResult, error) {
			args := v.(LookupRouteTableArgs)
			r, err := LookupRouteTable(ctx, &args, opts...)
			return *r, err
		}).(LookupRouteTableResultOutput)
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableOutputArgs struct {
	// One or more configuration blocks containing name-values filters. Detailed below.
	Filters GetRouteTableFilterArrayInput `pulumi:"filters"`
	// Identifier of the EC2 Transit Gateway Route Table.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Key-value tags for the EC2 Transit Gateway Route Table
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupRouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
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

// EC2 Transit Gateway Route Table Amazon Resource Name (ARN).
func (o LookupRouteTableResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Boolean whether this is the default association route table for the EC2 Transit Gateway
func (o LookupRouteTableResultOutput) DefaultAssociationRouteTable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRouteTableResult) bool { return v.DefaultAssociationRouteTable }).(pulumi.BoolOutput)
}

// Boolean whether this is the default propagation route table for the EC2 Transit Gateway
func (o LookupRouteTableResultOutput) DefaultPropagationRouteTable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRouteTableResult) bool { return v.DefaultPropagationRouteTable }).(pulumi.BoolOutput)
}

func (o LookupRouteTableResultOutput) Filters() GetRouteTableFilterArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []GetRouteTableFilter { return v.Filters }).(GetRouteTableFilterArrayOutput)
}

// EC2 Transit Gateway Route Table identifier
func (o LookupRouteTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Id }).(pulumi.StringOutput)
}

// Key-value tags for the EC2 Transit Gateway Route Table
func (o LookupRouteTableResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRouteTableResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// EC2 Transit Gateway identifier
func (o LookupRouteTableResultOutput) TransitGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.TransitGatewayId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTableResultOutput{})
}
