# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['NfsLocation']


class NfsLocation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 on_prem_config: Optional[pulumi.Input[pulumi.InputType['NfsLocationOnPremConfigArgs']]] = None,
                 server_hostname: Optional[pulumi.Input[str]] = None,
                 subdirectory: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an NFS Location within AWS DataSync.

        > **NOTE:** The DataSync Agents must be available before creating this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.datasync.NfsLocation("example",
            server_hostname="nfs.example.com",
            subdirectory="/exported/path",
            on_prem_config={
                "agent_arns": [aws_datasync_agent["example"]["arn"]],
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['NfsLocationOnPremConfigArgs']] on_prem_config: Configuration block containing information for connecting to the NFS File System.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
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

            if on_prem_config is None:
                raise TypeError("Missing required property 'on_prem_config'")
            __props__['on_prem_config'] = on_prem_config
            if server_hostname is None:
                raise TypeError("Missing required property 'server_hostname'")
            __props__['server_hostname'] = server_hostname
            if subdirectory is None:
                raise TypeError("Missing required property 'subdirectory'")
            __props__['subdirectory'] = subdirectory
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['uri'] = None
        super(NfsLocation, __self__).__init__(
            'aws:datasync/nfsLocation:NfsLocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            on_prem_config: Optional[pulumi.Input[pulumi.InputType['NfsLocationOnPremConfigArgs']]] = None,
            server_hostname: Optional[pulumi.Input[str]] = None,
            subdirectory: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uri: Optional[pulumi.Input[str]] = None) -> 'NfsLocation':
        """
        Get an existing NfsLocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the DataSync Location.
        :param pulumi.Input[pulumi.InputType['NfsLocationOnPremConfigArgs']] on_prem_config: Configuration block containing information for connecting to the NFS File System.
        :param pulumi.Input[str] server_hostname: Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
        :param pulumi.Input[str] subdirectory: Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value pairs of resource tags to assign to the DataSync Location.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["on_prem_config"] = on_prem_config
        __props__["server_hostname"] = server_hostname
        __props__["subdirectory"] = subdirectory
        __props__["tags"] = tags
        __props__["uri"] = uri
        return NfsLocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) of the DataSync Location.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="onPremConfig")
    def on_prem_config(self) -> 'outputs.NfsLocationOnPremConfig':
        """
        Configuration block containing information for connecting to the NFS File System.
        """
        return pulumi.get(self, "on_prem_config")

    @property
    @pulumi.getter(name="serverHostname")
    def server_hostname(self) -> str:
        """
        Specifies the IP address or DNS name of the NFS server. The DataSync Agent(s) use this to mount the NFS server.
        """
        return pulumi.get(self, "server_hostname")

    @property
    @pulumi.getter
    def subdirectory(self) -> str:
        """
        Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        """
        return pulumi.get(self, "subdirectory")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value pairs of resource tags to assign to the DataSync Location.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uri(self) -> str:
        return pulumi.get(self, "uri")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

