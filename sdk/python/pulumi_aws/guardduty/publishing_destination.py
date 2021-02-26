# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['PublishingDestination']


class PublishingDestination(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_arn: Optional[pulumi.Input[str]] = None,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 detector_id: Optional[pulumi.Input[str]] = None,
                 kms_key_arn: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage a GuardDuty PublishingDestination. Requires an existing GuardDuty Detector.

        ## Import

        GuardDuty PublishingDestination can be imported using the the master GuardDuty detector ID and PublishingDestinationID, e.g.

        ```sh
         $ pulumi import aws:guardduty/publishingDestination:PublishingDestination test a4b86f26fa42e7e7cf0d1c333ea77777:a4b86f27a0e464e4a7e0516d242f1234
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        :param pulumi.Input[str] destination_type: Currently there is only "S3" available as destination type which is also the default value
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty.
        :param pulumi.Input[str] kms_key_arn: The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
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

            if destination_arn is None and not opts.urn:
                raise TypeError("Missing required property 'destination_arn'")
            __props__['destination_arn'] = destination_arn
            __props__['destination_type'] = destination_type
            if detector_id is None and not opts.urn:
                raise TypeError("Missing required property 'detector_id'")
            __props__['detector_id'] = detector_id
            if kms_key_arn is None and not opts.urn:
                raise TypeError("Missing required property 'kms_key_arn'")
            __props__['kms_key_arn'] = kms_key_arn
        super(PublishingDestination, __self__).__init__(
            'aws:guardduty/publishingDestination:PublishingDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_arn: Optional[pulumi.Input[str]] = None,
            destination_type: Optional[pulumi.Input[str]] = None,
            detector_id: Optional[pulumi.Input[str]] = None,
            kms_key_arn: Optional[pulumi.Input[str]] = None) -> 'PublishingDestination':
        """
        Get an existing PublishingDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_arn: The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        :param pulumi.Input[str] destination_type: Currently there is only "S3" available as destination type which is also the default value
        :param pulumi.Input[str] detector_id: The detector ID of the GuardDuty.
        :param pulumi.Input[str] kms_key_arn: The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["destination_arn"] = destination_arn
        __props__["destination_type"] = destination_type
        __props__["detector_id"] = detector_id
        __props__["kms_key_arn"] = kms_key_arn
        return PublishingDestination(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationArn")
    def destination_arn(self) -> pulumi.Output[str]:
        """
        The bucket arn and prefix under which the findings get exported. Bucket-ARN is required, the prefix is optional and will be `AWSLogs/[Account-ID]/GuardDuty/[Region]/` if not provided
        """
        return pulumi.get(self, "destination_arn")

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Output[Optional[str]]:
        """
        Currently there is only "S3" available as destination type which is also the default value
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter(name="detectorId")
    def detector_id(self) -> pulumi.Output[str]:
        """
        The detector ID of the GuardDuty.
        """
        return pulumi.get(self, "detector_id")

    @property
    @pulumi.getter(name="kmsKeyArn")
    def kms_key_arn(self) -> pulumi.Output[str]:
        """
        The ARN of the KMS key used to encrypt GuardDuty findings. GuardDuty enforces this to be encrypted.
        """
        return pulumi.get(self, "kms_key_arn")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

