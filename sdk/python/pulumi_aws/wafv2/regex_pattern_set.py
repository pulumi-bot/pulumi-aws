# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RegexPatternSet']


class RegexPatternSet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AWS WAFv2 Regex Pattern Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.RegexPatternSet("example",
            description="Example regex pattern set",
            regular_expressions=[
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="one",
                ),
                aws.wafv2.RegexPatternSetRegularExpressionArgs(
                    regex_string="two",
                ),
            ],
            scope="REGIONAL",
            tags={
                "Tag1": "Value1",
                "Tag2": "Value2",
            })
        ```

        ## Import

        WAFv2 Regex Pattern Sets can be imported using `ID/name/scope` e.g.

        ```sh
         $ pulumi import aws:wafv2/regexPatternSet:RegexPatternSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            __props__['name'] = name
            __props__['regular_expressions'] = regular_expressions
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['lock_token'] = None
        super(RegexPatternSet, __self__).__init__(
            'aws:wafv2/regexPatternSet:RegexPatternSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            lock_token: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            regular_expressions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]]] = None,
            scope: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'RegexPatternSet':
        """
        Get an existing RegexPatternSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) that identifies the cluster.
        :param pulumi.Input[str] description: A friendly description of the regular expression pattern set.
        :param pulumi.Input[str] name: A friendly name of the regular expression pattern set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RegexPatternSetRegularExpressionArgs']]]] regular_expressions: One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        :param pulumi.Input[str] scope: Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: An array of key:value pairs to associate with the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["lock_token"] = lock_token
        __props__["name"] = name
        __props__["regular_expressions"] = regular_expressions
        __props__["scope"] = scope
        __props__["tags"] = tags
        return RegexPatternSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) that identifies the cluster.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A friendly description of the regular expression pattern set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lockToken")
    def lock_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lock_token")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A friendly name of the regular expression pattern set.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="regularExpressions")
    def regular_expressions(self) -> pulumi.Output[Optional[Sequence['outputs.RegexPatternSetRegularExpression']]]:
        """
        One or more blocks of regular expression patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`. See Regular Expression below for details.
        """
        return pulumi.get(self, "regular_expressions")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Specifies whether this is for an AWS CloudFront distribution or for a regional application. Valid values are `CLOUDFRONT` or `REGIONAL`. To work with CloudFront, you must also specify the region `us-east-1` (N. Virginia) on the AWS provider.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        An array of key:value pairs to associate with the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

