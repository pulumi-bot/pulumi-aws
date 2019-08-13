// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface AmiEbsBlockDevice {
    deleteOnTermination?: pulumi.Input<boolean>;
    deviceName: pulumi.Input<string>;
    encrypted?: pulumi.Input<boolean>;
    iops?: pulumi.Input<number>;
    snapshotId?: pulumi.Input<string>;
    volumeSize?: pulumi.Input<number>;
    volumeType?: pulumi.Input<string>;
}

export interface AmiEphemeralBlockDevice {
    deviceName: pulumi.Input<string>;
    virtualName: pulumi.Input<string>;
}
