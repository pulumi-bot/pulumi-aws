// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * S3 Control Bucket Lifecycle Configurations can be imported using the Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
 * ```
 */
export class BucketLifecycleConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing BucketLifecycleConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketLifecycleConfigurationState, opts?: pulumi.CustomResourceOptions): BucketLifecycleConfiguration {
        return new BucketLifecycleConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3control/bucketLifecycleConfiguration:BucketLifecycleConfiguration';

    /**
     * Returns true if the given object is an instance of BucketLifecycleConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketLifecycleConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketLifecycleConfiguration.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the bucket.
     */
    public readonly bucket!: pulumi.Output<string>;
    /**
     * Configuration block(s) containing lifecycle rules for the bucket.
     */
    public readonly rules!: pulumi.Output<outputs.s3control.BucketLifecycleConfigurationRule[]>;

    /**
     * Create a BucketLifecycleConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketLifecycleConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketLifecycleConfigurationArgs | BucketLifecycleConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as BucketLifecycleConfigurationState | undefined;
            inputs["bucket"] = state ? state.bucket : undefined;
            inputs["rules"] = state ? state.rules : undefined;
        } else {
            const args = argsOrState as BucketLifecycleConfigurationArgs | undefined;
            if (!args || args.bucket === undefined) {
                throw new Error("Missing required property 'bucket'");
            }
            if (!args || args.rules === undefined) {
                throw new Error("Missing required property 'rules'");
            }
            inputs["bucket"] = args ? args.bucket : undefined;
            inputs["rules"] = args ? args.rules : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BucketLifecycleConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketLifecycleConfiguration resources.
 */
export interface BucketLifecycleConfigurationState {
    /**
     * Amazon Resource Name (ARN) of the bucket.
     */
    readonly bucket?: pulumi.Input<string>;
    /**
     * Configuration block(s) containing lifecycle rules for the bucket.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.s3control.BucketLifecycleConfigurationRule>[]>;
}

/**
 * The set of arguments for constructing a BucketLifecycleConfiguration resource.
 */
export interface BucketLifecycleConfigurationArgs {
    /**
     * Amazon Resource Name (ARN) of the bucket.
     */
    readonly bucket: pulumi.Input<string>;
    /**
     * Configuration block(s) containing lifecycle rules for the bucket.
     */
    readonly rules: pulumi.Input<pulumi.Input<inputs.s3control.BucketLifecycleConfigurationRule>[]>;
}
