// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a VPC Endpoint resource.
 * 
 * > **NOTE on VPC Endpoints and VPC Endpoint Associations:** Terraform provides both standalone VPC Endpoint Associations for
 * Route Tables - (an association between a VPC endpoint and a single `route_table_id`) and
 * Subnets - (an association between a VPC endpoint and a single `subnet_id`) and
 * a VPC Endpoint resource with `route_table_ids` and `subnet_ids` attributes.
 * Do not use the same resource ID in both a VPC Endpoint resource and a VPC Endpoint Association resource.
 * Doing so will cause a conflict of associations and will overwrite the association.
 * 
 * ## Example Usage
 * 
 * Basic usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const s3 = new aws.ec2.VpcEndpoint("s3", {
 *     serviceName: "com.amazonaws.us-west-2.s3",
 *     vpcId: aws_vpc_main.id,
 * });
 * ```
 * 
 * Interface type usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ec2 = new aws.ec2.VpcEndpoint("ec2", {
 *     privateDnsEnabled: true,
 *     securityGroupIds: [aws_security_group_sg1.id],
 *     serviceName: "com.amazonaws.us-west-2.ec2",
 *     vpcEndpointType: "Interface",
 *     vpcId: aws_vpc_main.id,
 * });
 * ```
 * 
 * Custom Service Usage:
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const ptfeServiceVpcEndpoint = new aws.ec2.VpcEndpoint("ptfe_service", {
 *     privateDnsEnabled: false,
 *     securityGroupIds: [aws_security_group_ptfe_service.id],
 *     serviceName: var_ptfe_service,
 *     subnetIds: [local_subnet_ids],
 *     vpcEndpointType: "Interface",
 *     vpcId: var_vpc_id,
 * });
 * const internal = pulumi.output(aws.route53.getZone({
 *     name: "vpc.internal.",
 *     privateZone: true,
 *     vpcId: var_vpc_id,
 * }));
 * const ptfeServiceRecord = new aws.route53.Record("ptfe_service", {
 *     records: [ptfeServiceVpcEndpoint.dnsEntries.apply(dnsEntries => (<any>dnsEntries[0])["dns_name"])],
 *     ttl: 300,
 *     type: "CNAME",
 *     zoneId: internal.apply(internal => internal.zoneId),
 * });
 * ```
 * 
 * > **NOTE The `dns_entry` output is a list of maps:** Terraform interpolation support for lists of maps requires the `lookup` and `[]` until full support of lists of maps is available
 */
export class VpcEndpoint extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpoint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointState, opts?: pulumi.CustomResourceOptions): VpcEndpoint {
        return new VpcEndpoint(name, <any>state, { ...opts, id: id });
    }

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
    public /*out*/ readonly dnsEntries!: pulumi.Output<{ dnsName: string, hostedZoneId: string }[]>;
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     */
    public /*out*/ readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
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
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     */
    public readonly routeTableIds!: pulumi.Output<string[]>;
    /**
     * The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The service name, in the form `com.amazonaws.region.service` for AWS services.
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
            inputs["autoAccept"] = state ? state.autoAccept : undefined;
            inputs["cidrBlocks"] = state ? state.cidrBlocks : undefined;
            inputs["dnsEntries"] = state ? state.dnsEntries : undefined;
            inputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            inputs["policy"] = state ? state.policy : undefined;
            inputs["prefixListId"] = state ? state.prefixListId : undefined;
            inputs["privateDnsEnabled"] = state ? state.privateDnsEnabled : undefined;
            inputs["routeTableIds"] = state ? state.routeTableIds : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["serviceName"] = state ? state.serviceName : undefined;
            inputs["state"] = state ? state.state : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
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
            inputs["vpcEndpointType"] = args ? args.vpcEndpointType : undefined;
            inputs["vpcId"] = args ? args.vpcId : undefined;
            inputs["cidrBlocks"] = undefined /*out*/;
            inputs["dnsEntries"] = undefined /*out*/;
            inputs["networkInterfaceIds"] = undefined /*out*/;
            inputs["prefixListId"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super("aws:ec2/vpcEndpoint:VpcEndpoint", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpoint resources.
 */
export interface VpcEndpointState {
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
    readonly dnsEntries?: pulumi.Input<pulumi.Input<{ dnsName?: pulumi.Input<string>, hostedZoneId?: pulumi.Input<string> }>[]>;
    /**
     * One or more network interfaces for the VPC Endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
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
     * One or more route table IDs. Applicable for endpoints of type `Gateway`.
     */
    readonly routeTableIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The ID of one or more security groups to associate with the network interface. Required for endpoints of type `Interface`.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The service name, in the form `com.amazonaws.region.service` for AWS services.
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
     * A policy to attach to the endpoint that controls access to the service. Applicable for endpoints of type `Gateway`. Defaults to full access. For more information about building AWS IAM policy documents with Terraform, see the [AWS IAM Policy Document Guide](https://www.terraform.io/docs/providers/aws/guides/iam-policy-documents.html).
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
     * The service name, in the form `com.amazonaws.region.service` for AWS services.
     */
    readonly serviceName: pulumi.Input<string>;
    /**
     * The ID of one or more subnets in which to create a network interface for the endpoint. Applicable for endpoints of type `Interface`.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The VPC endpoint type, `Gateway` or `Interface`. Defaults to `Gateway`.
     */
    readonly vpcEndpointType?: pulumi.Input<string>;
    /**
     * The ID of the VPC in which the endpoint will be used.
     */
    readonly vpcId: pulumi.Input<string>;
}
