// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.
 *
 * > **NOTE:** Lake Formation introduces fine-grained access control for data in your data lake. Part of the changes include the `IAMAllowedPrincipals` principal in order to make Lake Formation backwards compatible with existing IAM and Glue permissions. For more information, see [Changing the Default Security Settings for Your Data Lake](https://docs.aws.amazon.com/lake-formation/latest/dg/change-settings.html) and [Upgrading AWS Glue Data Permissions to the AWS Lake Formation Model](https://docs.aws.amazon.com/lake-formation/latest/dg/upgrade-glue-lake-formation.html).
 *
 * ## Example Usage
 * ### Data Lake Admins
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lakeformation.DataLakeSettings("example", {admins: [
 *     aws_iam_user.test.arn,
 *     aws_iam_role.test.arn,
 * ]});
 * ```
 * ### Create Default Permissions
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lakeformation.DataLakeSettings("example", {
 *     admins: [
 *         aws_iam_user.test.arn,
 *         aws_iam_role.test.arn,
 *     ],
 *     createDatabaseDefaultPermissions: [{
 *         permissions: [
 *             "SELECT",
 *             "ALTER",
 *             "DROP",
 *         ],
 *         principal: aws_iam_user.test.arn,
 *     }],
 *     createTableDefaultPermissions: [{
 *         permissions: ["ALL"],
 *         principal: aws_iam_role.test.arn,
 *     }],
 * });
 * ```
 */
export class DataLakeSettings extends pulumi.CustomResource {
    /**
     * Get an existing DataLakeSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataLakeSettingsState, opts?: pulumi.CustomResourceOptions): DataLakeSettings {
        return new DataLakeSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lakeformation/dataLakeSettings:DataLakeSettings';

    /**
     * Returns true if the given object is an instance of DataLakeSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataLakeSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataLakeSettings.__pulumiType;
    }

    /**
     * Set of ARNs of AWS Lake Formation principals (IAM users or roles).
     */
    public readonly admins!: pulumi.Output<string[]>;
    /**
     * Identifier for the Data Catalog. By default, the account ID.
     */
    public readonly catalogId!: pulumi.Output<string | undefined>;
    /**
     * Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
     */
    public readonly createDatabaseDefaultPermissions!: pulumi.Output<outputs.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermission[]>;
    /**
     * Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
     */
    public readonly createTableDefaultPermissions!: pulumi.Output<outputs.lakeformation.DataLakeSettingsCreateTableDefaultPermission[]>;
    /**
     * List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
     */
    public readonly trustedResourceOwners!: pulumi.Output<string[]>;

    /**
     * Create a DataLakeSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DataLakeSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataLakeSettingsArgs | DataLakeSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataLakeSettingsState | undefined;
            inputs["admins"] = state ? state.admins : undefined;
            inputs["catalogId"] = state ? state.catalogId : undefined;
            inputs["createDatabaseDefaultPermissions"] = state ? state.createDatabaseDefaultPermissions : undefined;
            inputs["createTableDefaultPermissions"] = state ? state.createTableDefaultPermissions : undefined;
            inputs["trustedResourceOwners"] = state ? state.trustedResourceOwners : undefined;
        } else {
            const args = argsOrState as DataLakeSettingsArgs | undefined;
            inputs["admins"] = args ? args.admins : undefined;
            inputs["catalogId"] = args ? args.catalogId : undefined;
            inputs["createDatabaseDefaultPermissions"] = args ? args.createDatabaseDefaultPermissions : undefined;
            inputs["createTableDefaultPermissions"] = args ? args.createTableDefaultPermissions : undefined;
            inputs["trustedResourceOwners"] = args ? args.trustedResourceOwners : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DataLakeSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataLakeSettings resources.
 */
export interface DataLakeSettingsState {
    /**
     * Set of ARNs of AWS Lake Formation principals (IAM users or roles).
     */
    admins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the Data Catalog. By default, the account ID.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
     */
    createDatabaseDefaultPermissions?: pulumi.Input<pulumi.Input<inputs.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermission>[]>;
    /**
     * Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
     */
    createTableDefaultPermissions?: pulumi.Input<pulumi.Input<inputs.lakeformation.DataLakeSettingsCreateTableDefaultPermission>[]>;
    /**
     * List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
     */
    trustedResourceOwners?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DataLakeSettings resource.
 */
export interface DataLakeSettingsArgs {
    /**
     * Set of ARNs of AWS Lake Formation principals (IAM users or roles).
     */
    admins?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Identifier for the Data Catalog. By default, the account ID.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Up to three configuration blocks of principal permissions for default create database permissions. Detailed below.
     */
    createDatabaseDefaultPermissions?: pulumi.Input<pulumi.Input<inputs.lakeformation.DataLakeSettingsCreateDatabaseDefaultPermission>[]>;
    /**
     * Up to three configuration blocks of principal permissions for default create table permissions. Detailed below.
     */
    createTableDefaultPermissions?: pulumi.Input<pulumi.Input<inputs.lakeformation.DataLakeSettingsCreateTableDefaultPermission>[]>;
    /**
     * List of the resource-owning account IDs that the caller's account can use to share their user access details (user ARNs).
     */
    trustedResourceOwners?: pulumi.Input<pulumi.Input<string>[]>;
}
