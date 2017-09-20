// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Subnet extends fabric.Resource {
    public readonly assignIpv6AddressOnCreation?: fabric.Computed<boolean>;
    public readonly availabilityZone: fabric.Computed<string>;
    public readonly cidrBlock: fabric.Computed<string>;
    public readonly ipv6CidrBlock: fabric.Computed<string>;
    public /*out*/ readonly ipv6CidrBlockAssociationId: fabric.Computed<string>;
    public readonly mapPublicIpOnLaunch?: fabric.Computed<boolean>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public readonly vpcId: fabric.Computed<string>;

    constructor(urnName: string, args: SubnetArgs, dependsOn?: fabric.Resource[]) {
        if (args.cidrBlock === undefined) {
            throw new Error("Missing required property 'cidrBlock'");
        }
        if (args.vpcId === undefined) {
            throw new Error("Missing required property 'vpcId'");
        }
        super("aws:ec2/subnet:Subnet", urnName, {
            "assignIpv6AddressOnCreation": args.assignIpv6AddressOnCreation,
            "availabilityZone": args.availabilityZone,
            "cidrBlock": args.cidrBlock,
            "ipv6CidrBlock": args.ipv6CidrBlock,
            "mapPublicIpOnLaunch": args.mapPublicIpOnLaunch,
            "tags": args.tags,
            "vpcId": args.vpcId,
            "ipv6CidrBlockAssociationId": undefined,
        }, dependsOn);
    }
}

export interface SubnetArgs {
    readonly assignIpv6AddressOnCreation?: fabric.ComputedValue<boolean>;
    readonly availabilityZone?: fabric.ComputedValue<string>;
    readonly cidrBlock: fabric.ComputedValue<string>;
    readonly ipv6CidrBlock?: fabric.ComputedValue<string>;
    readonly mapPublicIpOnLaunch?: fabric.ComputedValue<boolean>;
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
    readonly vpcId: fabric.ComputedValue<string>;
}

