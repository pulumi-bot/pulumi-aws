// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides a VPC Endpoint resource.
 *
 * > **NOTE on VPC Endpoints and VPC Endpoint Associations:** This provider provides both standalone VPC Endpoint Associations for
 * Route Tables - (an association between a VPC endpoint and a single `routeTableId`) and
 * Subnets - (an association between a VPC endpoint and a single `subnetId`) and
 * a VPC Endpoint resource with `routeTableIds` and `subnetIds` attributes.
 * Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
 * Doing so will cause a conflict of associations and will overwrite the association.
 *
 * ## Example Usage
 * ### Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const s3 = new aws.ec2.VpcEndpoint("s3", {
 *     vpcId: aws_vpc.main.id,
 *     serviceName: "com.amazonaws.us-west-2.s3",
 * });
 * ```
 * ### Basic w/ Tags
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const s3 = new aws.ec2.VpcEndpoint("s3", {
 *     vpcId: aws_vpc.main.id,
 *     serviceName: "com.amazonaws.us-west-2.s3",
 *     tags: {
 *         Environment: "test",
 *     },
 * });
 * ```
 * ### Interface Endpoint Type
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const ec2 = new aws.ec2.VpcEndpoint("ec2", {
 *     vpcId: aws_vpc.main.id,
 *     serviceName: "com.amazonaws.us-west-2.ec2",
 *     vpcEndpointType: "Interface",
 *     securityGroupIds: [aws_security_group.sg1.id],
 *     privateDnsEnabled: true,
 * });
 * ```
 */
export class VpcEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointState, opts?: pulumi.CustomResourceOptions): VpcEndpoint {
        return new VpcEndpoint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpcEndpoint:VpcEndpoint';

    /**
     * Returns true if the given object is an instance of VpcEndpoint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpoint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpoint.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the VPC endpoint.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
     */
    public readonly autoAccept!: pulumi.Output<boolean | undefined>;
    /**
     * The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    public /*out*/ readonly cidrBlocks!: pulumi.Output<string[]>;
    /**
     * The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
     */
    public /*out*/ readonly dnsEntries!: pulumi.Output<outputs.ec2.VpcEndpointDnsEntry[]>;
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * The ID of the AWS account that owns the VPC endpoint.
     */
    public /*out*/ readonly ownerId!: pulumi.Output<string>;
    /**
     * A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
     */
    public readonly policy!: pulumi.Output<string>;
    /**
     * The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    public /*out*/ readonly prefixListId!: pulumi.Output<string>;
    /**
     * Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
     * Defaults to `false`.
     */
    public readonly privateDnsEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
     */
    public /*out*/ readonly requesterManaged!: pulumi.Output<boolean>;
    /**
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     */
    public readonly routeTableIds!: pulumi.Output<string[]>;
    /**
     * The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * The state of the VPC endpoint.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
     */
    public readonly vpcEndpointType!: pulumi.Output<string | undefined>;
    /**
     * The ID of the VPC in which the endpoint will be used.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpoint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointArgs | VpcEndpointState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcEndpointState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["autoAccept"] = state ? state.autoAccept : undefined;
            inputs["cidrBlocks"] = state ? state.cidrBlocks : undefined;
            inputs["dnsEntries"] = state ? state.dnsEntries : undefined;
            inputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            inputs["ownerId"] = state ? state.ownerId : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["prefixListId"] = state ? state.prefixListId : undefined;
            inputs["privateDnsEnabled"] = state ? state.privateDnsEnabled : undefined;
            inputs["requesterManaged"] = state ? state.requesterManaged : undefined;
            inputs["routeTableIds"] = state ? state.routeTableIds : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["vpcEndpointType"] = state ? state.vpcEndpointType : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as VpcEndpointArgs | undefined;
            if (!args || args.serviceName === undefined) {
                throw new Error("Missing required property 'serviceName'");
            }
            if (!args || args.vpcId === undefined) {
                throw new Error("Missing required property 'vpcId'");
            }
            inputs["autoAccept"] = args ? args.autoAccept : undefined;
            inputs["policy"] = args ? args.policy : undefined;
            inputs["privateDnsEnabled"] = args ? args.privateDnsEnabled : undefined;
            inputs["routeTableIds"] = args ? args.routeTableIds : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["serviceName"] = args ? args.serviceName : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["vpcEndpointType"] = args ? args.vpcEndpointType : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["cidrBlocks"] = undefined /*out*/;
            inputs["dnsEntries"] = undefined /*out*/;
            inputs["networkInterfaceIds"] = undefined /*out*/;
            inputs["ownerId"] = undefined /*out*/;
            inputs["prefixListId"] = undefined /*out*/;
            inputs["requesterManaged"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcEndpoint.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpoint resources.
 */
export interface VpcEndpointState {
    /**
     * The Amazon Resource Name (ARN) of the VPC endpoint.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
     */
    readonly autoAccept?: pulumi.Input<boolean>;
    /**
     * The list of CIDR blocks for the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    readonly cidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The DNS entries for the VPC Endpoint. Applicable for endpoints of type `Interface`. DNS blocks are documented below.
     */
    readonly dnsEntries?: pulumi.Input<pulumi.Input<inputs.ec2.VpcEndpointDnsEntry>[]>;
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of the AWS account that owns the VPC endpoint.
     */
    readonly ownerId?: pulumi.Input<string>;
    /**
     * A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * The prefix list ID of the exposed AWS service. Applicable for endpoints of type `Gateway`.
     */
    readonly prefixListId?: pulumi.Input<string>;
    /**
     * Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
     * Defaults to `false`.
     */
    readonly privateDnsEnabled?: pulumi.Input<boolean>;
    /**
     * Whether or not the VPC Endpoint is being managed by its service - `true` or `false`.
     */
    readonly requesterManaged?: pulumi.Input<boolean>;
    /**
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     */
    readonly routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
     */
    readonly serviceName?: pulumi.Input<string>;
    /**
     * The state of the VPC endpoint.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
     */
    readonly vpcEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the VPC in which the endpoint will be used.
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpoint resource.
 */
export interface VpcEndpointArgs {
    /**
     * Accept the VPC endpoint (the VPC endpoint and service need to be in the same AWS account).
     */
    readonly autoAccept?: pulumi.Input<boolean>;
    /**
     * A policy to attach to the endpoint that controls access to the service. Defaults to full access. All `Gateway` and some `Interface` endpoints support policies - see the [relevant AWS documentation](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-endpoints-access.html) for more details.
     */
    readonly policy?: pulumi.Input<string>;
    /**
     * Whether or not to associate a private hosted zone with the specified VPC. Applicable for endpoints of type `Interface`.
     * Defaults to `false`.
     */
    readonly privateDnsEnabled?: pulumi.Input<boolean>;
    /**
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     */
    readonly routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The service name. For AWS services the service name is usually in the form `com.amazonaws.<region>.<service>` (the SageMaker Notebook service is an exception to this rule, the service name is in the form `aws.sagemaker.<region>.notebook`).
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
     */
    readonly vpcEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the VPC in which the endpoint will be used.
     */
    readonly vpcId: pulumi.Input<string>;
}
