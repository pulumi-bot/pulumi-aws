// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class FlowLog extends fabric.Resource {
    public readonly eniId?: fabric.Computed<string>;
    public readonly iamRoleArn: fabric.Computed<string>;
    public readonly logGroupName: fabric.Computed<string>;
    public readonly subnetId?: fabric.Computed<string>;
    public readonly trafficType: fabric.Computed<string>;
    public readonly vpcId?: fabric.Computed<string>;

    constructor(urnName: string, args: FlowLogArgs, dependsOn?: fabric.Resource[]) {
        if (args.iamRoleArn === undefined) {
            throw new Error("Missing required property 'iamRoleArn'");
        }
        if (args.logGroupName === undefined) {
            throw new Error("Missing required property 'logGroupName'");
        }
        if (args.trafficType === undefined) {
            throw new Error("Missing required property 'trafficType'");
        }
        super("aws:ec2/flowLog:FlowLog", urnName, {
            "eniId": args.eniId,
            "iamRoleArn": args.iamRoleArn,
            "logGroupName": args.logGroupName,
            "subnetId": args.subnetId,
            "trafficType": args.trafficType,
            "vpcId": args.vpcId,
        }, dependsOn);
    }
}

export interface FlowLogArgs {
    readonly eniId?: fabric.ComputedValue<string>;
    readonly iamRoleArn: fabric.ComputedValue<string>;
    readonly logGroupName: fabric.ComputedValue<string>;
    readonly subnetId?: fabric.ComputedValue<string>;
    readonly trafficType: fabric.ComputedValue<string>;
    readonly vpcId?: fabric.ComputedValue<string>;
}

