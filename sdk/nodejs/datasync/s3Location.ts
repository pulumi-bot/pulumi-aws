// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

import {ARN} from "../index";

export class S3Location extends pulumi.CustomResource {
    /**
     * Get an existing S3Location resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: S3LocationState, opts?: pulumi.CustomResourceOptions): S3Location {
        return new S3Location(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly arn: pulumi.Output<string>;
    public readonly s3BucketArn: pulumi.Output<ARN>;
    public readonly s3Config: pulumi.Output<{ bucketAccessRoleArn: string }>;
    public readonly subdirectory: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: string} | undefined>;
    public /*out*/ readonly uri: pulumi.Output<string>;

    /**
     * Create a S3Location resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: S3LocationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: S3LocationArgs | S3LocationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: S3LocationState = argsOrState as S3LocationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["s3BucketArn"] = state ? state.s3BucketArn : undefined;
            inputs["s3Config"] = state ? state.s3Config : undefined;
            inputs["subdirectory"] = state ? state.subdirectory : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as S3LocationArgs | undefined;
            if (!args || args.s3BucketArn === undefined) {
                throw new Error("Missing required property 's3BucketArn'");
            }
            if (!args || args.s3Config === undefined) {
                throw new Error("Missing required property 's3Config'");
            }
            if (!args || args.subdirectory === undefined) {
                throw new Error("Missing required property 'subdirectory'");
            }
            inputs["s3BucketArn"] = args ? args.s3BucketArn : undefined;
            inputs["s3Config"] = args ? args.s3Config : undefined;
            inputs["subdirectory"] = args ? args.subdirectory : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["uri"] = undefined /*out*/;
        }
        super("aws:datasync/s3Location:S3Location", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering S3Location resources.
 */
export interface S3LocationState {
    readonly arn?: pulumi.Input<string>;
    readonly s3BucketArn?: pulumi.Input<ARN>;
    readonly s3Config?: pulumi.Input<{ bucketAccessRoleArn: pulumi.Input<string> }>;
    readonly subdirectory?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a S3Location resource.
 */
export interface S3LocationArgs {
    readonly s3BucketArn: pulumi.Input<ARN>;
    readonly s3Config: pulumi.Input<{ bucketAccessRoleArn: pulumi.Input<string> }>;
    readonly subdirectory: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
