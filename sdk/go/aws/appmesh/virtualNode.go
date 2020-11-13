// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS App Mesh virtual node resource.
//
// ## Breaking Changes
//
// Because of backward incompatible API changes (read [here](https://github.com/awslabs/aws-app-mesh-examples/issues/92)), `appmesh.VirtualNode` resource definitions created with provider versions earlier than v2.3.0 will need to be modified:
//
// * Rename the `serviceName` attribute of the `dns` object to `hostname`.
//
// * Replace the `backends` attribute of the `spec` object with one or more `backend` configuration blocks,
// setting `virtualServiceName` to the name of the service.
//
// The state associated with existing resources will automatically be migrated.
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
// 		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualNodeSpecArgs{
// 				Backends: appmesh.VirtualNodeSpecBackendArray{
// 					&appmesh.VirtualNodeSpecBackendArgs{
// 						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
// 							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
// 						},
// 					},
// 				},
// 				Listener: &appmesh.VirtualNodeSpecListenerArgs{
// 					PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
// 						Port:     pulumi.Int(8080),
// 						Protocol: pulumi.String("http"),
// 					},
// 				},
// 				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
// 					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
// 						Hostname: pulumi.String("serviceb.simpleapp.local"),
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
// ### AWS Cloud Map Service Discovery
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/appmesh"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/servicediscovery"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := servicediscovery.NewHttpNamespace(ctx, "example", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualNodeSpecArgs{
// 				Backends: appmesh.VirtualNodeSpecBackendArray{
// 					&appmesh.VirtualNodeSpecBackendArgs{
// 						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
// 							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
// 						},
// 					},
// 				},
// 				Listener: &appmesh.VirtualNodeSpecListenerArgs{
// 					PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
// 						Port:     pulumi.Int(8080),
// 						Protocol: pulumi.String("http"),
// 					},
// 				},
// 				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
// 					AwsCloudMap: &appmesh.VirtualNodeSpecServiceDiscoveryAwsCloudMapArgs{
// 						Attributes: pulumi.StringMap{
// 							"stack": pulumi.String("blue"),
// 						},
// 						ServiceName:   pulumi.String("serviceb1"),
// 						NamespaceName: example.Name,
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
// ### Listener Health Check
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
// 		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualNodeSpecArgs{
// 				Backends: appmesh.VirtualNodeSpecBackendArray{
// 					&appmesh.VirtualNodeSpecBackendArgs{
// 						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
// 							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
// 						},
// 					},
// 				},
// 				Listener: &appmesh.VirtualNodeSpecListenerArgs{
// 					PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
// 						Port:     pulumi.Int(8080),
// 						Protocol: pulumi.String("http"),
// 					},
// 					HealthCheck: &appmesh.VirtualNodeSpecListenerHealthCheckArgs{
// 						Protocol:           pulumi.String("http"),
// 						Path:               pulumi.String("/ping"),
// 						HealthyThreshold:   pulumi.Int(2),
// 						UnhealthyThreshold: pulumi.Int(2),
// 						TimeoutMillis:      pulumi.Int(2000),
// 						IntervalMillis:     pulumi.Int(5000),
// 					},
// 				},
// 				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
// 					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
// 						Hostname: pulumi.String("serviceb.simpleapp.local"),
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
// ### Logging
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
// 		_, err := appmesh.NewVirtualNode(ctx, "serviceb1", &appmesh.VirtualNodeArgs{
// 			MeshName: pulumi.Any(aws_appmesh_mesh.Simple.Id),
// 			Spec: &appmesh.VirtualNodeSpecArgs{
// 				Backends: appmesh.VirtualNodeSpecBackendArray{
// 					&appmesh.VirtualNodeSpecBackendArgs{
// 						VirtualService: &appmesh.VirtualNodeSpecBackendVirtualServiceArgs{
// 							VirtualServiceName: pulumi.String("servicea.simpleapp.local"),
// 						},
// 					},
// 				},
// 				Listener: &appmesh.VirtualNodeSpecListenerArgs{
// 					PortMapping: &appmesh.VirtualNodeSpecListenerPortMappingArgs{
// 						Port:     pulumi.Int(8080),
// 						Protocol: pulumi.String("http"),
// 					},
// 				},
// 				ServiceDiscovery: &appmesh.VirtualNodeSpecServiceDiscoveryArgs{
// 					Dns: &appmesh.VirtualNodeSpecServiceDiscoveryDnsArgs{
// 						Hostname: pulumi.String("serviceb.simpleapp.local"),
// 					},
// 				},
// 				Logging: &appmesh.VirtualNodeSpecLoggingArgs{
// 					AccessLog: &appmesh.VirtualNodeSpecLoggingAccessLogArgs{
// 						File: &appmesh.VirtualNodeSpecLoggingAccessLogFileArgs{
// 							Path: pulumi.String("/dev/stdout"),
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
type VirtualNode struct {
	pulumi.CustomResourceState

	// The ARN of the virtual node.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The creation date of the virtual node.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The last update date of the virtual node.
	LastUpdatedDate pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringOutput `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringOutput `pulumi:"meshOwner"`
	// The name to use for the virtual node.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// The virtual node specification to apply.
	Spec VirtualNodeSpecOutput `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewVirtualNode registers a new resource with the given unique name, arguments, and options.
func NewVirtualNode(ctx *pulumi.Context,
	name string, args *VirtualNodeArgs, opts ...pulumi.ResourceOption) (*VirtualNode, error) {
	if args == nil || args.MeshName == nil {
		return nil, errors.New("missing required argument 'MeshName'")
	}
	if args == nil || args.Spec == nil {
		return nil, errors.New("missing required argument 'Spec'")
	}
	if args == nil {
		args = &VirtualNodeArgs{}
	}
	var resource VirtualNode
	err := ctx.RegisterResource("aws:appmesh/virtualNode:VirtualNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNode gets an existing VirtualNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNodeState, opts ...pulumi.ResourceOption) (*VirtualNode, error) {
	var resource VirtualNode
	err := ctx.ReadResource("aws:appmesh/virtualNode:VirtualNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNode resources.
type virtualNodeState struct {
	// The ARN of the virtual node.
	Arn *string `pulumi:"arn"`
	// The creation date of the virtual node.
	CreatedDate *string `pulumi:"createdDate"`
	// The last update date of the virtual node.
	LastUpdatedDate *string `pulumi:"lastUpdatedDate"`
	// The name of the service mesh in which to create the virtual node.
	MeshName *string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual node.
	Name *string `pulumi:"name"`
	// The resource owner's AWS account ID.
	ResourceOwner *string `pulumi:"resourceOwner"`
	// The virtual node specification to apply.
	Spec *VirtualNodeSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type VirtualNodeState struct {
	// The ARN of the virtual node.
	Arn pulumi.StringPtrInput
	// The creation date of the virtual node.
	CreatedDate pulumi.StringPtrInput
	// The last update date of the virtual node.
	LastUpdatedDate pulumi.StringPtrInput
	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringPtrInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual node.
	Name pulumi.StringPtrInput
	// The resource owner's AWS account ID.
	ResourceOwner pulumi.StringPtrInput
	// The virtual node specification to apply.
	Spec VirtualNodeSpecPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VirtualNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNodeState)(nil)).Elem()
}

type virtualNodeArgs struct {
	// The name of the service mesh in which to create the virtual node.
	MeshName string `pulumi:"meshName"`
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner *string `pulumi:"meshOwner"`
	// The name to use for the virtual node.
	Name *string `pulumi:"name"`
	// The virtual node specification to apply.
	Spec VirtualNodeSpec `pulumi:"spec"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a VirtualNode resource.
type VirtualNodeArgs struct {
	// The name of the service mesh in which to create the virtual node.
	MeshName pulumi.StringInput
	// The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
	MeshOwner pulumi.StringPtrInput
	// The name to use for the virtual node.
	Name pulumi.StringPtrInput
	// The virtual node specification to apply.
	Spec VirtualNodeSpecInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (VirtualNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNodeArgs)(nil)).Elem()
}

type VirtualNodeInput interface {
	pulumi.Input

	ToVirtualNodeOutput() VirtualNodeOutput
	ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput
}

func (VirtualNode) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNode)(nil)).Elem()
}

func (i VirtualNode) ToVirtualNodeOutput() VirtualNodeOutput {
	return i.ToVirtualNodeOutputWithContext(context.Background())
}

func (i VirtualNode) ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNodeOutput)
}

type VirtualNodeOutput struct {
	*pulumi.OutputState
}

func (VirtualNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNodeOutput)(nil)).Elem()
}

func (o VirtualNodeOutput) ToVirtualNodeOutput() VirtualNodeOutput {
	return o
}

func (o VirtualNodeOutput) ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNodeOutput{})
}
