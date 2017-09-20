// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class VpnConnectionRoute extends fabric.Resource {
    public readonly destinationCidrBlock: fabric.Computed<string>;
    public readonly vpnConnectionId: fabric.Computed<string>;

    constructor(urnName: string, args: VpnConnectionRouteArgs, dependsOn?: fabric.Resource[]) {
        if (args.destinationCidrBlock === undefined) {
            throw new Error("Missing required property 'destinationCidrBlock'");
        }
        if (args.vpnConnectionId === undefined) {
            throw new Error("Missing required property 'vpnConnectionId'");
        }
        super("aws:ec2/vpnConnectionRoute:VpnConnectionRoute", urnName, {
            "destinationCidrBlock": args.destinationCidrBlock,
            "vpnConnectionId": args.vpnConnectionId,
        }, dependsOn);
    }
}

export interface VpnConnectionRouteArgs {
    readonly destinationCidrBlock: fabric.ComputedValue<string>;
    readonly vpnConnectionId: fabric.ComputedValue<string>;
}

