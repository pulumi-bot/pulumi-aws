# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['WebsiteCertificateAuthorityAssociation']


class WebsiteCertificateAuthorityAssociation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 fleet_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_aws as aws

        example = aws.worklink.Fleet("example")
        test = aws.worklink.WebsiteCertificateAuthorityAssociation("test",
            fleet_arn=aws_worklink_fleet["test"]["arn"],
            certificate=(lambda path: open(path).read())("certificate.pem"))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: The root certificate of the Certificate Authority.
        :param pulumi.Input[str] display_name: The certificate name to display.
        :param pulumi.Input[str] fleet_arn: The ARN of the fleet.
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

            if certificate is None:
                raise TypeError("Missing required property 'certificate'")
            __props__['certificate'] = certificate
            __props__['display_name'] = display_name
            if fleet_arn is None:
                raise TypeError("Missing required property 'fleet_arn'")
            __props__['fleet_arn'] = fleet_arn
            __props__['website_ca_id'] = None
        super(WebsiteCertificateAuthorityAssociation, __self__).__init__(
            'aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: str,
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            fleet_arn: Optional[pulumi.Input[str]] = None,
            website_ca_id: Optional[pulumi.Input[str]] = None) -> 'WebsiteCertificateAuthorityAssociation':
        """
        Get an existing WebsiteCertificateAuthorityAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate: The root certificate of the Certificate Authority.
        :param pulumi.Input[str] display_name: The certificate name to display.
        :param pulumi.Input[str] fleet_arn: The ARN of the fleet.
        :param pulumi.Input[str] website_ca_id: A unique identifier for the Certificate Authority.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate"] = certificate
        __props__["display_name"] = display_name
        __props__["fleet_arn"] = fleet_arn
        __props__["website_ca_id"] = website_ca_id
        return WebsiteCertificateAuthorityAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def certificate(self) -> str:
        """
        The root certificate of the Certificate Authority.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[str]:
        """
        The certificate name to display.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="fleetArn")
    def fleet_arn(self) -> str:
        """
        The ARN of the fleet.
        """
        return pulumi.get(self, "fleet_arn")

    @property
    @pulumi.getter(name="websiteCaId")
    def website_ca_id(self) -> str:
        """
        A unique identifier for the Certificate Authority.
        """
        return pulumi.get(self, "website_ca_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

