// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetInstanceTypeOfferings(ctx *pulumi.Context, args *GetInstanceTypeOfferingsArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeOfferingsResult, error) {
	var rv GetInstanceTypeOfferingsResult
	err := ctx.Invoke("aws:ec2/getInstanceTypeOfferings:getInstanceTypeOfferings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeOfferings.
type GetInstanceTypeOfferingsArgs struct {
	Filters      []GetInstanceTypeOfferingsFilter `pulumi:"filters"`
	LocationType *string                          `pulumi:"locationType"`
}

// A collection of values returned by getInstanceTypeOfferings.
type GetInstanceTypeOfferingsResult struct {
	Filters []GetInstanceTypeOfferingsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	InstanceTypes []string `pulumi:"instanceTypes"`
	LocationType  *string  `pulumi:"locationType"`
}
