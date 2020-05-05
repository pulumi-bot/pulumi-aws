# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetArnResult:
    """
    A collection of values returned by getArn.
    """
    def __init__(__self__, account=None, arn=None, id=None, partition=None, region=None, resource=None, service=None):
        if account and not isinstance(account, str):
            raise TypeError("Expected argument 'account' to be a str")
        __self__.account = account
        """
        The [ID](https://docs.aws.amazon.com/general/latest/gr/acct-identifiers.html) of the AWS account that owns the resource, without the hyphens.
        """
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        __self__.arn = arn
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if partition and not isinstance(partition, str):
            raise TypeError("Expected argument 'partition' to be a str")
        __self__.partition = partition
        """
        The partition that the resource is in.
        """
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        __self__.region = region
        """
        The region the resource resides in.
        Note that the ARNs for some resources do not require a region, so this component might be omitted.
        """
        if resource and not isinstance(resource, str):
            raise TypeError("Expected argument 'resource' to be a str")
        __self__.resource = resource
        """
        The content of this part of the ARN varies by service.
        It often includes an indicator of the type of resource—for example, an IAM user or Amazon RDS database —followed by a slash (/) or a colon (:), followed by the resource name itself.
        """
        if service and not isinstance(service, str):
            raise TypeError("Expected argument 'service' to be a str")
        __self__.service = service
        """
        The [service namespace](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces) that identifies the AWS product.
        """
class AwaitableGetArnResult(GetArnResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetArnResult(
            account=self.account,
            arn=self.arn,
            id=self.id,
            partition=self.partition,
            region=self.region,
            resource=self.resource,
            service=self.service)

def get_arn(arn=None,opts=None):
    """
    Parses an Amazon Resource Name (ARN) into its constituent parts.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_aws as aws

    db_instance = aws.get_arn(arn="arn:aws:rds:eu-west-1:123456789012:db:mysql-db")
    ```



    :param str arn: The ARN to parse.
    """
    __args__ = dict()


    __args__['arn'] = arn
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:index/getArn:getArn', __args__, opts=opts).value

    return AwaitableGetArnResult(
        account=__ret__.get('account'),
        arn=__ret__.get('arn'),
        id=__ret__.get('id'),
        partition=__ret__.get('partition'),
        region=__ret__.get('region'),
        resource=__ret__.get('resource'),
        service=__ret__.get('service'))
