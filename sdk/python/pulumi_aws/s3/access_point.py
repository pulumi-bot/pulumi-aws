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

__all__ = ['AccessPoint']


class AccessPoint(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 bucket: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 public_access_block_configuration: Optional[pulumi.Input[pulumi.InputType['AccessPointPublicAccessBlockConfigurationArgs']]] = None,
                 vpc_configuration: Optional[pulumi.Input[pulumi.InputType['AccessPointVpcConfigurationArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage an S3 Access Point.

        ## Example Usage

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
        :param pulumi.Input[str] bucket: The name of the bucket that you want to associate this access point with.
        :param pulumi.Input[str] name: The name you want to assign to this access point.
        :param pulumi.Input[str] policy: A valid JSON document that specifies the policy that you want to apply to this access point.
        :param pulumi.Input[pulumi.InputType['AccessPointPublicAccessBlockConfigurationArgs']] public_access_block_configuration: Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
        :param pulumi.Input[pulumi.InputType['AccessPointVpcConfigurationArgs']] vpc_configuration: Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Detailed below.
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

            __props__['account_id'] = account_id
            if bucket is None:
                raise TypeError("Missing required property 'bucket'")
            __props__['bucket'] = bucket
            __props__['name'] = name
            __props__['policy'] = policy
            __props__['public_access_block_configuration'] = public_access_block_configuration
            __props__['vpc_configuration'] = vpc_configuration
            __props__['arn'] = None
            __props__['domain_name'] = None
            __props__['has_public_access_policy'] = None
            __props__['network_origin'] = None
        super(AccessPoint, __self__).__init__(
            'aws:s3/accessPoint:AccessPoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bucket: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            has_public_access_policy: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_origin: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None,
            public_access_block_configuration: Optional[pulumi.Input[pulumi.InputType['AccessPointPublicAccessBlockConfigurationArgs']]] = None,
            vpc_configuration: Optional[pulumi.Input[pulumi.InputType['AccessPointVpcConfigurationArgs']]] = None) -> 'AccessPoint':
        """
        Get an existing AccessPoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of the S3 Access Point.
        :param pulumi.Input[str] bucket: The name of the bucket that you want to associate this access point with.
        :param pulumi.Input[str] domain_name: The DNS domain name of the S3 Access Point in the format _`name`_-_`account_id`_.s3-accesspoint._region_.amazonaws.com.
               Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
        :param pulumi.Input[bool] has_public_access_policy: Indicates whether this access point currently has a policy that allows public access.
        :param pulumi.Input[str] name: The name you want to assign to this access point.
        :param pulumi.Input[str] network_origin: Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
        :param pulumi.Input[str] policy: A valid JSON document that specifies the policy that you want to apply to this access point.
        :param pulumi.Input[pulumi.InputType['AccessPointPublicAccessBlockConfigurationArgs']] public_access_block_configuration: Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
        :param pulumi.Input[pulumi.InputType['AccessPointVpcConfigurationArgs']] vpc_configuration: Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Detailed below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["account_id"] = account_id
        __props__["arn"] = arn
        __props__["bucket"] = bucket
        __props__["domain_name"] = domain_name
        __props__["has_public_access_policy"] = has_public_access_policy
        __props__["name"] = name
        __props__["network_origin"] = network_origin
        __props__["policy"] = policy
        __props__["public_access_block_configuration"] = public_access_block_configuration
        __props__["vpc_configuration"] = vpc_configuration
        return AccessPoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID for the owner of the bucket for which you want to create an access point. Defaults to automatically determined account ID of the provider.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of the S3 Access Point.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def bucket(self) -> pulumi.Output[str]:
        """
        The name of the bucket that you want to associate this access point with.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        """
        The DNS domain name of the S3 Access Point in the format _`name`_-_`account_id`_.s3-accesspoint._region_.amazonaws.com.
        Note: S3 access points only support secure access by HTTPS. HTTP isn't supported.
        """
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="hasPublicAccessPolicy")
    def has_public_access_policy(self) -> pulumi.Output[bool]:
        """
        Indicates whether this access point currently has a policy that allows public access.
        """
        return pulumi.get(self, "has_public_access_policy")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name you want to assign to this access point.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkOrigin")
    def network_origin(self) -> pulumi.Output[str]:
        """
        Indicates whether this access point allows access from the public Internet. Values are `VPC` (the access point doesn't allow access from the public Internet) and `Internet` (the access point allows access from the public Internet, subject to the access point and bucket access policies).
        """
        return pulumi.get(self, "network_origin")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[Optional[str]]:
        """
        A valid JSON document that specifies the policy that you want to apply to this access point.
        """
        return pulumi.get(self, "policy")

    @property
    @pulumi.getter(name="publicAccessBlockConfiguration")
    def public_access_block_configuration(self) -> pulumi.Output[Optional['outputs.AccessPointPublicAccessBlockConfiguration']]:
        """
        Configuration block to manage the `PublicAccessBlock` configuration that you want to apply to this Amazon S3 bucket. You can enable the configuration options in any combination. Detailed below.
        """
        return pulumi.get(self, "public_access_block_configuration")

    @property
    @pulumi.getter(name="vpcConfiguration")
    def vpc_configuration(self) -> pulumi.Output[Optional['outputs.AccessPointVpcConfiguration']]:
        """
        Configuration block to restrict access to this access point to requests from the specified Virtual Private Cloud (VPC). Detailed below.
        """
        return pulumi.get(self, "vpc_configuration")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

