// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cognito
{
    /// <summary>
    /// Provides a Cognito Resource Server.
    /// 
    /// ## Example Usage
    /// </summary>
    public partial class ResourceServer : Pulumi.CustomResource
    {
        /// <summary>
        /// An identifier for the resource server.
        /// </summary>
        [Output("identifier")]
        public Output<string> Identifier { get; private set; } = null!;

        /// <summary>
        /// A name for the resource server.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of all scopes configured for this resource server in the format identifier/scope_name.
        /// </summary>
        [Output("scopeIdentifiers")]
        public Output<ImmutableArray<string>> ScopeIdentifiers { get; private set; } = null!;

        /// <summary>
        /// A list of Authorization Scope.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<Outputs.ResourceServerScope>> Scopes { get; private set; } = null!;

        [Output("userPoolId")]
        public Output<string> UserPoolId { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceServer(string name, ResourceServerArgs args, CustomResourceOptions? options = null)
            : base("aws:cognito/resourceServer:ResourceServer", name, args ?? new ResourceServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceServer(string name, Input<string> id, ResourceServerState? state = null, CustomResourceOptions? options = null)
            : base("aws:cognito/resourceServer:ResourceServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceServer Get(string name, Input<string> id, ResourceServerState? state = null, CustomResourceOptions? options = null)
        {
            return new ResourceServer(name, id, state, options);
        }
    }

    public sealed class ResourceServerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identifier for the resource server.
        /// </summary>
        [Input("identifier", required: true)]
        public Input<string> Identifier { get; set; } = null!;

        /// <summary>
        /// A name for the resource server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<Inputs.ResourceServerScopeArgs>? _scopes;

        /// <summary>
        /// A list of Authorization Scope.
        /// </summary>
        public InputList<Inputs.ResourceServerScopeArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.ResourceServerScopeArgs>());
            set => _scopes = value;
        }

        [Input("userPoolId", required: true)]
        public Input<string> UserPoolId { get; set; } = null!;

        public ResourceServerArgs()
        {
        }
    }

    public sealed class ResourceServerState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An identifier for the resource server.
        /// </summary>
        [Input("identifier")]
        public Input<string>? Identifier { get; set; }

        /// <summary>
        /// A name for the resource server.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopeIdentifiers")]
        private InputList<string>? _scopeIdentifiers;

        /// <summary>
        /// A list of all scopes configured for this resource server in the format identifier/scope_name.
        /// </summary>
        public InputList<string> ScopeIdentifiers
        {
            get => _scopeIdentifiers ?? (_scopeIdentifiers = new InputList<string>());
            set => _scopeIdentifiers = value;
        }

        [Input("scopes")]
        private InputList<Inputs.ResourceServerScopeGetArgs>? _scopes;

        /// <summary>
        /// A list of Authorization Scope.
        /// </summary>
        public InputList<Inputs.ResourceServerScopeGetArgs> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<Inputs.ResourceServerScopeGetArgs>());
            set => _scopes = value;
        }

        [Input("userPoolId")]
        public Input<string>? UserPoolId { get; set; }

        public ResourceServerState()
        {
        }
    }
}
