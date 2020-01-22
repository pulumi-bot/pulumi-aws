// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appmesh

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides an AWS App Mesh route resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/appmesh_route.html.markdown.
type Route struct {
	pulumi.CustomResourceState

	// The ARN of the route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the route.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The name to use for the route.
	Name pulumi.StringOutput `pulumi:"name"`
	// The route specification to apply.
	Spec RouteSpecOutput `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The name of the virtual router in which to create the route.
	VirtualRouterName pulumi.StringOutput `pulumi:"virtualRouterName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil || args.VirtualRouterName == nil {
		return nil, errors.New("missing required argument 'VirtualRouterName'")
	}
	if args == nil {
		args = &RouteArgs{}
	}
	var resource Route
	err := ctx.RegisterResource("aws:appmesh/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws:appmesh/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// The ARN of the route.
	Arn *string `pulumi:"arn"`
	// The creation date of the route.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the route.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the route.
	MeshName *string `pulumi:"meshName"`
	// The name to use for the route.
	Name *string `pulumi:"name"`
	// The route specification to apply.
	Spec *RouteSpec `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the virtual router in which to create the route.
	VirtualRouterName *string `pulumi:"virtualRouterName"`
}

type RouteState struct {
	// The ARN of the route.
	Arn pulumi.StringPtrInput
	// The creation date of the route.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the route.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the route.
	MeshName pulumi.StringPtrInput
	// The name to use for the route.
	Name pulumi.StringPtrInput
	// The route specification to apply.
	Spec RouteSpecPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The name of the virtual router in which to create the route.
	VirtualRouterName pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The name of the service mesh in which to create the route.
	MeshName string `pulumi:"meshName"`
	// The name to use for the route.
	Name *string `pulumi:"name"`
	// The route specification to apply.
	Spec RouteSpec `pulumi:"spec"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The name of the virtual router in which to create the route.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The name of the service mesh in which to create the route.
	MeshName pulumi.StringInput
	// The name to use for the route.
	Name pulumi.StringPtrInput
	// The route specification to apply.
	Spec RouteSpecInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The name of the virtual router in which to create the route.
	VirtualRouterName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

