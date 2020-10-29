// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieves the summary of a WAFv2 Web ACL.
func LookupWebAcl(ctx *pulumi.Context, args *LookupWebAclArgs, opts ...pulumi.InvokeOption) (*LookupWebAclResult, error) {
	var rv LookupWebAclResult
	err := ctx.Invoke("aws:wafv2/getWebAcl:getWebAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebAcl.
type LookupWebAclArgs struct {
	// The name of the WAFv2 Web ACL.
	Name string `pulumi:"name"`
	// Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
	Scope string `pulumi:"scope"`
}

// A collection of values returned by getWebAcl.
type LookupWebAclResult struct {
	// The Amazon Resource Name (ARN) of the entity.
	Arn string `pulumi:"arn"`
	// The description of the WebACL that helps with identification.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	Name  string `pulumi:"name"`
	Scope string `pulumi:"scope"`
}
