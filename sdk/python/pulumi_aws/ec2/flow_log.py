# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['FlowLogArgs', 'FlowLog']

@pulumi.input_type
class FlowLogArgs:
    def __init__(__self__, *,
                 traffic_type: pulumi.Input[str],
                 eni_id: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 log_destination: Optional[pulumi.Input[str]] = None,
                 log_destination_type: Optional[pulumi.Input[str]] = None,
                 log_format: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 max_aggregation_interval: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FlowLog resource.
        :param pulumi.Input[str] traffic_type: The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        :param pulumi.Input[str] eni_id: Elastic Network Interface ID to attach to
        :param pulumi.Input[str] iam_role_arn: The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        :param pulumi.Input[str] log_destination: The ARN of the logging destination.
        :param pulumi.Input[str] log_destination_type: The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        :param pulumi.Input[str] log_format: The fields to include in the flow log record, in the order in which they should appear.
        :param pulumi.Input[str] log_group_name: *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        :param pulumi.Input[int] max_aggregation_interval: The maximum interval of time
               during which a flow of packets is captured and aggregated into a flow
               log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
               minutes). Default: `600`.
        :param pulumi.Input[str] subnet_id: Subnet ID to attach to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] vpc_id: VPC ID to attach to
        """
        pulumi.set(__self__, "traffic_type", traffic_type)
        if eni_id is not None:
            pulumi.set(__self__, "eni_id", eni_id)
        if iam_role_arn is not None:
            pulumi.set(__self__, "iam_role_arn", iam_role_arn)
        if log_destination is not None:
            pulumi.set(__self__, "log_destination", log_destination)
        if log_destination_type is not None:
            pulumi.set(__self__, "log_destination_type", log_destination_type)
        if log_format is not None:
            pulumi.set(__self__, "log_format", log_format)
        if log_group_name is not None:
            warnings.warn("""use 'log_destination' argument instead""", DeprecationWarning)
            pulumi.log.warn("""log_group_name is deprecated: use 'log_destination' argument instead""")
        if log_group_name is not None:
            pulumi.set(__self__, "log_group_name", log_group_name)
        if max_aggregation_interval is not None:
            pulumi.set(__self__, "max_aggregation_interval", max_aggregation_interval)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Input[str]:
        """
        The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        """
        return pulumi.get(self, "traffic_type")

    @traffic_type.setter
    def traffic_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "traffic_type", value)

    @property
    @pulumi.getter(name="eniId")
    def eni_id(self) -> Optional[pulumi.Input[str]]:
        """
        Elastic Network Interface ID to attach to
        """
        return pulumi.get(self, "eni_id")

    @eni_id.setter
    def eni_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eni_id", value)

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        """
        return pulumi.get(self, "iam_role_arn")

    @iam_role_arn.setter
    def iam_role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_role_arn", value)

    @property
    @pulumi.getter(name="logDestination")
    def log_destination(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the logging destination.
        """
        return pulumi.get(self, "log_destination")

    @log_destination.setter
    def log_destination(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_destination", value)

    @property
    @pulumi.getter(name="logDestinationType")
    def log_destination_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        """
        return pulumi.get(self, "log_destination_type")

    @log_destination_type.setter
    def log_destination_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_destination_type", value)

    @property
    @pulumi.getter(name="logFormat")
    def log_format(self) -> Optional[pulumi.Input[str]]:
        """
        The fields to include in the flow log record, in the order in which they should appear.
        """
        return pulumi.get(self, "log_format")

    @log_format.setter
    def log_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_format", value)

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        """
        return pulumi.get(self, "log_group_name")

    @log_group_name.setter
    def log_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_group_name", value)

    @property
    @pulumi.getter(name="maxAggregationInterval")
    def max_aggregation_interval(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum interval of time
        during which a flow of packets is captured and aggregated into a flow
        log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
        minutes). Default: `600`.
        """
        return pulumi.get(self, "max_aggregation_interval")

    @max_aggregation_interval.setter
    def max_aggregation_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_aggregation_interval", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        Subnet ID to attach to
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        VPC ID to attach to
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class FlowLog(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowLogArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
        interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.

        ## Example Usage
        ### CloudWatch Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "vpc-flow-logs.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
            iam_role_arn=example_role.arn,
            log_destination=example_log_group.arn,
            traffic_type="ALL",
            vpc_id=aws_vpc["example"]["id"])
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams"
              ],
              "Effect": "Allow",
              "Resource": "*"
            }
          ]
        }
        \"\"\")
        ```
        ### S3 Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
            log_destination=example_bucket.arn,
            log_destination_type="s3",
            traffic_type="ALL",
            vpc_id=aws_vpc["example"]["id"])
        ```

        ## Import

        Flow Logs can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/flowLog:FlowLog test_flow_log fl-1a2b3c4d
        ```

        :param str resource_name: The name of the resource.
        :param FlowLogArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eni_id: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 log_destination: Optional[pulumi.Input[str]] = None,
                 log_destination_type: Optional[pulumi.Input[str]] = None,
                 log_format: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 max_aggregation_interval: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a VPC/Subnet/ENI Flow Log to capture IP traffic for a specific network
        interface, subnet, or VPC. Logs are sent to a CloudWatch Log Group or a S3 Bucket.

        ## Example Usage
        ### CloudWatch Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "",
              "Effect": "Allow",
              "Principal": {
                "Service": "vpc-flow-logs.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }
        \"\"\")
        example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
            iam_role_arn=example_role.arn,
            log_destination=example_log_group.arn,
            traffic_type="ALL",
            vpc_id=aws_vpc["example"]["id"])
        example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
            role=example_role.id,
            policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": [
                "logs:CreateLogGroup",
                "logs:CreateLogStream",
                "logs:PutLogEvents",
                "logs:DescribeLogGroups",
                "logs:DescribeLogStreams"
              ],
              "Effect": "Allow",
              "Resource": "*"
            }
          ]
        }
        \"\"\")
        ```
        ### S3 Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        example_bucket = aws.s3.Bucket("exampleBucket")
        example_flow_log = aws.ec2.FlowLog("exampleFlowLog",
            log_destination=example_bucket.arn,
            log_destination_type="s3",
            traffic_type="ALL",
            vpc_id=aws_vpc["example"]["id"])
        ```

        ## Import

        Flow Logs can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/flowLog:FlowLog test_flow_log fl-1a2b3c4d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eni_id: Elastic Network Interface ID to attach to
        :param pulumi.Input[str] iam_role_arn: The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        :param pulumi.Input[str] log_destination: The ARN of the logging destination.
        :param pulumi.Input[str] log_destination_type: The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        :param pulumi.Input[str] log_format: The fields to include in the flow log record, in the order in which they should appear.
        :param pulumi.Input[str] log_group_name: *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        :param pulumi.Input[int] max_aggregation_interval: The maximum interval of time
               during which a flow of packets is captured and aggregated into a flow
               log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
               minutes). Default: `600`.
        :param pulumi.Input[str] subnet_id: Subnet ID to attach to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] traffic_type: The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        :param pulumi.Input[str] vpc_id: VPC ID to attach to
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowLogArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eni_id: Optional[pulumi.Input[str]] = None,
                 iam_role_arn: Optional[pulumi.Input[str]] = None,
                 log_destination: Optional[pulumi.Input[str]] = None,
                 log_destination_type: Optional[pulumi.Input[str]] = None,
                 log_format: Optional[pulumi.Input[str]] = None,
                 log_group_name: Optional[pulumi.Input[str]] = None,
                 max_aggregation_interval: Optional[pulumi.Input[int]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 traffic_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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

            __props__['eni_id'] = eni_id
            __props__['iam_role_arn'] = iam_role_arn
            __props__['log_destination'] = log_destination
            __props__['log_destination_type'] = log_destination_type
            __props__['log_format'] = log_format
            if log_group_name is not None and not opts.urn:
                warnings.warn("""use 'log_destination' argument instead""", DeprecationWarning)
                pulumi.log.warn("""log_group_name is deprecated: use 'log_destination' argument instead""")
            __props__['log_group_name'] = log_group_name
            __props__['max_aggregation_interval'] = max_aggregation_interval
            __props__['subnet_id'] = subnet_id
            __props__['tags'] = tags
            if traffic_type is None and not opts.urn:
                raise TypeError("Missing required property 'traffic_type'")
            __props__['traffic_type'] = traffic_type
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
        super(FlowLog, __self__).__init__(
            'aws:ec2/flowLog:FlowLog',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            eni_id: Optional[pulumi.Input[str]] = None,
            iam_role_arn: Optional[pulumi.Input[str]] = None,
            log_destination: Optional[pulumi.Input[str]] = None,
            log_destination_type: Optional[pulumi.Input[str]] = None,
            log_format: Optional[pulumi.Input[str]] = None,
            log_group_name: Optional[pulumi.Input[str]] = None,
            max_aggregation_interval: Optional[pulumi.Input[int]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            traffic_type: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'FlowLog':
        """
        Get an existing FlowLog resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Flow Log.
        :param pulumi.Input[str] eni_id: Elastic Network Interface ID to attach to
        :param pulumi.Input[str] iam_role_arn: The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        :param pulumi.Input[str] log_destination: The ARN of the logging destination.
        :param pulumi.Input[str] log_destination_type: The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        :param pulumi.Input[str] log_format: The fields to include in the flow log record, in the order in which they should appear.
        :param pulumi.Input[str] log_group_name: *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        :param pulumi.Input[int] max_aggregation_interval: The maximum interval of time
               during which a flow of packets is captured and aggregated into a flow
               log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
               minutes). Default: `600`.
        :param pulumi.Input[str] subnet_id: Subnet ID to attach to
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
        :param pulumi.Input[str] traffic_type: The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        :param pulumi.Input[str] vpc_id: VPC ID to attach to
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["eni_id"] = eni_id
        __props__["iam_role_arn"] = iam_role_arn
        __props__["log_destination"] = log_destination
        __props__["log_destination_type"] = log_destination_type
        __props__["log_format"] = log_format
        __props__["log_group_name"] = log_group_name
        __props__["max_aggregation_interval"] = max_aggregation_interval
        __props__["subnet_id"] = subnet_id
        __props__["tags"] = tags
        __props__["traffic_type"] = traffic_type
        __props__["vpc_id"] = vpc_id
        return FlowLog(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the Flow Log.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="eniId")
    def eni_id(self) -> pulumi.Output[Optional[str]]:
        """
        Elastic Network Interface ID to attach to
        """
        return pulumi.get(self, "eni_id")

    @property
    @pulumi.getter(name="iamRoleArn")
    def iam_role_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN for the IAM role that's used to post flow logs to a CloudWatch Logs log group
        """
        return pulumi.get(self, "iam_role_arn")

    @property
    @pulumi.getter(name="logDestination")
    def log_destination(self) -> pulumi.Output[str]:
        """
        The ARN of the logging destination.
        """
        return pulumi.get(self, "log_destination")

    @property
    @pulumi.getter(name="logDestinationType")
    def log_destination_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the logging destination. Valid values: `cloud-watch-logs`, `s3`. Default: `cloud-watch-logs`.
        """
        return pulumi.get(self, "log_destination_type")

    @property
    @pulumi.getter(name="logFormat")
    def log_format(self) -> pulumi.Output[str]:
        """
        The fields to include in the flow log record, in the order in which they should appear.
        """
        return pulumi.get(self, "log_format")

    @property
    @pulumi.getter(name="logGroupName")
    def log_group_name(self) -> pulumi.Output[str]:
        """
        *Deprecated:* Use `log_destination` instead. The name of the CloudWatch log group.
        """
        return pulumi.get(self, "log_group_name")

    @property
    @pulumi.getter(name="maxAggregationInterval")
    def max_aggregation_interval(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum interval of time
        during which a flow of packets is captured and aggregated into a flow
        log record. Valid Values: `60` seconds (1 minute) or `600` seconds (10
        minutes). Default: `600`.
        """
        return pulumi.get(self, "max_aggregation_interval")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[Optional[str]]:
        """
        Subnet ID to attach to
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="trafficType")
    def traffic_type(self) -> pulumi.Output[str]:
        """
        The type of traffic to capture. Valid values: `ACCEPT`,`REJECT`, `ALL`.
        """
        return pulumi.get(self, "traffic_type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[Optional[str]]:
        """
        VPC ID to attach to
        """
        return pulumi.get(self, "vpc_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

