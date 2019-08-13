// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

import {InstanceType} from "../../../ec2/instanceType";
import {InstanceProfile} from "../../../iam";

export interface InstanceCreditSpecification {
    cpuCredits?: string;
}

export interface InstanceEbsBlockDevice {
    /**
     * Whether the volume should be destroyed
     * on instance termination (Default: `true`).
     */
    deleteOnTermination?: boolean;
    /**
     * The name of the block device to mount on the instance.
     */
    deviceName: string;
    /**
     * Enables [EBS
     * encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
     * on the volume (Default: `false`). Cannot be used with `snapshotId`. Must be configured to perform drift detection.
     */
    encrypted: boolean;
    /**
     * The amount of provisioned
     * [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
     * This must be set with a `volumeType` of `"io1"`.
     */
    iops: number;
    /**
     * Amazon Resource Name (ARN) of the KMS Key to use when encrypting the volume. Must be configured to perform drift detection.
     */
    kmsKeyId: string;
    /**
     * The Snapshot ID to mount.
     */
    snapshotId: string;
    volumeId: string;
    /**
     * The size of the volume in gibibytes (GiB).
     */
    volumeSize: number;
    /**
     * The type of volume. Can be `"standard"`, `"gp2"`,
     * or `"io1"`. (Default: `"standard"`).
     */
    volumeType: string;
}

export interface InstanceEphemeralBlockDevice {
    /**
     * The name of the block device to mount on the instance.
     */
    deviceName: string;
    /**
     * Suppresses the specified device included in the AMI's block device mapping.
     */
    noDevice?: boolean;
    /**
     * The [Instance Store Device
     * Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames)
     * (e.g. `"ephemeral0"`).
     */
    virtualName?: string;
}

export interface InstanceNetworkInterface {
    /**
     * Whether the volume should be destroyed
     * on instance termination (Default: `true`).
     */
    deleteOnTermination?: boolean;
    deviceIndex: number;
    networkInterfaceId: string;
}

export interface InstanceRootBlockDevice {
    /**
     * Whether the volume should be destroyed
     * on instance termination (Default: `true`).
     */
    deleteOnTermination?: boolean;
    /**
     * Enables [EBS
     * encryption](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
     * on the volume (Default: `false`). Cannot be used with `snapshotId`. Must be configured to perform drift detection.
     */
    encrypted: boolean;
    /**
     * The amount of provisioned
     * [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
     * This must be set with a `volumeType` of `"io1"`.
     */
    iops: number;
    /**
     * Amazon Resource Name (ARN) of the KMS Key to use when encrypting the volume. Must be configured to perform drift detection.
     */
    kmsKeyId: string;
    volumeId: string;
    /**
     * The size of the volume in gibibytes (GiB).
     */
    volumeSize: number;
    /**
     * The type of volume. Can be `"standard"`, `"gp2"`,
     * or `"io1"`. (Default: `"standard"`).
     */
    volumeType: string;
}
