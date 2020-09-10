// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIpSet(ctx *pulumi.Context, args *LookupIpSetArgs, opts ...pulumi.InvokeOption) (*LookupIpSetResult, error) {
	var rv LookupIpSetResult
	err := ctx.Invoke("aws:wafv2/getIpSet:getIpSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpSet.
type LookupIpSetArgs struct {
	Name  string `pulumi:"name"`
	Scope string `pulumi:"scope"`
}

// A collection of values returned by getIpSet.
type LookupIpSetResult struct {
	Addresses   []string `pulumi:"addresses"`
	Arn         string   `pulumi:"arn"`
	Description string   `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id               string `pulumi:"id"`
	IpAddressVersion string `pulumi:"ipAddressVersion"`
	Name             string `pulumi:"name"`
	Scope            string `pulumi:"scope"`
}
