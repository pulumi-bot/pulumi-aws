# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Webhook']


class Webhook(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 branch_filter: Optional[pulumi.Input[str]] = None,
                 filter_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.

        ## Example Usage
        ### Bitbucket and GitHub

        When working with [Bitbucket](https://bitbucket.org) and [GitHub](https://github.com) source CodeBuild webhooks, the CodeBuild service will automatically create (on `codebuild.Webhook` resource creation) and delete (on `codebuild.Webhook` resource deletion) the Bitbucket/GitHub repository webhook using its granted OAuth permissions. This behavior cannot be controlled by this provider.

        > **Note:** The AWS account that this provider uses to create this resource *must* have authorized CodeBuild to access Bitbucket/GitHub's OAuth API in each applicable region. This is a manual step that must be done *before* creating webhooks with this resource. If OAuth is not configured, AWS will return an error similar to `ResourceNotFoundException: Could not find access token for server type github`. More information can be found in the CodeBuild User Guide for [Bitbucket](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-bitbucket-pull-request.html) and [GitHub](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-pull-request.html).

        > **Note:** Further managing the automatically created Bitbucket/GitHub webhook with the `bitbucket_hook`/`github_repository_webhook` resource is only possible with importing that resource after creation of the `codebuild.Webhook` resource. The CodeBuild API does not ever provide the `secret` attribute for the `codebuild.Webhook` resource in this scenario.

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.codebuild.Webhook("example",
            project_name=aws_codebuild_project["example"]["name"],
            filter_groups=[aws.codebuild.WebhookFilterGroupArgs(
                filters=[
                    aws.codebuild.WebhookFilterGroupFilterArgs(
                        type="EVENT",
                        pattern="PUSH",
                    ),
                    aws.codebuild.WebhookFilterGroupFilterArgs(
                        type="HEAD_REF",
                        pattern="master",
                    ),
                ],
            )])
        ```
        ### GitHub Enterprise

        When working with [GitHub Enterprise](https://enterprise.github.com/) source CodeBuild webhooks, the GHE repository webhook must be separately managed (e.g. manually or with the `github_repository_webhook` resource).

        More information creating webhooks with GitHub Enterprise can be found in the [CodeBuild User Guide](https://docs.aws.amazon.com/codebuild/latest/userguide/sample-github-enterprise.html).

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_github as github

        example_webhook = aws.codebuild.Webhook("exampleWebhook", project_name=aws_codebuild_project["example"]["name"])
        example_repository_webhook = github.RepositoryWebhook("exampleRepositoryWebhook",
            active=True,
            events=["push"],
            repository=github_repository["example"]["name"],
            configuration={
                "url": example_webhook.payload_url,
                "secret": example_webhook.secret,
                "contentType": "json",
                "insecureSsl": False,
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch_filter: A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filter_group` over `branch_filter`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]] filter_groups: Information about the webhook's trigger. Filter group blocks are documented below.
        :param pulumi.Input[str] project_name: The name of the build project.
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

            __props__['branch_filter'] = branch_filter
            __props__['filter_groups'] = filter_groups
            if project_name is None:
                raise TypeError("Missing required property 'project_name'")
            __props__['project_name'] = project_name
            __props__['payload_url'] = None
            __props__['secret'] = None
            __props__['url'] = None
        super(Webhook, __self__).__init__(
            'aws:codebuild/webhook:Webhook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            branch_filter: Optional[pulumi.Input[str]] = None,
            filter_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]]] = None,
            payload_url: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            secret: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None) -> 'Webhook':
        """
        Get an existing Webhook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch_filter: A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filter_group` over `branch_filter`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WebhookFilterGroupArgs']]]] filter_groups: Information about the webhook's trigger. Filter group blocks are documented below.
        :param pulumi.Input[str] payload_url: The CodeBuild endpoint where webhook events are sent.
        :param pulumi.Input[str] project_name: The name of the build project.
        :param pulumi.Input[str] secret: The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
        :param pulumi.Input[str] url: The URL to the webhook.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["branch_filter"] = branch_filter
        __props__["filter_groups"] = filter_groups
        __props__["payload_url"] = payload_url
        __props__["project_name"] = project_name
        __props__["secret"] = secret
        __props__["url"] = url
        return Webhook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="branchFilter")
    def branch_filter(self) -> pulumi.Output[Optional[str]]:
        """
        A regular expression used to determine which branches get built. Default is all branches are built. It is recommended to use `filter_group` over `branch_filter`.
        """
        return pulumi.get(self, "branch_filter")

    @property
    @pulumi.getter(name="filterGroups")
    def filter_groups(self) -> pulumi.Output[Optional[Sequence['outputs.WebhookFilterGroup']]]:
        """
        Information about the webhook's trigger. Filter group blocks are documented below.
        """
        return pulumi.get(self, "filter_groups")

    @property
    @pulumi.getter(name="payloadUrl")
    def payload_url(self) -> pulumi.Output[str]:
        """
        The CodeBuild endpoint where webhook events are sent.
        """
        return pulumi.get(self, "payload_url")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[str]:
        """
        The name of the build project.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def secret(self) -> pulumi.Output[str]:
        """
        The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
        """
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL to the webhook.
        """
        return pulumi.get(self, "url")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

