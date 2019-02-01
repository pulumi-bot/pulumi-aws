# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import pulumi
import pulumi.runtime
from .. import utilities, tables

class InstanceGroup(pulumi.CustomResource):
    cluster_id: pulumi.Output[str]
    """
    ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
    """
    ebs_configs: pulumi.Output[list]
    """
    One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
    """
    ebs_optimized: pulumi.Output[bool]
    """
    Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
    """
    instance_count: pulumi.Output[int]
    """
    Target number of instances for the instance group. Defaults to 0.
    """
    instance_type: pulumi.Output[str]
    """
    The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    Human friendly name given to the instance group. Changing this forces a new resource to be created.
    """
    running_instance_count: pulumi.Output[int]
    status: pulumi.Output[str]
    def __init__(__self__, __name__, __opts__=None, cluster_id=None, ebs_configs=None, ebs_optimized=None, instance_count=None, instance_type=None, name=None):
        """
        Provides an Elastic MapReduce Cluster Instance Group configuration.
        See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.
        
        > **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
        web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
        Terraform will resize any Instance Group to zero when destroying the resource.
        
        :param str __name__: The name of the resource.
        :param pulumi.ResourceOptions __opts__: Options for the resource.
        :param pulumi.Input[str] cluster_id: ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        :param pulumi.Input[list] ebs_configs: One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] ebs_optimized: Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        :param pulumi.Input[int] instance_count: Target number of instances for the instance group. Defaults to 0.
        :param pulumi.Input[str] instance_type: The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Human friendly name given to the instance group. Changing this forces a new resource to be created.
        """
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not cluster_id:
            raise TypeError('Missing required property cluster_id')
        __props__['cluster_id'] = cluster_id

        __props__['ebs_configs'] = ebs_configs

        __props__['ebs_optimized'] = ebs_optimized

        __props__['instance_count'] = instance_count

        if not instance_type:
            raise TypeError('Missing required property instance_type')
        __props__['instance_type'] = instance_type

        __props__['name'] = name

        __props__['running_instance_count'] = None
        __props__['status'] = None

        super(InstanceGroup, __self__).__init__(
            'aws:emr/instanceGroup:InstanceGroup',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

