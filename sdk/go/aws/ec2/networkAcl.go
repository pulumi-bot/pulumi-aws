// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an network ACL resource. You might set up network ACLs with rules similar
// to your security groups in order to add an additional layer of security to your VPC.
//
// > **NOTE on Network ACLs and Network ACL Rules:** This provider currently
// provides both a standalone Network ACL Rule resource and a Network ACL resource with rules
// defined in-line. At this time you cannot use a Network ACL with in-line rules
// in conjunction with any Network ACL Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = ec2.NewNetworkAcl(ctx, "main", &ec2.NetworkAclArgs{
// 			Egress: ec2.NetworkAclEgressArray{
// 				&ec2.NetworkAclEgressArgs{
// 					Action:    pulumi.String("allow"),
// 					CidrBlock: pulumi.String("10.3.0.0/18"),
// 					FromPort:  pulumi.Int(443),
// 					Protocol:  pulumi.String("tcp"),
// 					RuleNo:    pulumi.Int(200),
// 					ToPort:    pulumi.Int(443),
// 				},
// 			},
// 			Ingress: ec2.NetworkAclIngressArray{
// 				&ec2.NetworkAclIngressArgs{
// 					Action:    pulumi.String("allow"),
// 					CidrBlock: pulumi.String("10.3.0.0/18"),
// 					FromPort:  pulumi.Int(80),
// 					Protocol:  pulumi.String("tcp"),
// 					RuleNo:    pulumi.Int(100),
// 					ToPort:    pulumi.Int(80),
// 				},
// 			},
// 			Tags: map[string]interface{}{
// 				"Name": "main",
// 			},
// 			VpcId: pulumi.String(aws_vpc.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type NetworkAcl struct {
	pulumi.CustomResourceState

	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayOutput `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayOutput `pulumi:"ingress"`
	// The ID of the AWS account that owns the network ACL.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// The ID of the associated VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	if args == nil || args.VpcId == nil {
		return nil, errors.New("missing required argument 'VpcId'")
	}
	if args == nil {
		args = &NetworkAclArgs{}
	}
	var resource NetworkAcl
	err := ctx.RegisterResource("aws:ec2/networkAcl:NetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclState, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	var resource NetworkAcl
	err := ctx.ReadResource("aws:ec2/networkAcl:NetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAcl resources.
type networkAclState struct {
	// Specifies an egress rule. Parameters defined below.
	Egress []NetworkAclEgress `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []NetworkAclIngress `pulumi:"ingress"`
	// The ID of the AWS account that owns the network ACL.
	OwnerId *string `pulumi:"ownerId"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the associated VPC.
	VpcId *string `pulumi:"vpcId"`
}

type NetworkAclState struct {
	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayInput
	// The ID of the AWS account that owns the network ACL.
	OwnerId pulumi.StringPtrInput
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the associated VPC.
	VpcId pulumi.StringPtrInput
}

func (NetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclState)(nil)).Elem()
}

type networkAclArgs struct {
	// Specifies an egress rule. Parameters defined below.
	Egress []NetworkAclEgress `pulumi:"egress"`
	// Specifies an ingress rule. Parameters defined below.
	Ingress []NetworkAclIngress `pulumi:"ingress"`
	// A list of Subnet IDs to apply the ACL to
	SubnetIds []string `pulumi:"subnetIds"`
	// A mapping of tags to assign to the resource.
	Tags map[string]interface{} `pulumi:"tags"`
	// The ID of the associated VPC.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	// Specifies an egress rule. Parameters defined below.
	Egress NetworkAclEgressArrayInput
	// Specifies an ingress rule. Parameters defined below.
	Ingress NetworkAclIngressArrayInput
	// A list of Subnet IDs to apply the ACL to
	SubnetIds pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.MapInput
	// The ID of the associated VPC.
	VpcId pulumi.StringInput
}

func (NetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclArgs)(nil)).Elem()
}
