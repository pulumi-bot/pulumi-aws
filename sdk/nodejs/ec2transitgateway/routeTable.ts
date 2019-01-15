// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class RouteTable extends pulumi.CustomResource {
    /**
     * Get an existing RouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableState, opts?: pulumi.CustomResourceOptions): RouteTable {
        return new RouteTable(name, <any>state, { ...opts, id: id });
    }

    public /*out*/ readonly defaultAssociationRouteTable: pulumi.Output<boolean>;
    public /*out*/ readonly defaultPropagationRouteTable: pulumi.Output<boolean>;
    public readonly tags: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly transitGatewayId: pulumi.Output<string>;

    /**
     * Create a RouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableArgs | RouteTableState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: RouteTableState = argsOrState as RouteTableState | undefined;
            inputs["defaultAssociationRouteTable"] = state ? state.defaultAssociationRouteTable : undefined;
            inputs["defaultPropagationRouteTable"] = state ? state.defaultPropagationRouteTable : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
        } else {
            const args = argsOrState as RouteTableArgs | undefined;
            if (!args || args.transitGatewayId === undefined) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            inputs["defaultAssociationRouteTable"] = undefined /*out*/;
            inputs["defaultPropagationRouteTable"] = undefined /*out*/;
        }
        super("aws:ec2transitgateway/routeTable:RouteTable", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTable resources.
 */
export interface RouteTableState {
    readonly defaultAssociationRouteTable?: pulumi.Input<boolean>;
    readonly defaultPropagationRouteTable?: pulumi.Input<boolean>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly transitGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTable resource.
 */
export interface RouteTableArgs {
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly transitGatewayId: pulumi.Input<string>;
}
