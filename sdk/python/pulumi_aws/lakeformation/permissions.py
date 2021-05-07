# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['PermissionsArgs', 'Permissions']

@pulumi.input_type
class PermissionsArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input[Sequence[pulumi.Input[str]]],
                 principal: pulumi.Input[str],
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 catalog_resource: Optional[pulumi.Input[bool]] = None,
                 data_location: Optional[pulumi.Input['PermissionsDataLocationArgs']] = None,
                 database: Optional[pulumi.Input['PermissionsDatabaseArgs']] = None,
                 permissions_with_grant_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 table: Optional[pulumi.Input['PermissionsTableArgs']] = None,
                 table_with_columns: Optional[pulumi.Input['PermissionsTableWithColumnsArgs']] = None):
        """
        The set of arguments for constructing a Permissions resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[str] principal: Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param pulumi.Input[bool] catalog_resource: Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        :param pulumi.Input['PermissionsDataLocationArgs'] data_location: Configuration block for a data location resource. Detailed below.
        :param pulumi.Input['PermissionsDatabaseArgs'] database: Configuration block for a database resource. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions_with_grant_options: Subset of `permissions` which the principal can pass.
        :param pulumi.Input['PermissionsTableArgs'] table: Configuration block for a table resource. Detailed below.
        :param pulumi.Input['PermissionsTableWithColumnsArgs'] table_with_columns: Configuration block for a table with columns resource. Detailed below.
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "principal", principal)
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if catalog_resource is not None:
            pulumi.set(__self__, "catalog_resource", catalog_resource)
        if data_location is not None:
            pulumi.set(__self__, "data_location", data_location)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if permissions_with_grant_options is not None:
            pulumi.set(__self__, "permissions_with_grant_options", permissions_with_grant_options)
        if table is not None:
            pulumi.set(__self__, "table", table)
        if table_with_columns is not None:
            pulumi.set(__self__, "table_with_columns", table_with_columns)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Input[str]:
        """
        Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="catalogResource")
    def catalog_resource(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        """
        return pulumi.get(self, "catalog_resource")

    @catalog_resource.setter
    def catalog_resource(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "catalog_resource", value)

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> Optional[pulumi.Input['PermissionsDataLocationArgs']]:
        """
        Configuration block for a data location resource. Detailed below.
        """
        return pulumi.get(self, "data_location")

    @data_location.setter
    def data_location(self, value: Optional[pulumi.Input['PermissionsDataLocationArgs']]):
        pulumi.set(self, "data_location", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input['PermissionsDatabaseArgs']]:
        """
        Configuration block for a database resource. Detailed below.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input['PermissionsDatabaseArgs']]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter(name="permissionsWithGrantOptions")
    def permissions_with_grant_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Subset of `permissions` which the principal can pass.
        """
        return pulumi.get(self, "permissions_with_grant_options")

    @permissions_with_grant_options.setter
    def permissions_with_grant_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions_with_grant_options", value)

    @property
    @pulumi.getter
    def table(self) -> Optional[pulumi.Input['PermissionsTableArgs']]:
        """
        Configuration block for a table resource. Detailed below.
        """
        return pulumi.get(self, "table")

    @table.setter
    def table(self, value: Optional[pulumi.Input['PermissionsTableArgs']]):
        pulumi.set(self, "table", value)

    @property
    @pulumi.getter(name="tableWithColumns")
    def table_with_columns(self) -> Optional[pulumi.Input['PermissionsTableWithColumnsArgs']]:
        """
        Configuration block for a table with columns resource. Detailed below.
        """
        return pulumi.get(self, "table_with_columns")

    @table_with_columns.setter
    def table_with_columns(self, value: Optional[pulumi.Input['PermissionsTableWithColumnsArgs']]):
        pulumi.set(self, "table_with_columns", value)


@pulumi.input_type
class _PermissionsState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 catalog_resource: Optional[pulumi.Input[bool]] = None,
                 data_location: Optional[pulumi.Input['PermissionsDataLocationArgs']] = None,
                 database: Optional[pulumi.Input['PermissionsDatabaseArgs']] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 permissions_with_grant_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input['PermissionsTableArgs']] = None,
                 table_with_columns: Optional[pulumi.Input['PermissionsTableWithColumnsArgs']] = None):
        """
        Input properties used for looking up and filtering Permissions resources.
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param pulumi.Input[bool] catalog_resource: Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        :param pulumi.Input['PermissionsDataLocationArgs'] data_location: Configuration block for a data location resource. Detailed below.
        :param pulumi.Input['PermissionsDatabaseArgs'] database: Configuration block for a database resource. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions_with_grant_options: Subset of `permissions` which the principal can pass.
        :param pulumi.Input[str] principal: Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input['PermissionsTableArgs'] table: Configuration block for a table resource. Detailed below.
        :param pulumi.Input['PermissionsTableWithColumnsArgs'] table_with_columns: Configuration block for a table with columns resource. Detailed below.
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if catalog_resource is not None:
            pulumi.set(__self__, "catalog_resource", catalog_resource)
        if data_location is not None:
            pulumi.set(__self__, "data_location", data_location)
        if database is not None:
            pulumi.set(__self__, "database", database)
        if permissions is not None:
            pulumi.set(__self__, "permissions", permissions)
        if permissions_with_grant_options is not None:
            pulumi.set(__self__, "permissions_with_grant_options", permissions_with_grant_options)
        if principal is not None:
            pulumi.set(__self__, "principal", principal)
        if table is not None:
            pulumi.set(__self__, "table", table)
        if table_with_columns is not None:
            pulumi.set(__self__, "table_with_columns", table_with_columns)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="catalogResource")
    def catalog_resource(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        """
        return pulumi.get(self, "catalog_resource")

    @catalog_resource.setter
    def catalog_resource(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "catalog_resource", value)

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> Optional[pulumi.Input['PermissionsDataLocationArgs']]:
        """
        Configuration block for a data location resource. Detailed below.
        """
        return pulumi.get(self, "data_location")

    @data_location.setter
    def data_location(self, value: Optional[pulumi.Input['PermissionsDataLocationArgs']]):
        pulumi.set(self, "data_location", value)

    @property
    @pulumi.getter
    def database(self) -> Optional[pulumi.Input['PermissionsDatabaseArgs']]:
        """
        Configuration block for a database resource. Detailed below.
        """
        return pulumi.get(self, "database")

    @database.setter
    def database(self, value: Optional[pulumi.Input['PermissionsDatabaseArgs']]):
        pulumi.set(self, "database", value)

    @property
    @pulumi.getter
    def permissions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="permissionsWithGrantOptions")
    def permissions_with_grant_options(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Subset of `permissions` which the principal can pass.
        """
        return pulumi.get(self, "permissions_with_grant_options")

    @permissions_with_grant_options.setter
    def permissions_with_grant_options(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "permissions_with_grant_options", value)

    @property
    @pulumi.getter
    def principal(self) -> Optional[pulumi.Input[str]]:
        """
        Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "principal")

    @principal.setter
    def principal(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal", value)

    @property
    @pulumi.getter
    def table(self) -> Optional[pulumi.Input['PermissionsTableArgs']]:
        """
        Configuration block for a table resource. Detailed below.
        """
        return pulumi.get(self, "table")

    @table.setter
    def table(self, value: Optional[pulumi.Input['PermissionsTableArgs']]):
        pulumi.set(self, "table", value)

    @property
    @pulumi.getter(name="tableWithColumns")
    def table_with_columns(self) -> Optional[pulumi.Input['PermissionsTableWithColumnsArgs']]:
        """
        Configuration block for a table with columns resource. Detailed below.
        """
        return pulumi.get(self, "table_with_columns")

    @table_with_columns.setter
    def table_with_columns(self, value: Optional[pulumi.Input['PermissionsTableWithColumnsArgs']]):
        pulumi.set(self, "table_with_columns", value)


class Permissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 catalog_resource: Optional[pulumi.Input[bool]] = None,
                 data_location: Optional[pulumi.Input[pulumi.InputType['PermissionsDataLocationArgs']]] = None,
                 database: Optional[pulumi.Input[pulumi.InputType['PermissionsDatabaseArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 permissions_with_grant_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[pulumi.InputType['PermissionsTableArgs']]] = None,
                 table_with_columns: Optional[pulumi.Input[pulumi.InputType['PermissionsTableWithColumnsArgs']]] = None,
                 __props__=None):
        """
        Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).

        > **NOTE:** Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. These implicit permissions cannot be revoked _per se_. If this resource reads implicit permissions, it will attempt to revoke them, which causes an error when the resource is destroyed. There are two ways to avoid these errors. First, grant explicit permissions (and `permissions_with_grant_option`) to "overwrite" a principal's implicit permissions, which you can then revoke with this resource. Second, avoid using this resource with principals that have implicit permissions. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).

        ## Example Usage
        ### Grant Permissions For A Lake Formation S3 Resource

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.lakeformation.Permissions("test",
            principal=aws_iam_role["workflow_role"]["arn"],
            permissions=["ALL"],
            data_location=aws.lakeformation.PermissionsDataLocationArgs(
                arn=aws_lakeformation_resource["test"]["arn"],
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param pulumi.Input[bool] catalog_resource: Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['PermissionsDataLocationArgs']] data_location: Configuration block for a data location resource. Detailed below.
        :param pulumi.Input[pulumi.InputType['PermissionsDatabaseArgs']] database: Configuration block for a database resource. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions_with_grant_options: Subset of `permissions` which the principal can pass.
        :param pulumi.Input[str] principal: Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[pulumi.InputType['PermissionsTableArgs']] table: Configuration block for a table resource. Detailed below.
        :param pulumi.Input[pulumi.InputType['PermissionsTableWithColumnsArgs']] table_with_columns: Configuration block for a table with columns resource. Detailed below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Grants permissions to the principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).

        > **NOTE:** Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. These implicit permissions cannot be revoked _per se_. If this resource reads implicit permissions, it will attempt to revoke them, which causes an error when the resource is destroyed. There are two ways to avoid these errors. First, grant explicit permissions (and `permissions_with_grant_option`) to "overwrite" a principal's implicit permissions, which you can then revoke with this resource. Second, avoid using this resource with principals that have implicit permissions. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).

        ## Example Usage
        ### Grant Permissions For A Lake Formation S3 Resource

        ```python
        import pulumi
        import pulumi_aws as aws

        test = aws.lakeformation.Permissions("test",
            principal=aws_iam_role["workflow_role"]["arn"],
            permissions=["ALL"],
            data_location=aws.lakeformation.PermissionsDataLocationArgs(
                arn=aws_lakeformation_resource["test"]["arn"],
            ))
        ```

        :param str resource_name: The name of the resource.
        :param PermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 catalog_resource: Optional[pulumi.Input[bool]] = None,
                 data_location: Optional[pulumi.Input[pulumi.InputType['PermissionsDataLocationArgs']]] = None,
                 database: Optional[pulumi.Input[pulumi.InputType['PermissionsDatabaseArgs']]] = None,
                 permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 permissions_with_grant_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 principal: Optional[pulumi.Input[str]] = None,
                 table: Optional[pulumi.Input[pulumi.InputType['PermissionsTableArgs']]] = None,
                 table_with_columns: Optional[pulumi.Input[pulumi.InputType['PermissionsTableWithColumnsArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PermissionsArgs.__new__(PermissionsArgs)

            __props__.__dict__["catalog_id"] = catalog_id
            __props__.__dict__["catalog_resource"] = catalog_resource
            __props__.__dict__["data_location"] = data_location
            __props__.__dict__["database"] = database
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["permissions_with_grant_options"] = permissions_with_grant_options
            if principal is None and not opts.urn:
                raise TypeError("Missing required property 'principal'")
            __props__.__dict__["principal"] = principal
            __props__.__dict__["table"] = table
            __props__.__dict__["table_with_columns"] = table_with_columns
        super(Permissions, __self__).__init__(
            'aws:lakeformation/permissions:Permissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            catalog_resource: Optional[pulumi.Input[bool]] = None,
            data_location: Optional[pulumi.Input[pulumi.InputType['PermissionsDataLocationArgs']]] = None,
            database: Optional[pulumi.Input[pulumi.InputType['PermissionsDatabaseArgs']]] = None,
            permissions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            permissions_with_grant_options: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            principal: Optional[pulumi.Input[str]] = None,
            table: Optional[pulumi.Input[pulumi.InputType['PermissionsTableArgs']]] = None,
            table_with_columns: Optional[pulumi.Input[pulumi.InputType['PermissionsTableWithColumnsArgs']]] = None) -> 'Permissions':
        """
        Get an existing Permissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: Identifier for the Data Catalog. By default, it is the account ID of the caller.
        :param pulumi.Input[bool] catalog_resource: Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        :param pulumi.Input[pulumi.InputType['PermissionsDataLocationArgs']] data_location: Configuration block for a data location resource. Detailed below.
        :param pulumi.Input[pulumi.InputType['PermissionsDatabaseArgs']] database: Configuration block for a database resource. Detailed below.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions: List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permissions_with_grant_options: Subset of `permissions` which the principal can pass.
        :param pulumi.Input[str] principal: Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        :param pulumi.Input[pulumi.InputType['PermissionsTableArgs']] table: Configuration block for a table resource. Detailed below.
        :param pulumi.Input[pulumi.InputType['PermissionsTableWithColumnsArgs']] table_with_columns: Configuration block for a table with columns resource. Detailed below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PermissionsState.__new__(_PermissionsState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["catalog_resource"] = catalog_resource
        __props__.__dict__["data_location"] = data_location
        __props__.__dict__["database"] = database
        __props__.__dict__["permissions"] = permissions
        __props__.__dict__["permissions_with_grant_options"] = permissions_with_grant_options
        __props__.__dict__["principal"] = principal
        __props__.__dict__["table"] = table
        __props__.__dict__["table_with_columns"] = table_with_columns
        return Permissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[Optional[str]]:
        """
        Identifier for the Data Catalog. By default, it is the account ID of the caller.
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="catalogResource")
    def catalog_resource(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        """
        return pulumi.get(self, "catalog_resource")

    @property
    @pulumi.getter(name="dataLocation")
    def data_location(self) -> pulumi.Output['outputs.PermissionsDataLocation']:
        """
        Configuration block for a data location resource. Detailed below.
        """
        return pulumi.get(self, "data_location")

    @property
    @pulumi.getter
    def database(self) -> pulumi.Output['outputs.PermissionsDatabase']:
        """
        Configuration block for a database resource. Detailed below.
        """
        return pulumi.get(self, "database")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output[Sequence[str]]:
        """
        List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="permissionsWithGrantOptions")
    def permissions_with_grant_options(self) -> pulumi.Output[Sequence[str]]:
        """
        Subset of `permissions` which the principal can pass.
        """
        return pulumi.get(self, "permissions_with_grant_options")

    @property
    @pulumi.getter
    def principal(self) -> pulumi.Output[str]:
        """
        Principal to be granted the permissions on the resource. Supported principals include IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        """
        return pulumi.get(self, "principal")

    @property
    @pulumi.getter
    def table(self) -> pulumi.Output['outputs.PermissionsTable']:
        """
        Configuration block for a table resource. Detailed below.
        """
        return pulumi.get(self, "table")

    @property
    @pulumi.getter(name="tableWithColumns")
    def table_with_columns(self) -> pulumi.Output['outputs.PermissionsTableWithColumns']:
        """
        Configuration block for a table with columns resource. Detailed below.
        """
        return pulumi.get(self, "table_with_columns")

