// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get the public key about the specified KMS Key with flexible key id input. This can be useful to reference key alias without having to hard code the ARN as input.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const byAlias = pulumi.output(aws.kms.getPublicKey({
 *     keyId: "alias/my-key",
 * }));
 * const byId = pulumi.output(aws.kms.getPublicKey({
 *     keyId: "1234abcd-12ab-34cd-56ef-1234567890ab",
 * }));
 * const byAliasArn = pulumi.output(aws.kms.getPublicKey({
 *     keyId: "arn:aws:kms:us-east-1:111122223333:alias/my-key",
 * }));
 * const byKeyArn = pulumi.output(aws.kms.getPublicKey({
 *     keyId: "arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab",
 * }));
 * ```
 */
export function getPublicKey(args: GetPublicKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetPublicKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:kms/getPublicKey:getPublicKey", {
        "grantTokens": args.grantTokens,
        "keyId": args.keyId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPublicKey.
 */
export interface GetPublicKeyArgs {
    /**
     * List of grant tokens
     */
    grantTokens?: string[];
    /**
     * Key identifier which can be one of the following format:
     * * Key ID. E.g - `1234abcd-12ab-34cd-56ef-1234567890ab`
     * * Key ARN. E.g. - `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
     * * Alias name. E.g. - `alias/my-key`
     * * Alias ARN - E.g. - `arn:aws:kms:us-east-1:111122223333:alias/my-key`
     */
    keyId: string;
}

/**
 * A collection of values returned by getPublicKey.
 */
export interface GetPublicKeyResult {
    /**
     * Key ARN of the asymmetric CMK from which the public key was downloaded.
     */
    readonly arn: string;
    /**
     * Type of the public key that was downloaded.
     */
    readonly customerMasterKeySpec: string;
    /**
     * Encryption algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `ENCRYPT_DECRYPT`.
     */
    readonly encryptionAlgorithms: string[];
    readonly grantTokens?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly keyId: string;
    /**
     * Permitted use of the public key. Valid values are `ENCRYPT_DECRYPT` or `SIGN_VERIFY`
     */
    readonly keyUsage: string;
    /**
     * Exported public key. The value is a DER-encoded X.509 public key, also known as SubjectPublicKeyInfo (SPKI), as defined in [RFC 5280](https://tools.ietf.org/html/rfc5280). The value is Base64-encoded.
     */
    readonly publicKey: string;
    /**
     * Signing algorithms that AWS KMS supports for this key. Only set when the `keyUsage` of the public key is `SIGN_VERIFY`.
     */
    readonly signingAlgorithms: string[];
}

export function getPublicKeyApply(args: GetPublicKeyApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPublicKeyResult> {
    return pulumi.output(args).apply(a => getPublicKey(a, opts))
}

/**
 * A collection of arguments for invoking getPublicKey.
 */
export interface GetPublicKeyApplyArgs {
    /**
     * List of grant tokens
     */
    grantTokens?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key identifier which can be one of the following format:
     * * Key ID. E.g - `1234abcd-12ab-34cd-56ef-1234567890ab`
     * * Key ARN. E.g. - `arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`
     * * Alias name. E.g. - `alias/my-key`
     * * Alias ARN - E.g. - `arn:aws:kms:us-east-1:111122223333:alias/my-key`
     */
    keyId: pulumi.Input<string>;
}
