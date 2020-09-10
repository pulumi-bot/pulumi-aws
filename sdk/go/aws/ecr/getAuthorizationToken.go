// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetAuthorizationToken(ctx *pulumi.Context, args *GetAuthorizationTokenArgs, opts ...pulumi.InvokeOption) (*GetAuthorizationTokenResult, error) {
	var rv GetAuthorizationTokenResult
	err := ctx.Invoke("aws:ecr/getAuthorizationToken:getAuthorizationToken", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthorizationToken.
type GetAuthorizationTokenArgs struct {
	RegistryId *string `pulumi:"registryId"`
}

// A collection of values returned by getAuthorizationToken.
type GetAuthorizationTokenResult struct {
	AuthorizationToken string `pulumi:"authorizationToken"`
	ExpiresAt          string `pulumi:"expiresAt"`
	// The provider-assigned unique ID for this managed resource.
	Id            string  `pulumi:"id"`
	Password      string  `pulumi:"password"`
	ProxyEndpoint string  `pulumi:"proxyEndpoint"`
	RegistryId    *string `pulumi:"registryId"`
	UserName      string  `pulumi:"userName"`
}
