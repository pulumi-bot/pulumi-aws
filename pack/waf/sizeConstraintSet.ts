// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class SizeConstraintSet extends fabric.Resource {
    public readonly name: fabric.Computed<string>;
    public readonly sizeConstraints?: fabric.Computed<{ comparisonOperator: string, fieldToMatch: { data?: string, type: string }[], size: number, textTransformation: string }[]>;

    constructor(urnName: string, args?: SizeConstraintSetArgs, dependsOn?: fabric.Resource[]) {
        super("aws:waf/sizeConstraintSet:SizeConstraintSet", urnName, {
            "name": args.name,
            "sizeConstraints": args.sizeConstraints,
        }, dependsOn);
    }
}

export interface SizeConstraintSetArgs {
    readonly name?: fabric.ComputedValue<string>;
    readonly sizeConstraints?: fabric.ComputedValue<{ comparisonOperator: fabric.ComputedValue<string>, fieldToMatch: fabric.ComputedValue<{ data?: fabric.ComputedValue<string>, type: fabric.ComputedValue<string> }>[], size: fabric.ComputedValue<number>, textTransformation: fabric.ComputedValue<string> }>[];
}

