// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// S3 Bucket Ownership Controls can be imported using S3 Bucket name, e.g.
//
// ```sh
//  $ pulumi import aws:s3/bucketOwnershipControls:BucketOwnershipControls example my-bucket
// ```
type BucketOwnershipControls struct {
	pulumi.CustomResourceState

	// The name of the bucket that you want to associate this access point with.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule BucketOwnershipControlsRuleOutput `pulumi:"rule"`
}

// NewBucketOwnershipControls registers a new resource with the given unique name, arguments, and options.
func NewBucketOwnershipControls(ctx *pulumi.Context,
	name string, args *BucketOwnershipControlsArgs, opts ...pulumi.ResourceOption) (*BucketOwnershipControls, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Rule == nil {
		return nil, errors.New("invalid value for required argument 'Rule'")
	}
	var resource BucketOwnershipControls
	err := ctx.RegisterResource("aws:s3/bucketOwnershipControls:BucketOwnershipControls", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketOwnershipControls gets an existing BucketOwnershipControls resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketOwnershipControls(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketOwnershipControlsState, opts ...pulumi.ResourceOption) (*BucketOwnershipControls, error) {
	var resource BucketOwnershipControls
	err := ctx.ReadResource("aws:s3/bucketOwnershipControls:BucketOwnershipControls", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketOwnershipControls resources.
type bucketOwnershipControlsState struct {
	// The name of the bucket that you want to associate this access point with.
	Bucket *string `pulumi:"bucket"`
	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule *BucketOwnershipControlsRule `pulumi:"rule"`
}

type BucketOwnershipControlsState struct {
	// The name of the bucket that you want to associate this access point with.
	Bucket pulumi.StringPtrInput
	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule BucketOwnershipControlsRulePtrInput
}

func (BucketOwnershipControlsState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketOwnershipControlsState)(nil)).Elem()
}

type bucketOwnershipControlsArgs struct {
	// The name of the bucket that you want to associate this access point with.
	Bucket string `pulumi:"bucket"`
	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule BucketOwnershipControlsRule `pulumi:"rule"`
}

// The set of arguments for constructing a BucketOwnershipControls resource.
type BucketOwnershipControlsArgs struct {
	// The name of the bucket that you want to associate this access point with.
	Bucket pulumi.StringInput
	// Configuration block(s) with Ownership Controls rules. Detailed below.
	Rule BucketOwnershipControlsRuleInput
}

func (BucketOwnershipControlsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketOwnershipControlsArgs)(nil)).Elem()
}

type BucketOwnershipControlsInput interface {
	pulumi.Input

	ToBucketOwnershipControlsOutput() BucketOwnershipControlsOutput
	ToBucketOwnershipControlsOutputWithContext(ctx context.Context) BucketOwnershipControlsOutput
}

type BucketOwnershipControlsPtrInput interface {
	pulumi.Input

	ToBucketOwnershipControlsPtrOutput() BucketOwnershipControlsPtrOutput
	ToBucketOwnershipControlsPtrOutputWithContext(ctx context.Context) BucketOwnershipControlsPtrOutput
}

func (BucketOwnershipControls) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketOwnershipControls)(nil)).Elem()
}

func (i BucketOwnershipControls) ToBucketOwnershipControlsOutput() BucketOwnershipControlsOutput {
	return i.ToBucketOwnershipControlsOutputWithContext(context.Background())
}

func (i BucketOwnershipControls) ToBucketOwnershipControlsOutputWithContext(ctx context.Context) BucketOwnershipControlsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOwnershipControlsOutput)
}

func (i BucketOwnershipControls) ToBucketOwnershipControlsPtrOutput() BucketOwnershipControlsPtrOutput {
	return i.ToBucketOwnershipControlsPtrOutputWithContext(context.Background())
}

func (i BucketOwnershipControls) ToBucketOwnershipControlsPtrOutputWithContext(ctx context.Context) BucketOwnershipControlsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOwnershipControlsPtrOutput)
}

type BucketOwnershipControlsOutput struct {
	*pulumi.OutputState
}

func (BucketOwnershipControlsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketOwnershipControlsOutput)(nil)).Elem()
}

func (o BucketOwnershipControlsOutput) ToBucketOwnershipControlsOutput() BucketOwnershipControlsOutput {
	return o
}

func (o BucketOwnershipControlsOutput) ToBucketOwnershipControlsOutputWithContext(ctx context.Context) BucketOwnershipControlsOutput {
	return o
}

type BucketOwnershipControlsPtrOutput struct {
	*pulumi.OutputState
}

func (BucketOwnershipControlsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketOwnershipControls)(nil)).Elem()
}

func (o BucketOwnershipControlsPtrOutput) ToBucketOwnershipControlsPtrOutput() BucketOwnershipControlsPtrOutput {
	return o
}

func (o BucketOwnershipControlsPtrOutput) ToBucketOwnershipControlsPtrOutputWithContext(ctx context.Context) BucketOwnershipControlsPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketOwnershipControlsOutput{})
	pulumi.RegisterOutputType(BucketOwnershipControlsPtrOutput{})
}
