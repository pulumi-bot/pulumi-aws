// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetRegions(ctx *pulumi.Context, args *GetRegionsArgs, opts ...pulumi.InvokeOption) (*GetRegionsResult, error) {
	var rv GetRegionsResult
	err := ctx.Invoke("aws:index/getRegions:getRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegions.
type GetRegionsArgs struct {
	AllRegions *bool              `pulumi:"allRegions"`
	Filters    []GetRegionsFilter `pulumi:"filters"`
}

// A collection of values returned by getRegions.
type GetRegionsResult struct {
	AllRegions *bool              `pulumi:"allRegions"`
	Filters    []GetRegionsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id    string   `pulumi:"id"`
	Names []string `pulumi:"names"`
}
