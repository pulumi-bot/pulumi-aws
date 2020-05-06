// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the default EBS encryption KMS key in the current region.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_default_kms_key.html.markdown.
 */
export function getDefaultKmsKey(opts?: pulumi.InvokeOptions): Promise<GetDefaultKmsKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ebs/getDefaultKmsKey:getDefaultKmsKey", {
    }, opts);
}

/**
 * A collection of values returned by getDefaultKmsKey.
 */
export interface GetDefaultKmsKeyResult {
    /**
     * Amazon Resource Name (ARN) of the default KMS key uses to encrypt an EBS volume in this region when no key is specified in an API call that creates the volume and encryption by default is enabled.
     */
    readonly keyArn: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
