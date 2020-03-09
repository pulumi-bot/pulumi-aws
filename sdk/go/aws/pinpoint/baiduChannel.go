// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package pinpoint

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a Pinpoint Baidu Channel resource.
// 
// > **Note:** All arguments including the Api Key and Secret Key will be stored in the raw state as plain-text.
// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_baidu_channel.markdown.
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
	if args == nil || args.ApiKey == nil {
		return nil, errors.New("missing required argument 'ApiKey'")
	}
	if args == nil || args.ApplicationId == nil {
		return nil, errors.New("missing required argument 'ApplicationId'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	if args == nil {
		args = &BaiduChannelArgs{}
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

