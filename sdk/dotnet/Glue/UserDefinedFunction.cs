// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Glue
{
    /// <summary>
    /// Provides a Glue User Defined Function Resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleCatalogDatabase = new Aws.Glue.CatalogDatabase("exampleCatalogDatabase", new Aws.Glue.CatalogDatabaseArgs
    ///         {
    ///             Name = "my_database",
    ///         });
    ///         var exampleUserDefinedFunction = new Aws.Glue.UserDefinedFunction("exampleUserDefinedFunction", new Aws.Glue.UserDefinedFunctionArgs
    ///         {
    ///             CatalogId = exampleCatalogDatabase.CatalogId,
    ///             DatabaseName = exampleCatalogDatabase.Name,
    ///             ClassName = "class",
    ///             OwnerName = "owner",
    ///             OwnerType = "GROUP",
    ///             ResourceUris = 
    ///             {
    ///                 new Aws.Glue.Inputs.UserDefinedFunctionResourceUriArgs
    ///                 {
    ///                     ResourceType = "ARCHIVE",
    ///                     Uri = "uri",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Glue User Defined Functions can be imported using the `catalog_id:database_name:function_name`. If you have not set a Catalog ID specify the AWS Account ID that the database is in, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:glue/userDefinedFunction:UserDefinedFunction func 123456789012:my_database:my_func
    /// ```
    /// </summary>
    public partial class UserDefinedFunction : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
        /// </summary>
        [Output("catalogId")]
        public Output<string?> CatalogId { get; private set; } = null!;

        /// <summary>
        /// The Java class that contains the function code.
        /// </summary>
        [Output("className")]
        public Output<string> ClassName { get; private set; } = null!;

        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The name of the Database to create the Function.
        /// </summary>
        [Output("databaseName")]
        public Output<string> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The owner of the function.
        /// </summary>
        [Output("ownerName")]
        public Output<string> OwnerName { get; private set; } = null!;

        /// <summary>
        /// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
        /// </summary>
        [Output("ownerType")]
        public Output<string> OwnerType { get; private set; } = null!;

        /// <summary>
        /// The configuration block for Resource URIs. See resource uris below for more details.
        /// </summary>
        [Output("resourceUris")]
        public Output<ImmutableArray<Outputs.UserDefinedFunctionResourceUri>> ResourceUris { get; private set; } = null!;


        /// <summary>
        /// Create a UserDefinedFunction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserDefinedFunction(string name, UserDefinedFunctionArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/userDefinedFunction:UserDefinedFunction", name, args ?? new UserDefinedFunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserDefinedFunction(string name, Input<string> id, UserDefinedFunctionState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/userDefinedFunction:UserDefinedFunction", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing UserDefinedFunction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserDefinedFunction Get(string name, Input<string> id, UserDefinedFunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new UserDefinedFunction(name, id, state, options);
        }
    }

    public sealed class UserDefinedFunctionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The Java class that contains the function code.
        /// </summary>
        [Input("className", required: true)]
        public Input<string> ClassName { get; set; } = null!;

        /// <summary>
        /// The name of the Database to create the Function.
        /// </summary>
        [Input("databaseName", required: true)]
        public Input<string> DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the function.
        /// </summary>
        [Input("ownerName", required: true)]
        public Input<string> OwnerName { get; set; } = null!;

        /// <summary>
        /// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
        /// </summary>
        [Input("ownerType", required: true)]
        public Input<string> OwnerType { get; set; } = null!;

        [Input("resourceUris")]
        private InputList<Inputs.UserDefinedFunctionResourceUriArgs>? _resourceUris;

        /// <summary>
        /// The configuration block for Resource URIs. See resource uris below for more details.
        /// </summary>
        public InputList<Inputs.UserDefinedFunctionResourceUriArgs> ResourceUris
        {
            get => _resourceUris ?? (_resourceUris = new InputList<Inputs.UserDefinedFunctionResourceUriArgs>());
            set => _resourceUris = value;
        }

        public UserDefinedFunctionArgs()
        {
        }
    }

    public sealed class UserDefinedFunctionState : Pulumi.ResourceArgs
    {
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The Java class that contains the function code.
        /// </summary>
        [Input("className")]
        public Input<string>? ClassName { get; set; }

        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The name of the Database to create the Function.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// The name of the function.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The owner of the function.
        /// </summary>
        [Input("ownerName")]
        public Input<string>? OwnerName { get; set; }

        /// <summary>
        /// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
        /// </summary>
        [Input("ownerType")]
        public Input<string>? OwnerType { get; set; }

        [Input("resourceUris")]
        private InputList<Inputs.UserDefinedFunctionResourceUriGetArgs>? _resourceUris;

        /// <summary>
        /// The configuration block for Resource URIs. See resource uris below for more details.
        /// </summary>
        public InputList<Inputs.UserDefinedFunctionResourceUriGetArgs> ResourceUris
        {
            get => _resourceUris ?? (_resourceUris = new InputList<Inputs.UserDefinedFunctionResourceUriGetArgs>());
            set => _resourceUris = value;
        }

        public UserDefinedFunctionState()
        {
        }
    }
}
