// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupLogGroup(ctx *pulumi.Context, args *LookupLogGroupArgs, opts ...pulumi.InvokeOption) (*LookupLogGroupResult, error) {
	var rv LookupLogGroupResult
	err := ctx.Invoke("aws:cloudwatch/getLogGroup:getLogGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogGroup.
type LookupLogGroupArgs struct {
	Name string            `pulumi:"name"`
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getLogGroup.
type LookupLogGroupResult struct {
	Arn          string `pulumi:"arn"`
	CreationTime int    `pulumi:"creationTime"`
	// The provider-assigned unique ID for this managed resource.
	Id              string            `pulumi:"id"`
	KmsKeyId        string            `pulumi:"kmsKeyId"`
	Name            string            `pulumi:"name"`
	RetentionInDays int               `pulumi:"retentionInDays"`
	Tags            map[string]string `pulumi:"tags"`
}
