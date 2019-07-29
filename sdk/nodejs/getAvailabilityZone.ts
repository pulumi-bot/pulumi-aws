// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * `aws_availability_zone` provides details about a specific availability zone (AZ)
 * in the current region.
 * 
 * This can be used both to validate an availability zone given in a variable
 * and to split the AZ name into its component parts of an AWS region and an
 * AZ identifier letter. The latter may be useful e.g. for implementing a
 * consistent subnet numbering scheme across several regions by mapping both
 * the region and the subnet letter to network numbers.
 * 
 * This is different from the `aws_availability_zones` (plural) data source,
 * which provides a list of the available zones.
 * 
 * ## Example Usage
 * 
 * The following example shows how this data source might be used to derive
 * VPC and subnet CIDR prefixes systematically for an availability zone.
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const azNumber = config.get("azNumber") || {
 *     a: 1,
 *     b: 2,
 *     c: 3,
 *     d: 4,
 *     e: 5,
 *     f: 6,
 * };
 * const regionNumber = config.get("regionNumber") || {
 *     "ap-northeast-1": 5,
 *     "eu-central-1": 4,
 *     "us-east-1": 1,
 *     "us-west-1": 2,
 *     "us-west-2": 3,
 * };
 * 
 * // Retrieve the AZ where we want to create network resources
 * // This must be in the region selected on the AWS provider.
 * const exampleAvailabilityZone = pulumi.output(aws.getAvailabilityZone({
 *     name: "eu-central-1a",
 * }));
 * // Create a VPC for the region associated with the AZ
 * const exampleVpc = new aws.ec2.Vpc("example", {
 *     cidrBlock: exampleAvailabilityZone.apply(exampleAvailabilityZone => (() => {
 *         throw "tf2pulumi error: NYI: call to cidrsubnet";
 *         return (() => { throw "NYI: call to cidrsubnet"; })();
 *     })()),
 * });
 * // Create a subnet for the AZ within the regional VPC
 * const exampleSubnet = new aws.ec2.Subnet("example", {
 *     cidrBlock: pulumi.all([exampleVpc.cidrBlock, exampleAvailabilityZone]).apply(([cidrBlock, exampleAvailabilityZone]) => (() => {
 *         throw "tf2pulumi error: NYI: call to cidrsubnet";
 *         return (() => { throw "NYI: call to cidrsubnet"; })();
 *     })()),
 *     vpcId: exampleVpc.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/availability_zone.html.markdown.
 */
export function getAvailabilityZone(args?: GetAvailabilityZoneArgs, opts?: pulumi.InvokeOptions): Promise<GetAvailabilityZoneResult> & GetAvailabilityZoneResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAvailabilityZoneResult> = pulumi.runtime.invoke("aws:index/getAvailabilityZone:getAvailabilityZone", {
        "name": args.name,
        "state": args.state,
        "zoneId": args.zoneId,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAvailabilityZone.
 */
export interface GetAvailabilityZoneArgs {
    /**
     * The full name of the availability zone to select.
     */
    readonly name?: string;
    /**
     * A specific availability zone state to require. May
     * be any of `"available"`, `"information"` or `"impaired"`.
     */
    readonly state?: string;
    /**
     * The zone ID of the availability zone to select.
     */
    readonly zoneId?: string;
}

/**
 * A collection of values returned by getAvailabilityZone.
 */
export interface GetAvailabilityZoneResult {
    /**
     * The name of the selected availability zone.
     */
    readonly name: string;
    /**
     * The part of the AZ name that appears after the region name,
     * uniquely identifying the AZ within its region.
     */
    readonly nameSuffix: string;
    /**
     * The region where the selected availability zone resides.
     * This is always the region selected on the provider, since this data source
     * searches only within that region.
     */
    readonly region: string;
    /**
     * The current state of the AZ.
     */
    readonly state: string;
    /**
     * (Optional) The zone ID of the selected availability zone.
     */
    readonly zoneId: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
