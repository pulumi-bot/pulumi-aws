// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides an MQ Configuration Resource.
 *
 * For more information on Amazon MQ, see [Amazon MQ documentation](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/welcome.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.mq.Configuration("example", {
 *     data: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
 * <broker xmlns="http://activemq.apache.org/schema/core">
 *   <plugins>
 *     <forcePersistencyModeBrokerPlugin persistenceFlag="true"/>
 *     <statisticsBrokerPlugin/>
 *     <timeStampingBrokerPlugin ttlCeiling="86400000" zeroExpirationOverride="86400000"/>
 *   </plugins>
 * </broker>
 * `,
 *     description: "Example Configuration",
 *     engineType: "ActiveMQ",
 *     engineVersion: "5.15.0",
 * });
 * ```
 *
 * ## Import
 *
 * MQ Configurations can be imported using the configuration ID, e.g.
 *
 * ```sh
 *  $ pulumi import aws:mq/configuration:Configuration example c-0187d1eb-88c8-475a-9b79-16ef5a10c94f
 * ```
 */
export class Configuration extends pulumi.CustomResource {
    /**
     * Get an existing Configuration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationState, opts?: pulumi.CustomResourceOptions): Configuration {
        return new Configuration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mq/configuration:Configuration';

    /**
     * Returns true if the given object is an instance of Configuration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Configuration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Configuration.__pulumiType;
    }

    /**
     * The ARN of the configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The broker configuration in XML format.
     * See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
     * for supported parameters and format of the XML.
     */
    public readonly data!: pulumi.Output<string>;
    /**
     * The description of the configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The type of broker engine.
     */
    public readonly engineType!: pulumi.Output<string>;
    /**
     * The version of the broker engine.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * The latest revision of the configuration.
     */
    public /*out*/ readonly latestRevision!: pulumi.Output<number>;
    /**
     * The name of the configuration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Configuration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConfigurationArgs | ConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConfigurationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["data"] = state ? state.data : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["engineType"] = state ? state.engineType : undefined;
            inputs["engineVersion"] = state ? state.engineVersion : undefined;
            inputs["latestRevision"] = state ? state.latestRevision : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ConfigurationArgs | undefined;
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.engineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineType'");
            }
            if ((!args || args.engineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineVersion'");
            }
            inputs["data"] = args ? args.data : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["engineType"] = args ? args.engineType : undefined;
            inputs["engineVersion"] = args ? args.engineVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["latestRevision"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Configuration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Configuration resources.
 */
export interface ConfigurationState {
    /**
     * The ARN of the configuration.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The broker configuration in XML format.
     * See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
     * for supported parameters and format of the XML.
     */
    readonly data?: pulumi.Input<string>;
    /**
     * The description of the configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The type of broker engine.
     */
    readonly engineType?: pulumi.Input<string>;
    /**
     * The version of the broker engine.
     */
    readonly engineVersion?: pulumi.Input<string>;
    /**
     * The latest revision of the configuration.
     */
    readonly latestRevision?: pulumi.Input<number>;
    /**
     * The name of the configuration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Configuration resource.
 */
export interface ConfigurationArgs {
    /**
     * The broker configuration in XML format.
     * See [official docs](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)
     * for supported parameters and format of the XML.
     */
    readonly data: pulumi.Input<string>;
    /**
     * The description of the configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The type of broker engine.
     */
    readonly engineType: pulumi.Input<string>;
    /**
     * The version of the broker engine.
     */
    readonly engineVersion: pulumi.Input<string>;
    /**
     * The name of the configuration
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
