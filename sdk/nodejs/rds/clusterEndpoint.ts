// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a RDS Aurora Cluster Endpoint.
 * You can refer to the [User Guide][1].
 * 
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const defaultCluster = new aws.rds.Cluster("default", {
 *     availabilityZones: [
 *         "us-west-2a",
 *         "us-west-2b",
 *         "us-west-2c",
 *     ],
 *     backupRetentionPeriod: 5,
 *     clusterIdentifier: "aurora-cluster-demo",
 *     databaseName: "mydb",
 *     masterPassword: "bar",
 *     masterUsername: "foo",
 *     preferredBackupWindow: "07:00-09:00",
 * });
 * const test1 = new aws.rds.ClusterInstance("test1", {
 *     applyImmediately: true,
 *     clusterIdentifier: defaultCluster.id,
 *     identifier: "test1",
 *     instanceClass: "db.t2.small",
 * });
 * const test2 = new aws.rds.ClusterInstance("test2", {
 *     applyImmediately: true,
 *     clusterIdentifier: defaultCluster.id,
 *     identifier: "test2",
 *     instanceClass: "db.t2.small",
 * });
 * const test3 = new aws.rds.ClusterInstance("test3", {
 *     applyImmediately: true,
 *     clusterIdentifier: defaultCluster.id,
 *     identifier: "test3",
 *     instanceClass: "db.t2.small",
 * });
 * const eligible = new aws.rds.ClusterEndpoint("eligible", {
 *     clusterEndpointIdentifier: "reader",
 *     clusterIdentifier: defaultCluster.id,
 *     customEndpointType: "READER",
 *     excludedMembers: [
 *         test1.id,
 *         test2.id,
 *     ],
 * });
 * const static = new aws.rds.ClusterEndpoint("static", {
 *     clusterEndpointIdentifier: "static",
 *     clusterIdentifier: defaultCluster.id,
 *     customEndpointType: "READER",
 *     staticMembers: [
 *         test1.id,
 *         test3.id,
 *     ],
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/rds_cluster_endpoint.html.markdown.
 */
export class ClusterEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing ClusterEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterEndpointState, opts?: pulumi.CustomResourceOptions): ClusterEndpoint {
        return new ClusterEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/clusterEndpoint:ClusterEndpoint';

    /**
     * Returns true if the given object is an instance of ClusterEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterEndpoint.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of cluster
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
     */
    public readonly clusterEndpointIdentifier!: pulumi.Output<string>;
    /**
     * The cluster identifier.
     */
    public readonly clusterIdentifier!: pulumi.Output<string>;
    /**
     * The type of the endpoint. One of: READER , ANY .
     */
    public readonly customEndpointType!: pulumi.Output<string>;
    /**
     * A custom endpoint for the Aurora cluster
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
     */
    public readonly excludedMembers!: pulumi.Output<string[] | undefined>;
    /**
     * List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
     */
    public readonly staticMembers!: pulumi.Output<string[] | undefined>;
    /**
     * Key-value mapping of resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a ClusterEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterEndpointArgs | ClusterEndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClusterEndpointState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["clusterEndpointIdentifier"] = state ? state.clusterEndpointIdentifier : undefined;
            inputs["clusterIdentifier"] = state ? state.clusterIdentifier : undefined;
            inputs["customEndpointType"] = state ? state.customEndpointType : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["excludedMembers"] = state ? state.excludedMembers : undefined;
            inputs["staticMembers"] = state ? state.staticMembers : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ClusterEndpointArgs | undefined;
            if (!args || args.clusterEndpointIdentifier === undefined) {
                throw new Error("Missing required property 'clusterEndpointIdentifier'");
            }
            if (!args || args.clusterIdentifier === undefined) {
                throw new Error("Missing required property 'clusterIdentifier'");
            }
            if (!args || args.customEndpointType === undefined) {
                throw new Error("Missing required property 'customEndpointType'");
            }
            inputs["clusterEndpointIdentifier"] = args ? args.clusterEndpointIdentifier : undefined;
            inputs["clusterIdentifier"] = args ? args.clusterIdentifier : undefined;
            inputs["customEndpointType"] = args ? args.customEndpointType : undefined;
            inputs["excludedMembers"] = args ? args.excludedMembers : undefined;
            inputs["staticMembers"] = args ? args.staticMembers : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClusterEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterEndpoint resources.
 */
export interface ClusterEndpointState {
    /**
     * Amazon Resource Name (ARN) of cluster
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
     */
    readonly clusterEndpointIdentifier?: pulumi.Input<string>;
    /**
     * The cluster identifier.
     */
    readonly clusterIdentifier?: pulumi.Input<string>;
    /**
     * The type of the endpoint. One of: READER , ANY .
     */
    readonly customEndpointType?: pulumi.Input<string>;
    /**
     * A custom endpoint for the Aurora cluster
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
     */
    readonly excludedMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
     */
    readonly staticMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a ClusterEndpoint resource.
 */
export interface ClusterEndpointArgs {
    /**
     * The identifier to use for the new endpoint. This parameter is stored as a lowercase string.
     */
    readonly clusterEndpointIdentifier: pulumi.Input<string>;
    /**
     * The cluster identifier.
     */
    readonly clusterIdentifier: pulumi.Input<string>;
    /**
     * The type of the endpoint. One of: READER , ANY .
     */
    readonly customEndpointType: pulumi.Input<string>;
    /**
     * List of DB instance identifiers that aren't part of the custom endpoint group. All other eligible instances are reachable through the custom endpoint. Only relevant if the list of static members is empty. Conflicts with `staticMembers`.
     */
    readonly excludedMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of DB instance identifiers that are part of the custom endpoint group. Conflicts with `excludedMembers`.
     */
    readonly staticMembers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value mapping of resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
