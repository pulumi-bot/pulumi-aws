// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a CodeCommit Trigger Resource.
 * 
 * > **NOTE on CodeCommit**: The CodeCommit is not yet rolled out
 * in all regions - available regions are listed
 * [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#codecommit_region).
 */
export class Trigger extends pulumi.CustomResource {
    /**
     * Get an existing Trigger resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TriggerState, opts?: pulumi.CustomResourceOptions): Trigger {
        return new Trigger(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly configurationId!: pulumi.Output<string>;
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    public readonly repositoryName!: pulumi.Output<string>;
    public readonly triggers!: pulumi.Output<{ branches?: string[], customData?: string, destinationArn: string, events: string[], name: string }[]>;

    /**
     * Create a Trigger resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TriggerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TriggerArgs | TriggerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TriggerState | undefined;
            inputs["configurationId"] = state ? state.configurationId : undefined;
            inputs["repositoryName"] = state ? state.repositoryName : undefined;
            inputs["triggers"] = state ? state.triggers : undefined;
        } else {
            const args = argsOrState as TriggerArgs | undefined;
            if (!args || args.repositoryName === undefined) {
                throw new Error("Missing required property 'repositoryName'");
            }
            if (!args || args.triggers === undefined) {
                throw new Error("Missing required property 'triggers'");
            }
            inputs["repositoryName"] = args ? args.repositoryName : undefined;
            inputs["triggers"] = args ? args.triggers : undefined;
            inputs["configurationId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:codecommit/trigger:Trigger", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Trigger resources.
 */
export interface TriggerState {
    readonly configurationId?: pulumi.Input<string>;
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    readonly repositoryName?: pulumi.Input<string>;
    readonly triggers?: pulumi.Input<pulumi.Input<{ branches?: pulumi.Input<pulumi.Input<string>[]>, customData?: pulumi.Input<string>, destinationArn: pulumi.Input<string>, events: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string> }>[]>;
}

/**
 * The set of arguments for constructing a Trigger resource.
 */
export interface TriggerArgs {
    /**
     * The name for the repository. This needs to be less than 100 characters.
     */
    readonly repositoryName: pulumi.Input<string>;
    readonly triggers: pulumi.Input<pulumi.Input<{ branches?: pulumi.Input<pulumi.Input<string>[]>, customData?: pulumi.Input<string>, destinationArn: pulumi.Input<string>, events: pulumi.Input<pulumi.Input<string>[]>, name: pulumi.Input<string> }>[]>;
}
