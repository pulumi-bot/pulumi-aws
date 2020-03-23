// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Sagemaker
{
    /// <summary>
    /// Provides a lifecycle configuration for SageMaker Notebook Instances.
    /// 
    /// {{% examples %}}
    /// {{% /examples %}}
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/sagemaker_notebook_instance_lifecycle_configuration.html.markdown.
    /// </summary>
    public partial class NotebookInstanceLifecycleConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
        /// </summary>
        [Output("onCreate")]
        public Output<string?> OnCreate { get; private set; } = null!;

        /// <summary>
        /// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
        /// </summary>
        [Output("onStart")]
        public Output<string?> OnStart { get; private set; } = null!;


        /// <summary>
        /// Create a NotebookInstanceLifecycleConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NotebookInstanceLifecycleConfiguration(string name, NotebookInstanceLifecycleConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private NotebookInstanceLifecycleConfiguration(string name, Input<string> id, NotebookInstanceLifecycleConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:sagemaker/notebookInstanceLifecycleConfiguration:NotebookInstanceLifecycleConfiguration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NotebookInstanceLifecycleConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NotebookInstanceLifecycleConfiguration Get(string name, Input<string> id, NotebookInstanceLifecycleConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new NotebookInstanceLifecycleConfiguration(name, id, state, options);
        }
    }

    public sealed class NotebookInstanceLifecycleConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
        /// </summary>
        [Input("onCreate")]
        public Input<string>? OnCreate { get; set; }

        /// <summary>
        /// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
        /// </summary>
        [Input("onStart")]
        public Input<string>? OnStart { get; set; }

        public NotebookInstanceLifecycleConfigurationArgs()
        {
        }
    }

    public sealed class NotebookInstanceLifecycleConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) assigned by AWS to this lifecycle configuration.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the lifecycle configuration (must be unique). If omitted, this provider will assign a random, unique name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A shell script (base64-encoded) that runs only once when the SageMaker Notebook Instance is created.
        /// </summary>
        [Input("onCreate")]
        public Input<string>? OnCreate { get; set; }

        /// <summary>
        /// A shell script (base64-encoded) that runs every time the SageMaker Notebook Instance is started including the time it's created.
        /// </summary>
        [Input("onStart")]
        public Input<string>? OnStart { get; set; }

        public NotebookInstanceLifecycleConfigurationState()
        {
        }
    }
}
