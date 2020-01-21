// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package plan

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"https:/github.com/pulumi/pulumi-aws/backup/PlanRule"
	"https:/github.com/pulumi/pulumi-aws/backup/PlanRuleLifecycle"
)

// Provides an AWS Backup plan resource.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/backup_plan.html.markdown.
type Plan struct {
	pulumi.CustomResourceState

	// The ARN of the backup plan.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The display name of a backup plan.
	Name pulumi.StringOutput `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules backupPlanRule.PlanRuleArrayOutput `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewPlan registers a new resource with the given unique name, arguments, and options.
func NewPlan(ctx *pulumi.Context,
	name string, args *PlanArgs, opts ...pulumi.ResourceOption) (*Plan, error) {
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	if args == nil {
		args = &PlanArgs{}
	}
	var resource Plan
	err := ctx.RegisterResource("aws:backup/plan:Plan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlan gets an existing Plan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlanState, opts ...pulumi.ResourceOption) (*Plan, error) {
	var resource Plan
	err := ctx.ReadResource("aws:backup/plan:Plan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Plan resources.
type planState struct {
	// The ARN of the backup plan.
	Arn *string `pulumi:"arn"`
	// The display name of a backup plan.
	Name *string `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules []backupPlanRule.PlanRule `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags map[string]interface{} `pulumi:"tags"`
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version *string `pulumi:"version"`
}

type PlanState struct {
	// The ARN of the backup plan.
	Arn pulumi.StringPtrInput
	// The display name of a backup plan.
	Name pulumi.StringPtrInput
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules backupPlanRule.PlanRuleArrayInput
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.MapInput
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version pulumi.StringPtrInput
}

func (PlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*planState)(nil)).Elem()
}

type planArgs struct {
	// The display name of a backup plan.
	Name *string `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules []backupPlanRule.PlanRule `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags map[string]interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Plan resource.
type PlanArgs struct {
	// The display name of a backup plan.
	Name pulumi.StringPtrInput
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules backupPlanRule.PlanRuleArrayInput
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.MapInput
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*planArgs)(nil)).Elem()
}

