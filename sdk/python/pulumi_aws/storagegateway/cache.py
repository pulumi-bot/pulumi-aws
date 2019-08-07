# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Cache(pulumi.CustomResource):
    disk_id: pulumi.Output[str]
    """
    Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
    """
    gateway_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the gateway.
    """
    def __init__(__self__, resource_name, opts=None, disk_id=None, gateway_arn=None, __name__=None, __opts__=None):
        """
        Manages an AWS Storage Gateway cache.
        
        > **NOTE:** The Storage Gateway API provides no method to remove a cache disk. Destroying this resource does not perform any Storage Gateway actions.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_cache.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if disk_id is None:
            raise TypeError("Missing required property 'disk_id'")
        __props__['disk_id'] = disk_id
        if gateway_arn is None:
            raise TypeError("Missing required property 'gateway_arn'")
        __props__['gateway_arn'] = gateway_arn
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Cache, __self__).__init__(
            'aws:storagegateway/cache:Cache',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

