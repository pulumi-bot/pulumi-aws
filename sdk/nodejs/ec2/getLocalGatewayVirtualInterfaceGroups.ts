// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLocalGatewayVirtualInterfaceGroups(args?: GetLocalGatewayVirtualInterfaceGroupsArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayVirtualInterfaceGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayVirtualInterfaceGroups:getLocalGatewayVirtualInterfaceGroups", {
        "filters": args.filters,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayVirtualInterfaceGroups.
 */
export interface GetLocalGatewayVirtualInterfaceGroupsArgs {
    readonly filters?: inputs.ec2.GetLocalGatewayVirtualInterfaceGroupsFilter[];
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayVirtualInterfaceGroups.
 */
export interface GetLocalGatewayVirtualInterfaceGroupsResult {
    readonly filters?: outputs.ec2.GetLocalGatewayVirtualInterfaceGroupsFilter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids: string[];
    readonly localGatewayVirtualInterfaceIds: string[];
    readonly tags: {[key: string]: string};
}
