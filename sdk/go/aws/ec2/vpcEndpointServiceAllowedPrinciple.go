// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to allow a principal to discover a VPC endpoint service.
//
// > **NOTE on VPC Endpoint Services and VPC Endpoint Service Allowed Principals:** This provider provides
// both a standalone VPC Endpoint Service Allowed Principal resource
// and a VPC Endpoint Service resource with an `allowedPrincipals` attribute. Do not use the same principal ARN in both
// a VPC Endpoint Service resource and a VPC Endpoint Service Allowed Principal resource. Doing so will cause a conflict
// and will overwrite the association.
//
// ## Example Usage
type VpcEndpointServiceAllowedPrinciple struct {
	pulumi.CustomResourceState

	// The ARN of the principal to allow permissions.
	PrincipalArn pulumi.StringOutput `pulumi:"principalArn"`
	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId pulumi.StringOutput `pulumi:"vpcEndpointServiceId"`
}

// NewVpcEndpointServiceAllowedPrinciple registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServiceAllowedPrinciple(ctx *pulumi.Context,
	name string, args *VpcEndpointServiceAllowedPrincipleArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServiceAllowedPrinciple, error) {
	if args == nil || args.PrincipalArn == nil {
		return nil, errors.New("missing required argument 'PrincipalArn'")
	}
	if args == nil || args.VpcEndpointServiceId == nil {
		return nil, errors.New("missing required argument 'VpcEndpointServiceId'")
	}
	if args == nil {
		args = &VpcEndpointServiceAllowedPrincipleArgs{}
	}
	var resource VpcEndpointServiceAllowedPrinciple
	err := ctx.RegisterResource("aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServiceAllowedPrinciple gets an existing VpcEndpointServiceAllowedPrinciple resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServiceAllowedPrinciple(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServiceAllowedPrincipleState, opts ...pulumi.ResourceOption) (*VpcEndpointServiceAllowedPrinciple, error) {
	var resource VpcEndpointServiceAllowedPrinciple
	err := ctx.ReadResource("aws:ec2/vpcEndpointServiceAllowedPrinciple:VpcEndpointServiceAllowedPrinciple", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServiceAllowedPrinciple resources.
type vpcEndpointServiceAllowedPrincipleState struct {
	// The ARN of the principal to allow permissions.
	PrincipalArn *string `pulumi:"principalArn"`
	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId *string `pulumi:"vpcEndpointServiceId"`
}

type VpcEndpointServiceAllowedPrincipleState struct {
	// The ARN of the principal to allow permissions.
	PrincipalArn pulumi.StringPtrInput
	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId pulumi.StringPtrInput
}

func (VpcEndpointServiceAllowedPrincipleState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceAllowedPrincipleState)(nil)).Elem()
}

type vpcEndpointServiceAllowedPrincipleArgs struct {
	// The ARN of the principal to allow permissions.
	PrincipalArn string `pulumi:"principalArn"`
	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId string `pulumi:"vpcEndpointServiceId"`
}

// The set of arguments for constructing a VpcEndpointServiceAllowedPrinciple resource.
type VpcEndpointServiceAllowedPrincipleArgs struct {
	// The ARN of the principal to allow permissions.
	PrincipalArn pulumi.StringInput
	// The ID of the VPC endpoint service to allow permission.
	VpcEndpointServiceId pulumi.StringInput
}

func (VpcEndpointServiceAllowedPrincipleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServiceAllowedPrincipleArgs)(nil)).Elem()
}
