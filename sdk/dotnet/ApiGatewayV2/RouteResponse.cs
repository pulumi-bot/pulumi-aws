// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2
{
    public partial class RouteResponse : Pulumi.CustomResource
    {
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        [Output("modelSelectionExpression")]
        public Output<string?> ModelSelectionExpression { get; private set; } = null!;

        [Output("responseModels")]
        public Output<ImmutableDictionary<string, string>?> ResponseModels { get; private set; } = null!;

        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        [Output("routeResponseKey")]
        public Output<string> RouteResponseKey { get; private set; } = null!;


        /// <summary>
        /// Create a RouteResponse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteResponse(string name, RouteResponseArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/routeResponse:RouteResponse", name, args ?? new RouteResponseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteResponse(string name, Input<string> id, RouteResponseState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/routeResponse:RouteResponse", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteResponse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteResponse Get(string name, Input<string> id, RouteResponseState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteResponse(name, id, state, options);
        }
    }

    public sealed class RouteResponseArgs : Pulumi.ResourceArgs
    {
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        [Input("modelSelectionExpression")]
        public Input<string>? ModelSelectionExpression { get; set; }

        [Input("responseModels")]
        private InputMap<string>? _responseModels;
        public InputMap<string> ResponseModels
        {
            get => _responseModels ?? (_responseModels = new InputMap<string>());
            set => _responseModels = value;
        }

        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        [Input("routeResponseKey", required: true)]
        public Input<string> RouteResponseKey { get; set; } = null!;

        public RouteResponseArgs()
        {
        }
    }

    public sealed class RouteResponseState : Pulumi.ResourceArgs
    {
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        [Input("modelSelectionExpression")]
        public Input<string>? ModelSelectionExpression { get; set; }

        [Input("responseModels")]
        private InputMap<string>? _responseModels;
        public InputMap<string> ResponseModels
        {
            get => _responseModels ?? (_responseModels = new InputMap<string>());
            set => _responseModels = value;
        }

        [Input("routeId")]
        public Input<string>? RouteId { get; set; }

        [Input("routeResponseKey")]
        public Input<string>? RouteResponseKey { get; set; }

        public RouteResponseState()
        {
        }
    }
}
