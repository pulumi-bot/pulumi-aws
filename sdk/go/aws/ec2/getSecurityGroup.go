// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// `ec2.SecurityGroup` provides details about a specific Security Group.
//
// This resource can prove useful when a module accepts a Security Group id as
// an input variable and needs to, for example, determine the id of the
// VPC that the security group belongs to.
func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("aws:ec2/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupArgs struct {
	// Custom filter block as described below.
	Filters []GetSecurityGroupFilter `pulumi:"filters"`
	// The id of the specific security group to retrieve.
	Id *string `pulumi:"id"`
	// The name that the desired security group must have.
	Name *string `pulumi:"name"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired security group.
	Tags map[string]interface{} `pulumi:"tags"`
	// The id of the VPC that the desired security group belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResult struct {
	// The computed ARN of the security group.
	Arn string `pulumi:"arn"`
	// The description of the security group.
	Description string                   `pulumi:"description"`
	Filters     []GetSecurityGroupFilter `pulumi:"filters"`
	Id          string                   `pulumi:"id"`
	Name        string                   `pulumi:"name"`
	Tags        map[string]interface{}   `pulumi:"tags"`
	VpcId       string                   `pulumi:"vpcId"`
}
