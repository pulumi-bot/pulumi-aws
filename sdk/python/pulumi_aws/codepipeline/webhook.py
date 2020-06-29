# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Webhook(pulumi.CustomResource):
    authentication: pulumi.Output[str]
    """
    The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
    """
    authentication_configuration: pulumi.Output[dict]
    """
    An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.

      * `allowedIpRange` (`str`) - A valid CIDR block for `IP` filtering. Required for `IP`.
      * `secretToken` (`str`) - The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.
    """
    filters: pulumi.Output[list]
    """
    One or more `filter` blocks. Filter blocks are documented below.

      * `jsonPath` (`str`) - The [JSON path](https://github.com/json-path/JsonPath) to filter on.
      * `matchEquals` (`str`) - The value to match on (e.g. `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
    """
    name: pulumi.Output[str]
    """
    The name of the webhook.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    target_action: pulumi.Output[str]
    """
    The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
    """
    target_pipeline: pulumi.Output[str]
    """
    The name of the pipeline.
    """
    url: pulumi.Output[str]
    """
    The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.
    """
    def __init__(__self__, resource_name, opts=None, authentication=None, authentication_configuration=None, filters=None, name=None, tags=None, target_action=None, target_pipeline=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a CodePipeline Webhook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_github as github

        bar_pipeline = aws.codepipeline.Pipeline("barPipeline",
            artifact_store={
                "encryption_key": {
                    "id": data["aws_kms_alias"]["s3kmskey"]["arn"],
                    "type": "KMS",
                },
                "location": aws_s3_bucket["bar"]["bucket"],
                "type": "S3",
            },
            role_arn=aws_iam_role["bar"]["arn"],
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
                        "outputArtifacts": ["test"],
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
                        "inputArtifacts": ["test"],
                        "name": "Build",
                        "owner": "AWS",
                        "provider": "CodeBuild",
                        "version": "1",
                    }],
                    "name": "Build",
                },
            ])
        webhook_secret = "super-secret"
        bar_webhook = aws.codepipeline.Webhook("barWebhook",
            authentication="GITHUB_HMAC",
            authentication_configuration={
                "secretToken": webhook_secret,
            },
            filters=[{
                "jsonPath": "$.ref",
                "matchEquals": "refs/heads/{Branch}",
            }],
            target_action="Source",
            target_pipeline=bar_pipeline.name)
        # Wire the CodePipeline webhook into a GitHub repository.
        bar_repository_webhook = github.RepositoryWebhook("barRepositoryWebhook",
            configuration={
                "contentType": "json",
                "insecureSsl": True,
                "secret": webhook_secret,
                "url": bar_webhook.url,
            },
            events=["push"],
            repository=github_repository["repo"]["name"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authentication: The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
        :param pulumi.Input[dict] authentication_configuration: An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
        :param pulumi.Input[list] filters: One or more `filter` blocks. Filter blocks are documented below.
        :param pulumi.Input[str] name: The name of the webhook.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] target_action: The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
        :param pulumi.Input[str] target_pipeline: The name of the pipeline.

        The **authentication_configuration** object supports the following:

          * `allowedIpRange` (`pulumi.Input[str]`) - A valid CIDR block for `IP` filtering. Required for `IP`.
          * `secretToken` (`pulumi.Input[str]`) - The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.

        The **filters** object supports the following:

          * `jsonPath` (`pulumi.Input[str]`) - The [JSON path](https://github.com/json-path/JsonPath) to filter on.
          * `matchEquals` (`pulumi.Input[str]`) - The value to match on (e.g. `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
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

            if authentication is None:
                raise TypeError("Missing required property 'authentication'")
            __props__['authentication'] = authentication
            __props__['authentication_configuration'] = authentication_configuration
            if filters is None:
                raise TypeError("Missing required property 'filters'")
            __props__['filters'] = filters
            __props__['name'] = name
            __props__['tags'] = tags
            if target_action is None:
                raise TypeError("Missing required property 'target_action'")
            __props__['target_action'] = target_action
            if target_pipeline is None:
                raise TypeError("Missing required property 'target_pipeline'")
            __props__['target_pipeline'] = target_pipeline
            __props__['url'] = None
        super(Webhook, __self__).__init__(
            'aws:codepipeline/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, authentication=None, authentication_configuration=None, filters=None, name=None, tags=None, target_action=None, target_pipeline=None, url=None):
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authentication: The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
        :param pulumi.Input[dict] authentication_configuration: An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
        :param pulumi.Input[list] filters: One or more `filter` blocks. Filter blocks are documented below.
        :param pulumi.Input[str] name: The name of the webhook.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] target_action: The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
        :param pulumi.Input[str] target_pipeline: The name of the pipeline.
        :param pulumi.Input[str] url: The CodePipeline webhook's URL. POST events to this endpoint to trigger the target.

        The **authentication_configuration** object supports the following:

          * `allowedIpRange` (`pulumi.Input[str]`) - A valid CIDR block for `IP` filtering. Required for `IP`.
          * `secretToken` (`pulumi.Input[str]`) - The shared secret for the GitHub repository webhook. Set this as `secret` in your `github_repository_webhook`'s `configuration` block. Required for `GITHUB_HMAC`.

        The **filters** object supports the following:

          * `jsonPath` (`pulumi.Input[str]`) - The [JSON path](https://github.com/json-path/JsonPath) to filter on.
          * `matchEquals` (`pulumi.Input[str]`) - The value to match on (e.g. `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["authentication"] = authentication
        __props__["authentication_configuration"] = authentication_configuration
        __props__["filters"] = filters
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["target_action"] = target_action
        __props__["target_pipeline"] = target_pipeline
        __props__["url"] = url
        return Webhook(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
