# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['DataCatalogEncryptionSettings']


class DataCatalogEncryptionSettings(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 data_catalog_encryption_settings: Optional[pulumi.Input[pulumi.InputType['DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a Glue Data Catalog Encryption Settings resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        :param pulumi.Input[pulumi.InputType['DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs']] data_catalog_encryption_settings: The security configuration to set. see Data Catalog Encryption Settings.
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

            __props__['catalog_id'] = catalog_id
            if data_catalog_encryption_settings is None:
                raise TypeError("Missing required property 'data_catalog_encryption_settings'")
            __props__['data_catalog_encryption_settings'] = data_catalog_encryption_settings
        super(DataCatalogEncryptionSettings, __self__).__init__(
            'aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            data_catalog_encryption_settings: Optional[pulumi.Input[pulumi.InputType['DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs']]] = None) -> 'DataCatalogEncryptionSettings':
        """
        Get an existing DataCatalogEncryptionSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        :param pulumi.Input[pulumi.InputType['DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs']] data_catalog_encryption_settings: The security configuration to set. see Data Catalog Encryption Settings.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["catalog_id"] = catalog_id
        __props__["data_catalog_encryption_settings"] = data_catalog_encryption_settings
        return DataCatalogEncryptionSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="dataCatalogEncryptionSettings")
    def data_catalog_encryption_settings(self) -> pulumi.Output['outputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettings']:
        """
        The security configuration to set. see Data Catalog Encryption Settings.
        """
        return pulumi.get(self, "data_catalog_encryption_settings")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

