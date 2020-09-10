// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Efs
{
    public partial class FileSystemPolicy : Pulumi.CustomResource
    {
        [Output("fileSystemId")]
        public Output<string> FileSystemId { get; private set; } = null!;

        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a FileSystemPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FileSystemPolicy(string name, FileSystemPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:efs/fileSystemPolicy:FileSystemPolicy", name, args ?? new FileSystemPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FileSystemPolicy(string name, Input<string> id, FileSystemPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:efs/fileSystemPolicy:FileSystemPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FileSystemPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FileSystemPolicy Get(string name, Input<string> id, FileSystemPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new FileSystemPolicy(name, id, state, options);
        }
    }

    public sealed class FileSystemPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("fileSystemId", required: true)]
        public Input<string> FileSystemId { get; set; } = null!;

        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public FileSystemPolicyArgs()
        {
        }
    }

    public sealed class FileSystemPolicyState : Pulumi.ResourceArgs
    {
        [Input("fileSystemId")]
        public Input<string>? FileSystemId { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public FileSystemPolicyState()
        {
        }
    }
}
