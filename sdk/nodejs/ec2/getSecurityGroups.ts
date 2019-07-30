// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/security_groups.html.markdown.
 */
export function getSecurityGroups(args?: GetSecurityGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupsResult> & GetSecurityGroupsResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetSecurityGroupsResult> = pulumi.runtime.invoke("aws:ec2/getSecurityGroups:getSecurityGroups", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroups.
 */
export interface GetSecurityGroupsArgs {
    /**
     * One or more name/value pairs to use as filters. There are
     * several valid keys, for a full reference, check out
     * [describe-security-groups in the AWS CLI reference][1].
     */
    readonly filters?: { name: string, values: string[] }[];
    /**
     * A mapping of tags, each pair of which must exactly match for
     * desired security groups.
     */
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getSecurityGroups.
 */
export interface GetSecurityGroupsResult {
    readonly filters?: { name: string, values: string[] }[];
    /**
     * IDs of the matches security groups.
     */
    readonly ids: string[];
    readonly tags: {[key: string]: any};
    /**
     * The VPC IDs of the matched security groups. The data source's tag or filter *will span VPCs*
     * unless the `vpc-id` filter is also used.
     */
    readonly vpcIds: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
