# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule", DeprecationWarning)
class ListenerRule(pulumi.CustomResource):
    actions: pulumi.Output[list]
    """
    An Action block. Action blocks are documented below.

      * `authenticateCognito` (`dict`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
        * `authenticationRequestExtraParams` (`dict`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        * `onUnauthenticatedRequest` (`str`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        * `scope` (`str`) - The set of user claims to be requested from the IdP.
        * `sessionCookieName` (`str`) - The name of the cookie used to maintain session information.
        * `sessionTimeout` (`float`) - The maximum duration of the authentication session, in seconds.
        * `userPoolArn` (`str`) - The ARN of the Cognito user pool.
        * `userPoolClientId` (`str`) - The ID of the Cognito user pool client.
        * `userPoolDomain` (`str`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

      * `authenticateOidc` (`dict`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
        * `authenticationRequestExtraParams` (`dict`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        * `authorizationEndpoint` (`str`) - The authorization endpoint of the IdP.
        * `client_id` (`str`) - The OAuth 2.0 client identifier.
        * `client_secret` (`str`) - The OAuth 2.0 client secret.
        * `issuer` (`str`) - The OIDC issuer identifier of the IdP.
        * `onUnauthenticatedRequest` (`str`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        * `scope` (`str`) - The set of user claims to be requested from the IdP.
        * `sessionCookieName` (`str`) - The name of the cookie used to maintain session information.
        * `sessionTimeout` (`float`) - The maximum duration of the authentication session, in seconds.
        * `tokenEndpoint` (`str`) - The token endpoint of the IdP.
        * `userInfoEndpoint` (`str`) - The user info endpoint of the IdP.

      * `fixedResponse` (`dict`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
        * `content_type` (`str`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
        * `messageBody` (`str`) - The message body.
        * `status_code` (`str`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

      * `order` (`float`)
      * `redirect` (`dict`) - Information for creating a redirect action. Required if `type` is `redirect`.
        * `host` (`str`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
        * `path` (`str`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
        * `port` (`str`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        * `protocol` (`str`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        * `query` (`str`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
        * `status_code` (`str`) - The HTTP redirect code. The redirect is either permanent (`HTTP_301`) or temporary (`HTTP_302`).

      * `target_group_arn` (`str`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
      * `type` (`str`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the rule (matches `id`)
    """
    conditions: pulumi.Output[list]
    """
    A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.

      * `field` (`str`) - The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
      * `hostHeader` (`dict`) - Contains a single `values` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
        * `values` (`list`) - List of exactly one pattern to match. Required when `field` is set.

      * `httpHeader` (`dict`) - HTTP headers to match. HTTP Header block fields documented below.
        * `httpHeaderName` (`str`) - Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
        * `values` (`list`) - List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.

      * `httpRequestMethod` (`dict`) - Contains a single `values` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
        * `values` (`list`) - List of exactly one pattern to match. Required when `field` is set.

      * `pathPattern` (`dict`) - Contains a single `values` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
        * `values` (`list`) - List of exactly one pattern to match. Required when `field` is set.

      * `queryStrings` (`list`) - Query strings to match. Query String block fields documented below.
        * `key` (`str`) - Query string key pattern to match.
        * `value` (`str`) - Query string value pattern to match.

      * `sourceIp` (`dict`) - Contains a single `values` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
        * `values` (`list`) - List of exactly one pattern to match. Required when `field` is set.

      * `values` (`str`) - List of exactly one pattern to match. Required when `field` is set.
    """
    listener_arn: pulumi.Output[str]
    """
    The ARN of the listener to which to attach the rule.
    """
    priority: pulumi.Output[float]
    """
    The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
    """
    warnings.warn("aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule", DeprecationWarning)
    def __init__(__self__, resource_name, opts=None, actions=None, conditions=None, listener_arn=None, priority=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer Listener Rule resource.

        > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        front_end_listener = aws.lb.Listener("frontEndListener")
        static = aws.lb.ListenerRule("static",
            actions=[{
                "targetGroupArn": aws_lb_target_group["static"]["arn"],
                "type": "forward",
            }],
            conditions=[
                {
                    "pathPattern": {
                        "values": ["/static/*"],
                    },
                },
                {
                    "hostHeader": {
                        "values": ["example.com"],
                    },
                },
            ],
            listener_arn=front_end_listener.arn,
            priority=100)
        host_based_routing = aws.lb.ListenerRule("hostBasedRouting",
            actions=[{
                "targetGroupArn": aws_lb_target_group["static"]["arn"],
                "type": "forward",
            }],
            conditions=[{
                "hostHeader": {
                    "values": ["my-service.*.mydomain.io"],
                },
            }],
            listener_arn=front_end_listener.arn,
            priority=99)
        redirect_http_to_https = aws.lb.ListenerRule("redirectHttpToHttps",
            actions=[{
                "redirect": {
                    "port": "443",
                    "protocol": "HTTPS",
                    "statusCode": "HTTP_301",
                },
                "type": "redirect",
            }],
            conditions=[{
                "httpHeader": {
                    "httpHeaderName": "X-Forwarded-For",
                    "values": ["192.168.1.*"],
                },
            }],
            listener_arn=front_end_listener.arn)
        health_check = aws.lb.ListenerRule("healthCheck",
            actions=[{
                "fixedResponse": {
                    "contentType": "text/plain",
                    "messageBody": "HEALTHY",
                    "statusCode": "200",
                },
                "type": "fixed-response",
            }],
            conditions=[{
                "queryString": [
                    {
                        "key": "health",
                        "value": "check",
                    },
                    {
                        "value": "bar",
                    },
                ],
            }],
            listener_arn=front_end_listener.arn)
        pool = aws.cognito.UserPool("pool")
        client = aws.cognito.UserPoolClient("client")
        domain = aws.cognito.UserPoolDomain("domain")
        admin = aws.lb.ListenerRule("admin",
            actions=[
                {
                    "authenticateOidc": {
                        "authorizationEndpoint": "https://example.com/authorization_endpoint",
                        "clientId": "client_id",
                        "clientSecret": "client_secret",
                        "issuer": "https://example.com",
                        "tokenEndpoint": "https://example.com/token_endpoint",
                        "userInfoEndpoint": "https://example.com/user_info_endpoint",
                    },
                    "type": "authenticate-oidc",
                },
                {
                    "targetGroupArn": aws_lb_target_group["static"]["arn"],
                    "type": "forward",
                },
            ],
            listener_arn=front_end_listener.arn)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[list] conditions: A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.

        The **actions** object supports the following:

          * `authenticateCognito` (`pulumi.Input[dict]`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `userPoolArn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `userPoolClientId` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `userPoolDomain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

          * `authenticateOidc` (`pulumi.Input[dict]`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `authorizationEndpoint` (`pulumi.Input[str]`) - The authorization endpoint of the IdP.
            * `client_id` (`pulumi.Input[str]`) - The OAuth 2.0 client identifier.
            * `client_secret` (`pulumi.Input[str]`) - The OAuth 2.0 client secret.
            * `issuer` (`pulumi.Input[str]`) - The OIDC issuer identifier of the IdP.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `tokenEndpoint` (`pulumi.Input[str]`) - The token endpoint of the IdP.
            * `userInfoEndpoint` (`pulumi.Input[str]`) - The user info endpoint of the IdP.

          * `fixedResponse` (`pulumi.Input[dict]`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
            * `content_type` (`pulumi.Input[str]`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
            * `messageBody` (`pulumi.Input[str]`) - The message body.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

          * `order` (`pulumi.Input[float]`)
          * `redirect` (`pulumi.Input[dict]`) - Information for creating a redirect action. Required if `type` is `redirect`.
            * `host` (`pulumi.Input[str]`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
            * `path` (`pulumi.Input[str]`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
            * `port` (`pulumi.Input[str]`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
            * `protocol` (`pulumi.Input[str]`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
            * `query` (`pulumi.Input[str]`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
            * `status_code` (`pulumi.Input[str]`) - The HTTP redirect code. The redirect is either permanent (`HTTP_301`) or temporary (`HTTP_302`).

          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.

        The **conditions** object supports the following:

          * `field` (`pulumi.Input[str]`) - The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
          * `hostHeader` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `httpHeader` (`pulumi.Input[dict]`) - HTTP headers to match. HTTP Header block fields documented below.
            * `httpHeaderName` (`pulumi.Input[str]`) - Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
            * `values` (`pulumi.Input[list]`) - List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.

          * `httpRequestMethod` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `pathPattern` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `queryStrings` (`pulumi.Input[list]`) - Query strings to match. Query String block fields documented below.
            * `key` (`pulumi.Input[str]`) - Query string key pattern to match.
            * `value` (`pulumi.Input[str]`) - Query string value pattern to match.

          * `sourceIp` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `values` (`pulumi.Input[str]`) - List of exactly one pattern to match. Required when `field` is set.
        """
        pulumi.log.warn("ListenerRule is deprecated: aws.applicationloadbalancing.ListenerRule has been deprecated in favor of aws.alb.ListenerRule")
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

            if actions is None:
                raise TypeError("Missing required property 'actions'")
            __props__['actions'] = actions
            if conditions is None:
                raise TypeError("Missing required property 'conditions'")
            __props__['conditions'] = conditions
            if listener_arn is None:
                raise TypeError("Missing required property 'listener_arn'")
            __props__['listener_arn'] = listener_arn
            __props__['priority'] = priority
            __props__['arn'] = None
        super(ListenerRule, __self__).__init__(
            'aws:applicationloadbalancing/listenerRule:ListenerRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, actions=None, arn=None, conditions=None, listener_arn=None, priority=None):
        """
        Get an existing ListenerRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] arn: The ARN of the rule (matches `id`)
        :param pulumi.Input[list] conditions: A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.

        The **actions** object supports the following:

          * `authenticateCognito` (`pulumi.Input[dict]`) - Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `userPoolArn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `userPoolClientId` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `userPoolDomain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

          * `authenticateOidc` (`pulumi.Input[dict]`) - Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `authorizationEndpoint` (`pulumi.Input[str]`) - The authorization endpoint of the IdP.
            * `client_id` (`pulumi.Input[str]`) - The OAuth 2.0 client identifier.
            * `client_secret` (`pulumi.Input[str]`) - The OAuth 2.0 client secret.
            * `issuer` (`pulumi.Input[str]`) - The OIDC issuer identifier of the IdP.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `tokenEndpoint` (`pulumi.Input[str]`) - The token endpoint of the IdP.
            * `userInfoEndpoint` (`pulumi.Input[str]`) - The user info endpoint of the IdP.

          * `fixedResponse` (`pulumi.Input[dict]`) - Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
            * `content_type` (`pulumi.Input[str]`) - The content type. Valid values are `text/plain`, `text/css`, `text/html`, `application/javascript` and `application/json`.
            * `messageBody` (`pulumi.Input[str]`) - The message body.
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

          * `order` (`pulumi.Input[float]`)
          * `redirect` (`pulumi.Input[dict]`) - Information for creating a redirect action. Required if `type` is `redirect`.
            * `host` (`pulumi.Input[str]`) - The hostname. This component is not percent-encoded. The hostname can contain `#{host}`. Defaults to `#{host}`.
            * `path` (`pulumi.Input[str]`) - The absolute path, starting with the leading "/". This component is not percent-encoded. The path can contain #{host}, #{path}, and #{port}. Defaults to `/#{path}`.
            * `port` (`pulumi.Input[str]`) - The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
            * `protocol` (`pulumi.Input[str]`) - The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
            * `query` (`pulumi.Input[str]`) - The query parameters, URL-encoded when necessary, but not percent-encoded. Do not include the leading "?". Defaults to `#{query}`.
            * `status_code` (`pulumi.Input[str]`) - The HTTP redirect code. The redirect is either permanent (`HTTP_301`) or temporary (`HTTP_302`).

          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.

        The **conditions** object supports the following:

          * `field` (`pulumi.Input[str]`) - The type of condition. Valid values are `host-header` or `path-pattern`. Must also set `values`.
          * `hostHeader` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of host header patterns to match. The maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `httpHeader` (`pulumi.Input[dict]`) - HTTP headers to match. HTTP Header block fields documented below.
            * `httpHeaderName` (`pulumi.Input[str]`) - Name of HTTP header to search. The maximum size is 40 characters. Comparison is case insensitive. Only RFC7240 characters are supported. Wildcards are not supported. You cannot use HTTP header condition to specify the host header, use a `host-header` condition instead.
            * `values` (`pulumi.Input[list]`) - List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.

          * `httpRequestMethod` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of HTTP request methods or verbs to match. Maximum size is 40 characters. Only allowed characters are A-Z, hyphen (-) and underscore (\_). Comparison is case sensitive. Wildcards are not supported. Only one needs to match for the condition to be satisfied. AWS recommends that GET and HEAD requests are routed in the same way because the response to a HEAD request may be cached.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `pathPattern` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of path patterns to match against the request URL. Maximum size of each pattern is 128 characters. Comparison is case sensitive. Wildcard characters supported: * (matches 0 or more characters) and ? (matches exactly 1 character). Only one pattern needs to match for the condition to be satisfied. Path pattern is compared only to the path of the URL, not to its query string. To compare against the query string, use a `query-string` condition.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `queryStrings` (`pulumi.Input[list]`) - Query strings to match. Query String block fields documented below.
            * `key` (`pulumi.Input[str]`) - Query string key pattern to match.
            * `value` (`pulumi.Input[str]`) - Query string value pattern to match.

          * `sourceIp` (`pulumi.Input[dict]`) - Contains a single `values` item which is a list of source IP CIDR notations to match. You can use both IPv4 and IPv6 addresses. Wildcards are not supported. Condition is satisfied if the source IP address of the request matches one of the CIDR blocks. Condition is not satisfied by the addresses in the `X-Forwarded-For` header, use `http-header` condition instead.
            * `values` (`pulumi.Input[list]`) - List of exactly one pattern to match. Required when `field` is set.

          * `values` (`pulumi.Input[str]`) - List of exactly one pattern to match. Required when `field` is set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["actions"] = actions
        __props__["arn"] = arn
        __props__["conditions"] = conditions
        __props__["listener_arn"] = listener_arn
        __props__["priority"] = priority
        return ListenerRule(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

