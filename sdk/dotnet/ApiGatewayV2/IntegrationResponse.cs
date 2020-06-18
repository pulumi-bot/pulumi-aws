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
    /// Manages an Amazon API Gateway Version 2 integration response.
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
    ///         var example = new Aws.ApiGatewayV2.IntegrationResponse("example", new Aws.ApiGatewayV2.IntegrationResponseArgs
    ///         {
    ///             ApiId = aws_apigatewayv2_api.Example.Id,
    ///             IntegrationId = aws_apigatewayv2_integration.Example.Id,
    ///             IntegrationResponseKey = "/200/",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class IntegrationResponse : Pulumi.CustomResource
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
        /// </summary>
        [Output("contentHandlingStrategy")]
        public Output<string?> ContentHandlingStrategy { get; private set; } = null!;

        /// <summary>
        /// The identifier of the `aws.apigatewayv2.Integration`.
        /// </summary>
        [Output("integrationId")]
        public Output<string> IntegrationId { get; private set; } = null!;

        /// <summary>
        /// The integration response key.
        /// </summary>
        [Output("integrationResponseKey")]
        public Output<string> IntegrationResponseKey { get; private set; } = null!;

        /// <summary>
        /// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
        /// </summary>
        [Output("responseTemplates")]
        public Output<ImmutableDictionary<string, string>?> ResponseTemplates { get; private set; } = null!;

        /// <summary>
        /// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
        /// </summary>
        [Output("templateSelectionExpression")]
        public Output<string?> TemplateSelectionExpression { get; private set; } = null!;


        /// <summary>
        /// Create a IntegrationResponse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IntegrationResponse(string name, IntegrationResponseArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/integrationResponse:IntegrationResponse", name, args ?? new IntegrationResponseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IntegrationResponse(string name, Input<string> id, IntegrationResponseState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/integrationResponse:IntegrationResponse", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IntegrationResponse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IntegrationResponse Get(string name, Input<string> id, IntegrationResponseState? state = null, CustomResourceOptions? options = null)
        {
            return new IntegrationResponse(name, id, state, options);
        }
    }

    public sealed class IntegrationResponseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
        /// </summary>
        [Input("contentHandlingStrategy")]
        public Input<string>? ContentHandlingStrategy { get; set; }

        /// <summary>
        /// The identifier of the `aws.apigatewayv2.Integration`.
        /// </summary>
        [Input("integrationId", required: true)]
        public Input<string> IntegrationId { get; set; } = null!;

        /// <summary>
        /// The integration response key.
        /// </summary>
        [Input("integrationResponseKey", required: true)]
        public Input<string> IntegrationResponseKey { get; set; } = null!;

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
        /// </summary>
        [Input("templateSelectionExpression")]
        public Input<string>? TemplateSelectionExpression { get; set; }

        public IntegrationResponseArgs()
        {
        }
    }

    public sealed class IntegrationResponseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// How to handle response payload content type conversions. Valid values: `CONVERT_TO_BINARY`, `CONVERT_TO_TEXT`.
        /// </summary>
        [Input("contentHandlingStrategy")]
        public Input<string>? ContentHandlingStrategy { get; set; }

        /// <summary>
        /// The identifier of the `aws.apigatewayv2.Integration`.
        /// </summary>
        [Input("integrationId")]
        public Input<string>? IntegrationId { get; set; }

        /// <summary>
        /// The integration response key.
        /// </summary>
        [Input("integrationResponseKey")]
        public Input<string>? IntegrationResponseKey { get; set; }

        [Input("responseTemplates")]
        private InputMap<string>? _responseTemplates;

        /// <summary>
        /// A map of Velocity templates that are applied on the request payload based on the value of the Content-Type header sent by the client.
        /// </summary>
        public InputMap<string> ResponseTemplates
        {
            get => _responseTemplates ?? (_responseTemplates = new InputMap<string>());
            set => _responseTemplates = value;
        }

        /// <summary>
        /// The [template selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-template-selection-expressions) for the integration response.
        /// </summary>
        [Input("templateSelectionExpression")]
        public Input<string>? TemplateSelectionExpression { get; set; }

        public IntegrationResponseState()
        {
        }
    }
}
