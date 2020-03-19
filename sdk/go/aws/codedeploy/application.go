// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package codedeploy

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a CodeDeploy application to be used as a basis for deployments
//
// {{% examples %}}
// {{% /examples %}}
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codedeploy_app.html.markdown.
type Application struct {
	pulumi.CustomResourceState

	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform pulumi.StringPtrOutput `pulumi:"computePlatform"`
	// The name of the application.
	Name pulumi.StringOutput `pulumi:"name"`
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}
	var resource Application
	err := ctx.RegisterResource("aws:codedeploy/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("aws:codedeploy/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform *string `pulumi:"computePlatform"`
	// The name of the application.
	Name *string `pulumi:"name"`
	UniqueId *string `pulumi:"uniqueId"`
}

type ApplicationState struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform pulumi.StringPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	UniqueId pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform *string `pulumi:"computePlatform"`
	// The name of the application.
	Name *string `pulumi:"name"`
	UniqueId *string `pulumi:"uniqueId"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// The compute platform can either be `ECS`, `Lambda`, or `Server`. Default is `Server`.
	ComputePlatform pulumi.StringPtrInput
	// The name of the application.
	Name pulumi.StringPtrInput
	UniqueId pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

