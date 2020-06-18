// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC DHCP Options resource.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const dnsResolver = new aws.ec2.VpcDhcpOptions("dns_resolver", {
 *     domainNameServers: [
 *         "8.8.8.8",
 *         "8.8.4.4",
 *     ],
 * });
 * ```
 *
 * Full usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ec2.VpcDhcpOptions("foo", {
 *     domainName: "service.consul",
 *     domainNameServers: [
 *         "127.0.0.1",
 *         "10.0.0.2",
 *     ],
 *     netbiosNameServers: ["127.0.0.1"],
 *     netbiosNodeType: "2",
 *     ntpServers: ["127.0.0.1"],
 *     tags: {
 *         Name: "foo-name",
 *     },
 * });
 * ```
 * ## Remarks
 *
 * * Notice that all arguments are optional but you have to specify at least one argument.
 * * `domainNameServers`, `netbiosNameServers`, `ntpServers` are limited by AWS to maximum four servers only.
 * * To actually use the DHCP Options Set you need to associate it to a VPC using `aws.ec2.VpcDhcpOptionsAssociation`.
 * * If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
 * * In most cases unless you're configuring your own DNS you'll want to set `domainNameServers` to `AmazonProvidedDNS`.
 */
export class VpcDhcpOptions extends pulumi.CustomResource {
    /**
     * Get an existing VpcDhcpOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcDhcpOptionsState, opts?: pulumi.CustomResourceOptions): VpcDhcpOptions {
        return new VpcDhcpOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpcDhcpOptions:VpcDhcpOptions';

    /**
     * Returns true if the given object is an instance of VpcDhcpOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcDhcpOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcDhcpOptions.__pulumiType;
    }

    /**
     * the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
     */
    public readonly domainName!: pulumi.Output<string | undefined>;
    /**
     * List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
     */
    public readonly domainNameServers!: pulumi.Output<string[] | undefined>;
    /**
     * List of NETBIOS name servers.
     */
    public readonly netbiosNameServers!: pulumi.Output<string[] | undefined>;
    /**
     * The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     */
    public readonly netbiosNodeType!: pulumi.Output<string | undefined>;
    /**
     * List of NTP servers to configure.
     */
    public readonly ntpServers!: pulumi.Output<string[] | undefined>;
    /**
     * The ID of the AWS account that owns the DHCP options set.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a VpcDhcpOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcDhcpOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcDhcpOptionsArgs | VpcDhcpOptionsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcDhcpOptionsState | undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["domainNameServers"] = state ? state.domainNameServers : undefined;
            inputs["netbiosNameServers"] = state ? state.netbiosNameServers : undefined;
            inputs["netbiosNodeType"] = state ? state.netbiosNodeType : undefined;
            inputs["ntpServers"] = state ? state.ntpServers : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VpcDhcpOptionsArgs | undefined;
            inputs["domainName"] = args ? args.domainName : undefined;
            inputs["domainNameServers"] = args ? args.domainNameServers : undefined;
            inputs["netbiosNameServers"] = args ? args.netbiosNameServers : undefined;
            inputs["netbiosNodeType"] = args ? args.netbiosNodeType : undefined;
            inputs["ntpServers"] = args ? args.ntpServers : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["ownerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcDhcpOptions.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcDhcpOptions resources.
 */
export interface VpcDhcpOptionsState {
    /**
     * the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
     */
    readonly domainNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of NETBIOS name servers.
     */
    readonly netbiosNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     */
    readonly netbiosNodeType?: pulumi.Input<string>;
    /**
     * List of NTP servers to configure.
     */
    readonly ntpServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the AWS account that owns the DHCP options set.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a VpcDhcpOptions resource.
 */
export interface VpcDhcpOptionsArgs {
    /**
     * the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the `search` value in the `/etc/resolv.conf` file.
     */
    readonly domainName?: pulumi.Input<string>;
    /**
     * List of name servers to configure in `/etc/resolv.conf`. If you want to use the default AWS nameservers you should set this to `AmazonProvidedDNS`.
     */
    readonly domainNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of NETBIOS name servers.
     */
    readonly netbiosNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see [RFC 2132](http://www.ietf.org/rfc/rfc2132.txt).
     */
    readonly netbiosNodeType?: pulumi.Input<string>;
    /**
     * List of NTP servers to configure.
     */
    readonly ntpServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
