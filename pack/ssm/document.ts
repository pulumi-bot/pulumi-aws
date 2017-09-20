// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Document extends fabric.Resource {
    public /*out*/ readonly arn: fabric.Computed<string>;
    public readonly content: fabric.Computed<string>;
    public /*out*/ readonly createdDate: fabric.Computed<string>;
    public /*out*/ readonly defaultVersion: fabric.Computed<string>;
    public /*out*/ readonly description: fabric.Computed<string>;
    public readonly documentType: fabric.Computed<string>;
    public /*out*/ readonly hash: fabric.Computed<string>;
    public /*out*/ readonly hashType: fabric.Computed<string>;
    public /*out*/ readonly latestVersion: fabric.Computed<string>;
    public readonly name: fabric.Computed<string>;
    public /*out*/ readonly owner: fabric.Computed<string>;
    public /*out*/ readonly parameter: fabric.Computed<{ defaultValue?: string, description?: string, name?: string, type?: string }[]>;
    public readonly permissions?: fabric.Computed<{[key: string]: { accountIds: string, type: string }}>;
    public /*out*/ readonly platformTypes: fabric.Computed<string[]>;
    public /*out*/ readonly schemaVersion: fabric.Computed<string>;
    public /*out*/ readonly status: fabric.Computed<string>;

    constructor(urnName: string, args: DocumentArgs, dependsOn?: fabric.Resource[]) {
        if (args.content === undefined) {
            throw new Error("Missing required property 'content'");
        }
        if (args.documentType === undefined) {
            throw new Error("Missing required property 'documentType'");
        }
        super("aws:ssm/document:Document", urnName, {
            "content": args.content,
            "documentType": args.documentType,
            "name": args.name,
            "permissions": args.permissions,
            "arn": undefined,
            "createdDate": undefined,
            "defaultVersion": undefined,
            "description": undefined,
            "hash": undefined,
            "hashType": undefined,
            "latestVersion": undefined,
            "owner": undefined,
            "parameter": undefined,
            "platformTypes": undefined,
            "schemaVersion": undefined,
            "status": undefined,
        }, dependsOn);
    }
}

export interface DocumentArgs {
    readonly content: fabric.ComputedValue<string>;
    readonly documentType: fabric.ComputedValue<string>;
    readonly name?: fabric.ComputedValue<string>;
    readonly permissions?: fabric.ComputedValue<{[key: string]: { accountIds: fabric.ComputedValue<string>, type: fabric.ComputedValue<string> }}>;
}

