# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['RealtimeLogConfigArgs', 'RealtimeLogConfig']

@pulumi.input_type
class RealtimeLogConfigArgs:
    def __init__(__self__, *,
                 endpoint: pulumi.Input['RealtimeLogConfigEndpointArgs'],
                 fields: pulumi.Input[Sequence[pulumi.Input[str]]],
                 sampling_rate: pulumi.Input[int],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RealtimeLogConfig resource.
        :param pulumi.Input['RealtimeLogConfigEndpointArgs'] endpoint: The Amazon Kinesis data streams where real-time log data is sent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fields: The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
        :param pulumi.Input[int] sampling_rate: The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
        :param pulumi.Input[str] name: The unique name to identify this real-time log configuration.
        """
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "fields", fields)
        pulumi.set(__self__, "sampling_rate", sampling_rate)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input['RealtimeLogConfigEndpointArgs']:
        """
        The Amazon Kinesis data streams where real-time log data is sent.
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input['RealtimeLogConfigEndpointArgs']):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter(name="samplingRate")
    def sampling_rate(self) -> pulumi.Input[int]:
        """
        The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
        """
        return pulumi.get(self, "sampling_rate")

    @sampling_rate.setter
    def sampling_rate(self, value: pulumi.Input[int]):
        pulumi.set(self, "sampling_rate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name to identify this real-time log configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class RealtimeLogConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['RealtimeLogConfigEndpointArgs']]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sampling_rate: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a CloudFront real-time log configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "cloudfront.amazonaws.com"
              },
              "Effect": "Allow"
            }
          ]
        }
        \"\"\")
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
                "Effect": "Allow",
                "Action": [
                  "kinesis:DescribeStreamSummary",
                  "kinesis:DescribeStream",
                  "kinesis:PutRecord",
                  "kinesis:PutRecords"
                ],
                "Resource": "{aws_kinesis_stream["example"]["arn"]}"
            }}
          ]
        }}
        \"\"\")
        example_realtime_log_config = aws.cloudfront.RealtimeLogConfig("exampleRealtimeLogConfig",
            sampling_rate=75,
            fields=[
                "timestamp",
                "c-ip",
            ],
            endpoint=aws.cloudfront.RealtimeLogConfigEndpointArgs(
                stream_type="Kinesis",
                kinesis_stream_config=aws.cloudfront.RealtimeLogConfigEndpointKinesisStreamConfigArgs(
                    role_arn=example_role.arn,
                    stream_arn=aws_kinesis_stream["example"]["arn"],
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[example_role_policy]))
        ```

        ## Import

        CloudFront real-time log configurations can be imported using the ARN, e.g.

        ```sh
         $ pulumi import aws:cloudfront/realtimeLogConfig:RealtimeLogConfig example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RealtimeLogConfigEndpointArgs']] endpoint: The Amazon Kinesis data streams where real-time log data is sent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fields: The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
        :param pulumi.Input[str] name: The unique name to identify this real-time log configuration.
        :param pulumi.Input[int] sampling_rate: The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealtimeLogConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a CloudFront real-time log configuration resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": "sts:AssumeRole",
              "Principal": {
                "Service": "cloudfront.amazonaws.com"
              },
              "Effect": "Allow"
            }
          ]
        }
        \"\"\")
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
                "Effect": "Allow",
                "Action": [
                  "kinesis:DescribeStreamSummary",
                  "kinesis:DescribeStream",
                  "kinesis:PutRecord",
                  "kinesis:PutRecords"
                ],
                "Resource": "{aws_kinesis_stream["example"]["arn"]}"
            }}
          ]
        }}
        \"\"\")
        example_realtime_log_config = aws.cloudfront.RealtimeLogConfig("exampleRealtimeLogConfig",
            sampling_rate=75,
            fields=[
                "timestamp",
                "c-ip",
            ],
            endpoint=aws.cloudfront.RealtimeLogConfigEndpointArgs(
                stream_type="Kinesis",
                kinesis_stream_config=aws.cloudfront.RealtimeLogConfigEndpointKinesisStreamConfigArgs(
                    role_arn=example_role.arn,
                    stream_arn=aws_kinesis_stream["example"]["arn"],
                ),
            ),
            opts=pulumi.ResourceOptions(depends_on=[example_role_policy]))
        ```

        ## Import

        CloudFront real-time log configurations can be imported using the ARN, e.g.

        ```sh
         $ pulumi import aws:cloudfront/realtimeLogConfig:RealtimeLogConfig example arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig
        ```

        :param str resource_name: The name of the resource.
        :param RealtimeLogConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealtimeLogConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint: Optional[pulumi.Input[pulumi.InputType['RealtimeLogConfigEndpointArgs']]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sampling_rate: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__['endpoint'] = endpoint
            if fields is None and not opts.urn:
                raise TypeError("Missing required property 'fields'")
            __props__['fields'] = fields
            __props__['name'] = name
            if sampling_rate is None and not opts.urn:
                raise TypeError("Missing required property 'sampling_rate'")
            __props__['sampling_rate'] = sampling_rate
            __props__['arn'] = None
        super(RealtimeLogConfig, __self__).__init__(
            'aws:cloudfront/realtimeLogConfig:RealtimeLogConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            endpoint: Optional[pulumi.Input[pulumi.InputType['RealtimeLogConfigEndpointArgs']]] = None,
            fields: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            sampling_rate: Optional[pulumi.Input[int]] = None) -> 'RealtimeLogConfig':
        """
        Get an existing RealtimeLogConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
        :param pulumi.Input[pulumi.InputType['RealtimeLogConfigEndpointArgs']] endpoint: The Amazon Kinesis data streams where real-time log data is sent.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] fields: The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
        :param pulumi.Input[str] name: The unique name to identify this real-time log configuration.
        :param pulumi.Input[int] sampling_rate: The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["endpoint"] = endpoint
        __props__["fields"] = fields
        __props__["name"] = name
        __props__["sampling_rate"] = sampling_rate
        return RealtimeLogConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN (Amazon Resource Name) of the CloudFront real-time log configuration.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output['outputs.RealtimeLogConfigEndpoint']:
        """
        The Amazon Kinesis data streams where real-time log data is sent.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence[str]]:
        """
        The fields that are included in each real-time log record. See the [AWS documentation](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields) for supported values.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name to identify this real-time log configuration.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="samplingRate")
    def sampling_rate(self) -> pulumi.Output[int]:
        """
        The sampling rate for this real-time log configuration. The sampling rate determines the percentage of viewer requests that are represented in the real-time log data. An integer between `1` and `100`, inclusive.
        """
        return pulumi.get(self, "sampling_rate")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

