// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a CodePipeline Webhook.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/codepipeline_webhook.html.markdown.
 */
export class Webhook extends pulumi.CustomResource {
    /**
     * Get an existing Webhook resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WebhookState, opts?: pulumi.CustomResourceOptions): Webhook {
        return new Webhook(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:codepipeline/webhook:Webhook';

    /**
     * Returns true if the given object is an instance of Webhook.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Webhook {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Webhook.__pulumiType;
    }

    /**
     * The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
     */
    public readonly authentication!: pulumi.Output<string>;
    /**
     * An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
     */
    public readonly authenticationConfiguration!: pulumi.Output<outputApi.codepipeline.WebhookAuthenticationConfiguration | undefined>;
    /**
     * One or more `filter` blocks. Filter blocks are documented below.
     */
    public readonly filters!: pulumi.Output<outputApi.codepipeline.WebhookFilter[]>;
    /**
     * The name of the webhook.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
     */
    public readonly targetAction!: pulumi.Output<string>;
    /**
     * The name of the pipeline.
     */
    public readonly targetPipeline!: pulumi.Output<string>;
    /**
     * The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
     */
    public /*out*/ readonly url!: pulumi.Output<string>;

    /**
     * Create a Webhook resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WebhookArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WebhookArgs | WebhookState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as WebhookState | undefined;
            inputs["authentication"] = state ? state.authentication : undefined;
            inputs["authenticationConfiguration"] = state ? state.authenticationConfiguration : undefined;
            inputs["filters"] = state ? state.filters : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["targetAction"] = state ? state.targetAction : undefined;
            inputs["targetPipeline"] = state ? state.targetPipeline : undefined;
            inputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as WebhookArgs | undefined;
            if (!args || args.authentication === undefined) {
                throw new Error("Missing required property 'authentication'");
            }
            if (!args || args.filters === undefined) {
                throw new Error("Missing required property 'filters'");
            }
            if (!args || args.targetAction === undefined) {
                throw new Error("Missing required property 'targetAction'");
            }
            if (!args || args.targetPipeline === undefined) {
                throw new Error("Missing required property 'targetPipeline'");
            }
            inputs["authentication"] = args ? args.authentication : undefined;
            inputs["authenticationConfiguration"] = args ? args.authenticationConfiguration : undefined;
            inputs["filters"] = args ? args.filters : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetAction"] = args ? args.targetAction : undefined;
            inputs["targetPipeline"] = args ? args.targetPipeline : undefined;
            inputs["url"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Webhook.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Webhook resources.
 */
export interface WebhookState {
    /**
     * The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
     */
    readonly authentication?: pulumi.Input<string>;
    /**
     * An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
     */
    readonly authenticationConfiguration?: pulumi.Input<inputApi.codepipeline.WebhookAuthenticationConfiguration>;
    /**
     * One or more `filter` blocks. Filter blocks are documented below.
     */
    readonly filters?: pulumi.Input<pulumi.Input<inputApi.codepipeline.WebhookFilter>[]>;
    /**
     * The name of the webhook.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
     */
    readonly targetAction?: pulumi.Input<string>;
    /**
     * The name of the pipeline.
     */
    readonly targetPipeline?: pulumi.Input<string>;
    /**
     * The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
     */
    readonly url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Webhook resource.
 */
export interface WebhookArgs {
    /**
     * The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
     */
    readonly authentication: pulumi.Input<string>;
    /**
     * An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
     */
    readonly authenticationConfiguration?: pulumi.Input<inputApi.codepipeline.WebhookAuthenticationConfiguration>;
    /**
     * One or more `filter` blocks. Filter blocks are documented below.
     */
    readonly filters: pulumi.Input<pulumi.Input<inputApi.codepipeline.WebhookFilter>[]>;
    /**
     * The name of the webhook.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
     */
    readonly targetAction: pulumi.Input<string>;
    /**
     * The name of the pipeline.
     */
    readonly targetPipeline: pulumi.Input<string>;
}
