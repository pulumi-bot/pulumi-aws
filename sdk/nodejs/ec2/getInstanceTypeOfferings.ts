// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Information about EC2 Instance Type Offerings.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getInstanceTypeOfferings({
 *     filters: [
 *         {
 *             name: "instance-type",
 *             values: [
 *                 "t2.micro",
 *                 "t3.micro",
 *             ],
 *         },
 *         {
 *             name: "location",
 *             values: ["usw2-az4"],
 *         },
 *     ],
 *     locationType: "availability-zone-id",
 * }, { async: true }));
 * ```
 */
export function getInstanceTypeOfferings(args?: GetInstanceTypeOfferingsArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypeOfferingsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getInstanceTypeOfferings:getInstanceTypeOfferings", {
        "filters": args.filters,
        "locationType": args.locationType,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypeOfferings.
 */
export interface GetInstanceTypeOfferingsArgs {
    /**
     * One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
     */
    readonly filters?: inputs.ec2.GetInstanceTypeOfferingsFilter[];
    /**
     * Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
     */
    readonly locationType?: string;
}

/**
 * A collection of values returned by getInstanceTypeOfferings.
 */
export interface GetInstanceTypeOfferingsResult {
    readonly filters?: outputs.ec2.GetInstanceTypeOfferingsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Set of EC2 Instance Types.
     */
    readonly instanceTypes: string[];
    readonly locationType?: string;
}
