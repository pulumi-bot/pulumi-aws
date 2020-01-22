// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ecr

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetCredentials(ctx *pulumi.Context, args *GetCredentialsArgs, opts ...pulumi.InvokeOption) (*GetCredentialsResult, error) {
	var rv GetCredentialsResult
	err := ctx.Invoke("aws:ecr/getCredentials:getCredentials", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCredentials.
type GetCredentialsArgs struct {
	RegistryId string `pulumi:"registryId"`
}


// A collection of values returned by getCredentials.
type GetCredentialsResult struct {
	AuthorizationToken string `pulumi:"authorizationToken"`
	ExpiresAt string `pulumi:"expiresAt"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	ProxyEndpoint string `pulumi:"proxyEndpoint"`
	RegistryId string `pulumi:"registryId"`
}

