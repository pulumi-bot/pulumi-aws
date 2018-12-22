# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities, tables

class PeeringConnectionOptions(pulumi.CustomResource):
    """
    Provides a resource to manage VPC peering connection options.
    
    > **NOTE on VPC Peering Connections and VPC Peering Connection Options:** Terraform provides
    both a standalone VPC Peering Connection Options and a VPC Peering Connection
    resource with `accepter` and `requester` attributes. Do not manage options for the same VPC peering
    connection in both a VPC Peering Connection resource and a VPC Peering Connection Options resource.
    Doing so will cause a conflict of options and will overwrite the options.
    Using a VPC Peering Connection Options resource decouples management of the connection options from
    management of the VPC Peering Connection and allows options to be set correctly in cross-account scenarios.
    
    Basic usage:
    
    ```hcl
    resource "aws_vpc" "foo" {
      cidr_block = "10.0.0.0/16"
    }
    
    resource "aws_vpc" "bar" {
      cidr_block = "10.1.0.0/16"
    }
    
    resource "aws_vpc_peering_connection" "foo" {
      vpc_id      = "${aws_vpc.foo.id}"
      peer_vpc_id = "${aws_vpc.bar.id}"
      auto_accept = true
    }
    
    resource "aws_vpc_peering_connection_options" "foo" {
      vpc_peering_connection_id = "${aws_vpc_peering_connection.foo.id}"
    
      accepter {
        allow_remote_vpc_dns_resolution = true
      }
    
      requester {
        allow_vpc_to_remote_classic_link = true
        allow_classic_link_to_remote_vpc = true
      }
    }
    ```
    
    Basic cross-account usage:
    
    ```hcl
    provider "aws" {
      alias = "requester"
    
      # Requester's credentials.
    }
    
    provider "aws" {
      alias = "accepter"
    
      # Accepter's credentials.
    }
    
    resource "aws_vpc" "main" {
      provider = "aws.requester"
    
      cidr_block = "10.0.0.0/16"
    
      enable_dns_support   = true
      enable_dns_hostnames = true
    }
    
    resource "aws_vpc" "peer" {
      provider = "aws.accepter"
    
      cidr_block = "10.1.0.0/16"
    
      enable_dns_support   = true
      enable_dns_hostnames = true
    }
    
    data "aws_caller_identity" "peer" {
      provider = "aws.accepter"
    }
    
    # Requester's side of the connection.
    resource "aws_vpc_peering_connection" "peer" {
      provider = "aws.requester"
    
      vpc_id        = "${aws_vpc.main.id}"
      peer_vpc_id   = "${aws_vpc.peer.id}"
      peer_owner_id = "${data.aws_caller_identity.peer.account_id}"
      auto_accept   = false
    
      tags = {
        Side = "Requester"
      }
    }
    
    # Accepter's side of the connection.
    resource "aws_vpc_peering_connection_accepter" "peer" {
      provider = "aws.accepter"
    
      vpc_peering_connection_id = "${aws_vpc_peering_connection.peer.id}"
      auto_accept               = true
    
      tags = {
        Side = "Accepter"
      }
    }
    
    resource "aws_vpc_peering_connection_options" "requester" {
      provider = "aws.requester"
    
      # As options can't be set until the connection has been accepted
      # create an explicit dependency on the accepter.
      vpc_peering_connection_id = "${aws_vpc_peering_connection_accepter.peer.id}"
    
      requester {
        allow_remote_vpc_dns_resolution = true
      }
    }
    
    resource "aws_vpc_peering_connection_options" "accepter" {
      provider = "aws.accepter"
    
      vpc_peering_connection_id = "${aws_vpc_peering_connection_accepter.peer.id}"
    
      accepter {
        allow_remote_vpc_dns_resolution = true
      }
    }
    ```
    """
    def __init__(__self__, __name__, __opts__=None, accepter=None, requester=None, vpc_peering_connection_id=None):
        """Create a PeeringConnectionOptions resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['accepter'] = accepter

        __props__['requester'] = requester

        if not vpc_peering_connection_id:
            raise TypeError('Missing required property vpc_peering_connection_id')
        __props__['vpc_peering_connection_id'] = vpc_peering_connection_id

        super(PeeringConnectionOptions, __self__).__init__(
            'aws:ec2/peeringConnectionOptions:PeeringConnectionOptions',
            __name__,
            __props__,
            __opts__)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

