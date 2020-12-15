// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh route resource.
//
// ## Example Usage
// ### HTTP Routing
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
// 			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
// 			Spec: &appmesh.RouteSpecArgs{
// 				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
// 					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
// 						Prefix: pulumi.String("/"),
// 					},
// 					Action: &appmesh.RouteSpecHttpRouteActionArgs{
// 						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
// 							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
// 								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
// 								Weight:      pulumi.Int(90),
// 							},
// 							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
// 								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb2.Name),
// 								Weight:      pulumi.Int(10),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### HTTP Header Routing
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
// 			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
// 			Spec: &appmesh.RouteSpecArgs{
// 				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
// 					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
// 						Method: pulumi.String("POST"),
// 						Prefix: pulumi.String("/"),
// 						Scheme: pulumi.String("https"),
// 						Headers: appmesh.RouteSpecHttpRouteMatchHeaderArray{
// 							&appmesh.RouteSpecHttpRouteMatchHeaderArgs{
// 								Name: pulumi.String("clientRequestId"),
// 								Match: &appmesh.RouteSpecHttpRouteMatchHeaderMatchArgs{
// 									Prefix: pulumi.String("123"),
// 								},
// 							},
// 						},
// 					},
// 					Action: &appmesh.RouteSpecHttpRouteActionArgs{
// 						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
// 							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
// 								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
// 								Weight:      pulumi.Int(100),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Retry Policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
// 			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
// 			Spec: &appmesh.RouteSpecArgs{
// 				HttpRoute: &appmesh.RouteSpecHttpRouteArgs{
// 					Match: &appmesh.RouteSpecHttpRouteMatchArgs{
// 						Prefix: pulumi.String("/"),
// 					},
// 					RetryPolicy: &appmesh.RouteSpecHttpRouteRetryPolicyArgs{
// 						HttpRetryEvents: pulumi.StringArray{
// 							pulumi.String("server-error"),
// 						},
// 						MaxRetries: pulumi.Int(1),
// 						PerRetryTimeout: &appmesh.RouteSpecHttpRouteRetryPolicyPerRetryTimeoutArgs{
// 							Unit:  pulumi.String("s"),
// 							Value: pulumi.Int(15),
// 						},
// 					},
// 					Action: &appmesh.RouteSpecHttpRouteActionArgs{
// 						WeightedTargets: appmesh.RouteSpecHttpRouteActionWeightedTargetArray{
// 							&appmesh.RouteSpecHttpRouteActionWeightedTargetArgs{
// 								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb.Name),
// 								Weight:      pulumi.Int(100),
// 							},
// 						},
// 					},
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### TCP Routing
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := appmesh.NewRoute(ctx, "serviceb", &appmesh.RouteArgs{
// 			MeshName:          pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
// 			Spec: &appmesh.RouteSpecArgs{
// 				TcpRoute: &appmesh.RouteSpecTcpRouteArgs{
// 					Action: &appmesh.RouteSpecTcpRouteActionArgs{
// 						WeightedTargets: appmesh.RouteSpecTcpRouteActionWeightedTargetArray{
// 							&appmesh.RouteSpecTcpRouteActionWeightedTargetArgs{
// 								VirtualNode: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
// 								Weight:      pulumi.Int(100),
// 							},
// 						},
// 					},
// 				},
// 			},
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
// App Mesh virtual routes can be imported using `mesh_name` and `virtual_router_name` together with the route's `name`, e.g.
//
// ```sh
//  $ pulumi import aws:appmesh/route:Route serviceb simpleapp/serviceB/serviceB-route
// ```
//
//  [1]/docs/providers/aws/index.html
type Route struct {
	pulumi.CustomResourceState

	// The ARN of the route.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the route.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the route.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The route specification to apply.
	Spec RouteSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringOutput `pulumi:"virtualRouterName"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualRouterName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualRouterName'")
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
	// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The route specification to apply.
	Spec *RouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName *string `pulumi:"virtualRouterName"`
}

type RouteState struct {
	// The ARN of the route.
	Arn pulumi.StringPtrInput
	// The creation date of the route.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the route.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The route specification to apply.
	Spec RouteSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the route. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The route specification to apply.
	Spec RouteSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// The name of the service mesh in which to create the route. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the route. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The route specification to apply.
	Spec RouteSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The name of the virtual router in which to create the route. Must be between 1 and 255 characters in length.
	VirtualRouterName pulumi.StringInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
}
