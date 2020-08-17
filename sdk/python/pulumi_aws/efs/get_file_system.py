# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'GetFileSystemResult',
    'AwaitableGetFileSystemResult',
    'get_file_system',
]



@pulumi.output_type
class GetFileSystemResult:
    """
    A collection of values returned by getFileSystem.
    """
    def __init__(__self__, arn=None, creation_token=None, dns_name=None, encrypted=None, file_system_id=None, id=None, kms_key_id=None, lifecycle_policy=None, performance_mode=None, provisioned_throughput_in_mibps=None, size_in_bytes=None, tags=None, throughput_mode=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if creation_token and not isinstance(creation_token, str):
            raise TypeError("Expected argument 'creation_token' to be a str")
        pulumi.set(__self__, "creation_token", creation_token)
        if dns_name and not isinstance(dns_name, str):
            raise TypeError("Expected argument 'dns_name' to be a str")
        pulumi.set(__self__, "dns_name", dns_name)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if file_system_id and not isinstance(file_system_id, str):
            raise TypeError("Expected argument 'file_system_id' to be a str")
        pulumi.set(__self__, "file_system_id", file_system_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kms_key_id and not isinstance(kms_key_id, str):
            raise TypeError("Expected argument 'kms_key_id' to be a str")
        pulumi.set(__self__, "kms_key_id", kms_key_id)
        if lifecycle_policy and not isinstance(lifecycle_policy, dict):
            raise TypeError("Expected argument 'lifecycle_policy' to be a dict")
        pulumi.set(__self__, "lifecycle_policy", lifecycle_policy)
        if performance_mode and not isinstance(performance_mode, str):
            raise TypeError("Expected argument 'performance_mode' to be a str")
        pulumi.set(__self__, "performance_mode", performance_mode)
        if provisioned_throughput_in_mibps and not isinstance(provisioned_throughput_in_mibps, float):
            raise TypeError("Expected argument 'provisioned_throughput_in_mibps' to be a float")
        pulumi.set(__self__, "provisioned_throughput_in_mibps", provisioned_throughput_in_mibps)
        if size_in_bytes and not isinstance(size_in_bytes, float):
            raise TypeError("Expected argument 'size_in_bytes' to be a float")
        pulumi.set(__self__, "size_in_bytes", size_in_bytes)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if throughput_mode and not isinstance(throughput_mode, str):
            raise TypeError("Expected argument 'throughput_mode' to be a str")
        pulumi.set(__self__, "throughput_mode", throughput_mode)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name of the file system.
        """
        ...

    @property
    @pulumi.getter(name="creationToken")
    def creation_token(self) -> str:
        ...

    @property
    @pulumi.getter(name="dnsName")
    def dns_name(self) -> str:
        """
        The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
        """
        ...

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        """
        Whether EFS is encrypted.
        """
        ...

    @property
    @pulumi.getter(name="fileSystemId")
    def file_system_id(self) -> str:
        ...

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        ...

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> str:
        """
        The ARN for the KMS encryption key.
        """
        ...

    @property
    @pulumi.getter(name="lifecyclePolicy")
    def lifecycle_policy(self) -> 'outputs.GetFileSystemLifecyclePolicyResult':
        """
        A file system [lifecycle policy](https://docs.aws.amazon.com/efs/latest/ug/API_LifecyclePolicy.html) object.
        """
        ...

    @property
    @pulumi.getter(name="performanceMode")
    def performance_mode(self) -> str:
        """
        The file system performance mode.
        """
        ...

    @property
    @pulumi.getter(name="provisionedThroughputInMibps")
    def provisioned_throughput_in_mibps(self) -> float:
        """
        The throughput, measured in MiB/s, that you want to provision for the file system.
        * `tags` -A map of tags to assign to the file system.
        """
        ...

    @property
    @pulumi.getter(name="sizeInBytes")
    def size_in_bytes(self) -> float:
        """
        The current byte count used by the file system.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        ...

    @property
    @pulumi.getter(name="throughputMode")
    def throughput_mode(self) -> str:
        """
        Throughput mode for the file system.
        """
        ...



class AwaitableGetFileSystemResult(GetFileSystemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileSystemResult(
            arn=self.arn,
            creation_token=self.creation_token,
            dns_name=self.dns_name,
            encrypted=self.encrypted,
            file_system_id=self.file_system_id,
            id=self.id,
            kms_key_id=self.kms_key_id,
            lifecycle_policy=self.lifecycle_policy,
            performance_mode=self.performance_mode,
            provisioned_throughput_in_mibps=self.provisioned_throughput_in_mibps,
            size_in_bytes=self.size_in_bytes,
            tags=self.tags,
            throughput_mode=self.throughput_mode)


def get_file_system(creation_token: Optional[str] = None,
                    file_system_id: Optional[str] = None,
                    tags: Optional[Mapping[str, str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileSystemResult:
    """
    Provides information about an Elastic File System (EFS) File System.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    file_system_id = config.get("fileSystemId")
    if file_system_id is None:
        file_system_id = ""
    by_id = aws.efs.get_file_system(file_system_id=file_system_id)
    ```


    :param str creation_token: Restricts the list to the file system with this creation token.
    :param str file_system_id: The ID that identifies the file system (e.g. fs-ccfc0d65).
    """
    __args__ = dict()
    __args__['creationToken'] = creation_token
    __args__['fileSystemId'] = file_system_id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:efs/getFileSystem:getFileSystem', __args__, opts=opts, typ=GetFileSystemResult).value

    return AwaitableGetFileSystemResult(
        arn=__ret__.arn,
        creation_token=__ret__.creation_token,
        dns_name=__ret__.dns_name,
        encrypted=__ret__.encrypted,
        file_system_id=__ret__.file_system_id,
        id=__ret__.id,
        kms_key_id=__ret__.kms_key_id,
        lifecycle_policy=__ret__.lifecycle_policy,
        performance_mode=__ret__.performance_mode,
        provisioned_throughput_in_mibps=__ret__.provisioned_throughput_in_mibps,
        size_in_bytes=__ret__.size_in_bytes,
        tags=__ret__.tags,
        throughput_mode=__ret__.throughput_mode)
