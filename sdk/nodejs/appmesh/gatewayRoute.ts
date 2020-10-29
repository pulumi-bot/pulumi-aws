// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Provides an AWS App Mesh gateway route resource.
 */
export class GatewayRoute extends pulumi.CustomResource {
    /**
     * Get an existing GatewayRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GatewayRouteState, opts?: pulumi.CustomResourceOptions): GatewayRoute {
        return new GatewayRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:appmesh/gatewayRoute:GatewayRoute';

    /**
     * Returns true if the given object is an instance of GatewayRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GatewayRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GatewayRoute.__pulumiType;
    }

    /**
     * The ARN of the gateway route.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The creation date of the gateway route.
     */
    public /*out*/ readonly createdDate!: pulumi.Output<string>;
    /**
     * The last update date of the gateway route.
     */
    public /*out*/ readonly lastUpdatedDate!: pulumi.Output<string>;
    /**
     * The name of the service mesh in which to create the gateway route.
     */
    public readonly meshName!: pulumi.Output<string>;
    /**
     * The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    public readonly meshOwner!: pulumi.Output<string>;
    /**
     * The name to use for the gateway route.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource owner's AWS account ID.
     */
    public /*out*/ readonly resourceOwner!: pulumi.Output<string>;
    /**
     * The gateway route specification to apply.
     */
    public readonly spec!: pulumi.Output<outputs.appmesh.GatewayRouteSpec>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
     */
    public readonly virtualGatewayName!: pulumi.Output<string>;

    /**
     * Create a GatewayRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GatewayRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GatewayRouteArgs | GatewayRouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GatewayRouteState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["createdDate"] = state ? state.createdDate : undefined;
            inputs["lastUpdatedDate"] = state ? state.lastUpdatedDate : undefined;
            inputs["meshName"] = state ? state.meshName : undefined;
            inputs["meshOwner"] = state ? state.meshOwner : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceOwner"] = state ? state.resourceOwner : undefined;
            inputs["spec"] = state ? state.spec : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualGatewayName"] = state ? state.virtualGatewayName : undefined;
        } else {
            const args = argsOrState as GatewayRouteArgs | undefined;
            if (!args || args.meshName === undefined) {
                throw new Error("Missing required property 'meshName'");
            }
            if (!args || args.spec === undefined) {
                throw new Error("Missing required property 'spec'");
            }
            if (!args || args.virtualGatewayName === undefined) {
                throw new Error("Missing required property 'virtualGatewayName'");
            }
            inputs["meshName"] = args ? args.meshName : undefined;
            inputs["meshOwner"] = args ? args.meshOwner : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["spec"] = args ? args.spec : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualGatewayName"] = args ? args.virtualGatewayName : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdDate"] = undefined /*out*/;
            inputs["lastUpdatedDate"] = undefined /*out*/;
            inputs["resourceOwner"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GatewayRoute.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GatewayRoute resources.
 */
export interface GatewayRouteState {
    /**
     * The ARN of the gateway route.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The creation date of the gateway route.
     */
    readonly createdDate?: pulumi.Input<string>;
    /**
     * The last update date of the gateway route.
     */
    readonly lastUpdatedDate?: pulumi.Input<string>;
    /**
     * The name of the service mesh in which to create the gateway route.
     */
    readonly meshName?: pulumi.Input<string>;
    /**
     * The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    readonly meshOwner?: pulumi.Input<string>;
    /**
     * The name to use for the gateway route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource owner's AWS account ID.
     */
    readonly resourceOwner?: pulumi.Input<string>;
    /**
     * The gateway route specification to apply.
     */
    readonly spec?: pulumi.Input<inputs.appmesh.GatewayRouteSpec>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
     */
    readonly virtualGatewayName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GatewayRoute resource.
 */
export interface GatewayRouteArgs {
    /**
     * The name of the service mesh in which to create the gateway route.
     */
    readonly meshName: pulumi.Input<string>;
    /**
     * The AWS account ID of the service mesh's owner. Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    readonly meshOwner?: pulumi.Input<string>;
    /**
     * The name to use for the gateway route.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The gateway route specification to apply.
     */
    readonly spec: pulumi.Input<inputs.appmesh.GatewayRouteSpec>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the [virtual gateway](https://www.terraform.io/docs/providers/aws/r/appmesh_virtual_gateway.html) to associate the gateway route with.
     */
    readonly virtualGatewayName: pulumi.Input<string>;
}
