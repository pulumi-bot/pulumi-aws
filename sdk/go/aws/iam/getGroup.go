// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM group. By using this data source, you can reference IAM group
// properties without having to hard code ARNs as input.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.LookupGroup(ctx, &iam.LookupGroupArgs{
// 			GroupName: "an_example_group_name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	var rv LookupGroupResult
	err := ctx.Invoke("aws:iam/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// The friendly IAM group name to match.
	GroupName string `pulumi:"groupName"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The Amazon Resource Name (ARN) specifying the iam user.
	Arn string `pulumi:"arn"`
	// The stable and unique string identifying the group.
	GroupId   string `pulumi:"groupId"`
	GroupName string `pulumi:"groupName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The path to the iam user.
	Path string `pulumi:"path"`
	// List of objects containing group member information. See supported fields below.
	Users []GetGroupUser `pulumi:"users"`
}
