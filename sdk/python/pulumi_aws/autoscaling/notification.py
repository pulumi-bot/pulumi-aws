# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['NotificationArgs', 'Notification']

@pulumi.input_type
class NotificationArgs:
    def __init__(__self__, *,
                 group_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 notifications: pulumi.Input[Sequence[pulumi.Input[str]]],
                 topic_arn: pulumi.Input[str]):
        """
        The set of arguments for constructing a Notification resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: A list of AutoScaling Group Names
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: A list of Notification Types that trigger
               notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
        :param pulumi.Input[str] topic_arn: The Topic ARN for notifications to be sent through
        """
        pulumi.set(__self__, "group_names", group_names)
        pulumi.set(__self__, "notifications", notifications)
        pulumi.set(__self__, "topic_arn", topic_arn)

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of AutoScaling Group Names
        """
        return pulumi.get(self, "group_names")

    @group_names.setter
    def group_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "group_names", value)

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of Notification Types that trigger
        notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
        """
        return pulumi.get(self, "notifications")

    @notifications.setter
    def notifications(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "notifications", value)

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> pulumi.Input[str]:
        """
        The Topic ARN for notifications to be sent through
        """
        return pulumi.get(self, "topic_arn")

    @topic_arn.setter
    def topic_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_arn", value)


class Notification(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
        the `notifications` map to a [Notification Configuration](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html) inside Amazon Web
        Services, and are applied to each AutoScaling Group you supply.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sns.Topic("example")
        # arn is an exported attribute
        bar = aws.autoscaling.Group("bar")
        # ...
        foo = aws.autoscaling.Group("foo")
        # ...
        example_notifications = aws.autoscaling.Notification("exampleNotifications",
            group_names=[
                bar.name,
                foo.name,
            ],
            notifications=[
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
            ],
            topic_arn=example.arn)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: A list of AutoScaling Group Names
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: A list of Notification Types that trigger
               notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
        :param pulumi.Input[str] topic_arn: The Topic ARN for notifications to be sent through
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NotificationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an AutoScaling Group with Notification support, via SNS Topics. Each of
        the `notifications` map to a [Notification Configuration](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_DescribeNotificationConfigurations.html) inside Amazon Web
        Services, and are applied to each AutoScaling Group you supply.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.sns.Topic("example")
        # arn is an exported attribute
        bar = aws.autoscaling.Group("bar")
        # ...
        foo = aws.autoscaling.Group("foo")
        # ...
        example_notifications = aws.autoscaling.Notification("exampleNotifications",
            group_names=[
                bar.name,
                foo.name,
            ],
            notifications=[
                "autoscaling:EC2_INSTANCE_LAUNCH",
                "autoscaling:EC2_INSTANCE_TERMINATE",
                "autoscaling:EC2_INSTANCE_LAUNCH_ERROR",
                "autoscaling:EC2_INSTANCE_TERMINATE_ERROR",
            ],
            topic_arn=example.arn)
        ```

        :param str resource_name: The name of the resource.
        :param NotificationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NotificationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 topic_arn: Optional[pulumi.Input[str]] = None,
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

            if group_names is None and not opts.urn:
                raise TypeError("Missing required property 'group_names'")
            __props__['group_names'] = group_names
            if notifications is None and not opts.urn:
                raise TypeError("Missing required property 'notifications'")
            __props__['notifications'] = notifications
            if topic_arn is None and not opts.urn:
                raise TypeError("Missing required property 'topic_arn'")
            __props__['topic_arn'] = topic_arn
        super(Notification, __self__).__init__(
            'aws:autoscaling/notification:Notification',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            notifications: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            topic_arn: Optional[pulumi.Input[str]] = None) -> 'Notification':
        """
        Get an existing Notification resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] group_names: A list of AutoScaling Group Names
        :param pulumi.Input[Sequence[pulumi.Input[str]]] notifications: A list of Notification Types that trigger
               notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
        :param pulumi.Input[str] topic_arn: The Topic ARN for notifications to be sent through
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["group_names"] = group_names
        __props__["notifications"] = notifications
        __props__["topic_arn"] = topic_arn
        return Notification(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of AutoScaling Group Names
        """
        return pulumi.get(self, "group_names")

    @property
    @pulumi.getter
    def notifications(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of Notification Types that trigger
        notifications. Acceptable values are documented [in the AWS documentation here](https://docs.aws.amazon.com/AutoScaling/latest/APIReference/API_NotificationConfiguration.html)
        """
        return pulumi.get(self, "notifications")

    @property
    @pulumi.getter(name="topicArn")
    def topic_arn(self) -> pulumi.Output[str]:
        """
        The Topic ARN for notifications to be sent through
        """
        return pulumi.get(self, "topic_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

