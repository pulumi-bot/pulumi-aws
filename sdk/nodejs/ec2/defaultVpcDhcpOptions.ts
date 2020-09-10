// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DefaultVpcDhcpOptions extends pulumi.CustomResource {
    /**
     * Get an existing DefaultVpcDhcpOptions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultVpcDhcpOptionsState, opts?: pulumi.CustomResourceOptions): DefaultVpcDhcpOptions {
        return new DefaultVpcDhcpOptions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions';

    /**
     * Returns true if the given object is an instance of DefaultVpcDhcpOptions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultVpcDhcpOptions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultVpcDhcpOptions.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly domainName!: pulumi.Output<string>;
    public /*out*/ readonly domainNameServers!: pulumi.Output<string>;
    public readonly netbiosNameServers!: pulumi.Output<string[] | undefined>;
    public readonly netbiosNodeType!: pulumi.Output<string | undefined>;
    public /*out*/ readonly ntpServers!: pulumi.Output<string>;
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a DefaultVpcDhcpOptions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DefaultVpcDhcpOptionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultVpcDhcpOptionsArgs | DefaultVpcDhcpOptionsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DefaultVpcDhcpOptionsState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["domainName"] = state ? state.domainName : undefined;
            inputs["domainNameServers"] = state ? state.domainNameServers : undefined;
            inputs["netbiosNameServers"] = state ? state.netbiosNameServers : undefined;
            inputs["netbiosNodeType"] = state ? state.netbiosNodeType : undefined;
            inputs["ntpServers"] = state ? state.ntpServers : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DefaultVpcDhcpOptionsArgs | undefined;
            inputs["netbiosNameServers"] = args ? args.netbiosNameServers : undefined;
            inputs["netbiosNodeType"] = args ? args.netbiosNodeType : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["domainName"] = undefined /*out*/;
            inputs["domainNameServers"] = undefined /*out*/;
            inputs["ntpServers"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(DefaultVpcDhcpOptions.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultVpcDhcpOptions resources.
 */
export interface DefaultVpcDhcpOptionsState {
    readonly arn?: pulumi.Input<string>;
    readonly domainName?: pulumi.Input<string>;
    readonly domainNameServers?: pulumi.Input<string>;
    readonly netbiosNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly netbiosNodeType?: pulumi.Input<string>;
    readonly ntpServers?: pulumi.Input<string>;
    readonly ownerId?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a DefaultVpcDhcpOptions resource.
 */
export interface DefaultVpcDhcpOptionsArgs {
    readonly netbiosNameServers?: pulumi.Input<pulumi.Input<string>[]>;
    readonly netbiosNodeType?: pulumi.Input<string>;
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
