// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about an Elastic File System (EFS) Access Point.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.efs.getAccessPoint({
 *     accessPointId: "fsap-12345678",
 * }));
 * ```
 */
export function getAccessPoint(args: GetAccessPointArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessPointResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:efs/getAccessPoint:getAccessPoint", {
        "accessPointId": args.accessPointId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessPoint.
 */
export interface GetAccessPointArgs {
    /**
     * The ID that identifies the file system.
     */
    readonly accessPointId: string;
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getAccessPoint.
 */
export interface GetAccessPointResult {
    readonly accessPointId: string;
    /**
     * Amazon Resource Name of the file system.
     */
    readonly arn: string;
    /**
     * Amazon Resource Name of the file system.
     */
    readonly fileSystemArn: string;
    /**
     * The ID of the file system for which the access point is intended.
     */
    readonly fileSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ownerId: string;
    /**
     * Single element list containing operating system user and group applied to all file system requests made using the access point.
     */
    readonly posixUsers: outputs.efs.GetAccessPointPosixUser[];
    readonly rootDirectories: outputs.efs.GetAccessPointRootDirectory[];
    /**
     * Key-value mapping of resource tags.
     */
    readonly tags?: {[key: string]: string};
}
