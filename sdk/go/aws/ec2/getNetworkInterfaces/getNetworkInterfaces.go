// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package getNetworkInterfaces

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/ec2/getNetworkInterfacesFilter"
)

func GetNetworkInterfaces(ctx *pulumi.Context, args *GetNetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*GetNetworkInterfacesResult, error) {
	var rv GetNetworkInterfacesResult
	err := ctx.Invoke("aws:ec2/getNetworkInterfaces:getNetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkInterfaces.
type GetNetworkInterfacesArgs struct {
	// Custom filter block as described below.
	Filters []ec2getNetworkInterfacesFilter.GetNetworkInterfacesFilter `pulumi:"filters"`
	// A mapping of tags, each pair of which must exactly match
	// a pair on the desired network interfaces.
	Tags map[string]interface{} `pulumi:"tags"`
}


// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResult struct {
	Filters []ec2getNetworkInterfacesFilter.GetNetworkInterfacesFilter `pulumi:"filters"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of all the network interface ids found. This data source will fail if none are found.
	Ids []string `pulumi:"ids"`
	Tags map[string]interface{} `pulumi:"tags"`
}

