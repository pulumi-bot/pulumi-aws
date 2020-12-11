// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Associates an SSM Document to an instance or EC2 tag.
//
// ## Import
//
// SSM associations can be imported using the `association_id`, e.g.
//
// ```sh
//  $ pulumi import aws:ssm/association:Association test-association 10abcdef-0abc-1234-5678-90abcdef123456
// ```
type Association struct {
	pulumi.CustomResourceState

	// The ID of the SSM association.
	AssociationId pulumi.StringOutput `pulumi:"associationId"`
	// The descriptive name for the association.
	AssociationName pulumi.StringPtrOutput `pulumi:"associationName"`
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
	AutomationTargetParameterName pulumi.StringPtrOutput `pulumi:"automationTargetParameterName"`
	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity pulumi.StringPtrOutput `pulumi:"complianceSeverity"`
	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion pulumi.StringOutput `pulumi:"documentVersion"`
	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency pulumi.StringPtrOutput `pulumi:"maxConcurrency"`
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors pulumi.StringPtrOutput `pulumi:"maxErrors"`
	// The name of the SSM document to apply.
	Name pulumi.StringOutput `pulumi:"name"`
	// An output location block. Output Location is documented below.
	OutputLocation AssociationOutputLocationPtrOutput `pulumi:"outputLocation"`
	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression pulumi.StringPtrOutput `pulumi:"scheduleExpression"`
	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets AssociationTargetArrayOutput `pulumi:"targets"`
}

// NewAssociation registers a new resource with the given unique name, arguments, and options.
func NewAssociation(ctx *pulumi.Context,
	name string, args *AssociationArgs, opts ...pulumi.ResourceOption) (*Association, error) {
	if args == nil {
		args = &AssociationArgs{}
	}

	var resource Association
	err := ctx.RegisterResource("aws:ssm/association:Association", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssociation gets an existing Association resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssociationState, opts ...pulumi.ResourceOption) (*Association, error) {
	var resource Association
	err := ctx.ReadResource("aws:ssm/association:Association", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Association resources.
type associationState struct {
	// The ID of the SSM association.
	AssociationId *string `pulumi:"associationId"`
	// The descriptive name for the association.
	AssociationName *string `pulumi:"associationName"`
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
	AutomationTargetParameterName *string `pulumi:"automationTargetParameterName"`
	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity *string `pulumi:"complianceSeverity"`
	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion *string `pulumi:"documentVersion"`
	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
	InstanceId *string `pulumi:"instanceId"`
	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency *string `pulumi:"maxConcurrency"`
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors *string `pulumi:"maxErrors"`
	// The name of the SSM document to apply.
	Name *string `pulumi:"name"`
	// An output location block. Output Location is documented below.
	OutputLocation *AssociationOutputLocation `pulumi:"outputLocation"`
	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters map[string]string `pulumi:"parameters"`
	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression *string `pulumi:"scheduleExpression"`
	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets []AssociationTarget `pulumi:"targets"`
}

type AssociationState struct {
	// The ID of the SSM association.
	AssociationId pulumi.StringPtrInput
	// The descriptive name for the association.
	AssociationName pulumi.StringPtrInput
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
	AutomationTargetParameterName pulumi.StringPtrInput
	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity pulumi.StringPtrInput
	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion pulumi.StringPtrInput
	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
	InstanceId pulumi.StringPtrInput
	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency pulumi.StringPtrInput
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors pulumi.StringPtrInput
	// The name of the SSM document to apply.
	Name pulumi.StringPtrInput
	// An output location block. Output Location is documented below.
	OutputLocation AssociationOutputLocationPtrInput
	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters pulumi.StringMapInput
	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression pulumi.StringPtrInput
	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets AssociationTargetArrayInput
}

func (AssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*associationState)(nil)).Elem()
}

type associationArgs struct {
	// The descriptive name for the association.
	AssociationName *string `pulumi:"associationName"`
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
	AutomationTargetParameterName *string `pulumi:"automationTargetParameterName"`
	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity *string `pulumi:"complianceSeverity"`
	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion *string `pulumi:"documentVersion"`
	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
	InstanceId *string `pulumi:"instanceId"`
	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency *string `pulumi:"maxConcurrency"`
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors *string `pulumi:"maxErrors"`
	// The name of the SSM document to apply.
	Name *string `pulumi:"name"`
	// An output location block. Output Location is documented below.
	OutputLocation *AssociationOutputLocation `pulumi:"outputLocation"`
	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters map[string]string `pulumi:"parameters"`
	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression *string `pulumi:"scheduleExpression"`
	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets []AssociationTarget `pulumi:"targets"`
}

// The set of arguments for constructing a Association resource.
type AssociationArgs struct {
	// The descriptive name for the association.
	AssociationName pulumi.StringPtrInput
	// Specify the target for the association. This target is required for associations that use an `Automation` document and target resources by using rate controls.
	AutomationTargetParameterName pulumi.StringPtrInput
	// The compliance severity for the association. Can be one of the following: `UNSPECIFIED`, `LOW`, `MEDIUM`, `HIGH` or `CRITICAL`
	ComplianceSeverity pulumi.StringPtrInput
	// The document version you want to associate with the target(s). Can be a specific version or the default version.
	DocumentVersion pulumi.StringPtrInput
	// The instance ID to apply an SSM document to. Use `targets` with key `InstanceIds` for document schema versions 2.0 and above.
	InstanceId pulumi.StringPtrInput
	// The maximum number of targets allowed to run the association at the same time. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxConcurrency pulumi.StringPtrInput
	// The number of errors that are allowed before the system stops sending requests to run the association on additional targets. You can specify a number, for example 10, or a percentage of the target set, for example 10%.
	MaxErrors pulumi.StringPtrInput
	// The name of the SSM document to apply.
	Name pulumi.StringPtrInput
	// An output location block. Output Location is documented below.
	OutputLocation AssociationOutputLocationPtrInput
	// A block of arbitrary string parameters to pass to the SSM document.
	Parameters pulumi.StringMapInput
	// A cron expression when the association will be applied to the target(s).
	ScheduleExpression pulumi.StringPtrInput
	// A block containing the targets of the SSM association. Targets are documented below. AWS currently supports a maximum of 5 targets.
	Targets AssociationTargetArrayInput
}

func (AssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*associationArgs)(nil)).Elem()
}

type AssociationInput interface {
	pulumi.Input

	ToAssociationOutput() AssociationOutput
	ToAssociationOutputWithContext(ctx context.Context) AssociationOutput
}

type AssociationPtrInput interface {
	pulumi.Input

	ToAssociationPtrOutput() AssociationPtrOutput
	ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput
}

func (Association) ElementType() reflect.Type {
	return reflect.TypeOf((*Association)(nil)).Elem()
}

func (i Association) ToAssociationOutput() AssociationOutput {
	return i.ToAssociationOutputWithContext(context.Background())
}

func (i Association) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationOutput)
}

func (i Association) ToAssociationPtrOutput() AssociationPtrOutput {
	return i.ToAssociationPtrOutputWithContext(context.Background())
}

func (i Association) ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssociationPtrOutput)
}

type AssociationOutput struct {
	*pulumi.OutputState
}

func (AssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssociationOutput)(nil)).Elem()
}

func (o AssociationOutput) ToAssociationOutput() AssociationOutput {
	return o
}

func (o AssociationOutput) ToAssociationOutputWithContext(ctx context.Context) AssociationOutput {
	return o
}

type AssociationPtrOutput struct {
	*pulumi.OutputState
}

func (AssociationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Association)(nil)).Elem()
}

func (o AssociationPtrOutput) ToAssociationPtrOutput() AssociationPtrOutput {
	return o
}

func (o AssociationPtrOutput) ToAssociationPtrOutputWithContext(ctx context.Context) AssociationPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AssociationOutput{})
	pulumi.RegisterOutputType(AssociationPtrOutput{})
}
