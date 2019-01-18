# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
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
    The ARN of the ELB
    """
    availability_zones: pulumi.Output[list]
    """
    The AZ's to serve traffic in.
    """
    connection_draining: pulumi.Output[bool]
    """
    Boolean to enable connection draining. Default: `false`
    """
    connection_draining_timeout: pulumi.Output[int]
    """
    The time in seconds to allow for connections to drain. Default: `300`
    """
    cross_zone_load_balancing: pulumi.Output[bool]
    """
    Enable cross-zone load balancing. Default: `true`
    """
    dns_name: pulumi.Output[str]
    """
    The DNS name of the ELB
    """
    health_check: pulumi.Output[dict]
    """
    A health_check block. Health Check documented below.
    """
    idle_timeout: pulumi.Output[int]
    """
    The time in seconds that the connection is allowed to be idle. Default: `60`
    """
    instances: pulumi.Output[list]
    """
    A list of instance ids to place in the ELB pool.
    """
    internal: pulumi.Output[bool]
    """
    If true, ELB will be an internal ELB.
    """
    listeners: pulumi.Output[list]
    """
    A list of listener blocks. Listeners documented below.
    """
    name: pulumi.Output[str]
    """
    The name of the ELB. By default generated by Terraform.
    """
    name_prefix: pulumi.Output[str]
    """
    Creates a unique name beginning with the specified
    prefix. Conflicts with `name`.
    """
    security_groups: pulumi.Output[list]
    """
    A list of security group IDs to assign to the ELB.
    Only valid if creating an ELB within a VPC
    """
    source_security_group: pulumi.Output[str]
    """
    The name of the security group that you can use as
    part of your inbound rules for your load balancer's back-end application
    instances. Use this for Classic or Default VPC only.
    """
    source_security_group_id: pulumi.Output[str]
    """
    The ID of the security group that you can use as
    part of your inbound rules for your load balancer's back-end application
    instances. Only available on ELBs launched in a VPC.
    """
    subnets: pulumi.Output[list]
    """
    A list of subnet IDs to attach to the ELB.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    zone_id: pulumi.Output[str]
    """
    The canonical hosted zone ID of the ELB (to be used in a Route 53 Alias record)
    """
    def __init__(__self__, __name__, __opts__=None, access_logs=None, availability_zones=None, connection_draining=None, connection_draining_timeout=None, cross_zone_load_balancing=None, health_check=None, idle_timeout=None, instances=None, internal=None, listeners=None, name=None, name_prefix=None, security_groups=None, source_security_group=None, subnets=None, tags=None):
        """
        Provides an Elastic Load Balancer resource, also known as a "Classic
        Load Balancer" after the release of
        [Application/Network Load Balancers](https://www.terraform.io/docs/providers/aws/r/lb.html).
        
        > **NOTE on ELB Instances and ELB Attachments:** Terraform currently
        provides both a standalone ELB Attachment resource
        (describing an instance attached to an ELB), and an ELB resource with
        `instances` defined in-line. At this time you cannot use an ELB with in-line
        instances in conjunction with a ELB Attachment resources. Doing so will cause a
        conflict and will overwrite attachments.
        
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[dict] access_logs: An Access Logs block. Access Logs documented below.
        :param pulumi.Input[list] availability_zones: The AZ's to serve traffic in.
        :param pulumi.Input[bool] connection_draining: Boolean to enable connection draining. Default: `false`
        :param pulumi.Input[int] connection_draining_timeout: The time in seconds to allow for connections to drain. Default: `300`
        :param pulumi.Input[bool] cross_zone_load_balancing: Enable cross-zone load balancing. Default: `true`
        :param pulumi.Input[dict] health_check: A health_check block. Health Check documented below.
        :param pulumi.Input[int] idle_timeout: The time in seconds that the connection is allowed to be idle. Default: `60`
        :param pulumi.Input[list] instances: A list of instance ids to place in the ELB pool.
        :param pulumi.Input[bool] internal: If true, ELB will be an internal ELB.
        :param pulumi.Input[list] listeners: A list of listener blocks. Listeners documented below.
        :param pulumi.Input[str] name: The name of the ELB. By default generated by Terraform.
        :param pulumi.Input[str] name_prefix: Creates a unique name beginning with the specified
               prefix. Conflicts with `name`.
        :param pulumi.Input[list] security_groups: A list of security group IDs to assign to the ELB.
               Only valid if creating an ELB within a VPC
        :param pulumi.Input[str] source_security_group: The name of the security group that you can use as
               part of your inbound rules for your load balancer's back-end application
               instances. Use this for Classic or Default VPC only.
        :param pulumi.Input[list] subnets: A list of subnet IDs to attach to the ELB.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['access_logs'] = access_logs

        __props__['availability_zones'] = availability_zones

        __props__['connection_draining'] = connection_draining

        __props__['connection_draining_timeout'] = connection_draining_timeout

        __props__['cross_zone_load_balancing'] = cross_zone_load_balancing

        __props__['health_check'] = health_check

        __props__['idle_timeout'] = idle_timeout

        __props__['instances'] = instances

        __props__['internal'] = internal

        if not listeners:
            raise TypeError('Missing required property listeners')
        __props__['listeners'] = listeners

        __props__['name'] = name

        __props__['name_prefix'] = name_prefix

        __props__['security_groups'] = security_groups

        __props__['source_security_group'] = source_security_group

        __props__['subnets'] = subnets

        __props__['tags'] = tags

        __props__['arn'] = None
        __props__['dns_name'] = None
        __props__['source_security_group_id'] = None
        __props__['zone_id'] = None

        super(LoadBalancer, __self__).__init__(
            'aws:elasticloadbalancing/loadBalancer:LoadBalancer',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

