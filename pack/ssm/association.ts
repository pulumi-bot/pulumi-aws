// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class Association extends fabric.Resource {
    public /*out*/ readonly associationId: fabric.Computed<string>;
    public readonly instanceId?: fabric.Computed<string>;
    public readonly name: fabric.Computed<string>;
    public readonly parameters: fabric.Computed<{[key: string]: any}>;
    public readonly targets: fabric.Computed<{ key: string, values: string[] }[]>;

    constructor(urnName: string, args?: AssociationArgs, dependsOn?: fabric.Resource[]) {
        super("aws:ssm/association:Association", urnName, {
            "instanceId": args.instanceId,
            "name": args.name,
            "parameters": args.parameters,
            "targets": args.targets,
            "associationId": undefined,
        }, dependsOn);
    }
}

export interface AssociationArgs {
    readonly instanceId?: fabric.ComputedValue<string>;
    readonly name?: fabric.ComputedValue<string>;
    readonly parameters?: fabric.ComputedValue<{[key: string]: any}>;
    readonly targets?: fabric.ComputedValue<{ key: fabric.ComputedValue<string>, values: fabric.ComputedValue<fabric.ComputedValue<string>>[] }>[];
}

