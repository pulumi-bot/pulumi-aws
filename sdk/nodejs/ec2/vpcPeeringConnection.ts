// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage a VPC peering connection.
 *
 * > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
 * both a standalone VPC Peering Connection Options and a VPC Peering Connection
 * resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
 * connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
 * Doing so will cause a conflict of options and will overwrite the options.
 * Using a VPC Peering Connection Options resource decouples management of the connection options from
 * management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
 *
 * > **Note:** For cross-account (requester's AWS account differs from the accepter's AWS account) or inter-region
 * VPC Peering Connections use the `aws.ec2.VpcPeeringConnection` resource to manage the requester's side of the
 * connection and use the `aws.ec2.VpcPeeringConnectionAccepter` resource to manage the accepter's side of the connection.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ec2.VpcPeeringConnection("foo", {
 *     peerOwnerId: _var.peer_owner_id,
 *     peerVpcId: aws_vpc.bar.id,
 *     vpcId: aws_vpc.foo.id,
 * });
 * ```
 *
 * Basic usage with connection options:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const foo = new aws.ec2.VpcPeeringConnection("foo", {
 *     peerOwnerId: _var.peer_owner_id,
 *     peerVpcId: aws_vpc.bar.id,
 *     vpcId: aws_vpc.foo.id,
 *     accepter: {
 *         allowRemoteVpcDnsResolution: true,
 *     },
 *     requester: {
 *         allowRemoteVpcDnsResolution: true,
 *     },
 * });
 * ```
 *
 * Basic usage with tags:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"});
 * const bar = new aws.ec2.Vpc("bar", {cidrBlock: "10.2.0.0/16"});
 * const fooVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection", {
 *     peerOwnerId: _var.peer_owner_id,
 *     peerVpcId: bar.id,
 *     vpcId: fooVpc.id,
 *     autoAccept: true,
 *     tags: {
 *         Name: "VPC Peering between foo and bar",
 *     },
 * });
 * ```
 *
 * Basic usage with region:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const fooVpc = new aws.ec2.Vpc("fooVpc", {cidrBlock: "10.1.0.0/16"}, {
 *     provider: aws["us-west-2"],
 * });
 * const bar = new aws.ec2.Vpc("bar", {cidrBlock: "10.2.0.0/16"}, {
 *     provider: aws["us-east-1"],
 * });
 * const fooVpcPeeringConnection = new aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection", {
 *     peerOwnerId: _var.peer_owner_id,
 *     peerVpcId: bar.id,
 *     vpcId: fooVpc.id,
 *     peerRegion: "us-east-1",
 * });
 * ```
 * ## Notes
 *
 * If both VPCs are not in the same AWS account do not enable the `autoAccept` attribute.
 * The accepter can manage its side of the connection using the `aws.ec2.VpcPeeringConnectionAccepter` resource
 * or accept the connection manually using the AWS Management Console, AWS CLI, through SDKs, etc.
 *
 * ## Import
 *
 * VPC Peering resources can be imported using the `vpc peering id`, e.g.
 *
 * ```sh
 *  $ pulumi import aws:ec2/vpcPeeringConnection:VpcPeeringConnection test_connection pcx-111aaa111
 * ```
 *
 *  [1]/docs/providers/aws/index.html
 */
export class VpcPeeringConnection extends pulumi.CustomResource {
    /**
     * Get an existing VpcPeeringConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcPeeringConnectionState, opts?: pulumi.CustomResourceOptions): VpcPeeringConnection {
        return new VpcPeeringConnection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpcPeeringConnection:VpcPeeringConnection';

    /**
     * Returns true if the given object is an instance of VpcPeeringConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcPeeringConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcPeeringConnection.__pulumiType;
    }

    /**
     * The status of the VPC Peering Connection request.
     */
    public /*out*/ readonly acceptStatus!: pulumi.Output<string>;
    /**
     * An optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
     * the peering connection (a maximum of one).
     */
    public readonly accepter!: pulumi.Output<outputs.ec2.VpcPeeringConnectionAccepter>;
    /**
     * Accept the peering (both VPCs need to be in the same AWS account).
     */
    public readonly autoAccept!: pulumi.Output<boolean | undefined>;
    /**
     * The AWS account ID of the owner of the peer VPC.
     * Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    public readonly peerOwnerId!: pulumi.Output<string>;
    /**
     * The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
     * and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
     */
    public readonly peerRegion!: pulumi.Output<string>;
    /**
     * The ID of the VPC with which you are creating the VPC Peering Connection.
     */
    public readonly peerVpcId!: pulumi.Output<string>;
    /**
     * A optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
     * the peering connection (a maximum of one).
     */
    public readonly requester!: pulumi.Output<outputs.ec2.VpcPeeringConnectionRequester>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the requester VPC.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcPeeringConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcPeeringConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcPeeringConnectionArgs | VpcPeeringConnectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcPeeringConnectionState | undefined;
            inputs["acceptStatus"] = state ? state.acceptStatus : undefined;
            inputs["accepter"] = state ? state.accepter : undefined;
            inputs["autoAccept"] = state ? state.autoAccept : undefined;
            inputs["peerOwnerId"] = state ? state.peerOwnerId : undefined;
            inputs["peerRegion"] = state ? state.peerRegion : undefined;
            inputs["peerVpcId"] = state ? state.peerVpcId : undefined;
            inputs["requester"] = state ? state.requester : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcPeeringConnectionArgs | undefined;
            if ((!args || args.peerVpcId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'peerVpcId'");
            }
            if ((!args || args.vpcId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["accepter"] = args ? args.accepter : undefined;
            inputs["autoAccept"] = args ? args.autoAccept : undefined;
            inputs["peerOwnerId"] = args ? args.peerOwnerId : undefined;
            inputs["peerRegion"] = args ? args.peerRegion : undefined;
            inputs["peerVpcId"] = args ? args.peerVpcId : undefined;
            inputs["requester"] = args ? args.requester : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["acceptStatus"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcPeeringConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcPeeringConnection resources.
 */
export interface VpcPeeringConnectionState {
    /**
     * The status of the VPC Peering Connection request.
     */
    readonly acceptStatus?: pulumi.Input<string>;
    /**
     * An optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
     * the peering connection (a maximum of one).
     */
    readonly accepter?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepter>;
    /**
     * Accept the peering (both VPCs need to be in the same AWS account).
     */
    readonly autoAccept?: pulumi.Input<boolean>;
    /**
     * The AWS account ID of the owner of the peer VPC.
     * Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    readonly peerOwnerId?: pulumi.Input<string>;
    /**
     * The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
     * and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
     */
    readonly peerRegion?: pulumi.Input<string>;
    /**
     * The ID of the VPC with which you are creating the VPC Peering Connection.
     */
    readonly peerVpcId?: pulumi.Input<string>;
    /**
     * A optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
     * the peering connection (a maximum of one).
     */
    readonly requester?: pulumi.Input<inputs.ec2.VpcPeeringConnectionRequester>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the requester VPC.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcPeeringConnection resource.
 */
export interface VpcPeeringConnectionArgs {
    /**
     * An optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
     * the peering connection (a maximum of one).
     */
    readonly accepter?: pulumi.Input<inputs.ec2.VpcPeeringConnectionAccepter>;
    /**
     * Accept the peering (both VPCs need to be in the same AWS account).
     */
    readonly autoAccept?: pulumi.Input<boolean>;
    /**
     * The AWS account ID of the owner of the peer VPC.
     * Defaults to the account ID the [AWS provider](https://www.terraform.io/docs/providers/aws/index.html) is currently connected to.
     */
    readonly peerOwnerId?: pulumi.Input<string>;
    /**
     * The region of the accepter VPC of the [VPC Peering Connection]. `autoAccept` must be `false`,
     * and use the `aws.ec2.VpcPeeringConnectionAccepter` to manage the accepter side.
     */
    readonly peerRegion?: pulumi.Input<string>;
    /**
     * The ID of the VPC with which you are creating the VPC Peering Connection.
     */
    readonly peerVpcId: pulumi.Input<string>;
    /**
     * A optional configuration block that allows for [VPC Peering Connection]
     * (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
     * the peering connection (a maximum of one).
     */
    readonly requester?: pulumi.Input<inputs.ec2.VpcPeeringConnectionRequester>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the requester VPC.
     */
    readonly vpcId: pulumi.Input<string>;
}
