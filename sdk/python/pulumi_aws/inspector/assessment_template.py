# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['AssessmentTemplate']


class AssessmentTemplate(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The template assessment ARN.
    """

    duration: pulumi.Output[float] = pulumi.property("duration")
    """
    The duration of the inspector run.
    """

    name: pulumi.Output[str] = pulumi.property("name")
    """
    The name of the assessment template.
    """

    rules_package_arns: pulumi.Output[List[str]] = pulumi.property("rulesPackageArns")
    """
    The rules to be used during the run.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    Key-value map of tags for the Inspector assessment template.
    """

    target_arn: pulumi.Output[str] = pulumi.property("targetArn")
    """
    The assessment target ARN to attach the template to.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 duration: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rules_package_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 target_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Inspector assessment template

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.inspector.AssessmentTemplate("example",
            duration=3600,
            rules_package_arns=[
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
            ],
            target_arn=aws_inspector_assessment_target["example"]["arn"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] duration: The duration of the inspector run.
        :param pulumi.Input[str] name: The name of the assessment template.
        :param pulumi.Input[List[pulumi.Input[str]]] rules_package_arns: The rules to be used during the run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Inspector assessment template.
        :param pulumi.Input[str] target_arn: The assessment target ARN to attach the template to.
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

            if duration is None:
                raise TypeError("Missing required property 'duration'")
            __props__['duration'] = duration
            __props__['name'] = name
            if rules_package_arns is None:
                raise TypeError("Missing required property 'rules_package_arns'")
            __props__['rules_package_arns'] = rules_package_arns
            __props__['tags'] = tags
            if target_arn is None:
                raise TypeError("Missing required property 'target_arn'")
            __props__['target_arn'] = target_arn
            __props__['arn'] = None
        super(AssessmentTemplate, __self__).__init__(
            'aws:inspector/assessmentTemplate:AssessmentTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            duration: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rules_package_arns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            target_arn: Optional[pulumi.Input[str]] = None) -> 'AssessmentTemplate':
        """
        Get an existing AssessmentTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The template assessment ARN.
        :param pulumi.Input[float] duration: The duration of the inspector run.
        :param pulumi.Input[str] name: The name of the assessment template.
        :param pulumi.Input[List[pulumi.Input[str]]] rules_package_arns: The rules to be used during the run.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of tags for the Inspector assessment template.
        :param pulumi.Input[str] target_arn: The assessment target ARN to attach the template to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["duration"] = duration
        __props__["name"] = name
        __props__["rules_package_arns"] = rules_package_arns
        __props__["tags"] = tags
        __props__["target_arn"] = target_arn
        return AssessmentTemplate(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

