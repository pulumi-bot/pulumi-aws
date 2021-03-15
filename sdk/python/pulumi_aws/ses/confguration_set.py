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

__all__ = ['ConfgurationSetArgs', 'ConfgurationSet']

@pulumi.input_type
class ConfgurationSetArgs:
    def __init__(__self__, *,
                 delivery_options: Optional[pulumi.Input['ConfgurationSetDeliveryOptionsArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ConfgurationSet resource.
        :param pulumi.Input['ConfgurationSetDeliveryOptionsArgs'] delivery_options: Configuration block. Detailed below.
        :param pulumi.Input[str] name: Name of the configuration set.
        """
        if delivery_options is not None:
            pulumi.set(__self__, "delivery_options", delivery_options)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="deliveryOptions")
    def delivery_options(self) -> Optional[pulumi.Input['ConfgurationSetDeliveryOptionsArgs']]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "delivery_options")

    @delivery_options.setter
    def delivery_options(self, value: Optional[pulumi.Input['ConfgurationSetDeliveryOptionsArgs']]):
        pulumi.set(self, "delivery_options", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the configuration set.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


warnings.warn("""aws.ses.ConfgurationSet has been deprecated in favor of aws.ses.ConfigurationSet""", DeprecationWarning)


class ConfgurationSet(pulumi.CustomResource):
    warnings.warn("""aws.ses.ConfgurationSet has been deprecated in favor of aws.ses.ConfigurationSet""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ConfgurationSetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an SES configuration set resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.ses.ConfigurationSet("test")
        ```
        ### Require TLS Connections

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.ses.ConfigurationSet("test", delivery_options=aws.ses.ConfigurationSetDeliveryOptionsArgs(
            tls_policy="Require",
        ))
        ```

        ## Import

        SES Configuration Sets can be imported using their `name`, e.g.

        ```sh
         $ pulumi import aws:ses/confgurationSet:ConfgurationSet test some-configuration-set-test
        ```

        :param str resource_name: The name of the resource.
        :param ConfgurationSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delivery_options: Optional[pulumi.Input[pulumi.InputType['ConfgurationSetDeliveryOptionsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an SES configuration set resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.ses.ConfigurationSet("test")
        ```
        ### Require TLS Connections

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.ses.ConfigurationSet("test", delivery_options=aws.ses.ConfigurationSetDeliveryOptionsArgs(
            tls_policy="Require",
        ))
        ```

        ## Import

        SES Configuration Sets can be imported using their `name`, e.g.

        ```sh
         $ pulumi import aws:ses/confgurationSet:ConfgurationSet test some-configuration-set-test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['ConfgurationSetDeliveryOptionsArgs']] delivery_options: Configuration block. Detailed below.
        :param pulumi.Input[str] name: Name of the configuration set.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConfgurationSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delivery_options: Optional[pulumi.Input[pulumi.InputType['ConfgurationSetDeliveryOptionsArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""ConfgurationSet is deprecated: aws.ses.ConfgurationSet has been deprecated in favor of aws.ses.ConfigurationSet""")
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

            __props__['delivery_options'] = delivery_options
            __props__['name'] = name
            __props__['arn'] = None
        super(ConfgurationSet, __self__).__init__(
            'aws:ses/confgurationSet:ConfgurationSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            delivery_options: Optional[pulumi.Input[pulumi.InputType['ConfgurationSetDeliveryOptionsArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ConfgurationSet':
        """
        Get an existing ConfgurationSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: SES configuration set ARN.
        :param pulumi.Input[pulumi.InputType['ConfgurationSetDeliveryOptionsArgs']] delivery_options: Configuration block. Detailed below.
        :param pulumi.Input[str] name: Name of the configuration set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["delivery_options"] = delivery_options
        __props__["name"] = name
        return ConfgurationSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        SES configuration set ARN.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="deliveryOptions")
    def delivery_options(self) -> pulumi.Output[Optional['outputs.ConfgurationSetDeliveryOptions']]:
        """
        Configuration block. Detailed below.
        """
        return pulumi.get(self, "delivery_options")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the configuration set.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

