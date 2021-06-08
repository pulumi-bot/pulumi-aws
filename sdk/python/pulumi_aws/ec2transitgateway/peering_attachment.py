# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PeeringAttachmentArgs', 'PeeringAttachment']

@pulumi.input_type
class PeeringAttachmentArgs:
    def __init__(__self__, *,
                 peer_region: pulumi.Input[str],
                 peer_transit_gateway_id: pulumi.Input[str],
                 transit_gateway_id: pulumi.Input[str],
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a PeeringAttachment resource.
        :param pulumi.Input[str] peer_region: Region of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] peer_transit_gateway_id: Identifier of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        :param pulumi.Input[str] peer_account_id: Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        """
        pulumi.set(__self__, "peer_region", peer_region)
        pulumi.set(__self__, "peer_transit_gateway_id", peer_transit_gateway_id)
        pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)
        if peer_account_id is not None:
            pulumi.set(__self__, "peer_account_id", peer_account_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Input[str]:
        """
        Region of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_region", value)

    @property
    @pulumi.getter(name="peerTransitGatewayId")
    def peer_transit_gateway_id(self) -> pulumi.Input[str]:
        """
        Identifier of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_transit_gateway_id")

    @peer_transit_gateway_id.setter
    def peer_transit_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_transit_gateway_id", value)

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> pulumi.Input[str]:
        """
        Identifier of EC2 Transit Gateway.
        """
        return pulumi.get(self, "transit_gateway_id")

    @transit_gateway_id.setter
    def transit_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "transit_gateway_id", value)

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        """
        return pulumi.get(self, "peer_account_id")

    @peer_account_id.setter
    def peer_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_account_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


@pulumi.input_type
class _PeeringAttachmentState:
    def __init__(__self__, *,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PeeringAttachment resources.
        :param pulumi.Input[str] peer_account_id: Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        :param pulumi.Input[str] peer_region: Region of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] peer_transit_gateway_id: Identifier of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        """
        if peer_account_id is not None:
            pulumi.set(__self__, "peer_account_id", peer_account_id)
        if peer_region is not None:
            pulumi.set(__self__, "peer_region", peer_region)
        if peer_transit_gateway_id is not None:
            pulumi.set(__self__, "peer_transit_gateway_id", peer_transit_gateway_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)
        if transit_gateway_id is not None:
            pulumi.set(__self__, "transit_gateway_id", transit_gateway_id)

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        """
        return pulumi.get(self, "peer_account_id")

    @peer_account_id.setter
    def peer_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_account_id", value)

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> Optional[pulumi.Input[str]]:
        """
        Region of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_region")

    @peer_region.setter
    def peer_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_region", value)

    @property
    @pulumi.getter(name="peerTransitGatewayId")
    def peer_transit_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_transit_gateway_id")

    @peer_transit_gateway_id.setter
    def peer_transit_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_transit_gateway_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier of EC2 Transit Gateway.
        """
        return pulumi.get(self, "transit_gateway_id")

    @transit_gateway_id.setter
    def transit_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "transit_gateway_id", value)


class PeeringAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an EC2 Transit Gateway Peering Attachment.
        For examples of custom route table association and propagation, see the [EC2 Transit Gateway Networking Examples Guide](https://docs.aws.amazon.com/vpc/latest/tgw/TGW_Scenarios.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        local = pulumi.providers.Aws("local", region="us-east-1")
        peer = pulumi.providers.Aws("peer", region="us-west-2")
        peer_region = aws.get_region()
        local_transit_gateway = aws.ec2transitgateway.TransitGateway("localTransitGateway", tags={
            "Name": "Local TGW",
        },
        opts=pulumi.ResourceOptions(provider=aws["local"]))
        peer_transit_gateway = aws.ec2transitgateway.TransitGateway("peerTransitGateway", tags={
            "Name": "Peer TGW",
        },
        opts=pulumi.ResourceOptions(provider=aws["peer"]))
        example = aws.ec2transitgateway.PeeringAttachment("example",
            peer_account_id=peer_transit_gateway.owner_id,
            peer_region=peer_region.name,
            peer_transit_gateway_id=peer_transit_gateway.id,
            transit_gateway_id=local_transit_gateway.id,
            tags={
                "Name": "TGW Peering Requestor",
            })
        ```

        ## Import

        `aws_ec2_transit_gateway_peering_attachment` can be imported by using the EC2 Transit Gateway Attachment identifier, e.g.

        ```sh
         $ pulumi import aws:ec2transitgateway/peeringAttachment:PeeringAttachment example tgw-attach-12345678
        ```

        :param str resource_name_: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_account_id: Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        :param pulumi.Input[str] peer_region: Region of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] peer_transit_gateway_id: Identifier of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name_: str,
                 args: PeeringAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an EC2 Transit Gateway Peering Attachment.
        For examples of custom route table association and propagation, see the [EC2 Transit Gateway Networking Examples Guide](https://docs.aws.amazon.com/vpc/latest/tgw/TGW_Scenarios.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        local = pulumi.providers.Aws("local", region="us-east-1")
        peer = pulumi.providers.Aws("peer", region="us-west-2")
        peer_region = aws.get_region()
        local_transit_gateway = aws.ec2transitgateway.TransitGateway("localTransitGateway", tags={
            "Name": "Local TGW",
        },
        opts=pulumi.ResourceOptions(provider=aws["local"]))
        peer_transit_gateway = aws.ec2transitgateway.TransitGateway("peerTransitGateway", tags={
            "Name": "Peer TGW",
        },
        opts=pulumi.ResourceOptions(provider=aws["peer"]))
        example = aws.ec2transitgateway.PeeringAttachment("example",
            peer_account_id=peer_transit_gateway.owner_id,
            peer_region=peer_region.name,
            peer_transit_gateway_id=peer_transit_gateway.id,
            transit_gateway_id=local_transit_gateway.id,
            tags={
                "Name": "TGW Peering Requestor",
            })
        ```

        ## Import

        `aws_ec2_transit_gateway_peering_attachment` can be imported by using the EC2 Transit Gateway Attachment identifier, e.g.

        ```sh
         $ pulumi import aws:ec2transitgateway/peeringAttachment:PeeringAttachment example tgw-attach-12345678
        ```

        :param str resource_name_: The name of the resource.
        :param PeeringAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name_: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeeringAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name_, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name_, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name_: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 peer_account_id: Optional[pulumi.Input[str]] = None,
                 peer_region: Optional[pulumi.Input[str]] = None,
                 peer_transit_gateway_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 transit_gateway_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = PeeringAttachmentArgs.__new__(PeeringAttachmentArgs)

            __props__.__dict__["peer_account_id"] = peer_account_id
            if peer_region is None and not opts.urn:
                raise TypeError("Missing required property 'peer_region'")
            __props__.__dict__["peer_region"] = peer_region
            if peer_transit_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_transit_gateway_id'")
            __props__.__dict__["peer_transit_gateway_id"] = peer_transit_gateway_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["tags_all"] = tags_all
            if transit_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'transit_gateway_id'")
            __props__.__dict__["transit_gateway_id"] = transit_gateway_id
        super(PeeringAttachment, __self__).__init__(
            'aws:ec2transitgateway/peeringAttachment:PeeringAttachment',
            resource_name_,
            __props__,
            opts)

    @staticmethod
    def get(resource_name_: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            peer_account_id: Optional[pulumi.Input[str]] = None,
            peer_region: Optional[pulumi.Input[str]] = None,
            peer_transit_gateway_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            transit_gateway_id: Optional[pulumi.Input[str]] = None) -> 'PeeringAttachment':
        """
        Get an existing PeeringAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name_: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] peer_account_id: Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        :param pulumi.Input[str] peer_region: Region of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] peer_transit_gateway_id: Identifier of EC2 Transit Gateway to peer with.
        :param pulumi.Input[str] transit_gateway_id: Identifier of EC2 Transit Gateway.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PeeringAttachmentState.__new__(_PeeringAttachmentState)

        __props__.__dict__["peer_account_id"] = peer_account_id
        __props__.__dict__["peer_region"] = peer_region
        __props__.__dict__["peer_transit_gateway_id"] = peer_transit_gateway_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        __props__.__dict__["transit_gateway_id"] = transit_gateway_id
        return PeeringAttachment(resource_name_, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="peerAccountId")
    def peer_account_id(self) -> pulumi.Output[str]:
        """
        Account ID of EC2 Transit Gateway to peer with. Defaults to the account ID the current provider is currently connected to.
        """
        return pulumi.get(self, "peer_account_id")

    @property
    @pulumi.getter(name="peerRegion")
    def peer_region(self) -> pulumi.Output[str]:
        """
        Region of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_region")

    @property
    @pulumi.getter(name="peerTransitGatewayId")
    def peer_transit_gateway_id(self) -> pulumi.Output[str]:
        """
        Identifier of EC2 Transit Gateway to peer with.
        """
        return pulumi.get(self, "peer_transit_gateway_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        return pulumi.get(self, "tags_all")

    @property
    @pulumi.getter(name="transitGatewayId")
    def transit_gateway_id(self) -> pulumi.Output[str]:
        """
        Identifier of EC2 Transit Gateway.
        """
        return pulumi.get(self, "transit_gateway_id")

