// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type EventPermission struct {
	pulumi.CustomResourceState

	Action      pulumi.StringPtrOutput            `pulumi:"action"`
	Condition   EventPermissionConditionPtrOutput `pulumi:"condition"`
	Principal   pulumi.StringOutput               `pulumi:"principal"`
	StatementId pulumi.StringOutput               `pulumi:"statementId"`
}

// NewEventPermission registers a new resource with the given unique name, arguments, and options.
func NewEventPermission(ctx *pulumi.Context,
	name string, args *EventPermissionArgs, opts ...pulumi.ResourceOption) (*EventPermission, error) {
	if args == nil || args.Principal == nil {
		return nil, errors.New("missing required argument 'Principal'")
	}
	if args == nil || args.StatementId == nil {
		return nil, errors.New("missing required argument 'StatementId'")
	}
	if args == nil {
		args = &EventPermissionArgs{}
	}
	var resource EventPermission
	err := ctx.RegisterResource("aws:cloudwatch/eventPermission:EventPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventPermission gets an existing EventPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventPermissionState, opts ...pulumi.ResourceOption) (*EventPermission, error) {
	var resource EventPermission
	err := ctx.ReadResource("aws:cloudwatch/eventPermission:EventPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventPermission resources.
type eventPermissionState struct {
	Action      *string                   `pulumi:"action"`
	Condition   *EventPermissionCondition `pulumi:"condition"`
	Principal   *string                   `pulumi:"principal"`
	StatementId *string                   `pulumi:"statementId"`
}

type EventPermissionState struct {
	Action      pulumi.StringPtrInput
	Condition   EventPermissionConditionPtrInput
	Principal   pulumi.StringPtrInput
	StatementId pulumi.StringPtrInput
}

func (EventPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPermissionState)(nil)).Elem()
}

type eventPermissionArgs struct {
	Action      *string                   `pulumi:"action"`
	Condition   *EventPermissionCondition `pulumi:"condition"`
	Principal   string                    `pulumi:"principal"`
	StatementId string                    `pulumi:"statementId"`
}

// The set of arguments for constructing a EventPermission resource.
type EventPermissionArgs struct {
	Action      pulumi.StringPtrInput
	Condition   EventPermissionConditionPtrInput
	Principal   pulumi.StringInput
	StatementId pulumi.StringInput
}

func (EventPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventPermissionArgs)(nil)).Elem()
}
