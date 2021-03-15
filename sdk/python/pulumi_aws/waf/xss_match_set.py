# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['XssMatchSetArgs', 'XssMatchSet']

@pulumi.input_type
class XssMatchSetArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]] = None):
        """
        The set of arguments for constructing a XssMatchSet resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if xss_match_tuples is not None:
            pulumi.set(__self__, "xss_match_tuples", xss_match_tuples)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the SizeConstraintSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]:
        """
        The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        return pulumi.get(self, "xss_match_tuples")

    @xss_match_tuples.setter
    def xss_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['XssMatchSetXssMatchTupleArgs']]]]):
        pulumi.set(self, "xss_match_tuples", value)


class XssMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[XssMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF XSS Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        xss_match_set = aws.waf.XssMatchSet("xssMatchSet", xss_match_tuples=[
            aws.waf.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.waf.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="URI",
                ),
                text_transformation="NONE",
            ),
            aws.waf.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.waf.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="QUERY_STRING",
                ),
                text_transformation="NONE",
            ),
        ])
        ```

        ## Import

        WAF XSS Match Set can be imported using their ID, e.g.

        ```sh
         $ pulumi import aws:waf/xssMatchSet:XssMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param XssMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a WAF XSS Match Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        xss_match_set = aws.waf.XssMatchSet("xssMatchSet", xss_match_tuples=[
            aws.waf.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.waf.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="URI",
                ),
                text_transformation="NONE",
            ),
            aws.waf.XssMatchSetXssMatchTupleArgs(
                field_to_match=aws.waf.XssMatchSetXssMatchTupleFieldToMatchArgs(
                    type="QUERY_STRING",
                ),
                text_transformation="NONE",
            ),
        ])
        ```

        ## Import

        WAF XSS Match Set can be imported using their ID, e.g.

        ```sh
         $ pulumi import aws:waf/xssMatchSet:XssMatchSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(XssMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['name'] = name
            __props__['xss_match_tuples'] = xss_match_tuples
            __props__['arn'] = None
        super(XssMatchSet, __self__).__init__(
            'aws:waf/xssMatchSet:XssMatchSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            xss_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]]] = None) -> 'XssMatchSet':
        """
        Get an existing XssMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] name: The name or description of the SizeConstraintSet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['XssMatchSetXssMatchTupleArgs']]]] xss_match_tuples: The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["xss_match_tuples"] = xss_match_tuples
        return XssMatchSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the SizeConstraintSet.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="xssMatchTuples")
    def xss_match_tuples(self) -> pulumi.Output[Optional[Sequence['outputs.XssMatchSetXssMatchTuple']]]:
        """
        The parts of web requests that you want to inspect for cross-site scripting attacks.
        """
        return pulumi.get(self, "xss_match_tuples")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

