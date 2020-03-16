// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides network associations for AWS Client VPN endpoints. For more information on usage, please see the 
 * [AWS Client VPN Administrator's Guide](https://docs.aws.amazon.com/vpn/latest/clientvpn-admin/what-is.html).
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ec2clientvpn.NetworkAssociation("example", {
 *     clientVpnEndpointId: aws_ec2_client_vpn_endpoint_example.id,
 *     subnetId: aws_subnet_example.id,
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_client_vpn_network_association.html.markdown.
 */
export class NetworkAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkAssociationState, opts?: pulumi.CustomResourceOptions): NetworkAssociation {
        return new NetworkAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2clientvpn/networkAssociation:NetworkAssociation';

    /**
     * Returns true if the given object is an instance of NetworkAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkAssociation.__pulumiType;
    }

    /**
     * The ID of the Client VPN endpoint.
     */
    public readonly clientVpnEndpointId!: pulumi.Output<string>;
    /**
     * The IDs of the security groups applied to the target network association.
     */
    public /*out*/ readonly securityGroups!: pulumi.Output<string[]>;
    /**
     * The current state of the target network association.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The ID of the VPC in which the target network (subnet) is located. 
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a NetworkAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkAssociationArgs | NetworkAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkAssociationState | undefined;
            inputs["clientVpnEndpointId"] = state ? state.clientVpnEndpointId : undefined;
            inputs["securityGroups"] = state ? state.securityGroups : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
            inputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as NetworkAssociationArgs | undefined;
            if (!args || args.clientVpnEndpointId === undefined) {
                throw new Error("Missing required property 'clientVpnEndpointId'");
            }
            if (!args || args.subnetId === undefined) {
                throw new Error("Missing required property 'subnetId'");
            }
            inputs["clientVpnEndpointId"] = args ? args.clientVpnEndpointId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
            inputs["securityGroups"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["vpcId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkAssociation resources.
 */
export interface NetworkAssociationState {
    /**
     * The ID of the Client VPN endpoint.
     */
    readonly clientVpnEndpointId?: pulumi.Input<string>;
    /**
     * The IDs of the security groups applied to the target network association.
     */
    readonly securityGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The current state of the target network association.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    readonly subnetId?: pulumi.Input<string>;
    /**
     * The ID of the VPC in which the target network (subnet) is located. 
     */
    readonly vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkAssociation resource.
 */
export interface NetworkAssociationArgs {
    /**
     * The ID of the Client VPN endpoint.
     */
    readonly clientVpnEndpointId: pulumi.Input<string>;
    /**
     * The ID of the subnet to associate with the Client VPN endpoint.
     */
    readonly subnetId: pulumi.Input<string>;
}
