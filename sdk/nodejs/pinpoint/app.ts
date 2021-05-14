// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Pinpoint App resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.pinpoint.App("example", {
 *     limits: {
 *         maximumDuration: 600,
 *     },
 *     quietTime: {
 *         end: "06:00",
 *         start: "00:00",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Pinpoint App can be imported using the `application-id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:pinpoint/app:App name application-id
 * ```
 */
export class App extends pulumi.CustomResource {
    /**
     * Get an existing App resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AppState, opts?: pulumi.CustomResourceOptions): App {
        return new App(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:pinpoint/app:App';

    /**
     * Returns true if the given object is an instance of App.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is App {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === App.__pulumiType;
    }

    /**
     * The Application ID of the Pinpoint App.
     */
    public /*out*/ readonly applicationId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) of the PinPoint Application
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    public readonly campaignHook!: pulumi.Output<outputs.pinpoint.AppCampaignHook | undefined>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    public readonly limits!: pulumi.Output<outputs.pinpoint.AppLimits | undefined>;
    /**
     * The application name. By default generated by this provider
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    public readonly namePrefix!: pulumi.Output<string | undefined>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    public readonly quietTime!: pulumi.Output<outputs.pinpoint.AppQuietTime | undefined>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a App resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AppArgs | AppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AppState | undefined;
            inputs["applicationId"] = state ? state.applicationId : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["campaignHook"] = state ? state.campaignHook : undefined;
            inputs["limits"] = state ? state.limits : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["namePrefix"] = state ? state.namePrefix : undefined;
            inputs["quietTime"] = state ? state.quietTime : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as AppArgs | undefined;
            inputs["campaignHook"] = args ? args.campaignHook : undefined;
            inputs["limits"] = args ? args.limits : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["namePrefix"] = args ? args.namePrefix : undefined;
            inputs["quietTime"] = args ? args.quietTime : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["applicationId"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(App.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering App resources.
 */
export interface AppState {
    /**
     * The Application ID of the Pinpoint App.
     */
    applicationId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) of the PinPoint Application
     */
    arn?: pulumi.Input<string>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    campaignHook?: pulumi.Input<inputs.pinpoint.AppCampaignHook>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    limits?: pulumi.Input<inputs.pinpoint.AppLimits>;
    /**
     * The application name. By default generated by this provider
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    quietTime?: pulumi.Input<inputs.pinpoint.AppQuietTime>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a App resource.
 */
export interface AppArgs {
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    campaignHook?: pulumi.Input<inputs.pinpoint.AppCampaignHook>;
    /**
     * The default campaign limits for the app. These limits apply to each campaign for the app, unless the campaign overrides the default with limits of its own
     */
    limits?: pulumi.Input<inputs.pinpoint.AppLimits>;
    /**
     * The application name. By default generated by this provider
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the Pinpoint application. Conflicts with `name`
     */
    namePrefix?: pulumi.Input<string>;
    /**
     * The default quiet time for the app. Each campaign for this app sends no messages during this time unless the campaign overrides the default with a quiet time of its own
     */
    quietTime?: pulumi.Input<inputs.pinpoint.AppQuietTime>;
    /**
     * Key-value map of resource tags. If configured with a provider [`defaultTags` configuration block](https://www.terraform.io/docs/providers/aws/index.html#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
