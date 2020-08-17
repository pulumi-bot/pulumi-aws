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
                 resource_name,
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
        Registers a custom domain name for use with AWS API Gateway. Additional information about this functionality
        can be found in the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).

        This resource just establishes ownership of and the TLS settings for
        a particular domain name. An API can be attached to a particular path
        under the registered domain name using
        the `apigateway.BasePathMapping` resource.

        API Gateway domains can be defined as either 'edge-optimized' or 'regional'.  In an edge-optimized configuration,
        API Gateway internally creates and manages a CloudFront distribution to route requests on the given hostname. In
        addition to this resource it's necessary to create a DNS record corresponding to the given domain name which is an alias
        (either Route53 alias or traditional CNAME) to the Cloudfront domain name exported in the `cloudfront_domain_name`
        attribute.

        In a regional configuration, API Gateway does not create a CloudFront distribution to route requests to the API, though
        a distribution can be created if needed. In either case, it is necessary to create a DNS record corresponding to the
        given domain name which is an alias (either Route53 alias or traditional CNAME) to the regional domain name exported in
        the `regional_domain_name` attribute.

        > **Note:** API Gateway requires the use of AWS Certificate Manager (ACM) certificates instead of Identity and Access Management (IAM) certificates in regions that support ACM. Regions that support ACM can be found in the [Regions and Endpoints Documentation](https://docs.aws.amazon.com/general/latest/gr/rande.html#acm_region). To import an existing private key and certificate into ACM or request an ACM certificate, see the `acm.Certificate` resource.

        > **Note:** The `apigateway.DomainName` resource expects dependency on the `acm.CertificateValidation` as
        only verified certificates can be used. This can be made either explicitly by adding the
        `depends_on = [aws_acm_certificate_validation.cert]` attribute. Or implicitly by referring certificate ARN
        from the validation resource where it will be available after the resource creation:
        `regional_certificate_arn = aws_acm_certificate_validation.cert.certificate_arn`.

        > **Note:** All arguments including the private key will be stored in the raw state as plain-text.

        ## Example Usage
        ### Edge Optimized (ACM Certificate)

        ```python
        import pulumi
        import pulumi_aws as aws

        example_domain_name = aws.apigateway.DomainName("exampleDomainName",
            certificate_arn=aws_acm_certificate_validation["example"]["certificate_arn"],
            domain_name="api.example.com")
        # Example DNS record using Route53.
        # Route53 is not specifically required; any DNS host can be used.
        example_record = aws.route53.Record("exampleRecord",
            name=example_domain_name.domain_name,
            type="A",
            zone_id=aws_route53_zone["example"]["id"],
            aliases=[{
                "evaluateTargetHealth": True,
                "name": example_domain_name.cloudfront_domain_name,
                "zone_id": example_domain_name.cloudfront_zone_id,
            }])
        ```
        ### Edge Optimized (IAM Certificate)

        ```python
        import pulumi
        import pulumi_aws as aws

        example_domain_name = aws.apigateway.DomainName("exampleDomainName",
            domain_name="api.example.com",
            certificate_name="example-api",
            certificate_body=(lambda path: open(path).read())(f"{path['module']}/example.com/example.crt"),
            certificate_chain=(lambda path: open(path).read())(f"{path['module']}/example.com/ca.crt"),
            certificate_private_key=(lambda path: open(path).read())(f"{path['module']}/example.com/example.key"))
        # Example DNS record using Route53.
        # Route53 is not specifically required; any DNS host can be used.
        example_record = aws.route53.Record("exampleRecord",
            zone_id=aws_route53_zone["example"]["id"],
            name=example_domain_name.domain_name,
            type="A",
            aliases=[{
                "name": example_domain_name.cloudfront_domain_name,
                "zone_id": example_domain_name.cloudfront_zone_id,
                "evaluateTargetHealth": True,
            }])
        ```
        ### Regional (ACM Certificate)

        ```python
        import pulumi
        import pulumi_aws as aws

        example_domain_name = aws.apigateway.DomainName("exampleDomainName",
            domain_name="api.example.com",
            regional_certificate_arn=aws_acm_certificate_validation["example"]["certificate_arn"],
            endpoint_configuration={
                "types": ["REGIONAL"],
            })
        # Example DNS record using Route53.
        # Route53 is not specifically required; any DNS host can be used.
        example_record = aws.route53.Record("exampleRecord",
            name=example_domain_name.domain_name,
            type="A",
            zone_id=aws_route53_zone["example"]["id"],
            aliases=[{
                "evaluateTargetHealth": True,
                "name": example_domain_name.regional_domain_name,
                "zone_id": example_domain_name.regional_zone_id,
            }])
        ```
        ### Regional (IAM Certificate)

        ```python
        import pulumi
        import pulumi_aws as aws

        example_domain_name = aws.apigateway.DomainName("exampleDomainName",
            certificate_body=(lambda path: open(path).read())(f"{path['module']}/example.com/example.crt"),
            certificate_chain=(lambda path: open(path).read())(f"{path['module']}/example.com/ca.crt"),
            certificate_private_key=(lambda path: open(path).read())(f"{path['module']}/example.com/example.key"),
            domain_name="api.example.com",
            regional_certificate_name="example-api",
            endpoint_configuration={
                "types": ["REGIONAL"],
            })
        # Example DNS record using Route53.
        # Route53 is not specifically required; any DNS host can be used.
        example_record = aws.route53.Record("exampleRecord",
            name=example_domain_name.domain_name,
            type="A",
            zone_id=aws_route53_zone["example"]["id"],
            aliases=[{
                "evaluateTargetHealth": True,
                "name": example_domain_name.regional_domain_name,
                "zone_id": example_domain_name.regional_zone_id,
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_body: The certificate issued for the domain name
               being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`.
        :param pulumi.Input[str] certificate_chain: The certificate for the CA that issued the
               certificate, along with any intermediate CA certificates required to
               create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
               `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_name: The unique name to use when registering this
               certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`. Required if `certificate_arn` is not set.
        :param pulumi.Input[str] certificate_private_key: The private key associated with the
               domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] domain_name: The fully-qualified domain name to register
        :param pulumi.Input[pulumi.InputType['DomainNameEndpointConfigurationArgs']] endpoint_configuration: Configuration block defining API endpoint information including type. Defined below.
        :param pulumi.Input[str] regional_certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        :param pulumi.Input[str] regional_certificate_name: The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
               `certificate_private_key`.
        :param pulumi.Input[str] security_policy: The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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
            id: str,
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
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN)
        :param pulumi.Input[str] certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_body: The certificate issued for the domain name
               being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`.
        :param pulumi.Input[str] certificate_chain: The certificate for the CA that issued the
               certificate, along with any intermediate CA certificates required to
               create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
               `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_name: The unique name to use when registering this
               certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
               `regional_certificate_name`. Required if `certificate_arn` is not set.
        :param pulumi.Input[str] certificate_private_key: The private key associated with the
               domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        :param pulumi.Input[str] certificate_upload_date: The upload date associated with the domain certificate.
        :param pulumi.Input[str] cloudfront_domain_name: The hostname created by Cloudfront to represent
               the distribution that implements this domain name mapping.
        :param pulumi.Input[str] cloudfront_zone_id: For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
               that can be used to create a Route53 alias record for the distribution.
        :param pulumi.Input[str] domain_name: The fully-qualified domain name to register
        :param pulumi.Input[pulumi.InputType['DomainNameEndpointConfigurationArgs']] endpoint_configuration: Configuration block defining API endpoint information including type. Defined below.
        :param pulumi.Input[str] regional_certificate_arn: The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        :param pulumi.Input[str] regional_certificate_name: The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
               `certificate_private_key`.
        :param pulumi.Input[str] regional_domain_name: The hostname for the custom domain's regional endpoint.
        :param pulumi.Input[str] regional_zone_id: The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        :param pulumi.Input[str] security_policy: The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Key-value map of resource tags
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
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN)
        """
        ...

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> Optional[str]:
        """
        The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when an edge-optimized domain name is desired. Conflicts with `certificate_name`, `certificate_body`, `certificate_chain`, `certificate_private_key`, `regional_certificate_arn`, and `regional_certificate_name`.
        """
        ...

    @property
    @pulumi.getter(name="certificateBody")
    def certificate_body(self) -> Optional[str]:
        """
        The certificate issued for the domain name
        being registered, in PEM format. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        `regional_certificate_name`.
        """
        ...

    @property
    @pulumi.getter(name="certificateChain")
    def certificate_chain(self) -> Optional[str]:
        """
        The certificate for the CA that issued the
        certificate, along with any intermediate CA certificates required to
        create an unbroken chain to a certificate trusted by the intended API clients. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`,
        `regional_certificate_arn`, and `regional_certificate_name`.
        """
        ...

    @property
    @pulumi.getter(name="certificateName")
    def certificate_name(self) -> Optional[str]:
        """
        The unique name to use when registering this
        certificate as an IAM server certificate. Conflicts with `certificate_arn`, `regional_certificate_arn`, and
        `regional_certificate_name`. Required if `certificate_arn` is not set.
        """
        ...

    @property
    @pulumi.getter(name="certificatePrivateKey")
    def certificate_private_key(self) -> Optional[str]:
        """
        The private key associated with the
        domain certificate given in `certificate_body`. Only valid for `EDGE` endpoint configuration type. Conflicts with `certificate_arn`, `regional_certificate_arn`, and `regional_certificate_name`.
        """
        ...

    @property
    @pulumi.getter(name="certificateUploadDate")
    def certificate_upload_date(self) -> str:
        """
        The upload date associated with the domain certificate.
        """
        ...

    @property
    @pulumi.getter(name="cloudfrontDomainName")
    def cloudfront_domain_name(self) -> str:
        """
        The hostname created by Cloudfront to represent
        the distribution that implements this domain name mapping.
        """
        ...

    @property
    @pulumi.getter(name="cloudfrontZoneId")
    def cloudfront_zone_id(self) -> str:
        """
        For convenience, the hosted zone ID (`Z2FDTNDATAQYW2`)
        that can be used to create a Route53 alias record for the distribution.
        """
        ...

    @property
    @pulumi.getter(name="domainName")
    def domain_name(self) -> str:
        """
        The fully-qualified domain name to register
        """
        ...

    @property
    @pulumi.getter(name="endpointConfiguration")
    def endpoint_configuration(self) -> 'outputs.DomainNameEndpointConfiguration':
        """
        Configuration block defining API endpoint information including type. Defined below.
        """
        ...

    @property
    @pulumi.getter(name="regionalCertificateArn")
    def regional_certificate_arn(self) -> Optional[str]:
        """
        The ARN for an AWS-managed certificate. AWS Certificate Manager is the only supported source. Used when a regional domain name is desired. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and `certificate_private_key`.
        """
        ...

    @property
    @pulumi.getter(name="regionalCertificateName")
    def regional_certificate_name(self) -> Optional[str]:
        """
        The user-friendly name of the certificate that will be used by regional endpoint for this domain name. Conflicts with `certificate_arn`, `certificate_name`, `certificate_body`, `certificate_chain`, and
        `certificate_private_key`.
        """
        ...

    @property
    @pulumi.getter(name="regionalDomainName")
    def regional_domain_name(self) -> str:
        """
        The hostname for the custom domain's regional endpoint.
        """
        ...

    @property
    @pulumi.getter(name="regionalZoneId")
    def regional_zone_id(self) -> str:
        """
        The hosted zone ID that can be used to create a Route53 alias record for the regional endpoint.
        """
        ...

    @property
    @pulumi.getter(name="securityPolicy")
    def security_policy(self) -> str:
        """
        The Transport Layer Security (TLS) version + cipher suite for this DomainName. The valid values are `TLS_1_0` and `TLS_1_2`. Must be configured to perform drift detection.
        """
        ...

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Key-value map of resource tags
        """
        ...

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

