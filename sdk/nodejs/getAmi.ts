// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getAmi(args?: GetAmiArgs, opts?: pulumi.InvokeOptions): Promise<GetAmiResult> {
    args = args || {};
    return pulumi.runtime.invoke("aws:index/getAmi:getAmi", {
        "executableUsers": args.executableUsers,
        "filters": args.filters,
        "mostRecent": args.mostRecent,
        "nameRegex": args.nameRegex,
        "owners": args.owners,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getAmi.
 */
export interface GetAmiArgs {
    readonly executableUsers?: string[];
    readonly filters?: { name: string, values: string[] }[];
    readonly mostRecent?: boolean;
    readonly nameRegex?: string;
    readonly owners?: string[];
    readonly tags?: {[key: string]: any};
}

/**
 * A collection of values returned by getAmi.
 */
export interface GetAmiResult {
    readonly architecture: string;
    readonly blockDeviceMappings: { deviceName: string, ebs: {[key: string]: any}, noDevice: string, virtualName: string }[];
    readonly creationDate: string;
    readonly description: string;
    readonly hypervisor: string;
    readonly imageId: string;
    readonly imageLocation: string;
    readonly imageOwnerAlias: string;
    readonly imageType: string;
    readonly kernelId: string;
    readonly name: string;
    readonly ownerId: string;
    readonly platform: string;
    readonly productCodes: { productCodeId: string, productCodeType: string }[];
    readonly public: boolean;
    readonly ramdiskId: string;
    readonly rootDeviceName: string;
    readonly rootDeviceType: string;
    readonly rootSnapshotId: string;
    readonly sriovNetSupport: string;
    readonly state: string;
    readonly stateReason: {[key: string]: any};
    readonly tags: {[key: string]: any};
    readonly virtualizationType: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
