// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides information about an Elastic File System (EFS).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const config = new pulumi.Config();
 * const var_file_system_id = config.get("fileSystemId") || "";
 * 
 * const aws_efs_file_system_by_id = pulumi.output(aws.efs.getFileSystem({
 *     fileSystemId: var_file_system_id,
 * }));
 * ```
 */
export function getFileSystem(args?: GetFileSystemArgs, opts?: pulumi.InvokeOptions): Promise<GetFileSystemResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:efs/getFileSystem:getFileSystem", {
        "creationToken": args.creationToken,
        "fileSystemId": args.fileSystemId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getFileSystem.
 */
export interface GetFileSystemArgs {
    /**
     * Restricts the list to the file system with this creation token.
     */
    readonly creationToken?: string;
    /**
     * The ID that identifies the file system (e.g. fs-ccfc0d65).
     */
    readonly fileSystemId?: string;
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getFileSystem.
 */
export interface GetFileSystemResult {
    /**
     * Amazon Resource Name of the file system.
     */
    readonly arn: string;
    readonly creationToken: string;
    /**
     * The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
     */
    readonly dnsName: string;
    /**
     * Whether EFS is encrypted.
     */
    readonly encrypted: boolean;
    readonly fileSystemId: string;
    /**
     * The ARN for the KMS encryption key.
     */
    readonly kmsKeyId: string;
    /**
     * The PerformanceMode of the file system.
     */
    readonly performanceMode: string;
    /**
     * The list of tags assigned to the file system.
     */
    readonly tags: {[key: string]: any};
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
