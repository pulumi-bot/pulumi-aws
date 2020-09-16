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

__all__ = ['SizeConstraintSet']


class SizeConstraintSet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 size_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SizeConstraintSetSizeConstraintArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a WAF Size Constraint Set Resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        size_constraint_set = aws.waf.SizeConstraintSet("sizeConstraintSet", size_constraints=[aws.waf.SizeConstraintSetSizeConstraintArgs(
            comparison_operator="EQ",
            field_to_match=aws.waf.SizeConstraintSetSizeConstraintFieldToMatchArgs(
                type="BODY",
            ),
            size=4096,
            text_transformation="NONE",
        )])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name or description of the Size Constraint Set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SizeConstraintSetSizeConstraintArgs']]]] size_constraints: Specifies the parts of web requests that you want to inspect the size of.
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

            __props__['name'] = name
            __props__['size_constraints'] = size_constraints
            __props__['arn'] = None
        super(SizeConstraintSet, __self__).__init__(
            'aws:waf/sizeConstraintSet:SizeConstraintSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            size_constraints: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SizeConstraintSetSizeConstraintArgs']]]]] = None) -> 'SizeConstraintSet':
        """
        Get an existing SizeConstraintSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] name: The name or description of the Size Constraint Set.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SizeConstraintSetSizeConstraintArgs']]]] size_constraints: Specifies the parts of web requests that you want to inspect the size of.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["name"] = name
        __props__["size_constraints"] = size_constraints
        return SizeConstraintSet(resource_name, opts=opts, __props__=__props__)

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
        The name or description of the Size Constraint Set.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sizeConstraints")
    def size_constraints(self) -> pulumi.Output[Optional[Sequence['outputs.SizeConstraintSetSizeConstraint']]]:
        """
        Specifies the parts of web requests that you want to inspect the size of.
        """
        return pulumi.get(self, "size_constraints")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

