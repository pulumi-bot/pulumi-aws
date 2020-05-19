// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync
{
    /// <summary>
    /// Provides an AppSync Resolver.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var testGraphQLApi = new Aws.AppSync.GraphQLApi("testGraphQLApi", new Aws.AppSync.GraphQLApiArgs
    ///         {
    ///             AuthenticationType = "API_KEY",
    ///             Schema = @"type Mutation {
    /// 	putPost(id: ID!, title: String!): Post
    /// }
    /// 
    /// type Post {
    /// 	id: ID!
    /// 	title: String!
    /// }
    /// 
    /// type Query {
    /// 	singlePost(id: ID!): Post
    /// }
    /// 
    /// schema {
    /// 	query: Query
    /// 	mutation: Mutation
    /// }
    /// 
    /// ",
    ///         });
    ///         var testDataSource = new Aws.AppSync.DataSource("testDataSource", new Aws.AppSync.DataSourceArgs
    ///         {
    ///             ApiId = testGraphQLApi.Id,
    ///             HttpConfig = new Aws.AppSync.Inputs.DataSourceHttpConfigArgs
    ///             {
    ///                 Endpoint = "http://example.com",
    ///             },
    ///             Type = "HTTP",
    ///         });
    ///         // UNIT type resolver (default)
    ///         var testResolver = new Aws.AppSync.Resolver("testResolver", new Aws.AppSync.ResolverArgs
    ///         {
    ///             ApiId = testGraphQLApi.Id,
    ///             DataSource = testDataSource.Name,
    ///             Field = "singlePost",
    ///             RequestTemplate = @"{
    ///     ""version"": ""2018-05-29"",
    ///     ""method"": ""GET"",
    ///     ""resourcePath"": ""/"",
    ///     ""params"":{
    ///         ""headers"": $$utils.http.copyheaders($$ctx.request.headers)
    ///     }
    /// }
    /// 
    /// ",
    ///             ResponseTemplate = @"#if($$ctx.result.statusCode == 200)
    ///     $$ctx.result.body
    /// #else
    ///     $$utils.appendError($$ctx.result.body, $$ctx.result.statusCode)
    /// #end
    /// 
    /// ",
    ///             Type = "Query",
    ///         });
    ///         // PIPELINE type resolver
    ///         var mutationPipelineTest = new Aws.AppSync.Resolver("mutationPipelineTest", new Aws.AppSync.ResolverArgs
    ///         {
    ///             ApiId = testGraphQLApi.Id,
    ///             Field = "pipelineTest",
    ///             Kind = "PIPELINE",
    ///             PipelineConfig = new Aws.AppSync.Inputs.ResolverPipelineConfigArgs
    ///             {
    ///                 Functions = 
    ///                 {
    ///                     aws_appsync_function.Test1.Function_id,
    ///                     aws_appsync_function.Test2.Function_id,
    ///                     aws_appsync_function.Test3.Function_id,
    ///                 },
    ///             },
    ///             RequestTemplate = "{}",
    ///             ResponseTemplate = "$$util.toJson($$ctx.result)",
    ///             Type = "Mutation",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Resolver : Pulumi.CustomResource
    {
        /// <summary>
        /// The API ID for the GraphQL API.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The ARN
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The DataSource name.
        /// </summary>
        [Output("dataSource")]
        public Output<string?> DataSource { get; private set; } = null!;

        /// <summary>
        /// The field name from the schema defined in the GraphQL API.
        /// </summary>
        [Output("field")]
        public Output<string> Field { get; private set; } = null!;

        /// <summary>
        /// The resolver type. Valid values are `UNIT` and `PIPELINE`.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The PipelineConfig. A `pipeline_config` block is documented below.
        /// </summary>
        [Output("pipelineConfig")]
        public Output<Outputs.ResolverPipelineConfig?> PipelineConfig { get; private set; } = null!;

        /// <summary>
        /// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
        /// </summary>
        [Output("requestTemplate")]
        public Output<string> RequestTemplate { get; private set; } = null!;

        /// <summary>
        /// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
        /// </summary>
        [Output("responseTemplate")]
        public Output<string> ResponseTemplate { get; private set; } = null!;

        /// <summary>
        /// The type name from the schema defined in the GraphQL API.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Resolver resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Resolver(string name, ResolverArgs args, CustomResourceOptions? options = null)
            : base("aws:appsync/resolver:Resolver", name, args ?? new ResolverArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Resolver(string name, Input<string> id, ResolverState? state = null, CustomResourceOptions? options = null)
            : base("aws:appsync/resolver:Resolver", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Resolver resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Resolver Get(string name, Input<string> id, ResolverState? state = null, CustomResourceOptions? options = null)
        {
            return new Resolver(name, id, state, options);
        }
    }

    public sealed class ResolverArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API ID for the GraphQL API.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The DataSource name.
        /// </summary>
        [Input("dataSource")]
        public Input<string>? DataSource { get; set; }

        /// <summary>
        /// The field name from the schema defined in the GraphQL API.
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// The resolver type. Valid values are `UNIT` and `PIPELINE`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The PipelineConfig. A `pipeline_config` block is documented below.
        /// </summary>
        [Input("pipelineConfig")]
        public Input<Inputs.ResolverPipelineConfigArgs>? PipelineConfig { get; set; }

        /// <summary>
        /// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
        /// </summary>
        [Input("requestTemplate", required: true)]
        public Input<string> RequestTemplate { get; set; } = null!;

        /// <summary>
        /// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
        /// </summary>
        [Input("responseTemplate", required: true)]
        public Input<string> ResponseTemplate { get; set; } = null!;

        /// <summary>
        /// The type name from the schema defined in the GraphQL API.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ResolverArgs()
        {
        }
    }

    public sealed class ResolverState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API ID for the GraphQL API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The ARN
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The DataSource name.
        /// </summary>
        [Input("dataSource")]
        public Input<string>? DataSource { get; set; }

        /// <summary>
        /// The field name from the schema defined in the GraphQL API.
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// The resolver type. Valid values are `UNIT` and `PIPELINE`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The PipelineConfig. A `pipeline_config` block is documented below.
        /// </summary>
        [Input("pipelineConfig")]
        public Input<Inputs.ResolverPipelineConfigGetArgs>? PipelineConfig { get; set; }

        /// <summary>
        /// The request mapping template for UNIT resolver or 'before mapping template' for PIPELINE resolver.
        /// </summary>
        [Input("requestTemplate")]
        public Input<string>? RequestTemplate { get; set; }

        /// <summary>
        /// The response mapping template for UNIT resolver or 'after mapping template' for PIPELINE resolver.
        /// </summary>
        [Input("responseTemplate")]
        public Input<string>? ResponseTemplate { get; set; }

        /// <summary>
        /// The type name from the schema defined in the GraphQL API.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ResolverState()
        {
        }
    }
}
