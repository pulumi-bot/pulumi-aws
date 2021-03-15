# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = [
    'DataLakeSettingsCreateDatabaseDefaultPermission',
    'DataLakeSettingsCreateTableDefaultPermission',
    'PermissionsDataLocation',
    'PermissionsDatabase',
    'PermissionsTable',
    'PermissionsTableWithColumns',
    'GetDataLakeSettingsCreateDatabaseDefaultPermissionResult',
    'GetDataLakeSettingsCreateTableDefaultPermissionResult',
    'GetPermissionsDataLocationResult',
    'GetPermissionsDatabaseResult',
    'GetPermissionsTableResult',
    'GetPermissionsTableWithColumnsResult',
]

@pulumi.output_type
class DataLakeSettingsCreateDatabaseDefaultPermission(dict):
    def __init__(__self__, *,
                 permissions: Optional[Sequence[str]] = None,
                 principal: Optional[str] = None):
        """
        :param Sequence[str] permissions: List of permissions that are granted to the principal. Valid values may include `ALL`, `SELECT`, `ALTER`, `DROP`, `DELETE`, `INSERT`, and `DESCRIBE`. For more details, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param str principal: Principal who is granted permissions. To enforce metadata and underlying data access control only by IAM on new databases and tables set `principal` to `IAM_ALLOWED_PRINCIPALS` and `permissions` to `["ALL"]`.
        """
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[Sequence[str]]:
        """
        List of permissions that are granted to the principal. Valid values may include `ALL`, `SELECT`, `ALTER`, `DROP`, `DELETE`, `INSERT`, and `DESCRIBE`. For more details, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def principal(self) -> Optional[str]:
        """
        Principal who is granted permissions. To enforce metadata and underlying data access control only by IAM on new databases and tables set `principal` to `IAM_ALLOWED_PRINCIPALS` and `permissions` to `["ALL"]`.
        """
        return pulumi.get(self, "principal")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class DataLakeSettingsCreateTableDefaultPermission(dict):
    def __init__(__self__, *,
                 permissions: Optional[Sequence[str]] = None,
                 principal: Optional[str] = None):
        """
        :param Sequence[str] permissions: List of permissions that are granted to the principal. Valid values may include `ALL`, `SELECT`, `ALTER`, `DROP`, `DELETE`, `INSERT`, and `DESCRIBE`. For more details, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param str principal: Principal who is granted permissions. To enforce metadata and underlying data access control only by IAM on new databases and tables set `principal` to `IAM_ALLOWED_PRINCIPALS` and `permissions` to `["ALL"]`.
        """
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[Sequence[str]]:
        """
        List of permissions that are granted to the principal. Valid values may include `ALL`, `SELECT`, `ALTER`, `DROP`, `DELETE`, `INSERT`, and `DESCRIBE`. For more details, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def principal(self) -> Optional[str]:
        """
        Principal who is granted permissions. To enforce metadata and underlying data access control only by IAM on new databases and tables set `principal` to `IAM_ALLOWED_PRINCIPALS` and `permissions` to `["ALL"]`.
        """
        return pulumi.get(self, "principal")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PermissionsDataLocation(dict):
    def __init__(__self__, *,
                 arn: str,
                 catalog_id: Optional[str] = None):
        """
        :param str arn: Amazon Resource Name (ARN) that uniquely identifies the data location resource.
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        pulumi.set(__self__, "arn", arn)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) that uniquely identifies the data location resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PermissionsDatabase(dict):
    def __init__(__self__, *,
                 name: str,
                 catalog_id: Optional[str] = None):
        """
        :param str name: Name of the table resource.
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        pulumi.set(__self__, "name", name)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PermissionsTable(dict):
    def __init__(__self__, *,
                 database_name: str,
                 catalog_id: Optional[str] = None,
                 name: Optional[str] = None,
                 wildcard: Optional[bool] = None):
        """
        :param str database_name: Name of the database for the table with columns resource. Unique to the Data Catalog.
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param str name: Name of the table resource.
        :param bool wildcard: Whether to use a wildcard representing every table under a database. Defaults to `false`.
        """
        pulumi.set(__self__, "database_name", database_name)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if wildcard is not None:
            pulumi.set(__self__, "wildcard", wildcard)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Name of the database for the table with columns resource. Unique to the Data Catalog.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def wildcard(self) -> Optional[bool]:
        """
        Whether to use a wildcard representing every table under a database. Defaults to `false`.
        """
        return pulumi.get(self, "wildcard")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PermissionsTableWithColumns(dict):
    def __init__(__self__, *,
                 database_name: str,
                 name: str,
                 catalog_id: Optional[str] = None,
                 column_names: Optional[Sequence[str]] = None,
                 excluded_column_names: Optional[Sequence[str]] = None):
        """
        :param str database_name: Name of the database for the table with columns resource. Unique to the Data Catalog.
        :param str name: Name of the table resource.
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param Sequence[str] column_names: List of column names for the table.
        :param Sequence[str] excluded_column_names: List of column names for the table to exclude.
        """
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "name", name)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if column_names is not None:
            pulumi.set(__self__, "column_names", column_names)
        if excluded_column_names is not None:
            pulumi.set(__self__, "excluded_column_names", excluded_column_names)

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Name of the database for the table with columns resource. Unique to the Data Catalog.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="columnNames")
    def column_names(self) -> Optional[Sequence[str]]:
        """
        List of column names for the table.
        """
        return pulumi.get(self, "column_names")

    @property
    @pulumi.getter(name="excludedColumnNames")
    def excluded_column_names(self) -> Optional[Sequence[str]]:
        """
        List of column names for the table to exclude.
        """
        return pulumi.get(self, "excluded_column_names")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetDataLakeSettingsCreateDatabaseDefaultPermissionResult(dict):
    def __init__(__self__, *,
                 permissions: Sequence[str],
                 principal: str):
        """
        :param Sequence[str] permissions: List of permissions granted to the principal.
        :param str principal: Principal who is granted permissions.
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "principal", principal)

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        List of permissions granted to the principal.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def principal(self) -> str:
        """
        Principal who is granted permissions.
        """
        return pulumi.get(self, "principal")


@pulumi.output_type
class GetDataLakeSettingsCreateTableDefaultPermissionResult(dict):
    def __init__(__self__, *,
                 permissions: Sequence[str],
                 principal: str):
        """
        :param Sequence[str] permissions: List of permissions granted to the principal.
        :param str principal: Principal who is granted permissions.
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "principal", principal)

    @property
    @pulumi.getter
    def permissions(self) -> Sequence[str]:
        """
        List of permissions granted to the principal.
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter
    def principal(self) -> str:
        """
        Principal who is granted permissions.
        """
        return pulumi.get(self, "principal")


@pulumi.output_type
class GetPermissionsDataLocationResult(dict):
    def __init__(__self__, *,
                 arn: str,
                 catalog_id: str):
        """
        :param str arn: Amazon Resource Name (ARN) that uniquely identifies the data location resource.
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        pulumi.set(__self__, "arn", arn)
        pulumi.set(__self__, "catalog_id", catalog_id)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) that uniquely identifies the data location resource.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> str:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")


@pulumi.output_type
class GetPermissionsDatabaseResult(dict):
    def __init__(__self__, *,
                 catalog_id: str,
                 name: str):
        """
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param str name: Name of the table resource.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> str:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")


@pulumi.output_type
class GetPermissionsTableResult(dict):
    def __init__(__self__, *,
                 catalog_id: str,
                 database_name: str,
                 name: str,
                 wildcard: Optional[bool] = None):
        """
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param str database_name: Name of the database for the table with columns resource. Unique to the Data Catalog.
        :param str name: Name of the table resource.
        :param bool wildcard: Whether to use a wildcard representing every table under a database. At least one of `name` or `wildcard` is required. Defaults to `false`.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "name", name)
        if wildcard is not None:
            pulumi.set(__self__, "wildcard", wildcard)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> str:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Name of the database for the table with columns resource. Unique to the Data Catalog.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def wildcard(self) -> Optional[bool]:
        """
        Whether to use a wildcard representing every table under a database. At least one of `name` or `wildcard` is required. Defaults to `false`.
        """
        return pulumi.get(self, "wildcard")


@pulumi.output_type
class GetPermissionsTableWithColumnsResult(dict):
    def __init__(__self__, *,
                 catalog_id: str,
                 database_name: str,
                 name: str,
                 column_names: Optional[Sequence[str]] = None,
                 excluded_column_names: Optional[Sequence[str]] = None):
        """
        :param str catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param str database_name: Name of the database for the table with columns resource. Unique to the Data Catalog.
        :param str name: Name of the table resource.
        :param Sequence[str] column_names: List of column names for the table. At least one of `column_names` or `excluded_column_names` is required.
        :param Sequence[str] excluded_column_names: List of column names for the table to exclude. At least one of `column_names` or `excluded_column_names` is required.
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "database_name", database_name)
        pulumi.set(__self__, "name", name)
        if column_names is not None:
            pulumi.set(__self__, "column_names", column_names)
        if excluded_column_names is not None:
            pulumi.set(__self__, "excluded_column_names", excluded_column_names)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> str:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="databaseName")
    def database_name(self) -> str:
        """
        Name of the database for the table with columns resource. Unique to the Data Catalog.
        """
        return pulumi.get(self, "database_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the table resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="columnNames")
    def column_names(self) -> Optional[Sequence[str]]:
        """
        List of column names for the table. At least one of `column_names` or `excluded_column_names` is required.
        """
        return pulumi.get(self, "column_names")

    @property
    @pulumi.getter(name="excludedColumnNames")
    def excluded_column_names(self) -> Optional[Sequence[str]]:
        """
        List of column names for the table to exclude. At least one of `column_names` or `excluded_column_names` is required.
        """
        return pulumi.get(self, "excluded_column_names")


