// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Trigger struct {
	pulumi.CustomResourceState

	Actions      TriggerActionArrayOutput  `pulumi:"actions"`
	Arn          pulumi.StringOutput       `pulumi:"arn"`
	Description  pulumi.StringPtrOutput    `pulumi:"description"`
	Enabled      pulumi.BoolPtrOutput      `pulumi:"enabled"`
	Name         pulumi.StringOutput       `pulumi:"name"`
	Predicate    TriggerPredicatePtrOutput `pulumi:"predicate"`
	Schedule     pulumi.StringPtrOutput    `pulumi:"schedule"`
	Tags         pulumi.StringMapOutput    `pulumi:"tags"`
	Type         pulumi.StringOutput       `pulumi:"type"`
	WorkflowName pulumi.StringPtrOutput    `pulumi:"workflowName"`
}

// NewTrigger registers a new resource with the given unique name, arguments, and options.
func NewTrigger(ctx *pulumi.Context,
	name string, args *TriggerArgs, opts ...pulumi.ResourceOption) (*Trigger, error) {
	if args == nil || args.Actions == nil {
		return nil, errors.New("missing required argument 'Actions'")
	}
	if args == nil || args.Type == nil {
		return nil, errors.New("missing required argument 'Type'")
	}
	if args == nil {
		args = &TriggerArgs{}
	}
	var resource Trigger
	err := ctx.RegisterResource("aws:glue/trigger:Trigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrigger gets an existing Trigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerState, opts ...pulumi.ResourceOption) (*Trigger, error) {
	var resource Trigger
	err := ctx.ReadResource("aws:glue/trigger:Trigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trigger resources.
type triggerState struct {
	Actions      []TriggerAction   `pulumi:"actions"`
	Arn          *string           `pulumi:"arn"`
	Description  *string           `pulumi:"description"`
	Enabled      *bool             `pulumi:"enabled"`
	Name         *string           `pulumi:"name"`
	Predicate    *TriggerPredicate `pulumi:"predicate"`
	Schedule     *string           `pulumi:"schedule"`
	Tags         map[string]string `pulumi:"tags"`
	Type         *string           `pulumi:"type"`
	WorkflowName *string           `pulumi:"workflowName"`
}

type TriggerState struct {
	Actions      TriggerActionArrayInput
	Arn          pulumi.StringPtrInput
	Description  pulumi.StringPtrInput
	Enabled      pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	Predicate    TriggerPredicatePtrInput
	Schedule     pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
	Type         pulumi.StringPtrInput
	WorkflowName pulumi.StringPtrInput
}

func (TriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerState)(nil)).Elem()
}

type triggerArgs struct {
	Actions      []TriggerAction   `pulumi:"actions"`
	Description  *string           `pulumi:"description"`
	Enabled      *bool             `pulumi:"enabled"`
	Name         *string           `pulumi:"name"`
	Predicate    *TriggerPredicate `pulumi:"predicate"`
	Schedule     *string           `pulumi:"schedule"`
	Tags         map[string]string `pulumi:"tags"`
	Type         string            `pulumi:"type"`
	WorkflowName *string           `pulumi:"workflowName"`
}

// The set of arguments for constructing a Trigger resource.
type TriggerArgs struct {
	Actions      TriggerActionArrayInput
	Description  pulumi.StringPtrInput
	Enabled      pulumi.BoolPtrInput
	Name         pulumi.StringPtrInput
	Predicate    TriggerPredicatePtrInput
	Schedule     pulumi.StringPtrInput
	Tags         pulumi.StringMapInput
	Type         pulumi.StringInput
	WorkflowName pulumi.StringPtrInput
}

func (TriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerArgs)(nil)).Elem()
}
