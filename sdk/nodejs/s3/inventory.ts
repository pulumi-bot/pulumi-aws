// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a S3 bucket [inventory configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) resource.
 *
 * ## Example Usage
 *
 * ### Add inventory configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const testBucket = new aws.s3.Bucket("test", {});
 * const inventory = new aws.s3.Bucket("inventory", {});
 * const testInventory = new aws.s3.Inventory("test", {
 *     bucket: testBucket.id,
 *     destination: {
 *         bucket: {
 *             bucketArn: inventory.arn,
 *             format: "ORC",
 *         },
 *     },
 *     includedObjectVersions: "All",
 *     schedule: {
 *         frequency: "Daily",
 *     },
 * });
 * ```
 *
 * ### Add inventory configuration with S3 bucket object prefix
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = new aws.s3.Bucket("test", {});
 * const inventory = new aws.s3.Bucket("inventory", {});
 * const test_prefix = new aws.s3.Inventory("test-prefix", {
 *     bucket: test.id,
 *     destination: {
 *         bucket: {
 *             bucketArn: inventory.arn,
 *             format: "ORC",
 *             prefix: "inventory",
 *         },
 *     },
 *     filter: {
 *         prefix: "documents/",
 *     },
 *     includedObjectVersions: "All",
 *     schedule: {
 *         frequency: "Daily",
 *     },
 * });
 * ```
 */
export class Inventory extends pulumi.CustomResource {
    /**
     * Get an existing Inventory resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InventoryState, opts?: pulumi.CustomResourceOptions): Inventory {
        return new Inventory(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/inventory:Inventory';

    /**
     * Returns true if the given object is an instance of Inventory.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Inventory {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Inventory.__pulumiType;
    }

    /**
     * The name of the bucket where the inventory configuration will be stored.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Contains information about where to publish the inventory results (documented below).
     */
    public readonly destination!: pulumi.Output<outputs.s3.InventoryDestination>;
    /**
     * Specifies whether the inventory is enabled or disabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
     */
    public readonly filter!: pulumi.Output<outputs.s3.InventoryFilter | undefined>;
    /**
     * Object versions to include in the inventory list. Valid values: `All`, `Current`.
     */
    public readonly includedObjectVersions!: pulumi.Output<string>;
    /**
     * Unique identifier of the inventory configuration for the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * List of optional fields that are included in the inventory results.
     * Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
     */
    public readonly optionalFields!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the schedule for generating inventory results (documented below).
     */
    public readonly schedule!: pulumi.Output<outputs.s3.InventorySchedule>;

    /**
     * Create a Inventory resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InventoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InventoryArgs | InventoryState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InventoryState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["destination"] = state ? state.destination : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["includedObjectVersions"] = state ? state.includedObjectVersions : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["optionalFields"] = state ? state.optionalFields : undefined;
            inputs["schedule"] = state ? state.schedule : undefined;
        } else {
            const args = argsOrState as InventoryArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.destination === undefined) {
                throw new Error("Missing required property 'destination'");
            }
            if (!args || args.includedObjectVersions === undefined) {
                throw new Error("Missing required property 'includedObjectVersions'");
            }
            if (!args || args.schedule === undefined) {
                throw new Error("Missing required property 'schedule'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["destination"] = args ? args.destination : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["includedObjectVersions"] = args ? args.includedObjectVersions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["optionalFields"] = args ? args.optionalFields : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Inventory.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Inventory resources.
 */
export interface InventoryState {
    /**
     * The name of the bucket where the inventory configuration will be stored.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * Contains information about where to publish the inventory results (documented below).
     */
    readonly destination?: pulumi.Input<inputs.s3.InventoryDestination>;
    /**
     * Specifies whether the inventory is enabled or disabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.InventoryFilter>;
    /**
     * Object versions to include in the inventory list. Valid values: `All`, `Current`.
     */
    readonly includedObjectVersions?: pulumi.Input<string>;
    /**
     * Unique identifier of the inventory configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of optional fields that are included in the inventory results.
     * Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
     */
    readonly optionalFields?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the schedule for generating inventory results (documented below).
     */
    readonly schedule?: pulumi.Input<inputs.s3.InventorySchedule>;
}

/**
 * The set of arguments for constructing a Inventory resource.
 */
export interface InventoryArgs {
    /**
     * The name of the bucket where the inventory configuration will be stored.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * Contains information about where to publish the inventory results (documented below).
     */
    readonly destination: pulumi.Input<inputs.s3.InventoryDestination>;
    /**
     * Specifies whether the inventory is enabled or disabled.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Specifies an inventory filter. The inventory only includes objects that meet the filter's criteria (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.InventoryFilter>;
    /**
     * Object versions to include in the inventory list. Valid values: `All`, `Current`.
     */
    readonly includedObjectVersions: pulumi.Input<string>;
    /**
     * Unique identifier of the inventory configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * List of optional fields that are included in the inventory results.
     * Valid values: `Size`, `LastModifiedDate`, `StorageClass`, `ETag`, `IsMultipartUploaded`, `ReplicationStatus`, `EncryptionStatus`, `ObjectLockRetainUntilDate`, `ObjectLockMode`, `ObjectLockLegalHoldStatus`, `IntelligentTieringAccessTier`.
     */
    readonly optionalFields?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the schedule for generating inventory results (documented below).
     */
    readonly schedule: pulumi.Input<inputs.s3.InventorySchedule>;
}
