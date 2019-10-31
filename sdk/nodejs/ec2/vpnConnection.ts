// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an EC2 VPN connection. These objects can be connected to customer gateways, and allow you to establish tunnels between your network and Amazon.
 * 
 * > **Note:** All arguments including `tunnel1PresharedKey` and `tunnel2PresharedKey` will be stored in the raw state as plain-text.
 * [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
 * 
 * > **Note:** The CIDR blocks in the arguments `tunnel1InsideCidr` and `tunnel2InsideCidr` must have a prefix of /30 and be a part of a specific range.
 * [Read more about this in the AWS documentation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_VpnTunnelOptionsSpecification.html).
 * 
 * ## Example Usage
 * 
 * ### EC2 Transit Gateway
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const exampleTransitGateway = new aws.ec2transitgateway.TransitGateway("example", {});
 * const exampleCustomerGateway = new aws.ec2.CustomerGateway("example", {
 *     bgpAsn: 65000,
 *     ipAddress: "172.0.0.1",
 *     type: "ipsec.1",
 * });
 * const exampleVpnConnection = new aws.ec2.VpnConnection("example", {
 *     customerGatewayId: exampleCustomerGateway.id,
 *     transitGatewayId: exampleTransitGateway.id,
 *     type: exampleCustomerGateway.type,
 * });
 * ```
 * 
 * ### Virtual Private Gateway
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const vpc = new aws.ec2.Vpc("vpc", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const vpnGateway = new aws.ec2.VpnGateway("vpnGateway", {
 *     vpcId: vpc.id,
 * });
 * const customerGateway = new aws.ec2.CustomerGateway("customerGateway", {
 *     bgpAsn: 65000,
 *     ipAddress: "172.0.0.1",
 *     type: "ipsec.1",
 * });
 * const main = new aws.ec2.VpnConnection("main", {
 *     customerGatewayId: customerGateway.id,
 *     staticRoutesOnly: true,
 *     type: "ipsec.1",
 *     vpnGatewayId: vpnGateway.id,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpn_connection.html.markdown.
 */
export class VpnConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpnConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnConnectionState, opts?: pulumi.CustomResourceOptions): VpnConnection {
        return new VpnConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpnConnection:VpnConnection';

    /**
     * Returns true if the given object is an instance of VpnConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnConnection.__pulumiType;
    }

    /**
     * The configuration information for the VPN connection's customer gateway (in the native XML format).
     */
    public /*out*/ readonly customerGatewayConfiguration!: pulumi.Output<string>;
    /**
     * The ID of the customer gateway.
     */
    public readonly customerGatewayId!: pulumi.Output<string>;
    public /*out*/ readonly routes!: pulumi.Output<outputs.ec2.VpnConnectionRoute[]>;
    /**
     * Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
     */
    public readonly staticRoutesOnly!: pulumi.Output<boolean>;
    /**
     * Tags to apply to the connection.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * When associated with an EC2 Transit Gateway (`transitGatewayId` argument), the attachment ID.
     */
    public /*out*/ readonly transitGatewayAttachmentId!: pulumi.Output<string>;
    /**
     * The ID of the EC2 Transit Gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string | undefined>;
    /**
     * The public IP address of the first VPN tunnel.
     */
    public /*out*/ readonly tunnel1Address!: pulumi.Output<string>;
    /**
     * The bgp asn number of the first VPN tunnel.
     */
    public /*out*/ readonly tunnel1BgpAsn!: pulumi.Output<string>;
    /**
     * The bgp holdtime of the first VPN tunnel.
     */
    public /*out*/ readonly tunnel1BgpHoldtime!: pulumi.Output<number>;
    /**
     * The RFC 6890 link-local address of the first VPN tunnel (Customer Gateway Side).
     */
    public /*out*/ readonly tunnel1CgwInsideAddress!: pulumi.Output<string>;
    /**
     * The CIDR block of the inside IP addresses for the first VPN tunnel.
     */
    public readonly tunnel1InsideCidr!: pulumi.Output<string>;
    /**
     * The preshared key of the first VPN tunnel.
     */
    public readonly tunnel1PresharedKey!: pulumi.Output<string>;
    /**
     * The RFC 6890 link-local address of the first VPN tunnel (VPN Gateway Side).
     */
    public /*out*/ readonly tunnel1VgwInsideAddress!: pulumi.Output<string>;
    /**
     * The public IP address of the second VPN tunnel.
     */
    public /*out*/ readonly tunnel2Address!: pulumi.Output<string>;
    /**
     * The bgp asn number of the second VPN tunnel.
     */
    public /*out*/ readonly tunnel2BgpAsn!: pulumi.Output<string>;
    /**
     * The bgp holdtime of the second VPN tunnel.
     */
    public /*out*/ readonly tunnel2BgpHoldtime!: pulumi.Output<number>;
    /**
     * The RFC 6890 link-local address of the second VPN tunnel (Customer Gateway Side).
     */
    public /*out*/ readonly tunnel2CgwInsideAddress!: pulumi.Output<string>;
    /**
     * The CIDR block of the inside IP addresses for the second VPN tunnel.
     */
    public readonly tunnel2InsideCidr!: pulumi.Output<string>;
    /**
     * The preshared key of the second VPN tunnel.
     */
    public readonly tunnel2PresharedKey!: pulumi.Output<string>;
    /**
     * The RFC 6890 link-local address of the second VPN tunnel (VPN Gateway Side).
     */
    public /*out*/ readonly tunnel2VgwInsideAddress!: pulumi.Output<string>;
    /**
     * The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
     */
    public readonly type!: pulumi.Output<string>;
    public /*out*/ readonly vgwTelemetries!: pulumi.Output<outputs.ec2.VpnConnectionVgwTelemetry[]>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string | undefined>;

    /**
     * Create a VpnConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnConnectionArgs | VpnConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpnConnectionState | undefined;
            inputs["customerGatewayConfiguration"] = state ? state.customerGatewayConfiguration : undefined;
            inputs["customerGatewayId"] = state ? state.customerGatewayId : undefined;
            inputs["routes"] = state ? state.routes : undefined;
            inputs["staticRoutesOnly"] = state ? state.staticRoutesOnly : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
            inputs["tunnel1Address"] = state ? state.tunnel1Address : undefined;
            inputs["tunnel1BgpAsn"] = state ? state.tunnel1BgpAsn : undefined;
            inputs["tunnel1BgpHoldtime"] = state ? state.tunnel1BgpHoldtime : undefined;
            inputs["tunnel1CgwInsideAddress"] = state ? state.tunnel1CgwInsideAddress : undefined;
            inputs["tunnel1InsideCidr"] = state ? state.tunnel1InsideCidr : undefined;
            inputs["tunnel1PresharedKey"] = state ? state.tunnel1PresharedKey : undefined;
            inputs["tunnel1VgwInsideAddress"] = state ? state.tunnel1VgwInsideAddress : undefined;
            inputs["tunnel2Address"] = state ? state.tunnel2Address : undefined;
            inputs["tunnel2BgpAsn"] = state ? state.tunnel2BgpAsn : undefined;
            inputs["tunnel2BgpHoldtime"] = state ? state.tunnel2BgpHoldtime : undefined;
            inputs["tunnel2CgwInsideAddress"] = state ? state.tunnel2CgwInsideAddress : undefined;
            inputs["tunnel2InsideCidr"] = state ? state.tunnel2InsideCidr : undefined;
            inputs["tunnel2PresharedKey"] = state ? state.tunnel2PresharedKey : undefined;
            inputs["tunnel2VgwInsideAddress"] = state ? state.tunnel2VgwInsideAddress : undefined;
            inputs["type"] = state ? state.type : undefined;
            inputs["vgwTelemetries"] = state ? state.vgwTelemetries : undefined;
            inputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as VpnConnectionArgs | undefined;
            if (!args || args.customerGatewayId === undefined) {
                throw new Error("Missing required property 'customerGatewayId'");
            }
            if (!args || args.type === undefined) {
                throw new Error("Missing required property 'type'");
            }
            inputs["customerGatewayId"] = args ? args.customerGatewayId : undefined;
            inputs["staticRoutesOnly"] = args ? args.staticRoutesOnly : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            inputs["tunnel1InsideCidr"] = args ? args.tunnel1InsideCidr : undefined;
            inputs["tunnel1PresharedKey"] = args ? args.tunnel1PresharedKey : undefined;
            inputs["tunnel2InsideCidr"] = args ? args.tunnel2InsideCidr : undefined;
            inputs["tunnel2PresharedKey"] = args ? args.tunnel2PresharedKey : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
            inputs["customerGatewayConfiguration"] = undefined /*out*/;
            inputs["routes"] = undefined /*out*/;
            inputs["transitGatewayAttachmentId"] = undefined /*out*/;
            inputs["tunnel1Address"] = undefined /*out*/;
            inputs["tunnel1BgpAsn"] = undefined /*out*/;
            inputs["tunnel1BgpHoldtime"] = undefined /*out*/;
            inputs["tunnel1CgwInsideAddress"] = undefined /*out*/;
            inputs["tunnel1VgwInsideAddress"] = undefined /*out*/;
            inputs["tunnel2Address"] = undefined /*out*/;
            inputs["tunnel2BgpAsn"] = undefined /*out*/;
            inputs["tunnel2BgpHoldtime"] = undefined /*out*/;
            inputs["tunnel2CgwInsideAddress"] = undefined /*out*/;
            inputs["tunnel2VgwInsideAddress"] = undefined /*out*/;
            inputs["vgwTelemetries"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpnConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnConnection resources.
 */
export interface VpnConnectionState {
    /**
     * The configuration information for the VPN connection's customer gateway (in the native XML format).
     */
    readonly customerGatewayConfiguration?: pulumi.Input<string>;
    /**
     * The ID of the customer gateway.
     */
    readonly customerGatewayId?: pulumi.Input<string>;
    readonly routes?: pulumi.Input<pulumi.Input<inputs.ec2.VpnConnectionRoute>[]>;
    /**
     * Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
     */
    readonly staticRoutesOnly?: pulumi.Input<boolean>;
    /**
     * Tags to apply to the connection.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * When associated with an EC2 Transit Gateway (`transitGatewayId` argument), the attachment ID.
     */
    readonly transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * The ID of the EC2 Transit Gateway.
     */
    readonly transitGatewayId?: pulumi.Input<string>;
    /**
     * The public IP address of the first VPN tunnel.
     */
    readonly tunnel1Address?: pulumi.Input<string>;
    /**
     * The bgp asn number of the first VPN tunnel.
     */
    readonly tunnel1BgpAsn?: pulumi.Input<string>;
    /**
     * The bgp holdtime of the first VPN tunnel.
     */
    readonly tunnel1BgpHoldtime?: pulumi.Input<number>;
    /**
     * The RFC 6890 link-local address of the first VPN tunnel (Customer Gateway Side).
     */
    readonly tunnel1CgwInsideAddress?: pulumi.Input<string>;
    /**
     * The CIDR block of the inside IP addresses for the first VPN tunnel.
     */
    readonly tunnel1InsideCidr?: pulumi.Input<string>;
    /**
     * The preshared key of the first VPN tunnel.
     */
    readonly tunnel1PresharedKey?: pulumi.Input<string>;
    /**
     * The RFC 6890 link-local address of the first VPN tunnel (VPN Gateway Side).
     */
    readonly tunnel1VgwInsideAddress?: pulumi.Input<string>;
    /**
     * The public IP address of the second VPN tunnel.
     */
    readonly tunnel2Address?: pulumi.Input<string>;
    /**
     * The bgp asn number of the second VPN tunnel.
     */
    readonly tunnel2BgpAsn?: pulumi.Input<string>;
    /**
     * The bgp holdtime of the second VPN tunnel.
     */
    readonly tunnel2BgpHoldtime?: pulumi.Input<number>;
    /**
     * The RFC 6890 link-local address of the second VPN tunnel (Customer Gateway Side).
     */
    readonly tunnel2CgwInsideAddress?: pulumi.Input<string>;
    /**
     * The CIDR block of the inside IP addresses for the second VPN tunnel.
     */
    readonly tunnel2InsideCidr?: pulumi.Input<string>;
    /**
     * The preshared key of the second VPN tunnel.
     */
    readonly tunnel2PresharedKey?: pulumi.Input<string>;
    /**
     * The RFC 6890 link-local address of the second VPN tunnel (VPN Gateway Side).
     */
    readonly tunnel2VgwInsideAddress?: pulumi.Input<string>;
    /**
     * The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
     */
    readonly type?: pulumi.Input<string>;
    readonly vgwTelemetries?: pulumi.Input<pulumi.Input<inputs.ec2.VpnConnectionVgwTelemetry>[]>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    readonly vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnConnection resource.
 */
export interface VpnConnectionArgs {
    /**
     * The ID of the customer gateway.
     */
    readonly customerGatewayId: pulumi.Input<string>;
    /**
     * Whether the VPN connection uses static routes exclusively. Static routes must be used for devices that don't support BGP.
     */
    readonly staticRoutesOnly?: pulumi.Input<boolean>;
    /**
     * Tags to apply to the connection.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the EC2 Transit Gateway.
     */
    readonly transitGatewayId?: pulumi.Input<string>;
    /**
     * The CIDR block of the inside IP addresses for the first VPN tunnel.
     */
    readonly tunnel1InsideCidr?: pulumi.Input<string>;
    /**
     * The preshared key of the first VPN tunnel.
     */
    readonly tunnel1PresharedKey?: pulumi.Input<string>;
    /**
     * The CIDR block of the inside IP addresses for the second VPN tunnel.
     */
    readonly tunnel2InsideCidr?: pulumi.Input<string>;
    /**
     * The preshared key of the second VPN tunnel.
     */
    readonly tunnel2PresharedKey?: pulumi.Input<string>;
    /**
     * The type of VPN connection. The only type AWS supports at this time is "ipsec.1".
     */
    readonly type: pulumi.Input<string>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    readonly vpnGatewayId?: pulumi.Input<string>;
}
