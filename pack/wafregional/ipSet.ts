// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class IpSet extends fabric.Resource {
    public readonly ipSetDescriptor?: fabric.Computed<{ type: string, value: string }[]>;
    public readonly name: fabric.Computed<string>;

    constructor(urnName: string, args?: IpSetArgs, dependsOn?: fabric.Resource[]) {
        super("aws:wafregional/ipSet:IpSet", urnName, {
            "ipSetDescriptor": args.ipSetDescriptor,
            "name": args.name,
        }, dependsOn);
    }
}

export interface IpSetArgs {
    readonly ipSetDescriptor?: fabric.ComputedValue<{ type: fabric.ComputedValue<string>, value: fabric.ComputedValue<string> }>[];
    readonly name?: fabric.ComputedValue<string>;
}

