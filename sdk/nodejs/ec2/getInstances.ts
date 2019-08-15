// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/instances.html.markdown.
 */
export function getInstances(args?: GetInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetInstancesResult> & GetInstancesResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetInstancesResult> = pulumi.runtime.invoke("aws:ec2/getInstances:getInstances", {
        "filters": args.filters,
        "instanceStateNames": args.instanceStateNames,
        "instanceTags": args.instanceTags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getInstances.
 */
export interface GetInstancesArgs {
    /**
     * One or more name/value pairs to use as filters. There are
     * several valid keys, for a full reference, check out
     * [describe-instances in the AWS CLI reference][1].
     */
    readonly filters?: inputApi.ec2.GetInstancesFilter[];
    /**
     * A list of instance states that should be applicable to the desired instances. The permitted values are: `pending, running, shutting-down, stopped, stopping, terminated`. The default value is `running`.
     */
    readonly instanceStateNames?: string[];
    /**
     * A mapping of tags, each pair of which must
     * exactly match a pair on desired instances.
     */
    readonly instanceTags?: {[key: string]: any};
}

/**
 * A collection of values returned by getInstances.
 */
export interface GetInstancesResult {
    readonly filters?: outputApi.ec2.GetInstancesFilter[];
    /**
     * IDs of instances found through the filter
     */
    readonly ids: string[];
    readonly instanceStateNames?: string[];
    readonly instanceTags: {[key: string]: any};
    /**
     * Private IP addresses of instances found through the filter
     */
    readonly privateIps: string[];
    /**
     * Public IP addresses of instances found through the filter
     */
    readonly publicIps: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
