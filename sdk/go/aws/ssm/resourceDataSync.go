// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a SSM resource data sync.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		hogeBucket, err := s3.NewBucket(ctx, "hogeBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPolicy(ctx, "hogeBucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: hogeBucket.Bucket,
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "    \"Version\": \"2012-10-17\",\n", "    \"Statement\": [\n", "        {\n", "            \"Sid\": \"SSMBucketPermissionsCheck\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "                \"Service\": \"ssm.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:GetBucketAcl\",\n", "            \"Resource\": \"arn:aws:s3:::tf-test-bucket-1234\"\n", "        },\n", "        {\n", "            \"Sid\": \" SSMBucketDelivery\",\n", "            \"Effect\": \"Allow\",\n", "            \"Principal\": {\n", "                \"Service\": \"ssm.amazonaws.com\"\n", "            },\n", "            \"Action\": \"s3:PutObject\",\n", "            \"Resource\": [\"arn:aws:s3:::tf-test-bucket-1234/*\"],\n", "            \"Condition\": {\n", "                \"StringEquals\": {\n", "                    \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                }\n", "            }\n", "        }\n", "    ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewResourceDataSync(ctx, "foo", &ssm.ResourceDataSyncArgs{
// 			S3Destination: &ssm.ResourceDataSyncS3DestinationArgs{
// 				BucketName: hogeBucket.Bucket,
// 				Region:     hogeBucket.Region,
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
// SSM resource data sync can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:ssm/resourceDataSync:ResourceDataSync example example-name
// ```
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
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3Destination == nil {
		return nil, errors.New("invalid value for required argument 'S3Destination'")
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

type ResourceDataSyncInput interface {
	pulumi.Input

	ToResourceDataSyncOutput() ResourceDataSyncOutput
	ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput
}

func (*ResourceDataSync) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDataSync)(nil))
}

func (i *ResourceDataSync) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return i.ToResourceDataSyncOutputWithContext(context.Background())
}

func (i *ResourceDataSync) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncOutput)
}

func (i *ResourceDataSync) ToResourceDataSyncPtrOutput() ResourceDataSyncPtrOutput {
	return i.ToResourceDataSyncPtrOutputWithContext(context.Background())
}

func (i *ResourceDataSync) ToResourceDataSyncPtrOutputWithContext(ctx context.Context) ResourceDataSyncPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncPtrOutput)
}

type ResourceDataSyncPtrInput interface {
	pulumi.Input

	ToResourceDataSyncPtrOutput() ResourceDataSyncPtrOutput
	ToResourceDataSyncPtrOutputWithContext(ctx context.Context) ResourceDataSyncPtrOutput
}

type resourceDataSyncPtrType ResourceDataSyncArgs

func (*resourceDataSyncPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDataSync)(nil))
}

func (i *resourceDataSyncPtrType) ToResourceDataSyncPtrOutput() ResourceDataSyncPtrOutput {
	return i.ToResourceDataSyncPtrOutputWithContext(context.Background())
}

func (i *resourceDataSyncPtrType) ToResourceDataSyncPtrOutputWithContext(ctx context.Context) ResourceDataSyncPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncOutput).ToResourceDataSyncPtrOutput()
}

// ResourceDataSyncArrayInput is an input type that accepts ResourceDataSyncArray and ResourceDataSyncArrayOutput values.
// You can construct a concrete instance of `ResourceDataSyncArrayInput` via:
//
//          ResourceDataSyncArray{ ResourceDataSyncArgs{...} }
type ResourceDataSyncArrayInput interface {
	pulumi.Input

	ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput
	ToResourceDataSyncArrayOutputWithContext(context.Context) ResourceDataSyncArrayOutput
}

type ResourceDataSyncArray []ResourceDataSyncInput

func (ResourceDataSyncArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceDataSync)(nil))
}

func (i ResourceDataSyncArray) ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput {
	return i.ToResourceDataSyncArrayOutputWithContext(context.Background())
}

func (i ResourceDataSyncArray) ToResourceDataSyncArrayOutputWithContext(ctx context.Context) ResourceDataSyncArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncArrayOutput)
}

// ResourceDataSyncMapInput is an input type that accepts ResourceDataSyncMap and ResourceDataSyncMapOutput values.
// You can construct a concrete instance of `ResourceDataSyncMapInput` via:
//
//          ResourceDataSyncMap{ "key": ResourceDataSyncArgs{...} }
type ResourceDataSyncMapInput interface {
	pulumi.Input

	ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput
	ToResourceDataSyncMapOutputWithContext(context.Context) ResourceDataSyncMapOutput
}

type ResourceDataSyncMap map[string]ResourceDataSyncInput

func (ResourceDataSyncMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceDataSync)(nil))
}

func (i ResourceDataSyncMap) ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput {
	return i.ToResourceDataSyncMapOutputWithContext(context.Background())
}

func (i ResourceDataSyncMap) ToResourceDataSyncMapOutputWithContext(ctx context.Context) ResourceDataSyncMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDataSyncMapOutput)
}

type ResourceDataSyncOutput struct {
	*pulumi.OutputState
}

func (ResourceDataSyncOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDataSync)(nil))
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutput() ResourceDataSyncOutput {
	return o
}

func (o ResourceDataSyncOutput) ToResourceDataSyncOutputWithContext(ctx context.Context) ResourceDataSyncOutput {
	return o
}

func (o ResourceDataSyncOutput) ToResourceDataSyncPtrOutput() ResourceDataSyncPtrOutput {
	return o.ToResourceDataSyncPtrOutputWithContext(context.Background())
}

func (o ResourceDataSyncOutput) ToResourceDataSyncPtrOutputWithContext(ctx context.Context) ResourceDataSyncPtrOutput {
	return o.ApplyT(func(v ResourceDataSync) *ResourceDataSync {
		return &v
	}).(ResourceDataSyncPtrOutput)
}

type ResourceDataSyncPtrOutput struct {
	*pulumi.OutputState
}

func (ResourceDataSyncPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceDataSync)(nil))
}

func (o ResourceDataSyncPtrOutput) ToResourceDataSyncPtrOutput() ResourceDataSyncPtrOutput {
	return o
}

func (o ResourceDataSyncPtrOutput) ToResourceDataSyncPtrOutputWithContext(ctx context.Context) ResourceDataSyncPtrOutput {
	return o
}

type ResourceDataSyncArrayOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceDataSync)(nil))
}

func (o ResourceDataSyncArrayOutput) ToResourceDataSyncArrayOutput() ResourceDataSyncArrayOutput {
	return o
}

func (o ResourceDataSyncArrayOutput) ToResourceDataSyncArrayOutputWithContext(ctx context.Context) ResourceDataSyncArrayOutput {
	return o
}

func (o ResourceDataSyncArrayOutput) Index(i pulumi.IntInput) ResourceDataSyncOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceDataSync {
		return vs[0].([]ResourceDataSync)[vs[1].(int)]
	}).(ResourceDataSyncOutput)
}

type ResourceDataSyncMapOutput struct{ *pulumi.OutputState }

func (ResourceDataSyncMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ResourceDataSync)(nil))
}

func (o ResourceDataSyncMapOutput) ToResourceDataSyncMapOutput() ResourceDataSyncMapOutput {
	return o
}

func (o ResourceDataSyncMapOutput) ToResourceDataSyncMapOutputWithContext(ctx context.Context) ResourceDataSyncMapOutput {
	return o
}

func (o ResourceDataSyncMapOutput) MapIndex(k pulumi.StringInput) ResourceDataSyncOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ResourceDataSync {
		return vs[0].(map[string]ResourceDataSync)[vs[1].(string)]
	}).(ResourceDataSyncOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceDataSyncOutput{})
	pulumi.RegisterOutputType(ResourceDataSyncPtrOutput{})
	pulumi.RegisterOutputType(ResourceDataSyncArrayOutput{})
	pulumi.RegisterOutputType(ResourceDataSyncMapOutput{})
}
