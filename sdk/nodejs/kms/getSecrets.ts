// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Decrypt multiple secrets from data encrypted with the AWS KMS service.
 */
export function getSecrets(args: GetSecretsArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:kms/getSecrets:getSecrets", {
        "secrets": args.secrets,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecrets.
 */
export interface GetSecretsArgs {
    /**
     * One or more encrypted payload definitions from the KMS service. See the Secret Definitions below.
     */
    readonly secrets: inputs.kms.GetSecretsSecret[];
}

/**
 * A collection of values returned by getSecrets.
 */
export interface GetSecretsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Map containing each `secret` `name` as the key with its decrypted plaintext value
     */
    readonly plaintext: {[key: string]: string};
    readonly secrets: outputs.kms.GetSecretsSecret[];
}
