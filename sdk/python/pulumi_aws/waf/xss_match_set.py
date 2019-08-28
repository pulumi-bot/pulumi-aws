# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class XssMatchSet(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name or description of the SizeConstraintSet.
    """
    xss_match_tuples: pulumi.Output[list]
    """
    The parts of web requests that you want to inspect for cross-site scripting attacks.
    
      * `fieldToMatch` (`dict`) - Specifies where in a web request to look for cross-site scripting attacks.
    
        * `data` (`str`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
          If `type` is any other value, omit this field.
        * `type` (`str`) - The part of the web request that you want AWS WAF to search for a specified string.
          e.g. `HEADER`, `METHOD` or `BODY`.
          See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
          for all supported values.
    
      * `textTransformation` (`str`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
        e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
        for all supported values.
    """
    def __init__(__self__, resource_name, opts=None, name=None, xss_match_tuples=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a WAF XSS Match Set Resource
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[list] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        
        The **xss_match_tuples** object supports the following:
        
          * `fieldToMatch` (`pulumi.Input[dict]`) - Specifies where in a web request to look for cross-site scripting attacks.
        
            * `data` (`pulumi.Input[str]`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
              If `type` is any other value, omit this field.
            * `type` (`pulumi.Input[str]`) - The part of the web request that you want AWS WAF to search for a specified string.
              e.g. `HEADER`, `METHOD` or `BODY`.
              See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
              for all supported values.
        
          * `textTransformation` (`pulumi.Input[str]`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
            If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
            e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
            See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
            for all supported values.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_xss_match_set.html.markdown.
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

            __props__['name'] = name
            __props__['xss_match_tuples'] = xss_match_tuples
        super(XssMatchSet, __self__).__init__(
            'aws:waf/xssMatchSet:XssMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, xss_match_tuples=None):
        """
        Get an existing XssMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[list] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        
        The **xss_match_tuples** object supports the following:
        
          * `fieldToMatch` (`pulumi.Input[dict]`) - Specifies where in a web request to look for cross-site scripting attacks.
        
            * `data` (`pulumi.Input[str]`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
              If `type` is any other value, omit this field.
            * `type` (`pulumi.Input[str]`) - The part of the web request that you want AWS WAF to search for a specified string.
              e.g. `HEADER`, `METHOD` or `BODY`.
              See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
              for all supported values.
        
          * `textTransformation` (`pulumi.Input[str]`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
            If you specify a transformation, AWS WAF performs the transformation on `target_string` before inspecting a request for a match.
            e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
            See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_XssMatchTuple.html#WAF-Type-XssMatchTuple-TextTransformation)
            for all supported values.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/waf_xss_match_set.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["name"] = name
        __props__["xss_match_tuples"] = xss_match_tuples
        return XssMatchSet(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

