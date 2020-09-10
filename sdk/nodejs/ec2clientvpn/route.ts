// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2clientvpn/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    public readonly clientVpnEndpointId!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly destinationCidrBlock!: pulumi.Output<string>;
    public /*out*/ readonly origin!: pulumi.Output<string>;
    public readonly targetVpcSubnetId!: pulumi.Output<string>;
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteState | undefined;
            inputs["clientVpnEndpointId"] = state ? state.clientVpnEndpointId : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["destinationCidrBlock"] = state ? state.destinationCidrBlock : undefined;
            inputs["origin"] = state ? state.origin : undefined;
            inputs["targetVpcSubnetId"] = state ? state.targetVpcSubnetId : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if (!args || args.clientVpnEndpointId === undefined) {
                throw new Error("Missing required property 'clientVpnEndpointId'");
            }
            if (!args || args.destinationCidrBlock === undefined) {
                throw new Error("Missing required property 'destinationCidrBlock'");
            }
            if (!args || args.targetVpcSubnetId === undefined) {
                throw new Error("Missing required property 'targetVpcSubnetId'");
            }
            inputs["clientVpnEndpointId"] = args ? args.clientVpnEndpointId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            inputs["targetVpcSubnetId"] = args ? args.targetVpcSubnetId : undefined;
            inputs["origin"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    readonly clientVpnEndpointId?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly destinationCidrBlock?: pulumi.Input<string>;
    readonly origin?: pulumi.Input<string>;
    readonly targetVpcSubnetId?: pulumi.Input<string>;
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    readonly clientVpnEndpointId: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly destinationCidrBlock: pulumi.Input<string>;
    readonly targetVpcSubnetId: pulumi.Input<string>;
}
