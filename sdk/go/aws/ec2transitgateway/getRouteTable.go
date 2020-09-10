// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2transitgateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

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
	Filters []GetRouteTableFilter `pulumi:"filters"`
	Id      *string               `pulumi:"id"`
	Tags    map[string]string     `pulumi:"tags"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	DefaultAssociationRouteTable bool                  `pulumi:"defaultAssociationRouteTable"`
	DefaultPropagationRouteTable bool                  `pulumi:"defaultPropagationRouteTable"`
	Filters                      []GetRouteTableFilter `pulumi:"filters"`
	Id                           *string               `pulumi:"id"`
	Tags                         map[string]string     `pulumi:"tags"`
	TransitGatewayId             string                `pulumi:"transitGatewayId"`
}
