# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Listener']


class Listener(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name (ARN) of the target group.
    """

    certificate_arn: pulumi.Output[Optional[str]] = pulumi.property("certificateArn")
    """
    The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
    """

    default_actions: pulumi.Output[List['outputs.ListenerDefaultAction']] = pulumi.property("defaultActions")
    """
    An Action block. Action blocks are documented below.
    """

    load_balancer_arn: pulumi.Output[str] = pulumi.property("loadBalancerArn")
    """
    The ARN of the load balancer.
    """

    port: pulumi.Output[float] = pulumi.property("port")
    """
    The port on which the load balancer is listening.
    """

    protocol: pulumi.Output[Optional[str]] = pulumi.property("protocol")
    """
    The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
    """

    ssl_policy: pulumi.Output[str] = pulumi.property("sslPolicy")
    """
    The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 default_actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]]] = None,
                 load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[float]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 ssl_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Load Balancer Listener resource.

        > **Note:** `alb.Listener` is known as `lb.Listener`. The functionality is identical.

        ## Example Usage
        ### Forward Action

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener",
            load_balancer_arn=front_end_load_balancer.arn,
            port=443,
            protocol="HTTPS",
            ssl_policy="ELBSecurityPolicy-2016-08",
            certificate_arn="arn:aws:iam::187416307283:server-certificate/test_cert_rab3wuqwgja25ct3n4jdj2tzu4",
            default_actions=[{
                "type": "forward",
                "target_group_arn": front_end_target_group.arn,
            }])
        ```
        ### Redirect Action

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener",
            load_balancer_arn=front_end_load_balancer.arn,
            port=80,
            protocol="HTTP",
            default_actions=[{
                "type": "redirect",
                "redirect": {
                    "port": "443",
                    "protocol": "HTTPS",
                    "status_code": "HTTP_301",
                },
            }])
        ```
        ### Fixed-response Action

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener",
            load_balancer_arn=front_end_load_balancer.arn,
            port=80,
            protocol="HTTP",
            default_actions=[{
                "type": "fixed-response",
                "fixedResponse": {
                    "content_type": "text/plain",
                    "messageBody": "Fixed response content",
                    "status_code": "200",
                },
            }])
        ```
        ### Authenticate-cognito Action

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
        # ...
        pool = aws.cognito.UserPool("pool")
        # ...
        client = aws.cognito.UserPoolClient("client")
        # ...
        domain = aws.cognito.UserPoolDomain("domain")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener",
            load_balancer_arn=front_end_load_balancer.arn,
            port=80,
            protocol="HTTP",
            default_actions=[
                {
                    "type": "authenticate-cognito",
                    "authenticateCognito": {
                        "userPoolArn": pool.arn,
                        "userPoolClientId": client.id,
                        "userPoolDomain": domain.domain,
                    },
                },
                {
                    "type": "forward",
                    "target_group_arn": front_end_target_group.arn,
                },
            ])
        ```
        ### Authenticate-oidc Action

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_target_group = aws.lb.TargetGroup("frontEndTargetGroup")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener",
            load_balancer_arn=front_end_load_balancer.arn,
            port=80,
            protocol="HTTP",
            default_actions=[
                {
                    "type": "authenticate-oidc",
                    "authenticateOidc": {
                        "authorizationEndpoint": "https://example.com/authorization_endpoint",
                        "client_id": "client_id",
                        "client_secret": "client_secret",
                        "issuer": "https://example.com",
                        "tokenEndpoint": "https://example.com/token_endpoint",
                        "userInfoEndpoint": "https://example.com/user_info_endpoint",
                    },
                },
                {
                    "type": "forward",
                    "target_group_arn": front_end_target_group.arn,
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[float] port: The port on which the load balancer is listening.
        :param pulumi.Input[str] protocol: The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:applicationloadbalancing/listener:Listener")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Listener, __self__).__init__(
            'aws:alb/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            default_actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]]] = None,
            load_balancer_arn: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[float]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            ssl_policy: Optional[pulumi.Input[str]] = None) -> 'Listener':
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the target group.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[float] port: The port on which the load balancer is listening.
        :param pulumi.Input[str] protocol: The protocol for connections from clients to the load balancer. Valid values are `TCP`, `TLS`, `UDP`, `TCP_UDP`, `HTTP` and `HTTPS`. Defaults to `HTTP`.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
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
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

