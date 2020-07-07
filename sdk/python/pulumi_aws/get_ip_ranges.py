# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class GetIpRangesResult:
    """
    A collection of values returned by getIpRanges.
    """
    def __init__(__self__, cidr_blocks=None, create_date=None, id=None, ipv6_cidr_blocks=None, regions=None, services=None, sync_token=None, url=None):
        if cidr_blocks and not isinstance(cidr_blocks, list):
            raise TypeError("Expected argument 'cidr_blocks' to be a list")
        __self__.cidr_blocks = cidr_blocks
        """
        The lexically ordered list of CIDR blocks.
        """
        if create_date and not isinstance(create_date, str):
            raise TypeError("Expected argument 'create_date' to be a str")
        __self__.create_date = create_date
        """
        The publication time of the IP ranges (e.g. `2016-08-03-23-46-05`).
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if ipv6_cidr_blocks and not isinstance(ipv6_cidr_blocks, list):
            raise TypeError("Expected argument 'ipv6_cidr_blocks' to be a list")
        __self__.ipv6_cidr_blocks = ipv6_cidr_blocks
        """
        The lexically ordered list of IPv6 CIDR blocks.
        """
        if regions and not isinstance(regions, list):
            raise TypeError("Expected argument 'regions' to be a list")
        __self__.regions = regions
        if services and not isinstance(services, list):
            raise TypeError("Expected argument 'services' to be a list")
        __self__.services = services
        if sync_token and not isinstance(sync_token, float):
            raise TypeError("Expected argument 'sync_token' to be a float")
        __self__.sync_token = sync_token
        """
        The publication time of the IP ranges, in Unix epoch time format
        (e.g. `1470267965`).
        """
        if url and not isinstance(url, str):
            raise TypeError("Expected argument 'url' to be a str")
        __self__.url = url


class AwaitableGetIpRangesResult(GetIpRangesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpRangesResult(
            cidr_blocks=self.cidr_blocks,
            create_date=self.create_date,
            id=self.id,
            ipv6_cidr_blocks=self.ipv6_cidr_blocks,
            regions=self.regions,
            services=self.services,
            sync_token=self.sync_token,
            url=self.url)


def get_ip_ranges(regions=None, services=None, url=None, opts=None):
    """
    Use this data source to get the IP ranges of various AWS products and services. For more information about the contents of this data source and required JSON syntax if referencing a custom URL, see the [AWS IP Address Ranges documention](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    european_ec2 = aws.get_ip_ranges(regions=[
            "eu-west-1",
            "eu-central-1",
        ],
        services=["ec2"])
    from_europe = aws.ec2.SecurityGroup("fromEurope",
        ingress=[{
            "from_port": "443",
            "to_port": "443",
            "protocol": "tcp",
            "cidr_blocks": european_ec2.cidr_blocks,
            "ipv6_cidr_blocks": european_ec2.ipv6_cidr_blocks,
        }],
        tags={
            "CreateDate": european_ec2.create_date,
            "SyncToken": european_ec2.sync_token,
        })
    ```


    :param list regions: Filter IP ranges by regions (or include all regions, if
           omitted). Valid items are `global` (for `cloudfront`) as well as all AWS regions
           (e.g. `eu-central-1`)
    :param list services: Filter IP ranges by services. Valid items are `amazon`
           (for amazon.com), `amazon_connect`, `api_gateway`, `cloud9`, `cloudfront`,
           `codebuild`, `dynamodb`, `ec2`, `ec2_instance_connect`, `globalaccelerator`,
           `route53`, `route53_healthchecks`, `s3` and `workspaces_gateways`. See the
           [`service` attribute][2] documentation for other possible values.
    :param str url: Custom URL for source JSON file. Syntax must match [AWS IP Address Ranges documention](https://docs.aws.amazon.com/general/latest/gr/aws-ip-ranges.html). Defaults to `https://ip-ranges.amazonaws.com/ip-ranges.json`.
    """
    __args__ = dict()
    __args__['regions'] = regions
    __args__['services'] = services
    __args__['url'] = url
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getIpRanges:getIpRanges', __args__, opts=opts).value

    return AwaitableGetIpRangesResult(
        cidr_blocks=__ret__.get('cidrBlocks'),
        create_date=__ret__.get('createDate'),
        id=__ret__.get('id'),
        ipv6_cidr_blocks=__ret__.get('ipv6CidrBlocks'),
        regions=__ret__.get('regions'),
        services=__ret__.get('services'),
        sync_token=__ret__.get('syncToken'),
        url=__ret__.get('url'))
