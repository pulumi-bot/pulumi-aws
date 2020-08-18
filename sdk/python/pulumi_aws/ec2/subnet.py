# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Subnet']


class Subnet(pulumi.CustomResource):
    arn: pulumi.Output[str] = pulumi.property("arn")
    """
    The ARN of the subnet.
    """

    assign_ipv6_address_on_creation: pulumi.Output[Optional[bool]] = pulumi.property("assignIpv6AddressOnCreation")
    """
    Specify true to indicate
    that network interfaces created in the specified subnet should be
    assigned an IPv6 address. Default is `false`
    """

    availability_zone: pulumi.Output[str] = pulumi.property("availabilityZone")
    """
    The AZ for the subnet.
    """

    availability_zone_id: pulumi.Output[str] = pulumi.property("availabilityZoneId")
    """
    The AZ ID of the subnet.
    """

    cidr_block: pulumi.Output[str] = pulumi.property("cidrBlock")
    """
    The CIDR block for the subnet.
    """

    ipv6_cidr_block: pulumi.Output[str] = pulumi.property("ipv6CidrBlock")
    """
    The IPv6 network range for the subnet,
    in CIDR notation. The subnet size must use a /64 prefix length.
    """

    ipv6_cidr_block_association_id: pulumi.Output[str] = pulumi.property("ipv6CidrBlockAssociationId")
    """
    The association ID for the IPv6 CIDR block.
    """

    map_public_ip_on_launch: pulumi.Output[Optional[bool]] = pulumi.property("mapPublicIpOnLaunch")
    """
    Specify true to indicate
    that instances launched into the subnet should be assigned
    a public IP address. Default is `false`.
    """

    outpost_arn: pulumi.Output[Optional[str]] = pulumi.property("outpostArn")
    """
    The Amazon Resource Name (ARN) of the Outpost.
    """

    owner_id: pulumi.Output[str] = pulumi.property("ownerId")
    """
    The ID of the AWS account that owns the subnet.
    """

    tags: pulumi.Output[Optional[Mapping[str, str]]] = pulumi.property("tags")
    """
    A map of tags to assign to the resource.
    """

    vpc_id: pulumi.Output[str] = pulumi.property("vpcId")
    """
    The VPC ID.
    """

    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 assign_ipv6_address_on_creation: Optional[pulumi.Input[bool]] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 availability_zone_id: Optional[pulumi.Input[str]] = None,
                 cidr_block: Optional[pulumi.Input[str]] = None,
                 ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
                 map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
                 outpost_arn: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides an VPC subnet resource.

        > **NOTE:** Due to [AWS Lambda improved VPC networking changes that began deploying in September 2019](https://aws.amazon.com/blogs/compute/announcing-improved-vpc-networking-for-aws-lambda-functions/), subnets associated with Lambda Functions can take up to 45 minutes to successfully delete.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        main = aws.ec2.Subnet("main",
            vpc_id=aws_vpc["main"]["id"],
            cidr_block="10.0.1.0/24",
            tags={
                "Name": "Main",
            })
        ```
        ### Subnets In Secondary VPC CIDR Blocks

        When managing subnets in one of a VPC's secondary CIDR blocks created using a `ec2.VpcIpv4CidrBlockAssociation`
        resource, it is recommended to reference that resource's `vpc_id` attribute to ensure correct dependency ordering.

        ```python
        import pulumi
        import pulumi_aws as aws

        secondary_cidr = aws.ec2.VpcIpv4CidrBlockAssociation("secondaryCidr",
            vpc_id=aws_vpc["main"]["id"],
            cidr_block="172.2.0.0/16")
        in_secondary_cidr = aws.ec2.Subnet("inSecondaryCidr",
            vpc_id=secondary_cidr.vpc_id,
            cidr_block="172.2.0.0/24")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] assign_ipv6_address_on_creation: Specify true to indicate
               that network interfaces created in the specified subnet should be
               assigned an IPv6 address. Default is `false`
        :param pulumi.Input[str] availability_zone: The AZ for the subnet.
        :param pulumi.Input[str] availability_zone_id: The AZ ID of the subnet.
        :param pulumi.Input[str] cidr_block: The CIDR block for the subnet.
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 network range for the subnet,
               in CIDR notation. The subnet size must use a /64 prefix length.
        :param pulumi.Input[bool] map_public_ip_on_launch: Specify true to indicate
               that instances launched into the subnet should be assigned
               a public IP address. Default is `false`.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the Outpost.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
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

            __props__['assign_ipv6_address_on_creation'] = assign_ipv6_address_on_creation
            __props__['availability_zone'] = availability_zone
            __props__['availability_zone_id'] = availability_zone_id
            if cidr_block is None:
                raise TypeError("Missing required property 'cidr_block'")
            __props__['cidr_block'] = cidr_block
            __props__['ipv6_cidr_block'] = ipv6_cidr_block
            __props__['map_public_ip_on_launch'] = map_public_ip_on_launch
            __props__['outpost_arn'] = outpost_arn
            __props__['tags'] = tags
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['arn'] = None
            __props__['ipv6_cidr_block_association_id'] = None
            __props__['owner_id'] = None
        super(Subnet, __self__).__init__(
            'aws:ec2/subnet:Subnet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            assign_ipv6_address_on_creation: Optional[pulumi.Input[bool]] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            availability_zone_id: Optional[pulumi.Input[str]] = None,
            cidr_block: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block: Optional[pulumi.Input[str]] = None,
            ipv6_cidr_block_association_id: Optional[pulumi.Input[str]] = None,
            map_public_ip_on_launch: Optional[pulumi.Input[bool]] = None,
            outpost_arn: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Subnet':
        """
        Get an existing Subnet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the subnet.
        :param pulumi.Input[bool] assign_ipv6_address_on_creation: Specify true to indicate
               that network interfaces created in the specified subnet should be
               assigned an IPv6 address. Default is `false`
        :param pulumi.Input[str] availability_zone: The AZ for the subnet.
        :param pulumi.Input[str] availability_zone_id: The AZ ID of the subnet.
        :param pulumi.Input[str] cidr_block: The CIDR block for the subnet.
        :param pulumi.Input[str] ipv6_cidr_block: The IPv6 network range for the subnet,
               in CIDR notation. The subnet size must use a /64 prefix length.
        :param pulumi.Input[str] ipv6_cidr_block_association_id: The association ID for the IPv6 CIDR block.
        :param pulumi.Input[bool] map_public_ip_on_launch: Specify true to indicate
               that instances launched into the subnet should be assigned
               a public IP address. Default is `false`.
        :param pulumi.Input[str] outpost_arn: The Amazon Resource Name (ARN) of the Outpost.
        :param pulumi.Input[str] owner_id: The ID of the AWS account that owns the subnet.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource.
        :param pulumi.Input[str] vpc_id: The VPC ID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["assign_ipv6_address_on_creation"] = assign_ipv6_address_on_creation
        __props__["availability_zone"] = availability_zone
        __props__["availability_zone_id"] = availability_zone_id
        __props__["cidr_block"] = cidr_block
        __props__["ipv6_cidr_block"] = ipv6_cidr_block
        __props__["ipv6_cidr_block_association_id"] = ipv6_cidr_block_association_id
        __props__["map_public_ip_on_launch"] = map_public_ip_on_launch
        __props__["outpost_arn"] = outpost_arn
        __props__["owner_id"] = owner_id
        __props__["tags"] = tags
        __props__["vpc_id"] = vpc_id
        return Subnet(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

