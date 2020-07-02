# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class GetPeeringAttachmentResult:
    """
    A collection of values returned by getPeeringAttachment.
    """
    def __init__(__self__, filters=None, id=None, peer_account_id=None, peer_region=None, peer_transit_gateway_id=None, tags=None, transit_gateway_id=None):
        if filters and not isinstance(filters, list):
            raise TypeError("Expected argument 'filters' to be a list")
        __self__.filters = filters
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        if peer_account_id and not isinstance(peer_account_id, str):
            raise TypeError("Expected argument 'peer_account_id' to be a str")
        __self__.peer_account_id = peer_account_id
        """
        Identifier of the peer AWS account
        """
        if peer_region and not isinstance(peer_region, str):
            raise TypeError("Expected argument 'peer_region' to be a str")
        __self__.peer_region = peer_region
        """
        Identifier of the peer AWS region
        """
        if peer_transit_gateway_id and not isinstance(peer_transit_gateway_id, str):
            raise TypeError("Expected argument 'peer_transit_gateway_id' to be a str")
        __self__.peer_transit_gateway_id = peer_transit_gateway_id
        """
        Identifier of the peer EC2 Transit Gateway
        """
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        __self__.tags = tags
        if transit_gateway_id and not isinstance(transit_gateway_id, str):
            raise TypeError("Expected argument 'transit_gateway_id' to be a str")
        __self__.transit_gateway_id = transit_gateway_id
        """
        Identifier of the local EC2 Transit Gateway
        """


class AwaitableGetPeeringAttachmentResult(GetPeeringAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPeeringAttachmentResult(
            filters=self.filters,
            id=self.id,
            peer_account_id=self.peer_account_id,
            peer_region=self.peer_region,
            peer_transit_gateway_id=self.peer_transit_gateway_id,
            tags=self.tags,
            transit_gateway_id=self.transit_gateway_id)


def get_peering_attachment(filters=None,id=None,tags=None,opts=None):
    """
    Get information on an EC2 Transit Gateway Peering Attachment.

    ## Example Usage
    ### By Filter

    ```python
    import pulumi
    import pulumi_aws as aws

    example = aws.ec2transitgateway.get_peering_attachment(filters=[{
        "name": "transit-gateway-attachment-id",
        "values": ["tgw-attach-12345678"],
    }])
    ```
    ### By Identifier

    ```python
    import pulumi
    import pulumi_aws as aws

    attachment = aws.ec2transitgateway.get_peering_attachment(id="tgw-attach-12345678")
    ```


    :param list filters: One or more configuration blocks containing name-values filters. Detailed below.
    :param str id: Identifier of the EC2 Transit Gateway Peering Attachment.
    :param dict tags: A mapping of tags, each pair of which must exactly match
           a pair on the specific EC2 Transit Gateway Peering Attachment to retrieve.

    The **filters** object supports the following:

      * `name` (`str`) - The name of the field to filter by, as defined by
        [the underlying AWS API](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTransitGatewayPeeringAttachments.html).
      * `values` (`list`) - Set of values that are accepted for the given field.
        An EC2 Transit Gateway Peering Attachment be selected if any one of the given values matches.
    """
    __args__ = dict()

    __args__['filters'] = filters
    __args__['id'] = id
    __args__['tags'] = tags
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:ec2transitgateway/getPeeringAttachment:getPeeringAttachment', __args__, opts=opts).value

    return AwaitableGetPeeringAttachmentResult(
        filters=__ret__.get('filters'),
        id=__ret__.get('id'),
        peer_account_id=__ret__.get('peerAccountId'),
        peer_region=__ret__.get('peerRegion'),
        peer_transit_gateway_id=__ret__.get('peerTransitGatewayId'),
        tags=__ret__.get('tags'),
        transit_gateway_id=__ret__.get('transitGatewayId'))
