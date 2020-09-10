// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Connection extends pulumi.CustomResource {
    /**
     * Get an existing Connection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectionState, opts?: pulumi.CustomResourceOptions): Connection {
        return new Connection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/connection:Connection';

    /**
     * Returns true if the given object is an instance of Connection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Connection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Connection.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly awsDevice!: pulumi.Output<string>;
    public readonly bandwidth!: pulumi.Output<string>;
    public /*out*/ readonly hasLogicalRedundancy!: pulumi.Output<string>;
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    public readonly location!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a Connection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectionArgs | ConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ConnectionState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsDevice"] = state ? state.awsDevice : undefined;
            inputs["bandwidth"] = state ? state.bandwidth : undefined;
            inputs["hasLogicalRedundancy"] = state ? state.hasLogicalRedundancy : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ConnectionArgs | undefined;
            if (!args || args.bandwidth === undefined) {
                throw new Error("Missing required property 'bandwidth'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            inputs["bandwidth"] = args ? args.bandwidth : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["awsDevice"] = undefined /*out*/;
            inputs["hasLogicalRedundancy"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Connection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Connection resources.
 */
export interface ConnectionState {
    readonly arn?: pulumi.Input<string>;
    readonly awsDevice?: pulumi.Input<string>;
    readonly bandwidth?: pulumi.Input<string>;
    readonly hasLogicalRedundancy?: pulumi.Input<string>;
    readonly jumboFrameCapable?: pulumi.Input<boolean>;
    readonly location?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a Connection resource.
 */
export interface ConnectionArgs {
    readonly bandwidth: pulumi.Input<string>;
    readonly location: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
