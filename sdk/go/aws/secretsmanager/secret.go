// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Secret struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput    `pulumi:"arn"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	KmsKeyId             pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	NamePrefix           pulumi.StringOutput    `pulumi:"namePrefix"`
	Policy               pulumi.StringPtrOutput `pulumi:"policy"`
	RecoveryWindowInDays pulumi.IntPtrOutput    `pulumi:"recoveryWindowInDays"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationEnabled pulumi.BoolOutput `pulumi:"rotationEnabled"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationLambdaArn pulumi.StringOutput `pulumi:"rotationLambdaArn"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationRules SecretRotationRulesOutput `pulumi:"rotationRules"`
	Tags          pulumi.StringMapOutput    `pulumi:"tags"`
}

// NewSecret registers a new resource with the given unique name, arguments, and options.
func NewSecret(ctx *pulumi.Context,
	name string, args *SecretArgs, opts ...pulumi.ResourceOption) (*Secret, error) {
	if args == nil {
		args = &SecretArgs{}
	}
	var resource Secret
	err := ctx.RegisterResource("aws:secretsmanager/secret:Secret", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecret gets an existing Secret resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecret(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretState, opts ...pulumi.ResourceOption) (*Secret, error) {
	var resource Secret
	err := ctx.ReadResource("aws:secretsmanager/secret:Secret", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Secret resources.
type secretState struct {
	Arn                  *string `pulumi:"arn"`
	Description          *string `pulumi:"description"`
	KmsKeyId             *string `pulumi:"kmsKeyId"`
	Name                 *string `pulumi:"name"`
	NamePrefix           *string `pulumi:"namePrefix"`
	Policy               *string `pulumi:"policy"`
	RecoveryWindowInDays *int    `pulumi:"recoveryWindowInDays"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationEnabled *bool `pulumi:"rotationEnabled"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationLambdaArn *string `pulumi:"rotationLambdaArn"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationRules *SecretRotationRules `pulumi:"rotationRules"`
	Tags          map[string]string    `pulumi:"tags"`
}

type SecretState struct {
	Arn                  pulumi.StringPtrInput
	Description          pulumi.StringPtrInput
	KmsKeyId             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NamePrefix           pulumi.StringPtrInput
	Policy               pulumi.StringPtrInput
	RecoveryWindowInDays pulumi.IntPtrInput
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationEnabled pulumi.BoolPtrInput
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationLambdaArn pulumi.StringPtrInput
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationRules SecretRotationRulesPtrInput
	Tags          pulumi.StringMapInput
}

func (SecretState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretState)(nil)).Elem()
}

type secretArgs struct {
	Description          *string `pulumi:"description"`
	KmsKeyId             *string `pulumi:"kmsKeyId"`
	Name                 *string `pulumi:"name"`
	NamePrefix           *string `pulumi:"namePrefix"`
	Policy               *string `pulumi:"policy"`
	RecoveryWindowInDays *int    `pulumi:"recoveryWindowInDays"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationLambdaArn *string `pulumi:"rotationLambdaArn"`
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationRules *SecretRotationRules `pulumi:"rotationRules"`
	Tags          map[string]string    `pulumi:"tags"`
}

// The set of arguments for constructing a Secret resource.
type SecretArgs struct {
	Description          pulumi.StringPtrInput
	KmsKeyId             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NamePrefix           pulumi.StringPtrInput
	Policy               pulumi.StringPtrInput
	RecoveryWindowInDays pulumi.IntPtrInput
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationLambdaArn pulumi.StringPtrInput
	// Deprecated: Use the aws_secretsmanager_secret_rotation resource instead
	RotationRules SecretRotationRulesPtrInput
	Tags          pulumi.StringMapInput
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretArgs)(nil)).Elem()
}
