// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class VpcPeeringConnectionAccepter extends fabric.Resource {
    public /*out*/ readonly acceptStatus: fabric.Computed<string>;
    public readonly accepter: fabric.Computed<{ allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[]>;
    public readonly autoAccept?: fabric.Computed<boolean>;
    public /*out*/ readonly peerOwnerId: fabric.Computed<string>;
    public /*out*/ readonly peerVpcId: fabric.Computed<string>;
    public readonly requester: fabric.Computed<{ allowClassicLinkToRemoteVpc?: boolean, allowRemoteVpcDnsResolution?: boolean, allowVpcToRemoteClassicLink?: boolean }[]>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public /*out*/ readonly vpcId: fabric.Computed<string>;
    public readonly vpcPeeringConnectionId: fabric.Computed<string>;

    constructor(urnName: string, args: VpcPeeringConnectionAccepterArgs, dependsOn?: fabric.Resource[]) {
        if (args.vpcPeeringConnectionId === undefined) {
            throw new Error("Missing required property 'vpcPeeringConnectionId'");
        }
        super("aws:ec2/vpcPeeringConnectionAccepter:VpcPeeringConnectionAccepter", urnName, {
            "accepter": args.accepter,
            "autoAccept": args.autoAccept,
            "requester": args.requester,
            "tags": args.tags,
            "vpcPeeringConnectionId": args.vpcPeeringConnectionId,
            "acceptStatus": undefined,
            "peerOwnerId": undefined,
            "peerVpcId": undefined,
            "vpcId": undefined,
        }, dependsOn);
    }
}

export interface VpcPeeringConnectionAccepterArgs {
    readonly accepter?: fabric.ComputedValue<{ allowClassicLinkToRemoteVpc?: fabric.ComputedValue<boolean>, allowRemoteVpcDnsResolution?: fabric.ComputedValue<boolean>, allowVpcToRemoteClassicLink?: fabric.ComputedValue<boolean> }>[];
    readonly autoAccept?: fabric.ComputedValue<boolean>;
    readonly requester?: fabric.ComputedValue<{ allowClassicLinkToRemoteVpc?: fabric.ComputedValue<boolean>, allowRemoteVpcDnsResolution?: fabric.ComputedValue<boolean>, allowVpcToRemoteClassicLink?: fabric.ComputedValue<boolean> }>[];
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
    readonly vpcPeeringConnectionId: fabric.ComputedValue<string>;
}

