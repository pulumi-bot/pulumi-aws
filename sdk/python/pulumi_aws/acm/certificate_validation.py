# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['CertificateValidation']


class CertificateValidation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_arn: Optional[pulumi.Input[str]] = None,
                 validation_record_fqdns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

        example_certificate = aws.acm.Certificate("exampleCertificate",
            domain_name="example.com",
            validation_method="DNS")
        example_zone = aws.route53.get_zone(name="example.com",
            private_zone=False)
        example_record = []
        for range in [{"key": k, "value": v} for [k, v] in enumerate({dvo.domainName: {
            name: dvo.resourceRecordName,
            record: dvo.resourceRecordValue,
            type: dvo.resourceRecordType,
        } for dvo in example_certificate.domainValidationOptions})]:
            example_record.append(aws.route53.Record(f"exampleRecord-{range['key']}",
                allow_overwrite=True,
                name=range["value"]["name"],
                records=[range["value"]["record"]],
                ttl=60,
                type=range["value"]["type"],
                zone_id=example_zone.zone_id))
        example_certificate_validation = aws.acm.CertificateValidation("exampleCertificateValidation",
            certificate_arn=example_certificate.arn,
            validation_record_fqdns=example_record.apply(lambda example_record: [record.fqdn for record in example_record]))
        # ... other configuration ...
        example_listener = aws.lb.Listener("exampleListener", certificate_arn=example_certificate_validation.certificate_arn)
        ```
        ### Email Validation

        In this situation, the resource is simply a waiter for manual email approval of ACM certificates.

        ```python
        import pulumi
        import pulumi_aws as aws

        example_certificate = aws.acm.Certificate("exampleCertificate",
            domain_name="example.com",
            validation_method="EMAIL")
        example_certificate_validation = aws.acm.CertificateValidation("exampleCertificateValidation", certificate_arn=example_certificate.arn)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate that is being validated.
        :param pulumi.Input[List[pulumi.Input[str]]] validation_record_fqdns: List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
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
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            certificate_arn: Optional[pulumi.Input[str]] = None,
            validation_record_fqdns: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'CertificateValidation':
        """
        Get an existing CertificateValidation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_arn: The ARN of the certificate that is being validated.
        :param pulumi.Input[List[pulumi.Input[str]]] validation_record_fqdns: List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["certificate_arn"] = certificate_arn
        __props__["validation_record_fqdns"] = validation_record_fqdns
        return CertificateValidation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="certificateArn")
    def certificate_arn(self) -> str:
        """
        The ARN of the certificate that is being validated.
        """
        return pulumi.get(self, "certificate_arn")

    @property
    @pulumi.getter(name="validationRecordFqdns")
    def validation_record_fqdns(self) -> Optional[List[str]]:
        """
        List of FQDNs that implement the validation. Only valid for DNS validation method ACM certificates. If this is set, the resource can implement additional sanity checks and has an explicit dependency on the resource that is implementing the validation
        """
        return pulumi.get(self, "validation_record_fqdns")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

