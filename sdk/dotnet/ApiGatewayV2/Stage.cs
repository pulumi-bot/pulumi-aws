// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    /// <summary>
    /// Manages an Amazon API Gateway Version 2 stage.
    /// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).
    /// 
    /// ## Example Usage
    /// ### Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.ApiGatewayV2.Stage("example", new Aws.ApiGatewayV2.StageArgs
    ///         {
    ///             ApiId = aws_apigatewayv2_api.Example.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Stage : Pulumi.CustomResource
    {
        /// <summary>
        /// Settings for logging access in this stage.
        /// Use the `aws.apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
        /// </summary>
        [Output("accessLogSettings")]
        public Output<Outputs.StageAccessLogSettings?> AccessLogSettings { get; private set; } = null!;

        /// <summary>
        /// The API identifier.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The ARN of the stage.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
        /// </summary>
        [Output("autoDeploy")]
        public Output<bool?> AutoDeploy { get; private set; } = null!;

        /// <summary>
        /// The identifier of a client certificate for the stage. Use the `aws.apigateway.ClientCertificate` resource to configure a client certificate.
        /// Supported only for WebSocket APIs.
        /// </summary>
        [Output("clientCertificateId")]
        public Output<string?> ClientCertificateId { get; private set; } = null!;

        /// <summary>
        /// The default route settings for the stage.
        /// </summary>
        [Output("defaultRouteSettings")]
        public Output<Outputs.StageDefaultRouteSettings?> DefaultRouteSettings { get; private set; } = null!;

        /// <summary>
        /// The deployment identifier of the stage. Use the `aws.apigatewayv2.Deployment` resource to configure a deployment.
        /// </summary>
        [Output("deploymentId")]
        public Output<string?> DeploymentId { get; private set; } = null!;

        /// <summary>
        /// The description for the stage.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ARN prefix to be used in an `aws.lambda.Permission`'s `source_arn` attribute
        /// or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        /// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        /// Set only for WebSocket APIs.
        /// </summary>
        [Output("executionArn")]
        public Output<string> ExecutionArn { get; private set; } = null!;

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
        /// </summary>
        [Output("invokeUrl")]
        public Output<string> InvokeUrl { get; private set; } = null!;

        /// <summary>
        /// The name of the stage.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Route settings for the stage.
        /// </summary>
        [Output("routeSettings")]
        public Output<ImmutableArray<Outputs.StageRouteSetting>> RouteSettings { get; private set; } = null!;

        /// <summary>
        /// A map that defines the stage variables for the stage.
        /// </summary>
        [Output("stageVariables")]
        public Output<ImmutableDictionary<string, string>?> StageVariables { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the stage.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Stage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stage(string name, StageArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/stage:Stage", name, args ?? new StageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stage(string name, Input<string> id, StageState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/stage:Stage", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Stage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stage Get(string name, Input<string> id, StageState? state = null, CustomResourceOptions? options = null)
        {
            return new Stage(name, id, state, options);
        }
    }

    public sealed class StageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for logging access in this stage.
        /// Use the `aws.apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
        /// </summary>
        [Input("accessLogSettings")]
        public Input<Inputs.StageAccessLogSettingsArgs>? AccessLogSettings { get; set; }

        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
        /// </summary>
        [Input("autoDeploy")]
        public Input<bool>? AutoDeploy { get; set; }

        /// <summary>
        /// The identifier of a client certificate for the stage. Use the `aws.apigateway.ClientCertificate` resource to configure a client certificate.
        /// Supported only for WebSocket APIs.
        /// </summary>
        [Input("clientCertificateId")]
        public Input<string>? ClientCertificateId { get; set; }

        /// <summary>
        /// The default route settings for the stage.
        /// </summary>
        [Input("defaultRouteSettings")]
        public Input<Inputs.StageDefaultRouteSettingsArgs>? DefaultRouteSettings { get; set; }

        /// <summary>
        /// The deployment identifier of the stage. Use the `aws.apigatewayv2.Deployment` resource to configure a deployment.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The description for the stage.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the stage.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("routeSettings")]
        private InputList<Inputs.StageRouteSettingArgs>? _routeSettings;

        /// <summary>
        /// Route settings for the stage.
        /// </summary>
        public InputList<Inputs.StageRouteSettingArgs> RouteSettings
        {
            get => _routeSettings ?? (_routeSettings = new InputList<Inputs.StageRouteSettingArgs>());
            set => _routeSettings = value;
        }

        [Input("stageVariables")]
        private InputMap<string>? _stageVariables;

        /// <summary>
        /// A map that defines the stage variables for the stage.
        /// </summary>
        public InputMap<string> StageVariables
        {
            get => _stageVariables ?? (_stageVariables = new InputMap<string>());
            set => _stageVariables = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the stage.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public StageArgs()
        {
        }
    }

    public sealed class StageState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for logging access in this stage.
        /// Use the `aws.apigateway.Account` resource to configure [permissions for CloudWatch Logging](https://docs.aws.amazon.com/apigateway/latest/developerguide/set-up-logging.html#set-up-access-logging-permissions).
        /// </summary>
        [Input("accessLogSettings")]
        public Input<Inputs.StageAccessLogSettingsGetArgs>? AccessLogSettings { get; set; }

        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The ARN of the stage.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Whether updates to an API automatically trigger a new deployment. Defaults to `false`.
        /// </summary>
        [Input("autoDeploy")]
        public Input<bool>? AutoDeploy { get; set; }

        /// <summary>
        /// The identifier of a client certificate for the stage. Use the `aws.apigateway.ClientCertificate` resource to configure a client certificate.
        /// Supported only for WebSocket APIs.
        /// </summary>
        [Input("clientCertificateId")]
        public Input<string>? ClientCertificateId { get; set; }

        /// <summary>
        /// The default route settings for the stage.
        /// </summary>
        [Input("defaultRouteSettings")]
        public Input<Inputs.StageDefaultRouteSettingsGetArgs>? DefaultRouteSettings { get; set; }

        /// <summary>
        /// The deployment identifier of the stage. Use the `aws.apigatewayv2.Deployment` resource to configure a deployment.
        /// </summary>
        [Input("deploymentId")]
        public Input<string>? DeploymentId { get; set; }

        /// <summary>
        /// The description for the stage.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ARN prefix to be used in an `aws.lambda.Permission`'s `source_arn` attribute
        /// or in an `aws.iam.Policy` to authorize access to the [`@connections` API](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html).
        /// See the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-control-access-iam.html) for details.
        /// Set only for WebSocket APIs.
        /// </summary>
        [Input("executionArn")]
        public Input<string>? ExecutionArn { get; set; }

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `wss://z4675bid1j.execute-api.eu-west-2.amazonaws.com/example-stage`, or `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/`
        /// </summary>
        [Input("invokeUrl")]
        public Input<string>? InvokeUrl { get; set; }

        /// <summary>
        /// The name of the stage.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("routeSettings")]
        private InputList<Inputs.StageRouteSettingGetArgs>? _routeSettings;

        /// <summary>
        /// Route settings for the stage.
        /// </summary>
        public InputList<Inputs.StageRouteSettingGetArgs> RouteSettings
        {
            get => _routeSettings ?? (_routeSettings = new InputList<Inputs.StageRouteSettingGetArgs>());
            set => _routeSettings = value;
        }

        [Input("stageVariables")]
        private InputMap<string>? _stageVariables;

        /// <summary>
        /// A map that defines the stage variables for the stage.
        /// </summary>
        public InputMap<string> StageVariables
        {
            get => _stageVariables ?? (_stageVariables = new InputMap<string>());
            set => _stageVariables = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the stage.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public StageState()
        {
        }
    }
}
