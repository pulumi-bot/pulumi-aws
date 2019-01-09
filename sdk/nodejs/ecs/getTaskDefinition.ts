// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getTaskDefinition(args: GetTaskDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskDefinitionResult> {
    return pulumi.runtime.invoke("aws:ecs/getTaskDefinition:getTaskDefinition", {
        "taskDefinition": args.taskDefinition,
    }, opts);
}

/**
 * A collection of arguments for invoking getTaskDefinition.
 */
export interface GetTaskDefinitionArgs {
    readonly taskDefinition: string;
}

/**
 * A collection of values returned by getTaskDefinition.
 */
export interface GetTaskDefinitionResult {
    readonly family: string;
    readonly networkMode: string;
    readonly revision: number;
    readonly status: string;
    readonly taskRoleArn: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
