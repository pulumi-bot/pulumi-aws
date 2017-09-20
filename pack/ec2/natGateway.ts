// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class NatGateway extends fabric.Resource {
    public readonly allocationId: fabric.Computed<string>;
    public readonly networkInterfaceId: fabric.Computed<string>;
    public readonly privateIp: fabric.Computed<string>;
    public readonly publicIp: fabric.Computed<string>;
    public readonly subnetId: fabric.Computed<string>;

    constructor(urnName: string, args: NatGatewayArgs, dependsOn?: fabric.Resource[]) {
        if (args.allocationId === undefined) {
            throw new Error("Missing required property 'allocationId'");
        }
        if (args.subnetId === undefined) {
            throw new Error("Missing required property 'subnetId'");
        }
        super("aws:ec2/natGateway:NatGateway", urnName, {
            "allocationId": args.allocationId,
            "networkInterfaceId": args.networkInterfaceId,
            "privateIp": args.privateIp,
            "publicIp": args.publicIp,
            "subnetId": args.subnetId,
        }, dependsOn);
    }
}

export interface NatGatewayArgs {
    readonly allocationId: fabric.ComputedValue<string>;
    readonly networkInterfaceId?: fabric.ComputedValue<string>;
    readonly privateIp?: fabric.ComputedValue<string>;
    readonly publicIp?: fabric.ComputedValue<string>;
    readonly subnetId: fabric.ComputedValue<string>;
}

