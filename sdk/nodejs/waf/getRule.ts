// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * `aws.waf.Rule` Retrieves a WAF Rule Resource Id.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.waf.getRule({
 *     name: "tfWAFRule",
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
    return pulumi.runtime.invoke("aws:waf/getRule:getRule", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRule.
 */
export interface GetRuleArgs {
    /**
     * The name of the WAF rule.
     */
    readonly name: string;
}

/**
 * A collection of values returned by getRule.
 */
export interface GetRuleResult {
    readonly name: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
