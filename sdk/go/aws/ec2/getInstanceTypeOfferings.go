// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about EC2 Instance Type Offerings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "availability-zone-id"
// 		_, err := ec2.GetInstanceTypeOfferings(ctx, &ec2.GetInstanceTypeOfferingsArgs{
// 			Filters: []ec2.GetInstanceTypeOfferingsFilter{
// 				ec2.GetInstanceTypeOfferingsFilter{
// 					Name: "instance-type",
// 					Values: []string{
// 						"t2.micro",
// 						"t3.micro",
// 					},
// 				},
// 				ec2.GetInstanceTypeOfferingsFilter{
// 					Name: "location",
// 					Values: []string{
// 						"usw2-az4",
// 					},
// 				},
// 			},
// 			LocationType: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetInstanceTypeOfferings(ctx *pulumi.Context, args *GetInstanceTypeOfferingsArgs, opts ...pulumi.InvokeOption) (*GetInstanceTypeOfferingsResult, error) {
	var rv GetInstanceTypeOfferingsResult
	err := ctx.Invoke("aws:ec2/getInstanceTypeOfferings:getInstanceTypeOfferings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstanceTypeOfferings.
type GetInstanceTypeOfferingsArgs struct {
	// One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
	Filters []GetInstanceTypeOfferingsFilter `pulumi:"filters"`
	// Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
	LocationType *string `pulumi:"locationType"`
}

// A collection of values returned by getInstanceTypeOfferings.
type GetInstanceTypeOfferingsResult struct {
	Filters []GetInstanceTypeOfferingsFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Set of EC2 Instance Types.
	InstanceTypes []string `pulumi:"instanceTypes"`
	LocationType  *string  `pulumi:"locationType"`
}
