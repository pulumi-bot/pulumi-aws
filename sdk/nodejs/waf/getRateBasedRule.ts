// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * `aws.waf.RateBasedRule` Retrieves a WAF Rate Based Rule Resource Id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.waf.getRateBasedRule({
 *     name: "tfWAFRateBasedRule",
 * }, { async: true }));
 * ```
 */
export function getRateBasedRule(args: GetRateBasedRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRateBasedRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:waf/getRateBasedRule:getRateBasedRule", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRateBasedRule.
 */
export interface GetRateBasedRuleArgs {
    /**
     * The name of the WAF rate based rule.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRateBasedRule.
 */
export interface GetRateBasedRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
