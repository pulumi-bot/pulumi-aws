// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AvailabilityZoneGroup struct {
	pulumi.CustomResourceState

	GroupName   pulumi.StringOutput `pulumi:"groupName"`
	OptInStatus pulumi.StringOutput `pulumi:"optInStatus"`
}

// NewAvailabilityZoneGroup registers a new resource with the given unique name, arguments, and options.
func NewAvailabilityZoneGroup(ctx *pulumi.Context,
	name string, args *AvailabilityZoneGroupArgs, opts ...pulumi.ResourceOption) (*AvailabilityZoneGroup, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.OptInStatus == nil {
		return nil, errors.New("missing required argument 'OptInStatus'")
	}
	if args == nil {
		args = &AvailabilityZoneGroupArgs{}
	}
	var resource AvailabilityZoneGroup
	err := ctx.RegisterResource("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilityZoneGroup gets an existing AvailabilityZoneGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilityZoneGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilityZoneGroupState, opts ...pulumi.ResourceOption) (*AvailabilityZoneGroup, error) {
	var resource AvailabilityZoneGroup
	err := ctx.ReadResource("aws:ec2/availabilityZoneGroup:AvailabilityZoneGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilityZoneGroup resources.
type availabilityZoneGroupState struct {
	GroupName   *string `pulumi:"groupName"`
	OptInStatus *string `pulumi:"optInStatus"`
}

type AvailabilityZoneGroupState struct {
	GroupName   pulumi.StringPtrInput
	OptInStatus pulumi.StringPtrInput
}

func (AvailabilityZoneGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityZoneGroupState)(nil)).Elem()
}

type availabilityZoneGroupArgs struct {
	GroupName   string `pulumi:"groupName"`
	OptInStatus string `pulumi:"optInStatus"`
}

// The set of arguments for constructing a AvailabilityZoneGroup resource.
type AvailabilityZoneGroupArgs struct {
	GroupName   pulumi.StringInput
	OptInStatus pulumi.StringInput
}

func (AvailabilityZoneGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilityZoneGroupArgs)(nil)).Elem()
}
