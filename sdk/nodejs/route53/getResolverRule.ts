// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * `aws.route53.ResolverRule` provides details about a specific Route53 Resolver rule.
 *
 * ## Example Usage
 *
 * The following example shows how to get a Route53 Resolver rule based on its associated domain name and rule type.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = pulumi.output(aws.route53.getResolverRule({
 *     domainName: "subdomain.example.com",
 *     ruleType: "SYSTEM",
 * }));
 * ```
 */
export function getResolverRule(args?: GetResolverRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetResolverRuleResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("aws:route53/getResolverRule:getResolverRule", {
        "domainName": args.domainName,
        "name": args.name,
        "resolverEndpointId": args.resolverEndpointId,
        "resolverRuleId": args.resolverRuleId,
        "ruleType": args.ruleType,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getResolverRule.
 */
export interface GetResolverRuleArgs {
    /**
     * The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolverRuleId`.
     */
    domainName?: string;
    /**
     * The friendly name of the desired resolver rule. Conflicts with `resolverRuleId`.
     */
    name?: string;
    /**
     * The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolverRuleId`.
     */
    resolverEndpointId?: string;
    /**
     * The ID of the desired resolver rule. Conflicts with `domainName`, `name`, `resolverEndpointId` and `ruleType`.
     */
    resolverRuleId?: string;
    /**
     * The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolverRuleId`.
     */
    ruleType?: string;
    /**
     * A map of tags assigned to the resolver rule.
     */
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getResolverRule.
 */
export interface GetResolverRuleResult {
    /**
     * The ARN (Amazon Resource Name) for the resolver rule.
     */
    readonly arn: string;
    readonly domainName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * When a rule is shared with another AWS account, the account ID of the account that the rule is shared with.
     */
    readonly ownerId: string;
    readonly resolverEndpointId: string;
    readonly resolverRuleId: string;
    readonly ruleType: string;
    /**
     * Whether the rules is shared and, if so, whether the current account is sharing the rule with another account, or another account is sharing the rule with the current account.
     * Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
     */
    readonly shareStatus: string;
    /**
     * A map of tags assigned to the resolver rule.
     */
    readonly tags: {[key: string]: string};
}

export function getResolverRuleApply(args?: GetResolverRuleApplyArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResolverRuleResult> {
    return pulumi.output(args).apply(a => getResolverRule(a, opts))
}

/**
 * A collection of arguments for invoking getResolverRule.
 */
export interface GetResolverRuleApplyArgs {
    /**
     * The domain name the desired resolver rule forwards DNS queries for. Conflicts with `resolverRuleId`.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The friendly name of the desired resolver rule. Conflicts with `resolverRuleId`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the outbound resolver endpoint of the desired resolver rule. Conflicts with `resolverRuleId`.
     */
    resolverEndpointId?: pulumi.Input<string>;
    /**
     * The ID of the desired resolver rule. Conflicts with `domainName`, `name`, `resolverEndpointId` and `ruleType`.
     */
    resolverRuleId?: pulumi.Input<string>;
    /**
     * The rule type of the desired resolver rule. Valid values are `FORWARD`, `SYSTEM` and `RECURSIVE`. Conflicts with `resolverRuleId`.
     */
    ruleType?: pulumi.Input<string>;
    /**
     * A map of tags assigned to the resolver rule.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
