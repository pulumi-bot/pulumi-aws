// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a Model for a REST API Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/apigateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewModel(ctx, "myDemoModel", &apigateway.ModelArgs{
// 			RestApi:     myDemoAPI.ID(),
// 			Description: pulumi.String("a JSON schema"),
// 			ContentType: pulumi.String("application/json"),
// 			Schema:      pulumi.String(fmt.Sprintf("%v%v%v", "{\n", "  \"type\": \"object\"\n", "}\n")),
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
// `aws_api_gateway_model` can be imported using `REST-API-ID/NAME`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/model:Model example 12345abcde/example
// ```
type Model struct {
	pulumi.CustomResourceState

	// The content type of the model
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// The description of the model
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the model
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrOutput `pulumi:"schema"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContentType == nil {
		return nil, errors.New("invalid value for required argument 'ContentType'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	var resource Model
	err := ctx.RegisterResource("aws:apigateway/model:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("aws:apigateway/model:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
	// The content type of the model
	ContentType *string `pulumi:"contentType"`
	// The description of the model
	Description *string `pulumi:"description"`
	// The name of the model
	Name *string `pulumi:"name"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema *string `pulumi:"schema"`
}

type ModelState struct {
	// The content type of the model
	ContentType pulumi.StringPtrInput
	// The description of the model
	Description pulumi.StringPtrInput
	// The name of the model
	Name pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrInput
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	// The content type of the model
	ContentType string `pulumi:"contentType"`
	// The description of the model
	Description *string `pulumi:"description"`
	// The name of the model
	Name *string `pulumi:"name"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
	// The schema of the model in a JSON form
	Schema *string `pulumi:"schema"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	// The content type of the model
	ContentType pulumi.StringInput
	// The description of the model
	Description pulumi.StringPtrInput
	// The name of the model
	Name pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.Input
	// The schema of the model in a JSON form
	Schema pulumi.StringPtrInput
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}

type ModelInput interface {
	pulumi.Input

	ToModelOutput() ModelOutput
	ToModelOutputWithContext(ctx context.Context) ModelOutput
}

func (*Model) ElementType() reflect.Type {
	return reflect.TypeOf((*Model)(nil))
}

func (i *Model) ToModelOutput() ModelOutput {
	return i.ToModelOutputWithContext(context.Background())
}

func (i *Model) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput)
}

func (i *Model) ToModelPtrOutput() ModelPtrOutput {
	return i.ToModelPtrOutputWithContext(context.Background())
}

func (i *Model) ToModelPtrOutputWithContext(ctx context.Context) ModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPtrOutput)
}

type ModelPtrInput interface {
	pulumi.Input

	ToModelPtrOutput() ModelPtrOutput
	ToModelPtrOutputWithContext(ctx context.Context) ModelPtrOutput
}

type modelPtrType ModelArgs

func (*modelPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil))
}

func (i *modelPtrType) ToModelPtrOutput() ModelPtrOutput {
	return i.ToModelPtrOutputWithContext(context.Background())
}

func (i *modelPtrType) ToModelPtrOutputWithContext(ctx context.Context) ModelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput).ToModelPtrOutput()
}

// ModelArrayInput is an input type that accepts ModelArray and ModelArrayOutput values.
// You can construct a concrete instance of `ModelArrayInput` via:
//
//          ModelArray{ ModelArgs{...} }
type ModelArrayInput interface {
	pulumi.Input

	ToModelArrayOutput() ModelArrayOutput
	ToModelArrayOutputWithContext(context.Context) ModelArrayOutput
}

type ModelArray []ModelInput

func (ModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Model)(nil))
}

func (i ModelArray) ToModelArrayOutput() ModelArrayOutput {
	return i.ToModelArrayOutputWithContext(context.Background())
}

func (i ModelArray) ToModelArrayOutputWithContext(ctx context.Context) ModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelArrayOutput)
}

// ModelMapInput is an input type that accepts ModelMap and ModelMapOutput values.
// You can construct a concrete instance of `ModelMapInput` via:
//
//          ModelMap{ "key": ModelArgs{...} }
type ModelMapInput interface {
	pulumi.Input

	ToModelMapOutput() ModelMapOutput
	ToModelMapOutputWithContext(context.Context) ModelMapOutput
}

type ModelMap map[string]ModelInput

func (ModelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Model)(nil))
}

func (i ModelMap) ToModelMapOutput() ModelMapOutput {
	return i.ToModelMapOutputWithContext(context.Background())
}

func (i ModelMap) ToModelMapOutputWithContext(ctx context.Context) ModelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelMapOutput)
}

type ModelOutput struct {
	*pulumi.OutputState
}

func (ModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Model)(nil))
}

func (o ModelOutput) ToModelOutput() ModelOutput {
	return o
}

func (o ModelOutput) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return o
}

func (o ModelOutput) ToModelPtrOutput() ModelPtrOutput {
	return o.ToModelPtrOutputWithContext(context.Background())
}

func (o ModelOutput) ToModelPtrOutputWithContext(ctx context.Context) ModelPtrOutput {
	return o.ApplyT(func(v Model) *Model {
		return &v
	}).(ModelPtrOutput)
}

type ModelPtrOutput struct {
	*pulumi.OutputState
}

func (ModelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil))
}

func (o ModelPtrOutput) ToModelPtrOutput() ModelPtrOutput {
	return o
}

func (o ModelPtrOutput) ToModelPtrOutputWithContext(ctx context.Context) ModelPtrOutput {
	return o
}

type ModelArrayOutput struct{ *pulumi.OutputState }

func (ModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Model)(nil))
}

func (o ModelArrayOutput) ToModelArrayOutput() ModelArrayOutput {
	return o
}

func (o ModelArrayOutput) ToModelArrayOutputWithContext(ctx context.Context) ModelArrayOutput {
	return o
}

func (o ModelArrayOutput) Index(i pulumi.IntInput) ModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Model {
		return vs[0].([]Model)[vs[1].(int)]
	}).(ModelOutput)
}

type ModelMapOutput struct{ *pulumi.OutputState }

func (ModelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Model)(nil))
}

func (o ModelMapOutput) ToModelMapOutput() ModelMapOutput {
	return o
}

func (o ModelMapOutput) ToModelMapOutputWithContext(ctx context.Context) ModelMapOutput {
	return o
}

func (o ModelMapOutput) MapIndex(k pulumi.StringInput) ModelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Model {
		return vs[0].(map[string]Model)[vs[1].(string)]
	}).(ModelOutput)
}

func init() {
	pulumi.RegisterOutputType(ModelOutput{})
	pulumi.RegisterOutputType(ModelPtrOutput{})
	pulumi.RegisterOutputType(ModelArrayOutput{})
	pulumi.RegisterOutputType(ModelMapOutput{})
}
