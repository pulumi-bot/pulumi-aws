# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetElasticIpResult',
    'AwaitableGetElasticIpResult',
    'get_elastic_ip',
]

@pulumi.output_type
class GetElasticIpResult:
    """
    A collection of values returned by getElasticIp.
    """
    def __init__(__self__, association_id=None, customer_owned_ip=None, customer_owned_ipv4_pool=None, domain=None, filters=None, id=None, instance_id=None, network_interface_id=None, network_interface_owner_id=None, private_dns=None, private_ip=None, public_dns=None, public_ip=None, public_ipv4_pool=None, tags=None):
        if association_id and not isinstance(association_id, str):
            raise TypeError("Expected argument 'association_id' to be a str")
        pulumi.set(__self__, "association_id", association_id)
        if customer_owned_ip and not isinstance(customer_owned_ip, str):
            raise TypeError("Expected argument 'customer_owned_ip' to be a str")
        pulumi.set(__self__, "customer_owned_ip", customer_owned_ip)
        if customer_owned_ipv4_pool and not isinstance(customer_owned_ipv4_pool, str):
            raise TypeError("Expected argument 'customer_owned_ipv4_pool' to be a str")
        pulumi.set(__self__, "customer_owned_ipv4_pool", customer_owned_ipv4_pool)
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        pulumi.set(__self__, "filters", filters)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_id and not isinstance(instance_id, str):
            raise TypeError("Expected argument 'instance_id' to be a str")
        pulumi.set(__self__, "instance_id", instance_id)
        if network_interface_id and not isinstance(network_interface_id, str):
            raise TypeError("Expected argument 'network_interface_id' to be a str")
        pulumi.set(__self__, "network_interface_id", network_interface_id)
        if network_interface_owner_id and not isinstance(network_interface_owner_id, str):
            raise TypeError("Expected argument 'network_interface_owner_id' to be a str")
        pulumi.set(__self__, "network_interface_owner_id", network_interface_owner_id)
        if private_dns and not isinstance(private_dns, str):
            raise TypeError("Expected argument 'private_dns' to be a str")
        pulumi.set(__self__, "private_dns", private_dns)
        if private_ip and not isinstance(private_ip, str):
            raise TypeError("Expected argument 'private_ip' to be a str")
        pulumi.set(__self__, "private_ip", private_ip)
        if public_dns and not isinstance(public_dns, str):
            raise TypeError("Expected argument 'public_dns' to be a str")
        pulumi.set(__self__, "public_dns", public_dns)
        if public_ip and not isinstance(public_ip, str):
            raise TypeError("Expected argument 'public_ip' to be a str")
        pulumi.set(__self__, "public_ip", public_ip)
        if public_ipv4_pool and not isinstance(public_ipv4_pool, str):
            raise TypeError("Expected argument 'public_ipv4_pool' to be a str")
        pulumi.set(__self__, "public_ipv4_pool", public_ipv4_pool)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> str:
        """
        The ID representing the association of the address with an instance in a VPC.
        """
        return pulumi.get(self, "association_id")

    @property
    @pulumi.getter(name="customerOwnedIp")
    def customer_owned_ip(self) -> str:
        """
        Customer Owned IP.
        """
        return pulumi.get(self, "customer_owned_ip")

    @property
    @pulumi.getter(name="customerOwnedIpv4Pool")
    def customer_owned_ipv4_pool(self) -> str:
        """
        The ID of a Customer Owned IP Pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
        """
        return pulumi.get(self, "customer_owned_ipv4_pool")

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        Indicates whether the address is for use in EC2-Classic (standard) or in a VPC (vpc).
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def filters(self) -> Optional[Sequence['outputs.GetElasticIpFilterResult']]:
        return pulumi.get(self, "filters")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        If VPC Elastic IP, the allocation identifier. If EC2-Classic Elastic IP, the public IP address.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The ID of the instance that the address is associated with (if any).
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> str:
        """
        The ID of the network interface.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="networkInterfaceOwnerId")
    def network_interface_owner_id(self) -> str:
        """
        The ID of the AWS account that owns the network interface.
        """
        return pulumi.get(self, "network_interface_owner_id")

    @property
    @pulumi.getter(name="privateDns")
    def private_dns(self) -> str:
        """
        The Private DNS associated with the Elastic IP address.
        """
        return pulumi.get(self, "private_dns")

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        The private IP address associated with the Elastic IP address.
        """
        return pulumi.get(self, "private_ip")

    @property
    @pulumi.getter(name="publicDns")
    def public_dns(self) -> str:
        """
        Public DNS associated with the Elastic IP address.
        """
        return pulumi.get(self, "public_dns")

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        Public IP address of Elastic IP.
        """
        return pulumi.get(self, "public_ip")

    @property
    @pulumi.getter(name="publicIpv4Pool")
    def public_ipv4_pool(self) -> str:
        """
        The ID of an address pool.
        """
        return pulumi.get(self, "public_ipv4_pool")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        Key-value map of tags associated with Elastic IP.
        """
        return pulumi.get(self, "tags")


class AwaitableGetElasticIpResult(GetElasticIpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetElasticIpResult(
            association_id=self.association_id,
            customer_owned_ip=self.customer_owned_ip,
            customer_owned_ipv4_pool=self.customer_owned_ipv4_pool,
            domain=self.domain,
            filters=self.filters,
            id=self.id,
            instance_id=self.instance_id,
            network_interface_id=self.network_interface_id,
            network_interface_owner_id=self.network_interface_owner_id,
            private_dns=self.private_dns,
            private_ip=self.private_ip,
            public_dns=self.public_dns,
            public_ip=self.public_ip,
            public_ipv4_pool=self.public_ipv4_pool,
            tags=self.tags)


def get_elastic_ip(filters: Optional[Sequence[pulumi.InputType['GetElasticIpFilterArgs']]] = None,
                   id: Optional[str] = None,
                   public_ip: Optional[str] = None,
                   tags: Optional[Mapping[str, str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetElasticIpResult:
    """
    `ec2.Eip` provides details about a specific Elastic IP.

    ## Example Usage
    ### Search By Allocation ID (VPC only)

    ```python
    import pulumi
    import pulumi_aws as aws

    by_allocation_id = aws.get_elastic_ip(id="eipalloc-12345678")
    ```
    ### Search By Filters (EC2-Classic or VPC)

    ```python
    import pulumi
    import pulumi_aws as aws

    by_filter = aws.get_elastic_ip(filters=[aws.GetElasticIpFilterArgs(
        name="tag:Name",
        values=["exampleNameTagValue"],
    )])
    ```
    ### Search By Public IP (EC2-Classic or VPC)

    ```python
    import pulumi
    import pulumi_aws as aws

    by_public_ip = aws.get_elastic_ip(public_ip="1.2.3.4")
    ```
    ### Search By Tags (EC2-Classic or VPC)

    ```python
    import pulumi
    import pulumi_aws as aws

    by_tags = aws.get_elastic_ip(tags={
        "Name": "exampleNameTagValue",
    })
    ```


    :param Sequence[pulumi.InputType['GetElasticIpFilterArgs']] filters: One or more name/value pairs to use as filters. There are several valid keys, for a full reference, check out the [EC2 API Reference](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAddresses.html).
    :param str id: The allocation id of the specific VPC EIP to retrieve. If a classic EIP is required, do NOT set `id`, only set `public_ip`
    :param str public_ip: The public IP of the specific EIP to retrieve.
    :param Mapping[str, str] tags: A map of tags, each pair of which must exactly match a pair on the desired Elastic IP
    """
    __args__ = dict()
    __args__['filters'] = filters
    __args__['id'] = id
    __args__['publicIp'] = public_ip
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getElasticIp:getElasticIp', __args__, opts=opts, typ=GetElasticIpResult).value

    return AwaitableGetElasticIpResult(
        association_id=__ret__.association_id,
        customer_owned_ip=__ret__.customer_owned_ip,
        customer_owned_ipv4_pool=__ret__.customer_owned_ipv4_pool,
        domain=__ret__.domain,
        filters=__ret__.filters,
        id=__ret__.id,
        instance_id=__ret__.instance_id,
        network_interface_id=__ret__.network_interface_id,
        network_interface_owner_id=__ret__.network_interface_owner_id,
        private_dns=__ret__.private_dns,
        private_ip=__ret__.private_ip,
        public_dns=__ret__.public_dns,
        public_ip=__ret__.public_ip,
        public_ipv4_pool=__ret__.public_ipv4_pool,
        tags=__ret__.tags)
