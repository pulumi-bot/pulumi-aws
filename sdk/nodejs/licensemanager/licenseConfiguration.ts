// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a License Manager license configuration resource.
 *
 * > **Note:** Removing the `licenseCount` attribute is not supported by the License Manager API - recreate the resource instead.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.licensemanager.LicenseConfiguration("example", {
 *     description: "Example",
 *     licenseCount: 10,
 *     licenseCountHardLimit: true,
 *     licenseCountingType: "Socket",
 *     licenseRules: ["#minimumSockets=2"],
 *     tags: {
 *         foo: "barr",
 *     },
 * });
 * ```
 * ## Rules
 *
 * License rules should be in the format of `#RuleType=RuleValue`. Supported rule types:
 *
 * * `minimumVcpus` - Resource must have minimum vCPU count in order to use the license. Default: 1
 * * `maximumVcpus` - Resource must have maximum vCPU count in order to use the license. Default: unbounded, limit: 10000
 * * `minimumCores` - Resource must have minimum core count in order to use the license. Default: 1
 * * `maximumCores` - Resource must have maximum core count in order to use the license. Default: unbounded, limit: 10000
 * * `minimumSockets` - Resource must have minimum socket count in order to use the license. Default: 1
 * * `maximumSockets` - Resource must have maximum socket count in order to use the license. Default: unbounded, limit: 10000
 * * `allowedTenancy` - Defines where the license can be used. If set, restricts license usage to selected tenancies. Specify a comma delimited list of `EC2-Default`, `EC2-DedicatedHost`, `EC2-DedicatedInstance`
 *
 * ## Import
 *
 * License configurations can be imported using the `id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:licensemanager/licenseConfiguration:LicenseConfiguration example arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef
 * ```
 */
export class LicenseConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing LicenseConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LicenseConfigurationState, opts?: pulumi.CustomResourceOptions): LicenseConfiguration {
        return new LicenseConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:licensemanager/licenseConfiguration:LicenseConfiguration';

    /**
     * Returns true if the given object is an instance of LicenseConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LicenseConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LicenseConfiguration.__pulumiType;
    }

    /**
     * Description of the license configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Number of licenses managed by the license configuration.
     */
    public readonly licenseCount!: pulumi.Output<number | undefined>;
    /**
     * Sets the number of available licenses as a hard limit.
     */
    public readonly licenseCountHardLimit!: pulumi.Output<boolean | undefined>;
    /**
     * Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
     */
    public readonly licenseCountingType!: pulumi.Output<string>;
    /**
     * Array of configured License Manager rules.
     */
    public readonly licenseRules!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the license configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a LicenseConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LicenseConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LicenseConfigurationArgs | LicenseConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LicenseConfigurationState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["licenseCount"] = state ? state.licenseCount : undefined;
            inputs["licenseCountHardLimit"] = state ? state.licenseCountHardLimit : undefined;
            inputs["licenseCountingType"] = state ? state.licenseCountingType : undefined;
            inputs["licenseRules"] = state ? state.licenseRules : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as LicenseConfigurationArgs | undefined;
            if ((!args || args.licenseCountingType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'licenseCountingType'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["licenseCount"] = args ? args.licenseCount : undefined;
            inputs["licenseCountHardLimit"] = args ? args.licenseCountHardLimit : undefined;
            inputs["licenseCountingType"] = args ? args.licenseCountingType : undefined;
            inputs["licenseRules"] = args ? args.licenseRules : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(LicenseConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LicenseConfiguration resources.
 */
export interface LicenseConfigurationState {
    /**
     * Description of the license configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Number of licenses managed by the license configuration.
     */
    readonly licenseCount?: pulumi.Input<number>;
    /**
     * Sets the number of available licenses as a hard limit.
     */
    readonly licenseCountHardLimit?: pulumi.Input<boolean>;
    /**
     * Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
     */
    readonly licenseCountingType?: pulumi.Input<string>;
    /**
     * Array of configured License Manager rules.
     */
    readonly licenseRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the license configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a LicenseConfiguration resource.
 */
export interface LicenseConfigurationArgs {
    /**
     * Description of the license configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Number of licenses managed by the license configuration.
     */
    readonly licenseCount?: pulumi.Input<number>;
    /**
     * Sets the number of available licenses as a hard limit.
     */
    readonly licenseCountHardLimit?: pulumi.Input<boolean>;
    /**
     * Dimension to use to track license inventory. Specify either `vCPU`, `Instance`, `Core` or `Socket`.
     */
    readonly licenseCountingType: pulumi.Input<string>;
    /**
     * Array of configured License Manager rules.
     */
    readonly licenseRules?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the license configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
