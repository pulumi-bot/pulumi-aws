// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to get the id of a VPC Link in
// API Gateway. To fetch the VPC Link you must provide a name to match against.
// As there is no unique name constraint on API Gateway VPC Links this data source will
// error if there is more than one match.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.LookupVpcLink(ctx, &apigateway.LookupVpcLinkArgs{
// 			Name: "my-vpc-link",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVpcLink(ctx *pulumi.Context, args *LookupVpcLinkArgs, opts ...pulumi.InvokeOption) (*LookupVpcLinkResult, error) {
	var rv LookupVpcLinkResult
	err := ctx.Invoke("aws:apigateway/getVpcLink:getVpcLink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpcLink.
type LookupVpcLinkArgs struct {
	// The name of the API Gateway VPC Link to look up. If no API Gateway VPC Link is found with this name, an error will be returned.
	// If multiple API Gateway VPC Links are found with this name, an error will be returned.
	Name string `pulumi:"name"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getVpcLink.
type LookupVpcLinkResult struct {
	// The description of the VPC link.
	Description string `pulumi:"description"`
	// Set to the ID of the found API Gateway VPC Link.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The status of the VPC link.
	Status string `pulumi:"status"`
	// The status message of the VPC link.
	StatusMessage string `pulumi:"statusMessage"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArns []string `pulumi:"targetArns"`
}
