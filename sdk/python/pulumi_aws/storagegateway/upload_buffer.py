# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UploadBuffer(pulumi.CustomResource):
    disk_id: pulumi.Output[str]
    """
    Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
    """
    gateway_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the gateway.
    """
    def __init__(__self__, resource_name, opts=None, disk_id=None, gateway_arn=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an AWS Storage Gateway upload buffer.
        
        > **NOTE:** The Storage Gateway API provides no method to remove an upload buffer disk. Destroying this resource does not perform any Storage Gateway actions.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_upload_buffer.html.markdown.
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

            if disk_id is None:
                raise TypeError("Missing required property 'disk_id'")
            __props__['disk_id'] = disk_id
            if gateway_arn is None:
                raise TypeError("Missing required property 'gateway_arn'")
            __props__['gateway_arn'] = gateway_arn
        super(UploadBuffer, __self__).__init__(
            'aws:storagegateway/uploadBuffer:UploadBuffer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, disk_id=None, gateway_arn=None):
        """
        Get an existing UploadBuffer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/storagegateway_upload_buffer.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["disk_id"] = disk_id
        __props__["gateway_arn"] = gateway_arn
        return UploadBuffer(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

