// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about an Image Builder Infrastructure Configuration.
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
// 		_, err := imagebuilder.LookupInfrastructureConfiguration(ctx, &imagebuilder.LookupInfrastructureConfigurationArgs{
// 			Arn: "arn:aws:imagebuilder:us-west-2:aws:infrastructure-configuration/example",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupInfrastructureConfiguration(ctx *pulumi.Context, args *LookupInfrastructureConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupInfrastructureConfigurationResult, error) {
	var rv LookupInfrastructureConfigurationResult
	err := ctx.Invoke("aws:imagebuilder/getInfrastructureConfiguration:getInfrastructureConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInfrastructureConfiguration.
type LookupInfrastructureConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn string `pulumi:"arn"`
	// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
	ResourceTags map[string]string `pulumi:"resourceTags"`
	// Key-value map of resource tags for the infrastructure configuration.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getInfrastructureConfiguration.
type LookupInfrastructureConfigurationResult struct {
	Arn string `pulumi:"arn"`
	// Date the infrastructure configuration was updated.
	DateCreated string `pulumi:"dateCreated"`
	DateUpdated string `pulumi:"dateUpdated"`
	// Description of the infrastructure configuration.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the IAM Instance Profile associated with the configuration.
	InstanceProfileName string `pulumi:"instanceProfileName"`
	// Set of EC2 Instance Types associated with the configuration.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// Name of the EC2 Key Pair associated with the configuration.
	KeyPair string `pulumi:"keyPair"`
	// Nested list of logging settings.
	Loggings []GetInfrastructureConfigurationLogging `pulumi:"loggings"`
	// Name of the infrastructure configuration.
	Name string `pulumi:"name"`
	// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
	ResourceTags map[string]string `pulumi:"resourceTags"`
	// Set of EC2 Security Group identifiers associated with the configuration.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// Amazon Resource Name (ARN) of the SNS Topic associated with the configuration.
	SnsTopicArn string `pulumi:"snsTopicArn"`
	// Identifier of the EC2 Subnet associated with the configuration.
	SubnetId string `pulumi:"subnetId"`
	// Key-value map of resource tags for the infrastructure configuration.
	Tags map[string]string `pulumi:"tags"`
	// Whether instances are terminated on failure.
	TerminateInstanceOnFailure bool `pulumi:"terminateInstanceOnFailure"`
}

func LookupInfrastructureConfigurationApply(ctx *pulumi.Context, args LookupInfrastructureConfigurationApplyInput, opts ...pulumi.InvokeOption) LookupInfrastructureConfigurationResultOutput {
	return args.ToLookupInfrastructureConfigurationApplyOutput().ApplyT(func(v LookupInfrastructureConfigurationArgs) (LookupInfrastructureConfigurationResult, error) {
		r, err := LookupInfrastructureConfiguration(ctx, &v, opts...)
		return *r, err

	}).(LookupInfrastructureConfigurationResultOutput)
}

// LookupInfrastructureConfigurationApplyInput is an input type that accepts LookupInfrastructureConfigurationApplyArgs and LookupInfrastructureConfigurationApplyOutput values.
// You can construct a concrete instance of `LookupInfrastructureConfigurationApplyInput` via:
//
//          LookupInfrastructureConfigurationApplyArgs{...}
type LookupInfrastructureConfigurationApplyInput interface {
	pulumi.Input

	ToLookupInfrastructureConfigurationApplyOutput() LookupInfrastructureConfigurationApplyOutput
	ToLookupInfrastructureConfigurationApplyOutputWithContext(context.Context) LookupInfrastructureConfigurationApplyOutput
}

// A collection of arguments for invoking getInfrastructureConfiguration.
type LookupInfrastructureConfigurationApplyArgs struct {
	// Amazon Resource Name (ARN) of the infrastructure configuration.
	Arn pulumi.StringInput `pulumi:"arn"`
	// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
	ResourceTags pulumi.StringMapInput `pulumi:"resourceTags"`
	// Key-value map of resource tags for the infrastructure configuration.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupInfrastructureConfigurationApplyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureConfigurationArgs)(nil)).Elem()
}

func (i LookupInfrastructureConfigurationApplyArgs) ToLookupInfrastructureConfigurationApplyOutput() LookupInfrastructureConfigurationApplyOutput {
	return i.ToLookupInfrastructureConfigurationApplyOutputWithContext(context.Background())
}

func (i LookupInfrastructureConfigurationApplyArgs) ToLookupInfrastructureConfigurationApplyOutputWithContext(ctx context.Context) LookupInfrastructureConfigurationApplyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LookupInfrastructureConfigurationApplyOutput)
}

// A collection of arguments for invoking getInfrastructureConfiguration.
type LookupInfrastructureConfigurationApplyOutput struct{ *pulumi.OutputState }

func (LookupInfrastructureConfigurationApplyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureConfigurationArgs)(nil)).Elem()
}

func (o LookupInfrastructureConfigurationApplyOutput) ToLookupInfrastructureConfigurationApplyOutput() LookupInfrastructureConfigurationApplyOutput {
	return o
}

func (o LookupInfrastructureConfigurationApplyOutput) ToLookupInfrastructureConfigurationApplyOutputWithContext(ctx context.Context) LookupInfrastructureConfigurationApplyOutput {
	return o
}

// Amazon Resource Name (ARN) of the infrastructure configuration.
func (o LookupInfrastructureConfigurationApplyOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationArgs) string { return v.Arn }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
func (o LookupInfrastructureConfigurationApplyOutput) ResourceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationArgs) map[string]string { return v.ResourceTags }).(pulumi.StringMapOutput)
}

// Key-value map of resource tags for the infrastructure configuration.
func (o LookupInfrastructureConfigurationApplyOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationArgs) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// A collection of values returned by getInfrastructureConfiguration.
type LookupInfrastructureConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupInfrastructureConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInfrastructureConfigurationResult)(nil)).Elem()
}

func (o LookupInfrastructureConfigurationResultOutput) ToLookupInfrastructureConfigurationResultOutput() LookupInfrastructureConfigurationResultOutput {
	return o
}

func (o LookupInfrastructureConfigurationResultOutput) ToLookupInfrastructureConfigurationResultOutputWithContext(ctx context.Context) LookupInfrastructureConfigurationResultOutput {
	return o
}

func (o LookupInfrastructureConfigurationResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.Arn }).(pulumi.StringOutput)
}

// Date the infrastructure configuration was updated.
func (o LookupInfrastructureConfigurationResultOutput) DateCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.DateCreated }).(pulumi.StringOutput)
}

func (o LookupInfrastructureConfigurationResultOutput) DateUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.DateUpdated }).(pulumi.StringOutput)
}

// Description of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupInfrastructureConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the IAM Instance Profile associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) InstanceProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.InstanceProfileName }).(pulumi.StringOutput)
}

// Set of EC2 Instance Types associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) []string { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// Name of the EC2 Key Pair associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) KeyPair() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.KeyPair }).(pulumi.StringOutput)
}

// Nested list of logging settings.
func (o LookupInfrastructureConfigurationResultOutput) Loggings() GetInfrastructureConfigurationLoggingArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) []GetInfrastructureConfigurationLogging {
		return v.Loggings
	}).(GetInfrastructureConfigurationLoggingArrayOutput)
}

// Name of the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the infrastructure created by the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) ResourceTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) map[string]string { return v.ResourceTags }).(pulumi.StringMapOutput)
}

// Set of EC2 Security Group identifiers associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// Amazon Resource Name (ARN) of the SNS Topic associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) SnsTopicArn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.SnsTopicArn }).(pulumi.StringOutput)
}

// Identifier of the EC2 Subnet associated with the configuration.
func (o LookupInfrastructureConfigurationResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

// Key-value map of resource tags for the infrastructure configuration.
func (o LookupInfrastructureConfigurationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Whether instances are terminated on failure.
func (o LookupInfrastructureConfigurationResultOutput) TerminateInstanceOnFailure() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInfrastructureConfigurationResult) bool { return v.TerminateInstanceOnFailure }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInfrastructureConfigurationApplyOutput{})
	pulumi.RegisterOutputType(LookupInfrastructureConfigurationResultOutput{})
}
