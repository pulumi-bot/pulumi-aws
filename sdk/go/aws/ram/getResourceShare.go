// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ram

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ram.ResourceShare` Retrieve information about a RAM Resource Share.
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
	// A filter used to scope the list e.g. by tags. See [related docs] (https://docs.aws.amazon.com/ram/latest/APIReference/API_TagFilter.html).
	Filters []GetResourceShareFilter `pulumi:"filters"`
	// The name of the tag key to filter on.
	Name string `pulumi:"name"`
	// The owner of the resource share. Valid values are SELF or OTHER-ACCOUNTS
	ResourceOwner string `pulumi:"resourceOwner"`
	// The Tags attached to the RAM share
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getResourceShare.
type LookupResourceShareResult struct {
	// The Amazon Resource Name (ARN) of the resource share.
	Arn     string                   `pulumi:"arn"`
	Filters []GetResourceShareFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The ID of the AWS account that owns the resource share.
	OwningAccountId string `pulumi:"owningAccountId"`
	ResourceOwner   string `pulumi:"resourceOwner"`
	// The Status of the RAM share.
	Status string `pulumi:"status"`
	// The Tags attached to the RAM share
	Tags map[string]string `pulumi:"tags"`
}
