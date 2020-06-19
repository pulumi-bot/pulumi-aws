// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an API Gateway VPC Link.
//
// > **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
// To enable private integration for HTTP APIs, use the `Amazon API Gateway Version 2 VPC Link` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/apigateway"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleLoadBalancer, err := lb.NewLoadBalancer(ctx, "exampleLoadBalancer", &lb.LoadBalancerArgs{
// 			Internal:         pulumi.Bool(true),
// 			LoadBalancerType: pulumi.String("network"),
// 			SubnetMappings: lb.LoadBalancerSubnetMappingArray{
// 				&lb.LoadBalancerSubnetMappingArgs{
// 					SubnetId: pulumi.String("12345"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = apigateway.NewVpcLink(ctx, "exampleVpcLink", &apigateway.VpcLinkArgs{
// 			Description: pulumi.String("example description"),
// 			TargetArn:   exampleLoadBalancer.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type VpcLink struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the VPC link.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name pulumi.StringOutput `pulumi:"name"`
	// Key-value map of resource tags
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringOutput `pulumi:"targetArn"`
}

// NewVpcLink registers a new resource with the given unique name, arguments, and options.
func NewVpcLink(ctx *pulumi.Context,
	name string, args *VpcLinkArgs, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	if args == nil || args.TargetArn == nil {
		return nil, errors.New("missing required argument 'TargetArn'")
	}
	if args == nil {
		args = &VpcLinkArgs{}
	}
	var resource VpcLink
	err := ctx.RegisterResource("aws:apigateway/vpcLink:VpcLink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcLink gets an existing VpcLink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcLinkState, opts ...pulumi.ResourceOption) (*VpcLink, error) {
	var resource VpcLink
	err := ctx.ReadResource("aws:apigateway/vpcLink:VpcLink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcLink resources.
type vpcLinkState struct {
	Arn *string `pulumi:"arn"`
	// The description of the VPC link.
	Description *string `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn *string `pulumi:"targetArn"`
}

type VpcLinkState struct {
	Arn pulumi.StringPtrInput
	// The description of the VPC link.
	Description pulumi.StringPtrInput
	// The name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringPtrInput
}

func (VpcLinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkState)(nil)).Elem()
}

type vpcLinkArgs struct {
	// The description of the VPC link.
	Description *string `pulumi:"description"`
	// The name used to label and identify the VPC link.
	Name *string `pulumi:"name"`
	// Key-value map of resource tags
	Tags map[string]interface{} `pulumi:"tags"`
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn string `pulumi:"targetArn"`
}

// The set of arguments for constructing a VpcLink resource.
type VpcLinkArgs struct {
	// The description of the VPC link.
	Description pulumi.StringPtrInput
	// The name used to label and identify the VPC link.
	Name pulumi.StringPtrInput
	// Key-value map of resource tags
	Tags pulumi.MapInput
	// The list of network load balancer arns in the VPC targeted by the VPC link. Currently AWS only supports 1 target.
	TargetArn pulumi.StringInput
}

func (VpcLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcLinkArgs)(nil)).Elem()
}
