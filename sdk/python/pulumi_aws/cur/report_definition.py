# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ReportDefinition']


class ReportDefinition(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_artifacts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 additional_schema_elements: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 compression: Optional[pulumi.Input[str]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 report_name: Optional[pulumi.Input[str]] = None,
                 s3_bucket: Optional[pulumi.Input[str]] = None,
                 s3_prefix: Optional[pulumi.Input[str]] = None,
                 s3_region: Optional[pulumi.Input[str]] = None,
                 time_unit: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ReportDefinition resource with the given unique name, props, and options.
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

            __props__['additional_artifacts'] = additional_artifacts
            if additional_schema_elements is None:
                raise TypeError("Missing required property 'additional_schema_elements'")
            __props__['additional_schema_elements'] = additional_schema_elements
            if compression is None:
                raise TypeError("Missing required property 'compression'")
            __props__['compression'] = compression
            if format is None:
                raise TypeError("Missing required property 'format'")
            __props__['format'] = format
            if report_name is None:
                raise TypeError("Missing required property 'report_name'")
            __props__['report_name'] = report_name
            if s3_bucket is None:
                raise TypeError("Missing required property 's3_bucket'")
            __props__['s3_bucket'] = s3_bucket
            __props__['s3_prefix'] = s3_prefix
            if s3_region is None:
                raise TypeError("Missing required property 's3_region'")
            __props__['s3_region'] = s3_region
            if time_unit is None:
                raise TypeError("Missing required property 'time_unit'")
            __props__['time_unit'] = time_unit
        super(ReportDefinition, __self__).__init__(
            'aws:cur/reportDefinition:ReportDefinition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_artifacts: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            additional_schema_elements: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            compression: Optional[pulumi.Input[str]] = None,
            format: Optional[pulumi.Input[str]] = None,
            report_name: Optional[pulumi.Input[str]] = None,
            s3_bucket: Optional[pulumi.Input[str]] = None,
            s3_prefix: Optional[pulumi.Input[str]] = None,
            s3_region: Optional[pulumi.Input[str]] = None,
            time_unit: Optional[pulumi.Input[str]] = None) -> 'ReportDefinition':
        """
        Get an existing ReportDefinition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["additional_artifacts"] = additional_artifacts
        __props__["additional_schema_elements"] = additional_schema_elements
        __props__["compression"] = compression
        __props__["format"] = format
        __props__["report_name"] = report_name
        __props__["s3_bucket"] = s3_bucket
        __props__["s3_prefix"] = s3_prefix
        __props__["s3_region"] = s3_region
        __props__["time_unit"] = time_unit
        return ReportDefinition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalArtifacts")
    def additional_artifacts(self) -> pulumi.Output[Optional[List[str]]]:
        return pulumi.get(self, "additional_artifacts")

    @property
    @pulumi.getter(name="additionalSchemaElements")
    def additional_schema_elements(self) -> pulumi.Output[List[str]]:
        return pulumi.get(self, "additional_schema_elements")

    @property
    @pulumi.getter
    def compression(self) -> pulumi.Output[str]:
        return pulumi.get(self, "compression")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="reportName")
    def report_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "report_name")

    @property
    @pulumi.getter(name="s3Bucket")
    def s3_bucket(self) -> pulumi.Output[str]:
        return pulumi.get(self, "s3_bucket")

    @property
    @pulumi.getter(name="s3Prefix")
    def s3_prefix(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "s3_prefix")

    @property
    @pulumi.getter(name="s3Region")
    def s3_region(self) -> pulumi.Output[str]:
        return pulumi.get(self, "s3_region")

    @property
    @pulumi.getter(name="timeUnit")
    def time_unit(self) -> pulumi.Output[str]:
        return pulumi.get(self, "time_unit")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

