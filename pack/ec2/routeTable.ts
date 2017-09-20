// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class RouteTable extends fabric.Resource {
    public readonly propagatingVgws: fabric.Computed<string[]>;
    public readonly route: fabric.Computed<{ cidrBlock?: string, egressOnlyGatewayId?: string, gatewayId?: string, instanceId?: string, ipv6CidrBlock?: string, natGatewayId?: string, networkInterfaceId?: string, vpcPeeringConnectionId?: string }[]>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public readonly vpcId: fabric.Computed<string>;

    constructor(urnName: string, args: RouteTableArgs, dependsOn?: fabric.Resource[]) {
        if (args.vpcId === undefined) {
            throw new Error("Missing required property 'vpcId'");
        }
        super("aws:ec2/routeTable:RouteTable", urnName, {
            "propagatingVgws": args.propagatingVgws,
            "route": args.route,
            "tags": args.tags,
            "vpcId": args.vpcId,
        }, dependsOn);
    }
}

export interface RouteTableArgs {
    readonly propagatingVgws?: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly route?: fabric.ComputedValue<{ cidrBlock?: fabric.ComputedValue<string>, egressOnlyGatewayId?: fabric.ComputedValue<string>, gatewayId?: fabric.ComputedValue<string>, instanceId?: fabric.ComputedValue<string>, ipv6CidrBlock?: fabric.ComputedValue<string>, natGatewayId?: fabric.ComputedValue<string>, networkInterfaceId?: fabric.ComputedValue<string>, vpcPeeringConnectionId?: fabric.ComputedValue<string> }>[];
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
    readonly vpcId: fabric.ComputedValue<string>;
}

