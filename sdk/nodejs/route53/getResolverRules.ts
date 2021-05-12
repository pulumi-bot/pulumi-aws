// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * `aws.route53.getResolverRules` provides details about a set of Route53 Resolver rules.
 *
 * ## Example Usage
 *
 * Retrieving the default resolver rule.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.route53.getResolverRules({
 *     ownerId: "Route 53 Resolver",
 *     ruleType: "RECURSIVE",
 *     shareStatus: "NOT_SHARED",
 * }, { async: true }));
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.route53.getResolverRules({
 *     ruleType: "FORWARD",
 *     shareStatus: "SHARED_WITH_ME",
 * }, { async: true }));
 * ```
 */
export function getResolverRules(args?: GetResolverRulesArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverRulesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:route53/getResolverRules:getResolverRules", {
        "ownerId": args.ownerId,
        "resolverEndpointId": args.resolverEndpointId,
        "ruleType": args.ruleType,
        "shareStatus": args.shareStatus,
    }, opts);
}

/**
 * A collection of arguments for invoking getResolverRules.
 */
export interface GetResolverRulesArgs {
    /**
     * When the desired resolver rules are shared with another AWS account, the account ID of the account that the rules are shared with.
     */
    ownerId?: string;
    /**
     * The ID of the outbound resolver endpoint for the desired resolver rules.
     */
    resolverEndpointId?: string;
    /**
     * The rule type of the desired resolver rules. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`.
     */
    ruleType?: string;
    /**
     * Whether the desired resolver rules are shared and, if so, whether the current account is sharing the rules with another account, or another account is sharing the rules with the current account. Valid values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
     */
    shareStatus?: string;
}

/**
 * A collection of values returned by getResolverRules.
 */
export interface GetResolverRulesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ownerId?: string;
    readonly resolverEndpointId?: string;
    /**
     * The IDs of the matched resolver rules.
     */
    readonly resolverRuleIds: string[];
    readonly ruleType?: string;
    readonly shareStatus?: string;
}
