// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
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
