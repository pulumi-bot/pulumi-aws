// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Certificate extends fabric.Resource {
    public /*out*/ readonly certificateArn: fabric.Computed<string>;
    public readonly certificateId: fabric.Computed<string>;
    public readonly certificatePem?: fabric.Computed<string>;
    public readonly certificateWallet?: fabric.Computed<string>;

    constructor(urnName: string, args: CertificateArgs, dependsOn?: fabric.Resource[]) {
        if (args.certificateId === undefined) {
            throw new Error("Missing required property 'certificateId'");
        }
        super("aws:dms/certificate:Certificate", urnName, {
            "certificateId": args.certificateId,
            "certificatePem": args.certificatePem,
            "certificateWallet": args.certificateWallet,
            "certificateArn": undefined,
        }, dependsOn);
    }
}

export interface CertificateArgs {
    readonly certificateId: fabric.ComputedValue<string>;
    readonly certificatePem?: fabric.ComputedValue<string>;
    readonly certificateWallet?: fabric.ComputedValue<string>;
}

