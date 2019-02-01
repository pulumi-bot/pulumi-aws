# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NetworkInterfaceAttachment(pulumi.CustomResource):
    attachment_id: pulumi.Output[str]
    """
    The ENI Attachment ID.
    """
    device_index: pulumi.Output[int]
    """
    Network interface index (int).
    """
    instance_id: pulumi.Output[str]
    """
    Instance ID to attach.
    """
    network_interface_id: pulumi.Output[str]
    """
    ENI ID to attach.
    """
    status: pulumi.Output[str]
    """
    The status of the Network Interface Attachment.
    """
    def __init__(__self__, __name__, __opts__=None, device_index=None, instance_id=None, network_interface_id=None):
        """
        Attach an Elastic network interface (ENI) resource with EC2 instance.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[int] device_index: Network interface index (int).
        :param pulumi.Input[str] instance_id: Instance ID to attach.
        :param pulumi.Input[str] network_interface_id: ENI ID to attach.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not device_index:
            raise TypeError('Missing required property device_index')
        __props__['device_index'] = device_index

        if not instance_id:
            raise TypeError('Missing required property instance_id')
        __props__['instance_id'] = instance_id

        if not network_interface_id:
            raise TypeError('Missing required property network_interface_id')
        __props__['network_interface_id'] = network_interface_id

        __props__['attachment_id'] = None
        __props__['status'] = None

        super(NetworkInterfaceAttachment, __self__).__init__(
            'aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

