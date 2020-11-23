// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an RDS security group resource. This is only for DB instances in the
// EC2-Classic Platform. For instances inside a VPC, use the
// `aws_db_instance.vpc_security_group_ids`
// attribute instead.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewSecurityGroup(ctx, "_default", &rds.SecurityGroupArgs{
// 			Ingress: rds.SecurityGroupIngressArray{
// 				&rds.SecurityGroupIngressArgs{
// 					Cidr: pulumi.String("10.0.0.0/24"),
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
// DB Security groups can be imported using the `name`, e.g.
//
// ```sh
//  $ pulumi import aws:rds/securityGroup:SecurityGroup default aws_rds_sg-1
// ```
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The arn of the DB security group.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The description of the DB security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// The name of the DB security group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ingress == nil {
		return nil, errors.New("invalid value for required argument 'Ingress'")
	}
	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:rds/securityGroup:SecurityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroup gets an existing SecurityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupState, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	var resource SecurityGroup
	err := ctx.ReadResource("aws:rds/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The arn of the DB security group.
	Arn *string `pulumi:"arn"`
	// The description of the DB security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the DB security group.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type SecurityGroupState struct {
	// The arn of the DB security group.
	Arn pulumi.StringPtrInput
	// The description of the DB security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the DB security group.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The description of the DB security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the DB security group.
	Name *string `pulumi:"name"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The description of the DB security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the DB security group.
	Name pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (SecurityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupArgs)(nil)).Elem()
}

type SecurityGroupInput interface {
	pulumi.Input

	ToSecurityGroupOutput() SecurityGroupOutput
	ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput
}

func (SecurityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroup)(nil)).Elem()
}

func (i SecurityGroup) ToSecurityGroupOutput() SecurityGroupOutput {
	return i.ToSecurityGroupOutputWithContext(context.Background())
}

func (i SecurityGroup) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupOutput)
}

type SecurityGroupOutput struct {
	*pulumi.OutputState
}

func (SecurityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityGroupOutput)(nil)).Elem()
}

func (o SecurityGroupOutput) ToSecurityGroupOutput() SecurityGroupOutput {
	return o
}

func (o SecurityGroupOutput) ToSecurityGroupOutputWithContext(ctx context.Context) SecurityGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecurityGroupOutput{})
}
