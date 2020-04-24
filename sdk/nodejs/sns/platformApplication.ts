// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SNS platform application resource
 * 
 * ## Example Usage
 * 
 * ### Apple Push Notification Service (APNS)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const apnsApplication = new aws.sns.PlatformApplication("apnsApplication", {
 *     platform: "APNS",
 *     platformCredential: "<APNS PRIVATE KEY>",
 *     platformPrincipal: "<APNS CERTIFICATE>",
 * });
 * ```
 * 
 * ### Google Cloud Messaging (GCM)
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const gcmApplication = new aws.sns.PlatformApplication("gcmApplication", {
 *     platform: "GCM",
 *     platformCredential: "<GCM API KEY>",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sns_platform_application.html.markdown.
 */
export class PlatformApplication extends pulumi.CustomResource {
    /**
     * Get an existing PlatformApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PlatformApplicationState, opts?: pulumi.CustomResourceOptions): PlatformApplication {
        return new PlatformApplication(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sns/platformApplication:PlatformApplication';

    /**
     * Returns true if the given object is an instance of PlatformApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PlatformApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PlatformApplication.__pulumiType;
    }

    /**
     * The ARN of the SNS platform application
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
     */
    public readonly eventDeliveryFailureTopicArn!: pulumi.Output<string | undefined>;
    /**
     * SNS Topic triggered when a new platform endpoint is added to your platform application.
     */
    public readonly eventEndpointCreatedTopicArn!: pulumi.Output<string | undefined>;
    /**
     * SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
     */
    public readonly eventEndpointDeletedTopicArn!: pulumi.Output<string | undefined>;
    /**
     * SNS Topic triggered when an existing platform endpoint is changed from your platform application.
     */
    public readonly eventEndpointUpdatedTopicArn!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive failure feedback for this application.
     */
    public readonly failureFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The friendly name for the SNS platform application
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
     */
    public readonly platform!: pulumi.Output<string>;
    /**
     * Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    public readonly platformCredential!: pulumi.Output<string>;
    /**
     * Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    public readonly platformPrincipal!: pulumi.Output<string | undefined>;
    /**
     * The IAM role permitted to receive success feedback for this application.
     */
    public readonly successFeedbackRoleArn!: pulumi.Output<string | undefined>;
    /**
     * The percentage of success to sample (0-100)
     */
    public readonly successFeedbackSampleRate!: pulumi.Output<string | undefined>;

    /**
     * Create a PlatformApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PlatformApplicationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PlatformApplicationArgs | PlatformApplicationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as PlatformApplicationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["eventDeliveryFailureTopicArn"] = state ? state.eventDeliveryFailureTopicArn : undefined;
            inputs["eventEndpointCreatedTopicArn"] = state ? state.eventEndpointCreatedTopicArn : undefined;
            inputs["eventEndpointDeletedTopicArn"] = state ? state.eventEndpointDeletedTopicArn : undefined;
            inputs["eventEndpointUpdatedTopicArn"] = state ? state.eventEndpointUpdatedTopicArn : undefined;
            inputs["failureFeedbackRoleArn"] = state ? state.failureFeedbackRoleArn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["platform"] = state ? state.platform : undefined;
            inputs["platformCredential"] = state ? state.platformCredential : undefined;
            inputs["platformPrincipal"] = state ? state.platformPrincipal : undefined;
            inputs["successFeedbackRoleArn"] = state ? state.successFeedbackRoleArn : undefined;
            inputs["successFeedbackSampleRate"] = state ? state.successFeedbackSampleRate : undefined;
        } else {
            const args = argsOrState as PlatformApplicationArgs | undefined;
            if (!args || args.platform === undefined) {
                throw new Error("Missing required property 'platform'");
            }
            if (!args || args.platformCredential === undefined) {
                throw new Error("Missing required property 'platformCredential'");
            }
            inputs["eventDeliveryFailureTopicArn"] = args ? args.eventDeliveryFailureTopicArn : undefined;
            inputs["eventEndpointCreatedTopicArn"] = args ? args.eventEndpointCreatedTopicArn : undefined;
            inputs["eventEndpointDeletedTopicArn"] = args ? args.eventEndpointDeletedTopicArn : undefined;
            inputs["eventEndpointUpdatedTopicArn"] = args ? args.eventEndpointUpdatedTopicArn : undefined;
            inputs["failureFeedbackRoleArn"] = args ? args.failureFeedbackRoleArn : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["platform"] = args ? args.platform : undefined;
            inputs["platformCredential"] = args ? args.platformCredential : undefined;
            inputs["platformPrincipal"] = args ? args.platformPrincipal : undefined;
            inputs["successFeedbackRoleArn"] = args ? args.successFeedbackRoleArn : undefined;
            inputs["successFeedbackSampleRate"] = args ? args.successFeedbackSampleRate : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PlatformApplication.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PlatformApplication resources.
 */
export interface PlatformApplicationState {
    /**
     * The ARN of the SNS platform application
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
     */
    readonly eventDeliveryFailureTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when a new platform endpoint is added to your platform application.
     */
    readonly eventEndpointCreatedTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
     */
    readonly eventEndpointDeletedTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when an existing platform endpoint is changed from your platform application.
     */
    readonly eventEndpointUpdatedTopicArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive failure feedback for this application.
     */
    readonly failureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The friendly name for the SNS platform application
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
     */
    readonly platform?: pulumi.Input<string>;
    /**
     * Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    readonly platformCredential?: pulumi.Input<string>;
    /**
     * Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    readonly platformPrincipal?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this application.
     */
    readonly successFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The percentage of success to sample (0-100)
     */
    readonly successFeedbackSampleRate?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PlatformApplication resource.
 */
export interface PlatformApplicationArgs {
    /**
     * SNS Topic triggered when a delivery to any of the platform endpoints associated with your platform application encounters a permanent failure.
     */
    readonly eventDeliveryFailureTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when a new platform endpoint is added to your platform application.
     */
    readonly eventEndpointCreatedTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when an existing platform endpoint is deleted from your platform application.
     */
    readonly eventEndpointDeletedTopicArn?: pulumi.Input<string>;
    /**
     * SNS Topic triggered when an existing platform endpoint is changed from your platform application.
     */
    readonly eventEndpointUpdatedTopicArn?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive failure feedback for this application.
     */
    readonly failureFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The friendly name for the SNS platform application
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The platform that the app is registered with. See [Platform](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for supported platforms.
     */
    readonly platform: pulumi.Input<string>;
    /**
     * Application Platform credential. See [Credential](http://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-register.html) for type of credential required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    readonly platformCredential: pulumi.Input<string>;
    /**
     * Application Platform principal. See [Principal](http://docs.aws.amazon.com/sns/latest/api/API_CreatePlatformApplication.html) for type of principal required for platform. The value of this attribute when stored into the state is only a hash of the real value, so therefore it is not practical to use this as an attribute for other resources.
     */
    readonly platformPrincipal?: pulumi.Input<string>;
    /**
     * The IAM role permitted to receive success feedback for this application.
     */
    readonly successFeedbackRoleArn?: pulumi.Input<string>;
    /**
     * The percentage of success to sample (0-100)
     */
    readonly successFeedbackSampleRate?: pulumi.Input<string>;
}
