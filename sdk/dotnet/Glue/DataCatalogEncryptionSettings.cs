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
    /// Provides a Glue Data Catalog Encryption Settings resource.
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
    ///         var example = new Aws.Glue.DataCatalogEncryptionSettings("example", new Aws.Glue.DataCatalogEncryptionSettingsArgs
    ///         {
    ///             DataCatalogEncryptionSettings = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs
    ///             {
    ///                 ConnectionPasswordEncryption = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsConnectionPasswordEncryptionArgs
    ///                 {
    ///                     AwsKmsKeyId = aws_kms_key.Test.Arn,
    ///                     ReturnConnectionPasswordEncrypted = true,
    ///                 },
    ///                 EncryptionAtRest = new Aws.Glue.Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsEncryptionAtRestArgs
    ///                 {
    ///                     CatalogEncryptionMode = "SSE-KMS",
    ///                     SseAwsKmsKeyId = aws_kms_key.Test.Arn,
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
    /// Glue Data Catalog Encryption Settings can be imported using `CATALOG-ID` (AWS account ID if not custom), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings example 123456789012
    /// ```
    /// </summary>
    public partial class DataCatalogEncryptionSettings : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        /// </summary>
        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// The security configuration to set. see Data Catalog Encryption Settings.
        /// </summary>
        [Output("dataCatalogEncryptionSettings")]
        public Output<Outputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettings> DataCatalogEncryptionSettingsConfig { get; private set; } = null!;


        /// <summary>
        /// Create a DataCatalogEncryptionSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataCatalogEncryptionSettings(string name, DataCatalogEncryptionSettingsArgs args, CustomResourceOptions? options = null)
            : base("aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings", name, args ?? new DataCatalogEncryptionSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataCatalogEncryptionSettings(string name, Input<string> id, DataCatalogEncryptionSettingsState? state = null, CustomResourceOptions? options = null)
            : base("aws:glue/dataCatalogEncryptionSettings:DataCatalogEncryptionSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataCatalogEncryptionSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataCatalogEncryptionSettings Get(string name, Input<string> id, DataCatalogEncryptionSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new DataCatalogEncryptionSettings(name, id, state, options);
        }
    }

    public sealed class DataCatalogEncryptionSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The security configuration to set. see Data Catalog Encryption Settings.
        /// </summary>
        [Input("dataCatalogEncryptionSettings", required: true)]
        public Input<Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsArgs> DataCatalogEncryptionSettingsConfig { get; set; } = null!;

        public DataCatalogEncryptionSettingsArgs()
        {
        }
    }

    public sealed class DataCatalogEncryptionSettingsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// The security configuration to set. see Data Catalog Encryption Settings.
        /// </summary>
        [Input("dataCatalogEncryptionSettings")]
        public Input<Inputs.DataCatalogEncryptionSettingsDataCatalogEncryptionSettingsGetArgs>? DataCatalogEncryptionSettingsConfig { get; set; }

        public DataCatalogEncryptionSettingsState()
        {
        }
    }
}
