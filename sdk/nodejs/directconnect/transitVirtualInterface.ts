// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect transit virtual interface resource.
 * A transit virtual interface is a VLAN that transports traffic from a Direct Connect gateway to one or more transit gateways.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleGateway = new aws.directconnect.Gateway("example", {
 *     amazonSideAsn: "64512",
 * });
 * const exampleTransitVirtualInterface = new aws.directconnect.TransitVirtualInterface("example", {
 *     addressFamily: "ipv4",
 *     bgpAsn: 65352,
 *     connectionId: aws_dx_connection_example.id,
 *     dxGatewayId: exampleGateway.id,
 *     vlan: 4094,
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/dx_transit_virtual_interface.html.markdown.
 */
export class TransitVirtualInterface extends pulumi.CustomResource {
    /**
     * Get an existing TransitVirtualInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions): TransitVirtualInterface {
        return new TransitVirtualInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/transitVirtualInterface:TransitVirtualInterface';

    /**
     * Returns true if the given object is an instance of TransitVirtualInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitVirtualInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitVirtualInterface.__pulumiType;
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    public readonly amazonAddress!: pulumi.Output<string>;
    public /*out*/ readonly amazonSideAsn!: pulumi.Output<string>;
    /**
     * The ARN of the virtual interface.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    public /*out*/ readonly awsDevice!: pulumi.Output<string>;
    /**
     * The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
     */
    public readonly bgpAsn!: pulumi.Output<number>;
    /**
     * The authentication key for BGP configuration.
     */
    public readonly bgpAuthKey!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
     */
    public readonly connectionId!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.
     */
    public readonly customerAddress!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    public readonly dxGatewayId!: pulumi.Output<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    public /*out*/ readonly jumboFrameCapable!: pulumi.Output<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    public readonly mtu!: pulumi.Output<number | undefined>;
    /**
     * The name for the virtual interface.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VLAN ID.
     */
    public readonly vlan!: pulumi.Output<number>;

    /**
     * Create a TransitVirtualInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitVirtualInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitVirtualInterfaceArgs | TransitVirtualInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TransitVirtualInterfaceState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["amazonSideAsn"] = state ? state.amazonSideAsn : undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["awsDevice"] = state ? state.awsDevice : undefined;
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
        } else {
            const args = argsOrState as TransitVirtualInterfaceArgs | undefined;
            if (!args || args.addressFamily === undefined) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if (!args || args.bgpAsn === undefined) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if (!args || args.connectionId === undefined) {
                throw new Error("Missing required property 'connectionId'");
            }
            if (!args || args.dxGatewayId === undefined) {
                throw new Error("Missing required property 'dxGatewayId'");
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
            inputs["amazonSideAsn"] = undefined /*out*/;
            inputs["arn"] = undefined /*out*/;
            inputs["awsDevice"] = undefined /*out*/;
            inputs["jumboFrameCapable"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TransitVirtualInterface.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitVirtualInterface resources.
 */
export interface TransitVirtualInterfaceState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    readonly amazonSideAsn?: pulumi.Input<string>;
    /**
     * The ARN of the virtual interface.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The Direct Connect endpoint on which the virtual interface terminates.
     */
    readonly awsDevice?: pulumi.Input<string>;
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
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    readonly dxGatewayId?: pulumi.Input<string>;
    /**
     * Indicates whether jumbo frames (8500 MTU) are supported.
     */
    readonly jumboFrameCapable?: pulumi.Input<boolean>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
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
 * The set of arguments for constructing a TransitVirtualInterface resource.
 */
export interface TransitVirtualInterfaceArgs {
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
     * The ID of the Direct Connect gateway to which to connect the virtual interface.
     */
    readonly dxGatewayId: pulumi.Input<string>;
    /**
     * The maximum transmission unit (MTU) is the size, in bytes, of the largest permissible packet that can be passed over the connection.
     * The MTU of a virtual transit interface can be either `1500` or `8500` (jumbo frames). Default is `1500`.
     */
    readonly mtu?: pulumi.Input<number>;
    /**
     * The name for the virtual interface.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VLAN ID.
     */
    readonly vlan: pulumi.Input<number>;
}
