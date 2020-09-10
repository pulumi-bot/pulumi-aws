// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupResourceShare(ctx *pulumi.Context, args *LookupResourceShareArgs, opts ...pulumi.InvokeOption) (*LookupResourceShareResult, error) {
	var rv LookupResourceShareResult
	err := ctx.Invoke("aws:ram/getResourceShare:getResourceShare", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceShare.
type LookupResourceShareArgs struct {
	Filters       []GetResourceShareFilter `pulumi:"filters"`
	Name          string                   `pulumi:"name"`
	ResourceOwner string                   `pulumi:"resourceOwner"`
	Tags          map[string]string        `pulumi:"tags"`
}

// A collection of values returned by getResourceShare.
type LookupResourceShareResult struct {
	Arn     string                   `pulumi:"arn"`
	Filters []GetResourceShareFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id              string            `pulumi:"id"`
	Name            string            `pulumi:"name"`
	OwningAccountId string            `pulumi:"owningAccountId"`
	ResourceOwner   string            `pulumi:"resourceOwner"`
	Status          string            `pulumi:"status"`
	Tags            map[string]string `pulumi:"tags"`
}
