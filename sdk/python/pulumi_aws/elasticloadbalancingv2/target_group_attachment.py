# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['TargetGroupAttachmentArgs', 'TargetGroupAttachment']

@pulumi.input_type
class TargetGroupAttachmentArgs:
    def __init__(__self__, *,
                 target_group_arn: pulumi.Input[str],
                 target_id: pulumi.Input[str],
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a TargetGroupAttachment resource.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
        :param pulumi.Input[int] port: The port on which targets receive traffic.
        """
        pulumi.set(__self__, "target_group_arn", target_group_arn)
        pulumi.set(__self__, "target_id", target_id)
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if port is not None:
            pulumi.set(__self__, "port", port)

    @property
    @pulumi.getter(name="targetGroupArn")
    def target_group_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the target group with which to register targets
        """
        return pulumi.get(self, "target_group_arn")

    @target_group_arn.setter
    def target_group_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_group_arn", value)

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Input[str]:
        """
        The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        return pulumi.get(self, "target_id")

    @target_id.setter
    def target_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_id", value)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port on which targets receive traffic.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)


warnings.warn("""aws.elasticloadbalancingv2.TargetGroupAttachment has been deprecated in favor of aws.lb.TargetGroupAttachment""", DeprecationWarning)


class TargetGroupAttachment(pulumi.CustomResource):
    warnings.warn("""aws.elasticloadbalancingv2.TargetGroupAttachment has been deprecated in favor of aws.lb.TargetGroupAttachment""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_arn: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the `elb.Attachment` resource.

        > **Note:** `alb.TargetGroupAttachment` is known as `lb.TargetGroupAttachment`. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_target_group = aws.lb.TargetGroup("testTargetGroup")
        # ... other configuration ...
        test_instance = aws.ec2.Instance("testInstance")
        # ... other configuration ...
        test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
            target_group_arn=test_target_group.arn,
            target_id=test_instance.id,
            port=80)
        ```
        ## Usage with lambda

        ```python
        import pulumi
        import pulumi_aws as aws

        test_target_group = aws.lb.TargetGroup("testTargetGroup", target_type="lambda")
        test_function = aws.lambda_.Function("testFunction")
        # ... other configuration ...
        with_lb = aws.lambda_.Permission("withLb",
            action="lambda:InvokeFunction",
            function=test_function.arn,
            principal="elasticloadbalancing.amazonaws.com",
            source_arn=test_target_group.arn)
        test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
            target_group_arn=test_target_group.arn,
            target_id=test_function.arn,
            opts=pulumi.ResourceOptions(depends_on=[with_lb]))
        ```

        ## Import

        Target Group Attachments cannot be imported.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
        :param pulumi.Input[int] port: The port on which targets receive traffic.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TargetGroupAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the `elb.Attachment` resource.

        > **Note:** `alb.TargetGroupAttachment` is known as `lb.TargetGroupAttachment`. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        test_target_group = aws.lb.TargetGroup("testTargetGroup")
        # ... other configuration ...
        test_instance = aws.ec2.Instance("testInstance")
        # ... other configuration ...
        test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
            target_group_arn=test_target_group.arn,
            target_id=test_instance.id,
            port=80)
        ```
        ## Usage with lambda

        ```python
        import pulumi
        import pulumi_aws as aws

        test_target_group = aws.lb.TargetGroup("testTargetGroup", target_type="lambda")
        test_function = aws.lambda_.Function("testFunction")
        # ... other configuration ...
        with_lb = aws.lambda_.Permission("withLb",
            action="lambda:InvokeFunction",
            function=test_function.arn,
            principal="elasticloadbalancing.amazonaws.com",
            source_arn=test_target_group.arn)
        test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
            target_group_arn=test_target_group.arn,
            target_id=test_function.arn,
            opts=pulumi.ResourceOptions(depends_on=[with_lb]))
        ```

        ## Import

        Target Group Attachments cannot be imported.

        :param str resource_name: The name of the resource.
        :param TargetGroupAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TargetGroupAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 target_group_arn: Optional[pulumi.Input[str]] = None,
                 target_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""TargetGroupAttachment is deprecated: aws.elasticloadbalancingv2.TargetGroupAttachment has been deprecated in favor of aws.lb.TargetGroupAttachment""")
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

            __props__['availability_zone'] = availability_zone
            __props__['port'] = port
            if target_group_arn is None and not opts.urn:
                raise TypeError("Missing required property 'target_group_arn'")
            __props__['target_group_arn'] = target_group_arn
            if target_id is None and not opts.urn:
                raise TypeError("Missing required property 'target_id'")
            __props__['target_id'] = target_id
        super(TargetGroupAttachment, __self__).__init__(
            'aws:elasticloadbalancingv2/targetGroupAttachment:TargetGroupAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            target_group_arn: Optional[pulumi.Input[str]] = None,
            target_id: Optional[pulumi.Input[str]] = None) -> 'TargetGroupAttachment':
        """
        Get an existing TargetGroupAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
        :param pulumi.Input[int] port: The port on which targets receive traffic.
        :param pulumi.Input[str] target_group_arn: The ARN of the target group with which to register targets
        :param pulumi.Input[str] target_id: The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["availability_zone"] = availability_zone
        __props__["port"] = port
        __props__["target_group_arn"] = target_group_arn
        __props__["target_id"] = target_id
        return TargetGroupAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[Optional[str]]:
        """
        The Availability Zone where the IP address of the target is to be registered. If the private ip address is outside of the VPC scope, this value must be set to 'all'.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port on which targets receive traffic.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="targetGroupArn")
    def target_group_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the target group with which to register targets
        """
        return pulumi.get(self, "target_group_arn")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> pulumi.Output[str]:
        """
        The ID of the target. This is the Instance ID for an instance, or the container ID for an ECS container. If the target type is ip, specify an IP address. If the target type is lambda, specify the arn of lambda.
        """
        return pulumi.get(self, "target_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

