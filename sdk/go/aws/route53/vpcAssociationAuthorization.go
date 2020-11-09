// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Authorizes a VPC in a peer account to be associated with a local Route53 Hosted Zone.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/providers"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/route53"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "alternate", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.6.0.0/16"),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleZone, err := route53.NewZone(ctx, "exampleZone", &route53.ZoneArgs{
// 			Vpcs: route53.ZoneVpcArray{
// 				&route53.ZoneVpcArgs{
// 					VpcId: exampleVpc.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		alternateVpc, err := ec2.NewVpc(ctx, "alternateVpc", &ec2.VpcArgs{
// 			CidrBlock:          pulumi.String("10.7.0.0/16"),
// 			EnableDnsHostnames: pulumi.Bool(true),
// 			EnableDnsSupport:   pulumi.Bool(true),
// 		}, pulumi.Provider("aws.alternate"))
// 		if err != nil {
// 			return err
// 		}
// 		exampleVpcAssociationAuthorization, err := route53.NewVpcAssociationAuthorization(ctx, "exampleVpcAssociationAuthorization", &route53.VpcAssociationAuthorizationArgs{
// 			VpcId:  alternateVpc.ID(),
// 			ZoneId: exampleZone.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = route53.NewZoneAssociation(ctx, "exampleZoneAssociation", &route53.ZoneAssociationArgs{
// 			VpcId:  exampleVpcAssociationAuthorization.VpcId,
// 			ZoneId: exampleVpcAssociationAuthorization.ZoneId,
// 		}, pulumi.Provider("aws.alternate"))
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
// Route 53 VPC Association Authorizations can be imported via the Hosted Zone ID and VPC ID, separated by a colon (`:`), e.g.
//
// ```sh
//  $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
// ```
type VpcAssociationAuthorization struct {
	pulumi.CustomResourceState

	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringOutput `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewVpcAssociationAuthorization registers a new resource with the given unique name, arguments, and options.
func NewVpcAssociationAuthorization(ctx *pulumi.Context,
	name string, args *VpcAssociationAuthorizationArgs, opts ...pulumi.ResourceOption) (*VpcAssociationAuthorization, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil || args.ZoneId == nil {
		return nil, errors.New("missing required argument 'ZoneId'")
	}
	if args == nil {
		args = &VpcAssociationAuthorizationArgs{}
	}
	var resource VpcAssociationAuthorization
	err := ctx.RegisterResource("aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcAssociationAuthorization gets an existing VpcAssociationAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcAssociationAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcAssociationAuthorizationState, opts ...pulumi.ResourceOption) (*VpcAssociationAuthorization, error) {
	var resource VpcAssociationAuthorization
	err := ctx.ReadResource("aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcAssociationAuthorization resources.
type vpcAssociationAuthorizationState struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId *string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId *string `pulumi:"zoneId"`
}

type VpcAssociationAuthorizationState struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringPtrInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringPtrInput
}

func (VpcAssociationAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAssociationAuthorizationState)(nil)).Elem()
}

type vpcAssociationAuthorizationArgs struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId string `pulumi:"vpcId"`
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion *string `pulumi:"vpcRegion"`
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a VpcAssociationAuthorization resource.
type VpcAssociationAuthorizationArgs struct {
	// The VPC to authorize for association with the private hosted zone.
	VpcId pulumi.StringInput
	// The VPC's region. Defaults to the region of the AWS provider.
	VpcRegion pulumi.StringPtrInput
	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneId pulumi.StringInput
}

func (VpcAssociationAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcAssociationAuthorizationArgs)(nil)).Elem()
}
