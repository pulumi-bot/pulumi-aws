// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package directconnect

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Associates a Direct Connect Gateway with a VGW or transit gateway.
//
// To create a cross-account association, create an `directconnect.GatewayAssociationProposal` resource
// in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
// by creating an `directconnect.GatewayAssociation` resource with the `proposalId` and `associatedGatewayOwnerAccountId` attributes set.
//
// ## Example Usage
// ### VPN Gateway Association
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
// 			AmazonSideAsn: pulumi.String("64512"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.255.255.0/28"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpnGateway, err := ec2.NewVpnGateway(ctx, "exampleVpnGateway", &ec2.VpnGatewayArgs{
// 			VpcId: exampleVpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
// 			DxGatewayId:         exampleGateway.ID(),
// 			AssociatedGatewayId: exampleVpnGateway.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Transit Gateway Association
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2transitgateway"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
// 			AmazonSideAsn: pulumi.String("64512"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleTransitGateway, err := ec2transitgateway.NewTransitGateway(ctx, "exampleTransitGateway", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
// 			DxGatewayId:         exampleGateway.ID(),
// 			AssociatedGatewayId: exampleTransitGateway.ID(),
// 			AllowedPrefixes: pulumi.StringArray{
// 				pulumi.String("10.255.255.0/30"),
// 				pulumi.String("10.255.255.8/30"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Allowed Prefixes
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/directconnect"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleGateway, err := directconnect.NewGateway(ctx, "exampleGateway", &directconnect.GatewayArgs{
// 			AmazonSideAsn: pulumi.String("64512"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock: pulumi.String("10.255.255.0/28"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpnGateway, err := ec2.NewVpnGateway(ctx, "exampleVpnGateway", &ec2.VpnGatewayArgs{
// 			VpcId: exampleVpc.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = directconnect.NewGatewayAssociation(ctx, "exampleGatewayAssociation", &directconnect.GatewayAssociationArgs{
// 			DxGatewayId:         exampleGateway.ID(),
// 			AssociatedGatewayId: exampleVpnGateway.ID(),
// 			AllowedPrefixes: pulumi.StringArray{
// 				pulumi.String("210.52.109.0/24"),
// 				pulumi.String("175.45.176.0/22"),
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
// Direct Connect gateway associations can be imported using `dx_gateway_id` together with `associated_gateway_id`, e.g.
//
// ```sh
//  $ pulumi import aws:directconnect/gatewayAssociation:GatewayAssociation example dxgw-12345678/vgw-98765432
// ```
type GatewayAssociation struct {
	pulumi.CustomResourceState

	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayOutput `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringOutput `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringOutput `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringOutput `pulumi:"associatedGatewayType"`
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringOutput `pulumi:"dxGatewayAssociationId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringOutput `pulumi:"dxGatewayId"`
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringOutput `pulumi:"dxGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrOutput `pulumi:"proposalId"`
}

// NewGatewayAssociation registers a new resource with the given unique name, arguments, and options.
func NewGatewayAssociation(ctx *pulumi.Context,
	name string, args *GatewayAssociationArgs, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DxGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'DxGatewayId'")
	}
	var resource GatewayAssociation
	err := ctx.RegisterResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayAssociation gets an existing GatewayAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayAssociationState, opts ...pulumi.ResourceOption) (*GatewayAssociation, error) {
	var resource GatewayAssociation
	err := ctx.ReadResource("aws:directconnect/gatewayAssociation:GatewayAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayAssociation resources.
type gatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId *string `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId *string `pulumi:"associatedGatewayOwnerAccountId"`
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType *string `pulumi:"associatedGatewayType"`
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId *string `pulumi:"dxGatewayAssociationId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId *string `pulumi:"dxGatewayId"`
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId *string `pulumi:"dxGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId *string `pulumi:"proposalId"`
}

type GatewayAssociationState struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringPtrInput
	// The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
	AssociatedGatewayType pulumi.StringPtrInput
	// The ID of the Direct Connect gateway association.
	DxGatewayAssociationId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the Direct Connect gateway.
	DxGatewayOwnerAccountId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrInput
}

func (GatewayAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayAssociationState)(nil)).Elem()
}

type gatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `pulumi:"allowedPrefixes"`
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId *string `pulumi:"associatedGatewayId"`
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId *string `pulumi:"associatedGatewayOwnerAccountId"`
	// The ID of the Direct Connect gateway.
	DxGatewayId string `pulumi:"dxGatewayId"`
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId *string `pulumi:"proposalId"`
}

// The set of arguments for constructing a GatewayAssociation resource.
type GatewayAssociationArgs struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes pulumi.StringArrayInput
	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for single account Direct Connect gateway associations.
	AssociatedGatewayId pulumi.StringPtrInput
	// The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	// Used for cross-account Direct Connect gateway associations.
	AssociatedGatewayOwnerAccountId pulumi.StringPtrInput
	// The ID of the Direct Connect gateway.
	DxGatewayId pulumi.StringInput
	// The ID of the Direct Connect gateway association proposal.
	// Used for cross-account Direct Connect gateway associations.
	ProposalId pulumi.StringPtrInput
}

func (GatewayAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayAssociationArgs)(nil)).Elem()
}

type GatewayAssociationInput interface {
	pulumi.Input

	ToGatewayAssociationOutput() GatewayAssociationOutput
	ToGatewayAssociationOutputWithContext(ctx context.Context) GatewayAssociationOutput
}

func (*GatewayAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayAssociation)(nil))
}

func (i *GatewayAssociation) ToGatewayAssociationOutput() GatewayAssociationOutput {
	return i.ToGatewayAssociationOutputWithContext(context.Background())
}

func (i *GatewayAssociation) ToGatewayAssociationOutputWithContext(ctx context.Context) GatewayAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayAssociationOutput)
}

type GatewayAssociationOutput struct {
	*pulumi.OutputState
}

func (GatewayAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayAssociation)(nil))
}

func (o GatewayAssociationOutput) ToGatewayAssociationOutput() GatewayAssociationOutput {
	return o
}

func (o GatewayAssociationOutput) ToGatewayAssociationOutputWithContext(ctx context.Context) GatewayAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayAssociationOutput{})
}
