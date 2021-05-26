// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an RDS subnet group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const database = pulumi.output(aws.rds.getSubnetGroup({
 *     name: "my-test-database-subnet-group",
 * }, { async: true }));
 * ```
 */
export function getSubnetGroup(args: GetSubnetGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getSubnetGroup:getSubnetGroup", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnetGroup.
 */
export interface GetSubnetGroupArgs {
    /**
     * The name of the RDS database subnet group.
     */
    name: string;
}

/**
 * A collection of values returned by getSubnetGroup.
 */
export interface GetSubnetGroupResult {
    /**
     * The Amazon Resource Name (ARN) for the DB subnet group.
     */
    readonly arn: string;
    /**
     * Provides the description of the DB subnet group.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * Provides the status of the DB subnet group.
     */
    readonly status: string;
    /**
     * Contains a list of subnet identifiers.
     */
    readonly subnetIds: string[];
    /**
     * Provides the VPC ID of the subnet group.
     */
    readonly vpcId: string;
}

export function getSubnetGroupApply(args: GetSubnetGroupApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetGroupResult> {
    return pulumi.output(args).apply(a => getSubnetGroup(a, opts))
}

/**
 * A collection of arguments for invoking getSubnetGroup.
 */
export interface GetSubnetGroupApplyArgs {
    /**
     * The name of the RDS database subnet group.
     */
    name: pulumi.Input<string>;
}
