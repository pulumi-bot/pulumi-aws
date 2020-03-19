// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package wafregional

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `wafregional.WebAcl` Retrieves a WAF Regional Web ACL Resource Id.
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/wafregional_web_acl.html.markdown.
func LookupWebAcl(ctx *pulumi.Context, args *LookupWebAclArgs, opts ...pulumi.InvokeOption) (*LookupWebAclResult, error) {
	var rv LookupWebAclResult
	err := ctx.Invoke("aws:wafregional/getWebAcl:getWebAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWebAcl.
type LookupWebAclArgs struct {
	// The name of the WAF Regional Web ACL.
	Name string `pulumi:"name"`
}


// A collection of values returned by getWebAcl.
type LookupWebAclResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
}

