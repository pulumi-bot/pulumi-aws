// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an association between a VPC endpoint and a subnet.
 * 
 * > **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
 * both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
 * and a single `subnetId`) and a VPC Endpoint resource with a `subnetIds`
 * attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
 * Association resource. Doing so will cause a conflict of associations and will overwrite the association.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/vpc_endpoint_subnet_association.html.markdown.
 */
export class VpcEndpointSubnetAssociation extends pulumi.CustomResource {
    /**
     * Get an existing VpcEndpointSubnetAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcEndpointSubnetAssociationState, opts?: pulumi.CustomResourceOptions): VpcEndpointSubnetAssociation {
        return new VpcEndpointSubnetAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation';

    /**
     * Returns true if the given object is an instance of VpcEndpointSubnetAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcEndpointSubnetAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcEndpointSubnetAssociation.__pulumiType;
    }

    /**
     * The ID of the subnet to be associated with the VPC endpoint.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The ID of the VPC endpoint with which the subnet will be associated.
     */
    public readonly vpcEndpointId!: pulumi.Output<string>;

    /**
     * Create a VpcEndpointSubnetAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcEndpointSubnetAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcEndpointSubnetAssociationArgs | VpcEndpointSubnetAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcEndpointSubnetAssociationState | undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["vpcEndpointId"] = state ? state.vpcEndpointId : undefined;
        } else {
            const args = argsOrState as VpcEndpointSubnetAssociationArgs | undefined;
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            if (!args || args.vpcEndpointId === undefined) {
                throw new Error("Missing required property 'vpcEndpointId'");
            }
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["vpcEndpointId"] = args ? args.vpcEndpointId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcEndpointSubnetAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcEndpointSubnetAssociation resources.
 */
export interface VpcEndpointSubnetAssociationState {
    /**
     * The ID of the subnet to be associated with the VPC endpoint.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The ID of the VPC endpoint with which the subnet will be associated.
     */
    readonly vpcEndpointId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VpcEndpointSubnetAssociation resource.
 */
export interface VpcEndpointSubnetAssociationArgs {
    /**
     * The ID of the subnet to be associated with the VPC endpoint.
     */
    readonly subnetId: pulumi.Input<string>;
    /**
     * The ID of the VPC endpoint with which the subnet will be associated.
     */
    readonly vpcEndpointId: pulumi.Input<string>;
}
