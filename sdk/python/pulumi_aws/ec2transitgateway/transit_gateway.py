# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['TransitGatewayArgs', 'TransitGateway']

@pulumi.input_type
class TransitGatewayArgs:
    def __init__(__self__, *,
                 amazon_side_asn: Optional[pulumi.Input[int]] = None,
                 auto_accept_shared_attachments: Optional[pulumi.Input[str]] = None,
                 default_route_table_association: Optional[pulumi.Input[str]] = None,
                 default_route_table_propagation: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_support: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_ecmp_support: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a TransitGateway resource.
        :param pulumi.Input[int] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        :param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway.
        :param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        if amazon_side_asn is not None:
            pulumi.set(__self__, "amazon_side_asn", amazon_side_asn)
        if auto_accept_shared_attachments is not None:
            pulumi.set(__self__, "auto_accept_shared_attachments", auto_accept_shared_attachments)
        if default_route_table_association is not None:
            pulumi.set(__self__, "default_route_table_association", default_route_table_association)
        if default_route_table_propagation is not None:
            pulumi.set(__self__, "default_route_table_propagation", default_route_table_propagation)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns_support is not None:
            pulumi.set(__self__, "dns_support", dns_support)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpn_ecmp_support is not None:
            pulumi.set(__self__, "vpn_ecmp_support", vpn_ecmp_support)

    @property
    @pulumi.getter(name="amazonSideAsn")
    def amazon_side_asn(self) -> Optional[pulumi.Input[int]]:
        """
        Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        """
        return pulumi.get(self, "amazon_side_asn")

    @amazon_side_asn.setter
    def amazon_side_asn(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "amazon_side_asn", value)

    @property
    @pulumi.getter(name="autoAcceptSharedAttachments")
    def auto_accept_shared_attachments(self) -> Optional[pulumi.Input[str]]:
        """
        Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        """
        return pulumi.get(self, "auto_accept_shared_attachments")

    @auto_accept_shared_attachments.setter
    def auto_accept_shared_attachments(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_accept_shared_attachments", value)

    @property
    @pulumi.getter(name="defaultRouteTableAssociation")
    def default_route_table_association(self) -> Optional[pulumi.Input[str]]:
        """
        Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "default_route_table_association")

    @default_route_table_association.setter
    def default_route_table_association(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_route_table_association", value)

    @property
    @pulumi.getter(name="defaultRouteTablePropagation")
    def default_route_table_propagation(self) -> Optional[pulumi.Input[str]]:
        """
        Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "default_route_table_propagation")

    @default_route_table_propagation.setter
    def default_route_table_propagation(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_route_table_propagation", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the EC2 Transit Gateway.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dnsSupport")
    def dns_support(self) -> Optional[pulumi.Input[str]]:
        """
        Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "dns_support")

    @dns_support.setter
    def dns_support(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_support", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Key-value tags for the EC2 Transit Gateway.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpnEcmpSupport")
    def vpn_ecmp_support(self) -> Optional[pulumi.Input[str]]:
        """
        Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "vpn_ecmp_support")

    @vpn_ecmp_support.setter
    def vpn_ecmp_support(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_ecmp_support", value)


class TransitGateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TransitGatewayArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EC2 Transit Gateway.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2transitgateway.TransitGateway("example", description="example")
        ```

        ## Import

        `aws_ec2_transit_gateway` can be imported by using the EC2 Transit Gateway identifier, e.g.

        ```sh
         $ pulumi import aws:ec2transitgateway/transitGateway:TransitGateway example tgw-12345678
        ```

        :param str resource_name: The name of the resource.
        :param TransitGatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amazon_side_asn: Optional[pulumi.Input[int]] = None,
                 auto_accept_shared_attachments: Optional[pulumi.Input[str]] = None,
                 default_route_table_association: Optional[pulumi.Input[str]] = None,
                 default_route_table_propagation: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_support: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_ecmp_support: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages an EC2 Transit Gateway.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.ec2transitgateway.TransitGateway("example", description="example")
        ```

        ## Import

        `aws_ec2_transit_gateway` can be imported by using the EC2 Transit Gateway identifier, e.g.

        ```sh
         $ pulumi import aws:ec2transitgateway/transitGateway:TransitGateway example tgw-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        :param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway.
        :param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransitGatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 amazon_side_asn: Optional[pulumi.Input[int]] = None,
                 auto_accept_shared_attachments: Optional[pulumi.Input[str]] = None,
                 default_route_table_association: Optional[pulumi.Input[str]] = None,
                 default_route_table_propagation: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns_support: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vpn_ecmp_support: Optional[pulumi.Input[str]] = None,
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

            __props__['amazon_side_asn'] = amazon_side_asn
            __props__['auto_accept_shared_attachments'] = auto_accept_shared_attachments
            __props__['default_route_table_association'] = default_route_table_association
            __props__['default_route_table_propagation'] = default_route_table_propagation
            __props__['description'] = description
            __props__['dns_support'] = dns_support
            __props__['tags'] = tags
            __props__['vpn_ecmp_support'] = vpn_ecmp_support
            __props__['arn'] = None
            __props__['association_default_route_table_id'] = None
            __props__['owner_id'] = None
            __props__['propagation_default_route_table_id'] = None
        super(TransitGateway, __self__).__init__(
            'aws:ec2transitgateway/transitGateway:TransitGateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            amazon_side_asn: Optional[pulumi.Input[int]] = None,
            arn: Optional[pulumi.Input[str]] = None,
            association_default_route_table_id: Optional[pulumi.Input[str]] = None,
            auto_accept_shared_attachments: Optional[pulumi.Input[str]] = None,
            default_route_table_association: Optional[pulumi.Input[str]] = None,
            default_route_table_propagation: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns_support: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            propagation_default_route_table_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            vpn_ecmp_support: Optional[pulumi.Input[str]] = None) -> 'TransitGateway':
        """
        Get an existing TransitGateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] amazon_side_asn: Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        :param pulumi.Input[str] arn: EC2 Transit Gateway Amazon Resource Name (ARN)
        :param pulumi.Input[str] association_default_route_table_id: Identifier of the default association route table
        :param pulumi.Input[str] auto_accept_shared_attachments: Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        :param pulumi.Input[str] default_route_table_association: Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] default_route_table_propagation: Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] description: Description of the EC2 Transit Gateway.
        :param pulumi.Input[str] dns_support: Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        :param pulumi.Input[str] owner_id: Identifier of the AWS account that owns the EC2 Transit Gateway
        :param pulumi.Input[str] propagation_default_route_table_id: Identifier of the default propagation route table
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value tags for the EC2 Transit Gateway.
        :param pulumi.Input[str] vpn_ecmp_support: Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["amazon_side_asn"] = amazon_side_asn
        __props__["arn"] = arn
        __props__["association_default_route_table_id"] = association_default_route_table_id
        __props__["auto_accept_shared_attachments"] = auto_accept_shared_attachments
        __props__["default_route_table_association"] = default_route_table_association
        __props__["default_route_table_propagation"] = default_route_table_propagation
        __props__["description"] = description
        __props__["dns_support"] = dns_support
        __props__["owner_id"] = owner_id
        __props__["propagation_default_route_table_id"] = propagation_default_route_table_id
        __props__["tags"] = tags
        __props__["vpn_ecmp_support"] = vpn_ecmp_support
        return TransitGateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="amazonSideAsn")
    def amazon_side_asn(self) -> pulumi.Output[Optional[int]]:
        """
        Private Autonomous System Number (ASN) for the Amazon side of a BGP session. The range is `64512` to `65534` for 16-bit ASNs and `4200000000` to `4294967294` for 32-bit ASNs. Default value: `64512`.
        """
        return pulumi.get(self, "amazon_side_asn")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        EC2 Transit Gateway Amazon Resource Name (ARN)
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="associationDefaultRouteTableId")
    def association_default_route_table_id(self) -> pulumi.Output[str]:
        """
        Identifier of the default association route table
        """
        return pulumi.get(self, "association_default_route_table_id")

    @property
    @pulumi.getter(name="autoAcceptSharedAttachments")
    def auto_accept_shared_attachments(self) -> pulumi.Output[Optional[str]]:
        """
        Whether resource attachment requests are automatically accepted. Valid values: `disable`, `enable`. Default value: `disable`.
        """
        return pulumi.get(self, "auto_accept_shared_attachments")

    @property
    @pulumi.getter(name="defaultRouteTableAssociation")
    def default_route_table_association(self) -> pulumi.Output[Optional[str]]:
        """
        Whether resource attachments are automatically associated with the default association route table. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "default_route_table_association")

    @property
    @pulumi.getter(name="defaultRouteTablePropagation")
    def default_route_table_propagation(self) -> pulumi.Output[Optional[str]]:
        """
        Whether resource attachments automatically propagate routes to the default propagation route table. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "default_route_table_propagation")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the EC2 Transit Gateway.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dnsSupport")
    def dns_support(self) -> pulumi.Output[Optional[str]]:
        """
        Whether DNS support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "dns_support")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        Identifier of the AWS account that owns the EC2 Transit Gateway
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="propagationDefaultRouteTableId")
    def propagation_default_route_table_id(self) -> pulumi.Output[str]:
        """
        Identifier of the default propagation route table
        """
        return pulumi.get(self, "propagation_default_route_table_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Key-value tags for the EC2 Transit Gateway.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpnEcmpSupport")
    def vpn_ecmp_support(self) -> pulumi.Output[Optional[str]]:
        """
        Whether VPN Equal Cost Multipath Protocol support is enabled. Valid values: `disable`, `enable`. Default value: `enable`.
        """
        return pulumi.get(self, "vpn_ecmp_support")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

