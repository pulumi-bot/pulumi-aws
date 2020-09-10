// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LocalGatewayRouteTableVpcAssociation struct {
	pulumi.CustomResourceState

	LocalGatewayId           pulumi.StringOutput    `pulumi:"localGatewayId"`
	LocalGatewayRouteTableId pulumi.StringOutput    `pulumi:"localGatewayRouteTableId"`
	Tags                     pulumi.StringMapOutput `pulumi:"tags"`
	VpcId                    pulumi.StringOutput    `pulumi:"vpcId"`
}

// NewLocalGatewayRouteTableVpcAssociation registers a new resource with the given unique name, arguments, and options.
func NewLocalGatewayRouteTableVpcAssociation(ctx *pulumi.Context,
	name string, args *LocalGatewayRouteTableVpcAssociationArgs, opts ...pulumi.ResourceOption) (*LocalGatewayRouteTableVpcAssociation, error) {
	if args == nil || args.LocalGatewayRouteTableId == nil {
		return nil, errors.New("missing required argument 'LocalGatewayRouteTableId'")
	}
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &LocalGatewayRouteTableVpcAssociationArgs{}
	}
	var resource LocalGatewayRouteTableVpcAssociation
	err := ctx.RegisterResource("aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGatewayRouteTableVpcAssociation gets an existing LocalGatewayRouteTableVpcAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGatewayRouteTableVpcAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGatewayRouteTableVpcAssociationState, opts ...pulumi.ResourceOption) (*LocalGatewayRouteTableVpcAssociation, error) {
	var resource LocalGatewayRouteTableVpcAssociation
	err := ctx.ReadResource("aws:ec2/localGatewayRouteTableVpcAssociation:LocalGatewayRouteTableVpcAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGatewayRouteTableVpcAssociation resources.
type localGatewayRouteTableVpcAssociationState struct {
	LocalGatewayId           *string           `pulumi:"localGatewayId"`
	LocalGatewayRouteTableId *string           `pulumi:"localGatewayRouteTableId"`
	Tags                     map[string]string `pulumi:"tags"`
	VpcId                    *string           `pulumi:"vpcId"`
}

type LocalGatewayRouteTableVpcAssociationState struct {
	LocalGatewayId           pulumi.StringPtrInput
	LocalGatewayRouteTableId pulumi.StringPtrInput
	Tags                     pulumi.StringMapInput
	VpcId                    pulumi.StringPtrInput
}

func (LocalGatewayRouteTableVpcAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteTableVpcAssociationState)(nil)).Elem()
}

type localGatewayRouteTableVpcAssociationArgs struct {
	LocalGatewayRouteTableId string            `pulumi:"localGatewayRouteTableId"`
	Tags                     map[string]string `pulumi:"tags"`
	VpcId                    string            `pulumi:"vpcId"`
}

// The set of arguments for constructing a LocalGatewayRouteTableVpcAssociation resource.
type LocalGatewayRouteTableVpcAssociationArgs struct {
	LocalGatewayRouteTableId pulumi.StringInput
	Tags                     pulumi.StringMapInput
	VpcId                    pulumi.StringInput
}

func (LocalGatewayRouteTableVpcAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGatewayRouteTableVpcAssociationArgs)(nil)).Elem()
}
