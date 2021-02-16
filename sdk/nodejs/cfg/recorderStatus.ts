// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages status (recording / stopped) of an AWS Config Configuration Recorder.
 *
 * > **Note:** Starting Configuration Recorder requires a `Delivery Channel` to be present. Use of `dependsOn` (as shown below) is recommended to avoid race conditions.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const bucket = new aws.s3.Bucket("bucket", {});
 * const fooDeliveryChannel = new aws.cfg.DeliveryChannel("fooDeliveryChannel", {s3BucketName: bucket.bucket});
 * const fooRecorderStatus = new aws.cfg.RecorderStatus("fooRecorderStatus", {isEnabled: true}, {
 *     dependsOn: [fooDeliveryChannel],
 * });
 * const role = new aws.iam.Role("role", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "config.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `});
 * const rolePolicyAttachment = new aws.iam.RolePolicyAttachment("rolePolicyAttachment", {
 *     role: role.name,
 *     policyArn: "arn:aws:iam::aws:policy/service-role/AWSConfigRole",
 * });
 * const fooRecorder = new aws.cfg.Recorder("fooRecorder", {roleArn: role.arn});
 * const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
 *     role: role.id,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": [
 *         "s3:*"
 *       ],
 *       "Effect": "Allow",
 *       "Resource": [
 *         "${bucket.arn}",
 *         "${bucket.arn}/*"
 *       ]
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * Configuration Recorder Status can be imported using the name of the Configuration Recorder, e.g.
 *
 * ```sh
 *  $ pulumi import aws:cfg/recorderStatus:RecorderStatus foo example
 * ```
 */
export class RecorderStatus extends pulumi.CustomResource {
    /**
     * Get an existing RecorderStatus resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RecorderStatusState, opts?: pulumi.CustomResourceOptions): RecorderStatus {
        return new RecorderStatus(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:cfg/recorderStatus:RecorderStatus';

    /**
     * Returns true if the given object is an instance of RecorderStatus.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RecorderStatus {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RecorderStatus.__pulumiType;
    }

    /**
     * Whether the configuration recorder should be enabled or disabled.
     */
    public readonly isEnabled!: pulumi.Output<boolean>;
    /**
     * The name of the recorder
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a RecorderStatus resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RecorderStatusArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RecorderStatusArgs | RecorderStatusState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RecorderStatusState | undefined;
            inputs["isEnabled"] = state ? state.isEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as RecorderStatusArgs | undefined;
            if ((!args || args.isEnabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'isEnabled'");
            }
            inputs["isEnabled"] = args ? args.isEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RecorderStatus.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RecorderStatus resources.
 */
export interface RecorderStatusState {
    /**
     * Whether the configuration recorder should be enabled or disabled.
     */
    readonly isEnabled?: pulumi.Input<boolean>;
    /**
     * The name of the recorder
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RecorderStatus resource.
 */
export interface RecorderStatusArgs {
    /**
     * Whether the configuration recorder should be enabled or disabled.
     */
    readonly isEnabled: pulumi.Input<boolean>;
    /**
     * The name of the recorder
     */
    readonly name?: pulumi.Input<string>;
}
