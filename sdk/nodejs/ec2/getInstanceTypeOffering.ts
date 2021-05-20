// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Information about single EC2 Instance Type Offering.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.ec2.getInstanceTypeOffering({
 *     filters: [{
 *         name: "instance-type",
 *         values: [
 *             "t2.micro",
 *             "t3.micro",
 *         ],
 *     }],
 *     preferredInstanceTypes: [
 *         "t3.micro",
 *         "t2.micro",
 *     ],
 * }));
 * ```
 */
export function getInstanceTypeOffering(args?: GetInstanceTypeOfferingArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceTypeOfferingResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getInstanceTypeOffering:getInstanceTypeOffering", {
        "filters": args.filters,
        "locationType": args.locationType,
        "preferredInstanceTypes": args.preferredInstanceTypes,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstanceTypeOffering.
 */
export interface GetInstanceTypeOfferingArgs {
    /**
     * One or more configuration blocks containing name-values filters. See the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstanceTypeOfferings.html) for supported filters. Detailed below.
     */
    readonly filters?: inputs.ec2.GetInstanceTypeOfferingFilter[];
    /**
     * Location type. Defaults to `region`. Valid values: `availability-zone`, `availability-zone-id`, and `region`.
     */
    readonly locationType?: string;
    /**
     * Ordered list of preferred EC2 Instance Types. The first match in this list will be returned. If no preferred matches are found and the original search returned more than one result, an error is returned.
     */
    readonly preferredInstanceTypes?: string[];
}

/**
 * A collection of values returned by getInstanceTypeOffering.
 */
export interface GetInstanceTypeOfferingResult {
    readonly filters?: outputs.ec2.GetInstanceTypeOfferingFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * EC2 Instance Type.
     */
    readonly instanceType: string;
    readonly locationType?: string;
    readonly preferredInstanceTypes?: string[];
}
