// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface SecurityConfigurationEncryptionConfiguration {
    cloudwatchEncryption: pulumi.Input<inputApi.glue.SecurityConfigurationEncryptionConfigurationCloudwatchEncryption>;
    jobBookmarksEncryption: pulumi.Input<inputApi.glue.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption>;
    /**
     * A `s3Encryption ` block as described below, which contains encryption configuration for S3 data.
     */
    s3Encryption: pulumi.Input<inputApi.glue.SecurityConfigurationEncryptionConfigurationS3Encryption>;
}

export interface SecurityConfigurationEncryptionConfigurationCloudwatchEncryption {
    /**
     * Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
     */
    cloudwatchEncryptionMode?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
     */
    kmsKeyArn?: pulumi.Input<string>;
}

export interface SecurityConfigurationEncryptionConfigurationJobBookmarksEncryption {
    /**
     * Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
     */
    jobBookmarksEncryptionMode?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
     */
    kmsKeyArn?: pulumi.Input<string>;
}

export interface SecurityConfigurationEncryptionConfigurationS3Encryption {
    /**
     * Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
     */
    kmsKeyArn?: pulumi.Input<string>;
    /**
     * Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
     */
    s3EncryptionMode?: pulumi.Input<string>;
}
