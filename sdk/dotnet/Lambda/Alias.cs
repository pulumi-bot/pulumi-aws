// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    /// <summary>
    /// Creates a Lambda function alias. Creates an alias that points to the specified Lambda function version.
    /// 
    /// For information about Lambda and how to use it, see [What is AWS Lambda?][1]
    /// For information about function aliases, see [CreateAlias][2] and [AliasRoutingConfiguration][3] in the API docs.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lambda_alias.html.markdown.
    /// </summary>
    public partial class Alias : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda function alias.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Description of the alias.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The function ARN of the Lambda function for which you want to create an alias.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
        /// </summary>
        [Output("functionVersion")]
        public Output<string> FunctionVersion { get; private set; } = null!;

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Output("invokeArn")]
        public Output<string> InvokeArn { get; private set; } = null!;

        /// <summary>
        /// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Lambda alias' route configuration settings. Fields documented below
        /// </summary>
        [Output("routingConfig")]
        public Output<Outputs.AliasRoutingConfig?> RoutingConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Alias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Alias(string name, AliasArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/alias:Alias", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Alias(string name, Input<string> id, AliasState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/alias:Alias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Alias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Alias Get(string name, Input<string> id, AliasState? state = null, CustomResourceOptions? options = null)
        {
            return new Alias(name, id, state, options);
        }
    }

    public sealed class AliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the alias.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The function ARN of the Lambda function for which you want to create an alias.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
        /// </summary>
        [Input("functionVersion", required: true)]
        public Input<string> FunctionVersion { get; set; } = null!;

        /// <summary>
        /// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Lambda alias' route configuration settings. Fields documented below
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.AliasRoutingConfigArgs>? RoutingConfig { get; set; }

        public AliasArgs()
        {
        }
    }

    public sealed class AliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) identifying your Lambda function alias.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// Description of the alias.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The function ARN of the Lambda function for which you want to create an alias.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Lambda function version for which you are creating the alias. Pattern: `(\$LATEST|[0-9]+)`.
        /// </summary>
        [Input("functionVersion")]
        public Input<string>? FunctionVersion { get; set; }

        /// <summary>
        /// The ARN to be used for invoking Lambda Function from API Gateway - to be used in [`aws.apigateway.Integration`](https://www.terraform.io/docs/providers/aws/r/api_gateway_integration.html)'s `uri`
        /// </summary>
        [Input("invokeArn")]
        public Input<string>? InvokeArn { get; set; }

        /// <summary>
        /// Name for the alias you are creating. Pattern: `(?!^[0-9]+$)([a-zA-Z0-9-_]+)`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Lambda alias' route configuration settings. Fields documented below
        /// </summary>
        [Input("routingConfig")]
        public Input<Inputs.AliasRoutingConfigGetArgs>? RoutingConfig { get; set; }

        public AliasState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class AliasRoutingConfigArgs : Pulumi.ResourceArgs
    {
        [Input("additionalVersionWeights")]
        private InputMap<double>? _additionalVersionWeights;

        /// <summary>
        /// A map that defines the proportion of events that should be sent to different versions of a lambda function.
        /// </summary>
        public InputMap<double> AdditionalVersionWeights
        {
            get => _additionalVersionWeights ?? (_additionalVersionWeights = new InputMap<double>());
            set => _additionalVersionWeights = value;
        }

        public AliasRoutingConfigArgs()
        {
        }
    }

    public sealed class AliasRoutingConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("additionalVersionWeights")]
        private InputMap<double>? _additionalVersionWeights;

        /// <summary>
        /// A map that defines the proportion of events that should be sent to different versions of a lambda function.
        /// </summary>
        public InputMap<double> AdditionalVersionWeights
        {
            get => _additionalVersionWeights ?? (_additionalVersionWeights = new InputMap<double>());
            set => _additionalVersionWeights = value;
        }

        public AliasRoutingConfigGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class AliasRoutingConfig
    {
        /// <summary>
        /// A map that defines the proportion of events that should be sent to different versions of a lambda function.
        /// </summary>
        public readonly ImmutableDictionary<string, double>? AdditionalVersionWeights;

        [OutputConstructor]
        private AliasRoutingConfig(ImmutableDictionary<string, double>? additionalVersionWeights)
        {
            AdditionalVersionWeights = additionalVersionWeights;
        }
    }
    }
}
