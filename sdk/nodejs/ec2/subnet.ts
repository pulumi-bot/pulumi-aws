// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an VPC subnet resource.
 *
 * > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.
 *
 * ## Example Usage
 *
 * ### Basic Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const main = new aws.ec2.Subnet("main", {
 *     cidrBlock: "10.0.1.0/24",
 *     tags: {
 *         Name: "Main",
 *     },
 *     vpcId: aws_vpc_main.id,
 * });
 * ```
 *
 * ### Subnets In Secondary VPC CIDR Blocks
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const secondaryCidr = new aws.ec2.VpcIpv4CidrBlockAssociation("secondary_cidr", {
 *     cidrBlock: "172.2.0.0/16",
 *     vpcId: aws_vpc_main.id,
 * });
 * const inSecondaryCidr = new aws.ec2.Subnet("in_secondary_cidr", {
 *     cidrBlock: "172.2.0.0/24",
 *     vpcId: secondaryCidr.vpcId,
 * });
 * ```
 */
export class Subnet extends pulumi.CustomResource {
    /**
     * Get an existing Subnet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubnetState, opts?: pulumi.CustomResourceOptions): Subnet {
        return new Subnet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/subnet:Subnet';

    /**
     * Returns true if the given object is an instance of Subnet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subnet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subnet.__pulumiType;
    }

    /**
     * The ARN of the subnet.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    public readonly assignIpv6AddressOnCreation!: pulumi.Output<boolean | undefined>;
    /**
     * The AZ for the subnet.
     */
    public readonly availabilityZone!: pulumi.Output<string>;
    /**
     * The AZ ID of the subnet.
     */
    public readonly availabilityZoneId!: pulumi.Output<string>;
    /**
     * The CIDR block for the subnet.
     */
    public readonly cidrBlock!: pulumi.Output<string>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    public readonly ipv6CidrBlock!: pulumi.Output<string>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    public /*out*/ readonly ipv6CidrBlockAssociationId!: pulumi.Output<string>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    public readonly mapPublicIpOnLaunch!: pulumi.Output<boolean | undefined>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    public readonly outpostArn!: pulumi.Output<string | undefined>;
    /**
     * The ID of the AWS account that owns the subnet.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The VPC ID.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a Subnet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubnetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubnetArgs | SubnetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SubnetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["assignIpv6AddressOnCreation"] = state ? state.assignIpv6AddressOnCreation : undefined;
            inputs["availabilityZone"] = state ? state.availabilityZone : undefined;
            inputs["availabilityZoneId"] = state ? state.availabilityZoneId : undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["ipv6CidrBlock"] = state ? state.ipv6CidrBlock : undefined;
            inputs["ipv6CidrBlockAssociationId"] = state ? state.ipv6CidrBlockAssociationId : undefined;
            inputs["mapPublicIpOnLaunch"] = state ? state.mapPublicIpOnLaunch : undefined;
            inputs["outpostArn"] = state ? state.outpostArn : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as SubnetArgs | undefined;
            if (!args || args.cidrBlock === undefined) {
                throw new Error("Missing required property 'cidrBlock'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["assignIpv6AddressOnCreation"] = args ? args.assignIpv6AddressOnCreation : undefined;
            inputs["availabilityZone"] = args ? args.availabilityZone : undefined;
            inputs["availabilityZoneId"] = args ? args.availabilityZoneId : undefined;
            inputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            inputs["ipv6CidrBlock"] = args ? args.ipv6CidrBlock : undefined;
            inputs["mapPublicIpOnLaunch"] = args ? args.mapPublicIpOnLaunch : undefined;
            inputs["outpostArn"] = args ? args.outpostArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["ipv6CidrBlockAssociationId"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Subnet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subnet resources.
 */
export interface SubnetState {
    /**
     * The ARN of the subnet.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    readonly assignIpv6AddressOnCreation?: pulumi.Input<boolean>;
    /**
     * The AZ for the subnet.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The AZ ID of the subnet.
     */
    readonly availabilityZoneId?: pulumi.Input<string>;
    /**
     * The CIDR block for the subnet.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    readonly ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * The association ID for the IPv6 CIDR block.
     */
    readonly ipv6CidrBlockAssociationId?: pulumi.Input<string>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    readonly mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    readonly outpostArn?: pulumi.Input<string>;
    /**
     * The ID of the AWS account that owns the subnet.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subnet resource.
 */
export interface SubnetArgs {
    /**
     * Specify true to indicate
     * that network interfaces created in the specified subnet should be
     * assigned an IPv6 address. Default is `false`
     */
    readonly assignIpv6AddressOnCreation?: pulumi.Input<boolean>;
    /**
     * The AZ for the subnet.
     */
    readonly availabilityZone?: pulumi.Input<string>;
    /**
     * The AZ ID of the subnet.
     */
    readonly availabilityZoneId?: pulumi.Input<string>;
    /**
     * The CIDR block for the subnet.
     */
    readonly cidrBlock: pulumi.Input<string>;
    /**
     * The IPv6 network range for the subnet,
     * in CIDR notation. The subnet size must use a /64 prefix length.
     */
    readonly ipv6CidrBlock?: pulumi.Input<string>;
    /**
     * Specify true to indicate
     * that instances launched into the subnet should be assigned
     * a public IP address. Default is `false`.
     */
    readonly mapPublicIpOnLaunch?: pulumi.Input<boolean>;
    /**
     * The Amazon Resource Name (ARN) of the Outpost.
     */
    readonly outpostArn?: pulumi.Input<string>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
    /**
     * The VPC ID.
     */
    readonly vpcId: pulumi.Input<string>;
}
