// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class SqlInjectionMatchSet extends fabric.Resource {
    public readonly name: fabric.Computed<string>;
    public readonly sqlInjectionMatchTuples?: fabric.Computed<{ fieldToMatch: { data?: string, type: string }[], textTransformation: string }[]>;

    constructor(urnName: string, args?: SqlInjectionMatchSetArgs, dependsOn?: fabric.Resource[]) {
        super("aws:waf/sqlInjectionMatchSet:SqlInjectionMatchSet", urnName, {
            "name": args.name,
            "sqlInjectionMatchTuples": args.sqlInjectionMatchTuples,
        }, dependsOn);
    }
}

export interface SqlInjectionMatchSetArgs {
    readonly name?: fabric.ComputedValue<string>;
    readonly sqlInjectionMatchTuples?: fabric.ComputedValue<{ fieldToMatch: fabric.ComputedValue<{ data?: fabric.ComputedValue<string>, type: fabric.ComputedValue<string> }>[], textTransformation: fabric.ComputedValue<string> }>[];
}

