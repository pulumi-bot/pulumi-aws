# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class LoadBalancer(pulumi.CustomResource):
    access_logs: pulumi.Output[dict]
    """
    An Access Logs block. Access Logs documented below.
    """
    arn: pulumi.Output[str]
    """
    The ARN of the load balancer (matches `id`).
    """
    arn_suffix: pulumi.Output[str]
    """
    The ARN suffix for use with CloudWatch Metrics.
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name of the load balancer.
    """
    enable_cross_zone_load_balancing: pulumi.Output[bool]
    """
    If true, cross-zone load balancing of the load balancer will be enabled.
    This is a `network` load balancer feature. Defaults to `false`.
    """
    enable_deletion_protection: pulumi.Output[bool]
    """
    If true, deletion of the load balancer will be disabled via
    the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
    """
    enable_http2: pulumi.Output[bool]
    """
    Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
    """
    idle_timeout: pulumi.Output[float]
    """
    The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
    """
    internal: pulumi.Output[bool]
    """
    If true, the LB will be internal.
    """
    ip_address_type: pulumi.Output[str]
    """
    The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
    """
    load_balancer_type: pulumi.Output[str]
    """
    The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
    """
    name: pulumi.Output[str]
    """
    The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
    must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
    this provider will autogenerate a name beginning with `tf-lb`.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified prefix. Conflicts with `name`.
    """
    security_groups: pulumi.Output[list]
    """
    A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
    """
    subnet_mappings: pulumi.Output[list]
    """
    A subnet mapping block as documented below.
    """
    subnets: pulumi.Output[list]
    """
    A list of subnet IDs to attach to the LB. Subnets
    cannot be updated for Load Balancers of type `network`. Changing this value
    for load balancers of type `network` will force a recreation of the resource.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    vpc_id: pulumi.Output[str]
    zone_id: pulumi.Output[str]
    """
    The canonical hosted zone ID of the load balancer (to be used in a Route 53 Alias record).
    """
    def __init__(__self__, resource_name, opts=None, access_logs=None, enable_cross_zone_load_balancing=None, enable_deletion_protection=None, enable_http2=None, idle_timeout=None, internal=None, ip_address_type=None, load_balancer_type=None, name=None, name_prefix=None, security_groups=None, subnet_mappings=None, subnets=None, tags=None, __name__=None, __opts__=None):
        """
        Provides a Load Balancer resource.
        
        > **Note:** `aws_alb` is known as `aws_lb`. The functionality is identical.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[dict] access_logs: An Access Logs block. Access Logs documented below.
        :param pulumi.Input[bool] enable_cross_zone_load_balancing: If true, cross-zone load balancing of the load balancer will be enabled.
               This is a `network` load balancer feature. Defaults to `false`.
        :param pulumi.Input[bool] enable_deletion_protection: If true, deletion of the load balancer will be disabled via
               the AWS API. This will prevent this provider from deleting the load balancer. Defaults to `false`.
        :param pulumi.Input[bool] enable_http2: Indicates whether HTTP/2 is enabled in `application` load balancers. Defaults to `true`.
        :param pulumi.Input[float] idle_timeout: The time in seconds that the connection is allowed to be idle. Only valid for Load Balancers of type `application`. Default: 60.
        :param pulumi.Input[bool] internal: If true, the LB will be internal.
        :param pulumi.Input[str] ip_address_type: The type of IP addresses used by the subnets for your load balancer. The possible values are `ipv4` and `dualstack`
        :param pulumi.Input[str] load_balancer_type: The type of load balancer to create. Possible values are `application` or `network`. The default value is `application`.
        :param pulumi.Input[str] name: The name of the LB. This name must be unique within your AWS account, can have a maximum of 32 characters,
               must contain only alphanumeric characters or hyphens, and must not begin or end with a hyphen. If not specified,
               this provider will autogenerate a name beginning with `tf-lb`.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified prefix. Conflicts with `name`.
        :param pulumi.Input[list] security_groups: A list of security group IDs to assign to the LB. Only valid for Load Balancers of type `application`.
        :param pulumi.Input[list] subnet_mappings: A subnet mapping block as documented below.
        :param pulumi.Input[list] subnets: A list of subnet IDs to attach to the LB. Subnets
               cannot be updated for Load Balancers of type `network`. Changing this value
               for load balancers of type `network` will force a recreation of the resource.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/alb.html.markdown.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_logs'] = access_logs
        __props__['enable_cross_zone_load_balancing'] = enable_cross_zone_load_balancing
        __props__['enable_deletion_protection'] = enable_deletion_protection
        __props__['enable_http2'] = enable_http2
        __props__['idle_timeout'] = idle_timeout
        __props__['internal'] = internal
        __props__['ip_address_type'] = ip_address_type
        __props__['load_balancer_type'] = load_balancer_type
        __props__['name'] = name
        __props__['name_prefix'] = name_prefix
        __props__['security_groups'] = security_groups
        __props__['subnet_mappings'] = subnet_mappings
        __props__['subnets'] = subnets
        __props__['tags'] = tags
        __props__['arn'] = None
        __props__['arn_suffix'] = None
        __props__['dns_name'] = None
        __props__['vpc_id'] = None
        __props__['zone_id'] = None

        if opts is None:
            opts = pulumi.ResourceOptions()
        if opts.version is None:
            opts.version = utilities.get_version()
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="aws:applicationloadbalancing/loadBalancer:LoadBalancer")])
        opts = alias_opts if opts is None else opts.merge(alias_opts)
        super(LoadBalancer, __self__).__init__(
            'aws:alb/loadBalancer:LoadBalancer',
            resource_name,
            __props__,
            opts)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

