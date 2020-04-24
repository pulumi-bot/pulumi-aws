// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Availability Zones data source allows access to the list of AWS
// Availability Zones which can be accessed by an AWS account within the region
// configured in the provider.
//
// This is different from the `.getAvailabilityZone` (singular) data source,
// which provides some details about a specific availability zone.
//
// > When [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) are enabled in a region, by default the API and this data source include both Local Zones and Availability Zones. To return only Availability Zones, see the example section below.
func GetAvailabilityZones(ctx *pulumi.Context, args *GetAvailabilityZonesArgs, opts ...pulumi.InvokeOption) (*GetAvailabilityZonesResult, error) {
	var rv GetAvailabilityZonesResult
	err := ctx.Invoke("aws:index/getAvailabilityZones:getAvailabilityZones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAvailabilityZones.
type GetAvailabilityZonesArgs struct {
	// Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
	AllAvailabilityZones *bool `pulumi:"allAvailabilityZones"`
	// List of blacklisted Availability Zone names.
	BlacklistedNames []string `pulumi:"blacklistedNames"`
	// List of blacklisted Availability Zone IDs.
	BlacklistedZoneIds []string `pulumi:"blacklistedZoneIds"`
	// Configuration block(s) for filtering. Detailed below.
	Filters    []GetAvailabilityZonesFilter `pulumi:"filters"`
	GroupNames []string                     `pulumi:"groupNames"`
	// Allows to filter list of Availability Zones based on their
	// current state. Can be either `"available"`, `"information"`, `"impaired"` or
	// `"unavailable"`. By default the list includes a complete set of Availability Zones
	// to which the underlying AWS account has access, regardless of their state.
	State *string `pulumi:"state"`
}

// A collection of values returned by getAvailabilityZones.
type GetAvailabilityZonesResult struct {
	AllAvailabilityZones *bool                        `pulumi:"allAvailabilityZones"`
	BlacklistedNames     []string                     `pulumi:"blacklistedNames"`
	BlacklistedZoneIds   []string                     `pulumi:"blacklistedZoneIds"`
	Filters              []GetAvailabilityZonesFilter `pulumi:"filters"`
	GroupNames           []string                     `pulumi:"groupNames"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Availability Zone names available to the account.
	Names []string `pulumi:"names"`
	State *string  `pulumi:"state"`
	// A list of the Availability Zone IDs available to the account.
	ZoneIds []string `pulumi:"zoneIds"`
}
