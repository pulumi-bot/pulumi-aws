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

__all__ = ['WebAclLoggingConfiguration']


class WebAclLoggingConfiguration(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 log_destination_configs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 redacted_fields: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationRedactedFieldArgs']]]]] = None,
                 resource_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a WAFv2 Web ACL Logging Configuration resource.

        > **Note:** To start logging from a WAFv2 Web ACL, an Amazon Kinesis Data Firehose (e.g. [`kinesis.FirehoseDeliveryStream` resource](https://www.terraform.io/docs/providers/aws/r/kinesis_firehose_delivery_stream.html) must also be created with a PUT source (not a stream) and in the region that you are operating.
        If you are capturing logs for Amazon CloudFront, always create the firehose in US East (N. Virginia).
        Be sure to give the data firehose a name that starts with the prefix `aws-waf-logs-`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.wafv2.WebAclLoggingConfiguration("example",
            log_destination_configs=[aws_kinesis_firehose_delivery_stream["example"]["arn"]],
            resource_arn=aws_wafv2_web_acl["example"]["arn"],
            redacted_fields=[{
                "singleHeader": {
                    "name": "user-agent",
                },
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] log_destination_configs: The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationRedactedFieldArgs']]]] redacted_fields: The parts of the request that you want to keep out of the logs. Up to 100 `redacted_fields` blocks are supported.
        :param pulumi.Input[str] resource_arn: The Amazon Resource Name (ARN) of the web ACL that you want to associate with `log_destination_configs`.
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

            if log_destination_configs is None:
                raise TypeError("Missing required property 'log_destination_configs'")
            __props__['log_destination_configs'] = log_destination_configs
            __props__['redacted_fields'] = redacted_fields
            if resource_arn is None:
                raise TypeError("Missing required property 'resource_arn'")
            __props__['resource_arn'] = resource_arn
        super(WebAclLoggingConfiguration, __self__).__init__(
            'aws:wafv2/webAclLoggingConfiguration:WebAclLoggingConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            log_destination_configs: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            redacted_fields: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationRedactedFieldArgs']]]]] = None,
            resource_arn: Optional[pulumi.Input[str]] = None) -> 'WebAclLoggingConfiguration':
        """
        Get an existing WebAclLoggingConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] log_destination_configs: The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['WebAclLoggingConfigurationRedactedFieldArgs']]]] redacted_fields: The parts of the request that you want to keep out of the logs. Up to 100 `redacted_fields` blocks are supported.
        :param pulumi.Input[str] resource_arn: The Amazon Resource Name (ARN) of the web ACL that you want to associate with `log_destination_configs`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["log_destination_configs"] = log_destination_configs
        __props__["redacted_fields"] = redacted_fields
        __props__["resource_arn"] = resource_arn
        return WebAclLoggingConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="logDestinationConfigs")
    def log_destination_configs(self) -> List[str]:
        """
        The Amazon Kinesis Data Firehose Amazon Resource Name (ARNs) that you want to associate with the web ACL. Currently, only 1 ARN is supported.
        """
        return pulumi.get(self, "log_destination_configs")

    @property
    @pulumi.getter(name="redactedFields")
    def redacted_fields(self) -> Optional[List['outputs.WebAclLoggingConfigurationRedactedField']]:
        """
        The parts of the request that you want to keep out of the logs. Up to 100 `redacted_fields` blocks are supported.
        """
        return pulumi.get(self, "redacted_fields")

    @property
    @pulumi.getter(name="resourceArn")
    def resource_arn(self) -> str:
        """
        The Amazon Resource Name (ARN) of the web ACL that you want to associate with `log_destination_configs`.
        """
        return pulumi.get(self, "resource_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

