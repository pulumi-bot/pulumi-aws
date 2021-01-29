// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a SageMaker Feature Group resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.sagemaker.FeatureGroup("example", {
 *     featureGroupName: "example",
 *     recordIdentifierFeatureName: "example",
 *     eventTimeFeatureName: "example",
 *     roleArn: aws_iam_role.test.arn,
 *     featureDefinitions: [{
 *         featureName: "example",
 *         featureType: "String",
 *     }],
 *     onlineStoreConfig: {
 *         enableOnlineStore: true,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Feature Groups can be imported using the `name`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:sagemaker/featureGroup:FeatureGroup test_feature_group feature_group-foo
 * ```
 */
export class FeatureGroup extends pulumi.CustomResource {
    /**
     * Get an existing FeatureGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FeatureGroupState, opts?: pulumi.CustomResourceOptions): FeatureGroup {
        return new FeatureGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:sagemaker/featureGroup:FeatureGroup';

    /**
     * Returns true if the given object is an instance of FeatureGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeatureGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeatureGroup.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A free-form description of a Feature Group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the feature that stores the EventTime of a Record in a Feature Group.
     */
    public readonly eventTimeFeatureName!: pulumi.Output<string>;
    /**
     * A list of Feature names and types. See Feature Definition Below.
     */
    public readonly featureDefinitions!: pulumi.Output<outputs.sagemaker.FeatureGroupFeatureDefinition[]>;
    /**
     * The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
     */
    public readonly featureGroupName!: pulumi.Output<string>;
    /**
     * The Offline Feature Store Configuration. See Offline Store Config Below.
     */
    public readonly offlineStoreConfig!: pulumi.Output<outputs.sagemaker.FeatureGroupOfflineStoreConfig | undefined>;
    /**
     * The Online Feature Store Configuration. See Online Store Config Below.
     */
    public readonly onlineStoreConfig!: pulumi.Output<outputs.sagemaker.FeatureGroupOnlineStoreConfig | undefined>;
    /**
     * The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
     */
    public readonly recordIdentifierFeatureName!: pulumi.Output<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offlineStoreConfig` is provided.
     */
    public readonly roleArn!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a FeatureGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeatureGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FeatureGroupArgs | FeatureGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as FeatureGroupState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["eventTimeFeatureName"] = state ? state.eventTimeFeatureName : undefined;
            inputs["featureDefinitions"] = state ? state.featureDefinitions : undefined;
            inputs["featureGroupName"] = state ? state.featureGroupName : undefined;
            inputs["offlineStoreConfig"] = state ? state.offlineStoreConfig : undefined;
            inputs["onlineStoreConfig"] = state ? state.onlineStoreConfig : undefined;
            inputs["recordIdentifierFeatureName"] = state ? state.recordIdentifierFeatureName : undefined;
            inputs["roleArn"] = state ? state.roleArn : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as FeatureGroupArgs | undefined;
            if ((!args || args.eventTimeFeatureName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'eventTimeFeatureName'");
            }
            if ((!args || args.featureDefinitions === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'featureDefinitions'");
            }
            if ((!args || args.featureGroupName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'featureGroupName'");
            }
            if ((!args || args.recordIdentifierFeatureName === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'recordIdentifierFeatureName'");
            }
            if ((!args || args.roleArn === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'roleArn'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["eventTimeFeatureName"] = args ? args.eventTimeFeatureName : undefined;
            inputs["featureDefinitions"] = args ? args.featureDefinitions : undefined;
            inputs["featureGroupName"] = args ? args.featureGroupName : undefined;
            inputs["offlineStoreConfig"] = args ? args.offlineStoreConfig : undefined;
            inputs["onlineStoreConfig"] = args ? args.onlineStoreConfig : undefined;
            inputs["recordIdentifierFeatureName"] = args ? args.recordIdentifierFeatureName : undefined;
            inputs["roleArn"] = args ? args.roleArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(FeatureGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FeatureGroup resources.
 */
export interface FeatureGroupState {
    /**
     * The Amazon Resource Name (ARN) assigned by AWS to this feature_group.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A free-form description of a Feature Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the feature that stores the EventTime of a Record in a Feature Group.
     */
    readonly eventTimeFeatureName?: pulumi.Input<string>;
    /**
     * A list of Feature names and types. See Feature Definition Below.
     */
    readonly featureDefinitions?: pulumi.Input<pulumi.Input<inputs.sagemaker.FeatureGroupFeatureDefinition>[]>;
    /**
     * The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
     */
    readonly featureGroupName?: pulumi.Input<string>;
    /**
     * The Offline Feature Store Configuration. See Offline Store Config Below.
     */
    readonly offlineStoreConfig?: pulumi.Input<inputs.sagemaker.FeatureGroupOfflineStoreConfig>;
    /**
     * The Online Feature Store Configuration. See Online Store Config Below.
     */
    readonly onlineStoreConfig?: pulumi.Input<inputs.sagemaker.FeatureGroupOnlineStoreConfig>;
    /**
     * The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
     */
    readonly recordIdentifierFeatureName?: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offlineStoreConfig` is provided.
     */
    readonly roleArn?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a FeatureGroup resource.
 */
export interface FeatureGroupArgs {
    /**
     * A free-form description of a Feature Group.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the feature that stores the EventTime of a Record in a Feature Group.
     */
    readonly eventTimeFeatureName: pulumi.Input<string>;
    /**
     * A list of Feature names and types. See Feature Definition Below.
     */
    readonly featureDefinitions: pulumi.Input<pulumi.Input<inputs.sagemaker.FeatureGroupFeatureDefinition>[]>;
    /**
     * The name of the Feature Group. The name must be unique within an AWS Region in an AWS account.
     */
    readonly featureGroupName: pulumi.Input<string>;
    /**
     * The Offline Feature Store Configuration. See Offline Store Config Below.
     */
    readonly offlineStoreConfig?: pulumi.Input<inputs.sagemaker.FeatureGroupOfflineStoreConfig>;
    /**
     * The Online Feature Store Configuration. See Online Store Config Below.
     */
    readonly onlineStoreConfig?: pulumi.Input<inputs.sagemaker.FeatureGroupOnlineStoreConfig>;
    /**
     * The name of the Feature whose value uniquely identifies a Record defined in the Feature Store. Only the latest record per identifier value will be stored in the Online Store.
     */
    readonly recordIdentifierFeatureName: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the IAM execution role used to persist data into the Offline Store if an `offlineStoreConfig` is provided.
     */
    readonly roleArn: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}