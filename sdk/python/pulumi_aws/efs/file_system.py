# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class FileSystem(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    Amazon Resource Name of the file system.
    """
    creation_token: pulumi.Output[str]
    """
    A unique name (a maximum of 64 characters are allowed)
    used as reference when creating the Elastic File System to ensure idempotent file
    system creation. By default generated by Terraform. See [Elastic File System]
    (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name for the filesystem per [documented convention](http://docs.aws.amazon.com/efs/latest/ug/mounting-fs-mount-cmd-dns-name.html).
    """
    encrypted: pulumi.Output[bool]
    """
    If true, the disk will be encrypted.
    """
    kms_key_id: pulumi.Output[str]
    """
    The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
    """
    performance_mode: pulumi.Output[str]
    """
    The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
    """
    provisioned_throughput_in_mibps: pulumi.Output[float]
    """
    The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughput_mode` set to `provisioned`.
    """
    reference_name: pulumi.Output[str]
    """
    **DEPRECATED** (Optional) A reference name used when creating the
    `Creation Token` which Amazon EFS uses to ensure idempotent file system creation. By
    default generated by Terraform.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the file system.
    """
    throughput_mode: pulumi.Output[str]
    """
    Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisioned_throughput_in_mibps`.
    """
    def __init__(__self__, resource_name, opts=None, creation_token=None, encrypted=None, kms_key_id=None, performance_mode=None, provisioned_throughput_in_mibps=None, reference_name=None, tags=None, throughput_mode=None, __name__=None, __opts__=None):
        """
        Provides an Elastic File System (EFS) resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_token: A unique name (a maximum of 64 characters are allowed)
               used as reference when creating the Elastic File System to ensure idempotent file
               system creation. By default generated by Terraform. See [Elastic File System]
               (http://docs.aws.amazon.com/efs/latest/ug/) user guide for more information.
        :param pulumi.Input[bool] encrypted: If true, the disk will be encrypted.
        :param pulumi.Input[str] kms_key_id: The ARN for the KMS encryption key. When specifying kms_key_id, encrypted needs to be set to true.
        :param pulumi.Input[str] performance_mode: The file system performance mode. Can be either `"generalPurpose"` or `"maxIO"` (Default: `"generalPurpose"`).
        :param pulumi.Input[float] provisioned_throughput_in_mibps: The throughput, measured in MiB/s, that you want to provision for the file system. Only applicable with `throughput_mode` set to `provisioned`.
        :param pulumi.Input[str] reference_name: **DEPRECATED** (Optional) A reference name used when creating the
               `Creation Token` which Amazon EFS uses to ensure idempotent file system creation. By
               default generated by Terraform.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the file system.
        :param pulumi.Input[str] throughput_mode: Throughput mode for the file system. Defaults to `bursting`. Valid values: `bursting`, `provisioned`. When using `provisioned`, also set `provisioned_throughput_in_mibps`.
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

        __props__['creation_token'] = creation_token

        __props__['encrypted'] = encrypted

        __props__['kms_key_id'] = kms_key_id

        __props__['performance_mode'] = performance_mode

        __props__['provisioned_throughput_in_mibps'] = provisioned_throughput_in_mibps

        __props__['reference_name'] = reference_name

        __props__['tags'] = tags

        __props__['throughput_mode'] = throughput_mode

        __props__['arn'] = None
        __props__['dns_name'] = None

        super(FileSystem, __self__).__init__(
            'aws:efs/fileSystem:FileSystem',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

