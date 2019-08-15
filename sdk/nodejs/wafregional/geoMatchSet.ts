// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../types/input";
import * as outputApi from "../types/output";
import * as utilities from "../utilities";

/**
 * Provides a WAF Regional Geo Match Set Resource
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * 
 * const geoMatchSet = new aws.wafregional.GeoMatchSet("geoMatchSet", {
 *     geoMatchConstraints: [
 *         {
 *             type: "Country",
 *             value: "US",
 *         },
 *         {
 *             type: "Country",
 *             value: "CA",
 *         },
 *     ],
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/wafregional_geo_match_set.html.markdown.
 */
export class GeoMatchSet extends pulumi.CustomResource {
    /**
     * Get an existing GeoMatchSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GeoMatchSetState, opts?: pulumi.CustomResourceOptions): GeoMatchSet {
        return new GeoMatchSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:wafregional/geoMatchSet:GeoMatchSet';

    /**
     * Returns true if the given object is an instance of GeoMatchSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GeoMatchSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GeoMatchSet.__pulumiType;
    }

    /**
     * The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
     */
    public readonly geoMatchConstraints!: pulumi.Output<outputApi.wafregional.GeoMatchSetGeoMatchConstraint[] | undefined>;
    /**
     * The name or description of the Geo Match Set.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a GeoMatchSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GeoMatchSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GeoMatchSetArgs | GeoMatchSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as GeoMatchSetState | undefined;
            inputs["geoMatchConstraints"] = state ? state.geoMatchConstraints : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as GeoMatchSetArgs | undefined;
            inputs["geoMatchConstraints"] = args ? args.geoMatchConstraints : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(GeoMatchSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GeoMatchSet resources.
 */
export interface GeoMatchSetState {
    /**
     * The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
     */
    readonly geoMatchConstraints?: pulumi.Input<pulumi.Input<inputApi.wafregional.GeoMatchSetGeoMatchConstraint>[]>;
    /**
     * The name or description of the Geo Match Set.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GeoMatchSet resource.
 */
export interface GeoMatchSetArgs {
    /**
     * The Geo Match Constraint objects which contain the country that you want AWS WAF to search for.
     */
    readonly geoMatchConstraints?: pulumi.Input<pulumi.Input<inputApi.wafregional.GeoMatchSetGeoMatchConstraint>[]>;
    /**
     * The name or description of the Geo Match Set.
     */
    readonly name?: pulumi.Input<string>;
}
