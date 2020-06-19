# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Pipeline(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The codepipeline ARN.
    """
    artifact_store: pulumi.Output[dict]
    """
    One or more artifact_store blocks. Artifact stores are documented below.

      * `encryption_key` (`dict`) - The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
        * `id` (`str`) - The KMS key ARN or ID
        * `type` (`str`) - The type of key; currently only `KMS` is supported

      * `location` (`str`) - The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
      * `region` (`str`) - The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
      * `type` (`str`) - The type of the artifact store, such as Amazon S3
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
    """
    A stage block. Stages are documented below.

      * `actions` (`list`) - The action(s) to include in the stage. Defined as an `action` block below
        * `category` (`str`) - A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
        * `configuration` (`dict`) - A Map of the action declaration's configuration. Find out more about configuring action configurations in the [Reference Pipeline Structure documentation](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements).
        * `inputArtifacts` (`list`) - A list of artifact names to be worked on.
        * `name` (`str`) - The action declaration's name.
        * `namespace` (`str`) - The namespace all output variables will be accessed from.
        * `outputArtifacts` (`list`) - A list of artifact names to output. Output artifact names must be unique within a pipeline.
        * `owner` (`str`) - The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
        * `provider` (`str`) - The provider of the service being called by the action. Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of AWS CodeDeploy, which would be specified as CodeDeploy.
        * `region` (`str`) - The region in which to run the action.
        * `role_arn` (`str`) - The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
        * `runOrder` (`float`) - The order in which actions are run.
        * `version` (`str`) - A string that identifies the action type.

      * `name` (`str`) - The name of the stage.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    def __init__(__self__, resource_name, opts=None, artifact_store=None, name=None, role_arn=None, stages=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CodePipeline.

        > **NOTE on `codepipeline.Pipeline`:** - the `GITHUB_TOKEN` environment variable must be set if the GitHub provider is specified.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        codepipeline_bucket = aws.s3.Bucket("codepipelineBucket", acl="private")
        codepipeline_role = aws.iam.Role("codepipelineRole", assume_role_policy=\"\"\"{
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": "codepipeline.amazonaws.com"
              },
              "Action": "sts:AssumeRole"
            }
          ]
        }

        \"\"\")
        codepipeline_policy = aws.iam.RolePolicy("codepipelinePolicy",
            policy=pulumi.Output.all(codepipeline_bucket.arn, codepipeline_bucket.arn).apply(lambda codepipelineBucketArn, codepipelineBucketArn1: f\"\"\"{{
          "Version": "2012-10-17",
          "Statement": [
            {{
              "Effect":"Allow",
              "Action": [
                "s3:GetObject",
                "s3:GetObjectVersion",
                "s3:GetBucketVersioning",
                "s3:PutObject"
              ],
              "Resource": [
                "{codepipeline_bucket_arn}",
                "{codepipeline_bucket_arn1}/*"
              ]
            }},
            {{
              "Effect": "Allow",
              "Action": [
                "codebuild:BatchGetBuilds",
                "codebuild:StartBuild"
              ],
              "Resource": "*"
            }}
          ]
        }}

        \"\"\"),
            role=codepipeline_role.id)
        s3kmskey = aws.kms.get_alias(name="alias/myKmsKey")
        codepipeline = aws.codepipeline.Pipeline("codepipeline",
            artifact_store={
                "encryption_key": {
                    "id": s3kmskey.arn,
                    "type": "KMS",
                },
                "location": codepipeline_bucket.bucket,
                "type": "S3",
            },
            role_arn=codepipeline_role.arn,
            stages=[
                {
                    "action": [{
                        "category": "Source",
                        "configuration": {
                            "Branch": "master",
                            "Owner": "my-organization",
                            "Repo": "test",
                        },
                        "name": "Source",
                        "outputArtifacts": ["source_output"],
                        "owner": "ThirdParty",
                        "provider": "GitHub",
                        "version": "1",
                    }],
                    "name": "Source",
                },
                {
                    "action": [{
                        "category": "Build",
                        "configuration": {
                            "ProjectName": "test",
                        },
                        "inputArtifacts": ["source_output"],
                        "name": "Build",
                        "outputArtifacts": ["build_output"],
                        "owner": "AWS",
                        "provider": "CodeBuild",
                        "version": "1",
                    }],
                    "name": "Build",
                },
                {
                    "action": [{
                        "category": "Deploy",
                        "configuration": {
                            "ActionMode": "REPLACE_ON_FAILURE",
                            "Capabilities": "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                            "OutputFileName": "CreateStackOutput.json",
                            "StackName": "MyStack",
                            "TemplatePath": "build_output::sam-templated.yaml",
                        },
                        "inputArtifacts": ["build_output"],
                        "name": "Deploy",
                        "owner": "AWS",
                        "provider": "CloudFormation",
                        "version": "1",
                    }],
                    "name": "Deploy",
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[list] stages: A stage block. Stages are documented below.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **artifact_store** object supports the following:

          * `encryption_key` (`pulumi.Input[dict]`) - The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
            * `id` (`pulumi.Input[str]`) - The KMS key ARN or ID
            * `type` (`pulumi.Input[str]`) - The type of key; currently only `KMS` is supported

          * `location` (`pulumi.Input[str]`) - The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
          * `region` (`pulumi.Input[str]`) - The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
          * `type` (`pulumi.Input[str]`) - The type of the artifact store, such as Amazon S3

        The **stages** object supports the following:

          * `actions` (`pulumi.Input[list]`) - The action(s) to include in the stage. Defined as an `action` block below
            * `category` (`pulumi.Input[str]`) - A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
            * `configuration` (`pulumi.Input[dict]`) - A Map of the action declaration's configuration. Find out more about configuring action configurations in the [Reference Pipeline Structure documentation](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements).
            * `inputArtifacts` (`pulumi.Input[list]`) - A list of artifact names to be worked on.
            * `name` (`pulumi.Input[str]`) - The action declaration's name.
            * `namespace` (`pulumi.Input[str]`) - The namespace all output variables will be accessed from.
            * `outputArtifacts` (`pulumi.Input[list]`) - A list of artifact names to output. Output artifact names must be unique within a pipeline.
            * `owner` (`pulumi.Input[str]`) - The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
            * `provider` (`pulumi.Input[str]`) - The provider of the service being called by the action. Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of AWS CodeDeploy, which would be specified as CodeDeploy.
            * `region` (`pulumi.Input[str]`) - The region in which to run the action.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
            * `runOrder` (`pulumi.Input[float]`) - The order in which actions are run.
            * `version` (`pulumi.Input[str]`) - A string that identifies the action type.

          * `name` (`pulumi.Input[str]`) - The name of the stage.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if artifact_store is None:
                raise TypeError("Missing required property 'artifact_store'")
            __props__['artifact_store'] = artifact_store
            __props__['name'] = name
            if role_arn is None:
                raise TypeError("Missing required property 'role_arn'")
            __props__['role_arn'] = role_arn
            if stages is None:
                raise TypeError("Missing required property 'stages'")
            __props__['stages'] = stages
            __props__['tags'] = tags
            __props__['arn'] = None
        super(Pipeline, __self__).__init__(
            'aws:codepipeline/pipeline:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, artifact_store=None, name=None, role_arn=None, stages=None, tags=None):
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The codepipeline ARN.
        :param pulumi.Input[dict] artifact_store: One or more artifact_store blocks. Artifact stores are documented below.
        :param pulumi.Input[str] name: The name of the pipeline.
        :param pulumi.Input[str] role_arn: A service role Amazon Resource Name (ARN) that grants AWS CodePipeline permission to make calls to AWS services on your behalf.
        :param pulumi.Input[list] stages: A stage block. Stages are documented below.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.

        The **artifact_store** object supports the following:

          * `encryption_key` (`pulumi.Input[dict]`) - The encryption key block AWS CodePipeline uses to encrypt the data in the artifact store, such as an AWS Key Management Service (AWS KMS) key. If you don't specify a key, AWS CodePipeline uses the default key for Amazon Simple Storage Service (Amazon S3). An `encryption_key` block is documented below.
            * `id` (`pulumi.Input[str]`) - The KMS key ARN or ID
            * `type` (`pulumi.Input[str]`) - The type of key; currently only `KMS` is supported

          * `location` (`pulumi.Input[str]`) - The location where AWS CodePipeline stores artifacts for a pipeline; currently only `S3` is supported.
          * `region` (`pulumi.Input[str]`) - The region where the artifact store is located. Required for a cross-region CodePipeline, do not provide for a single-region CodePipeline.
          * `type` (`pulumi.Input[str]`) - The type of the artifact store, such as Amazon S3

        The **stages** object supports the following:

          * `actions` (`pulumi.Input[list]`) - The action(s) to include in the stage. Defined as an `action` block below
            * `category` (`pulumi.Input[str]`) - A category defines what kind of action can be taken in the stage, and constrains the provider type for the action. Possible values are `Approval`, `Build`, `Deploy`, `Invoke`, `Source` and `Test`.
            * `configuration` (`pulumi.Input[dict]`) - A Map of the action declaration's configuration. Find out more about configuring action configurations in the [Reference Pipeline Structure documentation](http://docs.aws.amazon.com/codepipeline/latest/userguide/reference-pipeline-structure.html#action-requirements).
            * `inputArtifacts` (`pulumi.Input[list]`) - A list of artifact names to be worked on.
            * `name` (`pulumi.Input[str]`) - The action declaration's name.
            * `namespace` (`pulumi.Input[str]`) - The namespace all output variables will be accessed from.
            * `outputArtifacts` (`pulumi.Input[list]`) - A list of artifact names to output. Output artifact names must be unique within a pipeline.
            * `owner` (`pulumi.Input[str]`) - The creator of the action being called. Possible values are `AWS`, `Custom` and `ThirdParty`.
            * `provider` (`pulumi.Input[str]`) - The provider of the service being called by the action. Valid providers are determined by the action category. For example, an action in the Deploy category type might have a provider of AWS CodeDeploy, which would be specified as CodeDeploy.
            * `region` (`pulumi.Input[str]`) - The region in which to run the action.
            * `role_arn` (`pulumi.Input[str]`) - The ARN of the IAM service role that will perform the declared action. This is assumed through the roleArn for the pipeline.
            * `runOrder` (`pulumi.Input[float]`) - The order in which actions are run.
            * `version` (`pulumi.Input[str]`) - A string that identifies the action type.

          * `name` (`pulumi.Input[str]`) - The name of the stage.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["artifact_store"] = artifact_store
        __props__["name"] = name
        __props__["role_arn"] = role_arn
        __props__["stages"] = stages
        __props__["tags"] = tags
        return Pipeline(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
