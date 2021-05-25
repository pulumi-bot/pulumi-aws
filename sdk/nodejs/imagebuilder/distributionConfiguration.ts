// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an Image Builder Distribution Configuration.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.imagebuilder.DistributionConfiguration("example", {
 *     distributions: [{
 *         amiDistributionConfiguration: {
 *             amiTags: {
 *                 CostCenter: "IT",
 *             },
 *             launchPermission: {
 *                 userIds: ["123456789012"],
 *             },
 *             name: "example-{{ imagebuilder:buildDate }}",
 *         },
 *         region: "us-east-1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * `aws_imagebuilder_distribution_configurations` resources can be imported by using the Amazon Resource Name (ARN), e.g.
 *
 * ```sh
 *  $ pulumi import aws:imagebuilder/distributionConfiguration:DistributionConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
 * ```
 */
export class DistributionConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DistributionConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DistributionConfigurationState, opts?: pulumi.CustomResourceOptions): DistributionConfiguration {
        return new DistributionConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:imagebuilder/distributionConfiguration:DistributionConfiguration';

    /**
     * Returns true if the given object is an instance of DistributionConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DistributionConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DistributionConfiguration.__pulumiType;
    }

    /**
     * (Required) Amazon Resource Name (ARN) of the distribution configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Date the distribution configuration was created.
     */
    public /*out*/ readonly dateCreated!: pulumi.Output<string>;
    /**
     * Date the distribution configuration was updated.
     */
    public /*out*/ readonly dateUpdated!: pulumi.Output<string>;
    /**
     * Description to apply to the distributed AMI.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * One or more configuration blocks with distribution settings. Detailed below.
     */
    public readonly distributions!: pulumi.Output<outputs.imagebuilder.DistributionConfigurationDistribution[]>;
    /**
     * Name to apply to the distributed AMI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Key-value map of resource tags for the distribution configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a DistributionConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DistributionConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DistributionConfigurationArgs | DistributionConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DistributionConfigurationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["dateCreated"] = state ? state.dateCreated : undefined;
            inputs["dateUpdated"] = state ? state.dateUpdated : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["distributions"] = state ? state.distributions : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as DistributionConfigurationArgs | undefined;
            if ((!args || args.distributions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'distributions'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["distributions"] = args ? args.distributions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["tagsAll"] = args ? args.tagsAll : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["dateCreated"] = undefined /*out*/;
            inputs["dateUpdated"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DistributionConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DistributionConfiguration resources.
 */
export interface DistributionConfigurationState {
    /**
     * (Required) Amazon Resource Name (ARN) of the distribution configuration.
     */
    arn?: pulumi.Input<string>;
    /**
     * Date the distribution configuration was created.
     */
    dateCreated?: pulumi.Input<string>;
    /**
     * Date the distribution configuration was updated.
     */
    dateUpdated?: pulumi.Input<string>;
    /**
     * Description to apply to the distributed AMI.
     */
    description?: pulumi.Input<string>;
    /**
     * One or more configuration blocks with distribution settings. Detailed below.
     */
    distributions?: pulumi.Input<pulumi.Input<inputs.imagebuilder.DistributionConfigurationDistribution>[]>;
    /**
     * Name to apply to the distributed AMI.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags for the distribution configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DistributionConfiguration resource.
 */
export interface DistributionConfigurationArgs {
    /**
     * Description to apply to the distributed AMI.
     */
    description?: pulumi.Input<string>;
    /**
     * One or more configuration blocks with distribution settings. Detailed below.
     */
    distributions: pulumi.Input<pulumi.Input<inputs.imagebuilder.DistributionConfigurationDistribution>[]>;
    /**
     * Name to apply to the distributed AMI.
     */
    name?: pulumi.Input<string>;
    /**
     * Key-value map of resource tags for the distribution configuration. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
