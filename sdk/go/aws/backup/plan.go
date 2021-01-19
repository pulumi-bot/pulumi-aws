// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package backup

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an AWS Backup plan resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/backup"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := backup.NewPlan(ctx, "example", &backup.PlanArgs{
// 			Rules: backup.PlanRuleArray{
// 				&backup.PlanRuleArgs{
// 					RuleName:        pulumi.String("tf_example_backup_rule"),
// 					TargetVaultName: pulumi.Any(aws_backup_vault.Test.Name),
// 					Schedule:        pulumi.String("cron(0 12 * * ? *)"),
// 				},
// 			},
// 			AdvancedBackupSettings: backup.PlanAdvancedBackupSettingArray{
// 				&backup.PlanAdvancedBackupSettingArgs{
// 					BackupOptions: pulumi.StringMap{
// 						"WindowsVSS": pulumi.String("enabled"),
// 					},
// 					ResourceType: pulumi.String("EC2"),
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
// Backup Plan can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:backup/plan:Plan test <id>
// ```
type Plan struct {
	pulumi.CustomResourceState

	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings PlanAdvancedBackupSettingArrayOutput `pulumi:"advancedBackupSettings"`
	// The ARN of the backup plan.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The display name of a backup plan.
	Name pulumi.StringOutput `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules PlanRuleArrayOutput `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewPlan registers a new resource with the given unique name, arguments, and options.
func NewPlan(ctx *pulumi.Context,
	name string, args *PlanArgs, opts ...pulumi.ResourceOption) (*Plan, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
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
	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings []PlanAdvancedBackupSetting `pulumi:"advancedBackupSettings"`
	// The ARN of the backup plan.
	Arn *string `pulumi:"arn"`
	// The display name of a backup plan.
	Name *string `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules []PlanRule `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags map[string]string `pulumi:"tags"`
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version *string `pulumi:"version"`
}

type PlanState struct {
	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings PlanAdvancedBackupSettingArrayInput
	// The ARN of the backup plan.
	Arn pulumi.StringPtrInput
	// The display name of a backup plan.
	Name pulumi.StringPtrInput
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules PlanRuleArrayInput
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.StringMapInput
	// Unique, randomly generated, Unicode, UTF-8 encoded string that serves as the version ID of the backup plan.
	Version pulumi.StringPtrInput
}

func (PlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*planState)(nil)).Elem()
}

type planArgs struct {
	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings []PlanAdvancedBackupSetting `pulumi:"advancedBackupSettings"`
	// The display name of a backup plan.
	Name *string `pulumi:"name"`
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules []PlanRule `pulumi:"rules"`
	// Metadata that you can assign to help organize the plans you create.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Plan resource.
type PlanArgs struct {
	// An object that specifies backup options for each resource type.
	AdvancedBackupSettings PlanAdvancedBackupSettingArrayInput
	// The display name of a backup plan.
	Name pulumi.StringPtrInput
	// A rule object that specifies a scheduled task that is used to back up a selection of resources.
	Rules PlanRuleArrayInput
	// Metadata that you can assign to help organize the plans you create.
	Tags pulumi.StringMapInput
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*planArgs)(nil)).Elem()
}

type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(ctx context.Context) PlanOutput
}

func (*Plan) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil))
}

func (i *Plan) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i *Plan) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}

func (i *Plan) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *Plan) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

type PlanPtrInput interface {
	pulumi.Input

	ToPlanPtrOutput() PlanPtrOutput
	ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput
}

type planPtrType PlanArgs

func (*planPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil))
}

func (i *planPtrType) ToPlanPtrOutput() PlanPtrOutput {
	return i.ToPlanPtrOutputWithContext(context.Background())
}

func (i *planPtrType) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanPtrOutput)
}

// PlanArrayInput is an input type that accepts PlanArray and PlanArrayOutput values.
// You can construct a concrete instance of `PlanArrayInput` via:
//
//          PlanArray{ PlanArgs{...} }
type PlanArrayInput interface {
	pulumi.Input

	ToPlanArrayOutput() PlanArrayOutput
	ToPlanArrayOutputWithContext(context.Context) PlanArrayOutput
}

type PlanArray []PlanInput

func (PlanArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Plan)(nil))
}

func (i PlanArray) ToPlanArrayOutput() PlanArrayOutput {
	return i.ToPlanArrayOutputWithContext(context.Background())
}

func (i PlanArray) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanArrayOutput)
}

// PlanMapInput is an input type that accepts PlanMap and PlanMapOutput values.
// You can construct a concrete instance of `PlanMapInput` via:
//
//          PlanMap{ "key": PlanArgs{...} }
type PlanMapInput interface {
	pulumi.Input

	ToPlanMapOutput() PlanMapOutput
	ToPlanMapOutputWithContext(context.Context) PlanMapOutput
}

type PlanMap map[string]PlanInput

func (PlanMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Plan)(nil))
}

func (i PlanMap) ToPlanMapOutput() PlanMapOutput {
	return i.ToPlanMapOutputWithContext(context.Background())
}

func (i PlanMap) ToPlanMapOutputWithContext(ctx context.Context) PlanMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanMapOutput)
}

type PlanOutput struct {
	*pulumi.OutputState
}

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil))
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o.ToPlanPtrOutputWithContext(context.Background())
}

func (o PlanOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o.ApplyT(func(v Plan) *Plan {
		return &v
	}).(PlanPtrOutput)
}

type PlanPtrOutput struct {
	*pulumi.OutputState
}

func (PlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Plan)(nil))
}

func (o PlanPtrOutput) ToPlanPtrOutput() PlanPtrOutput {
	return o
}

func (o PlanPtrOutput) ToPlanPtrOutputWithContext(ctx context.Context) PlanPtrOutput {
	return o
}

type PlanArrayOutput struct{ *pulumi.OutputState }

func (PlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil))
}

func (o PlanArrayOutput) ToPlanArrayOutput() PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) Index(i pulumi.IntInput) PlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Plan {
		return vs[0].([]Plan)[vs[1].(int)]
	}).(PlanOutput)
}

type PlanMapOutput struct{ *pulumi.OutputState }

func (PlanMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Plan)(nil))
}

func (o PlanMapOutput) ToPlanMapOutput() PlanMapOutput {
	return o
}

func (o PlanMapOutput) ToPlanMapOutputWithContext(ctx context.Context) PlanMapOutput {
	return o
}

func (o PlanMapOutput) MapIndex(k pulumi.StringInput) PlanOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Plan {
		return vs[0].(map[string]Plan)[vs[1].(string)]
	}).(PlanOutput)
}

func init() {
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanPtrOutput{})
	pulumi.RegisterOutputType(PlanArrayOutput{})
	pulumi.RegisterOutputType(PlanMapOutput{})
}
