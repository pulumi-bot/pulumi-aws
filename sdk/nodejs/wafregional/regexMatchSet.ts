// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Regex Match Set Resource
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_regex_match_set.html.markdown.
 */
export class RegexMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing RegexMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RegexMatchSetState, opts?: pulumi.CustomResourceOptions): RegexMatchSet {
        return new RegexMatchSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/regexMatchSet:RegexMatchSet';

    /**
     * Returns true if the given object is an instance of RegexMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegexMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegexMatchSet.__pulumiType;
    }

    /**
     * The name or description of the Regex Match Set.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The regular expression pattern that you want AWS WAF to search for in web requests,
     * the location in requests that you want AWS WAF to search, and other settings. See below.
     */
    public readonly regexMatchTuples!: pulumi.Output<outputs.wafregional.RegexMatchSetRegexMatchTuple[] | undefined>;

    /**
     * Create a RegexMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RegexMatchSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RegexMatchSetArgs | RegexMatchSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as RegexMatchSetState | undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["regexMatchTuples"] = state ? state.regexMatchTuples : undefined;
        } else {
            const args = argsOrState as RegexMatchSetArgs | undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["regexMatchTuples"] = args ? args.regexMatchTuples : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(RegexMatchSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RegexMatchSet resources.
 */
export interface RegexMatchSetState {
    /**
     * The name or description of the Regex Match Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The regular expression pattern that you want AWS WAF to search for in web requests,
     * the location in requests that you want AWS WAF to search, and other settings. See below.
     */
    readonly regexMatchTuples?: pulumi.Input<pulumi.Input<inputs.wafregional.RegexMatchSetRegexMatchTuple>[]>;
}

/**
 * The set of arguments for constructing a RegexMatchSet resource.
 */
export interface RegexMatchSetArgs {
    /**
     * The name or description of the Regex Match Set.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The regular expression pattern that you want AWS WAF to search for in web requests,
     * the location in requests that you want AWS WAF to search, and other settings. See below.
     */
    readonly regexMatchTuples?: pulumi.Input<pulumi.Input<inputs.wafregional.RegexMatchSetRegexMatchTuple>[]>;
}
