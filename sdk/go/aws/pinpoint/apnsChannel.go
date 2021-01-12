// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint APNs Channel resource.
//
// > **Note:** All arguments, including certificates and tokens, will be stored in the raw state as plain-text.
//
// ## Import
//
// Pinpoint APNs Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/apnsChannel:ApnsChannel apns application-id
// ```
type ApnsChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrOutput `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrOutput `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrOutput `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrOutput `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewApnsChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsChannel(ctx *pulumi.Context,
	name string, args *ApnsChannelArgs, opts ...pulumi.ResourceOption) (*ApnsChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource ApnsChannel
	err := ctx.RegisterResource("aws:pinpoint/apnsChannel:ApnsChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsChannel gets an existing ApnsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsChannelState, opts ...pulumi.ResourceOption) (*ApnsChannel, error) {
	var resource ApnsChannel
	err := ctx.ReadResource("aws:pinpoint/apnsChannel:ApnsChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsChannel resources.
type apnsChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

type ApnsChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsChannelState)(nil)).Elem()
}

type apnsChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey *string `pulumi:"privateKey"`
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId *string `pulumi:"teamId"`
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey *string `pulumi:"tokenKey"`
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a ApnsChannel resource.
type ApnsChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs.
	// __NOTE__: Amazon Pinpoint uses this default for every APNs push notification that you send using the console.
	// You can override the default when you send a message programmatically using the Amazon Pinpoint API, the AWS CLI, or an AWS SDK.
	// If your default authentication type fails, Amazon Pinpoint doesn't attempt to use the other authentication type.
	DefaultAuthenticationMethod pulumi.StringPtrInput
	// Whether the channel is enabled or disabled. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Certificate Private Key file (ie. `.key` file).
	PrivateKey pulumi.StringPtrInput
	// The ID assigned to your Apple developer account team. This value is provided on the Membership page.
	TeamId pulumi.StringPtrInput
	// The `.p8` file that you download from your Apple developer account when you create an authentication key.
	TokenKey pulumi.StringPtrInput
	// The ID assigned to your signing key. To find this value, choose Certificates, IDs & Profiles, and choose your key in the Keys section.
	TokenKeyId pulumi.StringPtrInput
}

func (ApnsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsChannelArgs)(nil)).Elem()
}

type ApnsChannelInput interface {
	pulumi.Input

	ToApnsChannelOutput() ApnsChannelOutput
	ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput
}

func (*ApnsChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsChannel)(nil))
}

func (i *ApnsChannel) ToApnsChannelOutput() ApnsChannelOutput {
	return i.ToApnsChannelOutputWithContext(context.Background())
}

func (i *ApnsChannel) ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelOutput)
}

func (i *ApnsChannel) ToApnsChannelPtrOutput() ApnsChannelPtrOutput {
	return i.ToApnsChannelPtrOutputWithContext(context.Background())
}

func (i *ApnsChannel) ToApnsChannelPtrOutputWithContext(ctx context.Context) ApnsChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelPtrOutput)
}

type ApnsChannelPtrInput interface {
	pulumi.Input

	ToApnsChannelPtrOutput() ApnsChannelPtrOutput
	ToApnsChannelPtrOutputWithContext(ctx context.Context) ApnsChannelPtrOutput
}

type apnsChannelPtrType ApnsChannelArgs

func (*apnsChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsChannel)(nil))
}

func (i *apnsChannelPtrType) ToApnsChannelPtrOutput() ApnsChannelPtrOutput {
	return i.ToApnsChannelPtrOutputWithContext(context.Background())
}

func (i *apnsChannelPtrType) ToApnsChannelPtrOutputWithContext(ctx context.Context) ApnsChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelOutput).ToApnsChannelPtrOutput()
}

// ApnsChannelArrayInput is an input type that accepts ApnsChannelArray and ApnsChannelArrayOutput values.
// You can construct a concrete instance of `ApnsChannelArrayInput` via:
//
//          ApnsChannelArray{ ApnsChannelArgs{...} }
type ApnsChannelArrayInput interface {
	pulumi.Input

	ToApnsChannelArrayOutput() ApnsChannelArrayOutput
	ToApnsChannelArrayOutputWithContext(context.Context) ApnsChannelArrayOutput
}

type ApnsChannelArray []ApnsChannelInput

func (ApnsChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApnsChannel)(nil))
}

func (i ApnsChannelArray) ToApnsChannelArrayOutput() ApnsChannelArrayOutput {
	return i.ToApnsChannelArrayOutputWithContext(context.Background())
}

func (i ApnsChannelArray) ToApnsChannelArrayOutputWithContext(ctx context.Context) ApnsChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelArrayOutput)
}

// ApnsChannelMapInput is an input type that accepts ApnsChannelMap and ApnsChannelMapOutput values.
// You can construct a concrete instance of `ApnsChannelMapInput` via:
//
//          ApnsChannelMap{ "key": ApnsChannelArgs{...} }
type ApnsChannelMapInput interface {
	pulumi.Input

	ToApnsChannelMapOutput() ApnsChannelMapOutput
	ToApnsChannelMapOutputWithContext(context.Context) ApnsChannelMapOutput
}

type ApnsChannelMap map[string]ApnsChannelInput

func (ApnsChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApnsChannel)(nil))
}

func (i ApnsChannelMap) ToApnsChannelMapOutput() ApnsChannelMapOutput {
	return i.ToApnsChannelMapOutputWithContext(context.Background())
}

func (i ApnsChannelMap) ToApnsChannelMapOutputWithContext(ctx context.Context) ApnsChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsChannelMapOutput)
}

type ApnsChannelOutput struct {
	*pulumi.OutputState
}

func (ApnsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsChannel)(nil))
}

func (o ApnsChannelOutput) ToApnsChannelOutput() ApnsChannelOutput {
	return o
}

func (o ApnsChannelOutput) ToApnsChannelOutputWithContext(ctx context.Context) ApnsChannelOutput {
	return o
}

func (o ApnsChannelOutput) ToApnsChannelPtrOutput() ApnsChannelPtrOutput {
	return o.ToApnsChannelPtrOutputWithContext(context.Background())
}

func (o ApnsChannelOutput) ToApnsChannelPtrOutputWithContext(ctx context.Context) ApnsChannelPtrOutput {
	return o.ApplyT(func(v ApnsChannel) *ApnsChannel {
		return &v
	}).(ApnsChannelPtrOutput)
}

type ApnsChannelPtrOutput struct {
	*pulumi.OutputState
}

func (ApnsChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsChannel)(nil))
}

func (o ApnsChannelPtrOutput) ToApnsChannelPtrOutput() ApnsChannelPtrOutput {
	return o
}

func (o ApnsChannelPtrOutput) ToApnsChannelPtrOutputWithContext(ctx context.Context) ApnsChannelPtrOutput {
	return o
}

type ApnsChannelArrayOutput struct{ *pulumi.OutputState }

func (ApnsChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApnsChannel)(nil))
}

func (o ApnsChannelArrayOutput) ToApnsChannelArrayOutput() ApnsChannelArrayOutput {
	return o
}

func (o ApnsChannelArrayOutput) ToApnsChannelArrayOutputWithContext(ctx context.Context) ApnsChannelArrayOutput {
	return o
}

func (o ApnsChannelArrayOutput) Index(i pulumi.IntInput) ApnsChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApnsChannel {
		return vs[0].([]ApnsChannel)[vs[1].(int)]
	}).(ApnsChannelOutput)
}

type ApnsChannelMapOutput struct{ *pulumi.OutputState }

func (ApnsChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApnsChannel)(nil))
}

func (o ApnsChannelMapOutput) ToApnsChannelMapOutput() ApnsChannelMapOutput {
	return o
}

func (o ApnsChannelMapOutput) ToApnsChannelMapOutputWithContext(ctx context.Context) ApnsChannelMapOutput {
	return o
}

func (o ApnsChannelMapOutput) MapIndex(k pulumi.StringInput) ApnsChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApnsChannel {
		return vs[0].(map[string]ApnsChannel)[vs[1].(string)]
	}).(ApnsChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(ApnsChannelOutput{})
	pulumi.RegisterOutputType(ApnsChannelPtrOutput{})
	pulumi.RegisterOutputType(ApnsChannelArrayOutput{})
	pulumi.RegisterOutputType(ApnsChannelMapOutput{})
}
