// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getInstances

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/getInstancesFilter"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/instances.html.markdown.
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	var rv GetInstancesResult
	err := ctx.Invoke("aws:ec2/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// One or more name/value pairs to use as filters. There are
	// several valid keys, for a full reference, check out
	// [describe-instances in the AWS CLI reference][1].
	Filters []ec2getInstancesFilter.GetInstancesFilter `pulumi:"filters"`
	// A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
	InstanceStateNames []string `pulumi:"instanceStateNames"`
	// A mapping of tags, each pair of which must
	// exactly match a pair on desired instances.
	InstanceTags map[string]interface{} `pulumi:"instanceTags"`
}


// A collection of values returned by getInstances.
type GetInstancesResult struct {
	Filters []ec2getInstancesFilter.GetInstancesFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// IDs of instances found through the filter
	Ids []string `pulumi:"ids"`
	InstanceStateNames []string `pulumi:"instanceStateNames"`
	InstanceTags map[string]interface{} `pulumi:"instanceTags"`
	// Private IP addresses of instances found through the filter
	PrivateIps []string `pulumi:"privateIps"`
	// Public IP addresses of instances found through the filter
	PublicIps []string `pulumi:"publicIps"`
}

