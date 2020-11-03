// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type InstanceFleet struct {
	pulumi.CustomResourceState

	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Configuration block for instance fleet
	InstanceTypeConfigs InstanceFleetInstanceTypeConfigArrayOutput `pulumi:"instanceTypeConfigs"`
	// Configuration block for launch specification
	LaunchSpecifications InstanceFleetLaunchSpecificationsPtrOutput `pulumi:"launchSpecifications"`
	// Friendly name given to the instance fleet.
	Name                        pulumi.StringOutput `pulumi:"name"`
	ProvisionedOnDemandCapacity pulumi.IntOutput    `pulumi:"provisionedOnDemandCapacity"`
	ProvisionedSpotCapacity     pulumi.IntOutput    `pulumi:"provisionedSpotCapacity"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity pulumi.IntPtrOutput `pulumi:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity pulumi.IntPtrOutput `pulumi:"targetSpotCapacity"`
}

// NewInstanceFleet registers a new resource with the given unique name, arguments, and options.
func NewInstanceFleet(ctx *pulumi.Context,
	name string, args *InstanceFleetArgs, opts ...pulumi.ResourceOption) (*InstanceFleet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	var resource InstanceFleet
	err := ctx.RegisterResource("aws:emr/instanceFleet:InstanceFleet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceFleet gets an existing InstanceFleet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceFleet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceFleetState, opts ...pulumi.ResourceOption) (*InstanceFleet, error) {
	var resource InstanceFleet
	err := ctx.ReadResource("aws:emr/instanceFleet:InstanceFleet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceFleet resources.
type instanceFleetState struct {
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId *string `pulumi:"clusterId"`
	// Configuration block for instance fleet
	InstanceTypeConfigs []InstanceFleetInstanceTypeConfig `pulumi:"instanceTypeConfigs"`
	// Configuration block for launch specification
	LaunchSpecifications *InstanceFleetLaunchSpecifications `pulumi:"launchSpecifications"`
	// Friendly name given to the instance fleet.
	Name                        *string `pulumi:"name"`
	ProvisionedOnDemandCapacity *int    `pulumi:"provisionedOnDemandCapacity"`
	ProvisionedSpotCapacity     *int    `pulumi:"provisionedSpotCapacity"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity *int `pulumi:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity *int `pulumi:"targetSpotCapacity"`
}

type InstanceFleetState struct {
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringPtrInput
	// Configuration block for instance fleet
	InstanceTypeConfigs InstanceFleetInstanceTypeConfigArrayInput
	// Configuration block for launch specification
	LaunchSpecifications InstanceFleetLaunchSpecificationsPtrInput
	// Friendly name given to the instance fleet.
	Name                        pulumi.StringPtrInput
	ProvisionedOnDemandCapacity pulumi.IntPtrInput
	ProvisionedSpotCapacity     pulumi.IntPtrInput
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity pulumi.IntPtrInput
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity pulumi.IntPtrInput
}

func (InstanceFleetState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFleetState)(nil)).Elem()
}

type instanceFleetArgs struct {
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId string `pulumi:"clusterId"`
	// Configuration block for instance fleet
	InstanceTypeConfigs []InstanceFleetInstanceTypeConfig `pulumi:"instanceTypeConfigs"`
	// Configuration block for launch specification
	LaunchSpecifications *InstanceFleetLaunchSpecifications `pulumi:"launchSpecifications"`
	// Friendly name given to the instance fleet.
	Name *string `pulumi:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity *int `pulumi:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity *int `pulumi:"targetSpotCapacity"`
}

// The set of arguments for constructing a InstanceFleet resource.
type InstanceFleetArgs struct {
	// ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
	ClusterId pulumi.StringInput
	// Configuration block for instance fleet
	InstanceTypeConfigs InstanceFleetInstanceTypeConfigArrayInput
	// Configuration block for launch specification
	LaunchSpecifications InstanceFleetLaunchSpecificationsPtrInput
	// Friendly name given to the instance fleet.
	Name pulumi.StringPtrInput
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	TargetOnDemandCapacity pulumi.IntPtrInput
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	TargetSpotCapacity pulumi.IntPtrInput
}

func (InstanceFleetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceFleetArgs)(nil)).Elem()
}
