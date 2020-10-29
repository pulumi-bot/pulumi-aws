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
    /// Provides an API Gateway API Key.
    /// 
    /// &gt; **NOTE:** Since the API Gateway usage plans feature was launched on August 11, 2016, usage plans are now **required** to associate an API key with an API stage.
    /// </summary>
    public partial class ApiKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The creation date of the API key
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The API key description. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the API key can be used by callers. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The last update date of the API key
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// The name of the API key
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
        /// </summary>
        [Output("value")]
        public Output<string> Value { get; private set; } = null!;


        /// <summary>
        /// Create a ApiKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApiKey(string name, ApiKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/apiKey:ApiKey", name, args ?? new ApiKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApiKey(string name, Input<string> id, ApiKeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:apigateway/apiKey:ApiKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApiKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApiKey Get(string name, Input<string> id, ApiKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new ApiKey(name, id, state, options);
        }
    }

    public sealed class ApiKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API key description. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether the API key can be used by callers. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the API key
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApiKeyArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ApiKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN)
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The creation date of the API key
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The API key description. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether the API key can be used by callers. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The last update date of the API key
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// The name of the API key
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The value of the API key. If not specified, it will be automatically generated by AWS on creation.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public ApiKeyState()
        {
            Description = "Managed by Pulumi";
        }
    }
}
