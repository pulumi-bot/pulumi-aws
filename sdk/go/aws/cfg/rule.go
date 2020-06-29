// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cfg

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Config Rule.
//
// > **Note:** Config Rule requires an existing `Configuration Recorder` to be present. Use of `dependsOn` is recommended (as shown below) to avoid race conditions.
//
// ## Example Usage
// ### AWS Managed Rules
//
// AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cfg.NewRule(ctx, "rule", &cfg.RuleArgs{
// 			Source: &cfg.RuleSourceArgs{
// 				Owner:            pulumi.String("AWS"),
// 				SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			"aws_config_configuration_recorder.foo",
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": \"sts:AssumeRole\",\n", "      \"Principal\": {\n", "        \"Service\": \"config.amazonaws.com\"\n", "      },\n", "      \"Effect\": \"Allow\",\n", "      \"Sid\": \"\"\n", "    }\n", "  ]\n", "}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
// 			RoleArn: role.Arn,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
// 			Policy: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "  	{\n", "  		\"Action\": \"config:Put*\",\n", "  		\"Effect\": \"Allow\",\n", "  		\"Resource\": \"*\"\n", "\n", "  	}\n", "  ]\n", "}\n", "\n")),
// 			Role: role.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Custom Rules
//
// Custom rules can be used by setting the source owner to `CUSTOM_LAMBDA` and the source identifier to the Amazon Resource Name (ARN) of the Lambda Function. The AWS Config service must have permissions to invoke the Lambda Function, e.g. via the `lambda.Permission` resource. More information about custom rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/cfg"
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/lambda"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err = cfg.NewRecorder(ctx, "exampleRecorder", nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleFunction, err := lambda.NewFunction(ctx, "exampleFunction", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
// 			Action:    pulumi.String("lambda:InvokeFunction"),
// 			Function:  exampleFunction.Arn,
// 			Principal: pulumi.String("config.amazonaws.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cfg.NewRule(ctx, "exampleRule", &cfg.RuleArgs{
// 			Source: &cfg.RuleSourceArgs{
// 				Owner:            pulumi.String("CUSTOM_LAMBDA"),
// 				SourceIdentifier: exampleFunction.Arn,
// 			},
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			"aws_config_configuration_recorder.example",
// 			"aws_lambda_permission.example",
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Rule struct {
	pulumi.CustomResourceState

	// The ARN of the config rule
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Description of the rule
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrOutput `pulumi:"inputParameters"`
	// The frequency that you want AWS Config to run evaluations for a rule that
	// is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringPtrOutput `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the config rule
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Scope defines which resources can trigger an evaluation for the rule as documented below.
	Scope RuleScopePtrOutput `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause
	// the function to evaluate your AWS resources as documented below.
	Source RuleSourceOutput `pulumi:"source"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil || args.Source == nil {
		return nil, errors.New("missing required argument 'Source'")
	}
	if args == nil {
		args = &RuleArgs{}
	}
	var resource Rule
	err := ctx.RegisterResource("aws:cfg/rule:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws:cfg/rule:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
	// The ARN of the config rule
	Arn *string `pulumi:"arn"`
	// Description of the rule
	Description *string `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `pulumi:"inputParameters"`
	// The frequency that you want AWS Config to run evaluations for a rule that
	// is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// The ID of the config rule
	RuleId *string `pulumi:"ruleId"`
	// Scope defines which resources can trigger an evaluation for the rule as documented below.
	Scope *RuleScope `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause
	// the function to evaluate your AWS resources as documented below.
	Source *RuleSource `pulumi:"source"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type RuleState struct {
	// The ARN of the config rule
	Arn pulumi.StringPtrInput
	// Description of the rule
	Description pulumi.StringPtrInput
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrInput
	// The frequency that you want AWS Config to run evaluations for a rule that
	// is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// The ID of the config rule
	RuleId pulumi.StringPtrInput
	// Scope defines which resources can trigger an evaluation for the rule as documented below.
	Scope RuleScopePtrInput
	// Source specifies the rule owner, the rule identifier, and the notifications that cause
	// the function to evaluate your AWS resources as documented below.
	Source RuleSourcePtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	// Description of the rule
	Description *string `pulumi:"description"`
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters *string `pulumi:"inputParameters"`
	// The frequency that you want AWS Config to run evaluations for a rule that
	// is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
	MaximumExecutionFrequency *string `pulumi:"maximumExecutionFrequency"`
	// The name of the rule
	Name *string `pulumi:"name"`
	// Scope defines which resources can trigger an evaluation for the rule as documented below.
	Scope *RuleScope `pulumi:"scope"`
	// Source specifies the rule owner, the rule identifier, and the notifications that cause
	// the function to evaluate your AWS resources as documented below.
	Source RuleSource `pulumi:"source"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	// Description of the rule
	Description pulumi.StringPtrInput
	// A string in JSON format that is passed to the AWS Config rule Lambda function.
	InputParameters pulumi.StringPtrInput
	// The frequency that you want AWS Config to run evaluations for a rule that
	// is triggered periodically. If specified, requires `messageType` to be `ScheduledNotification`.
	MaximumExecutionFrequency pulumi.StringPtrInput
	// The name of the rule
	Name pulumi.StringPtrInput
	// Scope defines which resources can trigger an evaluation for the rule as documented below.
	Scope RuleScopePtrInput
	// Source specifies the rule owner, the rule identifier, and the notifications that cause
	// the function to evaluate your AWS resources as documented below.
	Source RuleSourceInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}
