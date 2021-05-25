// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides details about resource tagging.
 *
 * ## Example Usage
 * ### Get All Resource Tag Mappings
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.resourcegroupstaggingapi.getResources({ async: true }));
 * ```
 * ### Filter By Tag Key and Value
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.resourcegroupstaggingapi.getResources({
 *     tagFilters: [{
 *         key: "tag-key",
 *         values: [
 *             "tag-value-1",
 *             "tag-value-2",
 *         ],
 *     }],
 * }, { async: true }));
 * ```
 * ### Filter By Resource Type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.resourcegroupstaggingapi.getResources({
 *     resourceTypeFilters: ["ec2:instance"],
 * }, { async: true }));
 * ```
 */
export function getResources(args?: GetResourcesArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:resourcegroupstaggingapi/getResources:getResources", {
        "excludeCompliantResources": args.excludeCompliantResources,
        "includeComplianceDetails": args.includeComplianceDetails,
        "resourceArnLists": args.resourceArnLists,
        "resourceTypeFilters": args.resourceTypeFilters,
        "tagFilters": args.tagFilters,
    }, opts);
}

/**
 * A collection of arguments for invoking getResources.
 */
export interface GetResourcesArgs {
    /**
     * Specifies whether to exclude resources that are compliant with the tag policy. You can use this parameter only if the `includeComplianceDetails` argument is also set to `true`.
     */
    excludeCompliantResources?: boolean;
    /**
     * Specifies whether to include details regarding the compliance with the effective tag policy.
     */
    includeComplianceDetails?: boolean;
    /**
     * Specifies a list of ARNs of resources for which you want to retrieve tag data. Conflicts with `filter`.
     */
    resourceArnLists?: string[];
    /**
     * The constraints on the resources that you want returned. The format of each resource type is `service:resourceType`. For example, specifying a resource type of `ec2` returns all Amazon EC2 resources (which includes EC2 instances). Specifying a resource type of `ec2:instance` returns only EC2 instances.
     */
    resourceTypeFilters?: string[];
    /**
     * Specifies a list of Tag Filters (keys and values) to restrict the output to only those resources that have the specified tag and, if included, the specified value. See Tag Filter below. Conflicts with `resourceArnList`.
     */
    tagFilters?: inputs.resourcegroupstaggingapi.GetResourcesTagFilter[];
}

/**
 * A collection of values returned by getResources.
 */
export interface GetResourcesResult {
    readonly excludeCompliantResources?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeComplianceDetails?: boolean;
    readonly resourceArnLists?: string[];
    /**
     * List of objects matching the search criteria.
     */
    readonly resourceTagMappingLists: outputs.resourcegroupstaggingapi.GetResourcesResourceTagMappingList[];
    readonly resourceTypeFilters?: string[];
    readonly tagFilters?: outputs.resourcegroupstaggingapi.GetResourcesTagFilter[];
}
