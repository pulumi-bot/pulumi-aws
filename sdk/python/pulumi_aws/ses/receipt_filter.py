# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ReceiptFilterArgs', 'ReceiptFilter']

@pulumi.input_type
class ReceiptFilterArgs:
    def __init__(__self__, *,
                 cidr: pulumi.Input[str],
                 policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ReceiptFilter resource.
        :param pulumi.Input[str] cidr: The IP address or address range to filter, in CIDR notation
        :param pulumi.Input[str] policy: Block or Allow
        :param pulumi.Input[str] name: The name of the filter
        """
        pulumi.set(__self__, "cidr", cidr)
        pulumi.set(__self__, "policy", policy)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Input[str]:
        """
        The IP address or address range to filter, in CIDR notation
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        Block or Allow
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the filter
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ReceiptFilterState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ReceiptFilter resources.
        :param pulumi.Input[str] arn: The SES receipt filter ARN.
        :param pulumi.Input[str] cidr: The IP address or address range to filter, in CIDR notation
        :param pulumi.Input[str] name: The name of the filter
        :param pulumi.Input[str] policy: Block or Allow
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if cidr is not None:
            pulumi.set(__self__, "cidr", cidr)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The SES receipt filter ARN.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address or address range to filter, in CIDR notation
        """
        return pulumi.get(self, "cidr")

    @cidr.setter
    def cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cidr", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the filter
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Block or Allow
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class ReceiptFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides an SES receipt filter resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        filter = aws.ses.ReceiptFilter("filter",
            cidr="10.10.10.10",
            policy="Block")
        ```

        ## Import

        SES Receipt Filter can be imported using their `name`, e.g.

        ```sh
         $ pulumi import aws:ses/receiptFilter:ReceiptFilter test some-filter
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cidr: The IP address or address range to filter, in CIDR notation
        :param pulumi.Input[str] name: The name of the filter
        :param pulumi.Input[str] policy: Block or Allow
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: ReceiptFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an SES receipt filter resource

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        filter = aws.ses.ReceiptFilter("filter",
            cidr="10.10.10.10",
            policy="Block")
        ```

        ## Import

        SES Receipt Filter can be imported using their `name`, e.g.

        ```sh
         $ pulumi import aws:ses/receiptFilter:ReceiptFilter test some-filter
        ```

        :param str resource_name_: The name of the resource.
        :param ReceiptFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReceiptFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
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
            __props__ = ReceiptFilterArgs.__new__(ReceiptFilterArgs)

            if cidr is None and not opts.urn:
                raise TypeError("Missing required property 'cidr'")
            __props__.__dict__["cidr"] = cidr
            __props__.__dict__["name"] = name
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
            __props__.__dict__["arn"] = None
        super(ReceiptFilter, __self__).__init__(
            'aws:ses/receiptFilter:ReceiptFilter',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            cidr: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'ReceiptFilter':
        """
        Get an existing ReceiptFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The SES receipt filter ARN.
        :param pulumi.Input[str] cidr: The IP address or address range to filter, in CIDR notation
        :param pulumi.Input[str] name: The name of the filter
        :param pulumi.Input[str] policy: Block or Allow
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ReceiptFilterState.__new__(_ReceiptFilterState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["cidr"] = cidr
        __props__.__dict__["name"] = name
        __props__.__dict__["policy"] = policy
        return ReceiptFilter(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The SES receipt filter ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def cidr(self) -> pulumi.Output[str]:
        """
        The IP address or address range to filter, in CIDR notation
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the filter
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        Block or Allow
        """
        return pulumi.get(self, "policy")

