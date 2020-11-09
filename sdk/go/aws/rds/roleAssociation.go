// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an RDS DB Instance association with an IAM Role. Example use cases:
//
// * [Amazon RDS Oracle integration with Amazon S3](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-s3-integration.html)
// * [Importing Amazon S3 Data into an RDS PostgreSQL DB Instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_PostgreSQL.S3Import.html)
//
// > To manage the RDS DB Instance IAM Role for [Enhanced Monitoring](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html), see the `rds.Instance` resource `monitoringRoleArn` argument instead.
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
// 		_, err := rds.NewRoleAssociation(ctx, "example", &rds.RoleAssociationArgs{
// 			DbInstanceIdentifier: pulumi.Any(aws_db_instance.Example.Id),
// 			FeatureName:          pulumi.String("S3_INTEGRATION"),
// 			RoleArn:              pulumi.Any(aws_iam_role.Example.Arn),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RoleAssociation struct {
	pulumi.CustomResourceState

	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier pulumi.StringOutput `pulumi:"dbInstanceIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringOutput `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
}

// NewRoleAssociation registers a new resource with the given unique name, arguments, and options.
func NewRoleAssociation(ctx *pulumi.Context,
	name string, args *RoleAssociationArgs, opts ...pulumi.ResourceOption) (*RoleAssociation, error) {
	if args == nil || args.DbInstanceIdentifier == nil {
		return nil, errors.New("missing required argument 'DbInstanceIdentifier'")
	}
	if args == nil || args.FeatureName == nil {
		return nil, errors.New("missing required argument 'FeatureName'")
	}
	if args == nil || args.RoleArn == nil {
		return nil, errors.New("missing required argument 'RoleArn'")
	}
	if args == nil {
		args = &RoleAssociationArgs{}
	}
	var resource RoleAssociation
	err := ctx.RegisterResource("aws:rds/roleAssociation:RoleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleAssociation gets an existing RoleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleAssociationState, opts ...pulumi.ResourceOption) (*RoleAssociation, error) {
	var resource RoleAssociation
	err := ctx.ReadResource("aws:rds/roleAssociation:RoleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleAssociation resources.
type roleAssociationState struct {
	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier *string `pulumi:"dbInstanceIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName *string `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn *string `pulumi:"roleArn"`
}

type RoleAssociationState struct {
	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier pulumi.StringPtrInput
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringPtrInput
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn pulumi.StringPtrInput
}

func (RoleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssociationState)(nil)).Elem()
}

type roleAssociationArgs struct {
	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier string `pulumi:"dbInstanceIdentifier"`
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName string `pulumi:"featureName"`
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn string `pulumi:"roleArn"`
}

// The set of arguments for constructing a RoleAssociation resource.
type RoleAssociationArgs struct {
	// DB Instance Identifier to associate with the IAM Role.
	DbInstanceIdentifier pulumi.StringInput
	// Name of the feature for association. This can be found in the AWS documentation relevant to the integration or a full list is available in the `SupportedFeatureNames` list returned by [AWS CLI rds describe-db-engine-versions](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-engine-versions.html).
	FeatureName pulumi.StringInput
	// Amazon Resource Name (ARN) of the IAM Role to associate with the DB Instance.
	RoleArn pulumi.StringInput
}

func (RoleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleAssociationArgs)(nil)).Elem()
}

type RoleAssociationInput interface {
	pulumi.Input

	ToRoleAssociationOutput() RoleAssociationOutput
	ToRoleAssociationOutputWithContext(ctx context.Context) RoleAssociationOutput
}

func (RoleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssociation)(nil)).Elem()
}

func (i RoleAssociation) ToRoleAssociationOutput() RoleAssociationOutput {
	return i.ToRoleAssociationOutputWithContext(context.Background())
}

func (i RoleAssociation) ToRoleAssociationOutputWithContext(ctx context.Context) RoleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssociationOutput)
}

type RoleAssociationOutput struct {
	*pulumi.OutputState
}

func (RoleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssociationOutput)(nil)).Elem()
}

func (o RoleAssociationOutput) ToRoleAssociationOutput() RoleAssociationOutput {
	return o
}

func (o RoleAssociationOutput) ToRoleAssociationOutputWithContext(ctx context.Context) RoleAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoleAssociationOutput{})
}
