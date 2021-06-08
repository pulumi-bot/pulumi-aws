# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CustomerGatewayArgs', 'CustomerGateway']

@pulumi.input_type
class CustomerGatewayArgs:
    def __init__(__self__, *,
                 bgp_asn: pulumi.Input[str],
                 ip_address: pulumi.Input[str],
                 type: pulumi.Input[str],
                 device_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a CustomerGateway resource.
        :param pulumi.Input[str] bgp_asn: The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        :param pulumi.Input[str] ip_address: The IP address of the gateway's Internet-routable external interface.
        :param pulumi.Input[str] type: The type of customer gateway. The only type AWS
               supports at this time is "ipsec.1".
        :param pulumi.Input[str] device_name: A name for the customer gateway device.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        """
        pulumi.set(__self__, "bgp_asn", bgp_asn)
        pulumi.set(__self__, "ip_address", ip_address)
        pulumi.set(__self__, "type", type)
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> pulumi.Input[str]:
        """
        The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: pulumi.Input[str]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Input[str]:
        """
        The IP address of the gateway's Internet-routable external interface.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of customer gateway. The only type AWS
        supports at this time is "ipsec.1".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the customer gateway device.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _CustomerGatewayState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomerGateway resources.
        :param pulumi.Input[str] arn: The ARN of the customer gateway.
        :param pulumi.Input[str] bgp_asn: The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        :param pulumi.Input[str] device_name: A name for the customer gateway device.
        :param pulumi.Input[str] ip_address: The IP address of the gateway's Internet-routable external interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] type: The type of customer gateway. The only type AWS
               supports at this time is "ipsec.1".
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if bgp_asn is not None:
            pulumi.set(__self__, "bgp_asn", bgp_asn)
        if device_name is not None:
            pulumi.set(__self__, "device_name", device_name)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        The ARN of the customer gateway.
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> Optional[pulumi.Input[str]]:
        """
        The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        """
        return pulumi.get(self, "bgp_asn")

    @bgp_asn.setter
    def bgp_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bgp_asn", value)

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for the customer gateway device.
        """
        return pulumi.get(self, "device_name")

    @device_name.setter
    def device_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_name", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address of the gateway's Internet-routable external interface.
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of customer gateway. The only type AWS
        supports at this time is "ipsec.1".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class CustomerGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.CustomerGateway("main",
            bgp_asn="65000",
            ip_address="172.83.124.10",
            tags={
                "Name": "main-customer-gateway",
            },
            type="ipsec.1")
        ```

        ## Import

        Customer Gateways can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/customerGateway:CustomerGateway main cgw-b4dc3961
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bgp_asn: The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        :param pulumi.Input[str] device_name: A name for the customer gateway device.
        :param pulumi.Input[str] ip_address: The IP address of the gateway's Internet-routable external interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] type: The type of customer gateway. The only type AWS
               supports at this time is "ipsec.1".
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: CustomerGatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a customer gateway inside a VPC. These objects can be connected to VPN gateways via VPN connections, and allow you to establish tunnels between your network and the VPC.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.CustomerGateway("main",
            bgp_asn="65000",
            ip_address="172.83.124.10",
            tags={
                "Name": "main-customer-gateway",
            },
            type="ipsec.1")
        ```

        ## Import

        Customer Gateways can be imported using the `id`, e.g.

        ```sh
         $ pulumi import aws:ec2/customerGateway:CustomerGateway main cgw-b4dc3961
        ```

        :param str resource_name_: The name of the resource.
        :param CustomerGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomerGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bgp_asn: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomerGatewayArgs.__new__(CustomerGatewayArgs)

            if bgp_asn is None and not opts.urn:
                raise TypeError("Missing required property 'bgp_asn'")
            __props__.__dict__["bgp_asn"] = bgp_asn
            __props__.__dict__["device_name"] = device_name
            if ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'ip_address'")
            __props__.__dict__["ip_address"] = ip_address
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["arn"] = None
        super(CustomerGateway, __self__).__init__(
            'aws:ec2/customerGateway:CustomerGateway',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            bgp_asn: Optional[pulumi.Input[str]] = None,
            device_name: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'CustomerGateway':
        """
        Get an existing CustomerGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the customer gateway.
        :param pulumi.Input[str] bgp_asn: The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        :param pulumi.Input[str] device_name: A name for the customer gateway device.
        :param pulumi.Input[str] ip_address: The IP address of the gateway's Internet-routable external interface.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider .
        :param pulumi.Input[str] type: The type of customer gateway. The only type AWS
               supports at this time is "ipsec.1".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomerGatewayState.__new__(_CustomerGatewayState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["bgp_asn"] = bgp_asn
        __props__.__dict__["device_name"] = device_name
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["type"] = type
        return CustomerGateway(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the customer gateway.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bgpAsn")
    def bgp_asn(self) -> pulumi.Output[str]:
        """
        The gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
        """
        return pulumi.get(self, "bgp_asn")

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> pulumi.Output[Optional[str]]:
        """
        A name for the customer gateway device.
        """
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        The IP address of the gateway's Internet-routable external interface.
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Tags to apply to the gateway. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider .
        """
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of customer gateway. The only type AWS
        supports at this time is "ipsec.1".
        """
        return pulumi.get(self, "type")

