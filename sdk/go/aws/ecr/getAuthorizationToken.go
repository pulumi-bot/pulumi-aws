// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The ECR Authorization Token data source allows the authorization token, proxy endpoint, token expiration date, user name and password to be retrieved for an ECR repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ecr"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ecr.GetAuthorizationToken(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
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
	// AWS account ID of the ECR Repository. If not specified the default account is assumed.
	RegistryId *string `pulumi:"registryId"`
}

// A collection of values returned by getAuthorizationToken.
type GetAuthorizationTokenResult struct {
	// Temporary IAM authentication credentials to access the ECR repository encoded in base64 in the form of `user_name:password`.
	AuthorizationToken string `pulumi:"authorizationToken"`
	// The time in UTC RFC3339 format when the authorization token expires.
	ExpiresAt string `pulumi:"expiresAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Password decoded from the authorization token.
	Password string `pulumi:"password"`
	// The registry URL to use in the docker login command.
	ProxyEndpoint string  `pulumi:"proxyEndpoint"`
	RegistryId    *string `pulumi:"registryId"`
	// User name decoded from the authorization token.
	UserName string `pulumi:"userName"`
}

func GetAuthorizationTokenOutput(ctx *pulumi.Context, args GetAuthorizationTokenOutputArgs, opts ...pulumi.InvokeOption) GetAuthorizationTokenResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAuthorizationTokenResult, error) {
			args := v.(GetAuthorizationTokenArgs)
			r, err := GetAuthorizationToken(ctx, &args, opts...)
			return *r, err
		}).(GetAuthorizationTokenResultOutput)
}

// A collection of arguments for invoking getAuthorizationToken.
type GetAuthorizationTokenOutputArgs struct {
	// AWS account ID of the ECR Repository. If not specified the default account is assumed.
	RegistryId pulumi.StringPtrInput `pulumi:"registryId"`
}

func (GetAuthorizationTokenOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthorizationTokenArgs)(nil)).Elem()
}

// A collection of values returned by getAuthorizationToken.
type GetAuthorizationTokenResultOutput struct{ *pulumi.OutputState }

func (GetAuthorizationTokenResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAuthorizationTokenResult)(nil)).Elem()
}

func (o GetAuthorizationTokenResultOutput) ToGetAuthorizationTokenResultOutput() GetAuthorizationTokenResultOutput {
	return o
}

func (o GetAuthorizationTokenResultOutput) ToGetAuthorizationTokenResultOutputWithContext(ctx context.Context) GetAuthorizationTokenResultOutput {
	return o
}

// Temporary IAM authentication credentials to access the ECR repository encoded in base64 in the form of `user_name:password`.
func (o GetAuthorizationTokenResultOutput) AuthorizationToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.AuthorizationToken }).(pulumi.StringOutput)
}

// The time in UTC RFC3339 format when the authorization token expires.
func (o GetAuthorizationTokenResultOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.ExpiresAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAuthorizationTokenResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.Id }).(pulumi.StringOutput)
}

// Password decoded from the authorization token.
func (o GetAuthorizationTokenResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.Password }).(pulumi.StringOutput)
}

// The registry URL to use in the docker login command.
func (o GetAuthorizationTokenResultOutput) ProxyEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.ProxyEndpoint }).(pulumi.StringOutput)
}

func (o GetAuthorizationTokenResultOutput) RegistryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) *string { return v.RegistryId }).(pulumi.StringPtrOutput)
}

// User name decoded from the authorization token.
func (o GetAuthorizationTokenResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetAuthorizationTokenResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAuthorizationTokenResultOutput{})
}
