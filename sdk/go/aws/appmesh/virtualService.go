// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh virtual service resource.
//
// ## Example Usage
// ### Virtual Node Provider
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
// 		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualServiceSpecArgs{
// 				Provider: &appmesh.VirtualServiceSpecProviderArgs{
// 					VirtualNode: &appmesh.VirtualServiceSpecProviderVirtualNodeArgs{
// 						VirtualNodeName: pulumi.Any(aws_appmesh_virtual_node.Serviceb1.Name),
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
// ### Virtual Router Provider
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
// 		_, err := appmesh.NewVirtualService(ctx, "servicea", &appmesh.VirtualServiceArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualServiceSpecArgs{
// 				Provider: &appmesh.VirtualServiceSpecProviderArgs{
// 					VirtualRouter: &appmesh.VirtualServiceSpecProviderVirtualRouterArgs{
// 						VirtualRouterName: pulumi.Any(aws_appmesh_virtual_router.Serviceb.Name),
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
// App Mesh virtual services can be imported using `mesh_name` together with the virtual service's `name`, e.g.
//
// ```sh
//  $ pulumi import aws:appmesh/virtualService:VirtualService servicea simpleapp/servicea.simpleapp.local
// ```
//
//  [1]/docs/providers/aws/index.html
type VirtualService struct {
	pulumi.CustomResourceState

	// The ARN of the virtual service.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource VirtualService
	err := ctx.RegisterResource("aws:appmesh/virtualService:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("aws:appmesh/virtualService:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
	// The ARN of the virtual service.
	Arn *string `pulumi:"arn"`
	// The creation date of the virtual service.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the virtual service.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The virtual service specification to apply.
	Spec *VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VirtualServiceState struct {
	// The ARN of the virtual service.
	Arn pulumi.StringPtrInput
	// The creation date of the virtual service.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the virtual service.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name *string `pulumi:"name"`
	// The virtual service specification to apply.
	Spec VirtualServiceSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// The name of the service mesh in which to create the virtual service. Must be between 1 and 255 characters in length.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual service. Must be between 1 and 255 characters in length.
	Name pulumi.StringPtrInput
	// The virtual service specification to apply.
	Spec VirtualServiceSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}

type VirtualServiceInput interface {
	pulumi.Input

	ToVirtualServiceOutput() VirtualServiceOutput
	ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput
}

func (VirtualService) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualService)(nil)).Elem()
}

func (i VirtualService) ToVirtualServiceOutput() VirtualServiceOutput {
	return i.ToVirtualServiceOutputWithContext(context.Background())
}

func (i VirtualService) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceOutput)
}

type VirtualServiceOutput struct {
	*pulumi.OutputState
}

func (VirtualServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualServiceOutput)(nil)).Elem()
}

func (o VirtualServiceOutput) ToVirtualServiceOutput() VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualServiceOutput{})
}
