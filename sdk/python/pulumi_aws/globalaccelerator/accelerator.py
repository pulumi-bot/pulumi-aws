# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Accelerator(pulumi.CustomResource):
    attributes: pulumi.Output[dict]
    """
    The attributes of the accelerator. Fields documented below.

      * `flowLogsEnabled` (`bool`) - Indicates whether flow logs are enabled.
      * `flowLogsS3Bucket` (`str`) - The name of the Amazon S3 bucket for the flow logs.
      * `flowLogsS3Prefix` (`str`) - The prefix for the location in the Amazon S3 bucket for the flow logs.
    """
    enabled: pulumi.Output[bool]
    """
    Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
    """
    ip_address_type: pulumi.Output[str]
    """
    The value for the address type must be `IPV4`.
    """
    ip_sets: pulumi.Output[list]
    """
    IP address set associated with the accelerator.

      * `ip_addresses` (`list`) - The array of IP addresses in the IP address set.
      * `ipFamily` (`str`) - The types of IP addresses included in this IP set.
    """
    name: pulumi.Output[str]
    """
    The name of the accelerator.
    """
    def __init__(__self__, resource_name, opts=None, attributes=None, enabled=None, ip_address_type=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Global Accelerator accelerator.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/globalaccelerator_accelerator.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attributes: The attributes of the accelerator. Fields documented below.
        :param pulumi.Input[bool] enabled: Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
        :param pulumi.Input[str] ip_address_type: The value for the address type must be `IPV4`.
        :param pulumi.Input[str] name: The name of the accelerator.

        The **attributes** object supports the following:

          * `flowLogsEnabled` (`pulumi.Input[bool]`) - Indicates whether flow logs are enabled.
          * `flowLogsS3Bucket` (`pulumi.Input[str]`) - The name of the Amazon S3 bucket for the flow logs.
          * `flowLogsS3Prefix` (`pulumi.Input[str]`) - The prefix for the location in the Amazon S3 bucket for the flow logs.
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

            __props__['attributes'] = attributes
            __props__['enabled'] = enabled
            __props__['ip_address_type'] = ip_address_type
            __props__['name'] = name
            __props__['ip_sets'] = None
        super(Accelerator, __self__).__init__(
            'aws:globalaccelerator/accelerator:Accelerator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attributes=None, enabled=None, ip_address_type=None, ip_sets=None, name=None):
        """
        Get an existing Accelerator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attributes: The attributes of the accelerator. Fields documented below.
        :param pulumi.Input[bool] enabled: Indicates whether the accelerator is enabled. The value is true or false. The default value is true.
        :param pulumi.Input[str] ip_address_type: The value for the address type must be `IPV4`.
        :param pulumi.Input[list] ip_sets: IP address set associated with the accelerator.
        :param pulumi.Input[str] name: The name of the accelerator.

        The **attributes** object supports the following:

          * `flowLogsEnabled` (`pulumi.Input[bool]`) - Indicates whether flow logs are enabled.
          * `flowLogsS3Bucket` (`pulumi.Input[str]`) - The name of the Amazon S3 bucket for the flow logs.
          * `flowLogsS3Prefix` (`pulumi.Input[str]`) - The prefix for the location in the Amazon S3 bucket for the flow logs.

        The **ip_sets** object supports the following:

          * `ip_addresses` (`pulumi.Input[list]`) - The array of IP addresses in the IP address set.
          * `ipFamily` (`pulumi.Input[str]`) - The types of IP addresses included in this IP set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["enabled"] = enabled
        __props__["ip_address_type"] = ip_address_type
        __props__["ip_sets"] = ip_sets
        __props__["name"] = name
        return Accelerator(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

