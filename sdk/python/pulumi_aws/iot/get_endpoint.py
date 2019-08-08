# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class GetEndpointResult:
    """
    A collection of values returned by getEndpoint.
    """
    def __init__(__self__, endpoint_address=None, endpoint_type=None, id=None):
        if endpoint_address and not isinstance(endpoint_address, str):
            raise TypeError("Expected argument 'endpoint_address' to be a str")
        __self__.endpoint_address = endpoint_address
        """
        The endpoint based on `endpoint_type`:
        * No `endpoint_type`: Either `iot:Data` or `iot:Data-ATS` [depending on region](https://aws.amazon.com/blogs/iot/aws-iot-core-ats-endpoints/)
        * `iot:CredentialsProvider`: `IDENTIFIER.credentials.iot.REGION.amazonaws.com`
        * `iot:Data`: `IDENTIFIER.iot.REGION.amazonaws.com`
        * `iot:Data-ATS`: `IDENTIFIER-ats.iot.REGION.amazonaws.com`
        * `iot:Job`: `IDENTIFIER.jobs.iot.REGION.amazonaws.com`
        """
        if endpoint_type and not isinstance(endpoint_type, str):
            raise TypeError("Expected argument 'endpoint_type' to be a str")
        __self__.endpoint_type = endpoint_type
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            endpoint_address=self.endpoint_address,
            endpoint_type=self.endpoint_type,
            id=self.id)

def get_endpoint(endpoint_type=None,opts=None):
    """
    Returns a unique endpoint specific to the AWS account making the call.

    > This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/iot_endpoint.html.markdown.
    """
    __args__ = dict()

    __args__['endpointType'] = endpoint_type
    if opts is None:
        opts = pulumi.ResourceOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:iot/getEndpoint:getEndpoint', __args__, opts=opts).value

    return AwaitableGetEndpointResult(
        endpoint_address=__ret__.get('endpointAddress'),
        endpoint_type=__ret__.get('endpointType'),
        id=__ret__.get('id'))
