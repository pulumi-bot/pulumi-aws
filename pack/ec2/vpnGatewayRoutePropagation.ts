// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class VpnGatewayRoutePropagation extends fabric.Resource {
    public readonly routeTableId: fabric.Computed<string>;
    public readonly vpnGatewayId: fabric.Computed<string>;

    constructor(urnName: string, args: VpnGatewayRoutePropagationArgs, dependsOn?: fabric.Resource[]) {
        if (args.routeTableId === undefined) {
            throw new Error("Missing required property 'routeTableId'");
        }
        if (args.vpnGatewayId === undefined) {
            throw new Error("Missing required property 'vpnGatewayId'");
        }
        super("aws:ec2/vpnGatewayRoutePropagation:VpnGatewayRoutePropagation", urnName, {
            "routeTableId": args.routeTableId,
            "vpnGatewayId": args.vpnGatewayId,
        }, dependsOn);
    }
}

export interface VpnGatewayRoutePropagationArgs {
    readonly routeTableId: fabric.ComputedValue<string>;
    readonly vpnGatewayId: fabric.ComputedValue<string>;
}

