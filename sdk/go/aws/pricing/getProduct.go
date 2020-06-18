// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pricing

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the pricing information of all products in AWS.
// This data source is only available in a us-east-1 or ap-south-1 provider.
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
// 		_, err := pricing.LookupProduct(ctx, &pricing.LookupProductArgs{
// 			Filters: pricing.getProductFilterArray{
// 				&pricing.LookupProductFilter{
// 					Field: "instanceType",
// 					Value: "c5.xlarge",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "operatingSystem",
// 					Value: "Linux",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "location",
// 					Value: "US East (N. Virginia)",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "preInstalledSw",
// 					Value: "NA",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "licenseModel",
// 					Value: "No License required",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "tenancy",
// 					Value: "Shared",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "capacitystatus",
// 					Value: "Used",
// 				},
// 			},
// 			ServiceCode: "AmazonEC2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
// 		_, err := pricing.LookupProduct(ctx, &pricing.LookupProductArgs{
// 			Filters: pricing.getProductFilterArray{
// 				&pricing.LookupProductFilter{
// 					Field: "instanceType",
// 					Value: "ds1.xlarge",
// 				},
// 				&pricing.LookupProductFilter{
// 					Field: "location",
// 					Value: "US East (N. Virginia)",
// 				},
// 			},
// 			ServiceCode: "AmazonRedshift",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	var rv GetProductResult
	err := ctx.Invoke("aws:pricing/getProduct:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProduct.
type GetProductArgs struct {
	// A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
	Filters []GetProductFilter `pulumi:"filters"`
	// The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
	ServiceCode string `pulumi:"serviceCode"`
}

// A collection of values returned by getProduct.
type GetProductResult struct {
	Filters []GetProductFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set to the product returned from the API.
	Result      string `pulumi:"result"`
	ServiceCode string `pulumi:"serviceCode"`
}
