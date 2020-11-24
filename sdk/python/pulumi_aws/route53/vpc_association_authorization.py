# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['VpcAssociationAuthorization']


class VpcAssociationAuthorization(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_region: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Authorizes a VPC in a peer account to be associated with a local Route53 Hosted Zone.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        alternate = pulumi.providers.Aws("alternate")
        example_vpc = aws.ec2.Vpc("exampleVpc",
            cidr_block="10.6.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True)
        example_zone = aws.route53.Zone("exampleZone", vpcs=[aws.route53.ZoneVpcArgs(
            vpc_id=example_vpc.id,
        )])
        alternate_vpc = aws.ec2.Vpc("alternateVpc",
            cidr_block="10.7.0.0/16",
            enable_dns_hostnames=True,
            enable_dns_support=True,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        example_vpc_association_authorization = aws.route53.VpcAssociationAuthorization("exampleVpcAssociationAuthorization",
            vpc_id=alternate_vpc.id,
            zone_id=example_zone.id)
        example_zone_association = aws.route53.ZoneAssociation("exampleZoneAssociation",
            vpc_id=example_vpc_association_authorization.vpc_id,
            zone_id=example_vpc_association_authorization.zone_id,
            opts=pulumi.ResourceOptions(provider="aws.alternate"))
        ```

        ## Import

        Route 53 VPC Association Authorizations can be imported via the Hosted Zone ID and VPC ID, separated by a colon (`:`), e.g.

        ```sh
         $ pulumi import aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization example Z123456ABCDEFG:vpc-12345678
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
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

            if vpc_id is None:
                raise TypeError("Missing required property 'vpc_id'")
            __props__['vpc_id'] = vpc_id
            __props__['vpc_region'] = vpc_region
            if zone_id is None:
                raise TypeError("Missing required property 'zone_id'")
            __props__['zone_id'] = zone_id
        super(VpcAssociationAuthorization, __self__).__init__(
            'aws:route53/vpcAssociationAuthorization:VpcAssociationAuthorization',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_region: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None) -> 'VpcAssociationAuthorization':
        """
        Get an existing VpcAssociationAuthorization resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] vpc_id: The VPC to authorize for association with the private hosted zone.
        :param pulumi.Input[str] vpc_region: The VPC's region. Defaults to the region of the AWS provider.
        :param pulumi.Input[str] zone_id: The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["vpc_id"] = vpc_id
        __props__["vpc_region"] = vpc_region
        __props__["zone_id"] = zone_id
        return VpcAssociationAuthorization(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The VPC to authorize for association with the private hosted zone.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcRegion")
    def vpc_region(self) -> pulumi.Output[str]:
        """
        The VPC's region. Defaults to the region of the AWS provider.
        """
        return pulumi.get(self, "vpc_region")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The ID of the private hosted zone that you want to authorize associating a VPC with.
        """
        return pulumi.get(self, "zone_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

