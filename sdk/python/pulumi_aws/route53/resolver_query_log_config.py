# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['ResolverQueryLogConfig']


class ResolverQueryLogConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Route 53 Resolver query logging configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.route53.ResolverQueryLogConfig("example",
            destination_arn=aws_s3_bucket["example"]["arn"],
            tags={
                "Environment": "Prod",
            })
        ```

        ## Import

         Route 53 Resolver query logging configurations can be imported using the Route 53 Resolver query logging configuration ID, e.g.

        ```sh
         $ pulumi import aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig example rqlc-92edc3b1838248bf
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The ARN of the resource that you want Route 53 Resolver to send query logs.
               You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        :param pulumi.Input[str] name: The name of the Route 53 Resolver query logging configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
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

            if destination_arn is None:
                raise TypeError("Missing required property 'destination_arn'")
            __props__['destination_arn'] = destination_arn
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['owner_id'] = None
            __props__['share_status'] = None
        super(ResolverQueryLogConfig, __self__).__init__(
            'aws:route53/resolverQueryLogConfig:ResolverQueryLogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            destination_arn: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            share_status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'ResolverQueryLogConfig':
        """
        Get an existing ResolverQueryLogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
        :param pulumi.Input[str] destination_arn: The ARN of the resource that you want Route 53 Resolver to send query logs.
               You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        :param pulumi.Input[str] name: The name of the Route 53 Resolver query logging configuration.
        :param pulumi.Input[str] owner_id: The AWS account ID of the account that created the query logging configuration.
        :param pulumi.Input[str] share_status: An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
               Sharing is configured through AWS Resource Access Manager (AWS RAM).
               Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["destination_arn"] = destination_arn
        __props__["name"] = name
        __props__["owner_id"] = owner_id
        __props__["share_status"] = share_status
        __props__["tags"] = tags
        return ResolverQueryLogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN (Amazon Resource Name) of the Route 53 Resolver query logging configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the resource that you want Route 53 Resolver to send query logs.
        You can send query logs to an S3 bucket, a CloudWatch Logs log group, or a Kinesis Data Firehose delivery stream.
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Route 53 Resolver query logging configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        The AWS account ID of the account that created the query logging configuration.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="shareStatus")
    def share_status(self) -> pulumi.Output[str]:
        """
        An indication of whether the query logging configuration is shared with other AWS accounts, or was shared with the current account by another AWS account.
        Sharing is configured through AWS Resource Access Manager (AWS RAM).
        Values are `NOT_SHARED`, `SHARED_BY_ME` or `SHARED_WITH_ME`
        """
        return pulumi.get(self, "share_status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

