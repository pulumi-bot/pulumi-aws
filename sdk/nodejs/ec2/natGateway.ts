// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a VPC NAT Gateway.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const gw = new aws.ec2.NatGateway("gw", {
 *     allocationId: aws_eip_nat.id,
 *     subnetId: aws_subnet_public.id,
 * });
 * ```
 * 
 * Usage with tags:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const gw = new aws.ec2.NatGateway("gw", {
 *     allocationId: aws_eip_nat.id,
 *     subnetId: aws_subnet_public.id,
 *     tags: {
 *         Name: "gw NAT",
 *     },
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/nat_gateway.html.markdown.
 */
export class NatGateway extends pulumi.CustomResource {
    /**
     * Get an existing NatGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NatGatewayState, opts?: pulumi.CustomResourceOptions): NatGateway {
        return new NatGateway(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/natGateway:NatGateway';

    /**
     * Returns true if the given object is an instance of NatGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NatGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NatGateway.__pulumiType;
    }

    /**
     * The Allocation ID of the Elastic IP address for the gateway.
     */
    public readonly allocationId!: pulumi.Output<string>;
    /**
     * The ENI ID of the network interface created by the NAT gateway.
     */
    public /*out*/ readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The private IP address of the NAT Gateway.
     */
    public /*out*/ readonly privateIp!: pulumi.Output<string>;
    /**
     * The public IP address of the NAT Gateway.
     */
    public /*out*/ readonly publicIp!: pulumi.Output<string>;
    /**
     * The Subnet ID of the subnet in which to place the gateway.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a NatGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NatGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NatGatewayArgs | NatGatewayState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NatGatewayState | undefined;
            inputs["allocationId"] = state ? state.allocationId : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            inputs["privateIp"] = state ? state.privateIp : undefined;
            inputs["publicIp"] = state ? state.publicIp : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as NatGatewayArgs | undefined;
            if (!args || args.allocationId === undefined) {
                throw new Error("Missing required property 'allocationId'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["allocationId"] = args ? args.allocationId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["networkInterfaceId"] = undefined /*out*/;
            inputs["privateIp"] = undefined /*out*/;
            inputs["publicIp"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NatGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NatGateway resources.
 */
export interface NatGatewayState {
    /**
     * The Allocation ID of the Elastic IP address for the gateway.
     */
    readonly allocationId?: pulumi.Input<string>;
    /**
     * The ENI ID of the network interface created by the NAT gateway.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
    /**
     * The private IP address of the NAT Gateway.
     */
    readonly privateIp?: pulumi.Input<string>;
    /**
     * The public IP address of the NAT Gateway.
     */
    readonly publicIp?: pulumi.Input<string>;
    /**
     * The Subnet ID of the subnet in which to place the gateway.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a NatGateway resource.
 */
export interface NatGatewayArgs {
    /**
     * The Allocation ID of the Elastic IP address for the gateway.
     */
    readonly allocationId: pulumi.Input<string>;
    /**
     * The Subnet ID of the subnet in which to place the gateway.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
