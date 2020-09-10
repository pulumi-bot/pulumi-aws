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

__all__ = ['DomainName']


class DomainName(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 certificate_body: Optional[pulumi.Input[str]] = None,
                 certificate_chain: Optional[pulumi.Input[str]] = None,
                 certificate_name: Optional[pulumi.Input[str]] = None,
                 certificate_private_key: Optional[pulumi.Input[str]] = None,
                 domain_name: Optional[pulumi.Input[str]] = None,
                 endpoint_configuration: Optional[pulumi.Input[pulumi.InputType['DomainNameEndpointConfigurationArgs']]] = None,
                 regional_certificate_arn: Optional[pulumi.Input[str]] = None,
                 regional_certificate_name: Optional[pulumi.Input[str]] = None,
                 security_policy: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a DomainName resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            __props__['certificate_arn'] = certificate_arn
            __props__['certificate_body'] = certificate_body
            __props__['certificate_chain'] = certificate_chain
            __props__['certificate_name'] = certificate_name
            __props__['certificate_private_key'] = certificate_private_key
            if domain_name is None:
                raise TypeError("Missing required property 'domain_name'")
            __props__['domain_name'] = domain_name
            __props__['endpoint_configuration'] = endpoint_configuration
            __props__['regional_certificate_arn'] = regional_certificate_arn
            __props__['regional_certificate_name'] = regional_certificate_name
            __props__['security_policy'] = security_policy
            __props__['tags'] = tags
            __props__['arn'] = None
            __props__['certificate_upload_date'] = None
            __props__['cloudfront_domain_name'] = None
            __props__['cloudfront_zone_id'] = None
            __props__['regional_domain_name'] = None
            __props__['regional_zone_id'] = None
        super(DomainName, __self__).__init__(
            'aws:apigateway/domainName:DomainName',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            certificate_body: Optional[pulumi.Input[str]] = None,
            certificate_chain: Optional[pulumi.Input[str]] = None,
            certificate_name: Optional[pulumi.Input[str]] = None,
            certificate_private_key: Optional[pulumi.Input[str]] = None,
            certificate_upload_date: Optional[pulumi.Input[str]] = None,
            cloudfront_domain_name: Optional[pulumi.Input[str]] = None,
            cloudfront_zone_id: Optional[pulumi.Input[str]] = None,
            domain_name: Optional[pulumi.Input[str]] = None,
            endpoint_configuration: Optional[pulumi.Input[pulumi.InputType['DomainNameEndpointConfigurationArgs']]] = None,
            regional_certificate_arn: Optional[pulumi.Input[str]] = None,
            regional_certificate_name: Optional[pulumi.Input[str]] = None,
            regional_domain_name: Optional[pulumi.Input[str]] = None,
            regional_zone_id: Optional[pulumi.Input[str]] = None,
            security_policy: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'DomainName':
        """
        Get an existing DomainName resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["arn"] = arn
        __props__["certificate_arn"] = certificate_arn
        __props__["certificate_body"] = certificate_body
        __props__["certificate_chain"] = certificate_chain
        __props__["certificate_name"] = certificate_name
        __props__["certificate_private_key"] = certificate_private_key
        __props__["certificate_upload_date"] = certificate_upload_date
        __props__["cloudfront_domain_name"] = cloudfront_domain_name
        __props__["cloudfront_zone_id"] = cloudfront_zone_id
        __props__["domain_name"] = domain_name
        __props__["endpoint_configuration"] = endpoint_configuration
        __props__["regional_certificate_arn"] = regional_certificate_arn
        __props__["regional_certificate_name"] = regional_certificate_name
        __props__["regional_domain_name"] = regional_domain_name
        __props__["regional_zone_id"] = regional_zone_id
        __props__["security_policy"] = security_policy
        __props__["tags"] = tags
        return DomainName(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="certificateBody")
    def certificate_body(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate_body")

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate_chain")

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate_name")

    @property
    @pulumi.getter(name="certificatePrivateKey")
    def certificate_private_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "certificate_private_key")

    @property
    @pulumi.getter(name="certificateUploadDate")
    def certificate_upload_date(self) -> pulumi.Output[str]:
        return pulumi.get(self, "certificate_upload_date")

    @property
    @pulumi.getter(name="cloudfrontDomainName")
    def cloudfront_domain_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloudfront_domain_name")

    @property
    @pulumi.getter(name="cloudfrontZoneId")
    def cloudfront_zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "cloudfront_zone_id")

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "domain_name")

    @property
    @pulumi.getter(name="endpointConfiguration")
    def endpoint_configuration(self) -> pulumi.Output['outputs.DomainNameEndpointConfiguration']:
        return pulumi.get(self, "endpoint_configuration")

    @property
    @pulumi.getter(name="regionalCertificateArn")
    def regional_certificate_arn(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "regional_certificate_arn")

    @property
    @pulumi.getter(name="regionalCertificateName")
    def regional_certificate_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "regional_certificate_name")

    @property
    @pulumi.getter(name="regionalDomainName")
    def regional_domain_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "regional_domain_name")

    @property
    @pulumi.getter(name="regionalZoneId")
    def regional_zone_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "regional_zone_id")

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "security_policy")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        return pulumi.get(self, "tags")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

