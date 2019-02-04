# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class WorkingStorage(pulumi.CustomResource):
    disk_id: pulumi.Output[str]
    """
    Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
    """
    gateway_arn: pulumi.Output[str]
    """
    The Amazon Resource Name (ARN) of the gateway.
    """
    def __init__(__self__, __name__, __opts__=None, disk_id=None, gateway_arn=None):
        """
        Manages an AWS Storage Gateway working storage.
        
        > **NOTE:** The Storage Gateway API provides no method to remove a working storage disk. Destroying this Terraform resource does not perform any Storage Gateway actions.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if disk_id is None:
            raise TypeError('Missing required property disk_id')
        __props__['disk_id'] = disk_id

        if gateway_arn is None:
            raise TypeError('Missing required property gateway_arn')
        __props__['gateway_arn'] = gateway_arn

        super(WorkingStorage, __self__).__init__(
            'aws:storagegateway/workingStorage:WorkingStorage',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

