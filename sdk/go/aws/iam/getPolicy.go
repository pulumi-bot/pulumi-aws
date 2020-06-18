// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This data source can be used to fetch information about a specific
// IAM policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := iam.LookupPolicy(ctx, &iam.LookupPolicyArgs{
// 			Arn: "arn:aws:iam::123456789012:policy/UsersManageOwnCredentials",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("aws:iam/getPolicy:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicy.
type LookupPolicyArgs struct {
	// ARN of the IAM policy.
	Arn string `pulumi:"arn"`
}

// A collection of values returned by getPolicy.
type LookupPolicyResult struct {
	// The Amazon Resource Name (ARN) specifying the policy.
	Arn string `pulumi:"arn"`
	// The description of the policy.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the IAM policy.
	Name string `pulumi:"name"`
	// The path to the policy.
	Path string `pulumi:"path"`
	// The policy document of the policy.
	Policy string `pulumi:"policy"`
}
