// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Image Builder Distribution Configuration.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := imagebuilder.LookupDistributionConfiguration(ctx, &imagebuilder.LookupDistributionConfigurationArgs{
// 			Arn: "arn:aws:imagebuilder:us-west-2:aws:distribution-configuration/example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupDistributionConfiguration(ctx *pulumi.Context, args *LookupDistributionConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDistributionConfigurationResult, error) {
	var rv LookupDistributionConfigurationResult
	err := ctx.Invoke("aws:imagebuilder/getDistributionConfiguration:getDistributionConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDistributionConfiguration.
type LookupDistributionConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the distribution configuration.
	Arn string `pulumi:"arn"`
	// Key-value map of resource tags for the distribution configuration.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getDistributionConfiguration.
type LookupDistributionConfigurationResult struct {
	Arn string `pulumi:"arn"`
	// Date the distribution configuration was created.
	DateCreated string `pulumi:"dateCreated"`
	// Date the distribution configuration was updated.
	DateUpdated string `pulumi:"dateUpdated"`
	// Description to apply to distributed AMI.
	Description string `pulumi:"description"`
	// Set of distributions.
	Distributions []GetDistributionConfigurationDistribution `pulumi:"distributions"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the distribution configuration.
	Name string `pulumi:"name"`
	// Key-value map of resource tags for the distribution configuration.
	Tags map[string]string `pulumi:"tags"`
}

func LookupDistributionConfigurationApply(ctx *pulumi.Context, args LookupDistributionConfigurationApplyInput, opts ...pulumi.InvokeOption) LookupDistributionConfigurationResultOutput {
	return args.ToLookupDistributionConfigurationApplyOutput().ApplyT(func(v LookupDistributionConfigurationArgs) (LookupDistributionConfigurationResult, error) {
		r, err := LookupDistributionConfiguration(ctx, &v, opts...)
		return *r, err

	}).(LookupDistributionConfigurationResultOutput)
}

// LookupDistributionConfigurationApplyInput is an input type that accepts LookupDistributionConfigurationApplyArgs and LookupDistributionConfigurationApplyOutput values.
// You can construct a concrete instance of `LookupDistributionConfigurationApplyInput` via:
//
//          LookupDistributionConfigurationApplyArgs{...}
type LookupDistributionConfigurationApplyInput interface {
	pulumi.Input

	ToLookupDistributionConfigurationApplyOutput() LookupDistributionConfigurationApplyOutput
	ToLookupDistributionConfigurationApplyOutputWithContext(context.Context) LookupDistributionConfigurationApplyOutput
}

// A collection of arguments for invoking getDistributionConfiguration.
type LookupDistributionConfigurationApplyArgs struct {
	// Amazon Resource Name (ARN) of the distribution configuration.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Key-value map of resource tags for the distribution configuration.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupDistributionConfigurationApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributionConfigurationArgs)(nil)).Elem()
}

func (i LookupDistributionConfigurationApplyArgs) ToLookupDistributionConfigurationApplyOutput() LookupDistributionConfigurationApplyOutput {
	return i.ToLookupDistributionConfigurationApplyOutputWithContext(context.Background())
}

func (i LookupDistributionConfigurationApplyArgs) ToLookupDistributionConfigurationApplyOutputWithContext(ctx context.Context) LookupDistributionConfigurationApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupDistributionConfigurationApplyOutput)
}

// A collection of arguments for invoking getDistributionConfiguration.
type LookupDistributionConfigurationApplyOutput struct{ *pulumi.OutputState }

func (LookupDistributionConfigurationApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributionConfigurationArgs)(nil)).Elem()
}

func (o LookupDistributionConfigurationApplyOutput) ToLookupDistributionConfigurationApplyOutput() LookupDistributionConfigurationApplyOutput {
	return o
}

func (o LookupDistributionConfigurationApplyOutput) ToLookupDistributionConfigurationApplyOutputWithContext(ctx context.Context) LookupDistributionConfigurationApplyOutput {
	return o
}

// Amazon Resource Name (ARN) of the distribution configuration.
func (o LookupDistributionConfigurationApplyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationArgs) string { return v.Arn }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the distribution configuration.
func (o LookupDistributionConfigurationApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getDistributionConfiguration.
type LookupDistributionConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDistributionConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDistributionConfigurationResult)(nil)).Elem()
}

func (o LookupDistributionConfigurationResultOutput) ToLookupDistributionConfigurationResultOutput() LookupDistributionConfigurationResultOutput {
	return o
}

func (o LookupDistributionConfigurationResultOutput) ToLookupDistributionConfigurationResultOutputWithContext(ctx context.Context) LookupDistributionConfigurationResultOutput {
	return o
}

func (o LookupDistributionConfigurationResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Date the distribution configuration was created.
func (o LookupDistributionConfigurationResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

// Date the distribution configuration was updated.
func (o LookupDistributionConfigurationResultOutput) DateUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.DateUpdated }).(pulumi.StringOutput)
}

// Description to apply to distributed AMI.
func (o LookupDistributionConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

// Set of distributions.
func (o LookupDistributionConfigurationResultOutput) Distributions() GetDistributionConfigurationDistributionArrayOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) []GetDistributionConfigurationDistribution {
		return v.Distributions
	}).(GetDistributionConfigurationDistributionArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDistributionConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the distribution configuration.
func (o LookupDistributionConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the distribution configuration.
func (o LookupDistributionConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDistributionConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDistributionConfigurationApplyOutput{})
	pulumi.RegisterOutputType(LookupDistributionConfigurationResultOutput{})
}
