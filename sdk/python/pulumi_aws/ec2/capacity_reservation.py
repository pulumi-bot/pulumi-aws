# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class CapacityReservation(pulumi.CustomResource):
    availability_zone: pulumi.Output[str]
    """
    The Availability Zone in which to create the Capacity Reservation.
    """
    ebs_optimized: pulumi.Output[bool]
    """
    Indicates whether the Capacity Reservation supports EBS-optimized instances.
    """
    end_date: pulumi.Output[str]
    """
    The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
    """
    end_date_type: pulumi.Output[str]
    """
    Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
    """
    ephemeral_storage: pulumi.Output[bool]
    """
    Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
    """
    instance_count: pulumi.Output[float]
    """
    The number of instances for which to reserve capacity.
    """
    instance_match_criteria: pulumi.Output[str]
    """
    Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
    """
    instance_platform: pulumi.Output[str]
    """
    The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
    """
    instance_type: pulumi.Output[str]
    """
    The instance type for which to reserve capacity.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    tenancy: pulumi.Output[str]
    """
    Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
    """
    def __init__(__self__, resource_name, opts=None, availability_zone=None, ebs_optimized=None, end_date=None, end_date_type=None, ephemeral_storage=None, instance_count=None, instance_match_criteria=None, instance_platform=None, instance_type=None, tags=None, tenancy=None, __name__=None, __opts__=None):
        """
        Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: The Availability Zone in which to create the Capacity Reservation.
        :param pulumi.Input[bool] ebs_optimized: Indicates whether the Capacity Reservation supports EBS-optimized instances.
        :param pulumi.Input[str] end_date: The date and time at which the Capacity Reservation expires. When a Capacity Reservation expires, the reserved capacity is released and you can no longer launch instances into it. Valid values: [RFC3339 time string](https://tools.ietf.org/html/rfc3339#section-5.8) (`YYYY-MM-DDTHH:MM:SSZ`)
        :param pulumi.Input[str] end_date_type: Indicates the way in which the Capacity Reservation ends. Specify either `unlimited` or `limited`.
        :param pulumi.Input[bool] ephemeral_storage: Indicates whether the Capacity Reservation supports instances with temporary, block-level storage.
        :param pulumi.Input[float] instance_count: The number of instances for which to reserve capacity.
        :param pulumi.Input[str] instance_match_criteria: Indicates the type of instance launches that the Capacity Reservation accepts. Specify either `open` or `targeted`.
        :param pulumi.Input[str] instance_platform: The type of operating system for which to reserve capacity. Valid options are `Linux/UNIX`, `Red Hat Enterprise Linux`, `SUSE Linux`, `Windows`, `Windows with SQL Server`, `Windows with SQL Server Enterprise`, `Windows with SQL Server Standard` or `Windows with SQL Server Web`.
        :param pulumi.Input[str] instance_type: The instance type for which to reserve capacity.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[str] tenancy: Indicates the tenancy of the Capacity Reservation. Specify either `default` or `dedicated`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if availability_zone is None:
            raise TypeError("Missing required property 'availability_zone'")
        __props__['availability_zone'] = availability_zone

        __props__['ebs_optimized'] = ebs_optimized

        __props__['end_date'] = end_date

        __props__['end_date_type'] = end_date_type

        __props__['ephemeral_storage'] = ephemeral_storage

        if instance_count is None:
            raise TypeError("Missing required property 'instance_count'")
        __props__['instance_count'] = instance_count

        __props__['instance_match_criteria'] = instance_match_criteria

        if instance_platform is None:
            raise TypeError("Missing required property 'instance_platform'")
        __props__['instance_platform'] = instance_platform

        if instance_type is None:
            raise TypeError("Missing required property 'instance_type'")
        __props__['instance_type'] = instance_type

        __props__['tags'] = tags

        __props__['tenancy'] = tenancy

        super(CapacityReservation, __self__).__init__(
            'aws:ec2/capacityReservation:CapacityReservation',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

