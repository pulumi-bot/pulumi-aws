// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Config Organization Managed Rule. More information about these rules can be found in the [Enabling AWS Config Rules Across all Accounts in Your Organization](https://docs.aws.amazon.com/config/latest/developerguide/config-rule-multi-account-deployment.html) and [AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) documentation. For working with Organization Custom Rules (those invoking a custom Lambda Function), see the `cfg.OrganizationCustomRule` resource.
//
// > **NOTE:** This resource must be created in the Organization master account and rules will include the master account unless its ID is added to the `excludedAccounts` argument.
//
// > **NOTE:** Every Organization account except those configured in the `excludedAccounts` argument must have a Configuration Recorder with proper IAM permissions before the rule will successfully create or update. See also the `cfg.Recorder` resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
// 			AwsServiceAccessPrincipals: pulumi.StringArray{
// 				pulumi.String("config-multiaccountsetup.amazonaws.com"),
// 			},
// 			FeatureSet: pulumi.String("ALL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewOrganizationManagedRule(ctx, "exampleOrganizationManagedRule", &cfg.OrganizationManagedRuleArgs{
// 			RuleIdentifier: pulumi.String("IAM_PASSWORD_POLICY"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			exampleOrganization,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrganizationManagedRule struct {
	pulumi.CustomResourceState

	// Amazon Resource Name (ARN) of the rule
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the rule
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts pulumi.StringArrayOutput `pulumi:"excludedAccounts"`
	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters pulumi.StringPtrOutput `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrOutput `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name pulumi.StringOutput `pulumi:"name"`
	// Identifier of the AWS resource to evaluate
	ResourceIdScope pulumi.StringPtrOutput `pulumi:"resourceIdScope"`
	// List of types of AWS resources to evaluate
	ResourceTypesScopes pulumi.StringArrayOutput `pulumi:"resourceTypesScopes"`
	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier pulumi.StringOutput `pulumi:"ruleIdentifier"`
	// Tag key of AWS resources to evaluate
	TagKeyScope pulumi.StringPtrOutput `pulumi:"tagKeyScope"`
	// Tag value of AWS resources to evaluate
	TagValueScope pulumi.StringPtrOutput `pulumi:"tagValueScope"`
}

// NewOrganizationManagedRule registers a new resource with the given unique name, arguments, and options.
func NewOrganizationManagedRule(ctx *pulumi.Context,
	name string, args *OrganizationManagedRuleArgs, opts ...pulumi.ResourceOption) (*OrganizationManagedRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuleIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'RuleIdentifier'")
	}
	var resource OrganizationManagedRule
	err := ctx.RegisterResource("aws:cfg/organizationManagedRule:OrganizationManagedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationManagedRule gets an existing OrganizationManagedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationManagedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationManagedRuleState, opts ...pulumi.ResourceOption) (*OrganizationManagedRule, error) {
	var resource OrganizationManagedRule
	err := ctx.ReadResource("aws:cfg/organizationManagedRule:OrganizationManagedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationManagedRule resources.
type organizationManagedRuleState struct {
	// Amazon Resource Name (ARN) of the rule
	Arn *string `pulumi:"arn"`
	// Description of the rule
	Description *string `pulumi:"description"`
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts []string `pulumi:"excludedAccounts"`
	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters *string `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// Identifier of the AWS resource to evaluate
	ResourceIdScope *string `pulumi:"resourceIdScope"`
	// List of types of AWS resources to evaluate
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier *string `pulumi:"ruleIdentifier"`
	// Tag key of AWS resources to evaluate
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// Tag value of AWS resources to evaluate
	TagValueScope *string `pulumi:"tagValueScope"`
}

type OrganizationManagedRuleState struct {
	// Amazon Resource Name (ARN) of the rule
	Arn pulumi.StringPtrInput
	// Description of the rule
	Description pulumi.StringPtrInput
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts pulumi.StringArrayInput
	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters pulumi.StringPtrInput
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// Identifier of the AWS resource to evaluate
	ResourceIdScope pulumi.StringPtrInput
	// List of types of AWS resources to evaluate
	ResourceTypesScopes pulumi.StringArrayInput
	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier pulumi.StringPtrInput
	// Tag key of AWS resources to evaluate
	TagKeyScope pulumi.StringPtrInput
	// Tag value of AWS resources to evaluate
	TagValueScope pulumi.StringPtrInput
}

func (OrganizationManagedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagedRuleState)(nil)).Elem()
}

type organizationManagedRuleArgs struct {
	// Description of the rule
	Description *string `pulumi:"description"`
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts []string `pulumi:"excludedAccounts"`
	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters *string `pulumi:"inputParameters"`
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// Identifier of the AWS resource to evaluate
	ResourceIdScope *string `pulumi:"resourceIdScope"`
	// List of types of AWS resources to evaluate
	ResourceTypesScopes []string `pulumi:"resourceTypesScopes"`
	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier string `pulumi:"ruleIdentifier"`
	// Tag key of AWS resources to evaluate
	TagKeyScope *string `pulumi:"tagKeyScope"`
	// Tag value of AWS resources to evaluate
	TagValueScope *string `pulumi:"tagValueScope"`
}

// The set of arguments for constructing a OrganizationManagedRule resource.
type OrganizationManagedRuleArgs struct {
	// Description of the rule
	Description pulumi.StringPtrInput
	// List of AWS account identifiers to exclude from the rule
	ExcludedAccounts pulumi.StringArrayInput
	// A string in JSON format that is passed to the AWS Config Rule Lambda Function
	InputParameters pulumi.StringPtrInput
	// The maximum frequency with which AWS Config runs evaluations for a rule, if the rule is triggered at a periodic frequency. Defaults to `TwentyFour_Hours` for periodic frequency triggered rules. Valid values: `One_Hour`, `Three_Hours`, `Six_Hours`, `Twelve_Hours`, or `TwentyFour_Hours`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// Identifier of the AWS resource to evaluate
	ResourceIdScope pulumi.StringPtrInput
	// List of types of AWS resources to evaluate
	ResourceTypesScopes pulumi.StringArrayInput
	// Identifier of an available AWS Config Managed Rule to call. For available values, see the [List of AWS Config Managed Rules](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html) documentation
	RuleIdentifier pulumi.StringInput
	// Tag key of AWS resources to evaluate
	TagKeyScope pulumi.StringPtrInput
	// Tag value of AWS resources to evaluate
	TagValueScope pulumi.StringPtrInput
}

func (OrganizationManagedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationManagedRuleArgs)(nil)).Elem()
}

type OrganizationManagedRuleInput interface {
	pulumi.Input

	ToOrganizationManagedRuleOutput() OrganizationManagedRuleOutput
	ToOrganizationManagedRuleOutputWithContext(ctx context.Context) OrganizationManagedRuleOutput
}

func (OrganizationManagedRule) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationManagedRule)(nil)).Elem()
}

func (i OrganizationManagedRule) ToOrganizationManagedRuleOutput() OrganizationManagedRuleOutput {
	return i.ToOrganizationManagedRuleOutputWithContext(context.Background())
}

func (i OrganizationManagedRule) ToOrganizationManagedRuleOutputWithContext(ctx context.Context) OrganizationManagedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationManagedRuleOutput)
}

type OrganizationManagedRuleOutput struct {
	*pulumi.OutputState
}

func (OrganizationManagedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationManagedRuleOutput)(nil)).Elem()
}

func (o OrganizationManagedRuleOutput) ToOrganizationManagedRuleOutput() OrganizationManagedRuleOutput {
	return o
}

func (o OrganizationManagedRuleOutput) ToOrganizationManagedRuleOutputWithContext(ctx context.Context) OrganizationManagedRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationManagedRuleOutput{})
}
