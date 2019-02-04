# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class IdentityProvider(pulumi.CustomResource):
    attribute_mapping: pulumi.Output[dict]
    """
    The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
    """
    idp_identifiers: pulumi.Output[list]
    """
    The list of identity providers.
    """
    provider_details: pulumi.Output[dict]
    """
    The map of identity details, such as access token
    """
    provider_name: pulumi.Output[str]
    """
    The provider name
    """
    provider_type: pulumi.Output[str]
    """
    The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
    """
    user_pool_id: pulumi.Output[str]
    """
    The user pool id
    """
    def __init__(__self__, resource_name, opts=None, attribute_mapping=None, idp_identifiers=None, provider_details=None, provider_name=None, provider_type=None, user_pool_id=None, __name__=None, __opts__=None):
        """
        Provides a Cognito User Identity Provider resource.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] attribute_mapping: The map of attribute mapping of user pool attributes. [AttributeMapping in AWS API documentation](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-AttributeMapping)
        :param pulumi.Input[list] idp_identifiers: The list of identity providers.
        :param pulumi.Input[dict] provider_details: The map of identity details, such as access token
        :param pulumi.Input[str] provider_name: The provider name
        :param pulumi.Input[str] provider_type: The provider type.  [See AWS API for valid values](https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_CreateIdentityProvider.html#CognitoUserPools-CreateIdentityProvider-request-ProviderType)
        :param pulumi.Input[str] user_pool_id: The user pool id
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

        __props__['attribute_mapping'] = attribute_mapping

        __props__['idp_identifiers'] = idp_identifiers

        if not provider_details:
            raise TypeError('Missing required property provider_details')
        __props__['provider_details'] = provider_details

        if not provider_name:
            raise TypeError('Missing required property provider_name')
        __props__['provider_name'] = provider_name

        if not provider_type:
            raise TypeError('Missing required property provider_type')
        __props__['provider_type'] = provider_type

        if not user_pool_id:
            raise TypeError('Missing required property user_pool_id')
        __props__['user_pool_id'] = user_pool_id

        super(IdentityProvider, __self__).__init__(
            'aws:cognito/identityProvider:IdentityProvider',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

