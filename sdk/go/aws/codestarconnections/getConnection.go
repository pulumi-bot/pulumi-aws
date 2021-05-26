// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestarconnections

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about CodeStar Connection.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/codestarconnections"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := codestarconnections.LookupConnection(ctx, &codestarconnections.LookupConnectionArgs{
// 			Arn: aws_codestarconnections_connection.Example.Arn,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupConnection(ctx *pulumi.Context, args *LookupConnectionArgs, opts ...pulumi.InvokeOption) (*LookupConnectionResult, error) {
	var rv LookupConnectionResult
	err := ctx.Invoke("aws:codestarconnections/getConnection:getConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConnection.
type LookupConnectionArgs struct {
	// The CodeStar Connection ARN.
	Arn string `pulumi:"arn"`
	// Map of key-value resource tags to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getConnection.
type LookupConnectionResult struct {
	Arn string `pulumi:"arn"`
	// The CodeStar Connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
	ConnectionStatus string `pulumi:"connectionStatus"`
	// The Amazon Resource Name (ARN) of the host associated with the connection.
	HostArn string `pulumi:"hostArn"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the CodeStar Connection. The name is unique in the calling AWS account.
	Name string `pulumi:"name"`
	// The name of the external provider where your third-party code repository is configured. Possible values are `Bitbucket`, `GitHub`, or `GitHubEnterpriseServer`.
	ProviderType string `pulumi:"providerType"`
	// Map of key-value resource tags to associate with the resource.
	Tags map[string]string `pulumi:"tags"`
}

func LookupConnectionApply(ctx *pulumi.Context, args LookupConnectionApplyInput, opts ...pulumi.InvokeOption) LookupConnectionResultOutput {
	return args.ToLookupConnectionApplyOutput().ApplyT(func(v LookupConnectionArgs) (LookupConnectionResult, error) {
		r, err := LookupConnection(ctx, &v, opts...)
		return *r, err

	}).(LookupConnectionResultOutput)
}

// LookupConnectionApplyInput is an input type that accepts LookupConnectionApplyArgs and LookupConnectionApplyOutput values.
// You can construct a concrete instance of `LookupConnectionApplyInput` via:
//
//          LookupConnectionApplyArgs{...}
type LookupConnectionApplyInput interface {
	pulumi.Input

	ToLookupConnectionApplyOutput() LookupConnectionApplyOutput
	ToLookupConnectionApplyOutputWithContext(context.Context) LookupConnectionApplyOutput
}

// A collection of arguments for invoking getConnection.
type LookupConnectionApplyArgs struct {
	// The CodeStar Connection ARN.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Map of key-value resource tags to associate with the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupConnectionApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

func (i LookupConnectionApplyArgs) ToLookupConnectionApplyOutput() LookupConnectionApplyOutput {
	return i.ToLookupConnectionApplyOutputWithContext(context.Background())
}

func (i LookupConnectionApplyArgs) ToLookupConnectionApplyOutputWithContext(ctx context.Context) LookupConnectionApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupConnectionApplyOutput)
}

// A collection of arguments for invoking getConnection.
type LookupConnectionApplyOutput struct{ *pulumi.OutputState }

func (LookupConnectionApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionArgs)(nil)).Elem()
}

func (o LookupConnectionApplyOutput) ToLookupConnectionApplyOutput() LookupConnectionApplyOutput {
	return o
}

func (o LookupConnectionApplyOutput) ToLookupConnectionApplyOutputWithContext(ctx context.Context) LookupConnectionApplyOutput {
	return o
}

// The CodeStar Connection ARN.
func (o LookupConnectionApplyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionArgs) string { return v.Arn }).(pulumi.StringOutput)
}

// Map of key-value resource tags to associate with the resource.
func (o LookupConnectionApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getConnection.
type LookupConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConnectionResult)(nil)).Elem()
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutput() LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) ToLookupConnectionResultOutputWithContext(ctx context.Context) LookupConnectionResultOutput {
	return o
}

func (o LookupConnectionResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The CodeStar Connection status. Possible values are `PENDING`, `AVAILABLE` and `ERROR`.
func (o LookupConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the host associated with the connection.
func (o LookupConnectionResultOutput) HostArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.HostArn }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the CodeStar Connection. The name is unique in the calling AWS account.
func (o LookupConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// The name of the external provider where your third-party code repository is configured. Possible values are `Bitbucket`, `GitHub`, or `GitHubEnterpriseServer`.
func (o LookupConnectionResultOutput) ProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConnectionResult) string { return v.ProviderType }).(pulumi.StringOutput)
}

// Map of key-value resource tags to associate with the resource.
func (o LookupConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConnectionApplyOutput{})
	pulumi.RegisterOutputType(LookupConnectionResultOutput{})
}
