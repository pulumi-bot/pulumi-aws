// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
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
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	Tags    map[string]string            `pulumi:"tags"`
}

// A collection of values returned by getNetworkInterfaces.
type GetNetworkInterfacesResult struct {
	Filters []GetNetworkInterfacesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id   string            `pulumi:"id"`
	Ids  []string          `pulumi:"ids"`
	Tags map[string]string `pulumi:"tags"`
}
