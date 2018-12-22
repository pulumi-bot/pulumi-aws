# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class Repository(pulumi.CustomResource):
    """
    Provides a CodeCommit Repository Resource.
    
    > **NOTE on CodeCommit Availability**: The CodeCommit is not yet rolled out
    in all regions - available regions are listed
    [the AWS Docs](https://docs.aws.amazon.com/general/latest/gr/rande.html#codecommit_region).
    """
    def __init__(__self__, __name__, __opts__=None, default_branch=None, description=None, repository_name=None):
        """Create a Repository resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['default_branch'] = default_branch

        __props__['description'] = description

        if not repository_name:
            raise TypeError('Missing required property repository_name')
        __props__['repository_name'] = repository_name

        __props__['arn'] = None
        __props__['clone_url_http'] = None
        __props__['clone_url_ssh'] = None
        __props__['repository_id'] = None

        super(Repository, __self__).__init__(
            'aws:codecommit/repository:Repository',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

