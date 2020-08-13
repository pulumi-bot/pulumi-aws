# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import _utilities, _tables


class GatewayAssociationProposal(pulumi.CustomResource):
    allowed_prefixes: pulumi.Output[list]
    """
    VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
    """
    associated_gateway_id: pulumi.Output[str]
    """
    The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
    """
    associated_gateway_owner_account_id: pulumi.Output[str]
    """
    The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
    """
    associated_gateway_type: pulumi.Output[str]
    """
    The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
    """
    dx_gateway_id: pulumi.Output[str]
    """
    Direct Connect Gateway identifier.
    """
    dx_gateway_owner_account_id: pulumi.Output[str]
    """
    AWS Account identifier of the Direct Connect Gateway's owner.
    """
    def __init__(__self__, resource_name, opts=None, allowed_prefixes=None, associated_gateway_id=None, dx_gateway_id=None, dx_gateway_owner_account_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Direct Connect Gateway Association Proposal, typically for enabling cross-account associations. For single account associations, see the `directconnect.GatewayAssociation` resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.directconnect.GatewayAssociationProposal("example",
            dx_gateway_id=aws_dx_gateway["example"]["id"],
            dx_gateway_owner_account_id=aws_dx_gateway["example"]["owner_account_id"],
            associated_gateway_id=aws_vpn_gateway["example"]["id"])
        ```

        A full example of how to create a VPN Gateway in one AWS account, create a Direct Connect Gateway in a second AWS account, and associate the VPN Gateway with the Direct Connect Gateway via the `directconnect.GatewayAssociationProposal` and `directconnect.GatewayAssociation` resources can be found in [the `./examples/dx-gateway-cross-account-vgw-association` directory within the Github Repository](https://github.com/providers/provider-aws/tree/master/examples/dx-gateway-cross-account-vgw-association).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_prefixes: VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        :param pulumi.Input[str] associated_gateway_id: The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
        :param pulumi.Input[str] dx_gateway_id: Direct Connect Gateway identifier.
        :param pulumi.Input[str] dx_gateway_owner_account_id: AWS Account identifier of the Direct Connect Gateway's owner.
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

            __props__['allowedPrefixes'] = allowed_prefixes
            if associated_gateway_id is None:
                raise TypeError("Missing required property 'associated_gateway_id'")
            __props__['associatedGatewayId'] = associated_gateway_id
            if dx_gateway_id is None:
                raise TypeError("Missing required property 'dx_gateway_id'")
            __props__['dxGatewayId'] = dx_gateway_id
            if dx_gateway_owner_account_id is None:
                raise TypeError("Missing required property 'dx_gateway_owner_account_id'")
            __props__['dxGatewayOwnerAccountId'] = dx_gateway_owner_account_id
            __props__['associated_gateway_owner_account_id'] = None
            __props__['associated_gateway_type'] = None
        super(GatewayAssociationProposal, __self__).__init__(
            'aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_prefixes=None, associated_gateway_id=None, associated_gateway_owner_account_id=None, associated_gateway_type=None, dx_gateway_id=None, dx_gateway_owner_account_id=None):
        """
        Get an existing GatewayAssociationProposal resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_prefixes: VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
        :param pulumi.Input[str] associated_gateway_id: The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
        :param pulumi.Input[str] associated_gateway_owner_account_id: The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
        :param pulumi.Input[str] associated_gateway_type: The type of the associated gateway, `transitGateway` or `virtualPrivateGateway`.
        :param pulumi.Input[str] dx_gateway_id: Direct Connect Gateway identifier.
        :param pulumi.Input[str] dx_gateway_owner_account_id: AWS Account identifier of the Direct Connect Gateway's owner.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_prefixes"] = allowed_prefixes
        __props__["associated_gateway_id"] = associated_gateway_id
        __props__["associated_gateway_owner_account_id"] = associated_gateway_owner_account_id
        __props__["associated_gateway_type"] = associated_gateway_type
        __props__["dx_gateway_id"] = dx_gateway_id
        __props__["dx_gateway_owner_account_id"] = dx_gateway_owner_account_id
        return GatewayAssociationProposal(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
