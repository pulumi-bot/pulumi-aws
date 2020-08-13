# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class Stage(pulumi.CustomResource):
    access_log_settings: pulumi.Output[dict]
    """
    Settings for logging access in this stage.
    Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).

      * `destination_arn` (`str`) - The ARN of the CloudWatch Logs log group to receive access logs. Any trailing `:*` is trimmed from the ARN.
      * `format` (`str`) - A single line [format](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#apigateway-cloudwatch-log-formats) of the access logs of data, as specified by [selected $context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-logging.html).
    """
    api_id: pulumi.Output[str]
    """
    The API identifier.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the stage.
    """
    auto_deploy: pulumi.Output[bool]
    """
    Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
    """
    client_certificate_id: pulumi.Output[str]
    """
    The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
    Supported only for WebSocket APIs.
    """
    default_route_settings: pulumi.Output[dict]
    """
    The default route settings for the stage.

      * `dataTraceEnabled` (`bool`) - Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
        Defaults to `false`. Supported only for WebSocket APIs.
      * `detailedMetricsEnabled` (`bool`) - Whether detailed metrics are enabled for the default route. Defaults to `false`.
      * `loggingLevel` (`str`) - The logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
        Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
      * `throttlingBurstLimit` (`float`) - The throttling burst limit for the default route.
      * `throttlingRateLimit` (`float`) - The throttling rate limit for the default route.
    """
    deployment_id: pulumi.Output[str]
    """
    The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
    """
    description: pulumi.Output[str]
    """
    The description for the stage.
    """
    execution_arn: pulumi.Output[str]
    """
    The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
    or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
    See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
    Set only for WebSocket APIs.
    """
    invoke_url: pulumi.Output[str]
    """
    The URL to invoke the API pointing to the stage,
    e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
    """
    name: pulumi.Output[str]
    """
    The name of the stage.
    """
    route_settings: pulumi.Output[list]
    """
    Route settings for the stage.

      * `dataTraceEnabled` (`bool`) - Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
        Defaults to `false`. Supported only for WebSocket APIs.
      * `detailedMetricsEnabled` (`bool`) - Whether detailed metrics are enabled for the route. Defaults to `false`.
      * `loggingLevel` (`str`) - The logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
        Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
      * `route_key` (`str`) - Route key.
      * `throttlingBurstLimit` (`float`) - The throttling burst limit for the route.
      * `throttlingRateLimit` (`float`) - The throttling rate limit for the route.
    """
    stage_variables: pulumi.Output[dict]
    """
    A map that defines the stage variables for the stage.
    """
    tags: pulumi.Output[dict]
    """
    A map of tags to assign to the stage.
    """
    def __init__(__self__, resource_name, opts=None, access_log_settings=None, api_id=None, auto_deploy=None, client_certificate_id=None, default_route_settings=None, deployment_id=None, description=None, name=None, route_settings=None, stage_variables=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Amazon API Gateway Version 2 stage.
        More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

        ## Example Usage
        ### Basic

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.apigatewayv2.Stage("example", api_id=aws_apigatewayv2_api["example"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_log_settings: Settings for logging access in this stage.
               Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[bool] auto_deploy: Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
               Supported only for WebSocket APIs.
        :param pulumi.Input[dict] default_route_settings: The default route settings for the stage.
        :param pulumi.Input[str] deployment_id: The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
        :param pulumi.Input[str] description: The description for the stage.
        :param pulumi.Input[str] name: The name of the stage.
        :param pulumi.Input[list] route_settings: Route settings for the stage.
        :param pulumi.Input[dict] stage_variables: A map that defines the stage variables for the stage.
        :param pulumi.Input[dict] tags: A map of tags to assign to the stage.

        The **access_log_settings** object supports the following:

          * `destination_arn` (`pulumi.Input[str]`) - The ARN of the CloudWatch Logs log group to receive access logs. Any trailing `:*` is trimmed from the ARN.
          * `format` (`pulumi.Input[str]`) - A single line [format](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#apigateway-cloudwatch-log-formats) of the access logs of data, as specified by [selected $context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-logging.html).

        The **default_route_settings** object supports the following:

          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Defaults to `false`. Supported only for WebSocket APIs.
          * `detailedMetricsEnabled` (`pulumi.Input[bool]`) - Whether detailed metrics are enabled for the default route. Defaults to `false`.
          * `loggingLevel` (`pulumi.Input[str]`) - The logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - The throttling burst limit for the default route.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - The throttling rate limit for the default route.

        The **route_settings** object supports the following:

          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Defaults to `false`. Supported only for WebSocket APIs.
          * `detailedMetricsEnabled` (`pulumi.Input[bool]`) - Whether detailed metrics are enabled for the route. Defaults to `false`.
          * `loggingLevel` (`pulumi.Input[str]`) - The logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
          * `route_key` (`pulumi.Input[str]`) - Route key.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - The throttling burst limit for the route.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - The throttling rate limit for the route.
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

            __props__['accessLogSettings'] = access_log_settings
            if api_id is None:
                raise TypeError("Missing required property 'api_id'")
            __props__['apiId'] = api_id
            __props__['autoDeploy'] = auto_deploy
            __props__['clientCertificateId'] = client_certificate_id
            __props__['defaultRouteSettings'] = default_route_settings
            __props__['deploymentId'] = deployment_id
            __props__['description'] = description
            __props__['name'] = name
            __props__['routeSettings'] = route_settings
            __props__['stageVariables'] = stage_variables
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['execution_arn'] = None
            __props__['invoke_url'] = None
        super(Stage, __self__).__init__(
            'aws:apigatewayv2/stage:Stage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, access_log_settings=None, api_id=None, arn=None, auto_deploy=None, client_certificate_id=None, default_route_settings=None, deployment_id=None, description=None, execution_arn=None, invoke_url=None, name=None, route_settings=None, stage_variables=None, tags=None):
        """
        Get an existing Stage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_log_settings: Settings for logging access in this stage.
               Use the `apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
        :param pulumi.Input[str] api_id: The API identifier.
        :param pulumi.Input[str] arn: The ARN of the stage.
        :param pulumi.Input[bool] auto_deploy: Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
        :param pulumi.Input[str] client_certificate_id: The identifier of a client certificate for the stage. Use the `apigateway.ClientCertificate` resource to configure a client certificate.
               Supported only for WebSocket APIs.
        :param pulumi.Input[dict] default_route_settings: The default route settings for the stage.
        :param pulumi.Input[str] deployment_id: The deployment identifier of the stage. Use the [`apigatewayv2.Deployment`](https://www.terraform.io/docs/providers/aws/r/apigatewayv2_deployment.html) resource to configure a deployment.
        :param pulumi.Input[str] description: The description for the stage.
        :param pulumi.Input[str] execution_arn: The ARN prefix to be used in an `lambda.Permission`'s `source_arn` attribute
               or in an `iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
               See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
               Set only for WebSocket APIs.
        :param pulumi.Input[str] invoke_url: The URL to invoke the API pointing to the stage,
               e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
        :param pulumi.Input[str] name: The name of the stage.
        :param pulumi.Input[list] route_settings: Route settings for the stage.
        :param pulumi.Input[dict] stage_variables: A map that defines the stage variables for the stage.
        :param pulumi.Input[dict] tags: A map of tags to assign to the stage.

        The **access_log_settings** object supports the following:

          * `destination_arn` (`pulumi.Input[str]`) - The ARN of the CloudWatch Logs log group to receive access logs. Any trailing `:*` is trimmed from the ARN.
          * `format` (`pulumi.Input[str]`) - A single line [format](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#apigateway-cloudwatch-log-formats) of the access logs of data, as specified by [selected $context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-logging.html).

        The **default_route_settings** object supports the following:

          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Whether data trace logging is enabled for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Defaults to `false`. Supported only for WebSocket APIs.
          * `detailedMetricsEnabled` (`pulumi.Input[bool]`) - Whether detailed metrics are enabled for the default route. Defaults to `false`.
          * `loggingLevel` (`pulumi.Input[str]`) - The logging level for the default route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - The throttling burst limit for the default route.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - The throttling rate limit for the default route.

        The **route_settings** object supports the following:

          * `dataTraceEnabled` (`pulumi.Input[bool]`) - Whether data trace logging is enabled for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Defaults to `false`. Supported only for WebSocket APIs.
          * `detailedMetricsEnabled` (`pulumi.Input[bool]`) - Whether detailed metrics are enabled for the route. Defaults to `false`.
          * `loggingLevel` (`pulumi.Input[str]`) - The logging level for the route. Affects the log entries pushed to Amazon CloudWatch Logs.
            Valid values: `ERROR`, `INFO`, `OFF`. Defaults to `OFF`. Supported only for WebSocket APIs. This provider will only perform drift detection of its value when present in a configuration.
          * `route_key` (`pulumi.Input[str]`) - Route key.
          * `throttlingBurstLimit` (`pulumi.Input[float]`) - The throttling burst limit for the route.
          * `throttlingRateLimit` (`pulumi.Input[float]`) - The throttling rate limit for the route.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["access_log_settings"] = access_log_settings
        __props__["api_id"] = api_id
        __props__["arn"] = arn
        __props__["auto_deploy"] = auto_deploy
        __props__["client_certificate_id"] = client_certificate_id
        __props__["default_route_settings"] = default_route_settings
        __props__["deployment_id"] = deployment_id
        __props__["description"] = description
        __props__["execution_arn"] = execution_arn
        __props__["invoke_url"] = invoke_url
        __props__["name"] = name
        __props__["route_settings"] = route_settings
        __props__["stage_variables"] = stage_variables
        __props__["tags"] = tags
        return Stage(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
