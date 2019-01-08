// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getPartition(opts?: pulumi.InvokeOptions): Promise<GetPartitionResult> {
    return pulumi.runtime.invoke("aws:index/getPartition:getPartition", {
    }, opts);
}

/**
 * A collection of values returned by getPartition.
 */
export interface GetPartitionResult {
    readonly partition: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
