// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * Provides an AWS WAFv2 Regex Pattern Set Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafv2.RegexPatternSet("example", {
 *     description: "Example regex pattern set",
 *     regularExpressions: [
 *         {
 *             regexString: "one",
 *         },
 *         {
 *             regexString: "two",
 *         },
 *     ],
 *     scope: "REGIONAL",
 *     tags: {
 *         Tag1: "Value1",
 *         Tag2: "Value2",
 *     },
 * });
 * ```
 */
export class RegexPatternSet extends pulumi.CustomResource {
    /**
     * Get an existing RegexPatternSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegexPatternSetState, opts?: pulumi.CustomResourceOptions): RegexPatternSet {
        return new RegexPatternSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafv2/regexPatternSet:RegexPatternSet';

    /**
     * Returns true if the given object is an instance of RegexPatternSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegexPatternSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegexPatternSet.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) that identifies the cluster.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * A friendly description of the regular expression pattern set.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly lockToken!: pulumi.Output<string>;
    /**
     * A friendly name of the regular expression pattern set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
     */
    public readonly regularExpressions!: pulumi.Output<outputs.wafv2.RegexPatternSetRegularExpression[] | undefined>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    public readonly scope!: pulumi.Output<string>;
    /**
     * An array of key:value pairs to associate with the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;

    /**
     * Create a RegexPatternSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegexPatternSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegexPatternSetArgs | RegexPatternSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegexPatternSetState | undefined;
            inputs["arn"] = state ? state.arn : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["lockToken"] = state ? state.lockToken : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["regularExpressions"] = state ? state.regularExpressions : undefined;
            inputs["scope"] = state ? state.scope : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as RegexPatternSetArgs | undefined;
            if (!args || args.scope === undefined) {
                throw new Error("Missing required property 'scope'");
            }
            inputs["description"] = args ? args.description : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["regularExpressions"] = args ? args.regularExpressions : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["lockToken"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegexPatternSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegexPatternSet resources.
 */
export interface RegexPatternSetState {
    /**
     * The Amazon Resource Name (ARN) that identifies the cluster.
     */
    readonly arn?: pulumi.Input<string>;
    /**
     * A friendly description of the regular expression pattern set.
     */
    readonly description?: pulumi.Input<string>;
    readonly lockToken?: pulumi.Input<string>;
    /**
     * A friendly name of the regular expression pattern set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
     */
    readonly regularExpressions?: pulumi.Input<pulumi.Input<inputs.wafv2.RegexPatternSetRegularExpression>[]>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    readonly scope?: pulumi.Input<string>;
    /**
     * An array of key:value pairs to associate with the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a RegexPatternSet resource.
 */
export interface RegexPatternSetArgs {
    /**
     * A friendly description of the regular expression pattern set.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * A friendly name of the regular expression pattern set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
     */
    readonly regularExpressions?: pulumi.Input<pulumi.Input<inputs.wafv2.RegexPatternSetRegularExpression>[]>;
    /**
     * Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
     */
    readonly scope: pulumi.Input<string>;
    /**
     * An array of key:value pairs to associate with the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
