// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class OriginAccessIdentity extends fabric.Resource {
    public /*out*/ readonly callerReference: fabric.Computed<string>;
    public /*out*/ readonly cloudfrontAccessIdentityPath: fabric.Computed<string>;
    public readonly comment?: fabric.Computed<string>;
    public /*out*/ readonly etag: fabric.Computed<string>;
    public /*out*/ readonly iamArn: fabric.Computed<string>;
    public /*out*/ readonly s3CanonicalUserId: fabric.Computed<string>;

    constructor(urnName: string, args?: OriginAccessIdentityArgs, dependsOn?: fabric.Resource[]) {
        super("aws:cloudfront/originAccessIdentity:OriginAccessIdentity", urnName, {
            "comment": args.comment,
            "callerReference": undefined,
            "cloudfrontAccessIdentityPath": undefined,
            "etag": undefined,
            "iamArn": undefined,
            "s3CanonicalUserId": undefined,
        }, dependsOn);
    }
}

export interface OriginAccessIdentityArgs {
    readonly comment?: fabric.ComputedValue<string>;
}

