// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetInstanceTypeOffering(ctx *pulumi.Context, args *GetInstanceTypeOfferingArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeOfferingResult, error) {
	var rv GetInstanceTypeOfferingResult
	err := ctx.Invoke("aws:ec2/getInstanceTypeOffering:getInstanceTypeOffering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeOffering.
type GetInstanceTypeOfferingArgs struct {
	Filters                []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	LocationType           *string                         `pulumi:"locationType"`
	PreferredInstanceTypes []string                        `pulumi:"preferredInstanceTypes"`
}

// A collection of values returned by getInstanceTypeOffering.
type GetInstanceTypeOfferingResult struct {
	Filters []GetInstanceTypeOfferingFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string   `pulumi:"id"`
	InstanceType           string   `pulumi:"instanceType"`
	LocationType           *string  `pulumi:"locationType"`
	PreferredInstanceTypes []string `pulumi:"preferredInstanceTypes"`
}
