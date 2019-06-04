// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a Virtual Private Gateway attachment resource, allowing for an existing
 * hardware VPN gateway to be attached and/or detached from a VPC.
 * 
 * > **Note:** The `aws_vpn_gateway`
 * resource can also automatically attach the Virtual Private Gateway it creates
 * to an existing VPC by setting the `vpc_id` attribute accordingly.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const network = new aws.ec2.Vpc("network", {
 *     cidrBlock: "10.0.0.0/16",
 * });
 * const vpn = new aws.ec2.VpnGateway("vpn", {
 *     tags: {
 *         Name: "example-vpn-gateway",
 *     },
 * });
 * const vpnAttachment = new aws.ec2.VpnGatewayAttachment("vpn_attachment", {
 *     vpcId: network.id,
 *     vpnGatewayId: vpn.id,
 * });
 * ```
 * 
 * See [Virtual Private Cloud](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Introduction.html)
 * and [Virtual Private Gateway](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_VPN.html) user
 * guides for more information.
 */
export class VpnGatewayAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VpnGatewayAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpnGatewayAttachmentState, opts?: pulumi.CustomResourceOptions): VpnGatewayAttachment {
        return new VpnGatewayAttachment(name, <any>state, { ...opts, id: id });
    }

    private static readonly __pulumiType = 'aws:ec2/vpnGatewayAttachment:VpnGatewayAttachment';

    /**
     * Returns true if the given object is an instance of VpnGatewayAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnGatewayAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }

        const t = obj['__pulumiType'];
        if (typeof t !== 'string') {
            return false;
        }

        return t === VpnGatewayAttachment.__pulumiType;
    }

    /**
     * The ID of the VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    public readonly vpnGatewayId!: pulumi.Output<string>;

    /**
     * Create a VpnGatewayAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnGatewayAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnGatewayAttachmentArgs | VpnGatewayAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpnGatewayAttachmentState | undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vpnGatewayId"] = state ? state.vpnGatewayId : undefined;
        } else {
            const args = argsOrState as VpnGatewayAttachmentArgs | undefined;
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            if (!args || args.vpnGatewayId === undefined) {
                throw new Error("Missing required property 'vpnGatewayId'");
            }
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["vpnGatewayId"] = args ? args.vpnGatewayId : undefined;
        }
        super(VpnGatewayAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpnGatewayAttachment resources.
 */
export interface VpnGatewayAttachmentState {
    /**
     * The ID of the VPC.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    readonly vpnGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpnGatewayAttachment resource.
 */
export interface VpnGatewayAttachmentArgs {
    /**
     * The ID of the VPC.
     */
    readonly vpcId: pulumi.Input<string>;
    /**
     * The ID of the Virtual Private Gateway.
     */
    readonly vpnGatewayId: pulumi.Input<string>;
}
