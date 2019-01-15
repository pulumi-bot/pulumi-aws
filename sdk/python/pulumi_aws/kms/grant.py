# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class Grant(pulumi.CustomResource):
    """
    Provides a resource-based access control mechanism for a KMS customer master key.
    """
    def __init__(__self__, __name__, __opts__=None, constraints=None, grant_creation_tokens=None, grantee_principal=None, key_id=None, name=None, operations=None, retire_on_delete=None, retiring_principal=None):
        """Create a Grant resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['constraints'] = constraints

        __props__['grant_creation_tokens'] = grant_creation_tokens

        if not grantee_principal:
            raise TypeError('Missing required property grantee_principal')
        __props__['grantee_principal'] = grantee_principal

        if not key_id:
            raise TypeError('Missing required property key_id')
        __props__['key_id'] = key_id

        __props__['name'] = name

        if not operations:
            raise TypeError('Missing required property operations')
        __props__['operations'] = operations

        __props__['retire_on_delete'] = retire_on_delete

        __props__['retiring_principal'] = retiring_principal

        __props__['grant_id'] = None
        __props__['grant_token'] = None

        super(Grant, __self__).__init__(
            'aws:kms/grant:Grant',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

