// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `getRegion` provides details about a specific AWS region.
//
// As well as validating a given region name this resource can be used to
// discover the name of the region configured within the provider. The latter
// can be useful in a child module which is inheriting an AWS provider
// configuration from its parent module.
func GetRegion(ctx *pulumi.Context, args *GetRegionArgs, opts ...pulumi.InvokeOption) (*GetRegionResult, error) {
	var rv GetRegionResult
	err := ctx.Invoke("aws:index/getRegion:getRegion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegion.
type GetRegionArgs struct {
	// The EC2 endpoint of the region to select.
	Endpoint *string `pulumi:"endpoint"`
	// The full name of the region to select.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getRegion.
type GetRegionResult struct {
	// The region's description in this format: "Location (Region name)".
	Description string `pulumi:"description"`
	// The EC2 endpoint for the selected region.
	Endpoint string `pulumi:"endpoint"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the selected region.
	Name string `pulumi:"name"`
}
