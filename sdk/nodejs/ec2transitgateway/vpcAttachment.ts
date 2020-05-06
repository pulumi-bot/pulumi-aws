// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway VPC Attachment. For examples of custom route table association and propagation, see the EC2 Transit Gateway Networking Examples Guide.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_vpc_attachment.html.markdown.
 */
export class VpcAttachment extends pulumi.CustomResource {
    /**
     * Get an existing VpcAttachment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcAttachmentState, opts?: pulumi.CustomResourceOptions): VpcAttachment {
        return new VpcAttachment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/vpcAttachment:VpcAttachment';

    /**
     * Returns true if the given object is an instance of VpcAttachment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcAttachment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcAttachment.__pulumiType;
    }

    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     */
    public readonly dnsSupport!: pulumi.Output<string | undefined>;
    /**
     * Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
     */
    public readonly ipv6Support!: pulumi.Output<string | undefined>;
    /**
     * Identifiers of EC2 Subnets.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    public readonly transitGatewayDefaultRouteTableAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    public readonly transitGatewayDefaultRouteTablePropagation!: pulumi.Output<boolean | undefined>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    public readonly transitGatewayId!: pulumi.Output<string>;
    /**
     * Identifier of EC2 VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;
    /**
     * Identifier of the AWS account that owns the EC2 VPC.
     */
    public /*out*/ readonly vpcOwnerId!: pulumi.Output<string>;

    /**
     * Create a VpcAttachment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcAttachmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcAttachmentArgs | VpcAttachmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcAttachmentState | undefined;
            inputs["dnsSupport"] = state ? state.dnsSupport : undefined;
            inputs["ipv6Support"] = state ? state.ipv6Support : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitGatewayDefaultRouteTableAssociation"] = state ? state.transitGatewayDefaultRouteTableAssociation : undefined;
            inputs["transitGatewayDefaultRouteTablePropagation"] = state ? state.transitGatewayDefaultRouteTablePropagation : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vpcOwnerId"] = state ? state.vpcOwnerId : undefined;
        } else {
            const args = argsOrState as VpcAttachmentArgs | undefined;
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            if (!args || args.transitGatewayId === undefined) {
                throw new Error("Missing required property 'transitGatewayId'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["dnsSupport"] = args ? args.dnsSupport : undefined;
            inputs["ipv6Support"] = args ? args.ipv6Support : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayDefaultRouteTableAssociation"] = args ? args.transitGatewayDefaultRouteTableAssociation : undefined;
            inputs["transitGatewayDefaultRouteTablePropagation"] = args ? args.transitGatewayDefaultRouteTablePropagation : undefined;
            inputs["transitGatewayId"] = args ? args.transitGatewayId : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["vpcOwnerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcAttachment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcAttachment resources.
 */
export interface VpcAttachmentState {
    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     */
    readonly dnsSupport?: pulumi.Input<string>;
    /**
     * Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
     */
    readonly ipv6Support?: pulumi.Input<string>;
    /**
     * Identifiers of EC2 Subnets.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTableAssociation?: pulumi.Input<boolean>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTablePropagation?: pulumi.Input<boolean>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    readonly transitGatewayId?: pulumi.Input<string>;
    /**
     * Identifier of EC2 VPC.
     */
    readonly vpcId?: pulumi.Input<string>;
    /**
     * Identifier of the AWS account that owns the EC2 VPC.
     */
    readonly vpcOwnerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcAttachment resource.
 */
export interface VpcAttachmentArgs {
    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
     */
    readonly dnsSupport?: pulumi.Input<string>;
    /**
     * Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
     */
    readonly ipv6Support?: pulumi.Input<string>;
    /**
     * Identifiers of EC2 Subnets.
     */
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTableAssociation?: pulumi.Input<boolean>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTablePropagation?: pulumi.Input<boolean>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    readonly transitGatewayId: pulumi.Input<string>;
    /**
     * Identifier of EC2 VPC.
     */
    readonly vpcId: pulumi.Input<string>;
}
