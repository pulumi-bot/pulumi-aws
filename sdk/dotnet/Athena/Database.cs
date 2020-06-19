// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Athena
{
    /// <summary>
    /// Provides an Athena database.
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
    ///         var hogeBucket = new Aws.S3.Bucket("hogeBucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var hogeDatabase = new Aws.Athena.Database("hogeDatabase", new Aws.Athena.DatabaseArgs
    ///         {
    ///             Bucket = hogeBucket.BucketName,
    ///             Name = "database_name",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of s3 bucket to save the results of the query execution.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
        /// </summary>
        [Output("encryptionConfiguration")]
        public Output<Outputs.DatabaseEncryptionConfiguration?> EncryptionConfiguration { get; private set; } = null!;

        /// <summary>
        /// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// Name of the database to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("aws:athena/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("aws:athena/database:Database", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of s3 bucket to save the results of the query execution.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.DatabaseEncryptionConfigurationArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the database to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DatabaseArgs()
        {
        }
    }

    public sealed class DatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of s3 bucket to save the results of the query execution.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// The encryption key block AWS Athena uses to decrypt the data in S3, such as an AWS Key Management Service (AWS KMS) key. An `encryption_configuration` block is documented below.
        /// </summary>
        [Input("encryptionConfiguration")]
        public Input<Inputs.DatabaseEncryptionConfigurationGetArgs>? EncryptionConfiguration { get; set; }

        /// <summary>
        /// A boolean that indicates all tables should be deleted from the database so that the database can be destroyed without error. The tables are *not* recoverable.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        /// <summary>
        /// Name of the database to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DatabaseState()
        {
        }
    }
}
