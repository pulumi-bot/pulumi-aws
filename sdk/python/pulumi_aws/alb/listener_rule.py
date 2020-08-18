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

__all__ = ['ListenerRule']


class ListenerRule(pulumi.CustomResource):
    actions: pulumi.Output[List['outputs.ListenerRuleAction']] = pulumi.property("actions")
    """
    An Action block. Action blocks are documented below.
    """

    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The Amazon Resource Name (ARN) of the target group.
    """

    conditions: pulumi.Output[List['outputs.ListenerRuleCondition']] = pulumi.property("conditions")
    """
    A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
    """

    listener_arn: pulumi.Output[str] = pulumi.property("listenerArn")
    """
    The ARN of the listener to which to attach the rule.
    """

    priority: pulumi.Output[float] = pulumi.property("priority")
    """
    The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleActionArgs']]]]] = None,
                 conditions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleConditionArgs']]]]] = None,
                 listener_arn: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Load Balancer Listener Rule resource.

        > **Note:** `alb.ListenerRule` is known as `lb.ListenerRule`. The functionality is identical.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
        # ...
        front_end_listener = aws.lb.Listener("frontEndListener")
        # Other parameters
        static = aws.lb.ListenerRule("static",
            listener_arn=front_end_listener.arn,
            priority=100,
            actions=[{
                "type": "forward",
                "target_group_arn": aws_lb_target_group["static"]["arn"],
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
            ])
        # Forward action
        host_based_weighted_routing = aws.lb.ListenerRule("hostBasedWeightedRouting",
            listener_arn=front_end_listener.arn,
            priority=99,
            actions=[{
                "type": "forward",
                "target_group_arn": aws_lb_target_group["static"]["arn"],
            }],
            conditions=[{
                "hostHeader": {
                    "values": ["my-service.*.mycompany.io"],
                },
            }])
        # Weighted Forward action
        host_based_routing = aws.lb.ListenerRule("hostBasedRouting",
            listener_arn=front_end_listener.arn,
            priority=99,
            actions=[{
                "type": "forward",
                "forward": {
                    "targetGroups": [
                        {
                            "arn": aws_lb_target_group["main"]["arn"],
                            "weight": 80,
                        },
                        {
                            "arn": aws_lb_target_group["canary"]["arn"],
                            "weight": 20,
                        },
                    ],
                    "stickiness": {
                        "enabled": True,
                        "duration": 600,
                    },
                },
            }],
            conditions=[{
                "hostHeader": {
                    "values": ["my-service.*.mycompany.io"],
                },
            }])
        # Redirect action
        redirect_http_to_https = aws.lb.ListenerRule("redirectHttpToHttps",
            listener_arn=front_end_listener.arn,
            actions=[{
                "type": "redirect",
                "redirect": {
                    "port": "443",
                    "protocol": "HTTPS",
                    "status_code": "HTTP_301",
                },
            }],
            conditions=[{
                "httpHeader": {
                    "httpHeaderName": "X-Forwarded-For",
                    "values": ["192.168.1.*"],
                },
            }])
        # Fixed-response action
        health_check = aws.lb.ListenerRule("healthCheck",
            listener_arn=front_end_listener.arn,
            actions=[{
                "type": "fixed-response",
                "fixedResponse": {
                    "content_type": "text/plain",
                    "messageBody": "HEALTHY",
                    "status_code": "200",
                },
            }],
            conditions=[{
                "queryStrings": [
                    {
                        "key": "health",
                        "value": "check",
                    },
                    {
                        "value": "bar",
                    },
                ],
            }])
        # Authenticate-cognito Action
        pool = aws.cognito.UserPool("pool")
        # ...
        client = aws.cognito.UserPoolClient("client")
        # ...
        domain = aws.cognito.UserPoolDomain("domain")
        # ...
        admin_listener_rule = aws.lb.ListenerRule("adminListenerRule",
            listener_arn=front_end_listener.arn,
            actions=[
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
                    "target_group_arn": aws_lb_target_group["static"]["arn"],
                },
            ])
        # Authenticate-oidc Action
        admin_lb_listener_rule_listener_rule = aws.lb.ListenerRule("adminLb/listenerRuleListenerRule",
            listener_arn=front_end_listener.arn,
            actions=[
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
                    "target_group_arn": aws_lb_target_group["static"]["arn"],
                },
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleActionArgs']]]] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleConditionArgs']]]] conditions: A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
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
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:applicationloadbalancing/listenerRule:ListenerRule")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ListenerRule, __self__).__init__(
            'aws:alb/listenerRule:ListenerRule',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            actions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleActionArgs']]]]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            conditions: Optional[pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleConditionArgs']]]]] = None,
            listener_arn: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None) -> 'ListenerRule':
        """
        Get an existing ListenerRule resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleActionArgs']]]] actions: An Action block. Action blocks are documented below.
        :param pulumi.Input[str] arn: The Amazon Resource Name (ARN) of the target group.
        :param pulumi.Input[List[pulumi.Input[pulumi.InputType['ListenerRuleConditionArgs']]]] conditions: A Condition block. Multiple condition blocks of different types can be set and all must be satisfied for the rule to match. Condition blocks are documented below.
        :param pulumi.Input[str] listener_arn: The ARN of the listener to which to attach the rule.
        :param pulumi.Input[float] priority: The priority for the rule between `1` and `50000`. Leaving it unset will automatically set the rule with next available priority after currently existing highest rule. A listener can't have multiple rules with the same priority.
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
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

