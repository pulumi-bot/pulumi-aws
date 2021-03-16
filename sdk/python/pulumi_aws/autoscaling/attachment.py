# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['AttachmentArgs', 'Attachment']

@pulumi.input_type
class AttachmentArgs:
    def __init__(__self__, *,
                 autoscaling_group_name: pulumi.Input[str],
                 alb_target_group_arn: Optional[pulumi.Input[str]] = None,
                 elb: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Attachment resource.
        :param pulumi.Input[str] autoscaling_group_name: Name of ASG to associate with the ELB.
        :param pulumi.Input[str] alb_target_group_arn: The ARN of an ALB Target Group.
        :param pulumi.Input[str] elb: The name of the ELB.
        """
        pulumi.set(__self__, "autoscaling_group_name", autoscaling_group_name)
        if alb_target_group_arn is not None:
            pulumi.set(__self__, "alb_target_group_arn", alb_target_group_arn)
        if elb is not None:
            pulumi.set(__self__, "elb", elb)

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> pulumi.Input[str]:
        """
        Name of ASG to associate with the ELB.
        """
        return pulumi.get(self, "autoscaling_group_name")

    @autoscaling_group_name.setter
    def autoscaling_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "autoscaling_group_name", value)

    @property
    @pulumi.getter(name="albTargetGroupArn")
    def alb_target_group_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of an ALB Target Group.
        """
        return pulumi.get(self, "alb_target_group_arn")

    @alb_target_group_arn.setter
    def alb_target_group_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alb_target_group_arn", value)

    @property
    @pulumi.getter
    def elb(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ELB.
        """
        return pulumi.get(self, "elb")

    @elb.setter
    def elb(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "elb", value)


class Attachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_target_group_arn: Optional[pulumi.Input[str]] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 elb: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Auto Scaling Attachment resource.

        > **NOTE on Auto Scaling Groups and ASG Attachments:** This provider currently provides
        both a standalone `autoscaling.Attachment` resource
        (describing an ASG attached to an ELB or ALB), and an `autoscaling.Group`
        with `load_balancers` and `target_group_arns` defined in-line. These two methods are not
        mutually-exclusive. If `autoscaling.Attachment` resources are used, either alone or with inline
        `load_balancers` or `target_group_arns`, the `autoscaling.Group` resource must be configured
        to [ignore changes](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to the `load_balancers` and `target_group_arns` arguments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new load balancer attachment
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=aws_autoscaling_group["asg"]["id"],
            elb=aws_elb["bar"]["id"])
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new ALB Target Group attachment
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=aws_autoscaling_group["asg"]["id"],
            alb_target_group_arn=aws_alb_target_group["test"]["arn"])
        ```
        ## With An AutoScaling Group Resource

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration ...
        asg = aws.autoscaling.Group("asg")
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=asg.id,
            elb=aws_elb["test"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_target_group_arn: The ARN of an ALB Target Group.
        :param pulumi.Input[str] autoscaling_group_name: Name of ASG to associate with the ELB.
        :param pulumi.Input[str] elb: The name of the ELB.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Auto Scaling Attachment resource.

        > **NOTE on Auto Scaling Groups and ASG Attachments:** This provider currently provides
        both a standalone `autoscaling.Attachment` resource
        (describing an ASG attached to an ELB or ALB), and an `autoscaling.Group`
        with `load_balancers` and `target_group_arns` defined in-line. These two methods are not
        mutually-exclusive. If `autoscaling.Attachment` resources are used, either alone or with inline
        `load_balancers` or `target_group_arns`, the `autoscaling.Group` resource must be configured
        to [ignore changes](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) to the `load_balancers` and `target_group_arns` arguments.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new load balancer attachment
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=aws_autoscaling_group["asg"]["id"],
            elb=aws_elb["bar"]["id"])
        ```

        ```python
        import pulumi
        import pulumi_aws as aws

        # Create a new ALB Target Group attachment
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=aws_autoscaling_group["asg"]["id"],
            alb_target_group_arn=aws_alb_target_group["test"]["arn"])
        ```
        ## With An AutoScaling Group Resource

        ```python
        import pulumi
        import pulumi_aws as aws

        # ... other configuration ...
        asg = aws.autoscaling.Group("asg")
        asg_attachment_bar = aws.autoscaling.Attachment("asgAttachmentBar",
            autoscaling_group_name=asg.id,
            elb=aws_elb["test"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param AttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_target_group_arn: Optional[pulumi.Input[str]] = None,
                 autoscaling_group_name: Optional[pulumi.Input[str]] = None,
                 elb: Optional[pulumi.Input[str]] = None,
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

            __props__['alb_target_group_arn'] = alb_target_group_arn
            if autoscaling_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'autoscaling_group_name'")
            __props__['autoscaling_group_name'] = autoscaling_group_name
            __props__['elb'] = elb
        super(Attachment, __self__).__init__(
            'aws:autoscaling/attachment:Attachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alb_target_group_arn: Optional[pulumi.Input[str]] = None,
            autoscaling_group_name: Optional[pulumi.Input[str]] = None,
            elb: Optional[pulumi.Input[str]] = None) -> 'Attachment':
        """
        Get an existing Attachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_target_group_arn: The ARN of an ALB Target Group.
        :param pulumi.Input[str] autoscaling_group_name: Name of ASG to associate with the ELB.
        :param pulumi.Input[str] elb: The name of the ELB.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alb_target_group_arn"] = alb_target_group_arn
        __props__["autoscaling_group_name"] = autoscaling_group_name
        __props__["elb"] = elb
        return Attachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="albTargetGroupArn")
    def alb_target_group_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of an ALB Target Group.
        """
        return pulumi.get(self, "alb_target_group_arn")

    @property
    @pulumi.getter(name="autoscalingGroupName")
    def autoscaling_group_name(self) -> pulumi.Output[str]:
        """
        Name of ASG to associate with the ELB.
        """
        return pulumi.get(self, "autoscaling_group_name")

    @property
    @pulumi.getter
    def elb(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the ELB.
        """
        return pulumi.get(self, "elb")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

