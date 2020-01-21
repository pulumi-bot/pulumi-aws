// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getSubnetIds

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/getSubnetIdsFilter"
)

// `ec2.getSubnetIds` provides a list of ids for a vpcId
// 
// This resource can be useful for getting back a list of subnet ids for a vpc.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/subnet_ids.html.markdown.
func GetSubnetIds(ctx *pulumi.Context, args *GetSubnetIdsArgs, opts ...pulumi.InvokeOption) (*GetSubnetIdsResult, error) {
	var rv GetSubnetIdsResult
	err := ctx.Invoke("aws:ec2/getSubnetIds:getSubnetIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsArgs struct {
	// Custom filter block as described below.
	Filters []ec2getSubnetIdsFilter.GetSubnetIdsFilter `pulumi:"filters"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags map[string]interface{} `pulumi:"tags"`
	// The VPC ID that you want to filter from.
	VpcId string `pulumi:"vpcId"`
}


// A collection of values returned by getSubnetIds.
type GetSubnetIdsResult struct {
	Filters []ec2getSubnetIdsFilter.GetSubnetIdsFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of all the subnet ids found. This data source will fail if none are found.
	Ids []string `pulumi:"ids"`
	Tags map[string]interface{} `pulumi:"tags"`
	VpcId string `pulumi:"vpcId"`
}

