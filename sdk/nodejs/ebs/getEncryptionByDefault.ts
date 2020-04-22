// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a way to check whether default EBS encryption is enabled for your AWS account in the current AWS region.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const current = aws.ebs.getEncryptionByDefault();
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ebs_encryption_by_default.html.markdown.
 */
export function getEncryptionByDefault(opts?: pulumi.InvokeOptions): Promise<GetEncryptionByDefaultResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:ebs/getEncryptionByDefault:getEncryptionByDefault", {
    }, opts);
}

/**
 * A collection of values returned by getEncryptionByDefault.
 */
export interface GetEncryptionByDefaultResult {
    /**
     * Whether or not default EBS encryption is enabled. Returns as `true` or `false`.
     */
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
