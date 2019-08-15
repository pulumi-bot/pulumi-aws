// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a Network Interface.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const bar = pulumi.output(aws.ec2.getNetworkInterface({
 *     id: "eni-01234567",
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/network_interface.html.markdown.
 */
export function getNetworkInterface(args?: GetNetworkInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkInterfaceResult> & GetNetworkInterfaceResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetNetworkInterfaceResult> = pulumi.runtime.invoke("aws:ec2/getNetworkInterface:getNetworkInterface", {
        "filters": args.filters,
        "id": args.id,
        "tags": args.tags,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getNetworkInterface.
 */
export interface GetNetworkInterfaceArgs {
    /**
     * One or more name/value pairs to filter off of. There are several valid keys, for a full reference, check out [describe-network-interfaces](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-network-interfaces.html) in the AWS CLI reference.
     */
    readonly filters?: inputApi.ec2.GetNetworkInterfaceFilter[];
    /**
     * The identifier for the network interface.
     */
    readonly id?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getNetworkInterface.
 */
export interface GetNetworkInterfaceResult {
    /**
     * The association information for an Elastic IP address (IPv4) associated with the network interface. See supported fields below.
     */
    readonly associations: outputApi.ec2.GetNetworkInterfaceAssociation[];
    readonly attachments: outputApi.ec2.GetNetworkInterfaceAttachment[];
    /**
     * The Availability Zone.
     */
    readonly availabilityZone: string;
    /**
     * Description of the network interface.
     */
    readonly description: string;
    readonly filters?: outputApi.ec2.GetNetworkInterfaceFilter[];
    readonly id: string;
    /**
     * The type of interface.
     */
    readonly interfaceType: string;
    /**
     * List of IPv6 addresses to assign to the ENI.
     */
    readonly ipv6Addresses: string[];
    /**
     * The MAC address.
     */
    readonly macAddress: string;
    /**
     * The AWS account ID of the owner of the network interface.
     */
    readonly ownerId: string;
    /**
     * The private DNS name.
     */
    readonly privateDnsName: string;
    /**
     * The private IPv4 address of the network interface within the subnet.
     */
    readonly privateIp: string;
    /**
     * The private IPv4 addresses associated with the network interface.
     */
    readonly privateIps: string[];
    /**
     * The ID of the entity that launched the instance on your behalf.
     */
    readonly requesterId: string;
    /**
     * The list of security groups for the network interface.
     */
    readonly securityGroups: string[];
    /**
     * The ID of the subnet.
     */
    readonly subnetId: string;
    /**
     * Any tags assigned to the network interface.
     */
    readonly tags: {[key: string]: any};
    /**
     * The ID of the VPC.
     */
    readonly vpcId: string;
}
