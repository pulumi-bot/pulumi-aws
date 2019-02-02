# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class RegexMatchSet(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name or description of the Regex Match Set.
    """
    regex_match_tuples: pulumi.Output[list]
    """
    The regular expression pattern that you want AWS WAF to search for in web requests,
    the location in requests that you want AWS WAF to search, and other settings. See below.
    """
    def __init__(__self__, __name__, __opts__=None, name=None, regex_match_tuples=None):
        """
        Provides a WAF Regional Regex Match Set Resource
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the Regex Match Set.
        :param pulumi.Input[list] regex_match_tuples: The regular expression pattern that you want AWS WAF to search for in web requests,
               the location in requests that you want AWS WAF to search, and other settings. See below.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['name'] = name

        __props__['regex_match_tuples'] = regex_match_tuples

        super(RegexMatchSet, __self__).__init__(
            'aws:wafregional/regexMatchSet:RegexMatchSet',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

