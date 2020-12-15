// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides information about AWS Regions. Can be used to filter regions i.e. by Opt-In status or only regions enabled for current account. To get details like endpoint and description of each region the data source can be combined with the `getRegion` data source.
//
// ## Example Usage
//
// Enabled AWS Regions:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := aws.GetRegions(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// All the regions regardless of the availability
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := aws.GetRegions(ctx, &aws.GetRegionsArgs{
// 			AllRegions: _opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To see regions that are filtered by `"not-opted-in"`, the `allRegions` argument needs to be set to `true` or no results will be returned.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := true
// 		_, err := aws.GetRegions(ctx, &aws.GetRegionsArgs{
// 			AllRegions: _opt0,
// 			Filters: []aws.GetRegionsFilter{
// 				aws.GetRegionsFilter{
// 					Name: "opt-in-status",
// 					Values: []string{
// 						"not-opted-in",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetRegions(ctx *pulumi.Context, args *GetRegionsArgs, opts ...pulumi.InvokeOption) (*GetRegionsResult, error) {
	var rv GetRegionsResult
	err := ctx.Invoke("aws:index/getRegions:getRegions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRegions.
type GetRegionsArgs struct {
	// If true the source will query all regions regardless of availability.
	AllRegions *bool `pulumi:"allRegions"`
	// Configuration block(s) to use as filters. Detailed below.
	Filters []GetRegionsFilter `pulumi:"filters"`
}

// A collection of values returned by getRegions.
type GetRegionsResult struct {
	AllRegions *bool              `pulumi:"allRegions"`
	Filters    []GetRegionsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Names of regions that meets the criteria.
	Names []string `pulumi:"names"`
}
