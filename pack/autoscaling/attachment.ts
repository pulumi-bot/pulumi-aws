// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Attachment extends fabric.Resource {
    public readonly albTargetGroupArn?: fabric.Computed<string>;
    public readonly autoscalingGroupName: fabric.Computed<string>;
    public readonly elb?: fabric.Computed<string>;

    constructor(urnName: string, args: AttachmentArgs, dependsOn?: fabric.Resource[]) {
        if (args.autoscalingGroupName === undefined) {
            throw new Error("Missing required property 'autoscalingGroupName'");
        }
        super("aws:autoscaling/attachment:Attachment", urnName, {
            "albTargetGroupArn": args.albTargetGroupArn,
            "autoscalingGroupName": args.autoscalingGroupName,
            "elb": args.elb,
        }, dependsOn);
    }
}

export interface AttachmentArgs {
    readonly albTargetGroupArn?: fabric.ComputedValue<string>;
    readonly autoscalingGroupName: fabric.ComputedValue<string>;
    readonly elb?: fabric.ComputedValue<string>;
}

