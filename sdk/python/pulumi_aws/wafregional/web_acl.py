# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class WebAcl(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the WAF Regional WebACL.
    """
    default_action: pulumi.Output[dict]
    """
    The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
    """
    logging_configuration: pulumi.Output[dict]
    """
    Configuration block to enable WAF logging. Detailed below.
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
    Set of configuration blocks containing rules for the web ACL. Detailed below.
    """
    def __init__(__self__, resource_name, opts=None, default_action=None, logging_configuration=None, metric_name=None, name=None, rules=None, __name__=None, __opts__=None):
        """
        Provides a WAF Regional Web ACL Resource for use with Application Load Balancer.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] default_action: The action that you want AWS WAF Regional to take when a request doesn't match the criteria in any of the rules that are associated with the web ACL.
        :param pulumi.Input[dict] logging_configuration: Configuration block to enable WAF logging. Detailed below.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this web ACL.
        :param pulumi.Input[str] name: The name or description of the web ACL.
        :param pulumi.Input[list] rules: Set of configuration blocks containing rules for the web ACL. Detailed below.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if default_action is None:
            raise TypeError("Missing required property 'default_action'")
        __props__['default_action'] = default_action

        __props__['logging_configuration'] = logging_configuration

        if metric_name is None:
            raise TypeError("Missing required property 'metric_name'")
        __props__['metric_name'] = metric_name

        __props__['name'] = name

        __props__['rules'] = rules

        __props__['arn'] = None

        super(WebAcl, __self__).__init__(
            'aws:wafregional/webAcl:WebAcl',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

