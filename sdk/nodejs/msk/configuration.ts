// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an Amazon Managed Streaming for Kafka configuration. More information can be found on the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration.html).
 * 
 * > **NOTE:** The API does not support deleting MSK configurations. Removing this resource will only remove the this provider state for it.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/msk_configuration.html.markdown.
 */
export class Configuration extends pulumi.CustomResource {
    /**
     * Get an existing Configuration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConfigurationState, opts?: pulumi.CustomResourceOptions): Configuration {
        return new Configuration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:msk/configuration:Configuration';

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
     * Amazon Resource Name (ARN) of the configuration.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the configuration.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of Apache Kafka versions which can use this configuration.
     */
    public readonly kafkaVersions!: pulumi.Output<string[]>;
    /**
     * Latest revision of the configuration.
     */
    public /*out*/ readonly latestRevision!: pulumi.Output<number>;
    /**
     * Name of the configuration.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
     */
    public readonly serverProperties!: pulumi.Output<string>;

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
        if (opts && opts.id) {
            const state = argsOrState as ConfigurationState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["kafkaVersions"] = state ? state.kafkaVersions : undefined;
            inputs["latestRevision"] = state ? state.latestRevision : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["serverProperties"] = state ? state.serverProperties : undefined;
        } else {
            const args = argsOrState as ConfigurationArgs | undefined;
            if (!args || args.kafkaVersions === undefined) {
                throw new Error("Missing required property 'kafkaVersions'");
            }
            if (!args || args.serverProperties === undefined) {
                throw new Error("Missing required property 'serverProperties'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["kafkaVersions"] = args ? args.kafkaVersions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["serverProperties"] = args ? args.serverProperties : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["latestRevision"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Configuration.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Configuration resources.
 */
export interface ConfigurationState {
    /**
     * Amazon Resource Name (ARN) of the configuration.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Description of the configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of Apache Kafka versions which can use this configuration.
     */
    readonly kafkaVersions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Latest revision of the configuration.
     */
    readonly latestRevision?: pulumi.Input<number>;
    /**
     * Name of the configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
     */
    readonly serverProperties?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Configuration resource.
 */
export interface ConfigurationArgs {
    /**
     * Description of the configuration.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * List of Apache Kafka versions which can use this configuration.
     */
    readonly kafkaVersions: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the configuration.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Contents of the server.properties file. Supported properties are documented in the [MSK Developer Guide](https://docs.aws.amazon.com/msk/latest/developerguide/msk-configuration-properties.html).
     */
    readonly serverProperties: pulumi.Input<string>;
}
