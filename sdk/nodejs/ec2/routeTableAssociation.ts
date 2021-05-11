// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create an association between a route table and a subnet or a route table and an
 * internet gateway or virtual private gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const routeTableAssociation = new aws.ec2.RouteTableAssociation("routeTableAssociation", {
 *     subnetId: aws_subnet.foo.id,
 *     routeTableId: aws_route_table.bar.id,
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const routeTableAssociation = new aws.ec2.RouteTableAssociation("routeTableAssociation", {
 *     gatewayId: aws_internet_gateway.foo.id,
 *     routeTableId: aws_route_table.bar.id,
 * });
 * ```
 *
 * ## Import
 *
 * is already associated, will result in an error (e.g., `Resource.AlreadyAssociatedthe specified association for route table rtb-4176657279 conflicts with an existing association`) unless you first import the original association. EC2 Route Table Associations can be imported using the associated resource ID and Route Table ID separated by a forward slash (`/`). For example with EC2 Subnets
 *
 * ```sh
 *  $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc subnet-6777656e646f6c796e/rtb-656c65616e6f72
 * ```
 *
 *  For example with EC2 Internet Gateways
 *
 * ```sh
 *  $ pulumi import aws:ec2/routeTableAssociation:RouteTableAssociation assoc igw-01b3a60780f8d034a/rtb-656c65616e6f72
 * ```
 */
export class RouteTableAssociation extends pulumi.CustomResource {
    /**
     * Get an existing RouteTableAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteTableAssociationState, opts?: pulumi.CustomResourceOptions): RouteTableAssociation {
        return new RouteTableAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/routeTableAssociation:RouteTableAssociation';

    /**
     * Returns true if the given object is an instance of RouteTableAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteTableAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteTableAssociation.__pulumiType;
    }

    /**
     * The gateway ID to create an association. Conflicts with `subnetId`.
     */
    public readonly gatewayId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the routing table to associate with.
     */
    public readonly routeTableId!: pulumi.Output<string>;
    /**
     * The subnet ID to create an association. Conflicts with `gatewayId`.
     */
    public readonly subnetId!: pulumi.Output<string | undefined>;

    /**
     * Create a RouteTableAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteTableAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteTableAssociationArgs | RouteTableAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteTableAssociationState | undefined;
            inputs["gatewayId"] = state ? state.gatewayId : undefined;
            inputs["routeTableId"] = state ? state.routeTableId : undefined;
            inputs["subnetId"] = state ? state.subnetId : undefined;
        } else {
            const args = argsOrState as RouteTableAssociationArgs | undefined;
            if ((!args || args.routeTableId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTableId'");
            }
            inputs["gatewayId"] = args ? args.gatewayId : undefined;
            inputs["routeTableId"] = args ? args.routeTableId : undefined;
            inputs["subnetId"] = args ? args.subnetId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RouteTableAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteTableAssociation resources.
 */
export interface RouteTableAssociationState {
    /**
     * The gateway ID to create an association. Conflicts with `subnetId`.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The ID of the routing table to associate with.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The subnet ID to create an association. Conflicts with `gatewayId`.
     */
    subnetId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteTableAssociation resource.
 */
export interface RouteTableAssociationArgs {
    /**
     * The gateway ID to create an association. Conflicts with `subnetId`.
     */
    gatewayId?: pulumi.Input<string>;
    /**
     * The ID of the routing table to associate with.
     */
    routeTableId: pulumi.Input<string>;
    /**
     * The subnet ID to create an association. Conflicts with `gatewayId`.
     */
    subnetId?: pulumi.Input<string>;
}
