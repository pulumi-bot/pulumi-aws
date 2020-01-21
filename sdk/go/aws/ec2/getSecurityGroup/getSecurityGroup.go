// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getSecurityGroup

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/getSecurityGroupFilter"
)

// `ec2.SecurityGroup` provides details about a specific Security Group.
// 
// This resource can prove useful when a module accepts a Security Group id as
// an input variable and needs to, for example, determine the id of the
// VPC that the security group belongs to.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/security_group.html.markdown.
func GetSecurityGroup(ctx *pulumi.Context, args *GetSecurityGroupArgs, opts ...pulumi.InvokeOption) (*GetSecurityGroupResult, error) {
	var rv GetSecurityGroupResult
	err := ctx.Invoke("aws:ec2/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type GetSecurityGroupArgs struct {
	// Custom filter block as described below.
	Filters []ec2getSecurityGroupFilter.GetSecurityGroupFilter `pulumi:"filters"`
	// The id of the specific security group to retrieve.
	Id *string `pulumi:"id"`
	// The name of the field to filter by, as defined by
	// [the underlying AWS API](http://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html).
	Name *string `pulumi:"name"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired security group.
	Tags map[string]interface{} `pulumi:"tags"`
	// The id of the VPC that the desired security group belongs to.
	VpcId *string `pulumi:"vpcId"`
}


// A collection of values returned by getSecurityGroup.
type GetSecurityGroupResult struct {
	// The computed ARN of the security group.
	Arn string `pulumi:"arn"`
	// The description of the security group.
	Description string `pulumi:"description"`
	Filters []ec2getSecurityGroupFilter.GetSecurityGroupFilter `pulumi:"filters"`
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	Tags map[string]interface{} `pulumi:"tags"`
	VpcId string `pulumi:"vpcId"`
}

