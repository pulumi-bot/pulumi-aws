// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Elemental MediaPackage Channel.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/mediapackage"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := mediapackage.NewChannel(ctx, "kittens", &mediapackage.ChannelArgs{
// 			ChannelId:   pulumi.String("kitten-channel"),
// 			Description: pulumi.String("A channel dedicated to amusing videos of kittens."),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Channel struct {
	pulumi.CustomResourceState

	// The ARN of the channel
	Arn pulumi.StringOutput `pulumi:"arn"`
	// A unique identifier describing the channel
	ChannelId pulumi.StringOutput `pulumi:"channelId"`
	// A description of the channel
	Description pulumi.StringOutput `pulumi:"description"`
	// A single item list of HLS ingest information
	HlsIngests ChannelHlsIngestArrayOutput `pulumi:"hlsIngests"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ChannelId == nil {
		return nil, errors.New("invalid value for required argument 'ChannelId'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource Channel
	err := ctx.RegisterResource("aws:mediapackage/channel:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("aws:mediapackage/channel:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
	// The ARN of the channel
	Arn *string `pulumi:"arn"`
	// A unique identifier describing the channel
	ChannelId *string `pulumi:"channelId"`
	// A description of the channel
	Description *string `pulumi:"description"`
	// A single item list of HLS ingest information
	HlsIngests []ChannelHlsIngest `pulumi:"hlsIngests"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ChannelState struct {
	// The ARN of the channel
	Arn pulumi.StringPtrInput
	// A unique identifier describing the channel
	ChannelId pulumi.StringPtrInput
	// A description of the channel
	Description pulumi.StringPtrInput
	// A single item list of HLS ingest information
	HlsIngests ChannelHlsIngestArrayInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	// A unique identifier describing the channel
	ChannelId string `pulumi:"channelId"`
	// A description of the channel
	Description *string `pulumi:"description"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	// A unique identifier describing the channel
	ChannelId pulumi.StringInput
	// A description of the channel
	Description pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}
