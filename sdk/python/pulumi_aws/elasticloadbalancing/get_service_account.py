# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetServiceAccountResult',
    'AwaitableGetServiceAccountResult',
    'get_service_account',
]

warnings.warn("aws.elasticloadbalancing.getServiceAccount has been deprecated in favor of aws.elb.getServiceAccount", DeprecationWarning)

@pulumi.output_type
class GetServiceAccountResult:
    """
    A collection of values returned by getServiceAccount.
    """
    def __init__(__self__, arn=None, id=None, region=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the AWS ELB service account in the selected region.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def region(self) -> Optional[str]:
        return pulumi.get(self, "region")


class AwaitableGetServiceAccountResult(GetServiceAccountResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountResult(
            arn=self.arn,
            id=self.id,
            region=self.region)


def get_service_account(region: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountResult:
    """
    Use this data source to get the Account ID of the [AWS Elastic Load Balancing Service Account](http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy)
    in a given region for the purpose of permitting in S3 bucket policy.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    main = aws.elb.get_service_account()
    elb_logs = aws.s3.Bucket("elbLogs",
        acl="private",
        policy=f\"\"\"{{
      "Id": "Policy",
      "Version": "2012-10-17",
      "Statement": [
        {{
          "Action": [
            "s3:PutObject"
          ],
          "Effect": "Allow",
          "Resource": "arn:aws:s3:::my-elb-tf-test-bucket/AWSLogs/*",
          "Principal": {{
            "AWS": [
              "{main.arn}"
            ]
          }}
        }}
      ]
    }}
    \"\"\")
    bar = aws.elb.LoadBalancer("bar",
        availability_zones=["us-west-2a"],
        access_logs=aws.elb.LoadBalancerAccessLogsArgs(
            bucket=elb_logs.bucket,
            interval=5,
        ),
        listeners=[aws.elb.LoadBalancerListenerArgs(
            instance_port=8000,
            instance_protocol="http",
            lb_port=80,
            lb_protocol="http",
        )])
    ```


    :param str region: Name of the region whose AWS ELB account ID is desired.
           Defaults to the region from the AWS provider configuration.
    """
    pulumi.log.warn("get_service_account is deprecated: aws.elasticloadbalancing.getServiceAccount has been deprecated in favor of aws.elb.getServiceAccount")
    __args__ = dict()
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:elasticloadbalancing/getServiceAccount:getServiceAccount', __args__, opts=opts, typ=GetServiceAccountResult).value

    return AwaitableGetServiceAccountResult(
        arn=__ret__.arn,
        id=__ret__.id,
        region=__ret__.region)
