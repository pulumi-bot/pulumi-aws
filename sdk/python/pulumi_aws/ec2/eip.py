# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Eip']


class Eip(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associate_with_private_ip: Optional[pulumi.Input[str]] = None,
                 customer_owned_ipv4_pool: Optional[pulumi.Input[str]] = None,
                 instance: Optional[pulumi.Input[str]] = None,
                 network_interface: Optional[pulumi.Input[str]] = None,
                 public_ipv4_pool: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an Elastic IP resource.

        > **Note:** EIP may require IGW to exist prior to association. Use `depends_on` to set an explicit dependency on the IGW.

        > **Note:** Do not use `network_interface` to associate the EIP to `lb.LoadBalancer` or `ec2.NatGateway` resources. Instead use the `allocation_id` available in those resources to allow AWS to manage the association, otherwise you will see `AuthFailure` errors.

        ## Example Usage

        Single EIP associated with an instance:

        ```python
        import pulumi
        import pulumi_aws as aws

        lb = aws.ec2.Eip("lb",
            instance=aws_instance["web"]["id"],
            vpc=True)
        ```

        Multiple EIPs associated with a single network interface:

        ```python
        import pulumi
        import pulumi_aws as aws

        multi_ip = aws.ec2.NetworkInterface("multi-ip",
            subnet_id=aws_subnet["main"]["id"],
            private_ips=[
                "10.0.0.10",
                "10.0.0.11",
            ])
        one = aws.ec2.Eip("one",
            vpc=True,
            network_interface=multi_ip.id,
            associate_with_private_ip="10.0.0.10")
        two = aws.ec2.Eip("two",
            vpc=True,
            network_interface=multi_ip.id,
            associate_with_private_ip="10.0.0.11")
        ```

        Attaching an EIP to an Instance with a pre-assigned private ip (VPC Only):

        ```python
        import pulumi
        import pulumi_aws as aws

        default = aws.ec2.Vpc("default",
            cidr_block="10.0.0.0/16",
            enable_dns_hostnames=True)
        gw = aws.ec2.InternetGateway("gw", vpc_id=default.id)
        tf_test_subnet = aws.ec2.Subnet("tfTestSubnet",
            vpc_id=default.id,
            cidr_block="10.0.0.0/24",
            map_public_ip_on_launch=True,
            opts=ResourceOptions(depends_on=[gw]))
        foo = aws.ec2.Instance("foo",
            ami="ami-5189a661",
            instance_type="t2.micro",
            private_ip="10.0.0.12",
            subnet_id=tf_test_subnet.id)
        bar = aws.ec2.Eip("bar",
            vpc=True,
            instance=foo.id,
            associate_with_private_ip="10.0.0.12",
            opts=ResourceOptions(depends_on=[gw]))
        ```

        Allocating EIP from the BYOIP pool:

        ```python
        import pulumi
        import pulumi_aws as aws

        byoip_ip = aws.ec2.Eip("byoip-ip",
            public_ipv4_pool="ipv4pool-ec2-012345",
            vpc=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] associate_with_private_ip: A user specified primary or secondary private IP address to
               associate with the Elastic IP address. If no private IP address is specified,
               the Elastic IP address is associated with the primary private IP address.
        :param pulumi.Input[str] customer_owned_ipv4_pool: The  ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
        :param pulumi.Input[str] instance: EC2 instance ID.
        :param pulumi.Input[str] network_interface: Network interface ID to associate with.
        :param pulumi.Input[str] public_ipv4_pool: EC2 IPv4 address pool identifier or `amazon`. This option is only available for VPC EIPs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[bool] vpc: Boolean if the EIP is in a VPC or not.
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

            __props__['associate_with_private_ip'] = associate_with_private_ip
            __props__['customer_owned_ipv4_pool'] = customer_owned_ipv4_pool
            __props__['instance'] = instance
            __props__['network_interface'] = network_interface
            __props__['public_ipv4_pool'] = public_ipv4_pool
            __props__['tags'] = tags
            __props__['vpc'] = vpc
            __props__['allocation_id'] = None
            __props__['association_id'] = None
            __props__['customer_owned_ip'] = None
            __props__['domain'] = None
            __props__['private_dns'] = None
            __props__['private_ip'] = None
            __props__['public_dns'] = None
            __props__['public_ip'] = None
        super(Eip, __self__).__init__(
            'aws:ec2/eip:Eip',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_id: Optional[pulumi.Input[str]] = None,
            associate_with_private_ip: Optional[pulumi.Input[str]] = None,
            association_id: Optional[pulumi.Input[str]] = None,
            customer_owned_ip: Optional[pulumi.Input[str]] = None,
            customer_owned_ipv4_pool: Optional[pulumi.Input[str]] = None,
            domain: Optional[pulumi.Input[str]] = None,
            instance: Optional[pulumi.Input[str]] = None,
            network_interface: Optional[pulumi.Input[str]] = None,
            private_dns: Optional[pulumi.Input[str]] = None,
            private_ip: Optional[pulumi.Input[str]] = None,
            public_dns: Optional[pulumi.Input[str]] = None,
            public_ip: Optional[pulumi.Input[str]] = None,
            public_ipv4_pool: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc: Optional[pulumi.Input[bool]] = None) -> 'Eip':
        """
        Get an existing Eip resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] associate_with_private_ip: A user specified primary or secondary private IP address to
               associate with the Elastic IP address. If no private IP address is specified,
               the Elastic IP address is associated with the primary private IP address.
        :param pulumi.Input[str] customer_owned_ip: Customer owned IP.
        :param pulumi.Input[str] customer_owned_ipv4_pool: The  ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
        :param pulumi.Input[str] instance: EC2 instance ID.
        :param pulumi.Input[str] network_interface: Network interface ID to associate with.
        :param pulumi.Input[str] private_dns: The Private DNS associated with the Elastic IP address (if in VPC).
        :param pulumi.Input[str] private_ip: Contains the private IP address (if in VPC).
        :param pulumi.Input[str] public_dns: Public DNS associated with the Elastic IP address.
        :param pulumi.Input[str] public_ip: Contains the public IP address.
        :param pulumi.Input[str] public_ipv4_pool: EC2 IPv4 address pool identifier or `amazon`. This option is only available for VPC EIPs.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[bool] vpc: Boolean if the EIP is in a VPC or not.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allocation_id"] = allocation_id
        __props__["associate_with_private_ip"] = associate_with_private_ip
        __props__["association_id"] = association_id
        __props__["customer_owned_ip"] = customer_owned_ip
        __props__["customer_owned_ipv4_pool"] = customer_owned_ipv4_pool
        __props__["domain"] = domain
        __props__["instance"] = instance
        __props__["network_interface"] = network_interface
        __props__["private_dns"] = private_dns
        __props__["private_ip"] = private_ip
        __props__["public_dns"] = public_dns
        __props__["public_ip"] = public_ip
        __props__["public_ipv4_pool"] = public_ipv4_pool
        __props__["tags"] = tags
        __props__["vpc"] = vpc
        return Eip(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationId")
    def allocation_id(self) -> str:
        ...

    @property
    @pulumi.getter(name="associateWithPrivateIp")
    def associate_with_private_ip(self) -> Optional[str]:
        """
        A user specified primary or secondary private IP address to
        associate with the Elastic IP address. If no private IP address is specified,
        the Elastic IP address is associated with the primary private IP address.
        """
        ...

    @property
    @pulumi.getter(name="associationId")
    def association_id(self) -> str:
        ...

    @property
    @pulumi.getter(name="customerOwnedIp")
    def customer_owned_ip(self) -> str:
        """
        Customer owned IP.
        """
        ...

    @property
    @pulumi.getter(name="customerOwnedIpv4Pool")
    def customer_owned_ipv4_pool(self) -> Optional[str]:
        """
        The  ID  of a customer-owned address pool. For more on customer owned IP addressed check out [Customer-owned IP addresses guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#ip-addressing)
        """
        ...

    @property
    @pulumi.getter
    def domain(self) -> str:
        ...

    @property
    @pulumi.getter
    def instance(self) -> str:
        """
        EC2 instance ID.
        """
        ...

    @property
    @pulumi.getter(name="networkInterface")
    def network_interface(self) -> str:
        """
        Network interface ID to associate with.
        """
        ...

    @property
    @pulumi.getter(name="privateDns")
    def private_dns(self) -> str:
        """
        The Private DNS associated with the Elastic IP address (if in VPC).
        """
        ...

    @property
    @pulumi.getter(name="privateIp")
    def private_ip(self) -> str:
        """
        Contains the private IP address (if in VPC).
        """
        ...

    @property
    @pulumi.getter(name="publicDns")
    def public_dns(self) -> str:
        """
        Public DNS associated with the Elastic IP address.
        """
        ...

    @property
    @pulumi.getter(name="publicIp")
    def public_ip(self) -> str:
        """
        Contains the public IP address.
        """
        ...

    @property
    @pulumi.getter(name="publicIpv4Pool")
    def public_ipv4_pool(self) -> str:
        """
        EC2 IPv4 address pool identifier or `amazon`. This option is only available for VPC EIPs.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        A map of tags to assign to the resource.
        """
        ...

    @property
    @pulumi.getter
    def vpc(self) -> bool:
        """
        Boolean if the EIP is in a VPC or not.
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

