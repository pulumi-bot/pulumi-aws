# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class IdentityPool(pulumi.CustomResource):
    allow_unauthenticated_identities: pulumi.Output[bool]
    """
    Whether the identity pool supports unauthenticated logins or not.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the identity pool.
    """
    cognito_identity_providers: pulumi.Output[list]
    """
    An array of Amazon Cognito Identity user pools and their client IDs.

      * `client_id` (`str`) - The client ID for the Amazon Cognito Identity User Pool.
      * `provider_name` (`str`) - The provider name for an Amazon Cognito Identity User Pool.
      * `serverSideTokenCheck` (`bool`) - Whether server-side token validation is enabled for the identity provider’s token or not.
    """
    developer_provider_name: pulumi.Output[str]
    """
    The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
    backend and the Cognito service to communicate about the developer provider.
    """
    identity_pool_name: pulumi.Output[str]
    """
    The Cognito Identity Pool name.
    """
    openid_connect_provider_arns: pulumi.Output[list]
    """
    A list of OpendID Connect provider ARNs.
    """
    saml_provider_arns: pulumi.Output[list]
    """
    An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
    """
    supported_login_providers: pulumi.Output[dict]
    """
    Key-Value pairs mapping provider names to provider app IDs.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Identity Pool.
    """
    def __init__(__self__, resource_name, opts=None, allow_unauthenticated_identities=None, cognito_identity_providers=None, developer_provider_name=None, identity_pool_name=None, openid_connect_provider_arns=None, saml_provider_arns=None, supported_login_providers=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an AWS Cognito Identity Pool.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.iam.SamlProvider("default", saml_metadata_document=(lambda path: open(path).read())("saml-metadata.xml"))
        main = aws.cognito.IdentityPool("main",
            allow_unauthenticated_identities=False,
            cognito_identity_providers=[
                {
                    "clientId": "6lhlkkfbfb4q5kpp90urffae",
                    "providerName": "cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ",
                    "serverSideTokenCheck": False,
                },
                {
                    "clientId": "7kodkvfqfb4qfkp39eurffae",
                    "providerName": "cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu",
                    "serverSideTokenCheck": False,
                },
            ],
            identity_pool_name="identity pool",
            openid_connect_provider_arns=["arn:aws:iam::123456789012:oidc-provider/foo.example.com"],
            saml_provider_arns=[default.arn],
            supported_login_providers={
                "accounts.google.com": "123456789012.apps.googleusercontent.com",
                "graph.facebook.com": "7346241598935552",
            })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_unauthenticated_identities: Whether the identity pool supports unauthenticated logins or not.
        :param pulumi.Input[list] cognito_identity_providers: An array of Amazon Cognito Identity user pools and their client IDs.
        :param pulumi.Input[str] developer_provider_name: The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
               backend and the Cognito service to communicate about the developer provider.
        :param pulumi.Input[str] identity_pool_name: The Cognito Identity Pool name.
        :param pulumi.Input[list] openid_connect_provider_arns: A list of OpendID Connect provider ARNs.
        :param pulumi.Input[list] saml_provider_arns: An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
        :param pulumi.Input[dict] supported_login_providers: Key-Value pairs mapping provider names to provider app IDs.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Identity Pool.

        The **cognito_identity_providers** object supports the following:

          * `client_id` (`pulumi.Input[str]`) - The client ID for the Amazon Cognito Identity User Pool.
          * `provider_name` (`pulumi.Input[str]`) - The provider name for an Amazon Cognito Identity User Pool.
          * `serverSideTokenCheck` (`pulumi.Input[bool]`) - Whether server-side token validation is enabled for the identity provider’s token or not.
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

            __props__['allow_unauthenticated_identities'] = allow_unauthenticated_identities
            __props__['cognito_identity_providers'] = cognito_identity_providers
            __props__['developer_provider_name'] = developer_provider_name
            if identity_pool_name is None:
                raise TypeError("Missing required property 'identity_pool_name'")
            __props__['identity_pool_name'] = identity_pool_name
            __props__['openid_connect_provider_arns'] = openid_connect_provider_arns
            __props__['saml_provider_arns'] = saml_provider_arns
            __props__['supported_login_providers'] = supported_login_providers
            __props__['tags'] = tags
            __props__['arn'] = None
        super(IdentityPool, __self__).__init__(
            'aws:cognito/identityPool:IdentityPool',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allow_unauthenticated_identities=None, arn=None, cognito_identity_providers=None, developer_provider_name=None, identity_pool_name=None, openid_connect_provider_arns=None, saml_provider_arns=None, supported_login_providers=None, tags=None):
        """
        Get an existing IdentityPool resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_unauthenticated_identities: Whether the identity pool supports unauthenticated logins or not.
        :param pulumi.Input[str] arn: The ARN of the identity pool.
        :param pulumi.Input[list] cognito_identity_providers: An array of Amazon Cognito Identity user pools and their client IDs.
        :param pulumi.Input[str] developer_provider_name: The "domain" by which Cognito will refer to your users. This name acts as a placeholder that allows your
               backend and the Cognito service to communicate about the developer provider.
        :param pulumi.Input[str] identity_pool_name: The Cognito Identity Pool name.
        :param pulumi.Input[list] openid_connect_provider_arns: A list of OpendID Connect provider ARNs.
        :param pulumi.Input[list] saml_provider_arns: An array of Amazon Resource Names (ARNs) of the SAML provider for your identity.
        :param pulumi.Input[dict] supported_login_providers: Key-Value pairs mapping provider names to provider app IDs.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Identity Pool.

        The **cognito_identity_providers** object supports the following:

          * `client_id` (`pulumi.Input[str]`) - The client ID for the Amazon Cognito Identity User Pool.
          * `provider_name` (`pulumi.Input[str]`) - The provider name for an Amazon Cognito Identity User Pool.
          * `serverSideTokenCheck` (`pulumi.Input[bool]`) - Whether server-side token validation is enabled for the identity provider’s token or not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_unauthenticated_identities"] = allow_unauthenticated_identities
        __props__["arn"] = arn
        __props__["cognito_identity_providers"] = cognito_identity_providers
        __props__["developer_provider_name"] = developer_provider_name
        __props__["identity_pool_name"] = identity_pool_name
        __props__["openid_connect_provider_arns"] = openid_connect_provider_arns
        __props__["saml_provider_arns"] = saml_provider_arns
        __props__["supported_login_providers"] = supported_login_providers
        __props__["tags"] = tags
        return IdentityPool(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

