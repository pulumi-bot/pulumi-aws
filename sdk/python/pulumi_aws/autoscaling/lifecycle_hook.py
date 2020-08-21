# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['LifecycleHook']


class LifecycleHook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 default_result: Optional[pulumi.Input[str]] = None,
                 heartbeat_timeout: Optional[pulumi.Input[float]] = None,
                 lifecycle_transition: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 notification_metadata: Optional[pulumi.Input[str]] = None,
                 notification_target_arn: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AutoScaling Lifecycle Hook resource.

        > **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
        the `initial_lifecycle_hook` attribute from the
        `autoscaling.Group`
        resource, or via this one. Hooks added via this resource will not be added
        until the autoscaling group has been created, and depending on your
        `capacity`
        settings, after the initial instances have been launched, creating unintended
        behavior. If you need hooks to run on all instances, add them with
        `initial_lifecycle_hook` in
        `autoscaling.Group`,
        but take care to not duplicate those hooks with this resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foobar_group = aws.autoscaling.Group("foobarGroup",
            availability_zones=["us-west-2a"],
            health_check_type="EC2",
            termination_policies=["OldestInstance"],
            tags=[{
                "key": "Foo",
                "value": "foo-bar",
                "propagateAtLaunch": True,
            }])
        foobar_lifecycle_hook = aws.autoscaling.LifecycleHook("foobarLifecycleHook",
            autoscaling_group_name=foobar_group.name,
            default_result="CONTINUE",
            heartbeat_timeout=2000,
            lifecycle_transition="autoscaling:EC2_INSTANCE_LAUNCHING",
            notification_metadata=\"\"\"{
          "foo": "bar"
        }
        \"\"\",
            notification_target_arn="arn:aws:sqs:us-east-1:444455556666:queue1*",
            role_arn="arn:aws:iam::123456789012:role/S3Access")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_group_name: The name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[float] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] name: The name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
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

            if autoscaling_group_name is None:
                raise TypeError("Missing required property 'autoscaling_group_name'")
            __props__['autoscaling_group_name'] = autoscaling_group_name
            __props__['default_result'] = default_result
            __props__['heartbeat_timeout'] = heartbeat_timeout
            if lifecycle_transition is None:
                raise TypeError("Missing required property 'lifecycle_transition'")
            __props__['lifecycle_transition'] = lifecycle_transition
            __props__['name'] = name
            __props__['notification_metadata'] = notification_metadata
            __props__['notification_target_arn'] = notification_target_arn
            __props__['role_arn'] = role_arn
        super(LifecycleHook, __self__).__init__(
            'aws:autoscaling/lifecycleHook:LifecycleHook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_group_name: Optional[pulumi.Input[str]] = None,
            default_result: Optional[pulumi.Input[str]] = None,
            heartbeat_timeout: Optional[pulumi.Input[float]] = None,
            lifecycle_transition: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            notification_metadata: Optional[pulumi.Input[str]] = None,
            notification_target_arn: Optional[pulumi.Input[str]] = None,
            role_arn: Optional[pulumi.Input[str]] = None) -> 'LifecycleHook':
        """
        Get an existing LifecycleHook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_group_name: The name of the Auto Scaling group to which you want to assign the lifecycle hook
        :param pulumi.Input[str] default_result: Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        :param pulumi.Input[float] heartbeat_timeout: Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        :param pulumi.Input[str] lifecycle_transition: The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        :param pulumi.Input[str] name: The name of the lifecycle hook.
        :param pulumi.Input[str] notification_metadata: Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        :param pulumi.Input[str] notification_target_arn: The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        :param pulumi.Input[str] role_arn: The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_group_name"] = autoscaling_group_name
        __props__["default_result"] = default_result
        __props__["heartbeat_timeout"] = heartbeat_timeout
        __props__["lifecycle_transition"] = lifecycle_transition
        __props__["name"] = name
        __props__["notification_metadata"] = notification_metadata
        __props__["notification_target_arn"] = notification_target_arn
        __props__["role_arn"] = role_arn
        return LifecycleHook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> str:
        """
        The name of the Auto Scaling group to which you want to assign the lifecycle hook
        """
        return pulumi.get(self, "autoscaling_group_name")

    @property
    @pulumi.getter(name="defaultResult")
    def default_result(self) -> str:
        """
        Defines the action the Auto Scaling group should take when the lifecycle hook timeout elapses or if an unexpected failure occurs. The value for this parameter can be either CONTINUE or ABANDON. The default value for this parameter is ABANDON.
        """
        return pulumi.get(self, "default_result")

    @property
    @pulumi.getter(name="heartbeatTimeout")
    def heartbeat_timeout(self) -> Optional[float]:
        """
        Defines the amount of time, in seconds, that can elapse before the lifecycle hook times out. When the lifecycle hook times out, Auto Scaling performs the action defined in the DefaultResult parameter
        """
        return pulumi.get(self, "heartbeat_timeout")

    @property
    @pulumi.getter(name="lifecycleTransition")
    def lifecycle_transition(self) -> str:
        """
        The instance state to which you want to attach the lifecycle hook. For a list of lifecycle hook types, see [describe-lifecycle-hook-types](https://docs.aws.amazon.com/cli/latest/reference/autoscaling/describe-lifecycle-hook-types.html#examples)
        """
        return pulumi.get(self, "lifecycle_transition")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the lifecycle hook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationMetadata")
    def notification_metadata(self) -> Optional[str]:
        """
        Contains additional information that you want to include any time Auto Scaling sends a message to the notification target.
        """
        return pulumi.get(self, "notification_metadata")

    @property
    @pulumi.getter(name="notificationTargetArn")
    def notification_target_arn(self) -> Optional[str]:
        """
        The ARN of the notification target that Auto Scaling will use to notify you when an instance is in the transition state for the lifecycle hook. This ARN target can be either an SQS queue or an SNS topic.
        """
        return pulumi.get(self, "notification_target_arn")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
        """
        return pulumi.get(self, "role_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

