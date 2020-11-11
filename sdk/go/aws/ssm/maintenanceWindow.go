// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an SSM Maintenance Window resource
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewMaintenanceWindow(ctx, "production", &ssm.MaintenanceWindowArgs{
// 			Cutoff:   pulumi.Int(1),
// 			Duration: pulumi.Int(3),
// 			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrOutput `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntOutput `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntOutput `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// The name of the maintenance window.
	Name pulumi.StringOutput `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-cron.html) or rate expression.
	Schedule pulumi.StringOutput `pulumi:"schedule"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrOutput `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil || args.Cutoff == nil {
		return nil, errors.New("missing required argument 'Cutoff'")
	}
	if args == nil || args.Duration == nil {
		return nil, errors.New("missing required argument 'Duration'")
	}
	if args == nil || args.Schedule == nil {
		return nil, errors.New("missing required argument 'Schedule'")
	}
	if args == nil {
		args = &MaintenanceWindowArgs{}
	}
	var resource MaintenanceWindow
	err := ctx.RegisterResource("aws:ssm/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("aws:ssm/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff *int `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description *string `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration *int `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate *string `pulumi:"endDate"`
	// The name of the maintenance window.
	Name *string `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-cron.html) or rate expression.
	Schedule *string `pulumi:"schedule"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone *string `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate *string `pulumi:"startDate"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type MaintenanceWindowState struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrInput
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntPtrInput
	// A description for the maintenance window.
	Description pulumi.StringPtrInput
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntPtrInput
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrInput
	// The name of the maintenance window.
	Name pulumi.StringPtrInput
	// The schedule of the Maintenance Window in the form of a [cron](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-cron.html) or rate expression.
	Schedule pulumi.StringPtrInput
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `pulumi:"allowUnassociatedTargets"`
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff int `pulumi:"cutoff"`
	// A description for the maintenance window.
	Description *string `pulumi:"description"`
	// The duration of the Maintenance Window in hours.
	Duration int `pulumi:"duration"`
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled *bool `pulumi:"enabled"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate *string `pulumi:"endDate"`
	// The name of the maintenance window.
	Name *string `pulumi:"name"`
	// The schedule of the Maintenance Window in the form of a [cron](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-cron.html) or rate expression.
	Schedule string `pulumi:"schedule"`
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone *string `pulumi:"scheduleTimezone"`
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate *string `pulumi:"startDate"`
	// A map of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets pulumi.BoolPtrInput
	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff pulumi.IntInput
	// A description for the maintenance window.
	Description pulumi.StringPtrInput
	// The duration of the Maintenance Window in hours.
	Duration pulumi.IntInput
	// Whether the maintenance window is enabled. Default: `true`.
	Enabled pulumi.BoolPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to no longer run the maintenance window.
	EndDate pulumi.StringPtrInput
	// The name of the maintenance window.
	Name pulumi.StringPtrInput
	// The schedule of the Maintenance Window in the form of a [cron](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-maintenance-cron.html) or rate expression.
	Schedule pulumi.StringInput
	// Timezone for schedule in [Internet Assigned Numbers Authority (IANA) Time Zone Database format](https://www.iana.org/time-zones). For example: `America/Los_Angeles`, `etc/UTC`, or `Asia/Seoul`.
	ScheduleTimezone pulumi.StringPtrInput
	// Timestamp in [ISO-8601 extended format](https://www.iso.org/iso-8601-date-and-time-format.html) when to begin the maintenance window.
	StartDate pulumi.StringPtrInput
	// A map of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

type MaintenanceWindowOutput struct {
	*pulumi.OutputState
}

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MaintenanceWindowOutput)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
}
