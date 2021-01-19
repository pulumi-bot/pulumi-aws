// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway Resource.
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
// 		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
// 			Description: pulumi.String("This is my API for demonstration purposes"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
// 			RestApi:  myDemoAPI.ID(),
// 			ParentId: myDemoAPI.RootResourceId,
// 			PathPart: pulumi.String("mydemoresource"),
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
// `aws_api_gateway_resource` can be imported using `REST-API-ID/RESOURCE-ID`, e.g.
//
// ```sh
//  $ pulumi import aws:apigateway/resource:Resource example 12345abcde/67890fghij
// ```
type Resource struct {
	pulumi.CustomResourceState

	// The ID of the parent API resource
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// The complete path for this API resource, including all parent paths.
	Path pulumi.StringOutput `pulumi:"path"`
	// The last path segment of this API resource.
	PathPart pulumi.StringOutput `pulumi:"pathPart"`
	// The ID of the associated REST API
	RestApi pulumi.StringOutput `pulumi:"restApi"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.PathPart == nil {
		return nil, errors.New("invalid value for required argument 'PathPart'")
	}
	if args.RestApi == nil {
		return nil, errors.New("invalid value for required argument 'RestApi'")
	}
	var resource Resource
	err := ctx.RegisterResource("aws:apigateway/resource:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("aws:apigateway/resource:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
	// The ID of the parent API resource
	ParentId *string `pulumi:"parentId"`
	// The complete path for this API resource, including all parent paths.
	Path *string `pulumi:"path"`
	// The last path segment of this API resource.
	PathPart *string `pulumi:"pathPart"`
	// The ID of the associated REST API
	RestApi *string `pulumi:"restApi"`
}

type ResourceState struct {
	// The ID of the parent API resource
	ParentId pulumi.StringPtrInput
	// The complete path for this API resource, including all parent paths.
	Path pulumi.StringPtrInput
	// The last path segment of this API resource.
	PathPart pulumi.StringPtrInput
	// The ID of the associated REST API
	RestApi pulumi.StringPtrInput
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// The ID of the parent API resource
	ParentId string `pulumi:"parentId"`
	// The last path segment of this API resource.
	PathPart string `pulumi:"pathPart"`
	// The ID of the associated REST API
	RestApi interface{} `pulumi:"restApi"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// The ID of the parent API resource
	ParentId pulumi.StringInput
	// The last path segment of this API resource.
	PathPart pulumi.StringInput
	// The ID of the associated REST API
	RestApi pulumi.Input
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil))
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

func (i *Resource) ToResourcePtrOutput() ResourcePtrOutput {
	return i.ToResourcePtrOutputWithContext(context.Background())
}

func (i *Resource) ToResourcePtrOutputWithContext(ctx context.Context) ResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePtrOutput)
}

type ResourcePtrInput interface {
	pulumi.Input

	ToResourcePtrOutput() ResourcePtrOutput
	ToResourcePtrOutputWithContext(ctx context.Context) ResourcePtrOutput
}

type resourcePtrType ResourceArgs

func (*resourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil))
}

func (i *resourcePtrType) ToResourcePtrOutput() ResourcePtrOutput {
	return i.ToResourcePtrOutputWithContext(context.Background())
}

func (i *resourcePtrType) ToResourcePtrOutputWithContext(ctx context.Context) ResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourcePtrOutput)
}

// ResourceArrayInput is an input type that accepts ResourceArray and ResourceArrayOutput values.
// You can construct a concrete instance of `ResourceArrayInput` via:
//
//          ResourceArray{ ResourceArgs{...} }
type ResourceArrayInput interface {
	pulumi.Input

	ToResourceArrayOutput() ResourceArrayOutput
	ToResourceArrayOutputWithContext(context.Context) ResourceArrayOutput
}

type ResourceArray []ResourceInput

func (ResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Resource)(nil))
}

func (i ResourceArray) ToResourceArrayOutput() ResourceArrayOutput {
	return i.ToResourceArrayOutputWithContext(context.Background())
}

func (i ResourceArray) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceArrayOutput)
}

// ResourceMapInput is an input type that accepts ResourceMap and ResourceMapOutput values.
// You can construct a concrete instance of `ResourceMapInput` via:
//
//          ResourceMap{ "key": ResourceArgs{...} }
type ResourceMapInput interface {
	pulumi.Input

	ToResourceMapOutput() ResourceMapOutput
	ToResourceMapOutputWithContext(context.Context) ResourceMapOutput
}

type ResourceMap map[string]ResourceInput

func (ResourceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Resource)(nil))
}

func (i ResourceMap) ToResourceMapOutput() ResourceMapOutput {
	return i.ToResourceMapOutputWithContext(context.Background())
}

func (i ResourceMap) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceMapOutput)
}

type ResourceOutput struct {
	*pulumi.OutputState
}

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil))
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourcePtrOutput() ResourcePtrOutput {
	return o.ToResourcePtrOutputWithContext(context.Background())
}

func (o ResourceOutput) ToResourcePtrOutputWithContext(ctx context.Context) ResourcePtrOutput {
	return o.ApplyT(func(v Resource) *Resource {
		return &v
	}).(ResourcePtrOutput)
}

type ResourcePtrOutput struct {
	*pulumi.OutputState
}

func (ResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Resource)(nil))
}

func (o ResourcePtrOutput) ToResourcePtrOutput() ResourcePtrOutput {
	return o
}

func (o ResourcePtrOutput) ToResourcePtrOutputWithContext(ctx context.Context) ResourcePtrOutput {
	return o
}

type ResourceArrayOutput struct{ *pulumi.OutputState }

func (ResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Resource)(nil))
}

func (o ResourceArrayOutput) ToResourceArrayOutput() ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) ToResourceArrayOutputWithContext(ctx context.Context) ResourceArrayOutput {
	return o
}

func (o ResourceArrayOutput) Index(i pulumi.IntInput) ResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Resource {
		return vs[0].([]Resource)[vs[1].(int)]
	}).(ResourceOutput)
}

type ResourceMapOutput struct{ *pulumi.OutputState }

func (ResourceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Resource)(nil))
}

func (o ResourceMapOutput) ToResourceMapOutput() ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) ToResourceMapOutputWithContext(ctx context.Context) ResourceMapOutput {
	return o
}

func (o ResourceMapOutput) MapIndex(k pulumi.StringInput) ResourceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Resource {
		return vs[0].(map[string]Resource)[vs[1].(string)]
	}).(ResourceOutput)
}

func init() {
	pulumi.RegisterOutputType(ResourceOutput{})
	pulumi.RegisterOutputType(ResourcePtrOutput{})
	pulumi.RegisterOutputType(ResourceArrayOutput{})
	pulumi.RegisterOutputType(ResourceMapOutput{})
}
