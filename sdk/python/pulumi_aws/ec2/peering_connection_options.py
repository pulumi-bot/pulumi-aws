# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['PeeringConnectionOptionsArgs', 'PeeringConnectionOptions']

@pulumi.input_type
class PeeringConnectionOptionsArgs:
    def __init__(__self__, *,
                 vpc_peering_connection_id: pulumi.Input[str],
                 accepter: Optional[pulumi.Input['PeeringConnectionOptionsAccepterArgs']] = None,
                 requester: Optional[pulumi.Input['PeeringConnectionOptionsRequesterArgs']] = None):
        """
        The set of arguments for constructing a PeeringConnectionOptions resource.
        :param pulumi.Input[str] vpc_peering_connection_id: The ID of the requester VPC peering connection.
        :param pulumi.Input['PeeringConnectionOptionsAccepterArgs'] accepter: An optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
               the peering connection (a maximum of one).
        :param pulumi.Input['PeeringConnectionOptionsRequesterArgs'] requester: A optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
               the peering connection (a maximum of one).
        """
        pulumi.set(__self__, "vpc_peering_connection_id", vpc_peering_connection_id)
        if accepter is not None:
            pulumi.set(__self__, "accepter", accepter)
        if requester is not None:
            pulumi.set(__self__, "requester", requester)

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Input[str]:
        """
        The ID of the requester VPC peering connection.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    @vpc_peering_connection_id.setter
    def vpc_peering_connection_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_peering_connection_id", value)

    @property
    @pulumi.getter
    def accepter(self) -> Optional[pulumi.Input['PeeringConnectionOptionsAccepterArgs']]:
        """
        An optional configuration block that allows for [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "accepter")

    @accepter.setter
    def accepter(self, value: Optional[pulumi.Input['PeeringConnectionOptionsAccepterArgs']]):
        pulumi.set(self, "accepter", value)

    @property
    @pulumi.getter
    def requester(self) -> Optional[pulumi.Input['PeeringConnectionOptionsRequesterArgs']]:
        """
        A optional configuration block that allows for [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "requester")

    @requester.setter
    def requester(self, value: Optional[pulumi.Input['PeeringConnectionOptionsRequesterArgs']]):
        pulumi.set(self, "requester", value)


class PeeringConnectionOptions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsAccepterArgs']]] = None,
                 requester: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsRequesterArgs']]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage VPC peering connection options.

        > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
        both a standalone VPC Peering Connection Options and a VPC Peering Connection
        resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
        connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
        Doing so will cause a conflict of options and will overwrite the options.
        Using a VPC Peering Connection Options resource decouples management of the connection options from
        management of the VPC Peering Connection and allows options to be set correctly in cross-region and
        cross-account scenarios.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.0.0.0/16")
        bar = aws.ec2.Vpc("bar", cidr_block="10.1.0.0/16")
        foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
            vpc_id=foo_vpc.id,
            peer_vpc_id=bar.id,
            auto_accept=True)
        foo_peering_connection_options = aws.ec2.PeeringConnectionOptions("fooPeeringConnectionOptions",
            vpc_peering_connection_id=foo_vpc_peering_connection.id,
            accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            requester=aws.ec2.PeeringConnectionOptionsRequesterArgs(
                allow_vpc_to_remote_classic_link=True,
                allow_classic_link_to_remote_vpc=True,
            ))
        ```
        ### Cross-Account Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        requester = pulumi.providers.Aws("requester")
        # Requester's credentials.
        accepter = pulumi.providers.Aws("accepter")
        # Accepter's credentials.
        main = aws.ec2.Vpc("main",
            cidr_block="10.0.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True,
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        peer_vpc = aws.ec2.Vpc("peerVpc",
            cidr_block="10.1.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True,
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        peer_caller_identity = aws.get_caller_identity()
        # Requester's side of the connection.
        peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
            vpc_id=main.id,
            peer_vpc_id=peer_vpc.id,
            peer_owner_id=peer_caller_identity.account_id,
            auto_accept=False,
            tags={
                "Side": "Requester",
            },
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        # Accepter's side of the connection.
        peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
            vpc_peering_connection_id=peer_vpc_peering_connection.id,
            auto_accept=True,
            tags={
                "Side": "Accepter",
            },
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        requester_peering_connection_options = aws.ec2.PeeringConnectionOptions("requesterPeeringConnectionOptions",
            vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
            requester=aws.ec2.PeeringConnectionOptionsRequesterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        accepter_peering_connection_options = aws.ec2.PeeringConnectionOptions("accepterPeeringConnectionOptions",
            vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
            accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        ```

        ## Import

        VPC Peering Connection Options can be imported using the `vpc peering id`, e.g.

        ```sh
         $ pulumi import aws:ec2/peeringConnectionOptions:PeeringConnectionOptions foo pcx-111aaa111
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PeeringConnectionOptionsAccepterArgs']] accepter: An optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
               the peering connection (a maximum of one).
        :param pulumi.Input[pulumi.InputType['PeeringConnectionOptionsRequesterArgs']] requester: A optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
               the peering connection (a maximum of one).
        :param pulumi.Input[str] vpc_peering_connection_id: The ID of the requester VPC peering connection.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PeeringConnectionOptionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage VPC peering connection options.

        > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** This provider provides
        both a standalone VPC Peering Connection Options and a VPC Peering Connection
        resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
        connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
        Doing so will cause a conflict of options and will overwrite the options.
        Using a VPC Peering Connection Options resource decouples management of the connection options from
        management of the VPC Peering Connection and allows options to be set correctly in cross-region and
        cross-account scenarios.

        ## Example Usage
        ### Basic Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        foo_vpc = aws.ec2.Vpc("fooVpc", cidr_block="10.0.0.0/16")
        bar = aws.ec2.Vpc("bar", cidr_block="10.1.0.0/16")
        foo_vpc_peering_connection = aws.ec2.VpcPeeringConnection("fooVpcPeeringConnection",
            vpc_id=foo_vpc.id,
            peer_vpc_id=bar.id,
            auto_accept=True)
        foo_peering_connection_options = aws.ec2.PeeringConnectionOptions("fooPeeringConnectionOptions",
            vpc_peering_connection_id=foo_vpc_peering_connection.id,
            accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            requester=aws.ec2.PeeringConnectionOptionsRequesterArgs(
                allow_vpc_to_remote_classic_link=True,
                allow_classic_link_to_remote_vpc=True,
            ))
        ```
        ### Cross-Account Usage

        ```python
        import pulumi
        import pulumi_aws as aws
        import pulumi_pulumi as pulumi

        requester = pulumi.providers.Aws("requester")
        # Requester's credentials.
        accepter = pulumi.providers.Aws("accepter")
        # Accepter's credentials.
        main = aws.ec2.Vpc("main",
            cidr_block="10.0.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True,
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        peer_vpc = aws.ec2.Vpc("peerVpc",
            cidr_block="10.1.0.0/16",
            enable_dns_support=True,
            enable_dns_hostnames=True,
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        peer_caller_identity = aws.get_caller_identity()
        # Requester's side of the connection.
        peer_vpc_peering_connection = aws.ec2.VpcPeeringConnection("peerVpcPeeringConnection",
            vpc_id=main.id,
            peer_vpc_id=peer_vpc.id,
            peer_owner_id=peer_caller_identity.account_id,
            auto_accept=False,
            tags={
                "Side": "Requester",
            },
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        # Accepter's side of the connection.
        peer_vpc_peering_connection_accepter = aws.ec2.VpcPeeringConnectionAccepter("peerVpcPeeringConnectionAccepter",
            vpc_peering_connection_id=peer_vpc_peering_connection.id,
            auto_accept=True,
            tags={
                "Side": "Accepter",
            },
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        requester_peering_connection_options = aws.ec2.PeeringConnectionOptions("requesterPeeringConnectionOptions",
            vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
            requester=aws.ec2.PeeringConnectionOptionsRequesterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            opts=pulumi.ResourceOptions(provider=aws["requester"]))
        accepter_peering_connection_options = aws.ec2.PeeringConnectionOptions("accepterPeeringConnectionOptions",
            vpc_peering_connection_id=peer_vpc_peering_connection_accepter.id,
            accepter=aws.ec2.PeeringConnectionOptionsAccepterArgs(
                allow_remote_vpc_dns_resolution=True,
            ),
            opts=pulumi.ResourceOptions(provider=aws["accepter"]))
        ```

        ## Import

        VPC Peering Connection Options can be imported using the `vpc peering id`, e.g.

        ```sh
         $ pulumi import aws:ec2/peeringConnectionOptions:PeeringConnectionOptions foo pcx-111aaa111
        ```

        :param str resource_name: The name of the resource.
        :param PeeringConnectionOptionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PeeringConnectionOptionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 accepter: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsAccepterArgs']]] = None,
                 requester: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsRequesterArgs']]] = None,
                 vpc_peering_connection_id: Optional[pulumi.Input[str]] = None,
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

            __props__['accepter'] = accepter
            __props__['requester'] = requester
            if vpc_peering_connection_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_peering_connection_id'")
            __props__['vpc_peering_connection_id'] = vpc_peering_connection_id
        super(PeeringConnectionOptions, __self__).__init__(
            'aws:ec2/peeringConnectionOptions:PeeringConnectionOptions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accepter: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsAccepterArgs']]] = None,
            requester: Optional[pulumi.Input[pulumi.InputType['PeeringConnectionOptionsRequesterArgs']]] = None,
            vpc_peering_connection_id: Optional[pulumi.Input[str]] = None) -> 'PeeringConnectionOptions':
        """
        Get an existing PeeringConnectionOptions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['PeeringConnectionOptionsAccepterArgs']] accepter: An optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
               the peering connection (a maximum of one).
        :param pulumi.Input[pulumi.InputType['PeeringConnectionOptionsRequesterArgs']] requester: A optional configuration block that allows for [VPC Peering Connection]
               (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
               the peering connection (a maximum of one).
        :param pulumi.Input[str] vpc_peering_connection_id: The ID of the requester VPC peering connection.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accepter"] = accepter
        __props__["requester"] = requester
        __props__["vpc_peering_connection_id"] = vpc_peering_connection_id
        return PeeringConnectionOptions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accepter(self) -> pulumi.Output['outputs.PeeringConnectionOptionsAccepter']:
        """
        An optional configuration block that allows for [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that accepts
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "accepter")

    @property
    @pulumi.getter
    def requester(self) -> pulumi.Output['outputs.PeeringConnectionOptionsRequester']:
        """
        A optional configuration block that allows for [VPC Peering Connection]
        (https://docs.aws.amazon.com/vpc/latest/peering/what-is-vpc-peering.html) options to be set for the VPC that requests
        the peering connection (a maximum of one).
        """
        return pulumi.get(self, "requester")

    @property
    @pulumi.getter(name="vpcPeeringConnectionId")
    def vpc_peering_connection_id(self) -> pulumi.Output[str]:
        """
        The ID of the requester VPC peering connection.
        """
        return pulumi.get(self, "vpc_peering_connection_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

