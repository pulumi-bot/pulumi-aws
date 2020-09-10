// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupEncryptionByDefault(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*LookupEncryptionByDefaultResult, error) {
	var rv LookupEncryptionByDefaultResult
	err := ctx.Invoke("aws:ebs/getEncryptionByDefault:getEncryptionByDefault", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getEncryptionByDefault.
type LookupEncryptionByDefaultResult struct {
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
