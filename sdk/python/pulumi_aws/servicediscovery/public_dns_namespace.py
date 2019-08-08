# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class PublicDnsNamespace(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN that Amazon Route 53 assigns to the namespace when you create it.
    """
    description: pulumi.Output[str]
    """
    The description that you specify for the namespace when you create it.
    """
    hosted_zone: pulumi.Output[str]
    """
    The ID for the hosted zone that Amazon Route 53 creates when you create a namespace.
    """
    name: pulumi.Output[str]
    """
    The name of the namespace.
    """
    def __init__(__self__, resource_name, opts=None, description=None, name=None, __name__=None, __opts__=None):
        """
        Provides a Service Discovery Public DNS Namespace resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description that you specify for the namespace when you create it.
        :param pulumi.Input[str] name: The name of the namespace.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/service_discovery_public_dns_namespace.html.markdown.
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

        __props__['description'] = description
        __props__['name'] = name
        __props__['arn'] = None
        __props__['hosted_zone'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(PublicDnsNamespace, __self__).__init__(
            'aws:servicediscovery/publicDnsNamespace:PublicDnsNamespace',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

