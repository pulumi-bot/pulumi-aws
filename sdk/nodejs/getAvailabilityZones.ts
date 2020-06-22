// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The Availability Zones data source allows access to the list of AWS
 * Availability Zones which can be accessed by an AWS account within the region
 * configured in the provider.
 *
 * This is different from the `aws.getAvailabilityZone` (singular) data source,
 * which provides some details about a specific availability zone.
 *
 * > When [Local Zones](https://aws.amazon.com/about-aws/global-infrastructure/localzones/) are enabled in a region, by default the API and this data source include both Local Zones and Availability Zones. To return only Availability Zones, see the example section below.
 *
 * ## Example Usage
 * ### By State
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const available = aws.getAvailabilityZones({
 *     state: "available",
 * });
 * const primary = new aws.ec2.Subnet("primary", {availabilityZone: available.then(available => available.names[0])});
 * // ...
 * const secondary = new aws.ec2.Subnet("secondary", {availabilityZone: available.then(available => available.names[1])});
 * // ...
 * ```
 * ### By Filter
 *
 * All Local Zones (regardless of opt-in status):
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.getAvailabilityZones({
 *     allAvailabilityZones: true,
 *     filters: [{
 *         name: "opt-in-status",
 *         values: [
 *             "not-opted-in",
 *             "opted-in",
 *         ],
 *     }],
 * }, { async: true }));
 * ```
 *
 * Only Availability Zones (no Local Zones):
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.getAvailabilityZones({
 *     filters: [{
 *         name: "opt-in-status",
 *         values: ["opt-in-not-required"],
 *     }],
 * }, { async: true }));
 * ```
 */
export function getAvailabilityZones(args?: GetAvailabilityZonesArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailabilityZonesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:index/getAvailabilityZones:getAvailabilityZones", {
        "allAvailabilityZones": args.allAvailabilityZones,
        "blacklistedNames": args.blacklistedNames,
        "blacklistedZoneIds": args.blacklistedZoneIds,
        "filters": args.filters,
        "groupNames": args.groupNames,
        "state": args.state,
    }, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZones.
 */
export interface GetAvailabilityZonesArgs {
    /**
     * Set to `true` to include all Availability Zones and Local Zones regardless of your opt in status.
     */
    readonly allAvailabilityZones?: boolean;
    /**
     * List of blacklisted Availability Zone names.
     */
    readonly blacklistedNames?: string[];
    /**
     * List of blacklisted Availability Zone IDs.
     */
    readonly blacklistedZoneIds?: string[];
    /**
     * Configuration block(s) for filtering. Detailed below.
     */
    readonly filters?: inputs.GetAvailabilityZonesFilter[];
    readonly groupNames?: string[];
    /**
     * Allows to filter list of Availability Zones based on their
     * current state. Can be either `"available"`, `"information"`, `"impaired"` or
     * `"unavailable"`. By default the list includes a complete set of Availability Zones
     * to which the underlying AWS account has access, regardless of their state.
     */
    readonly state?: string;
}

/**
 * A collection of values returned by getAvailabilityZones.
 */
export interface GetAvailabilityZonesResult {
    readonly allAvailabilityZones?: boolean;
    readonly blacklistedNames?: string[];
    readonly blacklistedZoneIds?: string[];
    readonly filters?: outputs.GetAvailabilityZonesFilter[];
    readonly groupNames?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of the Availability Zone names available to the account.
     */
    readonly names: string[];
    readonly state?: string;
    /**
     * A list of the Availability Zone IDs available to the account.
     */
    readonly zoneIds: string[];
}
