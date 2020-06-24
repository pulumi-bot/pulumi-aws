// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages the accepter's side of an EC2 Transit Gateway VPC Attachment.
 *
 * When a cross-account (requester's AWS account differs from the accepter's AWS account) EC2 Transit Gateway VPC Attachment
 * is created, an EC2 Transit Gateway VPC Attachment resource is automatically created in the accepter's account.
 * The requester can use the `aws.ec2transitgateway.VpcAttachment` resource to manage its side of the connection
 * and the accepter can use the `aws.ec2transitgateway.VpcAttachmentAccepter` resource to "adopt" its side of the
 * connection into management.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.ec2transitgateway.VpcAttachmentAccepter("example", {
 *     tags: {
 *         Name: "Example cross-account attachment",
 *     },
 *     transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment_example.id,
 * });
 * ```
 *
 * A full example of how to how to create a Transit Gateway in one AWS account, share it with a second AWS account, and attach a VPC in the second account to the Transit Gateway via the `aws.ec2transitgateway.VpcAttachment` and `aws.ec2transitgateway.VpcAttachmentAccepter` resources can be found in [the `./examples/transit-gateway-cross-account-vpc-attachment` directory within the Github Repository](https://github.com/providers/provider-aws/tree/master/examples/transit-gateway-cross-account-vpc-attachment).
 */
export class VpcAttachmentAccepter extends pulumi.CustomResource {
    /**
     * Get an existing VpcAttachmentAccepter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcAttachmentAccepterState, opts?: pulumi.CustomResourceOptions): VpcAttachmentAccepter {
        return new VpcAttachmentAccepter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/vpcAttachmentAccepter:VpcAttachmentAccepter';

    /**
     * Returns true if the given object is an instance of VpcAttachmentAccepter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcAttachmentAccepter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcAttachmentAccepter.__pulumiType;
    }

    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`.
     */
    public /*out*/ readonly dnsSupport!: pulumi.Output<string>;
    /**
     * Whether IPv6 support is enabled. Valid values: `disable`, `enable`.
     */
    public /*out*/ readonly ipv6Support!: pulumi.Output<string>;
    /**
     * Identifiers of EC2 Subnets.
     */
    public /*out*/ readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the EC2 Transit Gateway Attachment to manage.
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
     */
    public readonly transitGatewayDefaultRouteTableAssociation!: pulumi.Output<boolean | undefined>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
     */
    public readonly transitGatewayDefaultRouteTablePropagation!: pulumi.Output<boolean | undefined>;
    /**
     * Identifier of EC2 Transit Gateway.
     */
    public /*out*/ readonly transitGatewayId!: pulumi.Output<string>;
    /**
     * Identifier of EC2 VPC.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * Identifier of the AWS account that owns the EC2 VPC.
     */
    public /*out*/ readonly vpcOwnerId!: pulumi.Output<string>;

    /**
     * Create a VpcAttachmentAccepter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcAttachmentAccepterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcAttachmentAccepterArgs | VpcAttachmentAccepterState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcAttachmentAccepterState | undefined;
            inputs["dnsSupport"] = state ? state.dnsSupport : undefined;
            inputs["ipv6Support"] = state ? state.ipv6Support : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayDefaultRouteTableAssociation"] = state ? state.transitGatewayDefaultRouteTableAssociation : undefined;
            inputs["transitGatewayDefaultRouteTablePropagation"] = state ? state.transitGatewayDefaultRouteTablePropagation : undefined;
            inputs["transitGatewayId"] = state ? state.transitGatewayId : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
            inputs["vpcOwnerId"] = state ? state.vpcOwnerId : undefined;
        } else {
            const args = argsOrState as VpcAttachmentAccepterArgs | undefined;
            if (!args || args.transitGatewayAttachmentId === undefined) {
                throw new Error("Missing required property 'transitGatewayAttachmentId'");
            }
            inputs["tags"] = args ? args.tags : undefined;
            inputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayDefaultRouteTableAssociation"] = args ? args.transitGatewayDefaultRouteTableAssociation : undefined;
            inputs["transitGatewayDefaultRouteTablePropagation"] = args ? args.transitGatewayDefaultRouteTablePropagation : undefined;
            inputs["dnsSupport"] = undefined /*out*/;
            inputs["ipv6Support"] = undefined /*out*/;
            inputs["subnetIds"] = undefined /*out*/;
            inputs["transitGatewayId"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
            inputs["vpcOwnerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcAttachmentAccepter.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcAttachmentAccepter resources.
 */
export interface VpcAttachmentAccepterState {
    /**
     * Whether DNS support is enabled. Valid values: `disable`, `enable`.
     */
    readonly dnsSupport?: pulumi.Input<string>;
    /**
     * Whether IPv6 support is enabled. Valid values: `disable`, `enable`.
     */
    readonly ipv6Support?: pulumi.Input<string>;
    /**
     * Identifiers of EC2 Subnets.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the EC2 Transit Gateway Attachment to manage.
     */
    readonly transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTableAssociation?: pulumi.Input<boolean>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
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
 * The set of arguments for constructing a VpcAttachmentAccepter resource.
 */
export interface VpcAttachmentAccepterArgs {
    /**
     * Key-value tags for the EC2 Transit Gateway VPC Attachment.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the EC2 Transit Gateway Attachment to manage.
     */
    readonly transitGatewayAttachmentId: pulumi.Input<string>;
    /**
     * Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTableAssociation?: pulumi.Input<boolean>;
    /**
     * Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. Default value: `true`.
     */
    readonly transitGatewayDefaultRouteTablePropagation?: pulumi.Input<boolean>;
}
