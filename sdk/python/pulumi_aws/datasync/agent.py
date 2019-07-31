# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Agent(pulumi.CustomResource):
    activation_key: pulumi.Output[str]
    arn: pulumi.Output[str]
    """
    Amazon Resource Name (ARN) of the DataSync Agent.
    """
    ip_address: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    Name of the DataSync Agent.
    """
    tags: pulumi.Output[dict]
    """
    Key-value pairs of resource tags to assign to the DataSync Agent.
    """
    def __init__(__self__, resource_name, opts=None, activation_key=None, ip_address=None, name=None, tags=None, __name__=None, __opts__=None):
        """
        Create a Agent resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the DataSync Agent.
        :param pulumi.Input[dict] tags: Key-value pairs of resource tags to assign to the DataSync Agent.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/datasync_agent.html.markdown.
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

        __props__['activation_key'] = activation_key

        __props__['ip_address'] = ip_address

        __props__['name'] = name

        __props__['tags'] = tags

        __props__['arn'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(Agent, __self__).__init__(
            'aws:datasync/agent:Agent',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

