// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides details about a specific Nat Gateway.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const subnetId = config.require("subnetId");
 *
 * const defaultNatGateway = aws_subnet_public.id.apply(id => aws.ec2.getNatGateway({
 *     subnetId: id,
 * }, { async: true }));
 * ```
 */
export function getNatGateway(args?: GetNatGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetNatGatewayResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getNatGateway:getNatGateway", {
        "filters": args.filters,
        "id": args.id,
        "state": args.state,
        "subnetId": args.subnetId,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNatGateway.
 */
export interface GetNatGatewayArgs {
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetNatGatewayFilter[];
    /**
     * The id of the specific Nat Gateway to retrieve.
     */
    readonly id?: string;
    /**
     * The state of the NAT gateway (pending | failed | available | deleting | deleted ).
     */
    readonly state?: string;
    /**
     * The id of subnet that the Nat Gateway resides in.
     */
    readonly subnetId?: string;
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired Nat Gateway.
     */
    readonly tags?: {[key: string]: any};
    /**
     * The id of the VPC that the Nat Gateway resides in.
     */
    readonly vpcId?: string;
}

/**
 * A collection of values returned by getNatGateway.
 */
export interface GetNatGatewayResult {
    /**
     * The Id of the EIP allocated to the selected Nat Gateway.
     */
    readonly allocationId: string;
    readonly filters?: outputs.ec2.GetNatGatewayFilter[];
    readonly id: string;
    /**
     * The Id of the ENI allocated to the selected Nat Gateway.
     */
    readonly networkInterfaceId: string;
    /**
     * The private Ip address of the selected Nat Gateway.
     */
    readonly privateIp: string;
    /**
     * The public Ip (EIP) address of the selected Nat Gateway.
     */
    readonly publicIp: string;
    readonly state: string;
    readonly subnetId: string;
    readonly tags: {[key: string]: any};
    readonly vpcId: string;
}
