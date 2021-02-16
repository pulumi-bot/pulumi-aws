// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Regex Pattern Set Resource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.wafregional.RegexPatternSet("example", {
 *     regexPatternStrings: [
 *         "one",
 *         "two",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * WAF Regional Regex Pattern Set can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import aws:wafregional/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
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
    public static readonly __pulumiType = 'aws:wafregional/regexPatternSet:RegexPatternSet';

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
     * The name or description of the Regex Pattern Set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
     */
    public readonly regexPatternStrings!: pulumi.Output<string[] | undefined>;

    /**
     * Create a RegexPatternSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegexPatternSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegexPatternSetArgs | RegexPatternSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RegexPatternSetState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["regexPatternStrings"] = state ? state.regexPatternStrings : undefined;
        } else {
            const args = argsOrState as RegexPatternSetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["regexPatternStrings"] = args ? args.regexPatternStrings : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegexPatternSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegexPatternSet resources.
 */
export interface RegexPatternSetState {
    /**
     * The name or description of the Regex Pattern Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
     */
    readonly regexPatternStrings?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a RegexPatternSet resource.
 */
export interface RegexPatternSetArgs {
    /**
     * The name or description of the Regex Pattern Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
     */
    readonly regexPatternStrings?: pulumi.Input<pulumi.Input<string>[]>;
}
