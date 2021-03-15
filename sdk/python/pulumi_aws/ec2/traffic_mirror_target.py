# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['TrafficMirrorTargetArgs', 'TrafficMirrorTarget']

@pulumi.input_type
class TrafficMirrorTargetArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 network_load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TrafficMirrorTarget resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: The network interface ID that is associated with the target.
        :param pulumi.Input[str] network_load_balancer_arn: The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if network_load_balancer_arn is not None:
            pulumi.set(__self__, "network_load_balancer_arn", network_load_balancer_arn)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the traffic mirror session.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The network interface ID that is associated with the target.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="networkLoadBalancerArn")
    def network_load_balancer_arn(self) -> Optional[pulumi.Input[str]]:
        """
        The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
        """
        return pulumi.get(self, "network_load_balancer_arn")

    @network_load_balancer_arn.setter
    def network_load_balancer_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_load_balancer_arn", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class TrafficMirrorTarget(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TrafficMirrorTargetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides an Traffic mirror target.\
        Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

        ## Example Usage

        To create a basic traffic mirror session

        ```python
        import pulumi
        import pulumi_aws as aws

        nlb = aws.ec2.TrafficMirrorTarget("nlb",
            description="NLB target",
            network_load_balancer_arn=aws_lb["lb"]["arn"])
        eni = aws.ec2.TrafficMirrorTarget("eni",
            description="ENI target",
            network_interface_id=aws_instance["test"]["primary_network_interface_id"])
        ```

        ## Import

        Traffic mirror targets can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/trafficMirrorTarget:TrafficMirrorTarget target tmt-0c13a005422b86606
        ```

        :param str resource_name: The name of the resource.
        :param TrafficMirrorTargetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 network_load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Traffic mirror target.\
        Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

        ## Example Usage

        To create a basic traffic mirror session

        ```python
        import pulumi
        import pulumi_aws as aws

        nlb = aws.ec2.TrafficMirrorTarget("nlb",
            description="NLB target",
            network_load_balancer_arn=aws_lb["lb"]["arn"])
        eni = aws.ec2.TrafficMirrorTarget("eni",
            description="ENI target",
            network_interface_id=aws_instance["test"]["primary_network_interface_id"])
        ```

        ## Import

        Traffic mirror targets can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/trafficMirrorTarget:TrafficMirrorTarget target tmt-0c13a005422b86606
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: The network interface ID that is associated with the target.
        :param pulumi.Input[str] network_load_balancer_arn: The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TrafficMirrorTargetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 network_load_balancer_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
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

            __props__['description'] = description
            __props__['network_interface_id'] = network_interface_id
            __props__['network_load_balancer_arn'] = network_load_balancer_arn
            __props__['tags'] = tags
            __props__['arn'] = None
        super(TrafficMirrorTarget, __self__).__init__(
            'aws:ec2/trafficMirrorTarget:TrafficMirrorTarget',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            network_load_balancer_arn: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'TrafficMirrorTarget':
        """
        Get an existing TrafficMirrorTarget resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the traffic mirror target.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: The network interface ID that is associated with the target.
        :param pulumi.Input[str] network_load_balancer_arn: The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["description"] = description
        __props__["network_interface_id"] = network_interface_id
        __props__["network_load_balancer_arn"] = network_load_balancer_arn
        __props__["tags"] = tags
        return TrafficMirrorTarget(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the traffic mirror target.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the traffic mirror session.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[Optional[str]]:
        """
        The network interface ID that is associated with the target.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkLoadBalancerArn")
    def network_load_balancer_arn(self) -> pulumi.Output[Optional[str]]:
        """
        The Amazon Resource Name (ARN) of the Network Load Balancer that is associated with the target.
        """
        return pulumi.get(self, "network_load_balancer_arn")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value map of resource tags.
        """
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

