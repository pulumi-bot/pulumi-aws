// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a MediaStore Container.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const example = new aws.mediastore.Container("example", {});
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/media_store_container.html.markdown.
 */
export class Container extends pulumi.CustomResource {
    /**
     * Get an existing Container resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContainerState, opts?: pulumi.CustomResourceOptions): Container {
        return new Container(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:mediastore/container:Container';

    /**
     * Returns true if the given object is an instance of Container.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Container {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Container.__pulumiType;
    }

    /**
     * The ARN of the container.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The DNS endpoint of the container.
     */
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    /**
     * The name of the container. Must contain alphanumeric characters or underscores.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Container resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ContainerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContainerArgs | ContainerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ContainerState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["endpoint"] = state ? state.endpoint : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as ContainerArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["endpoint"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Container.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Container resources.
 */
export interface ContainerState {
    /**
     * The ARN of the container.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The DNS endpoint of the container.
     */
    readonly endpoint?: pulumi.Input<string>;
    /**
     * The name of the container. Must contain alphanumeric characters or underscores.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Container resource.
 */
export interface ContainerArgs {
    /**
     * The name of the container. Must contain alphanumeric characters or underscores.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
