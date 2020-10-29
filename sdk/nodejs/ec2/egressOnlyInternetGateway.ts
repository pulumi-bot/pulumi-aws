// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * [IPv6 only] Creates an egress-only Internet gateway for your VPC.
 * An egress-only Internet gateway is used to enable outbound communication
 * over IPv6 from instances in your VPC to the Internet, and prevents hosts
 * outside of your VPC from initiating an IPv6 connection with your instance.
 */
export class EgressOnlyInternetGateway extends pulumi.CustomResource {
    /**
     * Get an existing EgressOnlyInternetGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EgressOnlyInternetGatewayState, opts?: pulumi.CustomResourceOptions): EgressOnlyInternetGateway {
        return new EgressOnlyInternetGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway';

    /**
     * Returns true if the given object is an instance of EgressOnlyInternetGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EgressOnlyInternetGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EgressOnlyInternetGateway.__pulumiType;
    }

    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The VPC ID to create in.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a EgressOnlyInternetGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EgressOnlyInternetGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EgressOnlyInternetGatewayArgs | EgressOnlyInternetGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EgressOnlyInternetGatewayState | undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as EgressOnlyInternetGatewayArgs | undefined;
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(EgressOnlyInternetGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EgressOnlyInternetGateway resources.
 */
export interface EgressOnlyInternetGatewayState {
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC ID to create in.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EgressOnlyInternetGateway resource.
 */
export interface EgressOnlyInternetGatewayArgs {
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC ID to create in.
     */
    readonly vpcId: pulumi.Input<string>;
}
