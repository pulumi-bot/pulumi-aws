// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package secretsmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage AWS Secrets Manager secret policy.
//
// ## Example Usage
// ### Basic
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/secretsmanager"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleSecret, err := secretsmanager.NewSecret(ctx, "exampleSecret", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = secretsmanager.NewSecretPolicy(ctx, "exampleSecretPolicy", &secretsmanager.SecretPolicyArgs{
// 			SecretArn: exampleSecret.Arn,
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "	{\n", "	  \"Sid\": \"EnableAllPermissions\",\n", "	  \"Effect\": \"Allow\",\n", "	  \"Principal\": {\n", "		\"AWS\": \"*\"\n", "	  },\n", "	  \"Action\": \"secretsmanager:GetSecretValue\",\n", "	  \"Resource\": \"*\"\n", "	}\n", "  ]\n", "}\n")),
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
// `aws_secretsmanager_secret_policy` can be imported by using the secret Amazon Resource Name (ARN), e.g.
//
// ```sh
//  $ pulumi import aws:secretsmanager/secretPolicy:SecretPolicy example arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
// ```
type SecretPolicy struct {
	pulumi.CustomResourceState

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy pulumi.BoolPtrOutput `pulumi:"blockPublicPolicy"`
	Policy            pulumi.StringOutput  `pulumi:"policy"`
	// Secret ARN.
	SecretArn pulumi.StringOutput `pulumi:"secretArn"`
}

// NewSecretPolicy registers a new resource with the given unique name, arguments, and options.
func NewSecretPolicy(ctx *pulumi.Context,
	name string, args *SecretPolicyArgs, opts ...pulumi.ResourceOption) (*SecretPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.SecretArn == nil {
		return nil, errors.New("invalid value for required argument 'SecretArn'")
	}
	var resource SecretPolicy
	err := ctx.RegisterResource("aws:secretsmanager/secretPolicy:SecretPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretPolicy gets an existing SecretPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretPolicyState, opts ...pulumi.ResourceOption) (*SecretPolicy, error) {
	var resource SecretPolicy
	err := ctx.ReadResource("aws:secretsmanager/secretPolicy:SecretPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretPolicy resources.
type secretPolicyState struct {
	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy *bool   `pulumi:"blockPublicPolicy"`
	Policy            *string `pulumi:"policy"`
	// Secret ARN.
	SecretArn *string `pulumi:"secretArn"`
}

type SecretPolicyState struct {
	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy pulumi.BoolPtrInput
	Policy            pulumi.StringPtrInput
	// Secret ARN.
	SecretArn pulumi.StringPtrInput
}

func (SecretPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretPolicyState)(nil)).Elem()
}

type secretPolicyArgs struct {
	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy *bool  `pulumi:"blockPublicPolicy"`
	Policy            string `pulumi:"policy"`
	// Secret ARN.
	SecretArn string `pulumi:"secretArn"`
}

// The set of arguments for constructing a SecretPolicy resource.
type SecretPolicyArgs struct {
	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent broad access to your secret.
	BlockPublicPolicy pulumi.BoolPtrInput
	Policy            pulumi.StringInput
	// Secret ARN.
	SecretArn pulumi.StringInput
}

func (SecretPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretPolicyArgs)(nil)).Elem()
}

type SecretPolicyInput interface {
	pulumi.Input

	ToSecretPolicyOutput() SecretPolicyOutput
	ToSecretPolicyOutputWithContext(ctx context.Context) SecretPolicyOutput
}

func (*SecretPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretPolicy)(nil))
}

func (i *SecretPolicy) ToSecretPolicyOutput() SecretPolicyOutput {
	return i.ToSecretPolicyOutputWithContext(context.Background())
}

func (i *SecretPolicy) ToSecretPolicyOutputWithContext(ctx context.Context) SecretPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPolicyOutput)
}

func (i *SecretPolicy) ToSecretPolicyPtrOutput() SecretPolicyPtrOutput {
	return i.ToSecretPolicyPtrOutputWithContext(context.Background())
}

func (i *SecretPolicy) ToSecretPolicyPtrOutputWithContext(ctx context.Context) SecretPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPolicyPtrOutput)
}

type SecretPolicyPtrInput interface {
	pulumi.Input

	ToSecretPolicyPtrOutput() SecretPolicyPtrOutput
	ToSecretPolicyPtrOutputWithContext(ctx context.Context) SecretPolicyPtrOutput
}

type secretPolicyPtrType SecretPolicyArgs

func (*secretPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretPolicy)(nil))
}

func (i *secretPolicyPtrType) ToSecretPolicyPtrOutput() SecretPolicyPtrOutput {
	return i.ToSecretPolicyPtrOutputWithContext(context.Background())
}

func (i *secretPolicyPtrType) ToSecretPolicyPtrOutputWithContext(ctx context.Context) SecretPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPolicyPtrOutput)
}

// SecretPolicyArrayInput is an input type that accepts SecretPolicyArray and SecretPolicyArrayOutput values.
// You can construct a concrete instance of `SecretPolicyArrayInput` via:
//
//          SecretPolicyArray{ SecretPolicyArgs{...} }
type SecretPolicyArrayInput interface {
	pulumi.Input

	ToSecretPolicyArrayOutput() SecretPolicyArrayOutput
	ToSecretPolicyArrayOutputWithContext(context.Context) SecretPolicyArrayOutput
}

type SecretPolicyArray []SecretPolicyInput

func (SecretPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretPolicy)(nil))
}

func (i SecretPolicyArray) ToSecretPolicyArrayOutput() SecretPolicyArrayOutput {
	return i.ToSecretPolicyArrayOutputWithContext(context.Background())
}

func (i SecretPolicyArray) ToSecretPolicyArrayOutputWithContext(ctx context.Context) SecretPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPolicyArrayOutput)
}

// SecretPolicyMapInput is an input type that accepts SecretPolicyMap and SecretPolicyMapOutput values.
// You can construct a concrete instance of `SecretPolicyMapInput` via:
//
//          SecretPolicyMap{ "key": SecretPolicyArgs{...} }
type SecretPolicyMapInput interface {
	pulumi.Input

	ToSecretPolicyMapOutput() SecretPolicyMapOutput
	ToSecretPolicyMapOutputWithContext(context.Context) SecretPolicyMapOutput
}

type SecretPolicyMap map[string]SecretPolicyInput

func (SecretPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretPolicy)(nil))
}

func (i SecretPolicyMap) ToSecretPolicyMapOutput() SecretPolicyMapOutput {
	return i.ToSecretPolicyMapOutputWithContext(context.Background())
}

func (i SecretPolicyMap) ToSecretPolicyMapOutputWithContext(ctx context.Context) SecretPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretPolicyMapOutput)
}

type SecretPolicyOutput struct {
	*pulumi.OutputState
}

func (SecretPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretPolicy)(nil))
}

func (o SecretPolicyOutput) ToSecretPolicyOutput() SecretPolicyOutput {
	return o
}

func (o SecretPolicyOutput) ToSecretPolicyOutputWithContext(ctx context.Context) SecretPolicyOutput {
	return o
}

func (o SecretPolicyOutput) ToSecretPolicyPtrOutput() SecretPolicyPtrOutput {
	return o.ToSecretPolicyPtrOutputWithContext(context.Background())
}

func (o SecretPolicyOutput) ToSecretPolicyPtrOutputWithContext(ctx context.Context) SecretPolicyPtrOutput {
	return o.ApplyT(func(v SecretPolicy) *SecretPolicy {
		return &v
	}).(SecretPolicyPtrOutput)
}

type SecretPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (SecretPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretPolicy)(nil))
}

func (o SecretPolicyPtrOutput) ToSecretPolicyPtrOutput() SecretPolicyPtrOutput {
	return o
}

func (o SecretPolicyPtrOutput) ToSecretPolicyPtrOutputWithContext(ctx context.Context) SecretPolicyPtrOutput {
	return o
}

type SecretPolicyArrayOutput struct{ *pulumi.OutputState }

func (SecretPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretPolicy)(nil))
}

func (o SecretPolicyArrayOutput) ToSecretPolicyArrayOutput() SecretPolicyArrayOutput {
	return o
}

func (o SecretPolicyArrayOutput) ToSecretPolicyArrayOutputWithContext(ctx context.Context) SecretPolicyArrayOutput {
	return o
}

func (o SecretPolicyArrayOutput) Index(i pulumi.IntInput) SecretPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretPolicy {
		return vs[0].([]SecretPolicy)[vs[1].(int)]
	}).(SecretPolicyOutput)
}

type SecretPolicyMapOutput struct{ *pulumi.OutputState }

func (SecretPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretPolicy)(nil))
}

func (o SecretPolicyMapOutput) ToSecretPolicyMapOutput() SecretPolicyMapOutput {
	return o
}

func (o SecretPolicyMapOutput) ToSecretPolicyMapOutputWithContext(ctx context.Context) SecretPolicyMapOutput {
	return o
}

func (o SecretPolicyMapOutput) MapIndex(k pulumi.StringInput) SecretPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretPolicy {
		return vs[0].(map[string]SecretPolicy)[vs[1].(string)]
	}).(SecretPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretPolicyOutput{})
	pulumi.RegisterOutputType(SecretPolicyPtrOutput{})
	pulumi.RegisterOutputType(SecretPolicyArrayOutput{})
	pulumi.RegisterOutputType(SecretPolicyMapOutput{})
}
