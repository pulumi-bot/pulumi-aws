# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['VpcAttachment']


class VpcAttachment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dns_support: Optional[pulumi.Input[str]] = None,
                 ipv6_support: Optional[pulumi.Input[str]] = None,
                 subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_default_route_table_association: Optional[pulumi.Input[bool]] = None,
                 transit_gateway_default_route_table_propagation: Optional[pulumi.Input[bool]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an EC2 Transit Gateway VPC Attachment. For examples of custom route table association and propagation, see the EC2 Transit Gateway Networking Examples Guide.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2transitgateway.VpcAttachment("example",
            subnet_ids=[aws_subnet["example"]["id"]],
            transit_gateway_id=aws_ec2_transit_gateway["example"]["id"],
            vpc_id=aws_vpc["example"]["id"])
        ```

        A full example of how to create a Transit Gateway in one AWS account, share it with a second AWS account, and attach a VPC in the second account to the Transit Gateway via the `ec2transitgateway.VpcAttachment` and `ec2transitgateway.VpcAttachmentAccepter` resources can be found in [the `./examples/transit-gateway-cross-account-vpc-attachment` directory within the Github Repository](https://github.com/providers/provider-aws/tree/master/examples/transit-gateway-cross-account-vpc-attachment).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] ipv6_support: Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: Identifiers of EC2 Subnets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway VPC Attachment.
        :param pulumi.Input[bool] transit_gateway_default_route_table_association: Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        :param pulumi.Input[bool] transit_gateway_default_route_table_propagation: Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_id: Identifier of EC2 VPC.
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

            __props__['dns_support'] = dns_support
            __props__['ipv6_support'] = ipv6_support
            if subnet_ids is None:
                raise TypeError("Missing required property 'subnet_ids'")
            __props__['subnet_ids'] = subnet_ids
            __props__['tags'] = tags
            __props__['transit_gateway_default_route_table_association'] = transit_gateway_default_route_table_association
            __props__['transit_gateway_default_route_table_propagation'] = transit_gateway_default_route_table_propagation
            if transit_gateway_id is None:
                raise TypeError("Missing required property 'transit_gateway_id'")
            __props__['transit_gateway_id'] = transit_gateway_id
            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['vpc_owner_id'] = None
        super(VpcAttachment, __self__).__init__(
            'aws:ec2transitgateway/vpcAttachment:VpcAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            dns_support: Optional[pulumi.Input[str]] = None,
            ipv6_support: Optional[pulumi.Input[str]] = None,
            subnet_ids: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transit_gateway_default_route_table_association: Optional[pulumi.Input[bool]] = None,
            transit_gateway_default_route_table_propagation: Optional[pulumi.Input[bool]] = None,
            transit_gateway_id: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_owner_id: Optional[pulumi.Input[str]] = None) -> 'VpcAttachment':
        """
        Get an existing VpcAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] ipv6_support: Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[List[pulumi.Input[str]]] subnet_ids: Identifiers of EC2 Subnets.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway VPC Attachment.
        :param pulumi.Input[bool] transit_gateway_default_route_table_association: Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        :param pulumi.Input[bool] transit_gateway_default_route_table_propagation: Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        :param pulumi.Input[str] vpc_id: Identifier of EC2 VPC.
        :param pulumi.Input[str] vpc_owner_id: Identifier of the AWS account that owns the EC2 VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["dns_support"] = dns_support
        __props__["ipv6_support"] = ipv6_support
        __props__["subnet_ids"] = subnet_ids
        __props__["tags"] = tags
        __props__["transit_gateway_default_route_table_association"] = transit_gateway_default_route_table_association
        __props__["transit_gateway_default_route_table_propagation"] = transit_gateway_default_route_table_propagation
        __props__["transit_gateway_id"] = transit_gateway_id
        __props__["vpc_id"] = vpc_id
        __props__["vpc_owner_id"] = vpc_owner_id
        return VpcAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dnsSupport")
    def dns_support(self) -> Optional[str]:
        """
        Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "dns_support")

    @property
    @pulumi.getter(name="ipv6Support")
    def ipv6_support(self) -> Optional[str]:
        """
        Whether IPv6 support is enabled. Valid values: `disable`, `enable`. Default value: `disable`.
        """
        return pulumi.get(self, "ipv6_support")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> List[str]:
        """
        Identifiers of EC2 Subnets.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value tags for the EC2 Transit Gateway VPC Attachment.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="transitGatewayDefaultRouteTableAssociation")
    def transit_gateway_default_route_table_association(self) -> Optional[bool]:
        """
        Boolean whether the VPC Attachment should be associated with the EC2 Transit Gateway association default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        """
        return pulumi.get(self, "transit_gateway_default_route_table_association")

    @property
    @pulumi.getter(name="transitGatewayDefaultRouteTablePropagation")
    def transit_gateway_default_route_table_propagation(self) -> Optional[bool]:
        """
        Boolean whether the VPC Attachment should propagate routes with the EC2 Transit Gateway propagation default route table. This cannot be configured or perform drift detection with Resource Access Manager shared EC2 Transit Gateways. Default value: `true`.
        """
        return pulumi.get(self, "transit_gateway_default_route_table_propagation")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> str:
        """
        Identifier of EC2 Transit Gateway.
        """
        return pulumi.get(self, "transit_gateway_id")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> str:
        """
        Identifier of EC2 VPC.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcOwnerId")
    def vpc_owner_id(self) -> str:
        """
        Identifier of the AWS account that owns the EC2 VPC.
        """
        return pulumi.get(self, "vpc_owner_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

