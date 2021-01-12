// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.
//
// ## Example Usage
// ### Add analytics configuration for entire S3 bucket and export results to a second S3 bucket
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := s3.NewBucket(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		analytics, err := s3.NewBucket(ctx, "analytics", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewAnalyticsConfiguration(ctx, "example_entire_bucket", &s3.AnalyticsConfigurationArgs{
// 			Bucket: example.Bucket,
// 			StorageClassAnalysis: &s3.AnalyticsConfigurationStorageClassAnalysisArgs{
// 				DataExport: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportArgs{
// 					Destination: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs{
// 						S3BucketDestination: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs{
// 							BucketArn: analytics.Arn,
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Add analytics configuration with S3 bucket object filter
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := s3.NewBucket(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewAnalyticsConfiguration(ctx, "example_filtered", &s3.AnalyticsConfigurationArgs{
// 			Bucket: example.Bucket,
// 			Filter: &s3.AnalyticsConfigurationFilterArgs{
// 				Prefix: pulumi.String("documents/"),
// 				Tags: pulumi.StringMap{
// 					"priority": pulumi.String("high"),
// 					"class":    pulumi.String("blue"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// S3 bucket analytics configurations can be imported using `bucket:analytics`, e.g.
//
// ```sh
//  $ pulumi import aws:s3/analyticsConfiguration:AnalyticsConfiguration my-bucket-entire-bucket my-bucket:EntireBucket
// ```
type AnalyticsConfiguration struct {
	pulumi.CustomResourceState

	// The name of the bucket this analytics configuration is associated with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter AnalyticsConfigurationFilterPtrOutput `pulumi:"filter"`
	// Unique identifier of the analytics configuration for the bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis AnalyticsConfigurationStorageClassAnalysisPtrOutput `pulumi:"storageClassAnalysis"`
}

// NewAnalyticsConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAnalyticsConfiguration(ctx *pulumi.Context,
	name string, args *AnalyticsConfigurationArgs, opts ...pulumi.ResourceOption) (*AnalyticsConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource AnalyticsConfiguration
	err := ctx.RegisterResource("aws:s3/analyticsConfiguration:AnalyticsConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalyticsConfiguration gets an existing AnalyticsConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalyticsConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalyticsConfigurationState, opts ...pulumi.ResourceOption) (*AnalyticsConfiguration, error) {
	var resource AnalyticsConfiguration
	err := ctx.ReadResource("aws:s3/analyticsConfiguration:AnalyticsConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AnalyticsConfiguration resources.
type analyticsConfigurationState struct {
	// The name of the bucket this analytics configuration is associated with.
	Bucket *string `pulumi:"bucket"`
	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *AnalyticsConfigurationFilter `pulumi:"filter"`
	// Unique identifier of the analytics configuration for the bucket.
	Name *string `pulumi:"name"`
	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis *AnalyticsConfigurationStorageClassAnalysis `pulumi:"storageClassAnalysis"`
}

type AnalyticsConfigurationState struct {
	// The name of the bucket this analytics configuration is associated with.
	Bucket pulumi.StringPtrInput
	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter AnalyticsConfigurationFilterPtrInput
	// Unique identifier of the analytics configuration for the bucket.
	Name pulumi.StringPtrInput
	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis AnalyticsConfigurationStorageClassAnalysisPtrInput
}

func (AnalyticsConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsConfigurationState)(nil)).Elem()
}

type analyticsConfigurationArgs struct {
	// The name of the bucket this analytics configuration is associated with.
	Bucket string `pulumi:"bucket"`
	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter *AnalyticsConfigurationFilter `pulumi:"filter"`
	// Unique identifier of the analytics configuration for the bucket.
	Name *string `pulumi:"name"`
	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis *AnalyticsConfigurationStorageClassAnalysis `pulumi:"storageClassAnalysis"`
}

// The set of arguments for constructing a AnalyticsConfiguration resource.
type AnalyticsConfigurationArgs struct {
	// The name of the bucket this analytics configuration is associated with.
	Bucket pulumi.StringInput
	// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
	Filter AnalyticsConfigurationFilterPtrInput
	// Unique identifier of the analytics configuration for the bucket.
	Name pulumi.StringPtrInput
	// Configuration for the analytics data export (documented below).
	StorageClassAnalysis AnalyticsConfigurationStorageClassAnalysisPtrInput
}

func (AnalyticsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analyticsConfigurationArgs)(nil)).Elem()
}

type AnalyticsConfigurationInput interface {
	pulumi.Input

	ToAnalyticsConfigurationOutput() AnalyticsConfigurationOutput
	ToAnalyticsConfigurationOutputWithContext(ctx context.Context) AnalyticsConfigurationOutput
}

func (*AnalyticsConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsConfiguration)(nil))
}

func (i *AnalyticsConfiguration) ToAnalyticsConfigurationOutput() AnalyticsConfigurationOutput {
	return i.ToAnalyticsConfigurationOutputWithContext(context.Background())
}

func (i *AnalyticsConfiguration) ToAnalyticsConfigurationOutputWithContext(ctx context.Context) AnalyticsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConfigurationOutput)
}

func (i *AnalyticsConfiguration) ToAnalyticsConfigurationPtrOutput() AnalyticsConfigurationPtrOutput {
	return i.ToAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *AnalyticsConfiguration) ToAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) AnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConfigurationPtrOutput)
}

type AnalyticsConfigurationPtrInput interface {
	pulumi.Input

	ToAnalyticsConfigurationPtrOutput() AnalyticsConfigurationPtrOutput
	ToAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) AnalyticsConfigurationPtrOutput
}

type analyticsConfigurationPtrType AnalyticsConfigurationArgs

func (*analyticsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsConfiguration)(nil))
}

func (i *analyticsConfigurationPtrType) ToAnalyticsConfigurationPtrOutput() AnalyticsConfigurationPtrOutput {
	return i.ToAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *analyticsConfigurationPtrType) ToAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) AnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConfigurationOutput).ToAnalyticsConfigurationPtrOutput()
}

// AnalyticsConfigurationArrayInput is an input type that accepts AnalyticsConfigurationArray and AnalyticsConfigurationArrayOutput values.
// You can construct a concrete instance of `AnalyticsConfigurationArrayInput` via:
//
//          AnalyticsConfigurationArray{ AnalyticsConfigurationArgs{...} }
type AnalyticsConfigurationArrayInput interface {
	pulumi.Input

	ToAnalyticsConfigurationArrayOutput() AnalyticsConfigurationArrayOutput
	ToAnalyticsConfigurationArrayOutputWithContext(context.Context) AnalyticsConfigurationArrayOutput
}

type AnalyticsConfigurationArray []AnalyticsConfigurationInput

func (AnalyticsConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsConfiguration)(nil))
}

func (i AnalyticsConfigurationArray) ToAnalyticsConfigurationArrayOutput() AnalyticsConfigurationArrayOutput {
	return i.ToAnalyticsConfigurationArrayOutputWithContext(context.Background())
}

func (i AnalyticsConfigurationArray) ToAnalyticsConfigurationArrayOutputWithContext(ctx context.Context) AnalyticsConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConfigurationArrayOutput)
}

// AnalyticsConfigurationMapInput is an input type that accepts AnalyticsConfigurationMap and AnalyticsConfigurationMapOutput values.
// You can construct a concrete instance of `AnalyticsConfigurationMapInput` via:
//
//          AnalyticsConfigurationMap{ "key": AnalyticsConfigurationArgs{...} }
type AnalyticsConfigurationMapInput interface {
	pulumi.Input

	ToAnalyticsConfigurationMapOutput() AnalyticsConfigurationMapOutput
	ToAnalyticsConfigurationMapOutputWithContext(context.Context) AnalyticsConfigurationMapOutput
}

type AnalyticsConfigurationMap map[string]AnalyticsConfigurationInput

func (AnalyticsConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnalyticsConfiguration)(nil))
}

func (i AnalyticsConfigurationMap) ToAnalyticsConfigurationMapOutput() AnalyticsConfigurationMapOutput {
	return i.ToAnalyticsConfigurationMapOutputWithContext(context.Background())
}

func (i AnalyticsConfigurationMap) ToAnalyticsConfigurationMapOutputWithContext(ctx context.Context) AnalyticsConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticsConfigurationMapOutput)
}

type AnalyticsConfigurationOutput struct {
	*pulumi.OutputState
}

func (AnalyticsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticsConfiguration)(nil))
}

func (o AnalyticsConfigurationOutput) ToAnalyticsConfigurationOutput() AnalyticsConfigurationOutput {
	return o
}

func (o AnalyticsConfigurationOutput) ToAnalyticsConfigurationOutputWithContext(ctx context.Context) AnalyticsConfigurationOutput {
	return o
}

func (o AnalyticsConfigurationOutput) ToAnalyticsConfigurationPtrOutput() AnalyticsConfigurationPtrOutput {
	return o.ToAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (o AnalyticsConfigurationOutput) ToAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) AnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v AnalyticsConfiguration) *AnalyticsConfiguration {
		return &v
	}).(AnalyticsConfigurationPtrOutput)
}

type AnalyticsConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (AnalyticsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticsConfiguration)(nil))
}

func (o AnalyticsConfigurationPtrOutput) ToAnalyticsConfigurationPtrOutput() AnalyticsConfigurationPtrOutput {
	return o
}

func (o AnalyticsConfigurationPtrOutput) ToAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) AnalyticsConfigurationPtrOutput {
	return o
}

type AnalyticsConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AnalyticsConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AnalyticsConfiguration)(nil))
}

func (o AnalyticsConfigurationArrayOutput) ToAnalyticsConfigurationArrayOutput() AnalyticsConfigurationArrayOutput {
	return o
}

func (o AnalyticsConfigurationArrayOutput) ToAnalyticsConfigurationArrayOutputWithContext(ctx context.Context) AnalyticsConfigurationArrayOutput {
	return o
}

func (o AnalyticsConfigurationArrayOutput) Index(i pulumi.IntInput) AnalyticsConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AnalyticsConfiguration {
		return vs[0].([]AnalyticsConfiguration)[vs[1].(int)]
	}).(AnalyticsConfigurationOutput)
}

type AnalyticsConfigurationMapOutput struct{ *pulumi.OutputState }

func (AnalyticsConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AnalyticsConfiguration)(nil))
}

func (o AnalyticsConfigurationMapOutput) ToAnalyticsConfigurationMapOutput() AnalyticsConfigurationMapOutput {
	return o
}

func (o AnalyticsConfigurationMapOutput) ToAnalyticsConfigurationMapOutputWithContext(ctx context.Context) AnalyticsConfigurationMapOutput {
	return o
}

func (o AnalyticsConfigurationMapOutput) MapIndex(k pulumi.StringInput) AnalyticsConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AnalyticsConfiguration {
		return vs[0].(map[string]AnalyticsConfiguration)[vs[1].(string)]
	}).(AnalyticsConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(AnalyticsConfigurationOutput{})
	pulumi.RegisterOutputType(AnalyticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AnalyticsConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AnalyticsConfigurationMapOutput{})
}
