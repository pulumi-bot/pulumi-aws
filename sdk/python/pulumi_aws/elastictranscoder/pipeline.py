# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Pipeline']


class Pipeline(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
                 content_config: Optional[pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']]] = None,
                 content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]]] = None,
                 input_bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notifications: Optional[pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']]] = None,
                 output_bucket: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 thumbnail_config: Optional[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']]] = None,
                 thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Elastic Transcoder pipeline resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
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

            __props__['aws_kms_key_arn'] = aws_kms_key_arn
            __props__['content_config'] = content_config
            __props__['content_config_permissions'] = content_config_permissions
            if input_bucket is None:
                raise TypeError("Missing required property 'input_bucket'")
            __props__['input_bucket'] = input_bucket
            __props__['name'] = name
            __props__['notifications'] = notifications
            __props__['output_bucket'] = output_bucket
            if role is None:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['thumbnail_config'] = thumbnail_config
            __props__['thumbnail_config_permissions'] = thumbnail_config_permissions
            __props__['arn'] = None
        super(Pipeline, __self__).__init__(
            'aws:elastictranscoder/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_kms_key_arn: Optional[pulumi.Input[str]] = None,
            content_config: Optional[pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']]] = None,
            content_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]]] = None,
            input_bucket: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notifications: Optional[pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']]] = None,
            output_bucket: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            thumbnail_config: Optional[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']]] = None,
            thumbnail_config_permissions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]]] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_kms_key_arn: The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineContentConfigArgs']] content_config: The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineContentConfigPermissionArgs']]]] content_config_permissions: The permissions for the `content_config` object. (documented below)
        :param pulumi.Input[str] input_bucket: The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        :param pulumi.Input[str] name: The name of the pipeline. Maximum 40 characters
        :param pulumi.Input[pulumi.InputType['PipelineNotificationsArgs']] notifications: The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        :param pulumi.Input[str] output_bucket: The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        :param pulumi.Input[str] role: The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        :param pulumi.Input[pulumi.InputType['PipelineThumbnailConfigArgs']] thumbnail_config: The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['PipelineThumbnailConfigPermissionArgs']]]] thumbnail_config_permissions: The permissions for the `thumbnail_config` object. (documented below)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["aws_kms_key_arn"] = aws_kms_key_arn
        __props__["content_config"] = content_config
        __props__["content_config_permissions"] = content_config_permissions
        __props__["input_bucket"] = input_bucket
        __props__["name"] = name
        __props__["notifications"] = notifications
        __props__["output_bucket"] = output_bucket
        __props__["role"] = role
        __props__["thumbnail_config"] = thumbnail_config
        __props__["thumbnail_config_permissions"] = thumbnail_config_permissions
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsKmsKeyArn")
    def aws_kms_key_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.
        """
        return pulumi.get(self, "aws_kms_key_arn")

    @property
    @pulumi.getter(name="contentConfig")
    def content_config(self) -> pulumi.Output['outputs.PipelineContentConfig']:
        """
        The ContentConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists. (documented below)
        """
        return pulumi.get(self, "content_config")

    @property
    @pulumi.getter(name="contentConfigPermissions")
    def content_config_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineContentConfigPermission']]]:
        """
        The permissions for the `content_config` object. (documented below)
        """
        return pulumi.get(self, "content_config_permissions")

    @property
    @pulumi.getter(name="inputBucket")
    def input_bucket(self) -> pulumi.Output[str]:
        """
        The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.
        """
        return pulumi.get(self, "input_bucket")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the pipeline. Maximum 40 characters
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Optional['outputs.PipelineNotifications']]:
        """
        The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status. (documented below)
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="outputBucket")
    def output_bucket(self) -> pulumi.Output[str]:
        """
        The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.
        """
        return pulumi.get(self, "output_bucket")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="thumbnailConfig")
    def thumbnail_config(self) -> pulumi.Output['outputs.PipelineThumbnailConfig']:
        """
        The ThumbnailConfig object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files. (documented below)
        """
        return pulumi.get(self, "thumbnail_config")

    @property
    @pulumi.getter(name="thumbnailConfigPermissions")
    def thumbnail_config_permissions(self) -> pulumi.Output[Optional[Sequence['outputs.PipelineThumbnailConfigPermission']]]:
        """
        The permissions for the `thumbnail_config` object. (documented below)
        """
        return pulumi.get(self, "thumbnail_config_permissions")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

