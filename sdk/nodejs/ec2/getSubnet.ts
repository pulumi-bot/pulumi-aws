// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * `aws.ec2.Subnet` provides details about a specific VPC subnet.
 *
 * This resource can prove useful when a module accepts a subnet id as
 * an input variable and needs to, for example, determine the id of the
 * VPC that the subnet belongs to.
 *
 * ## Example Usage
 *
 * The following example shows how one might accept a subnet id as a variable
 * and use this data source to obtain the data necessary to create a security
 * group that allows connections from hosts in that subnet.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const subnetId = config.requireObject("subnetId");
 * const selected = aws.ec2.getSubnet({
 *     id: subnetId,
 * });
 * const subnet = new aws.ec2.SecurityGroup("subnet", {
 *     vpcId: selected.then(selected => selected.vpcId),
 *     ingress: [{
 *         cidrBlocks: [selected.then(selected => selected.cidrBlock)],
 *         fromPort: 80,
 *         toPort: 80,
 *         protocol: "tcp",
 *     }],
 * });
 * ```
 */
export function getSubnet(args?: GetSubnetArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ec2/getSubnet:getSubnet", {
        "availabilityZone": args.availabilityZone,
        "availabilityZoneId": args.availabilityZoneId,
        "cidrBlock": args.cidrBlock,
        "defaultForAz": args.defaultForAz,
        "filters": args.filters,
        "id": args.id,
        "ipv6CidrBlock": args.ipv6CidrBlock,
        "state": args.state,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnet.
 */
export interface GetSubnetArgs {
    /**
     * The availability zone where the
     * subnet must reside.
     */
    readonly availabilityZone?: string;
    /**
     * The ID of the Availability Zone for the subnet.
     */
    readonly availabilityZoneId?: string;
    /**
     * The cidr block of the desired subnet.
     */
    readonly cidrBlock?: string;
    /**
     * Boolean constraint for whether the desired
     * subnet must be the default subnet for its associated availability zone.
     */
    readonly defaultForAz?: boolean;
    /**
     * Custom filter block as described below.
     */
    readonly filters?: inputs.ec2.GetSubnetFilter[];
    /**
     * The id of the specific subnet to retrieve.
     */
    readonly id?: string;
    /**
     * The Ipv6 cidr block of the desired subnet
     */
    readonly ipv6CidrBlock?: string;
    /**
     * The state that the desired subnet must have.
     */
    readonly state?: string;
    /**
     * A map of tags, each pair of which must exactly match
     * a pair on the desired subnet.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The id of the VPC that the desired subnet belongs to.
     */
    readonly vpcId?: string;
}

/**
 * A collection of values returned by getSubnet.
 */
export interface GetSubnetResult {
    /**
     * The ARN of the subnet.
     */
    readonly arn: string;
    readonly assignIpv6AddressOnCreation: boolean;
    readonly availabilityZone: string;
    readonly availabilityZoneId: string;
    readonly cidrBlock: string;
    readonly defaultForAz: boolean;
    readonly filters?: outputs.ec2.GetSubnetFilter[];
    readonly id: string;
    readonly ipv6CidrBlock: string;
    readonly ipv6CidrBlockAssociationId: string;
    readonly mapPublicIpOnLaunch: boolean;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    readonly outpostArn: string;
    /**
     * The ID of the AWS account that owns the subnet.
     */
    readonly ownerId: string;
    readonly state: string;
    readonly tags: {[key: string]: string};
    readonly vpcId: string;
}
