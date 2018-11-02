# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class RateBasedRule(pulumi.CustomResource):
    """
    Provides a WAF Rate Based Rule Resource
    """
    def __init__(__self__, __name__, __opts__=None, metric_name=None, name=None, predicates=None, rate_key=None, rate_limit=None):
        """Create a RateBasedRule resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not metric_name:
            raise TypeError('Missing required property metric_name')
        __props__['metricName'] = metric_name

        __props__['name'] = name

        __props__['predicates'] = predicates

        if not rate_key:
            raise TypeError('Missing required property rate_key')
        __props__['rateKey'] = rate_key

        if not rate_limit:
            raise TypeError('Missing required property rate_limit')
        __props__['rateLimit'] = rate_limit

        super(RateBasedRule, __self__).__init__(
            'aws:wafregional/rateBasedRule:RateBasedRule',
            __name__,
            __props__,
            __opts__)

