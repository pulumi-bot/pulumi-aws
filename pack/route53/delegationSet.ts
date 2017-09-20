// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class DelegationSet extends fabric.Resource {
    public /*out*/ readonly nameServers: fabric.Computed<string[]>;
    public readonly referenceName?: fabric.Computed<string>;

    constructor(urnName: string, args?: DelegationSetArgs, dependsOn?: fabric.Resource[]) {
        super("aws:route53/delegationSet:DelegationSet", urnName, {
            "referenceName": args.referenceName,
            "nameServers": undefined,
        }, dependsOn);
    }
}

export interface DelegationSetArgs {
    readonly referenceName?: fabric.ComputedValue<string>;
}

