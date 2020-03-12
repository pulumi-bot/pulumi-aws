// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Adds launch permission to Amazon Machine Image (AMI) from another AWS account.
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.ec2.AmiLaunchPermission("example", {
 *     accountId: "123456789012",
 *     imageId: "ami-12345678",
 * });
 * ```
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ami_launch_permission.html.markdown.
 */
export class AmiLaunchPermission extends pulumi.CustomResource {
    /**
     * Get an existing AmiLaunchPermission resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AmiLaunchPermissionState, opts?: pulumi.CustomResourceOptions): AmiLaunchPermission {
        return new AmiLaunchPermission(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:ec2/amiLaunchPermission:AmiLaunchPermission';

    /**
     * Returns true if the given object is an instance of AmiLaunchPermission.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AmiLaunchPermission {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AmiLaunchPermission.__pulumiType;
    }

    /**
     * An AWS Account ID to add launch permissions.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * A region-unique name for the AMI.
     */
    public readonly imageId!: pulumi.Output<string>;

    /**
     * Create a AmiLaunchPermission resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AmiLaunchPermissionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AmiLaunchPermissionArgs | AmiLaunchPermissionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AmiLaunchPermissionState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["imageId"] = state ? state.imageId : undefined;
        } else {
            const args = argsOrState as AmiLaunchPermissionArgs | undefined;
            if (!args || args.accountId === undefined) {
                throw new Error("Missing required property 'accountId'");
            }
            if (!args || args.imageId === undefined) {
                throw new Error("Missing required property 'imageId'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["imageId"] = args ? args.imageId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AmiLaunchPermission.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AmiLaunchPermission resources.
 */
export interface AmiLaunchPermissionState {
    /**
     * An AWS Account ID to add launch permissions.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * A region-unique name for the AMI.
     */
    readonly imageId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AmiLaunchPermission resource.
 */
export interface AmiLaunchPermissionArgs {
    /**
     * An AWS Account ID to add launch permissions.
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * A region-unique name for the AMI.
     */
    readonly imageId: pulumi.Input<string>;
}
