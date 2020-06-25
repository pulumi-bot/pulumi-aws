// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Amazon API Gateway Version 2 VPC Link.
 *
 * > **Note:** Amazon API Gateway Version 2 VPC Links enable private integrations that connect HTTP APIs to private resources in a VPC.
 * To enable private integration for REST APIs, use the `Amazon API Gateway Version 1 VPC Link` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.apigatewayv2.VpcLink("example", {
 *     securityGroupIds: [data.aws_security_group.example.id],
 *     subnetIds: data.aws_subnet_ids.example.ids,
 *     tags: {
 *         Usage: "example",
 *     },
 * });
 * ```
 */
export class VpcLink extends pulumi.CustomResource {
    /**
     * Get an existing VpcLink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcLinkState, opts?: pulumi.CustomResourceOptions): VpcLink {
        return new VpcLink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:apigatewayv2/vpcLink:VpcLink';

    /**
     * Returns true if the given object is an instance of VpcLink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcLink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcLink.__pulumiType;
    }

    /**
     * The VPC Link ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name of the VPC Link.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Security group IDs for the VPC Link.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * Subnet IDs for the VPC Link.
     */
    public readonly subnetIds!: pulumi.Output<string[]>;
    /**
     * A map of tags to assign to the VPC Link.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a VpcLink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcLinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcLinkArgs | VpcLinkState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as VpcLinkState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            inputs["subnetIds"] = state ? state.subnetIds : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as VpcLinkArgs | undefined;
            if (!args || args.securityGroupIds === undefined) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if (!args || args.subnetIds === undefined) {
                throw new Error("Missing required property 'subnetIds'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            inputs["subnetIds"] = args ? args.subnetIds : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VpcLink.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcLink resources.
 */
export interface VpcLinkState {
    /**
     * The VPC Link ARN.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name of the VPC Link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Security group IDs for the VPC Link.
     */
    readonly securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subnet IDs for the VPC Link.
     */
    readonly subnetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the VPC Link.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a VpcLink resource.
 */
export interface VpcLinkArgs {
    /**
     * The name of the VPC Link.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Security group IDs for the VPC Link.
     */
    readonly securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Subnet IDs for the VPC Link.
     */
    readonly subnetIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A map of tags to assign to the VPC Link.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
