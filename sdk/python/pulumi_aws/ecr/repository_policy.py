# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class RepositoryPolicy(pulumi.CustomResource):
    """
    Provides an ECR repository policy.
    
    Note that currently only one policy may be applied to a repository.
    
    ~> **NOTE on ECR Availability**: The EC2 Container Registry is not yet rolled out
    in all regions - available regions are listed
    [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#ecr_region).
    """
    def __init__(__self__, __name__, __opts__=None, policy=None, repository=None):
        """Create a RepositoryPolicy resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not policy:
            raise TypeError('Missing required property policy')
        __props__['policy'] = policy

        if not repository:
            raise TypeError('Missing required property repository')
        __props__['repository'] = repository

        __props__['registry_id'] = None

        super(RepositoryPolicy, __self__).__init__(
            'aws:ecr/repositoryPolicy:RepositoryPolicy',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

