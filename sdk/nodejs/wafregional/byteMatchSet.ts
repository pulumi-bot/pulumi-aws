// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Byte Match Set Resource for use with Application Load Balancer.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const byteSet = new aws.wafregional.ByteMatchSet("byteSet", {
 *     byteMatchTuples: [{
 *         fieldToMatch: {
 *             data: "referer",
 *             type: "HEADER",
 *         },
 *         positionalConstraint: "CONTAINS",
 *         targetString: "badrefer1",
 *         textTransformation: "NONE",
 *     }],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_byte_match_set.html.markdown.
 */
export class ByteMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing ByteMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ByteMatchSetState, opts?: pulumi.CustomResourceOptions): ByteMatchSet {
        return new ByteMatchSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/byteMatchSet:ByteMatchSet';

    /**
     * Returns true if the given object is an instance of ByteMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ByteMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ByteMatchSet.__pulumiType;
    }

    /**
     * Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
     */
    public readonly byteMatchTuples!: pulumi.Output<{ fieldToMatch: { data?: string, type: string }, positionalConstraint: string, targetString?: string, textTransformation: string }[] | undefined>;
    /**
     * The name or description of the ByteMatchSet.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a ByteMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ByteMatchSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ByteMatchSetArgs | ByteMatchSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ByteMatchSetState | undefined;
            inputs["byteMatchTuples"] = state ? state.byteMatchTuples : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as ByteMatchSetArgs | undefined;
            inputs["byteMatchTuples"] = args ? args.byteMatchTuples : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ByteMatchSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ByteMatchSet resources.
 */
export interface ByteMatchSetState {
    /**
     * Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
     */
    readonly byteMatchTuples?: pulumi.Input<pulumi.Input<{ fieldToMatch: pulumi.Input<{ data?: pulumi.Input<string>, type: pulumi.Input<string> }>, positionalConstraint: pulumi.Input<string>, targetString?: pulumi.Input<string>, textTransformation: pulumi.Input<string> }>[]>;
    /**
     * The name or description of the ByteMatchSet.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ByteMatchSet resource.
 */
export interface ByteMatchSetArgs {
    /**
     * Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
     */
    readonly byteMatchTuples?: pulumi.Input<pulumi.Input<{ fieldToMatch: pulumi.Input<{ data?: pulumi.Input<string>, type: pulumi.Input<string> }>, positionalConstraint: pulumi.Input<string>, targetString?: pulumi.Input<string>, textTransformation: pulumi.Input<string> }>[]>;
    /**
     * The name or description of the ByteMatchSet.
     */
    readonly name?: pulumi.Input<string>;
}
