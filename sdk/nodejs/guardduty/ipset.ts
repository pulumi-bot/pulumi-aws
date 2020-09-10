// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class IPSet extends pulumi.CustomResource {
    /**
     * Get an existing IPSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IPSetState, opts?: pulumi.CustomResourceOptions): IPSet {
        return new IPSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:guardduty/iPSet:IPSet';

    /**
     * Returns true if the given object is an instance of IPSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IPSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IPSet.__pulumiType;
    }

    public readonly activate!: pulumi.Output<boolean>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly detectorId!: pulumi.Output<string>;
    public readonly format!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a IPSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IPSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IPSetArgs | IPSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IPSetState | undefined;
            inputs["activate"] = state ? state.activate : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["detectorId"] = state ? state.detectorId : undefined;
            inputs["format"] = state ? state.format : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as IPSetArgs | undefined;
            if (!args || args.activate === undefined) {
                throw new Error("Missing required property 'activate'");
            }
            if (!args || args.detectorId === undefined) {
                throw new Error("Missing required property 'detectorId'");
            }
            if (!args || args.format === undefined) {
                throw new Error("Missing required property 'format'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            inputs["activate"] = args ? args.activate : undefined;
            inputs["detectorId"] = args ? args.detectorId : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IPSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IPSet resources.
 */
export interface IPSetState {
    readonly activate?: pulumi.Input<boolean>;
    readonly arn?: pulumi.Input<string>;
    readonly detectorId?: pulumi.Input<string>;
    readonly format?: pulumi.Input<string>;
    readonly location?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a IPSet resource.
 */
export interface IPSetArgs {
    readonly activate: pulumi.Input<boolean>;
    readonly detectorId: pulumi.Input<string>;
    readonly format: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
