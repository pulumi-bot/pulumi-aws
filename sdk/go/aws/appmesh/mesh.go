// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh service mesh resource.
//
// ## Example Usage
// ### Basic
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
// 		_, err := appmesh.NewMesh(ctx, "simple", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Egress Filter
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
// 		_, err := appmesh.NewMesh(ctx, "simple", &appmesh.MeshArgs{
// 			Spec: &appmesh.MeshSpecArgs{
// 				EgressFilter: &appmesh.MeshSpecEgressFilterArgs{
// 					Type: pulumi.String("ALLOW_ALL"),
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
type Mesh struct {
	pulumi.CustomResourceState

	// The ARN of the service mesh.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the service mesh.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the service mesh.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The AWS account ID of the service mesh's owner.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the service mesh.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The service mesh specification to apply.
	Spec MeshSpecPtrOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMesh registers a new resource with the given unique name, arguments, and options.
func NewMesh(ctx *pulumi.Context,
	name string, args *MeshArgs, opts ...pulumi.ResourceOption) (*Mesh, error) {
	if args == nil {
		args = &MeshArgs{}
	}

	var resource Mesh
	err := ctx.RegisterResource("aws:appmesh/mesh:Mesh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMesh gets an existing Mesh resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMesh(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MeshState, opts ...pulumi.ResourceOption) (*Mesh, error) {
	var resource Mesh
	err := ctx.ReadResource("aws:appmesh/mesh:Mesh", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mesh resources.
type meshState struct {
	// The ARN of the service mesh.
	Arn *string `pulumi:"arn"`
	// The creation date of the service mesh.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the service mesh.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The AWS account ID of the service mesh's owner.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the service mesh.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The service mesh specification to apply.
	Spec *MeshSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MeshState struct {
	// The ARN of the service mesh.
	Arn pulumi.StringPtrInput
	// The creation date of the service mesh.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the service mesh.
	LastUpdatedDate pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the service mesh.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The service mesh specification to apply.
	Spec MeshSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshState) ElementType() reflect.Type {
	return reflect.TypeOf((*meshState)(nil)).Elem()
}

type meshArgs struct {
	// The name to use for the service mesh.
	Name *string `pulumi:"name"`
	// The service mesh specification to apply.
	Spec *MeshSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Mesh resource.
type MeshArgs struct {
	// The name to use for the service mesh.
	Name pulumi.StringPtrInput
	// The service mesh specification to apply.
	Spec MeshSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MeshArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*meshArgs)(nil)).Elem()
}

type MeshInput interface {
	pulumi.Input

	ToMeshOutput() MeshOutput
	ToMeshOutputWithContext(ctx context.Context) MeshOutput
}

func (Mesh) ElementType() reflect.Type {
	return reflect.TypeOf((*Mesh)(nil)).Elem()
}

func (i Mesh) ToMeshOutput() MeshOutput {
	return i.ToMeshOutputWithContext(context.Background())
}

func (i Mesh) ToMeshOutputWithContext(ctx context.Context) MeshOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MeshOutput)
}

type MeshOutput struct {
	*pulumi.OutputState
}

func (MeshOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MeshOutput)(nil)).Elem()
}

func (o MeshOutput) ToMeshOutput() MeshOutput {
	return o
}

func (o MeshOutput) ToMeshOutputWithContext(ctx context.Context) MeshOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MeshOutput{})
}
