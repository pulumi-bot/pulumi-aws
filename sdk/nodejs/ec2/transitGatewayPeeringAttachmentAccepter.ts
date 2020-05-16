// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages the accepter's side of an EC2 Transit Gateway Peering Attachment.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2.TransitGatewayPeeringAttachmentAccepter("example", {
 *     tags: {
 *         Name: "Example cross-account attachment",
 *     },
 *     transitGatewayAttachmentId: aws_ec2_transit_gateway_peering_attachment_example.id,
 * });
 * ```
 */
export class TransitGatewayPeeringAttachmentAccepter extends pulumi.CustomResource {
    /**
     * Get an existing TransitGatewayPeeringAttachmentAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TransitGatewayPeeringAttachmentAccepterState, opts?: pulumi.CustomResourceOptions): TransitGatewayPeeringAttachmentAccepter {
        return new TransitGatewayPeeringAttachmentAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/transitGatewayPeeringAttachmentAccepter:TransitGatewayPeeringAttachmentAccepter';

    /**
     * Returns true if the given object is an instance of TransitGatewayPeeringAttachmentAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGatewayPeeringAttachmentAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGatewayPeeringAttachmentAccepter.__pulumiType;
    }

    /**
     * Identifier of the AWS account that owns the EC2 TGW peering.
     */
    public /*out*/ readonly peerAccountId!: pulumi.Output<string>;
    public /*out*/ readonly peerRegion!: pulumi.Output<string>;
    /**
     * Identifier of EC2 Transit Gateway to peer with.
     */
    public /*out*/ readonly peerTransitGatewayId!: pulumi.Output<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Peering Attachment.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The ID of the EC2 Transit Gateway Peering Attachment to manage.
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    public /*out*/ readonly transitGatewayId!: pulumi.Output<string>;

    /**
     * Create a TransitGatewayPeeringAttachmentAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TransitGatewayPeeringAttachmentAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TransitGatewayPeeringAttachmentAccepterArgs | TransitGatewayPeeringAttachmentAccepterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as TransitGatewayPeeringAttachmentAccepterState | undefined;
            inputs["peerAccountId"] = state ? state.peerAccountId : undefined;
            inputs["peerRegion"] = state ? state.peerRegion : undefined;
            inputs["peerTransitGatewayId"] = state ? state.peerTransitGatewayId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
        } else {
            const args = argsOrState as TransitGatewayPeeringAttachmentAccepterArgs | undefined;
            if (!args || args.transitGatewayAttachmentId === undefined) {
                throw new Error("Missing required property 'transitGatewayAttachmentId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            inputs["peerAccountId"] = undefined /*out*/;
            inputs["peerRegion"] = undefined /*out*/;
            inputs["peerTransitGatewayId"] = undefined /*out*/;
            inputs["transitGatewayId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(TransitGatewayPeeringAttachmentAccepter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TransitGatewayPeeringAttachmentAccepter resources.
 */
export interface TransitGatewayPeeringAttachmentAccepterState {
    /**
     * Identifier of the AWS account that owns the EC2 TGW peering.
     */
    readonly peerAccountId?: pulumi.Input<string>;
    readonly peerRegion?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway to peer with.
     */
    readonly peerTransitGatewayId?: pulumi.Input<string>;
    /**
     * Key-value tags for the EC2 Transit Gateway Peering Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the EC2 Transit Gateway Peering Attachment to manage.
     */
    readonly transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    readonly transitGatewayId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TransitGatewayPeeringAttachmentAccepter resource.
 */
export interface TransitGatewayPeeringAttachmentAccepterArgs {
    /**
     * Key-value tags for the EC2 Transit Gateway Peering Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The ID of the EC2 Transit Gateway Peering Attachment to manage.
     */
    readonly transitGatewayAttachmentId: pulumi.Input<string>;
}
