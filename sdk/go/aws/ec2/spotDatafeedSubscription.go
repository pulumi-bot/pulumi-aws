// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **Note:** There is only a single subscription allowed per account.
//
// To help you understand the charges for your Spot instances, Amazon EC2 provides a data feed that describes your Spot instance usage and pricing.
// This data feed is sent to an Amazon S3 bucket that you specify when you subscribe to the data feed.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		defaultBucket, err := s3.NewBucket(ctx, "defaultBucket", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSpotDatafeedSubscription(ctx, "defaultSpotDatafeedSubscription", &ec2.SpotDatafeedSubscriptionArgs{
// 			Bucket: defaultBucket.Bucket,
// 			Prefix: pulumi.String("my_subdirectory"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SpotDatafeedSubscription struct {
	pulumi.CustomResourceState

	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
}

// NewSpotDatafeedSubscription registers a new resource with the given unique name, arguments, and options.
func NewSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, args *SpotDatafeedSubscriptionArgs, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	var resource SpotDatafeedSubscription
	err := ctx.RegisterResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpotDatafeedSubscription gets an existing SpotDatafeedSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpotDatafeedSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpotDatafeedSubscriptionState, opts ...pulumi.ResourceOption) (*SpotDatafeedSubscription, error) {
	var resource SpotDatafeedSubscription
	err := ctx.ReadResource("aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpotDatafeedSubscription resources.
type spotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket *string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

type SpotDatafeedSubscriptionState struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringPtrInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionState)(nil)).Elem()
}

type spotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket string `pulumi:"bucket"`
	// Path of folder inside bucket to place spot pricing data.
	Prefix *string `pulumi:"prefix"`
}

// The set of arguments for constructing a SpotDatafeedSubscription resource.
type SpotDatafeedSubscriptionArgs struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket pulumi.StringInput
	// Path of folder inside bucket to place spot pricing data.
	Prefix pulumi.StringPtrInput
}

func (SpotDatafeedSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spotDatafeedSubscriptionArgs)(nil)).Elem()
}

type SpotDatafeedSubscriptionInput interface {
	pulumi.Input

	ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput
	ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput
}

func (SpotDatafeedSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotDatafeedSubscription)(nil)).Elem()
}

func (i SpotDatafeedSubscription) ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput {
	return i.ToSpotDatafeedSubscriptionOutputWithContext(context.Background())
}

func (i SpotDatafeedSubscription) ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpotDatafeedSubscriptionOutput)
}

type SpotDatafeedSubscriptionOutput struct {
	*pulumi.OutputState
}

func (SpotDatafeedSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpotDatafeedSubscriptionOutput)(nil)).Elem()
}

func (o SpotDatafeedSubscriptionOutput) ToSpotDatafeedSubscriptionOutput() SpotDatafeedSubscriptionOutput {
	return o
}

func (o SpotDatafeedSubscriptionOutput) ToSpotDatafeedSubscriptionOutputWithContext(ctx context.Context) SpotDatafeedSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SpotDatafeedSubscriptionOutput{})
}
