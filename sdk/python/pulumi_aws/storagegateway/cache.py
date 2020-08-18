# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Cache']


class Cache(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disk_id: Optional[pulumi.Input[str]] = None,
                 gateway_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an AWS Storage Gateway cache.

        > **NOTE:** The Storage Gateway API provides no method to remove a cache disk. Destroying this resource does not perform any Storage Gateway actions.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.storagegateway.Cache("example",
            disk_id=data["aws_storagegateway_local_disk"]["example"]["id"],
            gateway_arn=aws_storagegateway_gateway["example"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
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

            if disk_id is None:
                raise TypeError("Missing required property 'disk_id'")
            __props__['disk_id'] = disk_id
            if gateway_arn is None:
                raise TypeError("Missing required property 'gateway_arn'")
            __props__['gateway_arn'] = gateway_arn
        super(Cache, __self__).__init__(
            'aws:storagegateway/cache:Cache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disk_id: Optional[pulumi.Input[str]] = None,
            gateway_arn: Optional[pulumi.Input[str]] = None) -> 'Cache':
        """
        Get an existing Cache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] disk_id: Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        :param pulumi.Input[str] gateway_arn: The Amazon Resource Name (ARN) of the gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disk_id"] = disk_id
        __props__["gateway_arn"] = gateway_arn
        return Cache(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="diskId")
    def disk_id(self) -> str:
        """
        Local disk identifier. For example, `pci-0000:03:00.0-scsi-0:0:0:0`.
        """
        return pulumi.get(self, "disk_id")

    @property
    @pulumi.getter(name="gatewayArn")
    def gateway_arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the gateway.
        """
        return pulumi.get(self, "gateway_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

