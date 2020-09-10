// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupParameter(ctx *pulumi.Context, args *LookupParameterArgs, opts ...pulumi.InvokeOption) (*LookupParameterResult, error) {
	var rv LookupParameterResult
	err := ctx.Invoke("aws:ssm/getParameter:getParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameter.
type LookupParameterArgs struct {
	Name           string `pulumi:"name"`
	WithDecryption *bool  `pulumi:"withDecryption"`
}

// A collection of values returned by getParameter.
type LookupParameterResult struct {
	Arn string `pulumi:"arn"`
	// The provider-assigned unique ID for this managed resource.
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	Type           string `pulumi:"type"`
	Value          string `pulumi:"value"`
	Version        int    `pulumi:"version"`
	WithDecryption *bool  `pulumi:"withDecryption"`
}
