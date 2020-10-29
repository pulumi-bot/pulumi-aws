// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about multiple Elastic File System (EFS) Access Points.
func GetAccessPoints(ctx *pulumi.Context, args *GetAccessPointsArgs, opts ...pulumi.InvokeOption) (*GetAccessPointsResult, error) {
	var rv GetAccessPointsResult
	err := ctx.Invoke("aws:efs/getAccessPoints:getAccessPoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessPoints.
type GetAccessPointsArgs struct {
	// EFS File System identifier.
	FileSystemId string `pulumi:"fileSystemId"`
}

// A collection of values returned by getAccessPoints.
type GetAccessPointsResult struct {
	// Set of Amazon Resource Names (ARNs).
	Arns         []string `pulumi:"arns"`
	FileSystemId string   `pulumi:"fileSystemId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of identifiers.
	Ids []string `pulumi:"ids"`
}
