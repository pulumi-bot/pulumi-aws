// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

export function getEventCategories(args?: GetEventCategoriesArgs, opts?: pulumi.InvokeOptions): Promise<GetEventCategoriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:rds/getEventCategories:getEventCategories", {
        "sourceType": args.sourceType,
    }, opts);
}

/**
 * A collection of arguments for invoking getEventCategories.
 */
export interface GetEventCategoriesArgs {
    /**
     * The type of source that will be generating the events. Valid options are db-instance, db-security-group, db-parameter-group, db-snapshot, db-cluster or db-cluster-snapshot.
     */
    readonly sourceType?: string;
}

/**
 * A collection of values returned by getEventCategories.
 */
export interface GetEventCategoriesResult {
    /**
     * A list of the event categories.
     */
    readonly eventCategories: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly sourceType?: string;
}
