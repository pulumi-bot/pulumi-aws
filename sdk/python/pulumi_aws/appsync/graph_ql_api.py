# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GraphQLApi(pulumi.CustomResource):
    additional_authentication_providers: pulumi.Output[list]
    """
    One or more additional authentication providers for the GraphqlApi. Defined below.

      * `authentication_type` (`str`) - The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
      * `openid_connect_config` (`dict`) - Nested argument containing OpenID Connect configuration. Defined below.
        * `authTtl` (`float`) - Number of milliseconds a token is valid after being authenticated.
        * `client_id` (`str`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
        * `iatTtl` (`float`) - Number of milliseconds a token is valid after being issued to a user.
        * `issuer` (`str`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.

      * `user_pool_config` (`dict`) - The Amazon Cognito User Pool configuration. Defined below.
        * `appIdClientRegex` (`str`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
        * `awsRegion` (`str`) - The AWS region in which the user pool was created.
        * `user_pool_id` (`str`) - The user pool ID.
    """
    arn: pulumi.Output[str]
    """
    The ARN
    """
    authentication_type: pulumi.Output[str]
    """
    The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
    """
    log_config: pulumi.Output[dict]
    """
    Nested argument containing logging configuration. Defined below.

      * `cloudwatchLogsRoleArn` (`str`) - Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
      * `excludeVerboseContent` (`bool`) - Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: `true`, `false`. Default value: `false`
      * `fieldLogLevel` (`str`) - Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.
    """
    name: pulumi.Output[str]
    """
    A user-supplied name for the GraphqlApi.
    """
    openid_connect_config: pulumi.Output[dict]
    """
    Nested argument containing OpenID Connect configuration. Defined below.

      * `authTtl` (`float`) - Number of milliseconds a token is valid after being authenticated.
      * `client_id` (`str`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
      * `iatTtl` (`float`) - Number of milliseconds a token is valid after being issued to a user.
      * `issuer` (`str`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.
    """
    schema: pulumi.Output[str]
    """
    The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the resource.
    """
    uris: pulumi.Output[dict]
    """
    Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
    """
    user_pool_config: pulumi.Output[dict]
    """
    The Amazon Cognito User Pool configuration. Defined below.

      * `appIdClientRegex` (`str`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
      * `awsRegion` (`str`) - The AWS region in which the user pool was created.
      * `default_action` (`str`) - The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
      * `user_pool_id` (`str`) - The user pool ID.
    """
    xray_enabled: pulumi.Output[bool]
    """
    Whether tracing with X-ray is enabled. Defaults to false.
    """
    def __init__(__self__, resource_name, opts=None, additional_authentication_providers=None, authentication_type=None, log_config=None, name=None, openid_connect_config=None, schema=None, tags=None, user_pool_config=None, xray_enabled=None, __props__=None, __name__=None, __opts__=None):
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
            user_pool_config={
                "awsRegion": data[".getRegion"]["current"]["name"],
                "default_action": "DENY",
                "user_pool_id": aws_cognito_user_pool["example"]["id"],
            })
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
            openid_connect_config={
                "issuer": "https://example.com",
            })
        ```

        ### With Multiple Authentication Providers

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.appsync.GraphQLApi("example",
            additional_authentication_providers=[{
                "authentication_type": "AWS_IAM",
            }],
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
        example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", log_config={
            "cloudwatchLogsRoleArn": example_role.arn,
            "fieldLogLevel": "ERROR",
        })
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] authentication_type: The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        :param pulumi.Input[dict] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: A user-supplied name for the GraphqlApi.
        :param pulumi.Input[dict] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[dict] user_pool_config: The Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.

        The **additional_authentication_providers** object supports the following:

          * `authentication_type` (`pulumi.Input[str]`) - The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
          * `openid_connect_config` (`pulumi.Input[dict]`) - Nested argument containing OpenID Connect configuration. Defined below.
            * `authTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being authenticated.
            * `client_id` (`pulumi.Input[str]`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
            * `iatTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being issued to a user.
            * `issuer` (`pulumi.Input[str]`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.

          * `user_pool_config` (`pulumi.Input[dict]`) - The Amazon Cognito User Pool configuration. Defined below.
            * `appIdClientRegex` (`pulumi.Input[str]`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
            * `awsRegion` (`pulumi.Input[str]`) - The AWS region in which the user pool was created.
            * `user_pool_id` (`pulumi.Input[str]`) - The user pool ID.

        The **log_config** object supports the following:

          * `cloudwatchLogsRoleArn` (`pulumi.Input[str]`) - Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
          * `excludeVerboseContent` (`pulumi.Input[bool]`) - Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: `true`, `false`. Default value: `false`
          * `fieldLogLevel` (`pulumi.Input[str]`) - Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.

        The **openid_connect_config** object supports the following:

          * `authTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being authenticated.
          * `client_id` (`pulumi.Input[str]`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
          * `iatTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being issued to a user.
          * `issuer` (`pulumi.Input[str]`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.

        The **user_pool_config** object supports the following:

          * `appIdClientRegex` (`pulumi.Input[str]`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
          * `awsRegion` (`pulumi.Input[str]`) - The AWS region in which the user pool was created.
          * `default_action` (`pulumi.Input[str]`) - The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
          * `user_pool_id` (`pulumi.Input[str]`) - The user pool ID.
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
    def get(resource_name, id, opts=None, additional_authentication_providers=None, arn=None, authentication_type=None, log_config=None, name=None, openid_connect_config=None, schema=None, tags=None, uris=None, user_pool_config=None, xray_enabled=None):
        """
        Get an existing GraphQLApi resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] additional_authentication_providers: One or more additional authentication providers for the GraphqlApi. Defined below.
        :param pulumi.Input[str] arn: The ARN
        :param pulumi.Input[str] authentication_type: The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
        :param pulumi.Input[dict] log_config: Nested argument containing logging configuration. Defined below.
        :param pulumi.Input[str] name: A user-supplied name for the GraphqlApi.
        :param pulumi.Input[dict] openid_connect_config: Nested argument containing OpenID Connect configuration. Defined below.
        :param pulumi.Input[str] schema: The schema definition, in GraphQL schema language format. This provider cannot perform drift detection of this configuration.
        :param pulumi.Input[dict] tags: A map of tags to assign to the resource.
        :param pulumi.Input[dict] uris: Map of URIs associated with the API. e.g. `uris["GRAPHQL"] = https://ID.appsync-api.REGION.amazonaws.com/graphql`
        :param pulumi.Input[dict] user_pool_config: The Amazon Cognito User Pool configuration. Defined below.
        :param pulumi.Input[bool] xray_enabled: Whether tracing with X-ray is enabled. Defaults to false.

        The **additional_authentication_providers** object supports the following:

          * `authentication_type` (`pulumi.Input[str]`) - The authentication type. Valid values: `API_KEY`, `AWS_IAM`, `AMAZON_COGNITO_USER_POOLS`, `OPENID_CONNECT`
          * `openid_connect_config` (`pulumi.Input[dict]`) - Nested argument containing OpenID Connect configuration. Defined below.
            * `authTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being authenticated.
            * `client_id` (`pulumi.Input[str]`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
            * `iatTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being issued to a user.
            * `issuer` (`pulumi.Input[str]`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.

          * `user_pool_config` (`pulumi.Input[dict]`) - The Amazon Cognito User Pool configuration. Defined below.
            * `appIdClientRegex` (`pulumi.Input[str]`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
            * `awsRegion` (`pulumi.Input[str]`) - The AWS region in which the user pool was created.
            * `user_pool_id` (`pulumi.Input[str]`) - The user pool ID.

        The **log_config** object supports the following:

          * `cloudwatchLogsRoleArn` (`pulumi.Input[str]`) - Amazon Resource Name of the service role that AWS AppSync will assume to publish to Amazon CloudWatch logs in your account.
          * `excludeVerboseContent` (`pulumi.Input[bool]`) - Set to TRUE to exclude sections that contain information such as headers, context, and evaluated mapping templates, regardless of logging  level. Valid values: `true`, `false`. Default value: `false`
          * `fieldLogLevel` (`pulumi.Input[str]`) - Field logging level. Valid values: `ALL`, `ERROR`, `NONE`.

        The **openid_connect_config** object supports the following:

          * `authTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being authenticated.
          * `client_id` (`pulumi.Input[str]`) - Client identifier of the Relying party at the OpenID identity provider. This identifier is typically obtained when the Relying party is registered with the OpenID identity provider. You can specify a regular expression so the AWS AppSync can validate against multiple client identifiers at a time.
          * `iatTtl` (`pulumi.Input[float]`) - Number of milliseconds a token is valid after being issued to a user.
          * `issuer` (`pulumi.Input[str]`) - Issuer for the OpenID Connect configuration. The issuer returned by discovery MUST exactly match the value of iss in the ID Token.

        The **user_pool_config** object supports the following:

          * `appIdClientRegex` (`pulumi.Input[str]`) - A regular expression for validating the incoming Amazon Cognito User Pool app client ID.
          * `awsRegion` (`pulumi.Input[str]`) - The AWS region in which the user pool was created.
          * `default_action` (`pulumi.Input[str]`) - The action that you want your GraphQL API to take when a request that uses Amazon Cognito User Pool authentication doesn't match the Amazon Cognito User Pool configuration. Valid: `ALLOW` and `DENY`
          * `user_pool_id` (`pulumi.Input[str]`) - The user pool ID.
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

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
