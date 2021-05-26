// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides information about an Elastic File System Mount Target (EFS).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const config = new pulumi.Config();
 * const mountTargetId = config.get("mountTargetId") || "";
 * const byId = aws.efs.getMountTarget({
 *     mountTargetId: mountTargetId,
 * });
 * ```
 */
export function getMountTarget(args?: GetMountTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:efs/getMountTarget:getMountTarget", {
        "accessPointId": args.accessPointId,
        "fileSystemId": args.fileSystemId,
        "mountTargetId": args.mountTargetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMountTarget.
 */
export interface GetMountTargetArgs {
    /**
     * ID or ARN of the access point whose mount target that you want to find. It must be included if a `fileSystemId` and `mountTargetId` are not included.
     */
    accessPointId?: string;
    /**
     * ID or ARN of the file system whose mount target that you want to find. It must be included if an `accessPointId` and `mountTargetId` are not included.
     */
    fileSystemId?: string;
    /**
     * ID or ARN of the mount target that you want to find. It must be included in your request if an `accessPointId` and `fileSystemId` are not included.
     */
    mountTargetId?: string;
}

/**
 * A collection of values returned by getMountTarget.
 */
export interface GetMountTargetResult {
    readonly accessPointId?: string;
    /**
     * The unique and consistent identifier of the Availability Zone (AZ) that the mount target resides in.
     */
    readonly availabilityZoneId: string;
    /**
     * The name of the Availability Zone (AZ) that the mount target resides in.
     */
    readonly availabilityZoneName: string;
    /**
     * The DNS name for the EFS file system.
     */
    readonly dnsName: string;
    /**
     * Amazon Resource Name of the file system for which the mount target is intended.
     */
    readonly fileSystemArn: string;
    readonly fileSystemId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Address at which the file system may be mounted via the mount target.
     */
    readonly ipAddress: string;
    /**
     * The DNS name for the given subnet/AZ per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    readonly mountTargetDnsName: string;
    readonly mountTargetId: string;
    /**
     * The ID of the network interface that Amazon EFS created when it created the mount target.
     */
    readonly networkInterfaceId: string;
    /**
     * AWS account ID that owns the resource.
     */
    readonly ownerId: string;
    /**
     * List of VPC security group IDs attached to the mount target.
     */
    readonly securityGroups: string[];
    /**
     * ID of the mount target's subnet.
     */
    readonly subnetId: string;
}

export function getMountTargetApply(args?: GetMountTargetApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMountTargetResult> {
    return pulumi.output(args).apply(a => getMountTarget(a, opts))
}

/**
 * A collection of arguments for invoking getMountTarget.
 */
export interface GetMountTargetApplyArgs {
    /**
     * ID or ARN of the access point whose mount target that you want to find. It must be included if a `fileSystemId` and `mountTargetId` are not included.
     */
    accessPointId?: pulumi.Input<string>;
    /**
     * ID or ARN of the file system whose mount target that you want to find. It must be included if an `accessPointId` and `mountTargetId` are not included.
     */
    fileSystemId?: pulumi.Input<string>;
    /**
     * ID or ARN of the mount target that you want to find. It must be included in your request if an `accessPointId` and `fileSystemId` are not included.
     */
    mountTargetId?: pulumi.Input<string>;
}
