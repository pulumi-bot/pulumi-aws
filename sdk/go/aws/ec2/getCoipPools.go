// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetCoipPools(ctx *pulumi.Context, args *GetCoipPoolsArgs, opts ...pulumi.InvokeOption) (*GetCoipPoolsResult, error) {
	var rv GetCoipPoolsResult
	err := ctx.Invoke("aws:ec2/getCoipPools:getCoipPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCoipPools.
type GetCoipPoolsArgs struct {
	Filters []GetCoipPoolsFilter `pulumi:"filters"`
	Tags    map[string]string    `pulumi:"tags"`
}

// A collection of values returned by getCoipPools.
type GetCoipPoolsResult struct {
	Filters []GetCoipPoolsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id      string            `pulumi:"id"`
	PoolIds []string          `pulumi:"poolIds"`
	Tags    map[string]string `pulumi:"tags"`
}
