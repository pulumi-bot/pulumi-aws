// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Pinpoint GCM Channel resource.
 * 
 * > **Note:** Api Key argument will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const app = new aws.pinpoint.App("app", {});
 * const gcm = new aws.pinpoint.GcmChannel("gcm", {
 *     apiKey: "api_key",
 *     applicationId: app.applicationId,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/pinpoint_gcm_channel.html.markdown.
 */
export class GcmChannel extends pulumi.CustomResource {
    /**
     * Get an existing GcmChannel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GcmChannelState, opts?: pulumi.CustomResourceOptions): GcmChannel {
        return new GcmChannel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:pinpoint/gcmChannel:GcmChannel';

    /**
     * Returns true if the given object is an instance of GcmChannel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GcmChannel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GcmChannel.__pulumiType;
    }

    /**
     * Platform credential API key from Google.
     */
    public readonly apiKey!: pulumi.Output<string>;
    /**
     * The application ID.
     */
    public readonly applicationId!: pulumi.Output<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a GcmChannel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GcmChannelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GcmChannelArgs | GcmChannelState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GcmChannelState | undefined;
            inputs["apiKey"] = state ? state.apiKey : undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
        } else {
            const args = argsOrState as GcmChannelArgs | undefined;
            if (!args || args.apiKey === undefined) {
                throw new Error("Missing required property 'apiKey'");
            }
            if (!args || args.applicationId === undefined) {
                throw new Error("Missing required property 'applicationId'");
            }
            inputs["apiKey"] = args ? args.apiKey : undefined;
            inputs["applicationId"] = args ? args.applicationId : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GcmChannel.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GcmChannel resources.
 */
export interface GcmChannelState {
    /**
     * Platform credential API key from Google.
     */
    readonly apiKey?: pulumi.Input<string>;
    /**
     * The application ID.
     */
    readonly applicationId?: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GcmChannel resource.
 */
export interface GcmChannelArgs {
    /**
     * Platform credential API key from Google.
     */
    readonly apiKey: pulumi.Input<string>;
    /**
     * The application ID.
     */
    readonly applicationId: pulumi.Input<string>;
    /**
     * Whether the channel is enabled or disabled. Defaults to `true`.
     */
    readonly enabled?: pulumi.Input<boolean>;
}
