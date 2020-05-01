// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Enables AWS Shield Advanced for a specific AWS resource.
 * The resource can be an Amazon CloudFront distribution, Elastic Load Balancing load balancer, AWS Global Accelerator accelerator, Elastic IP Address, or an Amazon Route 53 hosted zone.
 * 
 * ## Example Usage
 * 
 * ### Create protection
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const available = pulumi.output(aws.getAvailabilityZones({ async: true }));
 * const currentRegion = pulumi.output(aws.getRegion({ async: true }));
 * const currentCallerIdentity = pulumi.output(aws.getCallerIdentity({ async: true }));
 * const fooEip = new aws.ec2.Eip("foo", {
 *     vpc: true,
 * });
 * const fooProtection = new aws.shield.Protection("foo", {
 *     name: var_name,
 *     resourceArn: pulumi.interpolate`arn:aws:ec2:${currentRegion.name!}:${currentCallerIdentity.accountId}:eip-allocation/${fooEip.id}`,
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/shield_protection.html.markdown.
 */
export class Protection extends pulumi.CustomResource {
    /**
     * Get an existing Protection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProtectionState, opts?: pulumi.CustomResourceOptions): Protection {
        return new Protection(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:shield/protection:Protection';

    /**
     * Returns true if the given object is an instance of Protection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Protection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Protection.__pulumiType;
    }

    /**
     * A friendly name for the Protection you are creating.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource to be protected.
     */
    public readonly resourceArn!: pulumi.Output<string>;

    /**
     * Create a Protection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProtectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProtectionArgs | ProtectionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProtectionState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceArn"] = state ? state.resourceArn : undefined;
        } else {
            const args = argsOrState as ProtectionArgs | undefined;
            if (!args || args.resourceArn === undefined) {
                throw new Error("Missing required property 'resourceArn'");
            }
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceArn"] = args ? args.resourceArn : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Protection.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Protection resources.
 */
export interface ProtectionState {
    /**
     * A friendly name for the Protection you are creating.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource to be protected.
     */
    readonly resourceArn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Protection resource.
 */
export interface ProtectionArgs {
    /**
     * A friendly name for the Protection you are creating.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ARN (Amazon Resource Name) of the resource to be protected.
     */
    readonly resourceArn: pulumi.Input<string>;
}
