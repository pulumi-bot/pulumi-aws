// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class ReplicationSubnetGroup extends fabric.Resource {
    public /*out*/ readonly replicationSubnetGroupArn: fabric.Computed<string>;
    public readonly replicationSubnetGroupDescription: fabric.Computed<string>;
    public readonly replicationSubnetGroupId: fabric.Computed<string>;
    public readonly subnetIds: fabric.Computed<string[]>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;
    public /*out*/ readonly vpcId: fabric.Computed<string>;

    constructor(urnName: string, args: ReplicationSubnetGroupArgs, dependsOn?: fabric.Resource[]) {
        if (args.replicationSubnetGroupDescription === undefined) {
            throw new Error("Missing required property 'replicationSubnetGroupDescription'");
        }
        if (args.replicationSubnetGroupId === undefined) {
            throw new Error("Missing required property 'replicationSubnetGroupId'");
        }
        if (args.subnetIds === undefined) {
            throw new Error("Missing required property 'subnetIds'");
        }
        super("aws:dms/replicationSubnetGroup:ReplicationSubnetGroup", urnName, {
            "replicationSubnetGroupDescription": args.replicationSubnetGroupDescription,
            "replicationSubnetGroupId": args.replicationSubnetGroupId,
            "subnetIds": args.subnetIds,
            "tags": args.tags,
            "replicationSubnetGroupArn": undefined,
            "vpcId": undefined,
        }, dependsOn);
    }
}

export interface ReplicationSubnetGroupArgs {
    readonly replicationSubnetGroupDescription: fabric.ComputedValue<string>;
    readonly replicationSubnetGroupId: fabric.ComputedValue<string>;
    readonly subnetIds: fabric.ComputedValue<fabric.ComputedValue<string>>[];
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
}

