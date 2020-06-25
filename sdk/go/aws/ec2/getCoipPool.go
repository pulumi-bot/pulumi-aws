// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides details about a specific EC2 Customer-Owned IP Pool.
//
// This data source can prove useful when a module accepts a coip pool id as
// an input variable and needs to, for example, determine the CIDR block of that
// COIP Pool.
//
// ## Example Usage
//
// The following example returns a specific coip pool ID
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
// 		_, err := ec2.GetCoipPool(ctx, &ec2.GetCoipPoolArgs{
// 			Id: coipPoolId,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetCoipPool(ctx *pulumi.Context, args *GetCoipPoolArgs, opts ...pulumi.InvokeOption) (*GetCoipPoolResult, error) {
	var rv GetCoipPoolResult
	err := ctx.Invoke("aws:ec2/getCoipPool:getCoipPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCoipPool.
type GetCoipPoolArgs struct {
	Filters []GetCoipPoolFilter `pulumi:"filters"`
	// Local Gateway Route Table Id assigned to desired COIP Pool
	LocalGatewayRouteTableId *string `pulumi:"localGatewayRouteTableId"`
	// The id of the specific COIP Pool to retrieve.
	PoolId *string `pulumi:"poolId"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired COIP Pool.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCoipPool.
type GetCoipPoolResult struct {
	Filters []GetCoipPoolFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string `pulumi:"id"`
	LocalGatewayRouteTableId string `pulumi:"localGatewayRouteTableId"`
	// Set of CIDR blocks in pool
	PoolCidrs []string          `pulumi:"poolCidrs"`
	PoolId    string            `pulumi:"poolId"`
	Tags      map[string]string `pulumi:"tags"`
}
