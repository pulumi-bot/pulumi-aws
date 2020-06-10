# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class CertificateValidation(pulumi.CustomResource):
    certificate_arn: pulumi.Output[str]
    """
    The ARN of the certificate that is being validated.
    """
    validation_record_fqdns: pulumi.Output[list]
    """
    List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
    """
    def __init__(__self__, resource_name, opts=None, certificate_arn=None, validation_record_fqdns=None, __props__=None, __name__=None, __opts__=None):
        """
        This resource represents a successful validation of an ACM certificate in concert
        with other resources.

        Most commonly, this resource is used together with `route53.Record` and
        `acm.Certificate` to request a DNS validated certificate,
        deploy the required validation records and wait for validation to complete.

        > **WARNING:** This resource implements a part of the validation workflow. It does not represent a real-world entity in AWS, therefore changing or deleting this resource on its own has no immediate effect.


        ## Example Usage

        ### DNS Validation with Route 53

        ```python
        import pulumi
        import pulumi_aws as aws

        cert_certificate = aws.acm.Certificate("certCertificate",
            domain_name="example.com",
            validation_method="DNS")
        zone = aws.route53.get_zone(name="example.com.",
            private_zone=False)
        cert_validation = aws.route53.Record("certValidation",
            name=cert_certificate.domain_validation_options[0]["resourceRecordName"],
            records=[cert_certificate.domain_validation_options[0]["resourceRecordValue"]],
            ttl=60,
            type=cert_certificate.domain_validation_options[0]["resourceRecordType"],
            zone_id=zone.zone_id)
        cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation",
            certificate_arn=cert_certificate.arn,
            validation_record_fqdns=[cert_validation.fqdn])
        front_end = aws.lb.Listener("frontEnd", certificate_arn=cert_certificate_validation.certificate_arn)
        ```

        ### Alternative Domains DNS Validation with Route 53

        ```python
        import pulumi
        import pulumi_aws as aws

        cert_certificate = aws.acm.Certificate("certCertificate",
            domain_name="example.com",
            subject_alternative_names=[
                "www.example.com",
                "example.org",
            ],
            validation_method="DNS")
        zone = aws.route53.get_zone(name="example.com.",
            private_zone=False)
        zone_alt = aws.route53.get_zone(name="example.org.",
            private_zone=False)
        cert_validation = aws.route53.Record("certValidation",
            name=cert_certificate.domain_validation_options[0]["resourceRecordName"],
            records=[cert_certificate.domain_validation_options[0]["resourceRecordValue"]],
            ttl=60,
            type=cert_certificate.domain_validation_options[0]["resourceRecordType"],
            zone_id=zone.zone_id)
        cert_validation_alt1 = aws.route53.Record("certValidationAlt1",
            name=cert_certificate.domain_validation_options[1]["resourceRecordName"],
            records=[cert_certificate.domain_validation_options[1]["resourceRecordValue"]],
            ttl=60,
            type=cert_certificate.domain_validation_options[1]["resourceRecordType"],
            zone_id=zone.zone_id)
        cert_validation_alt2 = aws.route53.Record("certValidationAlt2",
            name=cert_certificate.domain_validation_options[2]["resourceRecordName"],
            records=[cert_certificate.domain_validation_options[2]["resourceRecordValue"]],
            ttl=60,
            type=cert_certificate.domain_validation_options[2]["resourceRecordType"],
            zone_id=zone_alt.zone_id)
        cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation",
            certificate_arn=cert_certificate.arn,
            validation_record_fqdns=[
                cert_validation.fqdn,
                cert_validation_alt1.fqdn,
                cert_validation_alt2.fqdn,
            ])
        front_end = aws.lb.Listener("frontEnd", certificate_arn=cert_certificate_validation.certificate_arn)
        ```

        ### Email Validation

        ```python
        import pulumi
        import pulumi_aws as aws

        cert_certificate = aws.acm.Certificate("certCertificate",
            domain_name="example.com",
            validation_method="EMAIL")
        cert_certificate_validation = aws.acm.CertificateValidation("certCertificateValidation", certificate_arn=cert_certificate.arn)
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate that is being validated.
        :param pulumi.Input[list] validation_record_fqdns: List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if certificate_arn is None:
                raise TypeError("Missing required property 'certificate_arn'")
            __props__['certificate_arn'] = certificate_arn
            __props__['validation_record_fqdns'] = validation_record_fqdns
        super(CertificateValidation, __self__).__init__(
            'aws:acm/certificateValidation:CertificateValidation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, certificate_arn=None, validation_record_fqdns=None):
        """
        Get an existing CertificateValidation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate that is being validated.
        :param pulumi.Input[list] validation_record_fqdns: List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_arn"] = certificate_arn
        __props__["validation_record_fqdns"] = validation_record_fqdns
        return CertificateValidation(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
