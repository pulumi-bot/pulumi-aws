// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3control

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type BucketLifecycleConfiguration struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayOutput `pulumi:"rules"`
}

// NewBucketLifecycleConfiguration registers a new resource with the given unique name, arguments, and options.
func NewBucketLifecycleConfiguration(ctx *pulumi.Context,
	name string, args *BucketLifecycleConfigurationArgs, opts ...pulumi.ResourceOption) (*BucketLifecycleConfiguration, error) {
	if args == nil || args.Bucket == nil {
		return nil, errors.New("missing required argument 'Bucket'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	if args == nil {
		args = &BucketLifecycleConfigurationArgs{}
	}
	var resource BucketLifecycleConfiguration
	err := ctx.RegisterResource("aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketLifecycleConfiguration gets an existing BucketLifecycleConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketLifecycleConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketLifecycleConfigurationState, opts ...pulumi.ResourceOption) (*BucketLifecycleConfiguration, error) {
	var resource BucketLifecycleConfiguration
	err := ctx.ReadResource("aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketLifecycleConfiguration resources.
type bucketLifecycleConfigurationState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket *string `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules []BucketLifecycleConfigurationRule `pulumi:"rules"`
}

type BucketLifecycleConfigurationState struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringPtrInput
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayInput
}

func (BucketLifecycleConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationState)(nil)).Elem()
}

type bucketLifecycleConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket string `pulumi:"bucket"`
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules []BucketLifecycleConfigurationRule `pulumi:"rules"`
}

// The set of arguments for constructing a BucketLifecycleConfiguration resource.
type BucketLifecycleConfigurationArgs struct {
	// Amazon Resource Name (ARN) of the bucket.
	Bucket pulumi.StringInput
	// Configuration block(s) containing lifecycle rules for the bucket.
	Rules BucketLifecycleConfigurationRuleArrayInput
}

func (BucketLifecycleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketLifecycleConfigurationArgs)(nil)).Elem()
}

type BucketLifecycleConfigurationInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput
	ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput
}

func (BucketLifecycleConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfiguration)(nil)).Elem()
}

func (i BucketLifecycleConfiguration) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return i.ToBucketLifecycleConfigurationOutputWithContext(context.Background())
}

func (i BucketLifecycleConfiguration) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationOutput)
}

type BucketLifecycleConfigurationOutput struct {
	*pulumi.OutputState
}

func (BucketLifecycleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfigurationOutput)(nil)).Elem()
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return o
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketLifecycleConfigurationOutput{})
}
