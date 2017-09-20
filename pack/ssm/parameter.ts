// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Parameter extends fabric.Resource {
    public readonly keyId?: fabric.Computed<string>;
    public readonly name: fabric.Computed<string>;
    public readonly overwrite?: fabric.Computed<boolean>;
    public readonly type: fabric.Computed<string>;
    public readonly value: fabric.Computed<string>;

    constructor(urnName: string, args: ParameterArgs, dependsOn?: fabric.Resource[]) {
        if (args.type === undefined) {
            throw new Error("Missing required property 'type'");
        }
        if (args.value === undefined) {
            throw new Error("Missing required property 'value'");
        }
        super("aws:ssm/parameter:Parameter", urnName, {
            "keyId": args.keyId,
            "name": args.name,
            "overwrite": args.overwrite,
            "type": args.type,
            "value": args.value,
        }, dependsOn);
    }
}

export interface ParameterArgs {
    readonly keyId?: fabric.ComputedValue<string>;
    readonly name?: fabric.ComputedValue<string>;
    readonly overwrite?: fabric.ComputedValue<boolean>;
    readonly type: fabric.ComputedValue<string>;
    readonly value: fabric.ComputedValue<string>;
}

