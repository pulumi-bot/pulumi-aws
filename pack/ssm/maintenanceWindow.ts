// *** WARNING: this file was generated by the Pulumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as fabric from "@pulumi/pulumi-fabric";

export class MaintenanceWindow extends fabric.Resource {
    public readonly allowUnassociatedTargets?: fabric.Computed<boolean>;
    public readonly cutoff: fabric.Computed<number>;
    public readonly duration: fabric.Computed<number>;
    public readonly enabled?: fabric.Computed<boolean>;
    public readonly name: fabric.Computed<string>;
    public readonly schedule: fabric.Computed<string>;

    constructor(urnName: string, args: MaintenanceWindowArgs, dependsOn?: fabric.Resource[]) {
        if (args.cutoff === undefined) {
            throw new Error("Missing required property 'cutoff'");
        }
        if (args.duration === undefined) {
            throw new Error("Missing required property 'duration'");
        }
        if (args.schedule === undefined) {
            throw new Error("Missing required property 'schedule'");
        }
        super("aws:ssm/maintenanceWindow:MaintenanceWindow", urnName, {
            "allowUnassociatedTargets": args.allowUnassociatedTargets,
            "cutoff": args.cutoff,
            "duration": args.duration,
            "enabled": args.enabled,
            "name": args.name,
            "schedule": args.schedule,
        }, dependsOn);
    }
}

export interface MaintenanceWindowArgs {
    readonly allowUnassociatedTargets?: fabric.ComputedValue<boolean>;
    readonly cutoff: fabric.ComputedValue<number>;
    readonly duration: fabric.ComputedValue<number>;
    readonly enabled?: fabric.ComputedValue<boolean>;
    readonly name?: fabric.ComputedValue<string>;
    readonly schedule: fabric.ComputedValue<string>;
}

