# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ListenerArgs', 'Listener']

@pulumi.input_type
class ListenerArgs:
    def __init__(__self__, *,
                 default_actions: pulumi.Input[Sequence[pulumi.Input['ListenerDefaultActionArgs']]],
                 load_balancer_arn: pulumi.Input[str],
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 ssl_policy: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Listener resource.
        :param pulumi.Input[Sequence[pulumi.Input['ListenerDefaultActionArgs']]] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        :param pulumi.Input[int] port: The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
        :param pulumi.Input[str] protocol: The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        """
        pulumi.set(__self__, "default_actions", default_actions)
        pulumi.set(__self__, "load_balancer_arn", load_balancer_arn)
        if certificate_arn is not None:
            pulumi.set(__self__, "certificate_arn", certificate_arn)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if ssl_policy is not None:
            pulumi.set(__self__, "ssl_policy", ssl_policy)

    @property
    @pulumi.getter(name="defaultActions")
    def default_actions(self) -> pulumi.Input[Sequence[pulumi.Input['ListenerDefaultActionArgs']]]:
        """
        An Action block. Action blocks are documented below.
        """
        return pulumi.get(self, "default_actions")

    @default_actions.setter
    def default_actions(self, value: pulumi.Input[Sequence[pulumi.Input['ListenerDefaultActionArgs']]]):
        pulumi.set(self, "default_actions", value)

    @property
    @pulumi.getter(name="loadBalancerArn")
    def load_balancer_arn(self) -> pulumi.Input[str]:
        """
        The ARN of the load balancer.
        """
        return pulumi.get(self, "load_balancer_arn")

    @load_balancer_arn.setter
    def load_balancer_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "load_balancer_arn", value)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        """
        return pulumi.get(self, "certificate_arn")

    @certificate_arn.setter
    def certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_arn", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        """
        return pulumi.get(self, "ssl_policy")

    @ssl_policy.setter
    def ssl_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ssl_policy", value)


warnings.warn("""aws.applicationloadbalancing.Listener has been deprecated in favor of aws.alb.Listener""", DeprecationWarning)


class Listener(pulumi.CustomResource):
    warnings.warn("""aws.applicationloadbalancing.Listener has been deprecated in favor of aws.alb.Listener""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 default_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]]] = None,
                 load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="forward",
                target_group_arn=front_end_target_group.arn,
            )])
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="redirect",
                redirect=aws.lb.ListenerDefaultActionRedirectArgs(
                    port="443",
                    protocol="HTTPS",
                    status_code="HTTP_301",
                ),
            )])
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="fixed-response",
                fixed_response=aws.lb.ListenerDefaultActionFixedResponseArgs(
                    content_type="text/plain",
                    message_body="Fixed response content",
                    status_code="200",
                ),
            )])
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
                aws.lb.ListenerDefaultActionArgs(
                    type="authenticate-cognito",
                    authenticate_cognito=aws.lb.ListenerDefaultActionAuthenticateCognitoArgs(
                        user_pool_arn=pool.arn,
                        user_pool_client_id=client.id,
                        user_pool_domain=domain.domain,
                    ),
                ),
                aws.lb.ListenerDefaultActionArgs(
                    type="forward",
                    target_group_arn=front_end_target_group.arn,
                ),
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
                aws.lb.ListenerDefaultActionArgs(
                    type="authenticate-oidc",
                    authenticate_oidc=aws.lb.ListenerDefaultActionAuthenticateOidcArgs(
                        authorization_endpoint="https://example.com/authorization_endpoint",
                        client_id="client_id",
                        client_secret="client_secret",
                        issuer="https://example.com",
                        token_endpoint="https://example.com/token_endpoint",
                        user_info_endpoint="https://example.com/user_info_endpoint",
                    ),
                ),
                aws.lb.ListenerDefaultActionArgs(
                    type="forward",
                    target_group_arn=front_end_target_group.arn,
                ),
            ])
        ```
        ### Gateway Load Balancer Listener

        ```python
        import pulumi
        import pulumi_aws as aws

        example_load_balancer = aws.lb.LoadBalancer("exampleLoadBalancer",
            load_balancer_type="gateway",
            subnet_mappings=[aws.lb.LoadBalancerSubnetMappingArgs(
                subnet_id=aws_subnet["example"]["id"],
            )])
        example_target_group = aws.lb.TargetGroup("exampleTargetGroup",
            port=6081,
            protocol="GENEVE",
            vpc_id=aws_vpc["example"]["id"],
            health_check=aws.lb.TargetGroupHealthCheckArgs(
                port="80",
                protocol="HTTP",
            ))
        example_listener = aws.lb.Listener("exampleListener",
            load_balancer_arn=example_load_balancer.id,
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                target_group_arn=example_target_group.id,
                type="forward",
            )])
        ```

        ## Import

        Listeners can be imported using their ARN, e.g.

        ```sh
         $ pulumi import aws:applicationloadbalancing/listener:Listener front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[int] port: The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
        :param pulumi.Input[str] protocol: The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
        :param pulumi.Input[str] ssl_policy: The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ListenerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="forward",
                target_group_arn=front_end_target_group.arn,
            )])
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="redirect",
                redirect=aws.lb.ListenerDefaultActionRedirectArgs(
                    port="443",
                    protocol="HTTPS",
                    status_code="HTTP_301",
                ),
            )])
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
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                type="fixed-response",
                fixed_response=aws.lb.ListenerDefaultActionFixedResponseArgs(
                    content_type="text/plain",
                    message_body="Fixed response content",
                    status_code="200",
                ),
            )])
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
                aws.lb.ListenerDefaultActionArgs(
                    type="authenticate-cognito",
                    authenticate_cognito=aws.lb.ListenerDefaultActionAuthenticateCognitoArgs(
                        user_pool_arn=pool.arn,
                        user_pool_client_id=client.id,
                        user_pool_domain=domain.domain,
                    ),
                ),
                aws.lb.ListenerDefaultActionArgs(
                    type="forward",
                    target_group_arn=front_end_target_group.arn,
                ),
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
                aws.lb.ListenerDefaultActionArgs(
                    type="authenticate-oidc",
                    authenticate_oidc=aws.lb.ListenerDefaultActionAuthenticateOidcArgs(
                        authorization_endpoint="https://example.com/authorization_endpoint",
                        client_id="client_id",
                        client_secret="client_secret",
                        issuer="https://example.com",
                        token_endpoint="https://example.com/token_endpoint",
                        user_info_endpoint="https://example.com/user_info_endpoint",
                    ),
                ),
                aws.lb.ListenerDefaultActionArgs(
                    type="forward",
                    target_group_arn=front_end_target_group.arn,
                ),
            ])
        ```
        ### Gateway Load Balancer Listener

        ```python
        import pulumi
        import pulumi_aws as aws

        example_load_balancer = aws.lb.LoadBalancer("exampleLoadBalancer",
            load_balancer_type="gateway",
            subnet_mappings=[aws.lb.LoadBalancerSubnetMappingArgs(
                subnet_id=aws_subnet["example"]["id"],
            )])
        example_target_group = aws.lb.TargetGroup("exampleTargetGroup",
            port=6081,
            protocol="GENEVE",
            vpc_id=aws_vpc["example"]["id"],
            health_check=aws.lb.TargetGroupHealthCheckArgs(
                port="80",
                protocol="HTTP",
            ))
        example_listener = aws.lb.Listener("exampleListener",
            load_balancer_arn=example_load_balancer.id,
            default_actions=[aws.lb.ListenerDefaultActionArgs(
                target_group_arn=example_target_group.id,
                type="forward",
            )])
        ```

        ## Import

        Listeners can be imported using their ARN, e.g.

        ```sh
         $ pulumi import aws:applicationloadbalancing/listener:Listener front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener/app/front-end-alb/8e4497da625e2d8a/9ab28ade35828f96
        ```

        :param str resource_name: The name of the resource.
        :param ListenerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ListenerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 default_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]]] = None,
                 load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 ssl_policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        pulumi.log.warn("""Listener is deprecated: aws.applicationloadbalancing.Listener has been deprecated in favor of aws.alb.Listener""")
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
            if default_actions is None and not opts.urn:
                raise TypeError("Missing required property 'default_actions'")
            __props__['default_actions'] = default_actions
            if load_balancer_arn is None and not opts.urn:
                raise TypeError("Missing required property 'load_balancer_arn'")
            __props__['load_balancer_arn'] = load_balancer_arn
            __props__['port'] = port
            __props__['protocol'] = protocol
            __props__['ssl_policy'] = ssl_policy
            __props__['arn'] = None
        super(Listener, __self__).__init__(
            'aws:applicationloadbalancing/listener:Listener',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            default_actions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]]] = None,
            load_balancer_arn: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            ssl_policy: Optional[pulumi.Input[str]] = None) -> 'Listener':
        """
        Get an existing Listener resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the target group.
        :param pulumi.Input[str] certificate_arn: The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ListenerDefaultActionArgs']]]] default_actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] load_balancer_arn: The ARN of the load balancer.
        :param pulumi.Input[int] port: The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
        :param pulumi.Input[str] protocol: The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
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

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The Amazon Resource Name (ARN) of the target group.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The ARN of the default SSL server certificate. Exactly one certificate is required if the protocol is HTTPS. For adding additional SSL certificates, see the `lb.ListenerCertificate` resource.
        """
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="defaultActions")
    def default_actions(self) -> pulumi.Output[Sequence['outputs.ListenerDefaultAction']]:
        """
        An Action block. Action blocks are documented below.
        """
        return pulumi.get(self, "default_actions")

    @property
    @pulumi.getter(name="loadBalancerArn")
    def load_balancer_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the load balancer.
        """
        return pulumi.get(self, "load_balancer_arn")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[Optional[int]]:
        """
        The port on which the load balancer is listening. Not valid for Gateway Load Balancers.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        The protocol for connections from clients to the load balancer. For Application Load Balancers, valid values are `HTTP` and `HTTPS`, with a default of `HTTP`. For Network Load Balancers, valid values are `TCP`, `TLS`, `UDP`, and `TCP_UDP`. Not valid to use `UDP` or `TCP_UDP` if dual-stack mode is enabled. Not valid for Gateway Load Balancers.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sslPolicy")
    def ssl_policy(self) -> pulumi.Output[str]:
        """
        The name of the SSL Policy for the listener. Required if `protocol` is `HTTPS` or `TLS`.
        """
        return pulumi.get(self, "ssl_policy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

