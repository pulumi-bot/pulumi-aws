# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Fleet']


class Fleet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_stream_arn: Optional[pulumi.Input[str]] = None,
                 device_ca_certificate: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 identity_provider: Optional[pulumi.Input[pulumi.InputType['FleetIdentityProviderArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network: Optional[pulumi.Input[pulumi.InputType['FleetNetworkArgs']]] = None,
                 optimize_for_end_user_location: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        Basic usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.worklink.Fleet("example")
        ```

        Network Configuration Usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.worklink.Fleet("example", network={
            "vpc_id": aws_vpc["test"]["id"],
            "subnet_ids": [[__item["id"] for __item in aws_subnet["test"]]],
            "security_group_ids": [aws_security_group["test"]["id"]],
        })
        ```

        Identity Provider Configuration Usage:

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.worklink.Fleet("test", identity_provider={
            "type": "SAML",
            "samlMetadata": (lambda path: open(path).read())("saml-metadata.xml"),
        })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] audit_stream_arn: The ARN of the Amazon Kinesis data stream that receives the audit events.
        :param pulumi.Input[str] device_ca_certificate: The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        :param pulumi.Input[str] display_name: The name of the fleet.
        :param pulumi.Input[pulumi.InputType['FleetIdentityProviderArgs']] identity_provider: Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        :param pulumi.Input[str] name: A region-unique name for the AMI.
        :param pulumi.Input[pulumi.InputType['FleetNetworkArgs']] network: Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        :param pulumi.Input[bool] optimize_for_end_user_location: The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
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

            __props__['audit_stream_arn'] = audit_stream_arn
            __props__['device_ca_certificate'] = device_ca_certificate
            __props__['display_name'] = display_name
            __props__['identity_provider'] = identity_provider
            __props__['name'] = name
            __props__['network'] = network
            __props__['optimize_for_end_user_location'] = optimize_for_end_user_location
            __props__['arn'] = None
            __props__['company_code'] = None
            __props__['created_time'] = None
            __props__['last_updated_time'] = None
        super(Fleet, __self__).__init__(
            'aws:worklink/fleet:Fleet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            audit_stream_arn: Optional[pulumi.Input[str]] = None,
            company_code: Optional[pulumi.Input[str]] = None,
            created_time: Optional[pulumi.Input[str]] = None,
            device_ca_certificate: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            identity_provider: Optional[pulumi.Input[pulumi.InputType['FleetIdentityProviderArgs']]] = None,
            last_updated_time: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network: Optional[pulumi.Input[pulumi.InputType['FleetNetworkArgs']]] = None,
            optimize_for_end_user_location: Optional[pulumi.Input[bool]] = None) -> 'Fleet':
        """
        Get an existing Fleet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: The ARN of the created WorkLink Fleet.
        :param pulumi.Input[str] audit_stream_arn: The ARN of the Amazon Kinesis data stream that receives the audit events.
        :param pulumi.Input[str] company_code: The identifier used by users to sign in to the Amazon WorkLink app.
        :param pulumi.Input[str] created_time: The time that the fleet was created.
        :param pulumi.Input[str] device_ca_certificate: The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        :param pulumi.Input[str] display_name: The name of the fleet.
        :param pulumi.Input[pulumi.InputType['FleetIdentityProviderArgs']] identity_provider: Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        :param pulumi.Input[str] last_updated_time: The time that the fleet was last updated.
        :param pulumi.Input[str] name: A region-unique name for the AMI.
        :param pulumi.Input[pulumi.InputType['FleetNetworkArgs']] network: Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        :param pulumi.Input[bool] optimize_for_end_user_location: The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["audit_stream_arn"] = audit_stream_arn
        __props__["company_code"] = company_code
        __props__["created_time"] = created_time
        __props__["device_ca_certificate"] = device_ca_certificate
        __props__["display_name"] = display_name
        __props__["identity_provider"] = identity_provider
        __props__["last_updated_time"] = last_updated_time
        __props__["name"] = name
        __props__["network"] = network
        __props__["optimize_for_end_user_location"] = optimize_for_end_user_location
        return Fleet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        The ARN of the created WorkLink Fleet.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="auditStreamArn")
    def audit_stream_arn(self) -> Optional[str]:
        """
        The ARN of the Amazon Kinesis data stream that receives the audit events.
        """
        return pulumi.get(self, "audit_stream_arn")

    @property
    @pulumi.getter(name="companyCode")
    def company_code(self) -> str:
        """
        The identifier used by users to sign in to the Amazon WorkLink app.
        """
        return pulumi.get(self, "company_code")

    @property
    @pulumi.getter(name="createdTime")
    def created_time(self) -> str:
        """
        The time that the fleet was created.
        """
        return pulumi.get(self, "created_time")

    @property
    @pulumi.getter(name="deviceCaCertificate")
    def device_ca_certificate(self) -> Optional[str]:
        """
        The certificate chain, including intermediate certificates and the root certificate authority certificate used to issue device certificates.
        """
        return pulumi.get(self, "device_ca_certificate")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The name of the fleet.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="identityProvider")
    def identity_provider(self) -> Optional['outputs.FleetIdentityProvider']:
        """
        Provide this to allow manage the identity provider configuration for the fleet. Fields documented below.
        """
        return pulumi.get(self, "identity_provider")

    @property
    @pulumi.getter(name="lastUpdatedTime")
    def last_updated_time(self) -> str:
        """
        The time that the fleet was last updated.
        """
        return pulumi.get(self, "last_updated_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A region-unique name for the AMI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def network(self) -> Optional['outputs.FleetNetwork']:
        """
        Provide this to allow manage the company network configuration for the fleet. Fields documented below.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter(name="optimizeForEndUserLocation")
    def optimize_for_end_user_location(self) -> Optional[bool]:
        """
        The option to optimize for better performance by routing traffic through the closest AWS Region to users, which may be outside of your home Region. Defaults to `true`.
        """
        return pulumi.get(self, "optimize_for_end_user_location")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

