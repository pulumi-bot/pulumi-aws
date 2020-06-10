# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class TrafficMirrorSession(pulumi.CustomResource):
    description: pulumi.Output[str]
    """
    A description of the traffic mirror session.
    """
    network_interface_id: pulumi.Output[str]
    """
    ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
    """
    packet_length: pulumi.Output[float]
    """
    The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
    """
    session_number: pulumi.Output[float]
    """
    - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets. 
    """
    tags: pulumi.Output[dict]
    """
    Key-value map of resource tags.
    """
    traffic_mirror_filter_id: pulumi.Output[str]
    """
    ID of the traffic mirror filter to be used
    """
    traffic_mirror_target_id: pulumi.Output[str]
    """
    ID of the traffic mirror target to be used
    """
    virtual_network_id: pulumi.Output[float]
    """
    - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
    """
    def __init__(__self__, resource_name, opts=None, description=None, network_interface_id=None, packet_length=None, session_number=None, tags=None, traffic_mirror_filter_id=None, traffic_mirror_target_id=None, virtual_network_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides an Traffic mirror session.  
        Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

        ## Example Usage



        ```python
        import pulumi
        import pulumi_aws as aws

        filter = aws.ec2.TrafficMirrorFilter("filter",
            description="traffic mirror filter - example",
            network_services=["amazon-dns"])
        target = aws.ec2.TrafficMirrorTarget("target", network_load_balancer_arn=aws_lb["lb"]["arn"])
        session = aws.ec2.TrafficMirrorSession("session",
            description="traffic mirror session - example",
            network_interface_id=aws_instance["test"]["primary_network_interface_id"],
            traffic_mirror_filter_id=filter.id,
            traffic_mirror_target_id=target.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[float] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[float] session_number: - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets. 
        :param pulumi.Input[dict] tags: Key-value map of resource tags.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[float] virtual_network_id: - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
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

            __props__['description'] = description
            if network_interface_id is None:
                raise TypeError("Missing required property 'network_interface_id'")
            __props__['network_interface_id'] = network_interface_id
            __props__['packet_length'] = packet_length
            if session_number is None:
                raise TypeError("Missing required property 'session_number'")
            __props__['session_number'] = session_number
            __props__['tags'] = tags
            if traffic_mirror_filter_id is None:
                raise TypeError("Missing required property 'traffic_mirror_filter_id'")
            __props__['traffic_mirror_filter_id'] = traffic_mirror_filter_id
            if traffic_mirror_target_id is None:
                raise TypeError("Missing required property 'traffic_mirror_target_id'")
            __props__['traffic_mirror_target_id'] = traffic_mirror_target_id
            __props__['virtual_network_id'] = virtual_network_id
        super(TrafficMirrorSession, __self__).__init__(
            'aws:ec2/trafficMirrorSession:TrafficMirrorSession',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, description=None, network_interface_id=None, packet_length=None, session_number=None, tags=None, traffic_mirror_filter_id=None, traffic_mirror_target_id=None, virtual_network_id=None):
        """
        Get an existing TrafficMirrorSession resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the traffic mirror session.
        :param pulumi.Input[str] network_interface_id: ID of the source network interface. Not all network interfaces are eligible as mirror sources. On EC2 instances only nitro based instances support mirroring.
        :param pulumi.Input[float] packet_length: The number of bytes in each packet to mirror. These are bytes after the VXLAN header. Do not specify this parameter when you want to mirror the entire packet. To mirror a subset of the packet, set this to the length (in bytes) that you want to mirror.
        :param pulumi.Input[float] session_number: - The session number determines the order in which sessions are evaluated when an interface is used by multiple sessions. The first session with a matching filter is the one that mirrors the packets. 
        :param pulumi.Input[dict] tags: Key-value map of resource tags.
        :param pulumi.Input[str] traffic_mirror_filter_id: ID of the traffic mirror filter to be used
        :param pulumi.Input[str] traffic_mirror_target_id: ID of the traffic mirror target to be used
        :param pulumi.Input[float] virtual_network_id: - The VXLAN ID for the Traffic Mirror session. For more information about the VXLAN protocol, see RFC 7348. If you do not specify a VirtualNetworkId, an account-wide unique id is chosen at random.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["network_interface_id"] = network_interface_id
        __props__["packet_length"] = packet_length
        __props__["session_number"] = session_number
        __props__["tags"] = tags
        __props__["traffic_mirror_filter_id"] = traffic_mirror_filter_id
        __props__["traffic_mirror_target_id"] = traffic_mirror_target_id
        __props__["virtual_network_id"] = virtual_network_id
        return TrafficMirrorSession(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
