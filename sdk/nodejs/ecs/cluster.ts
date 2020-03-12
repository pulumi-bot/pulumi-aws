// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an ECS cluster.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.ecs.Cluster("foo", {});
 * ```
 * {{% /examples %}}
 * 
 * ## setting
 * 
 * The `setting` configuration block supports the following:
 * 
 * * `name` - (Required) Name of the setting to manage. Valid values: `containerInsights`.
 * * `value` -  (Required) The value to assign to the setting. Value values are `enabled` and `disabled`.
 * 
 * ## defaultCapacityProviderStrategy
 * 
 * The `defaultCapacityProviderStrategy` configuration block supports the following:
 * 
 * * `capacityProvider` - (Required) The short name of the capacity provider.
 * * `weight` - (Optional) The relative percentage of the total number of launched tasks that should use the specified capacity provider.
 * * `base` - (Optional) The number of tasks, at a minimum, to run on the specified capacity provider. Only one capacity provider in a capacity provider strategy can have a base defined.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ecs_cluster.html.markdown.
 */
export class Cluster extends pulumi.CustomResource {
    /**
     * Get an existing Cluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterState, opts?: pulumi.CustomResourceOptions): Cluster {
        return new Cluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ecs/cluster:Cluster';

    /**
     * Returns true if the given object is an instance of Cluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cluster.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that identifies the cluster
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     */
    public readonly capacityProviders!: pulumi.Output<string[] | undefined>;
    /**
     * The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
     */
    public readonly defaultCapacityProviderStrategies!: pulumi.Output<outputs.ecs.ClusterDefaultCapacityProviderStrategy[] | undefined>;
    /**
     * The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
     */
    public readonly settings!: pulumi.Output<outputs.ecs.ClusterSetting[]>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Cluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterArgs | ClusterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["capacityProviders"] = state ? state.capacityProviders : undefined;
            inputs["defaultCapacityProviderStrategies"] = state ? state.defaultCapacityProviderStrategies : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["settings"] = state ? state.settings : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ClusterArgs | undefined;
            inputs["capacityProviders"] = args ? args.capacityProviders : undefined;
            inputs["defaultCapacityProviderStrategies"] = args ? args.defaultCapacityProviderStrategies : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Cluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Cluster resources.
 */
export interface ClusterState {
    /**
     * The Amazon Resource Name (ARN) that identifies the cluster
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     */
    readonly capacityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
     */
    readonly defaultCapacityProviderStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterDefaultCapacityProviderStrategy>[]>;
    /**
     * The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
     */
    readonly settings?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterSetting>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Cluster resource.
 */
export interface ClusterArgs {
    /**
     * List of short names of one or more capacity providers to associate with the cluster. Valid values also include `FARGATE` and `FARGATE_SPOT`.
     */
    readonly capacityProviders?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The capacity provider strategy to use by default for the cluster. Can be one or more.  Defined below.
     */
    readonly defaultCapacityProviderStrategies?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterDefaultCapacityProviderStrategy>[]>;
    /**
     * The name of the cluster (up to 255 letters, numbers, hyphens, and underscores)
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Configuration block(s) with cluster settings. For example, this can be used to enable CloudWatch Container Insights for a cluster. Defined below.
     */
    readonly settings?: pulumi.Input<pulumi.Input<inputs.ecs.ClusterSetting>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
