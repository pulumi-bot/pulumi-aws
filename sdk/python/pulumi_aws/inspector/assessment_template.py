# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class AssessmentTemplate(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The template assessment ARN.
    """
    duration: pulumi.Output[float]
    """
    The duration of the inspector run.
    """
    name: pulumi.Output[str]
    """
    The name of the assessment template.
    """
    rules_package_arns: pulumi.Output[list]
    """
    The rules to be used during the run.
    """
    target_arn: pulumi.Output[str]
    """
    The assessment target ARN to attach the template to.
    """
    def __init__(__self__, resource_name, opts=None, duration=None, name=None, rules_package_arns=None, target_arn=None, __name__=None, __opts__=None):
        """
        Provides a Inspector assessment template
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] duration: The duration of the inspector run.
        :param pulumi.Input[str] name: The name of the assessment template.
        :param pulumi.Input[list] rules_package_arns: The rules to be used during the run.
        :param pulumi.Input[str] target_arn: The assessment target ARN to attach the template to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/inspector_assessment_template.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if duration is None:
            raise TypeError("Missing required property 'duration'")
        __props__['duration'] = duration
        __props__['name'] = name
        if rules_package_arns is None:
            raise TypeError("Missing required property 'rules_package_arns'")
        __props__['rules_package_arns'] = rules_package_arns
        if target_arn is None:
            raise TypeError("Missing required property 'target_arn'")
        __props__['target_arn'] = target_arn
        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(AssessmentTemplate, __self__).__init__(
            'aws:inspector/assessmentTemplate:AssessmentTemplate',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

