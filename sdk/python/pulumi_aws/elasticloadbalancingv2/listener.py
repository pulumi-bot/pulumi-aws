# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Listener(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the listener (matches `id`)
    """
    certificate_arn: pulumi.Output[str]
    """
    The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
    """
    default_actions: pulumi.Output[list]
    """
    An Action block. Action blocks are documented below.

      * `authenticateCognito` (`dict`)
        * `authenticationRequestExtraParams` (`dict`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
        * `onUnauthenticatedRequest` (`str`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
        * `scope` (`str`) - The set of user claims to be requested from the IdP.
        * `sessionCookieName` (`str`) - The name of the cookie used to maintain session information.
        * `sessionTimeout` (`float`) - The maximum duration of the authentication session, in seconds.
        * `userPoolArn` (`str`) - The ARN of the Cognito user pool.
        * `userPoolClientId` (`str`) - The ID of the Cognito user pool client.
        * `userPoolDomain` (`str`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

      * `authenticateOidc` (`dict`)
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
        * `status_code` (`str`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

      * `target_group_arn` (`str`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
      * `type` (`str`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
    """
    load_balancer_arn: pulumi.Output[str]
    """
    The ARN of the load balancer.
    """
    port: pulumi.Output[float]
    """
    The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
    """
    protocol: pulumi.Output[str]
    """
    The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
    """
    ssl_policy: pulumi.Output[str]
    """
    The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, default_actions=None, load_balancer_arn=None, port=None, protocol=None, ssl_policy=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer Listener resource.

        > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.

        {{% examples %}}
        {{% /examples %}}

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lb_listener.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
        :param pulumi.Input[list] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[float] port: The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        :param pulumi.Input[str] protocol: The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.

        The **default_actions** object supports the following:

          * `authenticateCognito` (`pulumi.Input[dict]`)
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `userPoolArn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `userPoolClientId` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `userPoolDomain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

          * `authenticateOidc` (`pulumi.Input[dict]`)
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
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
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

            __props__['certificate_arn'] = certificate_arn
            if default_actions is None:
                raise TypeError("Missing required property 'default_actions'")
            __props__['default_actions'] = default_actions
            if load_balancer_arn is None:
                raise TypeError("Missing required property 'load_balancer_arn'")
            __props__['load_balancer_arn'] = load_balancer_arn
            if port is None:
                raise TypeError("Missing required property 'port'")
            __props__['port'] = port
            __props__['protocol'] = protocol
            __props__['ssl_policy'] = ssl_policy
            __props__['arn'] = None
        super(Listener, __self__).__init__(
            'aws:elasticloadbalancingv2/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, certificate_arn=None, default_actions=None, load_balancer_arn=None, port=None, protocol=None, ssl_policy=None):
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the listener (matches `id`)
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the [`lb.ListenerCertificate` resource](https://www.terraform.io/docs/providers/aws/r/lb_listener_certificate.html).
        :param pulumi.Input[list] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[float] port: The port. Specify a value from `1` to `65535` or `#{port}`. Defaults to `#{port}`.
        :param pulumi.Input[str] protocol: The protocol. Valid values are `HTTP`, `HTTPS`, or `#{protocol}`. Defaults to `#{protocol}`.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.

        The **default_actions** object supports the following:

          * `authenticateCognito` (`pulumi.Input[dict]`)
            * `authenticationRequestExtraParams` (`pulumi.Input[dict]`) - The query parameters to include in the redirect request to the authorization endpoint. Max: 10.
            * `onUnauthenticatedRequest` (`pulumi.Input[str]`) - The behavior if the user is not authenticated. Valid values: `deny`, `allow` and `authenticate`
            * `scope` (`pulumi.Input[str]`) - The set of user claims to be requested from the IdP.
            * `sessionCookieName` (`pulumi.Input[str]`) - The name of the cookie used to maintain session information.
            * `sessionTimeout` (`pulumi.Input[float]`) - The maximum duration of the authentication session, in seconds.
            * `userPoolArn` (`pulumi.Input[str]`) - The ARN of the Cognito user pool.
            * `userPoolClientId` (`pulumi.Input[str]`) - The ID of the Cognito user pool client.
            * `userPoolDomain` (`pulumi.Input[str]`) - The domain prefix or fully-qualified domain name of the Cognito user pool.

          * `authenticateOidc` (`pulumi.Input[dict]`)
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
            * `status_code` (`pulumi.Input[str]`) - The HTTP response code. Valid values are `2XX`, `4XX`, or `5XX`.

          * `target_group_arn` (`pulumi.Input[str]`) - The ARN of the Target Group to which to route traffic. Required if `type` is `forward`.
          * `type` (`pulumi.Input[str]`) - The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["certificate_arn"] = certificate_arn
        __props__["default_actions"] = default_actions
        __props__["load_balancer_arn"] = load_balancer_arn
        __props__["port"] = port
        __props__["protocol"] = protocol
        __props__["ssl_policy"] = ssl_policy
        return Listener(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

