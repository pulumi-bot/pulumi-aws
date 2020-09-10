// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getLocalGatewayVirtualInterfaceGroup(args?: GetLocalGatewayVirtualInterfaceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGatewayVirtualInterfaceGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getLocalGatewayVirtualInterfaceGroup:getLocalGatewayVirtualInterfaceGroup", {
        "filters": args.filters,
        "id": args.id,
        "localGatewayId": args.localGatewayId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGatewayVirtualInterfaceGroup.
 */
export interface GetLocalGatewayVirtualInterfaceGroupArgs {
    readonly filters?: inputs.ec2.GetLocalGatewayVirtualInterfaceGroupFilter[];
    readonly id?: string;
    readonly localGatewayId?: string;
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getLocalGatewayVirtualInterfaceGroup.
 */
export interface GetLocalGatewayVirtualInterfaceGroupResult {
    readonly filters?: outputs.ec2.GetLocalGatewayVirtualInterfaceGroupFilter[];
    readonly id: string;
    readonly localGatewayId: string;
    readonly localGatewayVirtualInterfaceIds: string[];
    readonly tags: {[key: string]: string};
}
