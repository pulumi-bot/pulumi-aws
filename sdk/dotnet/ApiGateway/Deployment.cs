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
    /// Manages an API Gateway REST Deployment. A deployment is a snapshot of the REST API configuration. The deployment can then be published to callable endpoints via the `aws.apigateway.Stage` resource and optionally managed further with the `aws.apigateway.BasePathMapping` resource, `aws.apigateway.DomainName` resource, and `aws_api_method_settings` resource. For more information, see the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html).
    /// 
    /// To properly capture all REST API configuration in a deployment, this resource must have dependencies on all prior resources that manage resources/paths, methods, integrations, etc.
    /// 
    /// * For REST APIs that are configured via OpenAPI specification (`aws.apigateway.RestApi` resource `body` argument), no special dependency setup is needed beyond referencing the  `id` attribute of that resource unless additional resources have further customized the REST API.
    /// * When the REST API configuration involves other resources (`aws.apigateway.Integration` resource), the dependency setup can be done with implicit resource references in the `triggers` argument or explicit resource references using the [resource `dependsOn` custom option](https://www.pulumi.com/docs/intro/concepts/resources/#dependson). The `triggers` argument should be preferred over `depends_on`, since `depends_on` can only capture dependency ordering and will not cause the resource to recreate (redeploy the REST API) with upstream configuration changes.
    /// 
    /// !&gt; **WARNING:** It is recommended to use the `aws.apigateway.Stage` resource instead of managing an API Gateway Stage via the `stage_name` argument of this resource. When this resource is recreated (REST API redeployment) with the `stage_name` configured, the stage is deleted and recreated. This will cause a temporary service interruption, increase provide plan differences, and can require a second apply to recreate any downstream stage configuration such as associated `aws_api_method_settings` resources.
    /// 
    /// ## Example Usage
    /// </summary>
    [AwsResourceType("aws:apigateway/deployment:Deployment")]
    public partial class Deployment : Pulumi.CustomResource
    {
        /// <summary>
        /// The creation date of the deployment
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// Description of the deployment
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The execution ARN to be used in `lambda_permission` resource's `source_arn`
        /// when allowing API Gateway to invoke a Lambda function,
        /// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        /// </summary>
        [Output("executionArn")]
        public Output<string> ExecutionArn { get; private set; } = null!;

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        /// </summary>
        [Output("invokeUrl")]
        public Output<string> InvokeUrl { get; private set; } = null!;

        /// <summary>
        /// REST API identifier.
        /// </summary>
        [Output("restApi")]
        public Output<string> RestApi { get; private set; } = null!;

        /// <summary>
        /// Description to set on the stage managed by the `stage_name` argument.
        /// </summary>
        [Output("stageDescription")]
        public Output<string?> StageDescription { get; private set; } = null!;

        /// <summary>
        /// Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
        /// </summary>
        [Output("stageName")]
        public Output<string?> StageName { get; private set; } = null!;

        /// <summary>
        /// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableDictionary<string, string>?> Triggers { get; private set; } = null!;

        /// <summary>
        /// Map to set on the stage managed by the `stage_name` argument.
        /// </summary>
        [Output("variables")]
        public Output<ImmutableDictionary<string, string>?> Variables { get; private set; } = null!;


        /// <summary>
        /// Create a Deployment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deployment(string name, DeploymentArgs args, CustomResourceOptions? options = null)
            : base("aws:apigateway/deployment:Deployment", name, args ?? new DeploymentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Deployment(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/deployment:Deployment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Deployment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deployment Get(string name, Input<string> id, DeploymentState? state = null, CustomResourceOptions? options = null)
        {
            return new Deployment(name, id, state, options);
        }
    }

    public sealed class DeploymentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the deployment
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// REST API identifier.
        /// </summary>
        [Input("restApi", required: true)]
        public Input<string> RestApi { get; set; } = null!;

        /// <summary>
        /// Description to set on the stage managed by the `stage_name` argument.
        /// </summary>
        [Input("stageDescription")]
        public Input<string>? StageDescription { get; set; }

        /// <summary>
        /// Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        [Input("triggers")]
        private InputMap<string>? _triggers;

        /// <summary>
        /// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
        /// </summary>
        public InputMap<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<string>());
            set => _triggers = value;
        }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// Map to set on the stage managed by the `stage_name` argument.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public DeploymentArgs()
        {
        }
    }

    public sealed class DeploymentState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation date of the deployment
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// Description of the deployment
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The execution ARN to be used in `lambda_permission` resource's `source_arn`
        /// when allowing API Gateway to invoke a Lambda function,
        /// e.g. `arn:aws:execute-api:eu-west-2:123456789012:z4675bid1j/prod`
        /// </summary>
        [Input("executionArn")]
        public Input<string>? ExecutionArn { get; set; }

        /// <summary>
        /// The URL to invoke the API pointing to the stage,
        /// e.g. `https://z4675bid1j.execute-api.eu-west-2.amazonaws.com/prod`
        /// </summary>
        [Input("invokeUrl")]
        public Input<string>? InvokeUrl { get; set; }

        /// <summary>
        /// REST API identifier.
        /// </summary>
        [Input("restApi")]
        public Input<string>? RestApi { get; set; }

        /// <summary>
        /// Description to set on the stage managed by the `stage_name` argument.
        /// </summary>
        [Input("stageDescription")]
        public Input<string>? StageDescription { get; set; }

        /// <summary>
        /// Name of the stage to create with this deployment. If the specified stage already exists, it will be updated to point to the new deployment. It is recommended to use the `aws.apigateway.Stage` resource instead to manage stages.
        /// </summary>
        [Input("stageName")]
        public Input<string>? StageName { get; set; }

        [Input("triggers")]
        private InputMap<string>? _triggers;

        /// <summary>
        /// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
        /// </summary>
        public InputMap<string> Triggers
        {
            get => _triggers ?? (_triggers = new InputMap<string>());
            set => _triggers = value;
        }

        [Input("variables")]
        private InputMap<string>? _variables;

        /// <summary>
        /// Map to set on the stage managed by the `stage_name` argument.
        /// </summary>
        public InputMap<string> Variables
        {
            get => _variables ?? (_variables = new InputMap<string>());
            set => _variables = value;
        }

        public DeploymentState()
        {
        }
    }
}
