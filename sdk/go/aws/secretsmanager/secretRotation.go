// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage AWS Secrets Manager secret rotation. To manage a secret, see the [`secretsmanager.Secret` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret.html). To manage a secret value, see the [`secretsmanager.SecretVersion` resource](https://www.terraform.io/docs/providers/aws/r/secretsmanager_secret_version.html).
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/secretsmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := secretsmanager.NewSecretRotation(ctx, "example", &secretsmanager.SecretRotationArgs{
// 			SecretId:          pulumi.Any(aws_secretsmanager_secret.Example.Id),
// 			RotationLambdaArn: pulumi.Any(aws_lambda_function.Example.Arn),
// 			RotationRules: &secretsmanager.SecretRotationRotationRulesArgs{
// 				AutomaticallyAfterDays: pulumi.Int(30),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Rotation Configuration
//
// To enable automatic secret rotation, the Secrets Manager service requires usage of a Lambda function. The [Rotate Secrets section in the Secrets Manager User Guide](https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotating-secrets.html) provides additional information about deploying a prebuilt Lambda functions for supported credential rotation (e.g. RDS) or deploying a custom Lambda function.
//
// > **NOTE:** Configuring rotation causes the secret to rotate once as soon as you enable rotation. Before you do this, you must ensure that all of your applications that use the credentials stored in the secret are updated to retrieve the secret from AWS Secrets Manager. The old credentials might no longer be usable after the initial rotation and any applications that you fail to update will break as soon as the old credentials are no longer valid.
//
// > **NOTE:** If you cancel a rotation that is in progress (by removing the `rotation` configuration), it can leave the VersionStage labels in an unexpected state. Depending on what step of the rotation was in progress, you might need to remove the staging label AWSPENDING from the partially created version, specified by the SecretVersionId response value. You should also evaluate the partially rotated new version to see if it should be deleted, which you can do by removing all staging labels from the new version's VersionStage field.
//
// ## Import
//
// `aws_secretsmanager_secret_rotation` can be imported by using the secret Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:secretsmanager/secretRotation:SecretRotation example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
// ```
type SecretRotation struct {
	pulumi.CustomResourceState

	// Specifies whether automatic rotation is enabled for this secret.
	RotationEnabled pulumi.BoolOutput `pulumi:"rotationEnabled"`
	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn pulumi.StringOutput `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules SecretRotationRotationRulesOutput `pulumi:"rotationRules"`
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringOutput    `pulumi:"secretId"`
	Tags     pulumi.StringMapOutput `pulumi:"tags"`
}

// NewSecretRotation registers a new resource with the given unique name, arguments, and options.
func NewSecretRotation(ctx *pulumi.Context,
	name string, args *SecretRotationArgs, opts ...pulumi.ResourceOption) (*SecretRotation, error) {
	if args == nil || args.RotationLambdaArn == nil {
		return nil, errors.New("missing required argument 'RotationLambdaArn'")
	}
	if args == nil || args.RotationRules == nil {
		return nil, errors.New("missing required argument 'RotationRules'")
	}
	if args == nil || args.SecretId == nil {
		return nil, errors.New("missing required argument 'SecretId'")
	}
	if args == nil {
		args = &SecretRotationArgs{}
	}
	var resource SecretRotation
	err := ctx.RegisterResource("aws:secretsmanager/secretRotation:SecretRotation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretRotation gets an existing SecretRotation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretRotation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretRotationState, opts ...pulumi.ResourceOption) (*SecretRotation, error) {
	var resource SecretRotation
	err := ctx.ReadResource("aws:secretsmanager/secretRotation:SecretRotation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretRotation resources.
type secretRotationState struct {
	// Specifies whether automatic rotation is enabled for this secret.
	RotationEnabled *bool `pulumi:"rotationEnabled"`
	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn *string `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules *SecretRotationRotationRules `pulumi:"rotationRules"`
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId *string           `pulumi:"secretId"`
	Tags     map[string]string `pulumi:"tags"`
}

type SecretRotationState struct {
	// Specifies whether automatic rotation is enabled for this secret.
	RotationEnabled pulumi.BoolPtrInput
	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn pulumi.StringPtrInput
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules SecretRotationRotationRulesPtrInput
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringPtrInput
	Tags     pulumi.StringMapInput
}

func (SecretRotationState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRotationState)(nil)).Elem()
}

type secretRotationArgs struct {
	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn string `pulumi:"rotationLambdaArn"`
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules SecretRotationRotationRules `pulumi:"rotationRules"`
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId string            `pulumi:"secretId"`
	Tags     map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SecretRotation resource.
type SecretRotationArgs struct {
	// Specifies the ARN of the Lambda function that can rotate the secret.
	RotationLambdaArn pulumi.StringInput
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules SecretRotationRotationRulesInput
	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId pulumi.StringInput
	Tags     pulumi.StringMapInput
}

func (SecretRotationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretRotationArgs)(nil)).Elem()
}

type SecretRotationInput interface {
	pulumi.Input

	ToSecretRotationOutput() SecretRotationOutput
	ToSecretRotationOutputWithContext(ctx context.Context) SecretRotationOutput
}

func (SecretRotation) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRotation)(nil)).Elem()
}

func (i SecretRotation) ToSecretRotationOutput() SecretRotationOutput {
	return i.ToSecretRotationOutputWithContext(context.Background())
}

func (i SecretRotation) ToSecretRotationOutputWithContext(ctx context.Context) SecretRotationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretRotationOutput)
}

type SecretRotationOutput struct {
	*pulumi.OutputState
}

func (SecretRotationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretRotationOutput)(nil)).Elem()
}

func (o SecretRotationOutput) ToSecretRotationOutput() SecretRotationOutput {
	return o
}

func (o SecretRotationOutput) ToSecretRotationOutputWithContext(ctx context.Context) SecretRotationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretRotationOutput{})
}
