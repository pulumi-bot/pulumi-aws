// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Direct Connect BGP peer resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const peer = new aws.directconnect.BgpPeer("peer", {
 *     virtualInterfaceId: aws_dx_private_virtual_interface.foo.id,
 *     addressFamily: "ipv6",
 *     bgpAsn: 65351,
 * });
 * ```
 */
export class BgpPeer extends pulumi.CustomResource {
    /**
     * Get an existing BgpPeer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpPeerState, opts?: pulumi.CustomResourceOptions): BgpPeer {
        return new BgpPeer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:directconnect/bgpPeer:BgpPeer';

    /**
     * Returns true if the given object is an instance of BgpPeer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpPeer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpPeer.__pulumiType;
    }

    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    public readonly amazonAddress!: pulumi.Output<string>;
    /**
     * The Direct Connect endpoint on which the BGP peer terminates.
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
     * The ID of the BGP peer.
     */
    public /*out*/ readonly bgpPeerId!: pulumi.Output<string>;
    /**
     * The Up/Down state of the BGP peer.
     */
    public /*out*/ readonly bgpStatus!: pulumi.Output<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    public readonly customerAddress!: pulumi.Output<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    public readonly virtualInterfaceId!: pulumi.Output<string>;

    /**
     * Create a BgpPeer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpPeerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpPeerArgs | BgpPeerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpPeerState | undefined;
            inputs["addressFamily"] = state ? state.addressFamily : undefined;
            inputs["amazonAddress"] = state ? state.amazonAddress : undefined;
            inputs["awsDevice"] = state ? state.awsDevice : undefined;
            inputs["bgpAsn"] = state ? state.bgpAsn : undefined;
            inputs["bgpAuthKey"] = state ? state.bgpAuthKey : undefined;
            inputs["bgpPeerId"] = state ? state.bgpPeerId : undefined;
            inputs["bgpStatus"] = state ? state.bgpStatus : undefined;
            inputs["customerAddress"] = state ? state.customerAddress : undefined;
            inputs["virtualInterfaceId"] = state ? state.virtualInterfaceId : undefined;
        } else {
            const args = argsOrState as BgpPeerArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.bgpAsn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bgpAsn'");
            }
            if ((!args || args.virtualInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualInterfaceId'");
            }
            inputs["addressFamily"] = args ? args.addressFamily : undefined;
            inputs["amazonAddress"] = args ? args.amazonAddress : undefined;
            inputs["bgpAsn"] = args ? args.bgpAsn : undefined;
            inputs["bgpAuthKey"] = args ? args.bgpAuthKey : undefined;
            inputs["customerAddress"] = args ? args.customerAddress : undefined;
            inputs["virtualInterfaceId"] = args ? args.virtualInterfaceId : undefined;
            inputs["awsDevice"] = undefined /*out*/;
            inputs["bgpPeerId"] = undefined /*out*/;
            inputs["bgpStatus"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BgpPeer.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpPeer resources.
 */
export interface BgpPeerState {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly amazonAddress?: pulumi.Input<string>;
    /**
     * The Direct Connect endpoint on which the BGP peer terminates.
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
     * The ID of the BGP peer.
     */
    readonly bgpPeerId?: pulumi.Input<string>;
    /**
     * The Up/Down state of the BGP peer.
     */
    readonly bgpStatus?: pulumi.Input<string>;
    /**
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    readonly virtualInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpPeer resource.
 */
export interface BgpPeerArgs {
    /**
     * The address family for the BGP peer. `ipv4 ` or `ipv6`.
     */
    readonly addressFamily: pulumi.Input<string>;
    /**
     * The IPv4 CIDR address to use to send traffic to Amazon.
     * Required for IPv4 BGP peers on public virtual interfaces.
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
     * The IPv4 CIDR destination address to which Amazon should send traffic.
     * Required for IPv4 BGP peers on public virtual interfaces.
     */
    readonly customerAddress?: pulumi.Input<string>;
    /**
     * The ID of the Direct Connect virtual interface on which to create the BGP peer.
     */
    readonly virtualInterfaceId: pulumi.Input<string>;
}
