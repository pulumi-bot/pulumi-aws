// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppMesh
{
    /// <summary>
    /// Provides an AWS App Mesh service mesh resource.
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
    ///         var simple = new Aws.AppMesh.Mesh("simple", new Aws.AppMesh.MeshArgs
    ///         {
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Egress Filter
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var simple = new Aws.AppMesh.Mesh("simple", new Aws.AppMesh.MeshArgs
    ///         {
    ///             Spec = new Aws.AppMesh.Inputs.MeshSpecArgs
    ///             {
    ///                 EgressFilter = new Aws.AppMesh.Inputs.MeshSpecEgressFilterArgs
    ///                 {
    ///                     Type = "ALLOW_ALL",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Mesh : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the service mesh.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The creation date of the service mesh.
        /// </summary>
        [Output("createdDate")]
        public Output<string> CreatedDate { get; private set; } = null!;

        /// <summary>
        /// The last update date of the service mesh.
        /// </summary>
        [Output("lastUpdatedDate")]
        public Output<string> LastUpdatedDate { get; private set; } = null!;

        /// <summary>
        /// The name to use for the service mesh.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The service mesh specification to apply.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.MeshSpec?> Spec { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Mesh resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Mesh(string name, MeshArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/mesh:Mesh", name, args ?? new MeshArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Mesh(string name, Input<string> id, MeshState? state = null, CustomResourceOptions? options = null)
            : base("aws:appmesh/mesh:Mesh", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Mesh resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Mesh Get(string name, Input<string> id, MeshState? state = null, CustomResourceOptions? options = null)
        {
            return new Mesh(name, id, state, options);
        }
    }

    public sealed class MeshArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name to use for the service mesh.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The service mesh specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.MeshSpecArgs>? Spec { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public MeshArgs()
        {
        }
    }

    public sealed class MeshState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the service mesh.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The creation date of the service mesh.
        /// </summary>
        [Input("createdDate")]
        public Input<string>? CreatedDate { get; set; }

        /// <summary>
        /// The last update date of the service mesh.
        /// </summary>
        [Input("lastUpdatedDate")]
        public Input<string>? LastUpdatedDate { get; set; }

        /// <summary>
        /// The name to use for the service mesh.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The service mesh specification to apply.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.MeshSpecGetArgs>? Spec { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public MeshState()
        {
        }
    }
}
