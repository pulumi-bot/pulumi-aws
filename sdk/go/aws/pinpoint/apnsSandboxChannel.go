// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Pinpoint APNs Sandbox Channel resource.
//
// > **Note:** All arguments, including certificates and tokens, will be stored in the raw state as plain-text.
//
// ## Import
//
// Pinpoint APNs Sandbox Channel can be imported using the `application-id`, e.g.
//
// ```sh
//  $ pulumi import aws:pinpoint/apnsSandboxChannel:ApnsSandboxChannel apns_sandbox application-id
// ```
type ApnsSandboxChannel struct {
	pulumi.CustomResourceState

	// The application ID.
	ApplicationId pulumi.StringOutput `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrOutput `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrOutput `pulumi:"certificate"`
	// The default authentication method used for APNs Sandbox.
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

// NewApnsSandboxChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsSandboxChannel(ctx *pulumi.Context,
	name string, args *ApnsSandboxChannelArgs, opts ...pulumi.ResourceOption) (*ApnsSandboxChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	var resource ApnsSandboxChannel
	err := ctx.RegisterResource("aws:pinpoint/apnsSandboxChannel:ApnsSandboxChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsSandboxChannel gets an existing ApnsSandboxChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsSandboxChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsSandboxChannelState, opts ...pulumi.ResourceOption) (*ApnsSandboxChannel, error) {
	var resource ApnsSandboxChannel
	err := ctx.ReadResource("aws:pinpoint/apnsSandboxChannel:ApnsSandboxChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsSandboxChannel resources.
type apnsSandboxChannelState struct {
	// The application ID.
	ApplicationId *string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs Sandbox.
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

type ApnsSandboxChannelState struct {
	// The application ID.
	ApplicationId pulumi.StringPtrInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs Sandbox.
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

func (ApnsSandboxChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsSandboxChannelState)(nil)).Elem()
}

type apnsSandboxChannelArgs struct {
	// The application ID.
	ApplicationId string `pulumi:"applicationId"`
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId *string `pulumi:"bundleId"`
	// The pem encoded TLS Certificate from Apple.
	Certificate *string `pulumi:"certificate"`
	// The default authentication method used for APNs Sandbox.
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

// The set of arguments for constructing a ApnsSandboxChannel resource.
type ApnsSandboxChannelArgs struct {
	// The application ID.
	ApplicationId pulumi.StringInput
	// The ID assigned to your iOS app. To find this value, choose Certificates, IDs & Profiles, choose App IDs in the Identifiers section, and choose your app.
	BundleId pulumi.StringPtrInput
	// The pem encoded TLS Certificate from Apple.
	Certificate pulumi.StringPtrInput
	// The default authentication method used for APNs Sandbox.
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

func (ApnsSandboxChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsSandboxChannelArgs)(nil)).Elem()
}

type ApnsSandboxChannelInput interface {
	pulumi.Input

	ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput
	ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput
}

func (*ApnsSandboxChannel) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsSandboxChannel)(nil))
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput {
	return i.ToApnsSandboxChannelOutputWithContext(context.Background())
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelOutput)
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelPtrOutput() ApnsSandboxChannelPtrOutput {
	return i.ToApnsSandboxChannelPtrOutputWithContext(context.Background())
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelPtrOutputWithContext(ctx context.Context) ApnsSandboxChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelPtrOutput)
}

type ApnsSandboxChannelPtrInput interface {
	pulumi.Input

	ToApnsSandboxChannelPtrOutput() ApnsSandboxChannelPtrOutput
	ToApnsSandboxChannelPtrOutputWithContext(ctx context.Context) ApnsSandboxChannelPtrOutput
}

type apnsSandboxChannelPtrType ApnsSandboxChannelArgs

func (*apnsSandboxChannelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsSandboxChannel)(nil))
}

func (i *apnsSandboxChannelPtrType) ToApnsSandboxChannelPtrOutput() ApnsSandboxChannelPtrOutput {
	return i.ToApnsSandboxChannelPtrOutputWithContext(context.Background())
}

func (i *apnsSandboxChannelPtrType) ToApnsSandboxChannelPtrOutputWithContext(ctx context.Context) ApnsSandboxChannelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelOutput).ToApnsSandboxChannelPtrOutput()
}

// ApnsSandboxChannelArrayInput is an input type that accepts ApnsSandboxChannelArray and ApnsSandboxChannelArrayOutput values.
// You can construct a concrete instance of `ApnsSandboxChannelArrayInput` via:
//
//          ApnsSandboxChannelArray{ ApnsSandboxChannelArgs{...} }
type ApnsSandboxChannelArrayInput interface {
	pulumi.Input

	ToApnsSandboxChannelArrayOutput() ApnsSandboxChannelArrayOutput
	ToApnsSandboxChannelArrayOutputWithContext(context.Context) ApnsSandboxChannelArrayOutput
}

type ApnsSandboxChannelArray []ApnsSandboxChannelInput

func (ApnsSandboxChannelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApnsSandboxChannel)(nil))
}

func (i ApnsSandboxChannelArray) ToApnsSandboxChannelArrayOutput() ApnsSandboxChannelArrayOutput {
	return i.ToApnsSandboxChannelArrayOutputWithContext(context.Background())
}

func (i ApnsSandboxChannelArray) ToApnsSandboxChannelArrayOutputWithContext(ctx context.Context) ApnsSandboxChannelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelArrayOutput)
}

// ApnsSandboxChannelMapInput is an input type that accepts ApnsSandboxChannelMap and ApnsSandboxChannelMapOutput values.
// You can construct a concrete instance of `ApnsSandboxChannelMapInput` via:
//
//          ApnsSandboxChannelMap{ "key": ApnsSandboxChannelArgs{...} }
type ApnsSandboxChannelMapInput interface {
	pulumi.Input

	ToApnsSandboxChannelMapOutput() ApnsSandboxChannelMapOutput
	ToApnsSandboxChannelMapOutputWithContext(context.Context) ApnsSandboxChannelMapOutput
}

type ApnsSandboxChannelMap map[string]ApnsSandboxChannelInput

func (ApnsSandboxChannelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApnsSandboxChannel)(nil))
}

func (i ApnsSandboxChannelMap) ToApnsSandboxChannelMapOutput() ApnsSandboxChannelMapOutput {
	return i.ToApnsSandboxChannelMapOutputWithContext(context.Background())
}

func (i ApnsSandboxChannelMap) ToApnsSandboxChannelMapOutputWithContext(ctx context.Context) ApnsSandboxChannelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelMapOutput)
}

type ApnsSandboxChannelOutput struct {
	*pulumi.OutputState
}

func (ApnsSandboxChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApnsSandboxChannel)(nil))
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput {
	return o
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput {
	return o
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelPtrOutput() ApnsSandboxChannelPtrOutput {
	return o.ToApnsSandboxChannelPtrOutputWithContext(context.Background())
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelPtrOutputWithContext(ctx context.Context) ApnsSandboxChannelPtrOutput {
	return o.ApplyT(func(v ApnsSandboxChannel) *ApnsSandboxChannel {
		return &v
	}).(ApnsSandboxChannelPtrOutput)
}

type ApnsSandboxChannelPtrOutput struct {
	*pulumi.OutputState
}

func (ApnsSandboxChannelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsSandboxChannel)(nil))
}

func (o ApnsSandboxChannelPtrOutput) ToApnsSandboxChannelPtrOutput() ApnsSandboxChannelPtrOutput {
	return o
}

func (o ApnsSandboxChannelPtrOutput) ToApnsSandboxChannelPtrOutputWithContext(ctx context.Context) ApnsSandboxChannelPtrOutput {
	return o
}

type ApnsSandboxChannelArrayOutput struct{ *pulumi.OutputState }

func (ApnsSandboxChannelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApnsSandboxChannel)(nil))
}

func (o ApnsSandboxChannelArrayOutput) ToApnsSandboxChannelArrayOutput() ApnsSandboxChannelArrayOutput {
	return o
}

func (o ApnsSandboxChannelArrayOutput) ToApnsSandboxChannelArrayOutputWithContext(ctx context.Context) ApnsSandboxChannelArrayOutput {
	return o
}

func (o ApnsSandboxChannelArrayOutput) Index(i pulumi.IntInput) ApnsSandboxChannelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApnsSandboxChannel {
		return vs[0].([]ApnsSandboxChannel)[vs[1].(int)]
	}).(ApnsSandboxChannelOutput)
}

type ApnsSandboxChannelMapOutput struct{ *pulumi.OutputState }

func (ApnsSandboxChannelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ApnsSandboxChannel)(nil))
}

func (o ApnsSandboxChannelMapOutput) ToApnsSandboxChannelMapOutput() ApnsSandboxChannelMapOutput {
	return o
}

func (o ApnsSandboxChannelMapOutput) ToApnsSandboxChannelMapOutputWithContext(ctx context.Context) ApnsSandboxChannelMapOutput {
	return o
}

func (o ApnsSandboxChannelMapOutput) MapIndex(k pulumi.StringInput) ApnsSandboxChannelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ApnsSandboxChannel {
		return vs[0].(map[string]ApnsSandboxChannel)[vs[1].(string)]
	}).(ApnsSandboxChannelOutput)
}

func init() {
	pulumi.RegisterOutputType(ApnsSandboxChannelOutput{})
	pulumi.RegisterOutputType(ApnsSandboxChannelPtrOutput{})
	pulumi.RegisterOutputType(ApnsSandboxChannelArrayOutput{})
	pulumi.RegisterOutputType(ApnsSandboxChannelMapOutput{})
}
