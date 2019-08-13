// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface ResourceDataSyncS3Destination {
    bucketName: pulumi.Input<string>;
    kmsKeyArn?: pulumi.Input<string>;
    prefix?: pulumi.Input<string>;
    region: pulumi.Input<string>;
    syncFormat?: pulumi.Input<string>;
}
