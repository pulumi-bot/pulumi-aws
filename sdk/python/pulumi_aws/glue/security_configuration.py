# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class SecurityConfiguration(pulumi.CustomResource):
    encryption_configuration: pulumi.Output[dict]
    """
    Configuration block containing encryption configuration. Detailed below.

      * `cloudwatchEncryption` (`dict`)
        * `cloudwatchEncryptionMode` (`str`) - Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
        * `kms_key_arn` (`str`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

      * `jobBookmarksEncryption` (`dict`)
        * `jobBookmarksEncryptionMode` (`str`) - Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
        * `kms_key_arn` (`str`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

      * `s3Encryption` (`dict`) - A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
        * `kms_key_arn` (`str`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
        * `s3EncryptionMode` (`str`) - Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
    """
    name: pulumi.Output[str]
    """
    Name of the security configuration.
    """
    def __init__(__self__, resource_name, opts=None, encryption_configuration=None, name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Glue Security Configuration.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.glue.SecurityConfiguration("example", encryption_configuration=aws.glue.SecurityConfigurationEncryptionConfigurationArgs(
            cloudwatch_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs(
                cloudwatch_encryption_mode="DISABLED",
            ),
            job_bookmarks_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs(
                job_bookmarks_encryption_mode="DISABLED",
            ),
            s3_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs(
                kms_key_arn=data["aws_kms_key"]["example"]["arn"],
                s3_encryption_mode="SSE-KMS",
            ),
        ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] encryption_configuration: Configuration block containing encryption configuration. Detailed below.
        :param pulumi.Input[str] name: Name of the security configuration.

        The **encryption_configuration** object supports the following:

          * `cloudwatchEncryption` (`pulumi.Input[dict]`)
            * `cloudwatchEncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

          * `jobBookmarksEncryption` (`pulumi.Input[dict]`)
            * `jobBookmarksEncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

          * `s3Encryption` (`pulumi.Input[dict]`) - A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
            * `s3EncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
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

            if encryption_configuration is None:
                raise TypeError("Missing required property 'encryption_configuration'")
            __props__['encryption_configuration'] = encryption_configuration
            __props__['name'] = name
        super(SecurityConfiguration, __self__).__init__(
            'aws:glue/securityConfiguration:SecurityConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, encryption_configuration=None, name=None):
        """
        Get an existing SecurityConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] encryption_configuration: Configuration block containing encryption configuration. Detailed below.
        :param pulumi.Input[str] name: Name of the security configuration.

        The **encryption_configuration** object supports the following:

          * `cloudwatchEncryption` (`pulumi.Input[dict]`)
            * `cloudwatchEncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for CloudWatch data. Valid values: `DISABLED`, `SSE-KMS`. Default value: `DISABLED`.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

          * `jobBookmarksEncryption` (`pulumi.Input[dict]`)
            * `jobBookmarksEncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for job bookmarks data. Valid values: `CSE-KMS`, `DISABLED`. Default value: `DISABLED`.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.

          * `s3Encryption` (`pulumi.Input[dict]`) - A `s3_encryption ` block as described below, which contains encryption configuration for S3 data.
            * `kms_key_arn` (`pulumi.Input[str]`) - Amazon Resource Name (ARN) of the KMS key to be used to encrypt the data.
            * `s3EncryptionMode` (`pulumi.Input[str]`) - Encryption mode to use for S3 data. Valid values: `DISABLED`, `SSE-KMS`, `SSE-S3`. Default value: `DISABLED`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["encryption_configuration"] = encryption_configuration
        __props__["name"] = name
        return SecurityConfiguration(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
