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
    /// Provides an API Gateway Method Settings, e.g. logging or monitoring.
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
    ///         var testRestApi = new Aws.ApiGateway.RestApi("testRestApi", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///             Description = "This is my API for demonstration purposes",
    ///         });
    ///         var testResource = new Aws.ApiGateway.Resource("testResource", new Aws.ApiGateway.ResourceArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ParentId = testRestApi.RootResourceId,
    ///             PathPart = "mytestresource",
    ///         });
    ///         var testMethod = new Aws.ApiGateway.Method("testMethod", new Aws.ApiGateway.MethodArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ResourceId = testResource.Id,
    ///             HttpMethod = "GET",
    ///             Authorization = "NONE",
    ///         });
    ///         var testIntegration = new Aws.ApiGateway.Integration("testIntegration", new Aws.ApiGateway.IntegrationArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             ResourceId = testResource.Id,
    ///             HttpMethod = testMethod.HttpMethod,
    ///             Type = "MOCK",
    ///             RequestTemplates = 
    ///             {
    ///                 { "application/xml", @"{
    ///    ""body"" : $input.json('$')
    /// }
    /// " },
    ///             },
    ///         });
    ///         var testDeployment = new Aws.ApiGateway.Deployment("testDeployment", new Aws.ApiGateway.DeploymentArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             StageName = "dev",
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 testIntegration,
    ///             },
    ///         });
    ///         var testStage = new Aws.ApiGateway.Stage("testStage", new Aws.ApiGateway.StageArgs
    ///         {
    ///             StageName = "prod",
    ///             RestApi = testRestApi.Id,
    ///             Deployment = testDeployment.Id,
    ///         });
    ///         var methodSettings = new Aws.ApiGateway.MethodSettings("methodSettings", new Aws.ApiGateway.MethodSettingsArgs
    ///         {
    ///             RestApi = testRestApi.Id,
    ///             StageName = testStage.StageName,
    ///             MethodPath = Output.Tuple(testResource.PathPart, testMethod.HttpMethod).Apply(values =&gt;
    ///             {
    ///                 var pathPart = values.Item1;
    ///                 var httpMethod = values.Item2;
    ///                 return $"{pathPart}/{httpMethod}";
    ///             }),
    ///             Settings = new Aws.ApiGateway.Inputs.MethodSettingsSettingsArgs
    ///             {
    ///                 MetricsEnabled = true,
    ///                 LoggingLevel = "INFO",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// `aws_api_gateway_method_settings` can be imported using `REST-API-ID/STAGE-NAME/METHOD-PATH`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:apigateway/methodSettings:MethodSettings example 12345abcde/example/test/GET
    /// ```
    /// </summary>
    [AwsResourceType("aws:apigateway/methodSettings:MethodSettings")]
    public partial class MethodSettings : Pulumi.CustomResource
    {
        /// <summary>
        /// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
        /// </summary>
        [Output("methodPath")]
        public Output<string> MethodPath { get; private set; } = null!;

        /// <summary>
        /// The ID of the REST API
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// The settings block, see below.
        /// </summary>
        [Output("settings")]
        public Output<Outputs.MethodSettingsSettings> Settings { get; private set; } = null!;

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Output("stageName")]
        public Output<string> StageName { get; private set; } = null!;


        /// <summary>
        /// Create a MethodSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MethodSettings(string name, MethodSettingsArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/methodSettings:MethodSettings", name, args ?? new MethodSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MethodSettings(string name, Input<string> id, MethodSettingsState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/methodSettings:MethodSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MethodSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MethodSettings Get(string name, Input<string> id, MethodSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new MethodSettings(name, id, state, options);
        }
    }

    public sealed class MethodSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
        /// </summary>
        [Input("methodPath", required: true)]
        public Input<string> MethodPath { get; set; } = null!;

        /// <summary>
        /// The ID of the REST API
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// The settings block, see below.
        /// </summary>
        [Input("settings", required: true)]
        public Input<Inputs.MethodSettingsSettingsArgs> Settings { get; set; } = null!;

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Input("stageName", required: true)]
        public Input<string> StageName { get; set; } = null!;

        public MethodSettingsArgs()
        {
        }
    }

    public sealed class MethodSettingsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Method path defined as `{resource_path}/{http_method}` for an individual method override, or `*/*` for overriding all methods in the stage.
        /// </summary>
        [Input("methodPath")]
        public Input<string>? MethodPath { get; set; }

        /// <summary>
        /// The ID of the REST API
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// The settings block, see below.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.MethodSettingsSettingsGetArgs>? Settings { get; set; }

        /// <summary>
        /// The name of the stage
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        public MethodSettingsState()
        {
        }
    }
}
