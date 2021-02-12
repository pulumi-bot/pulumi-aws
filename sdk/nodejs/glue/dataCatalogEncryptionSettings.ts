// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Glue Data Catalog Encryption Settings resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.glue.DataCatalogEncryptionSettings("example", {dataCatalogEncryptionSettings: {
 *     connectionPasswordEncryption: {
 *         awsKmsKeyId: aws_kms_key.test.arn,
 *         returnConnectionPasswordEncrypted: true,
 *     },
 *     encryptionAtRest: {
 *         catalogEncryptionMode: "SSE-KMS",
 *         sseAwsKmsKeyId: aws_kms_key.test.arn,
 *     },
 * }});
 * ```
 *
 * ## Import
 *
 * Glue Data Catalog Encryption Settings can be imported using `CATALOG-ID` (AWS account ID if not custom), e.g.
 *
 * ```sh
 *  $ pulumi import aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings example 123456789012
 * ```
 */
export class DataCatalogEncryptionSettings extends pulumi.CustomResource {
    /**
     * Get an existing DataCatalogEncryptionSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataCatalogEncryptionSettingsState, opts?: pulumi.CustomResourceOptions): DataCatalogEncryptionSettings {
        return new DataCatalogEncryptionSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings';

    /**
     * Returns true if the given object is an instance of DataCatalogEncryptionSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataCatalogEncryptionSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataCatalogEncryptionSettings.__pulumiType;
    }

    /**
     * The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * The security configuration to set. see Data Catalog Encryption Settings.
     */
    public readonly dataCatalogEncryptionSettings!: pulumi.Output<outputs.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettings>;

    /**
     * Create a DataCatalogEncryptionSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataCatalogEncryptionSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataCatalogEncryptionSettingsArgs | DataCatalogEncryptionSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataCatalogEncryptionSettingsState | undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["dataCatalogEncryptionSettings"] = state ? state.dataCatalogEncryptionSettings : undefined;
        } else {
            const args = argsOrState as DataCatalogEncryptionSettingsArgs | undefined;
            if ((!args || args.dataCatalogEncryptionSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataCatalogEncryptionSettings'");
            }
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["dataCatalogEncryptionSettings"] = args ? args.dataCatalogEncryptionSettings : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataCatalogEncryptionSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataCatalogEncryptionSettings resources.
 */
export interface DataCatalogEncryptionSettingsState {
    /**
     * The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
     */
    readonly catalogId?: pulumi.Input<string>;
    /**
     * The security configuration to set. see Data Catalog Encryption Settings.
     */
    readonly dataCatalogEncryptionSettings?: pulumi.Input<inputs.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettings>;
}

/**
 * The set of arguments for constructing a DataCatalogEncryptionSettings resource.
 */
export interface DataCatalogEncryptionSettingsArgs {
    /**
     * The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
     */
    readonly catalogId?: pulumi.Input<string>;
    /**
     * The security configuration to set. see Data Catalog Encryption Settings.
     */
    readonly dataCatalogEncryptionSettings: pulumi.Input<inputs.glue.DataCatalogEncryptionSettingsDataCatalogEncryptionSettings>;
}
