# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['VpcEndpointSubnetAssociationArgs', 'VpcEndpointSubnetAssociation']

@pulumi.input_type
class VpcEndpointSubnetAssociationArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 vpc_endpoint_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a VpcEndpointSubnetAssociation resource.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to be associated with the VPC endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC endpoint with which the subnet will be associated.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_endpoint_id", vpc_endpoint_id)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the subnet to be associated with the VPC endpoint.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC endpoint with which the subnet will be associated.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    @vpc_endpoint_id.setter
    def vpc_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_endpoint_id", value)


class VpcEndpointSubnetAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create an association between a VPC endpoint and a subnet.

        > **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
        both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
        and a single `subnet_id`) and a VPC Endpoint resource with a `subnet_ids`
        attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
        Association resource. Doing so will cause a conflict of associations and will overwrite the association.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        sn_ec2 = aws.ec2.VpcEndpointSubnetAssociation("snEc2",
            vpc_endpoint_id=aws_vpc_endpoint["ec2"]["id"],
            subnet_id=aws_subnet["sn"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to be associated with the VPC endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC endpoint with which the subnet will be associated.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointSubnetAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create an association between a VPC endpoint and a subnet.

        > **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
        both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
        and a single `subnet_id`) and a VPC Endpoint resource with a `subnet_ids`
        attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
        Association resource. Doing so will cause a conflict of associations and will overwrite the association.

        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        sn_ec2 = aws.ec2.VpcEndpointSubnetAssociation("snEc2",
            vpc_endpoint_id=aws_vpc_endpoint["ec2"]["id"],
            subnet_id=aws_subnet["sn"]["id"])
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointSubnetAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointSubnetAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 vpc_endpoint_id: Optional[pulumi.Input[str]] = None,
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

            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__['subnet_id'] = subnet_id
            if vpc_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_endpoint_id'")
            __props__['vpc_endpoint_id'] = vpc_endpoint_id
        super(VpcEndpointSubnetAssociation, __self__).__init__(
            'aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            vpc_endpoint_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointSubnetAssociation':
        """
        Get an existing VpcEndpointSubnetAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] subnet_id: The ID of the subnet to be associated with the VPC endpoint.
        :param pulumi.Input[str] vpc_endpoint_id: The ID of the VPC endpoint with which the subnet will be associated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["subnet_id"] = subnet_id
        __props__["vpc_endpoint_id"] = vpc_endpoint_id
        return VpcEndpointSubnetAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the subnet to be associated with the VPC endpoint.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="vpcEndpointId")
    def vpc_endpoint_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC endpoint with which the subnet will be associated.
        """
        return pulumi.get(self, "vpc_endpoint_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

