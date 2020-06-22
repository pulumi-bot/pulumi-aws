// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGateway
{
    /// <summary>
    /// Provides a settings of an API Gateway Documentation Part.
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
    ///         var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///         });
    ///         var exampleDocumentationPart = new Aws.ApiGateway.DocumentationPart("exampleDocumentationPart", new Aws.ApiGateway.DocumentationPartArgs
    ///         {
    ///             Location = new Aws.ApiGateway.Inputs.DocumentationPartLocationArgs
    ///             {
    ///                 Method = "GET",
    ///                 Path = "/example",
    ///                 Type = "METHOD",
    ///             },
    ///             Properties = "{\"description\":\"Example description\"}",
    ///             RestApiId = exampleRestApi.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class DocumentationPart : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the targeted API entity of the to-be-created documentation part. See below.
        /// </summary>
        [Output("location")]
        public Output<Outputs.DocumentationPartLocation> Location { get; private set; } = null!;

        /// <summary>
        /// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        /// </summary>
        [Output("properties")]
        public Output<string> Properties { get; private set; } = null!;

        /// <summary>
        /// The ID of the associated Rest API
        /// </summary>
        [Output("restApiId")]
        public Output<string> RestApiId { get; private set; } = null!;


        /// <summary>
        /// Create a DocumentationPart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DocumentationPart(string name, DocumentationPartArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/documentationPart:DocumentationPart", name, args ?? new DocumentationPartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DocumentationPart(string name, Input<string> id, DocumentationPartState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/documentationPart:DocumentationPart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DocumentationPart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DocumentationPart Get(string name, Input<string> id, DocumentationPartState? state = null, CustomResourceOptions? options = null)
        {
            return new DocumentationPart(name, id, state, options);
        }
    }

    public sealed class DocumentationPartArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the targeted API entity of the to-be-created documentation part. See below.
        /// </summary>
        [Input("location", required: true)]
        public Input<Inputs.DocumentationPartLocationArgs> Location { get; set; } = null!;

        /// <summary>
        /// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        /// </summary>
        [Input("properties", required: true)]
        public Input<string> Properties { get; set; } = null!;

        /// <summary>
        /// The ID of the associated Rest API
        /// </summary>
        [Input("restApiId", required: true)]
        public Input<string> RestApiId { get; set; } = null!;

        public DocumentationPartArgs()
        {
        }
    }

    public sealed class DocumentationPartState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the targeted API entity of the to-be-created documentation part. See below.
        /// </summary>
        [Input("location")]
        public Input<Inputs.DocumentationPartLocationGetArgs>? Location { get; set; }

        /// <summary>
        /// A content map of API-specific key-value pairs describing the targeted API entity. The map must be encoded as a JSON string, e.g., "{ \"description\": \"The API does ...\" }". Only Swagger-compliant key-value pairs can be exported and, hence, published.
        /// </summary>
        [Input("properties")]
        public Input<string>? Properties { get; set; }

        /// <summary>
        /// The ID of the associated Rest API
        /// </summary>
        [Input("restApiId")]
        public Input<string>? RestApiId { get; set; }

        public DocumentationPartState()
        {
        }
    }
}
