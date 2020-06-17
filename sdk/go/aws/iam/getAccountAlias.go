// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The IAM Account Alias data source allows access to the account alias
// for the effective account in which this provider is working.
//
// ## Example Usage
func LookupAccountAlias(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupAccountAliasResult, error) {
	var rv LookupAccountAliasResult
	err := ctx.Invoke("aws:iam/getAccountAlias:getAccountAlias", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getAccountAlias.
type LookupAccountAliasResult struct {
	// The alias associated with the AWS account.
	AccountAlias string `pulumi:"accountAlias"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
