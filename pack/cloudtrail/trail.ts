// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Trail extends fabric.Resource {
    public /*out*/ readonly arn: fabric.Computed<string>;
    public readonly cloudWatchLogsGroupArn?: fabric.Computed<string>;
    public readonly cloudWatchLogsRoleArn?: fabric.Computed<string>;
    public readonly enableLogFileValidation?: fabric.Computed<boolean>;
    public readonly enableLogging?: fabric.Computed<boolean>;
    public /*out*/ readonly homeRegion: fabric.Computed<string>;
    public readonly includeGlobalServiceEvents?: fabric.Computed<boolean>;
    public readonly isMultiRegionTrail?: fabric.Computed<boolean>;
    public readonly kmsKeyId?: fabric.Computed<string>;
    public readonly name: fabric.Computed<string>;
    public readonly s3BucketName: fabric.Computed<string>;
    public readonly s3KeyPrefix?: fabric.Computed<string>;
    public readonly snsTopicName?: fabric.Computed<string>;
    public readonly tags?: fabric.Computed<{[key: string]: any}>;

    constructor(urnName: string, args: TrailArgs, dependsOn?: fabric.Resource[]) {
        if (args.s3BucketName === undefined) {
            throw new Error("Missing required property 's3BucketName'");
        }
        super("aws:cloudtrail/trail:Trail", urnName, {
            "cloudWatchLogsGroupArn": args.cloudWatchLogsGroupArn,
            "cloudWatchLogsRoleArn": args.cloudWatchLogsRoleArn,
            "enableLogFileValidation": args.enableLogFileValidation,
            "enableLogging": args.enableLogging,
            "includeGlobalServiceEvents": args.includeGlobalServiceEvents,
            "isMultiRegionTrail": args.isMultiRegionTrail,
            "kmsKeyId": args.kmsKeyId,
            "name": args.name,
            "s3BucketName": args.s3BucketName,
            "s3KeyPrefix": args.s3KeyPrefix,
            "snsTopicName": args.snsTopicName,
            "tags": args.tags,
            "arn": undefined,
            "homeRegion": undefined,
        }, dependsOn);
    }
}

export interface TrailArgs {
    readonly cloudWatchLogsGroupArn?: fabric.ComputedValue<string>;
    readonly cloudWatchLogsRoleArn?: fabric.ComputedValue<string>;
    readonly enableLogFileValidation?: fabric.ComputedValue<boolean>;
    readonly enableLogging?: fabric.ComputedValue<boolean>;
    readonly includeGlobalServiceEvents?: fabric.ComputedValue<boolean>;
    readonly isMultiRegionTrail?: fabric.ComputedValue<boolean>;
    readonly kmsKeyId?: fabric.ComputedValue<string>;
    readonly name?: fabric.ComputedValue<string>;
    readonly s3BucketName: fabric.ComputedValue<string>;
    readonly s3KeyPrefix?: fabric.ComputedValue<string>;
    readonly snsTopicName?: fabric.ComputedValue<string>;
    readonly tags?: fabric.ComputedValue<{[key: string]: any}>;
}

