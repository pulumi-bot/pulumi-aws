// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Step Functions Activity data source
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
// 		_, err := sfn.LookupActivity(ctx, &sfn.LookupActivityArgs{
// 			Name: "my-activity",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupActivity(ctx *pulumi.Context, args *LookupActivityArgs, opts ...pulumi.InvokeOption) (*LookupActivityResult, error) {
	var rv LookupActivityResult
	err := ctx.Invoke("aws:sfn/getActivity:getActivity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActivity.
type LookupActivityArgs struct {
	// The Amazon Resource Name (ARN) that identifies the activity.
	Arn *string `pulumi:"arn"`
	// The name that identifies the activity.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getActivity.
type LookupActivityResult struct {
	Arn string `pulumi:"arn"`
	// The date the activity was created.
	CreationDate string `pulumi:"creationDate"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}
