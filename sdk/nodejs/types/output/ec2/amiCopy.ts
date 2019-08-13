// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface AmiCopyEbsBlockDevice {
    deleteOnTermination: boolean;
    deviceName: string;
    /**
     * Specifies whether the destination snapshots of the copied image should be encrypted. Defaults to `false`
     */
    encrypted: boolean;
    iops: number;
    snapshotId: string;
    volumeSize: number;
    volumeType: string;
}

export interface AmiCopyEphemeralBlockDevice {
    deviceName: string;
    virtualName: string;
}
