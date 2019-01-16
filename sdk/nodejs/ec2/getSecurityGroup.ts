// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSecurityGroup(args?: GetSecurityGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:ec2/getSecurityGroup:getSecurityGroup", {
        "filters": args.filters,
        "id": args.id,
        "name": args.name,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroup.
 */
export interface GetSecurityGroupArgs {
    readonly filters?: { name: string, values: string[] }[];
    readonly id?: string;
    readonly name?: string;
    readonly tags?: {[key: string]: any};
    readonly vpcId?: string;
}

/**
 * A collection of values returned by getSecurityGroup.
 */
export interface GetSecurityGroupResult {
    readonly arn: string;
    readonly description: string;
    readonly id: string;
    readonly name: string;
    readonly tags: {[key: string]: any};
    readonly vpcId: string;
}
