# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class TargetGroup(pulumi.CustomResource):
    arn: pulumi.Output[str]
    """
    The ARN of the Target Group (matches `id`)
    """
    arn_suffix: pulumi.Output[str]
    """
    The ARN suffix for use with CloudWatch Metrics.
    """
    deregistration_delay: pulumi.Output[float]
    """
    The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
    """
    health_check: pulumi.Output[dict]
    """
    A Health Check block. Health Check blocks are documented below.
    """
    lambda_multi_value_headers_enabled: pulumi.Output[bool]
    """
    Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`.
    """
    name: pulumi.Output[str]
    """
    The name of the target group. If omitted, this provider will assign a random, unique name.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
    """
    port: pulumi.Output[float]
    """
    The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
    """
    protocol: pulumi.Output[str]
    """
    The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `target_type` is `lambda`.
    """
    proxy_protocol_v2: pulumi.Output[bool]
    """
    Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
    """
    slow_start: pulumi.Output[float]
    """
    The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
    """
    stickiness: pulumi.Output[dict]
    """
    A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    target_type: pulumi.Output[str]
    """
    The type of target that you must specify when registering targets with this target group.
    The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
    The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
    If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
    the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
    You can't specify publicly routable IP addresses.
    """
    vpc_id: pulumi.Output[str]
    """
    The identifier of the VPC in which to create the target group. Required when `target_type` is `instance` or `ip`. Does not apply when `target_type` is `lambda`.
    """
    def __init__(__self__, resource_name, opts=None, deregistration_delay=None, health_check=None, lambda_multi_value_headers_enabled=None, name=None, name_prefix=None, port=None, protocol=None, proxy_protocol_v2=None, slow_start=None, stickiness=None, tags=None, target_type=None, vpc_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a Target Group resource for use with Load Balancer resources.
        
        > **Note:** `alb.TargetGroup` is known as `lb.TargetGroup`. The functionality is identical.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] deregistration_delay: The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
        :param pulumi.Input[dict] health_check: A Health Check block. Health Check blocks are documented below.
        :param pulumi.Input[bool] lambda_multi_value_headers_enabled: Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`.
        :param pulumi.Input[str] name: The name of the target group. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
        :param pulumi.Input[float] port: The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
        :param pulumi.Input[str] protocol: The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `target_type` is `lambda`.
        :param pulumi.Input[bool] proxy_protocol_v2: Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
        :param pulumi.Input[float] slow_start: The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
        :param pulumi.Input[dict] stickiness: A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_type: The type of target that you must specify when registering targets with this target group.
               The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
               The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
               If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
               the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
               You can't specify publicly routable IP addresses.
        :param pulumi.Input[str] vpc_id: The identifier of the VPC in which to create the target group. Required when `target_type` is `instance` or `ip`. Does not apply when `target_type` is `lambda`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/alb_target_group.html.markdown.
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

            __props__['deregistration_delay'] = deregistration_delay
            __props__['health_check'] = health_check
            __props__['lambda_multi_value_headers_enabled'] = lambda_multi_value_headers_enabled
            __props__['name'] = name
            __props__['name_prefix'] = name_prefix
            __props__['port'] = port
            __props__['protocol'] = protocol
            __props__['proxy_protocol_v2'] = proxy_protocol_v2
            __props__['slow_start'] = slow_start
            __props__['stickiness'] = stickiness
            __props__['tags'] = tags
            __props__['target_type'] = target_type
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['arn_suffix'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:applicationloadbalancing/targetGroup:TargetGroup")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(TargetGroup, __self__).__init__(
            'aws:alb/targetGroup:TargetGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, arn=None, arn_suffix=None, deregistration_delay=None, health_check=None, lambda_multi_value_headers_enabled=None, name=None, name_prefix=None, port=None, protocol=None, proxy_protocol_v2=None, slow_start=None, stickiness=None, tags=None, target_type=None, vpc_id=None):
        """
        Get an existing TargetGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the Target Group (matches `id`)
        :param pulumi.Input[str] arn_suffix: The ARN suffix for use with CloudWatch Metrics.
        :param pulumi.Input[float] deregistration_delay: The amount time for Elastic Load Balancing to wait before changing the state of a deregistering target from draining to unused. The range is 0-3600 seconds. The default value is 300 seconds.
        :param pulumi.Input[dict] health_check: A Health Check block. Health Check blocks are documented below.
        :param pulumi.Input[bool] lambda_multi_value_headers_enabled: Boolean whether the request and response headers exchanged between the load balancer and the Lambda function include arrays of values or strings. Only applies when `target_type` is `lambda`.
        :param pulumi.Input[str] name: The name of the target group. If omitted, this provider will assign a random, unique name.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`. Cannot be longer than 6 characters.
        :param pulumi.Input[float] port: The port to use to connect with the target. Valid values are either ports 1-65536, or `traffic-port`. Defaults to `traffic-port`.
        :param pulumi.Input[str] protocol: The protocol to use to connect with the target. Defaults to `HTTP`. Not applicable when `target_type` is `lambda`.
        :param pulumi.Input[bool] proxy_protocol_v2: Boolean to enable / disable support for proxy protocol v2 on Network Load Balancers. See [doc](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-target-groups.html#proxy-protocol) for more information.
        :param pulumi.Input[float] slow_start: The amount time for targets to warm up before the load balancer sends them a full share of requests. The range is 30-900 seconds or 0 to disable. The default value is 0 seconds.
        :param pulumi.Input[dict] stickiness: A Stickiness block. Stickiness blocks are documented below. `stickiness` is only valid if used with Load Balancers of type `Application`
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] target_type: The type of target that you must specify when registering targets with this target group.
               The possible values are `instance` (targets are specified by instance ID) or `ip` (targets are specified by IP address) or `lambda` (targets are specified by lambda arn).
               The default is `instance`. Note that you can't specify targets for a target group using both instance IDs and IP addresses.
               If the target type is `ip`, specify IP addresses from the subnets of the virtual private cloud (VPC) for the target group,
               the RFC 1918 range (10.0.0.0/8, 172.16.0.0/12, and 192.168.0.0/16), and the RFC 6598 range (100.64.0.0/10).
               You can't specify publicly routable IP addresses.
        :param pulumi.Input[str] vpc_id: The identifier of the VPC in which to create the target group. Required when `target_type` is `instance` or `ip`. Does not apply when `target_type` is `lambda`.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/alb_target_group.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["arn"] = arn
        __props__["arn_suffix"] = arn_suffix
        __props__["deregistration_delay"] = deregistration_delay
        __props__["health_check"] = health_check
        __props__["lambda_multi_value_headers_enabled"] = lambda_multi_value_headers_enabled
        __props__["name"] = name
        __props__["name_prefix"] = name_prefix
        __props__["port"] = port
        __props__["protocol"] = protocol
        __props__["proxy_protocol_v2"] = proxy_protocol_v2
        __props__["slow_start"] = slow_start
        __props__["stickiness"] = stickiness
        __props__["tags"] = tags
        __props__["target_type"] = target_type
        __props__["vpc_id"] = vpc_id
        return TargetGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

