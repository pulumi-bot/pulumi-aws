// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides an Athena Workgroup.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.athena.Workgroup("example", {configuration: {
 *     enforceWorkgroupConfiguration: true,
 *     publishCloudwatchMetricsEnabled: true,
 *     resultConfiguration: {
 *         outputLocation: "s3://{aws_s3_bucket.example.bucket}/output/",
 *         encryptionConfiguration: {
 *             encryptionOption: "SSE_KMS",
 *             kmsKeyArn: aws_kms_key.example.arn,
 *         },
 *     },
 * }});
 * ```
 */
export class Workgroup extends pulumi.CustomResource {
    /**
     * Get an existing Workgroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkgroupState, opts?: pulumi.CustomResourceOptions): Workgroup {
        return new Workgroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:athena/workgroup:Workgroup';

    /**
     * Returns true if the given object is an instance of Workgroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workgroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workgroup.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the workgroup
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Configuration block with various settings for the workgroup. Documented below.
     */
    public readonly configuration!: pulumi.Output<outputs.athena.WorkgroupConfiguration | undefined>;
    /**
     * Description of the workgroup.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The option to delete the workgroup and its contents even if the workgroup contains any named queries.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the workgroup.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * Key-value map of resource tags for the workgroup.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Workgroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: WorkgroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkgroupArgs | WorkgroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WorkgroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["configuration"] = state ? state.configuration : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as WorkgroupArgs | undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Workgroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workgroup resources.
 */
export interface WorkgroupState {
    /**
     * Amazon Resource Name (ARN) of the workgroup
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Configuration block with various settings for the workgroup. Documented below.
     */
    readonly configuration?: pulumi.Input<inputs.athena.WorkgroupConfiguration>;
    /**
     * Description of the workgroup.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The option to delete the workgroup and its contents even if the workgroup contains any named queries.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Name of the workgroup.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags for the workgroup.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Workgroup resource.
 */
export interface WorkgroupArgs {
    /**
     * Configuration block with various settings for the workgroup. Documented below.
     */
    readonly configuration?: pulumi.Input<inputs.athena.WorkgroupConfiguration>;
    /**
     * Description of the workgroup.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The option to delete the workgroup and its contents even if the workgroup contains any named queries.
     */
    readonly forceDestroy?: pulumi.Input<boolean>;
    /**
     * Name of the workgroup.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * State of the workgroup. Valid values are `DISABLED` or `ENABLED`. Defaults to `ENABLED`.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags for the workgroup.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
