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
    /// Manages an Amazon API Gateway Version 2 API mapping.
    /// More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html).
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class ApiMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
        /// </summary>
        [Output("apiMappingKey")]
        public Output<string?> ApiMappingKey { get; private set; } = null!;

        /// <summary>
        /// The domain name. Use the `aws.apigatewayv2.DomainName` resource to configure a domain name.
        /// </summary>
        [Output("domainName")]
        public Output<string> DomainName { get; private set; } = null!;

        /// <summary>
        /// The API stage. Use the `aws.apigatewayv2.Stage` resource to configure an API stage.
        /// </summary>
        [Output("stage")]
        public Output<string> Stage { get; private set; } = null!;


        /// <summary>
        /// Create a ApiMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiMapping(string name, ApiMappingArgs args, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/apiMapping:ApiMapping", name, args ?? new ApiMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiMapping(string name, Input<string> id, ApiMappingState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigatewayv2/apiMapping:ApiMapping", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiMapping Get(string name, Input<string> id, ApiMappingState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiMapping(name, id, state, options);
        }
    }

    public sealed class ApiMappingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
        /// </summary>
        [Input("apiMappingKey")]
        public Input<string>? ApiMappingKey { get; set; }

        /// <summary>
        /// The domain name. Use the `aws.apigatewayv2.DomainName` resource to configure a domain name.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The API stage. Use the `aws.apigatewayv2.Stage` resource to configure an API stage.
        /// </summary>
        [Input("stage", required: true)]
        public Input<string> Stage { get; set; } = null!;

        public ApiMappingArgs()
        {
        }
    }

    public sealed class ApiMappingState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API identifier.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The [API mapping key](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-mapping-template-reference.html).
        /// </summary>
        [Input("apiMappingKey")]
        public Input<string>? ApiMappingKey { get; set; }

        /// <summary>
        /// The domain name. Use the `aws.apigatewayv2.DomainName` resource to configure a domain name.
        /// </summary>
        [Input("domainName")]
        public Input<string>? DomainName { get; set; }

        /// <summary>
        /// The API stage. Use the `aws.apigatewayv2.Stage` resource to configure an API stage.
        /// </summary>
        [Input("stage")]
        public Input<string>? Stage { get; set; }

        public ApiMappingState()
        {
        }
    }
}
