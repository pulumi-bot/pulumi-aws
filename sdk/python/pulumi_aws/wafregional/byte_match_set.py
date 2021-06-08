# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ByteMatchSetArgs', 'ByteMatchSet']

@pulumi.input_type
class ByteMatchSetArgs:
    def __init__(__self__, *,
                 byte_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ByteMatchSet resource.
        :param pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]] byte_match_tuples: Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        :param pulumi.Input[str] name: The name or description of the ByteMatchSet.
        """
        if byte_match_tuples is not None:
            pulumi.set(__self__, "byte_match_tuples", byte_match_tuples)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="byteMatchTuples")
    def byte_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]]:
        """
        Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        """
        return pulumi.get(self, "byte_match_tuples")

    @byte_match_tuples.setter
    def byte_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]]):
        pulumi.set(self, "byte_match_tuples", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the ByteMatchSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ByteMatchSetState:
    def __init__(__self__, *,
                 byte_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ByteMatchSet resources.
        :param pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]] byte_match_tuples: Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        :param pulumi.Input[str] name: The name or description of the ByteMatchSet.
        """
        if byte_match_tuples is not None:
            pulumi.set(__self__, "byte_match_tuples", byte_match_tuples)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="byteMatchTuples")
    def byte_match_tuples(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]]:
        """
        Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        """
        return pulumi.get(self, "byte_match_tuples")

    @byte_match_tuples.setter
    def byte_match_tuples(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ByteMatchSetByteMatchTupleArgs']]]]):
        pulumi.set(self, "byte_match_tuples", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name or description of the ByteMatchSet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ByteMatchSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 byte_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ByteMatchSetByteMatchTupleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a WAF Regional Byte Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        byte_set = aws.wafregional.ByteMatchSet("byteSet", byte_match_tuples=[aws.wafregional.ByteMatchSetByteMatchTupleArgs(
            field_to_match=aws.wafregional.ByteMatchSetByteMatchTupleFieldToMatchArgs(
                data="referer",
                type="HEADER",
            ),
            positional_constraint="CONTAINS",
            target_string="badrefer1",
            text_transformation="NONE",
        )])
        ```

        ## Import

        WAF Regional Byte Match Set can be imported using the id, e.g.

        ```sh
         $ pulumi import aws:wafregional/byteMatchSet:ByteMatchSet byte_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ByteMatchSetByteMatchTupleArgs']]]] byte_match_tuples: Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        :param pulumi.Input[str] name: The name or description of the ByteMatchSet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: Optional[ByteMatchSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a WAF Regional Byte Match Set Resource for use with Application Load Balancer.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        byte_set = aws.wafregional.ByteMatchSet("byteSet", byte_match_tuples=[aws.wafregional.ByteMatchSetByteMatchTupleArgs(
            field_to_match=aws.wafregional.ByteMatchSetByteMatchTupleFieldToMatchArgs(
                data="referer",
                type="HEADER",
            ),
            positional_constraint="CONTAINS",
            target_string="badrefer1",
            text_transformation="NONE",
        )])
        ```

        ## Import

        WAF Regional Byte Match Set can be imported using the id, e.g.

        ```sh
         $ pulumi import aws:wafregional/byteMatchSet:ByteMatchSet byte_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
        ```

        :param str resource_name_: The name of the resource.
        :param ByteMatchSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ByteMatchSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 byte_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ByteMatchSetByteMatchTupleArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ByteMatchSetArgs.__new__(ByteMatchSetArgs)

            __props__.__dict__["byte_match_tuples"] = byte_match_tuples
            __props__.__dict__["name"] = name
        super(ByteMatchSet, __self__).__init__(
            'aws:wafregional/byteMatchSet:ByteMatchSet',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            byte_match_tuples: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ByteMatchSetByteMatchTupleArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ByteMatchSet':
        """
        Get an existing ByteMatchSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ByteMatchSetByteMatchTupleArgs']]]] byte_match_tuples: Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        :param pulumi.Input[str] name: The name or description of the ByteMatchSet.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ByteMatchSetState.__new__(_ByteMatchSetState)

        __props__.__dict__["byte_match_tuples"] = byte_match_tuples
        __props__.__dict__["name"] = name
        return ByteMatchSet(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="byteMatchTuples")
    def byte_match_tuples(self) -> pulumi.Output[Optional[Sequence['outputs.ByteMatchSetByteMatchTuple']]]:
        """
        Settings for the ByteMatchSet, such as the bytes (typically a string that corresponds with ASCII characters) that you want AWS WAF to search for in web requests. ByteMatchTuple documented below.
        """
        return pulumi.get(self, "byte_match_tuples")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name or description of the ByteMatchSet.
        """
        return pulumi.get(self, "name")

