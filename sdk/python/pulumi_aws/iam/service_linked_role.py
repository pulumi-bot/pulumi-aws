# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ServiceLinkedRole']


class ServiceLinkedRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 aws_service_name: Optional[pulumi.Input[str]] = None,
                 custom_suffix: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] aws_service_name: The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
        :param pulumi.Input[str] custom_suffix: Additional string appended to the role name. Not all AWS services support custom suffixes.
        :param pulumi.Input[str] description: The description of the role.
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

            if aws_service_name is None:
                raise TypeError("Missing required property 'aws_service_name'")
            __props__['aws_service_name'] = aws_service_name
            __props__['custom_suffix'] = custom_suffix
            __props__['description'] = description
            __props__['arn'] = None
            __props__['create_date'] = None
            __props__['name'] = None
            __props__['path'] = None
            __props__['unique_id'] = None
        super(ServiceLinkedRole, __self__).__init__(
            'aws:iam/serviceLinkedRole:ServiceLinkedRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            aws_service_name: Optional[pulumi.Input[str]] = None,
            create_date: Optional[pulumi.Input[str]] = None,
            custom_suffix: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            unique_id: Optional[pulumi.Input[str]] = None) -> 'ServiceLinkedRole':
        """
        Get an existing ServiceLinkedRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) specifying the role.
        :param pulumi.Input[str] aws_service_name: The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
        :param pulumi.Input[str] create_date: The creation date of the IAM role.
        :param pulumi.Input[str] custom_suffix: Additional string appended to the role name. Not all AWS services support custom suffixes.
        :param pulumi.Input[str] description: The description of the role.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] path: The path of the role.
        :param pulumi.Input[str] unique_id: The stable and unique string identifying the role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["aws_service_name"] = aws_service_name
        __props__["create_date"] = create_date
        __props__["custom_suffix"] = custom_suffix
        __props__["description"] = description
        __props__["name"] = name
        __props__["path"] = path
        __props__["unique_id"] = unique_id
        return ServiceLinkedRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) specifying the role.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="awsServiceName")
    def aws_service_name(self) -> pulumi.Output[str]:
        """
        The AWS service to which this role is attached. You use a string similar to a URL but without the `http://` in front. For example: `elasticbeanstalk.amazonaws.com`. To find the full list of services that support service-linked roles, check [the docs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-services-that-work-with-iam.html).
        """
        return pulumi.get(self, "aws_service_name")

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> pulumi.Output[str]:
        """
        The creation date of the IAM role.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter(name="customSuffix")
    def custom_suffix(self) -> pulumi.Output[Optional[str]]:
        """
        Additional string appended to the role name. Not all AWS services support custom suffixes.
        """
        return pulumi.get(self, "custom_suffix")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the role.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="uniqueId")
    def unique_id(self) -> pulumi.Output[str]:
        """
        The stable and unique string identifying the role.
        """
        return pulumi.get(self, "unique_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

