// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an Athena database.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const hogeBucket = new aws.s3.Bucket("hoge", {});
 * const hogeDatabase = new aws.athena.Database("hoge", {
 *     bucket: hogeBucket.bucket,
 *     name: "database_name",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/athena_database.html.markdown.
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:athena/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * Name of s3 bucket to save the results of the query execution.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
     */
    public readonly encryptionConfiguration!: pulumi.Output<{ encryptionOption: string, kmsKey?: string } | undefined>;
    /**
     * A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the database to create.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["encryptionConfiguration"] = state ? state.encryptionConfiguration : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["encryptionConfiguration"] = args ? args.encryptionConfiguration : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * Name of s3 bucket to save the results of the query execution.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
     */
    readonly encryptionConfiguration?: pulumi.Input<{ encryptionOption: pulumi.Input<string>, kmsKey?: pulumi.Input<string> }>;
    /**
     * A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Name of the database to create.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * Name of s3 bucket to save the results of the query execution.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
     */
    readonly encryptionConfiguration?: pulumi.Input<{ encryptionOption: pulumi.Input<string>, kmsKey?: pulumi.Input<string> }>;
    /**
     * A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Name of the database to create.
     */
    readonly name?: pulumi.Input<string>;
}
