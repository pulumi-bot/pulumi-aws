// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect public virtual interface resource.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const foo = new aws.directconnect.PublicVirtualInterface("foo", {
 *     addressFamily: "ipv4",
 *     amazonAddress: "175.45.176.2/30",
 *     bgpAsn: 65352,
 *     connectionId: "dxcon-zzzzzzzz",
 *     customerAddress: "175.45.176.1/30",
 *     routeFilterPrefixes: [
 *         "210.52.109.0/24",
 *         "175.45.176.0/22",
 *     ],
 *     vlan: 4094,
 * });
 * ```
 */
export class PublicVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing PublicVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PublicVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): PublicVirtualInterface {
        return new PublicVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    public readonly amazonAddress: pulumi.Output<string>;
    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    public readonly connectionId: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    public readonly customerAddress: pulumi.Output<string>;
    /**
     * The name for the virtual interface.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    public readonly routeFilterPrefixes: pulumi.Output<string[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VLAN ID.
     */
    public readonly vlan: pulumi.Output<number>;

    /**
     * Create a PublicVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublicVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PublicVirtualInterfaceArgs | PublicVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: PublicVirtualInterfaceState = argsOrState as PublicVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["connectionId"] = state ? state.connectionId : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["routeFilterPrefixes"] = state ? state.routeFilterPrefixes : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vlan"] = state ? state.vlan : undefined;
        } else {
            const args = argsOrState as PublicVirtualInterfaceArgs | undefined;
            if (!args || args.addressFamily === undefined) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if (!args || args.bgpAsn === undefined) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if (!args || args.connectionId === undefined) {
                throw new Error("Missing required property 'connectionId'");
            }
            if (!args || args.routeFilterPrefixes === undefined) {
                throw new Error("Missing required property 'routeFilterPrefixes'");
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
            inputs["name"] = args ? args.name : undefined;
            inputs["routeFilterPrefixes"] = args ? args.routeFilterPrefixes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vlan"] = args ? args.vlan : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        super("aws:directconnect/publicVirtualInterface:PublicVirtualInterface", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PublicVirtualInterface resources.
 */
export interface PublicVirtualInterfaceState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn?: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    readonly connectionId?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    readonly routeFilterPrefixes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VLAN ID.
     */
    readonly vlan?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PublicVirtualInterface resource.
 */
export interface PublicVirtualInterfaceArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    readonly bgpAsn: pulumi.Input<number>;
    /**
     * The authentication key for BGP configuration.
     */
    readonly bgpAuthKey?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    readonly connectionId: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of routes to be advertised to the AWS network in this region.
     */
    readonly routeFilterPrefixes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VLAN ID.
     */
    readonly vlan: pulumi.Input<number>;
}
