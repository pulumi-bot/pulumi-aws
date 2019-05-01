# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Webhook(pulumi.CustomResource):
    branch_filter: pulumi.Output[str]
    """
    A regular expression used to determine which branches get built. Default is all branches are built.
    """
    payload_url: pulumi.Output[str]
    """
    The CodeBuild endpoint where webhook events are sent.
    """
    project_name: pulumi.Output[str]
    """
    The name of the build project.
    """
    secret: pulumi.Output[str]
    """
    The secret token of the associated repository. Not returned by the CodeBuild API for all source types.
    """
    url: pulumi.Output[str]
    """
    The URL to the webhook.
    """
    def __init__(__self__, resource_name, opts=None, branch_filter=None, project_name=None, __name__=None, __opts__=None):
        """
        Manages a CodeBuild webhook, which is an endpoint accepted by the CodeBuild service to trigger builds from source code repositories. Depending on the source type of the CodeBuild project, the CodeBuild service may also automatically create and delete the actual repository webhook as well.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] branch_filter: A regular expression used to determine which branches get built. Default is all branches are built.
        :param pulumi.Input[str] project_name: The name of the build project.
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

        __props__['branch_filter'] = branch_filter

        if project_name is None:
            raise TypeError("Missing required property 'project_name'")
        __props__['project_name'] = project_name

        __props__['payload_url'] = None
        __props__['secret'] = None
        __props__['url'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Webhook, __self__).__init__(
            'aws:codebuild/webhook:Webhook',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

