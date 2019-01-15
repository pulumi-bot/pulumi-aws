// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class PrivateVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing PrivateVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): PrivateVirtualInterface {
        return new PrivateVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    public readonly addressFamily: pulumi.Output<string>;
    public readonly amazonAddress: pulumi.Output<string>;
    public /*out*/ readonly arn: pulumi.Output<string>;
    public readonly bgpAsn: pulumi.Output<number>;
    public readonly bgpAuthKey: pulumi.Output<string>;
    public readonly connectionId: pulumi.Output<string>;
    public readonly customerAddress: pulumi.Output<string>;
    public readonly dxGatewayId: pulumi.Output<string | undefined>;
    public /*out*/ readonly jumboFrameCapable: pulumi.Output<boolean>;
    public readonly mtu: pulumi.Output<number | undefined>;
    public readonly name: pulumi.Output<string>;
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    public readonly vlan: pulumi.Output<number>;
    public readonly vpnGatewayId: pulumi.Output<string | undefined>;

    /**
     * Create a PrivateVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateVirtualInterfaceArgs | PrivateVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PrivateVirtualInterfaceState = argsOrState as PrivateVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["dxGatewayId"] = state ? state.dxGatewayId : undefined;
            inputs["jumboFrameCapable"] = state ? state.jumboFrameCapable : undefined;
            inputs["mtu"] = state ? state.mtu : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
            inputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as PrivateVirtualInterfaceArgs | undefined;
            if (!args || args.addressFamily === undefined) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if (!args || args.bgpAsn === undefined) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if (!args || args.connectionId === undefined) {
                throw new Error("Missing required property 'connectionId'");
            }
            if (!args || args.vlan === undefined) {
                throw new Error("Missing required property 'vlan'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            inputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            inputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            inputs["connectionId"] = args ? args.connectionId : undefined;
            inputs["customerAddress"] = args ? args.customerAddress : undefined;
            inputs["dxGatewayId"] = args ? args.dxGatewayId : undefined;
            inputs["mtu"] = args ? args.mtu : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        super("aws:directconnect/privateVirtualInterface:PrivateVirtualInterface", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateVirtualInterface resources.
 */
export interface PrivateVirtualInterfaceState {
    readonly addressFamily?: pulumi.Input<string>;
    readonly amazonAddress?: pulumi.Input<string>;
    readonly arn?: pulumi.Input<string>;
    readonly bgpAsn?: pulumi.Input<number>;
    readonly bgpAuthKey?: pulumi.Input<string>;
    readonly connectionId?: pulumi.Input<string>;
    readonly customerAddress?: pulumi.Input<string>;
    readonly dxGatewayId?: pulumi.Input<string>;
    readonly jumboFrameCapable?: pulumi.Input<boolean>;
    readonly mtu?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly vlan?: pulumi.Input<number>;
    readonly vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PrivateVirtualInterface resource.
 */
export interface PrivateVirtualInterfaceArgs {
    readonly addressFamily: pulumi.Input<string>;
    readonly amazonAddress?: pulumi.Input<string>;
    readonly bgpAsn: pulumi.Input<number>;
    readonly bgpAuthKey?: pulumi.Input<string>;
    readonly connectionId: pulumi.Input<string>;
    readonly customerAddress?: pulumi.Input<string>;
    readonly dxGatewayId?: pulumi.Input<string>;
    readonly mtu?: pulumi.Input<number>;
    readonly name?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    readonly vlan: pulumi.Input<number>;
    readonly vpnGatewayId?: pulumi.Input<string>;
}
