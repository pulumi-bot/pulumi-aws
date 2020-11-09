// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage an S3 Outposts Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.s3outposts.Endpoint("example", {
 *     outpostId: data.aws_outposts_outpost.example.id,
 *     securityGroupId: aws_security_group.example.id,
 *     subnetId: aws_subnet.example.id,
 * });
 * ```
 *
 * ## Import
 *
 * S3 Outposts Endpoints can be imported using Amazon Resource Name (ARN), EC2 Security Group identifier, and EC2 Subnet identifier, separated by commas (`,`) e.g.
 *
 * ```sh
 *  $ pulumi import aws:s3outposts/endpoint:Endpoint example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/endpoint/0123456789abcdef,sg-12345678,subnet-12345678
 * ```
 */
export class Endpoint extends pulumi.CustomResource {
    /**
     * Get an existing Endpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointState, opts?: pulumi.CustomResourceOptions): Endpoint {
        return new Endpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:s3outposts/endpoint:Endpoint';

    /**
     * Returns true if the given object is an instance of Endpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Endpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Endpoint.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN) of the endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * VPC CIDR block of the endpoint.
     */
    public /*out*/ readonly cidrBlock!: pulumi.Output<string>;
    /**
     * UTC creation time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * Set of nested attributes for associated Elastic Network Interfaces (ENIs).
     */
    public /*out*/ readonly networkInterfaces!: pulumi.Output<outputs.s3outposts.EndpointNetworkInterface[]>;
    /**
     * Identifier of the Outpost to contain this endpoint.
     */
    public readonly outpostId!: pulumi.Output<string>;
    /**
     * Identifier of the EC2 Security Group.
     */
    public readonly securityGroupId!: pulumi.Output<string>;
    /**
     * Identifier of the EC2 Subnet.
     */
    public readonly subnetId!: pulumi.Output<string>;

    /**
     * Create a Endpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointArgs | EndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as EndpointState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            inputs["creationTime"] = state ? state.creationTime : undefined;
            inputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            inputs["outpostId"] = state ? state.outpostId : undefined;
            inputs["securityGroupId"] = state ? state.securityGroupId : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as EndpointArgs | undefined;
            if (!args || args.outpostId === undefined) {
                throw new Error("Missing required property 'outpostId'");
            }
            if (!args || args.securityGroupId === undefined) {
                throw new Error("Missing required property 'securityGroupId'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["outpostId"] = args ? args.outpostId : undefined;
            inputs["securityGroupId"] = args ? args.securityGroupId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["cidrBlock"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Endpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Endpoint resources.
 */
export interface EndpointState {
    /**
     * Amazon Resource Name (ARN) of the endpoint.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * VPC CIDR block of the endpoint.
     */
    readonly cidrBlock?: pulumi.Input<string>;
    /**
     * UTC creation time in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8).
     */
    readonly creationTime?: pulumi.Input<string>;
    /**
     * Set of nested attributes for associated Elastic Network Interfaces (ENIs).
     */
    readonly networkInterfaces?: pulumi.Input<pulumi.Input<inputs.s3outposts.EndpointNetworkInterface>[]>;
    /**
     * Identifier of the Outpost to contain this endpoint.
     */
    readonly outpostId?: pulumi.Input<string>;
    /**
     * Identifier of the EC2 Security Group.
     */
    readonly securityGroupId?: pulumi.Input<string>;
    /**
     * Identifier of the EC2 Subnet.
     */
    readonly subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Endpoint resource.
 */
export interface EndpointArgs {
    /**
     * Identifier of the Outpost to contain this endpoint.
     */
    readonly outpostId: pulumi.Input<string>;
    /**
     * Identifier of the EC2 Security Group.
     */
    readonly securityGroupId: pulumi.Input<string>;
    /**
     * Identifier of the EC2 Subnet.
     */
    readonly subnetId: pulumi.Input<string>;
}
