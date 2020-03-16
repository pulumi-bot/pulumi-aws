// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF Size Constraint Set Resource
 * 
 * ## Example Usage
 * 
 * {{% examples %}}
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const sizeConstraintSet = new aws.waf.SizeConstraintSet("sizeConstraintSet", {
 *     sizeConstraints: [{
 *         comparisonOperator: "EQ",
 *         fieldToMatch: {
 *             type: "BODY",
 *         },
 *         size: 4096,
 *         textTransformation: "NONE",
 *     }],
 * });
 * ```
 * 
 * {{% /examples %}}
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_size_constraint_set.html.markdown.
 */
export class SizeConstraintSet extends pulumi.CustomResource {
    /**
     * Get an existing SizeConstraintSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SizeConstraintSetState, opts?: pulumi.CustomResourceOptions): SizeConstraintSet {
        return new SizeConstraintSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:waf/sizeConstraintSet:SizeConstraintSet';

    /**
     * Returns true if the given object is an instance of SizeConstraintSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SizeConstraintSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SizeConstraintSet.__pulumiType;
    }

    /**
     * Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The name or description of the Size Constraint Set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the parts of web requests that you want to inspect the size of.
     */
    public readonly sizeConstraints!: pulumi.Output<outputs.waf.SizeConstraintSetSizeConstraint[] | undefined>;

    /**
     * Create a SizeConstraintSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SizeConstraintSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SizeConstraintSetArgs | SizeConstraintSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SizeConstraintSetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["sizeConstraints"] = state ? state.sizeConstraints : undefined;
        } else {
            const args = argsOrState as SizeConstraintSetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["sizeConstraints"] = args ? args.sizeConstraints : undefined;
            inputs["arn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SizeConstraintSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SizeConstraintSet resources.
 */
export interface SizeConstraintSetState {
    /**
     * Amazon Resource Name (ARN)
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * The name or description of the Size Constraint Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the parts of web requests that you want to inspect the size of.
     */
    readonly sizeConstraints?: pulumi.Input<pulumi.Input<inputs.waf.SizeConstraintSetSizeConstraint>[]>;
}

/**
 * The set of arguments for constructing a SizeConstraintSet resource.
 */
export interface SizeConstraintSetArgs {
    /**
     * The name or description of the Size Constraint Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the parts of web requests that you want to inspect the size of.
     */
    readonly sizeConstraints?: pulumi.Input<pulumi.Input<inputs.waf.SizeConstraintSetSizeConstraint>[]>;
}
