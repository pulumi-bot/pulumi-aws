# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GatewayAssociation(pulumi.CustomResource):
    allowed_prefixes: pulumi.Output[list]
    """
    VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
    """
    associated_gateway_id: pulumi.Output[str]
    """
    The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
    Used for single account Direct Connect gateway associations.
    """
    associated_gateway_owner_account_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
    Used for cross-account Direct Connect gateway associations.
    """
    associated_gateway_type: pulumi.Output[str]
    """
    The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
    """
    dx_gateway_association_id: pulumi.Output[str]
    """
    The ID of the Direct Connect gateway association.
    """
    dx_gateway_id: pulumi.Output[str]
    """
    The ID of the Direct Connect gateway.
    """
    dx_gateway_owner_account_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the Direct Connect gateway.
    """
    proposal_id: pulumi.Output[str]
    """
    The ID of the Direct Connect gateway association proposal.
    Used for cross-account Direct Connect gateway associations.
    """
    vpn_gateway_id: pulumi.Output[str]
    """
    *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
    Used for single account Direct Connect gateway associations.
    """
    def __init__(__self__, resource_name, opts=None, allowed_prefixes=None, associated_gateway_id=None, associated_gateway_owner_account_id=None, dx_gateway_id=None, proposal_id=None, vpn_gateway_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Associates a Direct Connect Gateway with a VGW or transit gateway.

        To create a cross-account association, create an [`directconnect.GatewayAssociationProposal` resource](https://www.terraform.io/docs/providers/aws/r/dx_gateway_association_proposal.html)
        in the AWS account that owns the VGW or transit gateway and then accept the proposal in the AWS account that owns the Direct Connect Gateway
        by creating an `directconnect.GatewayAssociation` resource with the `proposal_id` and `associated_gateway_owner_account_id` attributes set.

        ## Example Usage

        ### VPN Gateway Association

        ```python
        import pulumi
        import pulumi_aws as aws

        example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.255.255.0/28")
        example_vpn_gateway = aws.ec2.VpnGateway("exampleVpnGateway", vpc_id=example_vpc.id)
        example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
            associated_gateway_id=example_vpn_gateway.id,
            dx_gateway_id=example_gateway.id)
        ```

        ### Transit Gateway Association

        ```python
        import pulumi
        import pulumi_aws as aws

        example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
        example_transit_gateway = aws.ec2transitgateway.TransitGateway("exampleTransitGateway")
        example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
            allowed_prefixes=[
                "10.255.255.0/30",
                "10.255.255.8/30",
            ],
            associated_gateway_id=example_transit_gateway.id,
            dx_gateway_id=example_gateway.id)
        ```

        ### Allowed Prefixes

        ```python
        import pulumi
        import pulumi_aws as aws

        example_gateway = aws.directconnect.Gateway("exampleGateway", amazon_side_asn="64512")
        example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.255.255.0/28")
        example_vpn_gateway = aws.ec2.VpnGateway("exampleVpnGateway", vpc_id=example_vpc.id)
        example_gateway_association = aws.directconnect.GatewayAssociation("exampleGatewayAssociation",
            allowed_prefixes=[
                "210.52.109.0/24",
                "175.45.176.0/22",
            ],
            associated_gateway_id=example_vpn_gateway.id,
            dx_gateway_id=example_gateway.id)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_prefixes: VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        :param pulumi.Input[str] associated_gateway_id: The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
               Used for single account Direct Connect gateway associations.
        :param pulumi.Input[str] associated_gateway_owner_account_id: The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
               Used for cross-account Direct Connect gateway associations.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway.
        :param pulumi.Input[str] proposal_id: The ID of the Direct Connect gateway association proposal.
               Used for cross-account Direct Connect gateway associations.
        :param pulumi.Input[str] vpn_gateway_id: *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
               Used for single account Direct Connect gateway associations.
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

            __props__['allowed_prefixes'] = allowed_prefixes
            __props__['associated_gateway_id'] = associated_gateway_id
            __props__['associated_gateway_owner_account_id'] = associated_gateway_owner_account_id
            if dx_gateway_id is None:
                raise TypeError("Missing required property 'dx_gateway_id'")
            __props__['dx_gateway_id'] = dx_gateway_id
            __props__['proposal_id'] = proposal_id
            if vpn_gateway_id is not None:
                warnings.warn("use 'associated_gateway_id' argument instead", DeprecationWarning)
                pulumi.log.warn("vpn_gateway_id is deprecated: use 'associated_gateway_id' argument instead")
            __props__['vpn_gateway_id'] = vpn_gateway_id
            __props__['associated_gateway_type'] = None
            __props__['dx_gateway_association_id'] = None
            __props__['dx_gateway_owner_account_id'] = None
        super(GatewayAssociation, __self__).__init__(
            'aws:directconnect/gatewayAssociation:GatewayAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_prefixes=None, associated_gateway_id=None, associated_gateway_owner_account_id=None, associated_gateway_type=None, dx_gateway_association_id=None, dx_gateway_id=None, dx_gateway_owner_account_id=None, proposal_id=None, vpn_gateway_id=None):
        """
        Get an existing GatewayAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_prefixes: VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        :param pulumi.Input[str] associated_gateway_id: The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
               Used for single account Direct Connect gateway associations.
        :param pulumi.Input[str] associated_gateway_owner_account_id: The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
               Used for cross-account Direct Connect gateway associations.
        :param pulumi.Input[str] associated_gateway_type: The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
        :param pulumi.Input[str] dx_gateway_association_id: The ID of the Direct Connect gateway association.
        :param pulumi.Input[str] dx_gateway_id: The ID of the Direct Connect gateway.
        :param pulumi.Input[str] dx_gateway_owner_account_id: The ID of the AWS account that owns the Direct Connect gateway.
        :param pulumi.Input[str] proposal_id: The ID of the Direct Connect gateway association proposal.
               Used for cross-account Direct Connect gateway associations.
        :param pulumi.Input[str] vpn_gateway_id: *Deprecated:* Use `associated_gateway_id` instead. The ID of the VGW with which to associate the gateway.
               Used for single account Direct Connect gateway associations.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_prefixes"] = allowed_prefixes
        __props__["associated_gateway_id"] = associated_gateway_id
        __props__["associated_gateway_owner_account_id"] = associated_gateway_owner_account_id
        __props__["associated_gateway_type"] = associated_gateway_type
        __props__["dx_gateway_association_id"] = dx_gateway_association_id
        __props__["dx_gateway_id"] = dx_gateway_id
        __props__["dx_gateway_owner_account_id"] = dx_gateway_owner_account_id
        __props__["proposal_id"] = proposal_id
        __props__["vpn_gateway_id"] = vpn_gateway_id
        return GatewayAssociation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

