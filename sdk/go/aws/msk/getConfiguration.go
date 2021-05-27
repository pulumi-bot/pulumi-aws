// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information on an Amazon MSK Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/msk"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := msk.LookupConfiguration(ctx, &msk.LookupConfigurationArgs{
// 			Name: "example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	var rv LookupConfigurationResult
	err := ctx.Invoke("aws:msk/getConfiguration:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getConfiguration.
type LookupConfigurationArgs struct {
	// Name of the configuration.
	Name string `pulumi:"name"`
}

// A collection of values returned by getConfiguration.
type LookupConfigurationResult struct {
	// Amazon Resource Name (ARN) of the configuration.
	Arn string `pulumi:"arn"`
	// Description of the configuration.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of Apache Kafka versions which can use this configuration.
	KafkaVersions []string `pulumi:"kafkaVersions"`
	// Latest revision of the configuration.
	LatestRevision int    `pulumi:"latestRevision"`
	Name           string `pulumi:"name"`
	// Contents of the server.properties file.
	ServerProperties string `pulumi:"serverProperties"`
}

func LookupConfigurationOutput(ctx *pulumi.Context, args LookupConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationResult, error) {
			args := v.(LookupConfigurationArgs)
			r, err := LookupConfiguration(ctx, &args, opts...)
			return *r, err
		}).(LookupConfigurationResultOutput)
}

// A collection of arguments for invoking getConfiguration.
type LookupConfigurationOutputArgs struct {
	// Name of the configuration.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getConfiguration.
type LookupConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationResult)(nil)).Elem()
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutput() LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutputWithContext(ctx context.Context) LookupConfigurationResultOutput {
	return o
}

// Amazon Resource Name (ARN) of the configuration.
func (o LookupConfigurationResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Description of the configuration.
func (o LookupConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of Apache Kafka versions which can use this configuration.
func (o LookupConfigurationResultOutput) KafkaVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConfigurationResult) []string { return v.KafkaVersions }).(pulumi.StringArrayOutput)
}

// Latest revision of the configuration.
func (o LookupConfigurationResultOutput) LatestRevision() pulumi.IntOutput {
	return o.ApplyT(func(v LookupConfigurationResult) int { return v.LatestRevision }).(pulumi.IntOutput)
}

func (o LookupConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Contents of the server.properties file.
func (o LookupConfigurationResultOutput) ServerProperties() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConfigurationResult) string { return v.ServerProperties }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationResultOutput{})
}
