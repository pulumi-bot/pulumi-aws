// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides an IP access control group in AWS WorkSpaces Service
 * 
 * {{% examples %}}
 * ## Example Usage
 * {{% example %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const contractors = new aws.workspaces.IpGroup("contractors", {
 *     description: "Contractors IP access control group",
 * });
 * ```
 * 
 * {{% /example %}}
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/workspaces_ip_group.html.markdown.
 */
export class IpGroup extends pulumi.CustomResource {
    /**
     * Get an existing IpGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpGroupState, opts?: pulumi.CustomResourceOptions): IpGroup {
        return new IpGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:workspaces/ipGroup:IpGroup';

    /**
     * Returns true if the given object is an instance of IpGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpGroup.__pulumiType;
    }

    /**
     * The description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the IP group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    public readonly rules!: pulumi.Output<outputs.workspaces.IpGroupRule[] | undefined>;
    public readonly tags!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a IpGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: IpGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpGroupArgs | IpGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IpGroupState | undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["rules"] = state ? state.rules : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as IpGroupArgs | undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["tags"] = args ? args.tags : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IpGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpGroup resources.
 */
export interface IpGroupState {
    /**
     * The description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the IP group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.workspaces.IpGroupRule>[]>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a IpGroup resource.
 */
export interface IpGroupArgs {
    /**
     * The description.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the IP group.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more pairs specifying the IP group rule (in CIDR format) from which web requests originate.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.workspaces.IpGroupRule>[]>;
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
