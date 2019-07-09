# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Attachment(pulumi.CustomResource):
    alb_target_group_arn: pulumi.Output[str]
    """
    The ARN of an ALB Target Group.
    """
    autoscaling_group_name: pulumi.Output[str]
    """
    Name of ASG to associate with the ELB.
    """
    elb: pulumi.Output[str]
    """
    The name of the ELB.
    """
    def __init__(__self__, resource_name, opts=None, alb_target_group_arn=None, autoscaling_group_name=None, elb=None, __name__=None, __opts__=None):
        """
        Create a Attachment resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_target_group_arn: The ARN of an ALB Target Group.
        :param pulumi.Input[str] autoscaling_group_name: Name of ASG to associate with the ELB.
        :param pulumi.Input[str] elb: The name of the ELB.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/autoscaling_attachment.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['alb_target_group_arn'] = alb_target_group_arn

        if autoscaling_group_name is None:
            raise TypeError("Missing required property 'autoscaling_group_name'")
        __props__['autoscaling_group_name'] = autoscaling_group_name

        __props__['elb'] = elb

        super(Attachment, __self__).__init__(
            'aws:autoscaling/attachment:Attachment',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

