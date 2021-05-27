# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetDataCatalogEncryptionSettingsResult',
    'AwaitableGetDataCatalogEncryptionSettingsResult',
    'get_data_catalog_encryption_settings',
]

@pulumi.output_type
class GetDataCatalogEncryptionSettingsResult:
    """
    A collection of values returned by getDataCatalogEncryptionSettings.
    """
    def __init__(__self__, catalog_id=None, data_catalog_encryption_settings=None, id=None):
        if catalog_id and not isinstance(catalog_id, str):
            raise TypeError("Expected argument 'catalog_id' to be a str")
        pulumi.set(__self__, "catalog_id", catalog_id)
        if data_catalog_encryption_settings and not isinstance(data_catalog_encryption_settings, list):
            raise TypeError("Expected argument 'data_catalog_encryption_settings' to be a list")
        pulumi.set(__self__, "data_catalog_encryption_settings", data_catalog_encryption_settings)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> str:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="dataCatalogEncryptionSettings")
    def data_catalog_encryption_settings(self) -> Sequence['outputs.GetDataCatalogEncryptionSettingsDataCatalogEncryptionSettingResult']:
        """
        The security configuration to set. see Data Catalog Encryption Settings.
        """
        return pulumi.get(self, "data_catalog_encryption_settings")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetDataCatalogEncryptionSettingsResult(GetDataCatalogEncryptionSettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataCatalogEncryptionSettingsResult(
            catalog_id=self.catalog_id,
            data_catalog_encryption_settings=self.data_catalog_encryption_settings,
            id=self.id)


def get_data_catalog_encryption_settings(catalog_id: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataCatalogEncryptionSettingsResult:
    """
    This data source can be used to fetch information about AWS Glue Data Catalog Encryption Settings.


    :param str catalog_id: The ID of the Data Catalog. This is typically the AWS account ID.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:glue/getDataCatalogEncryptionSettings:getDataCatalogEncryptionSettings', __args__, opts=opts, typ=GetDataCatalogEncryptionSettingsResult).value

    return AwaitableGetDataCatalogEncryptionSettingsResult(
        catalog_id=__ret__.catalog_id,
        data_catalog_encryption_settings=__ret__.data_catalog_encryption_settings,
        id=__ret__.id)


@_utilities.lift_output_func(get_data_catalog_encryption_settings)
def get_data_catalog_encryption_settings_apply(catalog_id: Optional[pulumi.Input[str]] = None,
                                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataCatalogEncryptionSettingsResult]:
    ...
