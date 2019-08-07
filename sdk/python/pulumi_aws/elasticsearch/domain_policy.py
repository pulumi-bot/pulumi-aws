# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class DomainPolicy(pulumi.CustomResource):
    access_policies: pulumi.Output[str]
    """
    IAM policy document specifying the access policies for the domain
    """
    domain_name: pulumi.Output[str]
    """
    Name of the domain.
    """
    def __init__(__self__, resource_name, opts=None, access_policies=None, domain_name=None, __name__=None, __opts__=None):
        """
        Allows setting policy to an Elasticsearch domain while referencing domain attributes (e.g. ARN)
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_policies: IAM policy document specifying the access policies for the domain
        :param pulumi.Input[str] domain_name: Name of the domain.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/elasticsearch_domain_policy.html.markdown.
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

        if access_policies is None:
            raise TypeError("Missing required property 'access_policies'")
        __props__['access_policies'] = access_policies
        if domain_name is None:
            raise TypeError("Missing required property 'domain_name'")
        __props__['domain_name'] = domain_name
        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        super(DomainPolicy, __self__).__init__(
            'aws:elasticsearch/domainPolicy:DomainPolicy',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

