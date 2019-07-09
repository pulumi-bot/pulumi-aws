// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// `aws_subnet_ids` provides a list of ids for a vpc_id
// 
// This resource can be useful for getting back a list of subnet ids for a vpc.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/subnet_ids.html.markdown.
func LookupSubnetIds(ctx *pulumi.Context, args *GetSubnetIdsArgs) (*GetSubnetIdsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["filters"] = args.Filters
		inputs["tags"] = args.Tags
		inputs["vpcId"] = args.VpcId
	}
	outputs, err := ctx.Invoke("aws:ec2/getSubnetIds:getSubnetIds", inputs)
	if err != nil {
		return nil, err
	}
	return &GetSubnetIdsResult{
		Filters: outputs["filters"],
		Ids: outputs["ids"],
		Tags: outputs["tags"],
		VpcId: outputs["vpcId"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getSubnetIds.
type GetSubnetIdsArgs struct {
	// Custom filter block as described below.
	Filters interface{}
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired subnets.
	Tags interface{}
	// The VPC ID that you want to filter from.
	VpcId interface{}
}

// A collection of values returned by getSubnetIds.
type GetSubnetIdsResult struct {
	Filters interface{}
	// A list of all the subnet ids found. This data source will fail if none are found.
	Ids interface{}
	Tags interface{}
	VpcId interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
