// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.
 *
 * ## Example Usage
 */
export class AnalyticsConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing AnalyticsConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AnalyticsConfigurationState, opts?: pulumi.CustomResourceOptions): AnalyticsConfiguration {
        return new AnalyticsConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3/analyticsConfiguration:AnalyticsConfiguration';

    /**
     * Returns true if the given object is an instance of AnalyticsConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AnalyticsConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AnalyticsConfiguration.__pulumiType;
    }

    /**
     * The name of the bucket this analytics configuration is associated with.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    public readonly filter!: pulumi.Output<outputs.s3.AnalyticsConfigurationFilter | undefined>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    public readonly storageClassAnalysis!: pulumi.Output<outputs.s3.AnalyticsConfigurationStorageClassAnalysis | undefined>;

    /**
     * Create a AnalyticsConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AnalyticsConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AnalyticsConfigurationArgs | AnalyticsConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AnalyticsConfigurationState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["filter"] = state ? state.filter : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["storageClassAnalysis"] = state ? state.storageClassAnalysis : undefined;
        } else {
            const args = argsOrState as AnalyticsConfigurationArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["filter"] = args ? args.filter : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["storageClassAnalysis"] = args ? args.storageClassAnalysis : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AnalyticsConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AnalyticsConfiguration resources.
 */
export interface AnalyticsConfigurationState {
    /**
     * The name of the bucket this analytics configuration is associated with.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.AnalyticsConfigurationFilter>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    readonly storageClassAnalysis?: pulumi.Input<inputs.s3.AnalyticsConfigurationStorageClassAnalysis>;
}

/**
 * The set of arguments for constructing a AnalyticsConfiguration resource.
 */
export interface AnalyticsConfigurationArgs {
    /**
     * The name of the bucket this analytics configuration is associated with.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
     */
    readonly filter?: pulumi.Input<inputs.s3.AnalyticsConfigurationFilter>;
    /**
     * Unique identifier of the analytics configuration for the bucket.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration for the analytics data export (documented below).
     */
    readonly storageClassAnalysis?: pulumi.Input<inputs.s3.AnalyticsConfigurationStorageClassAnalysis>;
}
