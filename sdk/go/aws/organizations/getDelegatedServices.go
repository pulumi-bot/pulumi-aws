// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package organizations

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a list the AWS services for which the specified account is a delegated administrator
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.GetDelegatedServices(ctx, &organizations.GetDelegatedServicesArgs{
// 			AccountId: "AWS ACCOUNT ID",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDelegatedServices(ctx *pulumi.Context, args *GetDelegatedServicesArgs, opts ...pulumi.InvokeOption) (*GetDelegatedServicesResult, error) {
	var rv GetDelegatedServicesResult
	err := ctx.Invoke("aws:organizations/getDelegatedServices:getDelegatedServices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDelegatedServices.
type GetDelegatedServicesArgs struct {
	// The account ID number of a delegated administrator account in the organization.
	AccountId string `pulumi:"accountId"`
}

// A collection of values returned by getDelegatedServices.
type GetDelegatedServicesResult struct {
	AccountId string `pulumi:"accountId"`
	// The services for which the account is a delegated administrator, which have the following attributes:
	DelegatedServices []GetDelegatedServicesDelegatedService `pulumi:"delegatedServices"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

func GetDelegatedServicesApply(ctx *pulumi.Context, args GetDelegatedServicesApplyInput, opts ...pulumi.InvokeOption) GetDelegatedServicesResultOutput {
	return args.ToGetDelegatedServicesApplyOutput().ApplyT(func(v GetDelegatedServicesArgs) (GetDelegatedServicesResult, error) {
		r, err := GetDelegatedServices(ctx, &v, opts...)
		return *r, err

	}).(GetDelegatedServicesResultOutput)
}

// GetDelegatedServicesApplyInput is an input type that accepts GetDelegatedServicesApplyArgs and GetDelegatedServicesApplyOutput values.
// You can construct a concrete instance of `GetDelegatedServicesApplyInput` via:
//
//          GetDelegatedServicesApplyArgs{...}
type GetDelegatedServicesApplyInput interface {
	pulumi.Input

	ToGetDelegatedServicesApplyOutput() GetDelegatedServicesApplyOutput
	ToGetDelegatedServicesApplyOutputWithContext(context.Context) GetDelegatedServicesApplyOutput
}

// A collection of arguments for invoking getDelegatedServices.
type GetDelegatedServicesApplyArgs struct {
	// The account ID number of a delegated administrator account in the organization.
	AccountId pulumi.StringInput `pulumi:"accountId"`
}

func (GetDelegatedServicesApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDelegatedServicesArgs)(nil)).Elem()
}

func (i GetDelegatedServicesApplyArgs) ToGetDelegatedServicesApplyOutput() GetDelegatedServicesApplyOutput {
	return i.ToGetDelegatedServicesApplyOutputWithContext(context.Background())
}

func (i GetDelegatedServicesApplyArgs) ToGetDelegatedServicesApplyOutputWithContext(ctx context.Context) GetDelegatedServicesApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDelegatedServicesApplyOutput)
}

// A collection of arguments for invoking getDelegatedServices.
type GetDelegatedServicesApplyOutput struct{ *pulumi.OutputState }

func (GetDelegatedServicesApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDelegatedServicesArgs)(nil)).Elem()
}

func (o GetDelegatedServicesApplyOutput) ToGetDelegatedServicesApplyOutput() GetDelegatedServicesApplyOutput {
	return o
}

func (o GetDelegatedServicesApplyOutput) ToGetDelegatedServicesApplyOutputWithContext(ctx context.Context) GetDelegatedServicesApplyOutput {
	return o
}

// The account ID number of a delegated administrator account in the organization.
func (o GetDelegatedServicesApplyOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDelegatedServicesArgs) string { return v.AccountId }).(pulumi.StringOutput)
}

// A collection of values returned by getDelegatedServices.
type GetDelegatedServicesResultOutput struct{ *pulumi.OutputState }

func (GetDelegatedServicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDelegatedServicesResult)(nil)).Elem()
}

func (o GetDelegatedServicesResultOutput) ToGetDelegatedServicesResultOutput() GetDelegatedServicesResultOutput {
	return o
}

func (o GetDelegatedServicesResultOutput) ToGetDelegatedServicesResultOutputWithContext(ctx context.Context) GetDelegatedServicesResultOutput {
	return o
}

func (o GetDelegatedServicesResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDelegatedServicesResult) string { return v.AccountId }).(pulumi.StringOutput)
}

// The services for which the account is a delegated administrator, which have the following attributes:
func (o GetDelegatedServicesResultOutput) DelegatedServices() GetDelegatedServicesDelegatedServiceArrayOutput {
	return o.ApplyT(func(v GetDelegatedServicesResult) []GetDelegatedServicesDelegatedService { return v.DelegatedServices }).(GetDelegatedServicesDelegatedServiceArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDelegatedServicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDelegatedServicesResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDelegatedServicesApplyOutput{})
	pulumi.RegisterOutputType(GetDelegatedServicesResultOutput{})
}
