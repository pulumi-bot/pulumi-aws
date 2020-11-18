// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Creates a new Amazon Redshift security group. You use security groups to control access to non-VPC clusters
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/redshift"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := redshift.NewSecurityGroup(ctx, "_default", &redshift.SecurityGroupArgs{
// 			Ingress: redshift.SecurityGroupIngressArray{
// 				&redshift.SecurityGroupIngressArgs{
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
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringOutput `pulumi:"description"`
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name pulumi.StringOutput `pulumi:"name"`
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
	err := ctx.RegisterResource("aws:redshift/securityGroup:SecurityGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:redshift/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name *string `pulumi:"name"`
}

type SecurityGroupState struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the Redshift security group.
	Name pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description *string `pulumi:"description"`
	// A list of ingress rules.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the Redshift security group.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// The description of the Redshift security group. Defaults to "Managed by Pulumi".
	Description pulumi.StringPtrInput
	// A list of ingress rules.
	Ingress SecurityGroupIngressArrayInput
	// The name of the Redshift security group.
	Name pulumi.StringPtrInput
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
