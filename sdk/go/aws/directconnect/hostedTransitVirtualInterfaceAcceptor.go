// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage the accepter's side of a Direct Connect hosted transit virtual interface.
// This resource accepts ownership of a transit virtual interface created by another AWS account.
//
// > **NOTE:** AWS allows a Direct Connect hosted transit virtual interface to be deleted from either the allocator's or accepter's side. However, this provider only allows the Direct Connect hosted transit virtual interface to be deleted from the allocator's side by removing the corresponding `directconnect.HostedTransitVirtualInterface` resource from your configuration. Removing a `directconnect.HostedTransitVirtualInterfaceAcceptor` resource from your configuration will remove it from your statefile and management, **but will not delete the Direct Connect virtual interface.**
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "accepter", nil)
// 		if err != nil {
// 			return err
// 		}
// 		accepterCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		example, err := directconnect.NewGateway(ctx, "example", &directconnect.GatewayArgs{
// 			AmazonSideAsn: pulumi.String("64512"),
// 		}, pulumi.Provider(aws.Accepter))
// 		if err != nil {
// 			return err
// 		}
// 		creator, err := directconnect.NewHostedTransitVirtualInterface(ctx, "creator", &directconnect.HostedTransitVirtualInterfaceArgs{
// 			ConnectionId:   pulumi.String("dxcon-zzzzzzzz"),
// 			OwnerAccountId: pulumi.String(accepterCallerIdentity.AccountId),
// 			Vlan:           pulumi.Int(4094),
// 			AddressFamily:  pulumi.String("ipv4"),
// 			BgpAsn:         pulumi.Int(65352),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			example,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewHostedTransitVirtualInterfaceAcceptor(ctx, "accepterHostedTransitVirtualInterfaceAcceptor", &directconnect.HostedTransitVirtualInterfaceAcceptorArgs{
// 			VirtualInterfaceId: creator.ID(),
// 			DxGatewayId:        example.ID(),
// 			Tags: pulumi.StringMap{
// 				"Side": pulumi.String("Accepter"),
// 			},
// 		}, pulumi.Provider(aws.Accepter))
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
// Direct Connect hosted transit virtual interfaces can be imported using the `vif id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor test dxvif-33cc44dd
// ```
type HostedTransitVirtualInterfaceAcceptor struct {
	pulumi.CustomResourceState

	// The ARN of the virtual interface.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringOutput `pulumi:"virtualInterfaceId"`
}

// NewHostedTransitVirtualInterfaceAcceptor registers a new resource with the given unique name, arguments, and options.
func NewHostedTransitVirtualInterfaceAcceptor(ctx *pulumi.Context,
	name string, args *HostedTransitVirtualInterfaceAcceptorArgs, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterfaceAcceptor, error) {
	if args == nil || args.DxGatewayId == nil {
		return nil, errors.New("missing required argument 'DxGatewayId'")
	}
	if args == nil || args.VirtualInterfaceId == nil {
		return nil, errors.New("missing required argument 'VirtualInterfaceId'")
	}
	if args == nil {
		args = &HostedTransitVirtualInterfaceAcceptorArgs{}
	}
	var resource HostedTransitVirtualInterfaceAcceptor
	err := ctx.RegisterResource("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostedTransitVirtualInterfaceAcceptor gets an existing HostedTransitVirtualInterfaceAcceptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostedTransitVirtualInterfaceAcceptor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostedTransitVirtualInterfaceAcceptorState, opts ...pulumi.ResourceOption) (*HostedTransitVirtualInterfaceAcceptor, error) {
	var resource HostedTransitVirtualInterfaceAcceptor
	err := ctx.ReadResource("aws:directconnect/hostedTransitVirtualInterfaceAcceptor:HostedTransitVirtualInterfaceAcceptor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostedTransitVirtualInterfaceAcceptor resources.
type hostedTransitVirtualInterfaceAcceptorState struct {
	// The ARN of the virtual interface.
	Arn *string `pulumi:"arn"`
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId *string `pulumi:"virtualInterfaceId"`
}

type HostedTransitVirtualInterfaceAcceptorState struct {
	// The ARN of the virtual interface.
	Arn pulumi.StringPtrInput
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringPtrInput
}

func (HostedTransitVirtualInterfaceAcceptorState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceAcceptorState)(nil)).Elem()
}

type hostedTransitVirtualInterfaceAcceptorArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId string `pulumi:"dxGatewayId"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId string `pulumi:"virtualInterfaceId"`
}

// The set of arguments for constructing a HostedTransitVirtualInterfaceAcceptor resource.
type HostedTransitVirtualInterfaceAcceptorArgs struct {
	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayId pulumi.StringInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceId pulumi.StringInput
}

func (HostedTransitVirtualInterfaceAcceptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostedTransitVirtualInterfaceAcceptorArgs)(nil)).Elem()
}
