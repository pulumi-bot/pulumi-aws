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

__all__ = ['GraphQLApi']


class GraphQLApi(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
                 authentication_type: Optional[pulumi.Input[str]] = None,
                 log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 openid_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']]] = None,
                 schema: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
                 xray_enabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an AppSync GraphQL API.

        ## Example Usage
        ### API Key Authentication

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example", authentication_type="API_KEY")
        ```
        ### AWS Cognito User Pool Authentication

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example",
            authentication_type="AMAZON_COGNITO_USER_POOLS",
            user_pool_config=aws.appsync.GraphQLApiUserPoolConfigArgs(
                aws_region=data["aws_region"]["current"]["name"],
                default_action="DENY",
                user_pool_id=aws_cognito_user_pool["example"]["id"],
            ))
        ```
        ### AWS IAM Authentication

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example", authentication_type="AWS_IAM")
        ```
        ### With Schema

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example",
            authentication_type="AWS_IAM",
            schema=\"\"\"schema {
        	query: Query
        }
        type Query {
          test: Int
        }

        \"\"\")
        ```
        ### OpenID Connect Authentication

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example",
            authentication_type="OPENID_CONNECT",
            openid_connect_config=aws.appsync.GraphQLApiOpenidConnectConfigArgs(
                issuer="https://example.com",
            ))
        ```
        ### With Multiple Authentication Providers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example",
            additional_authentication_providers=[aws.appsync.GraphQLApiAdditionalAuthenticationProviderArgs(
                authentication_type="AWS_IAM",
            )],
            authentication_type="API_KEY")
        ```
        ### Enabling Logging

        ```python
        import pulumi
        import pulumi_aws as aws

        example_role = aws.iam.Role("exampleRole", assume_role_policy=\"\"\"{
            "Version": "2012-10-17",
            "Statement": [
                {
                "Effect": "Allow",
                "Principal": {
                    "Service": "appsync.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
                }
            ]
        }
        \"\"\")
        example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
            policy_arn="arn:aws:iam::aws:policy/service-role/AWSAppSyncPushToCloudWatchLogs",
            role=example_role.name)
        # ... other configuration ...
        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", log_config=aws.appsync.GraphQLApiLogConfigArgs(
            cloudwatch_logs_role_arn=example_role.arn,
            field_log_level="ERROR",
        ))
        ```

        ## Import

        AppSync GraphQL API can be imported using the GraphQL API ID, e.g.

        ```sh
         $ pulumi import aws:appsync/graphQLApi:GraphQLApi example 0123456789
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] authentication_type: The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        :param pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: A user-supplied name for the GraphqlApi.
        :param pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']] user_pool_config: The Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
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

            __props__['additional_authentication_providers'] = additional_authentication_providers
            if authentication_type is None:
                raise TypeError("Missing required property 'authentication_type'")
            __props__['authentication_type'] = authentication_type
            __props__['log_config'] = log_config
            __props__['name'] = name
            __props__['openid_connect_config'] = openid_connect_config
            __props__['schema'] = schema
            __props__['tags'] = tags
            __props__['user_pool_config'] = user_pool_config
            __props__['xray_enabled'] = xray_enabled
            __props__['arn'] = None
            __props__['uris'] = None
        super(GraphQLApi, __self__).__init__(
            'aws:appsync/graphQLApi:GraphQLApi',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_authentication_providers: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            authentication_type: Optional[pulumi.Input[str]] = None,
            log_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            openid_connect_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']]] = None,
            schema: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            uris: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            user_pool_config: Optional[pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']]] = None,
            xray_enabled: Optional[pulumi.Input[bool]] = None) -> 'GraphQLApi':
        """
        Get an existing GraphQLApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GraphQLApiAdditionalAuthenticationProviderArgs']]]] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] arn: The ARN
        :param pulumi.Input[str] authentication_type: The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        :param pulumi.Input[pulumi.InputType['GraphQLApiLogConfigArgs']] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: A user-supplied name for the GraphqlApi.
        :param pulumi.Input[pulumi.InputType['GraphQLApiOpenidConnectConfigArgs']] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] uris: Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        :param pulumi.Input[pulumi.InputType['GraphQLApiUserPoolConfigArgs']] user_pool_config: The Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_authentication_providers"] = additional_authentication_providers
        __props__["arn"] = arn
        __props__["authentication_type"] = authentication_type
        __props__["log_config"] = log_config
        __props__["name"] = name
        __props__["openid_connect_config"] = openid_connect_config
        __props__["schema"] = schema
        __props__["tags"] = tags
        __props__["uris"] = uris
        __props__["user_pool_config"] = user_pool_config
        __props__["xray_enabled"] = xray_enabled
        return GraphQLApi(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalAuthenticationProviders")
    def additional_authentication_providers(self) -> pulumi.Output[Optional[Sequence['outputs.GraphQLApiAdditionalAuthenticationProvider']]]:
        """
        One or more additional authentication providers for the GraphqlApi. Defined below.
        """
        return pulumi.get(self, "additional_authentication_providers")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="authenticationType")
    def authentication_type(self) -> pulumi.Output[str]:
        """
        The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        """
        return pulumi.get(self, "authentication_type")

    @property
    @pulumi.getter(name="logConfig")
    def log_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiLogConfig']]:
        """
        Nested argument containing logging configuration. Defined below.
        """
        return pulumi.get(self, "log_config")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A user-supplied name for the GraphqlApi.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="openidConnectConfig")
    def openid_connect_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiOpenidConnectConfig']]:
        """
        Nested argument containing OpenID Connect configuration. Defined below.
        """
        return pulumi.get(self, "openid_connect_config")

    @property
    @pulumi.getter
    def schema(self) -> pulumi.Output[Optional[str]]:
        """
        The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        """
        return pulumi.get(self, "schema")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def uris(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        """
        return pulumi.get(self, "uris")

    @property
    @pulumi.getter(name="userPoolConfig")
    def user_pool_config(self) -> pulumi.Output[Optional['outputs.GraphQLApiUserPoolConfig']]:
        """
        The Amazon Cognito User Pool configuration. Defined below.
        """
        return pulumi.get(self, "user_pool_config")

    @property
    @pulumi.getter(name="xrayEnabled")
    def xray_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether tracing with X-ray is enabled. Defaults to false.
        """
        return pulumi.get(self, "xray_enabled")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

