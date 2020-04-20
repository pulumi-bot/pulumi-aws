// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides information about a DynamoDB table.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const tableName = aws.dynamodb.getTable({
 *     name: "tableName",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/dynamodb_table.html.markdown.
 */
export function getTable(args: GetTableArgs, opts?: pulumi.InvokeOptions): Promise<GetTableResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:dynamodb/getTable:getTable", {
        "name": args.name,
        "serverSideEncryption": args.serverSideEncryption,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getTable.
 */
export interface GetTableArgs {
    /**
     * The name of the DynamoDB table.
     */
    readonly name: string;
    readonly serverSideEncryption?: inputs.dynamodb.GetTableServerSideEncryption;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getTable.
 */
export interface GetTableResult {
    readonly arn: string;
    readonly attributes: outputs.dynamodb.GetTableAttribute[];
    readonly billingMode: string;
    readonly globalSecondaryIndexes: outputs.dynamodb.GetTableGlobalSecondaryIndex[];
    readonly hashKey: string;
    readonly localSecondaryIndexes: outputs.dynamodb.GetTableLocalSecondaryIndex[];
    readonly name: string;
    readonly pointInTimeRecovery: outputs.dynamodb.GetTablePointInTimeRecovery;
    readonly rangeKey: string;
    readonly readCapacity: number;
    readonly replicas: outputs.dynamodb.GetTableReplica[];
    readonly serverSideEncryption: outputs.dynamodb.GetTableServerSideEncryption;
    readonly streamArn: string;
    readonly streamEnabled: boolean;
    readonly streamLabel: string;
    readonly streamViewType: string;
    readonly tags: {[key: string]: any};
    readonly ttl: outputs.dynamodb.GetTableTtl;
    readonly writeCapacity: number;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
