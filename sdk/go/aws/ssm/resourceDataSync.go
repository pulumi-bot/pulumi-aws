// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ssm

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a SSM resource data sync.
//
// {{% examples %}}
// {{% /examples %}}
// ## s3Destination
//
// `s3Destination` supports the following:
//
// * `bucketName` - (Required) Name of S3 bucket where the aggregated data is stored.
// * `region` - (Required) Region with the bucket targeted by the Resource Data Sync.
// * `kmsKeyArn` - (Optional) ARN of an encryption key for a destination in Amazon S3.
// * `prefix` - (Optional) Prefix for the bucket.
// * `syncFormat` - (Optional) A supported sync format. Only JsonSerDe is currently supported. Defaults to JsonSerDe.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_resource_data_sync.html.markdown.
type ResourceDataSync struct {
	pulumi.CustomResourceState

	// Name for the configuration.
	Name pulumi.StringOutput `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationOutput `pulumi:"s3Destination"`
}

// NewResourceDataSync registers a new resource with the given unique name, arguments, and options.
func NewResourceDataSync(ctx *pulumi.Context,
	name string, args *ResourceDataSyncArgs, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	if args == nil || args.S3Destination == nil {
		return nil, errors.New("missing required argument 'S3Destination'")
	}
	if args == nil {
		args = &ResourceDataSyncArgs{}
	}
	var resource ResourceDataSync
	err := ctx.RegisterResource("aws:ssm/resourceDataSync:ResourceDataSync", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDataSync gets an existing ResourceDataSync resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDataSync(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDataSyncState, opts ...pulumi.ResourceOption) (*ResourceDataSync, error) {
	var resource ResourceDataSync
	err := ctx.ReadResource("aws:ssm/resourceDataSync:ResourceDataSync", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDataSync resources.
type resourceDataSyncState struct {
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination *ResourceDataSyncS3Destination `pulumi:"s3Destination"`
}

type ResourceDataSyncState struct {
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationPtrInput
}

func (ResourceDataSyncState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncState)(nil)).Elem()
}

type resourceDataSyncArgs struct {
	// Name for the configuration.
	Name *string `pulumi:"name"`
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3Destination `pulumi:"s3Destination"`
}

// The set of arguments for constructing a ResourceDataSync resource.
type ResourceDataSyncArgs struct {
	// Name for the configuration.
	Name pulumi.StringPtrInput
	// Amazon S3 configuration details for the sync.
	S3Destination ResourceDataSyncS3DestinationInput
}

func (ResourceDataSyncArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDataSyncArgs)(nil)).Elem()
}

