# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SizeConstraintSet(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN)
    """
    name: pulumi.Output[str]
    """
    The name or description of the Size Constraint Set.
    """
    size_constraints: pulumi.Output[list]
    """
    Specifies the parts of web requests that you want to inspect the size of.

      * `comparison_operator` (`str`) - The type of comparison you want to perform.
        e.g. `EQ`, `NE`, `LT`, `GT`.
        See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_SizeConstraint.html) for all supported values.
      * `fieldToMatch` (`dict`) - Specifies where in a web request to look for the size constraint.
        * `data` (`str`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
          If `type` is any other value, omit this field.
        * `type` (`str`) - The part of the web request that you want AWS WAF to search for a specified string.
          e.g. `HEADER`, `METHOD` or `BODY`.
          See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
          for all supported values.

      * `size` (`float`) - The size in bytes that you want to compare against the size of the specified `field_to_match`.
        Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
      * `textTransformation` (`str`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
        If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
        e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
        See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
        for all supported values.
        **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
    """
    def __init__(__self__, resource_name, opts=None, name=None, size_constraints=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a WAF Size Constraint Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        size_constraint_set = aws.waf.SizeConstraintSet("sizeConstraintSet", size_constraints=[{
            "comparison_operator": "EQ",
            "fieldToMatch": {
                "type": "BODY",
            },
            "size": "4096",
            "textTransformation": "NONE",
        }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the Size Constraint Set.
        :param pulumi.Input[list] size_constraints: Specifies the parts of web requests that you want to inspect the size of.

        The **size_constraints** object supports the following:

          * `comparison_operator` (`pulumi.Input[str]`) - The type of comparison you want to perform.
            e.g. `EQ`, `NE`, `LT`, `GT`.
            See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_SizeConstraint.html) for all supported values.
          * `fieldToMatch` (`pulumi.Input[dict]`) - Specifies where in a web request to look for the size constraint.
            * `data` (`pulumi.Input[str]`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
              If `type` is any other value, omit this field.
            * `type` (`pulumi.Input[str]`) - The part of the web request that you want AWS WAF to search for a specified string.
              e.g. `HEADER`, `METHOD` or `BODY`.
              See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
              for all supported values.

          * `size` (`pulumi.Input[float]`) - The size in bytes that you want to compare against the size of the specified `field_to_match`.
            Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
          * `textTransformation` (`pulumi.Input[str]`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
            If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
            e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
            See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
            for all supported values.
            **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
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
            __props__['size_constraints'] = size_constraints
            __props__['arn'] = None
        super(SizeConstraintSet, __self__).__init__(
            'aws:waf/sizeConstraintSet:SizeConstraintSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, name=None, size_constraints=None):
        """
        Get an existing SizeConstraintSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] name: The name or description of the Size Constraint Set.
        :param pulumi.Input[list] size_constraints: Specifies the parts of web requests that you want to inspect the size of.

        The **size_constraints** object supports the following:

          * `comparison_operator` (`pulumi.Input[str]`) - The type of comparison you want to perform.
            e.g. `EQ`, `NE`, `LT`, `GT`.
            See [docs](https://docs.aws.amazon.com/waf/latest/APIReference/API_wafRegional_SizeConstraint.html) for all supported values.
          * `fieldToMatch` (`pulumi.Input[dict]`) - Specifies where in a web request to look for the size constraint.
            * `data` (`pulumi.Input[str]`) - When `type` is `HEADER`, enter the name of the header that you want to search, e.g. `User-Agent` or `Referer`.
              If `type` is any other value, omit this field.
            * `type` (`pulumi.Input[str]`) - The part of the web request that you want AWS WAF to search for a specified string.
              e.g. `HEADER`, `METHOD` or `BODY`.
              See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_FieldToMatch.html)
              for all supported values.

          * `size` (`pulumi.Input[float]`) - The size in bytes that you want to compare against the size of the specified `field_to_match`.
            Valid values are between 0 - 21474836480 bytes (0 - 20 GB).
          * `textTransformation` (`pulumi.Input[str]`) - Text transformations used to eliminate unusual formatting that attackers use in web requests in an effort to bypass AWS WAF.
            If you specify a transformation, AWS WAF performs the transformation on `field_to_match` before inspecting a request for a match.
            e.g. `CMD_LINE`, `HTML_ENTITY_DECODE` or `NONE`.
            See [docs](http://docs.aws.amazon.com/waf/latest/APIReference/API_SizeConstraint.html#WAF-Type-SizeConstraint-TextTransformation)
            for all supported values.
            **Note:** if you choose `BODY` as `type`, you must choose `NONE` because CloudFront forwards only the first 8192 bytes for inspection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["size_constraints"] = size_constraints
        return SizeConstraintSet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
