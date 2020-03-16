// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an SES configuration set resource
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const test = new aws.ses.ConfigurationSet("test", {});
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ses_configuration_set.markdown.
 */
export class ConfigurationSet extends pulumi.CustomResource {
    /**
     * Get an existing ConfigurationSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationSetState, opts?: pulumi.CustomResourceOptions): ConfigurationSet {
        return new ConfigurationSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ses/configurationSet:ConfigurationSet';

    /**
     * Returns true if the given object is an instance of ConfigurationSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConfigurationSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConfigurationSet.__pulumiType;
    }

    /**
     * The name of the configuration set
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ConfigurationSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ConfigurationSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationSetArgs | ConfigurationSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConfigurationSetState | undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ConfigurationSetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "aws:ses/confgurationSet:ConfgurationSet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ConfigurationSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConfigurationSet resources.
 */
export interface ConfigurationSetState {
    /**
     * The name of the configuration set
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConfigurationSet resource.
 */
export interface ConfigurationSetArgs {
    /**
     * The name of the configuration set
     */
    readonly name?: pulumi.Input<string>;
}
