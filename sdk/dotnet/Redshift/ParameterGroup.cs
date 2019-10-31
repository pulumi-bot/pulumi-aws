// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Redshift
{
    /// <summary>
    /// Provides a Redshift Cluster parameter group resource.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/redshift_parameter_group.html.markdown.
    /// </summary>
    public partial class ParameterGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of parameter group
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The family of the Redshift parameter group.
        /// </summary>
        [Output("family")]
        public Output<string> Family { get; private set; } = null!;

        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of Redshift parameters to apply.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ParameterGroupParameters>> Parameters { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ParameterGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ParameterGroup(string name, ParameterGroupArgs args, CustomResourceOptions? options = null)
            : base("aws:redshift/parameterGroup:ParameterGroup", name, args, MakeResourceOptions(options, ""))
        {
        }

        private ParameterGroup(string name, Input<string> id, ParameterGroupState? state = null, CustomResourceOptions? options = null)
            : base("aws:redshift/parameterGroup:ParameterGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ParameterGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ParameterGroup Get(string name, Input<string> id, ParameterGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ParameterGroup(name, id, state, options);
        }
    }

    public sealed class ParameterGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The family of the Redshift parameter group.
        /// </summary>
        [Input("family", required: true)]
        public Input<string> Family { get; set; } = null!;

        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ParameterGroupParametersArgs>? _parameters;

        /// <summary>
        /// A list of Redshift parameters to apply.
        /// </summary>
        public InputList<Inputs.ParameterGroupParametersArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ParameterGroupParametersArgs>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ParameterGroupArgs()
        {
            Description = "Managed by Pulumi";
        }
    }

    public sealed class ParameterGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Name (ARN) of parameter group
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The description of the Redshift parameter group. Defaults to "Managed by Pulumi".
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The family of the Redshift parameter group.
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.ParameterGroupParametersGetArgs>? _parameters;

        /// <summary>
        /// A list of Redshift parameters to apply.
        /// </summary>
        public InputList<Inputs.ParameterGroupParametersGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ParameterGroupParametersGetArgs>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ParameterGroupState()
        {
            Description = "Managed by Pulumi";
        }
    }

    namespace Inputs
    {

    public sealed class ParameterGroupParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the Redshift parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ParameterGroupParametersArgs()
        {
        }
    }

    public sealed class ParameterGroupParametersGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The value of the Redshift parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ParameterGroupParametersGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ParameterGroupParameters
    {
        /// <summary>
        /// The name of the Redshift parameter.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the Redshift parameter.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ParameterGroupParameters(
            string name,
            string value)
        {
            Name = name;
            Value = value;
        }
    }
    }
}
