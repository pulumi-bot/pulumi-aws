// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Retrieve information about a Storage Gateway local disk. The disk identifier is useful for adding the disk as a cache or upload buffer to a gateway.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const test = pulumi.all([aws_volume_attachment_test.deviceName, aws_storagegateway_gateway_test.arn]).apply(([deviceName, arn]) => aws.storagegateway.getLocalDisk({
 *     diskPath: deviceName,
 *     gatewayArn: arn,
 * }));
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/storagegateway_local_disk.html.markdown.
 */
export function getLocalDisk(args: GetLocalDiskArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalDiskResult> & GetLocalDiskResult {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetLocalDiskResult> = pulumi.runtime.invoke("aws:storagegateway/getLocalDisk:getLocalDisk", {
        "diskNode": args.diskNode,
        "diskPath": args.diskPath,
        "gatewayArn": args.gatewayArn,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getLocalDisk.
 */
export interface GetLocalDiskArgs {
    /**
     * The device node of the local disk to retrieve. For example, `/dev/sdb`.
     */
    readonly diskNode?: string;
    /**
     * The device path of the local disk to retrieve. For example, `/dev/xvdb` or `/dev/nvme1n1`.
     */
    readonly diskPath?: string;
    /**
     * The Amazon Resource Name (ARN) of the gateway.
     */
    readonly gatewayArn: string;
}

/**
 * A collection of values returned by getLocalDisk.
 */
export interface GetLocalDiskResult {
    /**
     * The disk identifier. e.g. `pci-0000:03:00.0-scsi-0:0:0:0`
     */
    readonly diskId: string;
    readonly diskNode?: string;
    readonly diskPath?: string;
    readonly gatewayArn: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
