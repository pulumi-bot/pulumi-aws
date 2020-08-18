# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'PipelineArtifactStore',
    'PipelineArtifactStoreEncryptionKey',
    'PipelineStage',
    'PipelineStageAction',
    'WebhookAuthenticationConfiguration',
    'WebhookFilter',
]

@pulumi.output_type
class PipelineArtifactStore(dict):
    @property
    @pulumi.getter(name="encryptionKey")
    def encryption_key(self) -> Optional['outputs.PipelineArtifactStoreEncryptionKey']:
        """
        The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
        """
        ...

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
        """
        ...

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the artifact store, such as Amazon S3
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PipelineArtifactStoreEncryptionKey(dict):
    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The KMS key ARN or ID
        """
        ...

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of key; currently only `KMS` is supported
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PipelineStage(dict):
    @property
    @pulumi.getter
    def actions(self) -> List['outputs.PipelineStageAction']:
        """
        The action(s) to include in the stage. Defined as an `action` block below
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the stage.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PipelineStageAction(dict):
    @property
    @pulumi.getter
    def category(self) -> str:
        """
        A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
        """
        ...

    @property
    @pulumi.getter
    def configuration(self) -> Optional[Mapping[str, str]]:
        """
        A map of the action declaration's configuration. Configurations options for action types and providers can be found in the [Pipeline Structure Reference](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements) and [Action Structure Reference](https://docs.aws.amazon.com/codepipeline/latest/userguide/action-reference.html) documentation.
        """
        ...

    @property
    @pulumi.getter(name="inputArtifacts")
    def input_artifacts(self) -> Optional[List[str]]:
        """
        A list of artifact names to be worked on.
        """
        ...

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The action declaration's name.
        """
        ...

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        """
        The namespace all output variables will be accessed from.
        """
        ...

    @property
    @pulumi.getter(name="outputArtifacts")
    def output_artifacts(self) -> Optional[List[str]]:
        """
        A list of artifact names to output. Output artifact names must be unique within a pipeline.
        """
        ...

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
        """
        ...

    @property
    @pulumi.getter
    def provider(self) -> str:
        """
        The provider of the service being called by the action. Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of AWS CodeDeploy, which would be specified as CodeDeploy.
        """
        ...

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        """
        The region in which to run the action.
        """
        ...

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[str]:
        """
        The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
        """
        ...

    @property
    @pulumi.getter(name="runOrder")
    def run_order(self) -> Optional[float]:
        """
        The order in which actions are run.
        """
        ...

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        A string that identifies the action type.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WebhookAuthenticationConfiguration(dict):
    @property
    @pulumi.getter(name="allowedIpRange")
    def allowed_ip_range(self) -> Optional[str]:
        """
        A valid CIDR block for `IP` filtering. Required for `IP`.
        """
        ...

    @property
    @pulumi.getter(name="secretToken")
    def secret_token(self) -> Optional[str]:
        """
        The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class WebhookFilter(dict):
    @property
    @pulumi.getter(name="jsonPath")
    def json_path(self) -> str:
        """
        The [JSON path](https://github.com/json-path/JsonPath) to filter on.
        """
        ...

    @property
    @pulumi.getter(name="matchEquals")
    def match_equals(self) -> str:
        """
        The value to match on (e.g. `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
        """
        ...

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


