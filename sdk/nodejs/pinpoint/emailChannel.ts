// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a Pinpoint SMS Channel resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const role = new aws.iam.Role("role", {
 *     assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Action": "sts:AssumeRole",
 *       "Principal": {
 *         "Service": "pinpoint.amazonaws.com"
 *       },
 *       "Effect": "Allow",
 *       "Sid": ""
 *     }
 *   ]
 * }
 * `,
 * });
 * const app = new aws.pinpoint.App("app", {});
 * const identity = new aws.ses.DomainIdentity("identity", {
 *     domain: "example.com",
 * });
 * const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
 *     policy: `{
 *   "Version": "2012-10-17",
 *   "Statement": {
 *     "Action": [
 *       "mobileanalytics:PutEvents",
 *       "mobileanalytics:PutItems"
 *     ],
 *     "Effect": "Allow",
 *     "Resource": [
 *       "*"
 *     ]
 *   }
 * }
 * `,
 *     role: role.id,
 * });
 * const email = new aws.pinpoint.EmailChannel("email", {
 *     applicationId: app.applicationId,
 *     fromAddress: "user@example.com",
 *     identity: identity.arn,
 *     roleArn: role.arn,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_email_channel.html.markdown.
 */
export class EmailChannel extends pulumi.CustomResource {
    /**
     * Get an existing EmailChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EmailChannelState, opts?: pulumi.CustomResourceOptions): EmailChannel {
        return new EmailChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:pinpoint/emailChannel:EmailChannel';

    /**
     * Returns true if the given object is an instance of EmailChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EmailChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EmailChannel.__pulumiType;
    }

    /**
     * The application ID.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The email address used to send emails from.
     */
    public readonly fromAddress!: pulumi.Output<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    public readonly identity!: pulumi.Output<string>;
    /**
     * Messages per second that can be sent.
     */
    public /*out*/ readonly messagesPerSecond!: pulumi.Output<number>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    public readonly roleArn!: pulumi.Output<string>;

    /**
     * Create a EmailChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EmailChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EmailChannelArgs | EmailChannelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EmailChannelState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["fromAddress"] = state ? state.fromAddress : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["messagesPerSecond"] = state ? state.messagesPerSecond : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
        } else {
            const args = argsOrState as EmailChannelArgs | undefined;
            if (!args || args.applicationId === undefined) {
                throw new Error("Missing required property 'applicationId'");
            }
            if (!args || args.fromAddress === undefined) {
                throw new Error("Missing required property 'fromAddress'");
            }
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.roleArn === undefined) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["fromAddress"] = args ? args.fromAddress : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["messagesPerSecond"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EmailChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EmailChannel resources.
 */
export interface EmailChannelState {
    /**
     * The application ID.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The email address used to send emails from.
     */
    readonly fromAddress?: pulumi.Input<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    readonly identity?: pulumi.Input<string>;
    /**
     * Messages per second that can be sent.
     */
    readonly messagesPerSecond?: pulumi.Input<number>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    readonly roleArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EmailChannel resource.
 */
export interface EmailChannelArgs {
    /**
     * The application ID.
     */
    readonly applicationId: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The email address used to send emails from.
     */
    readonly fromAddress: pulumi.Input<string>;
    /**
     * The ARN of an identity verified with SES.
     */
    readonly identity: pulumi.Input<string>;
    /**
     * The ARN of an IAM Role used to submit events to Mobile Analytics' event ingestion service.
     */
    readonly roleArn: pulumi.Input<string>;
}
