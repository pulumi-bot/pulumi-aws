// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint Baidu Channel resource.
//
// > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/pinpoint"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		app, err := pinpoint.NewApp(ctx, "app", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = pinpoint.NewBaiduChannel(ctx, "channel", &pinpoint.BaiduChannelArgs{
// 			ApplicationId: app.ApplicationId,
// 			ApiKey:        pulumi.String(""),
// 			SecretKey:     pulumi.String(""),
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
// Pinpoint Baidu Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/baiduChannel:BaiduChannel channel application-id
// ```
type BaiduChannel struct {
	pulumi.CustomResourceState

	// Platform credential API key from Baidu.
	ApiKey pulumi.StringOutput `pulumi:"apiKey"`
	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
}

// NewBaiduChannel registers a new resource with the given unique name, arguments, and options.
func NewBaiduChannel(ctx *pulumi.Context,
	name string, args *BaiduChannelArgs, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiKey == nil {
		return nil, errors.New("invalid value for required argument 'ApiKey'")
	}
	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	if args.SecretKey == nil {
		return nil, errors.New("invalid value for required argument 'SecretKey'")
	}
	var resource BaiduChannel
	err := ctx.RegisterResource("aws:pinpoint/baiduChannel:BaiduChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBaiduChannel gets an existing BaiduChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBaiduChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BaiduChannelState, opts ...pulumi.ResourceOption) (*BaiduChannel, error) {
	var resource BaiduChannel
	err := ctx.ReadResource("aws:pinpoint/baiduChannel:BaiduChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BaiduChannel resources.
type baiduChannelState struct {
	// Platform credential API key from Baidu.
	ApiKey *string `pulumi:"apiKey"`
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Platform credential Secret key from Baidu.
	SecretKey *string `pulumi:"secretKey"`
}

type BaiduChannelState struct {
	// Platform credential API key from Baidu.
	ApiKey pulumi.StringPtrInput
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringPtrInput
}

func (BaiduChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*baiduChannelState)(nil)).Elem()
}

type baiduChannelArgs struct {
	// Platform credential API key from Baidu.
	ApiKey string `pulumi:"apiKey"`
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Platform credential Secret key from Baidu.
	SecretKey string `pulumi:"secretKey"`
}

// The set of arguments for constructing a BaiduChannel resource.
type BaiduChannelArgs struct {
	// Platform credential API key from Baidu.
	ApiKey pulumi.StringInput
	// The application ID.
	ApplicationId pulumi.StringInput
	// Specifies whether to enable the channel. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Platform credential Secret key from Baidu.
	SecretKey pulumi.StringInput
}

func (BaiduChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*baiduChannelArgs)(nil)).Elem()
}

type BaiduChannelInput interface {
	pulumi.Input

	ToBaiduChannelOutput() BaiduChannelOutput
	ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput
}

func (*BaiduChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduChannel)(nil))
}

func (i *BaiduChannel) ToBaiduChannelOutput() BaiduChannelOutput {
	return i.ToBaiduChannelOutputWithContext(context.Background())
}

func (i *BaiduChannel) ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BaiduChannelOutput)
}

type BaiduChannelOutput struct {
	*pulumi.OutputState
}

func (BaiduChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BaiduChannel)(nil))
}

func (o BaiduChannelOutput) ToBaiduChannelOutput() BaiduChannelOutput {
	return o
}

func (o BaiduChannelOutput) ToBaiduChannelOutputWithContext(ctx context.Context) BaiduChannelOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BaiduChannelOutput{})
}
