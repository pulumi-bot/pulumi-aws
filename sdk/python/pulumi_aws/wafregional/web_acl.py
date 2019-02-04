# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class WebAcl(pulumi.CustomResource):
    default_action: pulumi.Output[dict]
    """
    The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
    """
    metric_name: pulumi.Output[str]
    """
    The name or description for the Amazon CloudWatch metric of this web ACL.
    """
    name: pulumi.Output[str]
    """
    The name or description of the web ACL.
    """
    rules: pulumi.Output[list]
    """
    The rules to associate with the web ACL and the settings for each rule.
    """
    def __init__(__self__, __name__, __opts__=None, default_action=None, metric_name=None, name=None, rules=None):
        """
        Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
        
        ## Nested Fields
        
        ### `rule`
        
        See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_regional_ActivatedRule.html) for all details and supported values.
        
        #### Arguments
        
        * `action` - (Required) The action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Not used if `type` is `GROUP`.
        * `override_action` - (Required) Override the action that a group requests CloudFront or AWS WAF takes when a web request matches the conditions in the rule.  Only used if `type` is `GROUP`.
        * `priority` - (Required) Specifies the order in which the rules in a WebACL are evaluated.
          Rules with a lower value are evaluated before rules with a higher value.
        * `rule_id` - (Required) ID of the associated WAF (Regional) rule (e.g. [`aws_wafregional_rule`](https://www.terraform.io/docs/providers/aws/r/wafregional_rule.html)). WAF (Global) rules cannot be used.
        * `type` - (Optional) The rule type, either `REGULAR`, as defined by [Rule](http://docs.aws.amazon.com/waf/latest/APIReference/API_Rule.html), `RATE_BASED`, as defined by [RateBasedRule](http://docs.aws.amazon.com/waf/latest/APIReference/API_RateBasedRule.html), or `GROUP`, as defined by [RuleGroup](https://docs.aws.amazon.com/waf/latest/APIReference/API_RuleGroup.html). The default is REGULAR. If you add a RATE_BASED rule, you need to set `type` as `RATE_BASED`. If you add a GROUP rule, you need to set `type` as `GROUP`.
        
        ### `default_action` / `action`
        
        #### Arguments
        
        * `type` - (Required) Specifies how you want AWS WAF Regional to respond to requests that match the settings in a rule.
          e.g. `ALLOW`, `BLOCK` or `COUNT`
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[list] rules: The rules to associate with the web ACL and the settings for each rule.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if default_action is None:
            raise TypeError('Missing required property default_action')
        __props__['default_action'] = default_action

        if metric_name is None:
            raise TypeError('Missing required property metric_name')
        __props__['metric_name'] = metric_name

        __props__['name'] = name

        __props__['rules'] = rules

        super(WebAcl, __self__).__init__(
            'aws:wafregional/webAcl:WebAcl',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

