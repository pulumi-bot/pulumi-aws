// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a VPC Endpoint Route Table Association
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ec2.VpcEndpointRouteTableAssociation("example", {
 *     routeTableId: aws_route_table_example.id,
 *     vpcEndpointId: aws_vpc_endpoint_example.id,
 * });
 * ```
 */
export class VpcEndpointRouteTableAssociation extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointRouteTableAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointRouteTableAssociationState, opts?: pulumi.CustomResourceOptions): VpcEndpointRouteTableAssociation {
        return new VpcEndpointRouteTableAssociation(name, <any>state, { ...opts, id: id });
    }

    /**
     * Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
     */
    public readonly vpcEndpointId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointRouteTableAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointRouteTableAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointRouteTableAssociationArgs | VpcEndpointRouteTableAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcEndpointRouteTableAssociationState | undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["vpcEndpointId"] = state ? state.vpcEndpointId : undefined;
        } else {
            const args = argsOrState as VpcEndpointRouteTableAssociationArgs | undefined;
            if (!args || args.routeTableId === undefined) {
                throw new Error("Missing required property 'routeTableId'");
            }
            if (!args || args.vpcEndpointId === undefined) {
                throw new Error("Missing required property 'vpcEndpointId'");
            }
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
        }
        super("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointRouteTableAssociation resources.
 */
export interface VpcEndpointRouteTableAssociationState {
    /**
     * Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
     */
    readonly routeTableId?: pulumi.Input<string>;
    /**
     * Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
     */
    readonly vpcEndpointId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointRouteTableAssociation resource.
 */
export interface VpcEndpointRouteTableAssociationArgs {
    /**
     * Identifier of the EC2 Route Table to be associated with the VPC Endpoint.
     */
    readonly routeTableId: pulumi.Input<string>;
    /**
     * Identifier of the VPC Endpoint with which the EC2 Route Table will be associated.
     */
    readonly vpcEndpointId: pulumi.Input<string>;
}
