# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class NfsLocation(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the DataSync Location.
    """
    on_prem_config: pulumi.Output[dict]
    """
    Configuration block containing information for connecting to the NFS File System.
    """
    server_hostname: pulumi.Output[str]
    """
    Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
    """
    subdirectory: pulumi.Output[str]
    """
    Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of resource tags to assign to the DataSync Location.
    """
    uri: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, on_prem_config=None, server_hostname=None, subdirectory=None, tags=None, __name__=None, __opts__=None):
        """
        Manages an NFS Location within AWS DataSync.
        
        > **NOTE:** The DataSync Agents must be available before creating this resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] on_prem_config: Configuration block containing information for connecting to the NFS File System.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not on_prem_config:
            raise TypeError('Missing required property on_prem_config')
        __props__['on_prem_config'] = on_prem_config

        if not server_hostname:
            raise TypeError('Missing required property server_hostname')
        __props__['server_hostname'] = server_hostname

        if not subdirectory:
            raise TypeError('Missing required property subdirectory')
        __props__['subdirectory'] = subdirectory

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['uri'] = None

        super(NfsLocation, __self__).__init__(
            'aws:datasync/nfsLocation:NfsLocation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

