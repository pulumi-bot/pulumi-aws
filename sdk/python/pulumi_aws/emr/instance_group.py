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

__all__ = ['InstanceGroupArgs', 'InstanceGroup']

@pulumi.input_type
class InstanceGroupArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 instance_type: pulumi.Input[str],
                 autoscaling_policy: Optional[pulumi.Input[str]] = None,
                 bid_price: Optional[pulumi.Input[str]] = None,
                 configurations_json: Optional[pulumi.Input[str]] = None,
                 ebs_configs: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupEbsConfigArgs']]]] = None,
                 ebs_optimized: Optional[pulumi.Input[bool]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a InstanceGroup resource.
        :param pulumi.Input[str] cluster_id: ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        :param pulumi.Input[str] instance_type: The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] autoscaling_policy: The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        :param pulumi.Input[str] bid_price: If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        :param pulumi.Input[str] configurations_json: A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        :param pulumi.Input[Sequence[pulumi.Input['InstanceGroupEbsConfigArgs']]] ebs_configs: One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] ebs_optimized: Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        :param pulumi.Input[int] instance_count: target number of instances for the instance group. defaults to 0.
        :param pulumi.Input[str] name: Human friendly name given to the instance group. Changing this forces a new resource to be created.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "instance_type", instance_type)
        if autoscaling_policy is not None:
            pulumi.set(__self__, "autoscaling_policy", autoscaling_policy)
        if bid_price is not None:
            pulumi.set(__self__, "bid_price", bid_price)
        if configurations_json is not None:
            pulumi.set(__self__, "configurations_json", configurations_json)
        if ebs_configs is not None:
            pulumi.set(__self__, "ebs_configs", ebs_configs)
        if ebs_optimized is not None:
            pulumi.set(__self__, "ebs_optimized", ebs_optimized)
        if instance_count is not None:
            pulumi.set(__self__, "instance_count", instance_count)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        ID of the EMR Cluster to attach to. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        The EC2 instance type for all instances in the instance group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="autoscalingPolicy")
    def autoscaling_policy(self) -> Optional[pulumi.Input[str]]:
        """
        The autoscaling policy document. This is a JSON formatted string. See [EMR Auto Scaling](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html)
        """
        return pulumi.get(self, "autoscaling_policy")

    @autoscaling_policy.setter
    def autoscaling_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "autoscaling_policy", value)

    @property
    @pulumi.getter(name="bidPrice")
    def bid_price(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
        """
        return pulumi.get(self, "bid_price")

    @bid_price.setter
    def bid_price(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bid_price", value)

    @property
    @pulumi.getter(name="configurationsJson")
    def configurations_json(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON string for supplying list of configurations specific to the EMR instance group. Note that this can only be changed when using EMR release 5.21 or later.
        """
        return pulumi.get(self, "configurations_json")

    @configurations_json.setter
    def configurations_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "configurations_json", value)

    @property
    @pulumi.getter(name="ebsConfigs")
    def ebs_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupEbsConfigArgs']]]]:
        """
        One or more `ebs_config` blocks as defined below. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ebs_configs")

    @ebs_configs.setter
    def ebs_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InstanceGroupEbsConfigArgs']]]]):
        pulumi.set(self, "ebs_configs", value)

    @property
    @pulumi.getter(name="ebsOptimized")
    def ebs_optimized(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether an Amazon EBS volume is EBS-optimized. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "ebs_optimized")

    @ebs_optimized.setter
    def ebs_optimized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ebs_optimized", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        """
        target number of instances for the instance group. defaults to 0.
        """
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Human friendly name given to the instance group. Changing this forces a new resource to be created.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class InstanceGroup(pulumi.CustomResource):
    @overload
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

        ## Import

        EMR task instance group can be imported using their EMR Cluster id and Instance Group id separated by a forward-slash `/`, e.g.

        ```sh
         $ pulumi import aws:emr/instanceGroup:InstanceGroup task_greoup j-123456ABCDEF/ig-15EK4O09RZLNR
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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InstanceGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
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

        ## Import

        EMR task instance group can be imported using their EMR Cluster id and Instance Group id separated by a forward-slash `/`, e.g.

        ```sh
         $ pulumi import aws:emr/instanceGroup:InstanceGroup task_greoup j-123456ABCDEF/ig-15EK4O09RZLNR
        ```

        :param str resource_name: The name of the resource.
        :param InstanceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InstanceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
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
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__['cluster_id'] = cluster_id
            __props__['configurations_json'] = configurations_json
            __props__['ebs_configs'] = ebs_configs
            __props__['ebs_optimized'] = ebs_optimized
            __props__['instance_count'] = instance_count
            if instance_type is None and not opts.urn:
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

