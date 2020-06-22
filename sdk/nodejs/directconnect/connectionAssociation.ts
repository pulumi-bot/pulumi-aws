// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a Direct Connect Connection with a LAG.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const exampleConnection = new aws.directconnect.Connection("example", {
 *     bandwidth: "1Gbps",
 *     location: "EqSe2",
 * });
 * const exampleLinkAggregationGroup = new aws.directconnect.LinkAggregationGroup("example", {
 *     connectionsBandwidth: "1Gbps",
 *     location: "EqSe2",
 * });
 * const exampleConnectionAssociation = new aws.directconnect.ConnectionAssociation("example", {
 *     connectionId: exampleConnection.id,
 *     lagId: exampleLinkAggregationGroup.id,
 * });
 * ```
 */
export class ConnectionAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionAssociationState, opts?: pulumi.CustomResourceOptions): ConnectionAssociation {
        return new ConnectionAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/connectionAssociation:ConnectionAssociation';

    /**
     * Returns true if the given object is an instance of ConnectionAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionAssociation.__pulumiType;
    }

    /**
     * The ID of the connection.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * The ID of the LAG with which to associate the connection.
     */
    public readonly lagId!: pulumi.Output<string>;

    /**
     * Create a ConnectionAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionAssociationArgs | ConnectionAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectionAssociationState | undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["lagId"] = state ? state.lagId : undefined;
        } else {
            const args = argsOrState as ConnectionAssociationArgs | undefined;
            if (!args || args.connectionId === undefined) {
                throw new Error("Missing required property 'connectionId'");
            }
            if (!args || args.lagId === undefined) {
                throw new Error("Missing required property 'lagId'");
            }
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["lagId"] = args ? args.lagId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConnectionAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectionAssociation resources.
 */
export interface ConnectionAssociationState {
    /**
     * The ID of the connection.
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The ID of the LAG with which to associate the connection.
     */
    readonly lagId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectionAssociation resource.
 */
export interface ConnectionAssociationArgs {
    /**
     * The ID of the connection.
     */
    readonly connectionId: pulumi.Input<string>;
    /**
     * The ID of the LAG with which to associate the connection.
     */
    readonly lagId: pulumi.Input<string>;
}
