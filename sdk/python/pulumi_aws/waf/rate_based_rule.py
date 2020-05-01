# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class RateBasedRule(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN)
    """
    metric_name: pulumi.Output[str]
    """
    The name or description for the Amazon CloudWatch metric of this rule.
    """
    name: pulumi.Output[str]
    """
    The name or description of the rule.
    """
    predicates: pulumi.Output[list]
    """
    The objects to include in a rule (documented below).

      * `dataId` (`str`) - A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
      * `negated` (`bool`) - Set this to `false` if you want to allow, block, or count requests
        based on the settings in the specified `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, or `SizeConstraintSet`.
        For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
        If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
      * `type` (`str`) - The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
    """
    rate_key: pulumi.Output[str]
    """
    Valid value is IP.
    """
    rate_limit: pulumi.Output[float]
    """
    The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
    """
    tags: pulumi.Output[dict]
    """
    Key-value mapping of resource tags
    """
    def __init__(__self__, resource_name, opts=None, metric_name=None, name=None, predicates=None, rate_key=None, rate_limit=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a WAF Rate Based Rule Resource

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        ipset = aws.waf.IpSet("ipset", ip_set_descriptors=[{
            "type": "IPV4",
            "value": "192.0.7.0/24",
        }])
        wafrule = aws.waf.RateBasedRule("wafrule",
            metric_name="tfWAFRule",
            predicates=[{
                "dataId": ipset.id,
                "negated": False,
                "type": "IPMatch",
            }],
            rate_key="IP",
            rate_limit=100)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[list] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[float] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **predicates** object supports the following:

          * `dataId` (`pulumi.Input[str]`) - A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
          * `negated` (`pulumi.Input[bool]`) - Set this to `false` if you want to allow, block, or count requests
            based on the settings in the specified `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, or `SizeConstraintSet`.
            For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
            If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
          * `type` (`pulumi.Input[str]`) - The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if metric_name is None:
                raise TypeError("Missing required property 'metric_name'")
            __props__['metric_name'] = metric_name
            __props__['name'] = name
            __props__['predicates'] = predicates
            if rate_key is None:
                raise TypeError("Missing required property 'rate_key'")
            __props__['rate_key'] = rate_key
            if rate_limit is None:
                raise TypeError("Missing required property 'rate_limit'")
            __props__['rate_limit'] = rate_limit
            __props__['tags'] = tags
            __props__['arn'] = None
        super(RateBasedRule, __self__).__init__(
            'aws:waf/rateBasedRule:RateBasedRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, metric_name=None, name=None, predicates=None, rate_key=None, rate_limit=None, tags=None):
        """
        Get an existing RateBasedRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] metric_name: The name or description for the Amazon CloudWatch metric of this rule.
        :param pulumi.Input[str] name: The name or description of the rule.
        :param pulumi.Input[list] predicates: The objects to include in a rule (documented below).
        :param pulumi.Input[str] rate_key: Valid value is IP.
        :param pulumi.Input[float] rate_limit: The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
        :param pulumi.Input[dict] tags: Key-value mapping of resource tags

        The **predicates** object supports the following:

          * `dataId` (`pulumi.Input[str]`) - A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
          * `negated` (`pulumi.Input[bool]`) - Set this to `false` if you want to allow, block, or count requests
            based on the settings in the specified `ByteMatchSet`, `IPSet`, `SqlInjectionMatchSet`, `XssMatchSet`, or `SizeConstraintSet`.
            For example, if an IPSet includes the IP address `192.0.2.44`, AWS WAF will allow or block requests based on that IP address.
            If set to `true`, AWS WAF will allow, block, or count requests based on all IP addresses _except_ `192.0.2.44`.
          * `type` (`pulumi.Input[str]`) - The type of predicate in a rule. Valid values: `ByteMatch`, `GeoMatch`, `IPMatch`, `RegexMatch`, `SizeConstraint`, `SqlInjectionMatch`, or `XssMatch`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["metric_name"] = metric_name
        __props__["name"] = name
        __props__["predicates"] = predicates
        __props__["rate_key"] = rate_key
        __props__["rate_limit"] = rate_limit
        __props__["tags"] = tags
        return RateBasedRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

