// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway API Key.
//
// > **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := apigateway.NewApiKey(ctx, "myDemoApiKey", nil)
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
// API Gateway Keys can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/apiKey:ApiKey my_demo_key 8bklk8bl1k3sB38D9B3l0enyWT8c09B30lkq0blk
// ```
type ApiKey struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the API key
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The last update date of the API key
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the API key
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewApiKey registers a new resource with the given unique name, arguments, and options.
func NewApiKey(ctx *pulumi.Context,
	name string, args *ApiKeyArgs, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	if args == nil {
		args = &ApiKeyArgs{}
	}

	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource ApiKey
	err := ctx.RegisterResource("aws:apigateway/apiKey:ApiKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiKey gets an existing ApiKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiKeyState, opts ...pulumi.ResourceOption) (*ApiKey, error) {
	var resource ApiKey
	err := ctx.ReadResource("aws:apigateway/apiKey:ApiKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiKey resources.
type apiKeyState struct {
	// Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// The creation date of the API key
	CreatedDate *string `pulumi:"createdDate"`
	// The API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The last update date of the API key
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the API key
	Name *string `pulumi:"name"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value *string `pulumi:"value"`
}

type ApiKeyState struct {
	// Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// The creation date of the API key
	CreatedDate pulumi.StringPtrInput
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The last update date of the API key
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the API key
	Name pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringPtrInput
}

func (ApiKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyState)(nil)).Elem()
}

type apiKeyArgs struct {
	// The API key description. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the API key
	Name *string `pulumi:"name"`
	// Key-value map of resource tags
	Tags map[string]string `pulumi:"tags"`
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value *string `pulumi:"value"`
}

// The set of arguments for constructing a ApiKey resource.
type ApiKeyArgs struct {
	// The API key description. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// Specifies whether the API key can be used by callers. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the API key
	Name pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.StringMapInput
	// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
	Value pulumi.StringPtrInput
}

func (ApiKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiKeyArgs)(nil)).Elem()
}

type ApiKeyInput interface {
	pulumi.Input

	ToApiKeyOutput() ApiKeyOutput
	ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput
}

type ApiKeyPtrInput interface {
	pulumi.Input

	ToApiKeyPtrOutput() ApiKeyPtrOutput
	ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput
}

func (ApiKey) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKey)(nil)).Elem()
}

func (i ApiKey) ToApiKeyOutput() ApiKeyOutput {
	return i.ToApiKeyOutputWithContext(context.Background())
}

func (i ApiKey) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyOutput)
}

func (i ApiKey) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return i.ToApiKeyPtrOutputWithContext(context.Background())
}

func (i ApiKey) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiKeyPtrOutput)
}

type ApiKeyOutput struct {
	*pulumi.OutputState
}

func (ApiKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiKeyOutput)(nil)).Elem()
}

func (o ApiKeyOutput) ToApiKeyOutput() ApiKeyOutput {
	return o
}

func (o ApiKeyOutput) ToApiKeyOutputWithContext(ctx context.Context) ApiKeyOutput {
	return o
}

type ApiKeyPtrOutput struct {
	*pulumi.OutputState
}

func (ApiKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiKey)(nil)).Elem()
}

func (o ApiKeyPtrOutput) ToApiKeyPtrOutput() ApiKeyPtrOutput {
	return o
}

func (o ApiKeyPtrOutput) ToApiKeyPtrOutputWithContext(ctx context.Context) ApiKeyPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ApiKeyOutput{})
	pulumi.RegisterOutputType(ApiKeyPtrOutput{})
}
