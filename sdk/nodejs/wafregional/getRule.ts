// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enum";
import * as utilities from "../utilities";

/**
 * `aws.wafregional.Rule` Retrieves a WAF Regional Rule Resource Id.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.wafregional.getRule({
 *     name: "tfWAFRegionalRule",
 * }, { async: true }));
 * ```
 */
export function getRule(args: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:wafregional/getRule:getRule", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRule.
 */
export interface GetRuleArgs {
    /**
     * The name of the WAF Regional rule.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRule.
 */
export interface GetRuleResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
}
