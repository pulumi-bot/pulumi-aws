// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a security group resource.
//
// > **NOTE on Security Groups and Security Group Rules:** This provider currently
// provides both a standalone Security Group Rule resource (a single `ingress` or
// `egress` rule), and a Security Group resource with `ingress` and `egress` rules
// defined in-line. At this time you cannot use a Security Group with in-line rules
// in conjunction with any Security Group Rule resources. Doing so will cause
// a conflict of rule settings and will overwrite rules.
//
// > **NOTE:** Referencing Security Groups across VPC peering has certain restrictions. More information is available in the [VPC Peering User Guide](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html).
//
// > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), security groups associated with Lambda Functions can take up to 45 minutes to successfully delete.
//
// ## Usage with prefix list IDs
//
// Prefix list IDs are managed by AWS internally. Prefix list IDs
// are associated with a prefix list name, or service name, that is linked to a specific region.
// Prefix list IDs are exported on VPC Endpoints, so you can use this format:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ec2"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myEndpoint, err := ec2.NewVpcEndpoint(ctx, "myEndpoint", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ec2.NewSecurityGroup(ctx, "example", &ec2.SecurityGroupArgs{
// 			Egress: ec2.SecurityGroupEgressArray{
// 				&ec2.SecurityGroupEgressArgs{
// 					FromPort: pulumi.Int(0),
// 					ToPort:   pulumi.Int(0),
// 					Protocol: pulumi.String("-1"),
// 					PrefixListIds: pulumi.StringArray{
// 						myEndpoint.PrefixListId,
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
type SecurityGroup struct {
	pulumi.CustomResourceState

	// The ARN of the security group
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of this egress rule.
	Description pulumi.StringOutput `pulumi:"description"`
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress SecurityGroupEgressArrayOutput `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress SecurityGroupIngressArrayOutput `pulumi:"ingress"`
	// The name of the security group. If omitted, this provider will
	// assign a random, unique name
	Name pulumi.StringOutput `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrOutput `pulumi:"namePrefix"`
	// The owner ID.
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// Instruct this provider to revoke all of the
	// Security Groups attached ingress and egress rules before deleting the rule
	// itself. This is normally not needed, however certain AWS services such as
	// Elastic Map Reduce may automatically add required rules to security groups used
	// with the service, and those rules may contain a cyclic dependency that prevent
	// the security groups from being destroyed without removing the dependency first.
	// Default `false`
	RevokeRulesOnDelete pulumi.BoolPtrOutput `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The VPC ID.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewSecurityGroup registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroup(ctx *pulumi.Context,
	name string, args *SecurityGroupArgs, opts ...pulumi.ResourceOption) (*SecurityGroup, error) {
	if args == nil {
		args = &SecurityGroupArgs{}
	}

	if args.Description == nil {
		args.Description = pulumi.StringPtr("Managed by Pulumi")
	}
	var resource SecurityGroup
	err := ctx.RegisterResource("aws:ec2/securityGroup:SecurityGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws:ec2/securityGroup:SecurityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroup resources.
type securityGroupState struct {
	// The ARN of the security group
	Arn *string `pulumi:"arn"`
	// Description of this egress rule.
	Description *string `pulumi:"description"`
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress []SecurityGroupEgress `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the security group. If omitted, this provider will
	// assign a random, unique name
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// The owner ID.
	OwnerId *string `pulumi:"ownerId"`
	// Instruct this provider to revoke all of the
	// Security Groups attached ingress and egress rules before deleting the rule
	// itself. This is normally not needed, however certain AWS services such as
	// Elastic Map Reduce may automatically add required rules to security groups used
	// with the service, and those rules may contain a cyclic dependency that prevent
	// the security groups from being destroyed without removing the dependency first.
	// Default `false`
	RevokeRulesOnDelete *bool `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

type SecurityGroupState struct {
	// The ARN of the security group
	Arn pulumi.StringPtrInput
	// Description of this egress rule.
	Description pulumi.StringPtrInput
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress SecurityGroupEgressArrayInput
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress SecurityGroupIngressArrayInput
	// The name of the security group. If omitted, this provider will
	// assign a random, unique name
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// The owner ID.
	OwnerId pulumi.StringPtrInput
	// Instruct this provider to revoke all of the
	// Security Groups attached ingress and egress rules before deleting the rule
	// itself. This is normally not needed, however certain AWS services such as
	// Elastic Map Reduce may automatically add required rules to security groups used
	// with the service, and those rules may contain a cyclic dependency that prevent
	// the security groups from being destroyed without removing the dependency first.
	// Default `false`
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
}

func (SecurityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupState)(nil)).Elem()
}

type securityGroupArgs struct {
	// Description of this egress rule.
	Description *string `pulumi:"description"`
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress []SecurityGroupEgress `pulumi:"egress"`
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress []SecurityGroupIngress `pulumi:"ingress"`
	// The name of the security group. If omitted, this provider will
	// assign a random, unique name
	Name *string `pulumi:"name"`
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix *string `pulumi:"namePrefix"`
	// Instruct this provider to revoke all of the
	// Security Groups attached ingress and egress rules before deleting the rule
	// itself. This is normally not needed, however certain AWS services such as
	// Elastic Map Reduce may automatically add required rules to security groups used
	// with the service, and those rules may contain a cyclic dependency that prevent
	// the security groups from being destroyed without removing the dependency first.
	// Default `false`
	RevokeRulesOnDelete *bool `pulumi:"revokeRulesOnDelete"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a SecurityGroup resource.
type SecurityGroupArgs struct {
	// Description of this egress rule.
	Description pulumi.StringPtrInput
	// Can be specified multiple times for each
	// egress rule. Each egress block supports fields documented below.
	Egress SecurityGroupEgressArrayInput
	// Can be specified multiple times for each
	// ingress rule. Each ingress block supports fields documented below.
	Ingress SecurityGroupIngressArrayInput
	// The name of the security group. If omitted, this provider will
	// assign a random, unique name
	Name pulumi.StringPtrInput
	// Creates a unique name beginning with the specified
	// prefix. Conflicts with `name`.
	NamePrefix pulumi.StringPtrInput
	// Instruct this provider to revoke all of the
	// Security Groups attached ingress and egress rules before deleting the rule
	// itself. This is normally not needed, however certain AWS services such as
	// Elastic Map Reduce may automatically add required rules to security groups used
	// with the service, and those rules may contain a cyclic dependency that prevent
	// the security groups from being destroyed without removing the dependency first.
	// Default `false`
	RevokeRulesOnDelete pulumi.BoolPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The VPC ID.
	VpcId pulumi.StringPtrInput
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
