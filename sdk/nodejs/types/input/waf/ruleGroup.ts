// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface RuleGroupActivatedRule {
    /**
     * Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
     */
    action: pulumi.Input<inputApi.waf.RuleGroupActivatedRuleAction>;
    /**
     * Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
     */
    priority: pulumi.Input<number>;
    /**
     * The ID of a [rule](https://www.terraform.io/docs/providers/aws/r/waf_rule.html)
     */
    ruleId: pulumi.Input<string>;
    /**
     * The rule type, either [`REGULAR`](https://www.terraform.io/docs/providers/aws/r/waf_rule.html), [`RATE_BASED`](https://www.terraform.io/docs/providers/aws/r/waf_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.
     */
    type?: pulumi.Input<string>;
}

export interface RuleGroupActivatedRuleAction {
    /**
     * The rule type, either [`REGULAR`](https://www.terraform.io/docs/providers/aws/r/waf_rule.html), [`RATE_BASED`](https://www.terraform.io/docs/providers/aws/r/waf_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.
     */
    type: pulumi.Input<string>;
}
