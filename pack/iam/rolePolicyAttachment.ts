// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

import {ARN} from "../index";
import {Role} from "./role";

export class RolePolicyAttachment extends fabric.Resource {
    public readonly policyArn: fabric.Computed<ARN>;
    public readonly role: fabric.Computed<Role>;

    constructor(urnName: string, args: RolePolicyAttachmentArgs, dependsOn?: fabric.Resource[]) {
        if (args.policyArn === undefined) {
            throw new Error("Missing required property 'policyArn'");
        }
        if (args.role === undefined) {
            throw new Error("Missing required property 'role'");
        }
        super("aws:iam/rolePolicyAttachment:RolePolicyAttachment", urnName, {
            "policyArn": args.policyArn,
            "role": args.role,
        }, dependsOn);
    }
}

export interface RolePolicyAttachmentArgs {
    readonly policyArn: fabric.ComputedValue<ARN>;
    readonly role: fabric.ComputedValue<Role>;
}

