// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LakeFormation
{
    public static class GetPermissions
    {
        /// <summary>
        /// Get permissions for a principal to access metadata in the Data Catalog and data organized in underlying data storage such as Amazon S3. Permissions are granted to a principal, in a Data Catalog, relative to a Lake Formation resource, which includes the Data Catalog, databases, and tables. For more information, see [Security and Access Control to Metadata and Data in Lake Formation](https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
        /// 
        /// &gt; **NOTE:** This data source deals with explicitly granted permissions. Lake Formation grants implicit permissions to data lake administrators, database creators, and table creators. For more information, see [Implicit Lake Formation Permissions](https://docs.aws.amazon.com/lake-formation/latest/dg/implicit-permissions.html).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// ### Permissions For A Glue Catalog Database
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var test = Output.Create(Aws.LakeFormation.GetPermissions.InvokeAsync(new Aws.LakeFormation.GetPermissionsArgs
        ///         {
        ///             Principal = aws_iam_role.Workflow_role.Arn,
        ///             Database = new Aws.LakeFormation.Inputs.GetPermissionsDatabaseArgs
        ///             {
        ///                 Name = aws_glue_catalog_database.Test.Name,
        ///                 CatalogId = "110376042874",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPermissionsResult> InvokeAsync(GetPermissionsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPermissionsResult>("aws:lakeformation/getPermissions:getPermissions", args ?? new GetPermissionsArgs(), options.WithVersion());

        public static Output<GetPermissionsResult> Apply(GetPermissionsApplyArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.CatalogId.Box(),
                args.CatalogResource.Box(),
                args.DataLocation.Box(),
                args.Database.Box(),
                args.Principal.Box(),
                args.Table.Box(),
                args.TableWithColumns.Box()
            ).Apply(a => {
                    var args = new GetPermissionsArgs();
                    a[0].Set(args, nameof(args.CatalogId));
                    a[1].Set(args, nameof(args.CatalogResource));
                    a[2].Set(args, nameof(args.DataLocation));
                    a[3].Set(args, nameof(args.Database));
                    a[4].Set(args, nameof(args.Principal));
                    a[5].Set(args, nameof(args.Table));
                    a[6].Set(args, nameof(args.TableWithColumns));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetPermissionsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Identifier for the Data Catalog. By default, it is the account ID of the caller.
        /// </summary>
        [Input("catalogId")]
        public string? CatalogId { get; set; }

        /// <summary>
        /// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        /// </summary>
        [Input("catalogResource")]
        public bool? CatalogResource { get; set; }

        /// <summary>
        /// Configuration block for a data location resource. Detailed below.
        /// </summary>
        [Input("dataLocation")]
        public Inputs.GetPermissionsDataLocationArgs? DataLocation { get; set; }

        /// <summary>
        /// Configuration block for a database resource. Detailed below.
        /// </summary>
        [Input("database")]
        public Inputs.GetPermissionsDatabaseArgs? Database { get; set; }

        /// <summary>
        /// Principal to be granted the permissions on the resource. Supported principals are IAM users or IAM roles.
        /// </summary>
        [Input("principal", required: true)]
        public string Principal { get; set; } = null!;

        /// <summary>
        /// Configuration block for a table resource. Detailed below.
        /// </summary>
        [Input("table")]
        public Inputs.GetPermissionsTableArgs? Table { get; set; }

        /// <summary>
        /// Configuration block for a table with columns resource. Detailed below.
        /// </summary>
        [Input("tableWithColumns")]
        public Inputs.GetPermissionsTableWithColumnsArgs? TableWithColumns { get; set; }

        public GetPermissionsArgs()
        {
        }
    }

    public sealed class GetPermissionsApplyArgs
    {
        /// <summary>
        /// Identifier for the Data Catalog. By default, it is the account ID of the caller.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        /// </summary>
        [Input("catalogResource")]
        public Input<bool>? CatalogResource { get; set; }

        /// <summary>
        /// Configuration block for a data location resource. Detailed below.
        /// </summary>
        [Input("dataLocation")]
        public Input<Inputs.GetPermissionsDataLocationArgs>? DataLocation { get; set; }

        /// <summary>
        /// Configuration block for a database resource. Detailed below.
        /// </summary>
        [Input("database")]
        public Input<Inputs.GetPermissionsDatabaseArgs>? Database { get; set; }

        /// <summary>
        /// Principal to be granted the permissions on the resource. Supported principals are IAM users or IAM roles.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// Configuration block for a table resource. Detailed below.
        /// </summary>
        [Input("table")]
        public Input<Inputs.GetPermissionsTableArgs>? Table { get; set; }

        /// <summary>
        /// Configuration block for a table with columns resource. Detailed below.
        /// </summary>
        [Input("tableWithColumns")]
        public Input<Inputs.GetPermissionsTableWithColumnsArgs>? TableWithColumns { get; set; }

        public GetPermissionsApplyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPermissionsResult
    {
        public readonly string? CatalogId;
        public readonly bool? CatalogResource;
        public readonly Outputs.GetPermissionsDataLocationResult DataLocation;
        public readonly Outputs.GetPermissionsDatabaseResult Database;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of permissions granted to the principal. For details on permissions, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        public readonly ImmutableArray<string> Permissions;
        /// <summary>
        /// Subset of `permissions` which the principal can pass.
        /// </summary>
        public readonly ImmutableArray<string> PermissionsWithGrantOptions;
        public readonly string Principal;
        public readonly Outputs.GetPermissionsTableResult Table;
        public readonly Outputs.GetPermissionsTableWithColumnsResult TableWithColumns;

        [OutputConstructor]
        private GetPermissionsResult(
            string? catalogId,

            bool? catalogResource,

            Outputs.GetPermissionsDataLocationResult dataLocation,

            Outputs.GetPermissionsDatabaseResult database,

            string id,

            ImmutableArray<string> permissions,

            ImmutableArray<string> permissionsWithGrantOptions,

            string principal,

            Outputs.GetPermissionsTableResult table,

            Outputs.GetPermissionsTableWithColumnsResult tableWithColumns)
        {
            CatalogId = catalogId;
            CatalogResource = catalogResource;
            DataLocation = dataLocation;
            Database = database;
            Id = id;
            Permissions = permissions;
            PermissionsWithGrantOptions = permissionsWithGrantOptions;
            Principal = principal;
            Table = table;
            TableWithColumns = tableWithColumns;
        }
    }
}
