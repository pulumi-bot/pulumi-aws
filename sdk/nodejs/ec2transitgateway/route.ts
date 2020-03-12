// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an EC2 Transit Gateway Route.
 * 
 * ## Example Usage
 * 
 * ### Standard usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ec2transitgateway.Route("example", {
 *     destinationCidrBlock: "0.0.0.0/0",
 *     transitGatewayAttachmentId: aws_ec2_transit_gateway_vpc_attachment_example.id,
 *     transitGatewayRouteTableId: aws_ec2_transit_gateway_example.associationDefaultRouteTableId,
 * });
 * ```
 * {{% /examples %}}
 * 
 * ### Blackhole route
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ec2transitgateway.Route("example", {
 *     blackhole: true,
 *     destinationCidrBlock: "0.0.0.0/0",
 *     transitGatewayRouteTableId: aws_ec2_transit_gateway_example.associationDefaultRouteTableId,
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ec2_transit_gateway_route.html.markdown.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteState, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2transitgateway/route:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * Indicates whether to drop traffic that matches this route (default to `false`).
     */
    public readonly blackhole!: pulumi.Output<boolean | undefined>;
    /**
     * IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
     */
    public readonly destinationCidrBlock!: pulumi.Output<string>;
    /**
     * Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
     */
    public readonly transitGatewayAttachmentId!: pulumi.Output<string | undefined>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    public readonly transitGatewayRouteTableId!: pulumi.Output<string>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteArgs | RouteState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RouteState | undefined;
            inputs["blackhole"] = state ? state.blackhole : undefined;
            inputs["destinationCidrBlock"] = state ? state.destinationCidrBlock : undefined;
            inputs["transitGatewayAttachmentId"] = state ? state.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayRouteTableId"] = state ? state.transitGatewayRouteTableId : undefined;
        } else {
            const args = argsOrState as RouteArgs | undefined;
            if (!args || args.destinationCidrBlock === undefined) {
                throw new Error("Missing required property 'destinationCidrBlock'");
            }
            if (!args || args.transitGatewayRouteTableId === undefined) {
                throw new Error("Missing required property 'transitGatewayRouteTableId'");
            }
            inputs["blackhole"] = args ? args.blackhole : undefined;
            inputs["destinationCidrBlock"] = args ? args.destinationCidrBlock : undefined;
            inputs["transitGatewayAttachmentId"] = args ? args.transitGatewayAttachmentId : undefined;
            inputs["transitGatewayRouteTableId"] = args ? args.transitGatewayRouteTableId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Route resources.
 */
export interface RouteState {
    /**
     * Indicates whether to drop traffic that matches this route (default to `false`).
     */
    readonly blackhole?: pulumi.Input<boolean>;
    /**
     * IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
     */
    readonly destinationCidrBlock?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
     */
    readonly transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    readonly transitGatewayRouteTableId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * Indicates whether to drop traffic that matches this route (default to `false`).
     */
    readonly blackhole?: pulumi.Input<boolean>;
    /**
     * IPv4 CIDR range used for destination matches. Routing decisions are based on the most specific match.
     */
    readonly destinationCidrBlock: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Attachment (required if `blackhole` is set to false).
     */
    readonly transitGatewayAttachmentId?: pulumi.Input<string>;
    /**
     * Identifier of EC2 Transit Gateway Route Table.
     */
    readonly transitGatewayRouteTableId: pulumi.Input<string>;
}
