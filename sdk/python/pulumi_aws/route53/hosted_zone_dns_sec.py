# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['HostedZoneDnsSec']


class HostedZoneDnsSec(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 hosted_zone_id: Optional[pulumi.Input[str]] = None,
                 signing_status: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages Route 53 Hosted Zone Domain Name System Security Extensions (DNSSEC). For more information about managing DNSSEC in Route 53, see the [Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-dnssec.html).

        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_aws as aws

        example_key = aws.kms.Key("exampleKey",
            customer_master_key_spec="ECC_NIST_P256",
            deletion_window_in_days=7,
            key_usage="SIGN_VERIFY",
            policy=json.dumps({
                "Statement": [
                    {
                        "Action": [
                            "kms:DescribeKey",
                            "kms:GetPublicKey",
                            "kms:Sign",
                        ],
                        "Effect": "Allow",
                        "Principal": {
                            "Service": "api-service.dnssec.route53.aws.internal",
                        },
                        "Sid": "Route 53 DNSSEC Permissions",
                    },
                    {
                        "Action": "kms:*",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "*",
                        },
                        "Resource": "*",
                        "Sid": "IAM User Permissions",
                    },
                ],
                "Version": "2012-10-17",
            }))
        example_zone = aws.route53.Zone("exampleZone")
        example_key_signing_key = aws.route53.KeySigningKey("exampleKeySigningKey",
            hosted_zone_id=aws_route53_zone["test"]["id"],
            key_management_service_arn=aws_kms_key["test"]["arn"])
        example_hosted_zone_dns_sec = aws.route53.HostedZoneDnsSec("exampleHostedZoneDnsSec", hosted_zone_id=example_key_signing_key.hosted_zone_id)
        ```

        ## Import

        `aws_route53_hosted_zone_dnssec` resources can be imported by using the Route 53 Hosted Zone identifier, e.g.

        ```sh
         $ pulumi import aws:route53/hostedZoneDnsSec:HostedZoneDnsSec example Z1D633PJN98FT9
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hosted_zone_id: Identifier of the Route 53 Hosted Zone.
        :param pulumi.Input[str] signing_status: Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
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

            if hosted_zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'hosted_zone_id'")
            __props__['hosted_zone_id'] = hosted_zone_id
            __props__['signing_status'] = signing_status
        super(HostedZoneDnsSec, __self__).__init__(
            'aws:route53/hostedZoneDnsSec:HostedZoneDnsSec',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            hosted_zone_id: Optional[pulumi.Input[str]] = None,
            signing_status: Optional[pulumi.Input[str]] = None) -> 'HostedZoneDnsSec':
        """
        Get an existing HostedZoneDnsSec resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] hosted_zone_id: Identifier of the Route 53 Hosted Zone.
        :param pulumi.Input[str] signing_status: Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["hosted_zone_id"] = hosted_zone_id
        __props__["signing_status"] = signing_status
        return HostedZoneDnsSec(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="hostedZoneId")
    def hosted_zone_id(self) -> pulumi.Output[str]:
        """
        Identifier of the Route 53 Hosted Zone.
        """
        return pulumi.get(self, "hosted_zone_id")

    @property
    @pulumi.getter(name="signingStatus")
    def signing_status(self) -> pulumi.Output[Optional[str]]:
        """
        Hosted Zone signing status. Valid values: `SIGNING`, `NOT_SIGNING`. Defaults to `SIGNING`.
        """
        return pulumi.get(self, "signing_status")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
