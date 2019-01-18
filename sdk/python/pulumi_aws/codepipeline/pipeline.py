# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Pipeline(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The codepipeline ARN.
    """
    artifact_store: pulumi.Output[dict]
    """
    An artifact_store block. Artifact stores are documented below.
    * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
    """
    name: pulumi.Output[str]
    """
    The name of the pipeline.
    """
    role_arn: pulumi.Output[str]
    """
    A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
    """
    stages: pulumi.Output[list]
    def __init__(__self__, __name__, __opts__=None, artifact_store=None, name=None, role_arn=None, stages=None):
        """
        Provides a CodePipeline.
        
        > **NOTE on `aws_codepipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] artifact_store: An artifact_store block. Artifact stores are documented below.
               * `stage` (Minimum of at least two `stage` blocks is required) A stage block. Stages are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[list] stages
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not artifact_store:
            raise TypeError('Missing required property artifact_store')
        __props__['artifact_store'] = artifact_store

        __props__['name'] = name

        if not role_arn:
            raise TypeError('Missing required property role_arn')
        __props__['role_arn'] = role_arn

        if not stages:
            raise TypeError('Missing required property stages')
        __props__['stages'] = stages

        __props__['arn'] = None

        super(Pipeline, __self__).__init__(
            'aws:codepipeline/pipeline:Pipeline',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

