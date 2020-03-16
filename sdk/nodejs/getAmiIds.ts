// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get a list of AMI IDs matching the specified criteria.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ubuntu = aws.getAmiIds({
 *     filters: [{
 *         name: "name",
 *         values: ["ubuntu/images/ubuntu-*-*-amd64-server-*"],
 *     }],
 *     owners: ["099720109477"],
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ami_ids.html.markdown.
 */
export function getAmiIds(args: GetAmiIdsArgs, opts?: pulumi.InvokeOptions): Promise<GetAmiIdsResult> & GetAmiIdsResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetAmiIdsResult> = pulumi.runtime.invoke("aws:index/getAmiIds:getAmiIds", {
        "executableUsers": args.executableUsers,
        "filters": args.filters,
        "nameRegex": args.nameRegex,
        "owners": args.owners,
        "sortAscending": args.sortAscending,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getAmiIds.
 */
export interface GetAmiIdsArgs {
    /**
     * Limit search to users with *explicit* launch
     * permission on  the image. Valid items are the numeric account ID or `self`.
     */
    readonly executableUsers?: string[];
    /**
     * One or more name/value pairs to filter off of. There
     * are several valid keys, for a full reference, check out
     * [describe-images in the AWS CLI reference][1].
     */
    readonly filters?: inputs.GetAmiIdsFilter[];
    /**
     * A regex string to apply to the AMI list returned
     * by AWS. This allows more advanced filtering not supported from the AWS API.
     * This filtering is done locally on what AWS returns, and could have a performance
     * impact if the result is large. It is recommended to combine this with other
     * options to narrow down the list AWS returns.
     */
    readonly nameRegex?: string;
    /**
     * List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
     */
    readonly owners: string[];
    /**
     * Used to sort AMIs by creation time.
     */
    readonly sortAscending?: boolean;
}

/**
 * A collection of values returned by getAmiIds.
 */
export interface GetAmiIdsResult {
    readonly executableUsers?: string[];
    readonly filters?: outputs.GetAmiIdsFilter[];
    readonly ids: string[];
    readonly nameRegex?: string;
    readonly owners: string[];
    readonly sortAscending?: boolean;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
