# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['InstanceGroup']


class InstanceGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscaling_policy: Optional[pulumi.Input[str]] = None,
                 bid_price: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 configurations_json: Optional[pulumi.Input[str]] = None,
                 ebs_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupEbsConfigArgs']]]]] = None,
                 ebs_optimized: Optional[pulumi.Input[bool]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Elastic MapReduce Cluster Instance Group configuration.
        See [Amazon Elastic MapReduce Documentation](https://aws.amazon.com/documentation/emr/) for more information.

        > **NOTE:** At this time, Instance Groups cannot be destroyed through the API nor
        web interface. Instance Groups are destroyed when the EMR Cluster is destroyed.
        this provider will resize any Instance Group to zero when destroying the resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        task = aws.emr.InstanceGroup("task",
            cluster_id=aws_emr_cluster["tf-test-cluster"]["id"],
            instance_count=1,
            instance_type="m5.xlarge")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_policy: The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        :param pulumi.Input[str] bid_price: If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        :param pulumi.Input[str] cluster_id: ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] configurations_json: A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupEbsConfigArgs']]]] ebs_configs: One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] ebs_optimized: Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        :param pulumi.Input[int] instance_count: target number of instances for the instance group. defaults to 0.
        :param pulumi.Input[str] instance_type: The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Human friendly name given to the instance group. Changing this forces a new resource to be created.
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

            __props__['autoscaling_policy'] = autoscaling_policy
            __props__['bid_price'] = bid_price
            if cluster_id is None:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['configurations_json'] = configurations_json
            __props__['ebs_configs'] = ebs_configs
            __props__['ebs_optimized'] = ebs_optimized
            __props__['instance_count'] = instance_count
            if instance_type is None:
                raise TypeError("Missing required property 'instance_type'")
            __props__['instance_type'] = instance_type
            __props__['name'] = name
            __props__['running_instance_count'] = None
            __props__['status'] = None
        super(InstanceGroup, __self__).__init__(
            'aws:emr/instanceGroup:InstanceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            autoscaling_policy: Optional[pulumi.Input[str]] = None,
            bid_price: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            configurations_json: Optional[pulumi.Input[str]] = None,
            ebs_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupEbsConfigArgs']]]]] = None,
            ebs_optimized: Optional[pulumi.Input[bool]] = None,
            instance_count: Optional[pulumi.Input[int]] = None,
            instance_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            running_instance_count: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'InstanceGroup':
        """
        Get an existing InstanceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] autoscaling_policy: The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        :param pulumi.Input[str] bid_price: If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        :param pulumi.Input[str] cluster_id: ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] configurations_json: A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['InstanceGroupEbsConfigArgs']]]] ebs_configs: One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] ebs_optimized: Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        :param pulumi.Input[int] instance_count: target number of instances for the instance group. defaults to 0.
        :param pulumi.Input[str] instance_type: The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: Human friendly name given to the instance group. Changing this forces a new resource to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["autoscaling_policy"] = autoscaling_policy
        __props__["bid_price"] = bid_price
        __props__["cluster_id"] = cluster_id
        __props__["configurations_json"] = configurations_json
        __props__["ebs_configs"] = ebs_configs
        __props__["ebs_optimized"] = ebs_optimized
        __props__["instance_count"] = instance_count
        __props__["instance_type"] = instance_type
        __props__["name"] = name
        __props__["running_instance_count"] = running_instance_count
        __props__["status"] = status
        return InstanceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> pulumi.Output[Optional[str]]:
        """
        The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        """
        return pulumi.get(self, "autoscaling_policy")

    @property
    @pulumi.getter(name="bidPrice")
    def bid_price(self) -> pulumi.Output[Optional[str]]:
        """
        If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        """
        return pulumi.get(self, "bid_price")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="configurationsJson")
    def configurations_json(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        """
        return pulumi.get(self, "configurations_json")

    @property
    @pulumi.getter(name="ebsConfigs")
    def ebs_configs(self) -> pulumi.Output[Sequence['outputs.InstanceGroupEbsConfig']]:
        """
        One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ebs_configs")

    @property
    @pulumi.getter(name="ebsOptimized")
    def ebs_optimized(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ebs_optimized")

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> pulumi.Output[Optional[int]]:
        """
        target number of instances for the instance group. defaults to 0.
        """
        return pulumi.get(self, "instance_count")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "instance_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Human friendly name given to the instance group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="runningInstanceCount")
    def running_instance_count(self) -> pulumi.Output[int]:
        return pulumi.get(self, "running_instance_count")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        return pulumi.get(self, "status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

