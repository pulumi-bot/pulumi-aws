// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.OpsWorks
{
    public partial class Permission : Pulumi.CustomResource
    {
        [Output("allowSsh")]
        public Output<bool> AllowSsh { get; private set; } = null!;

        [Output("allowSudo")]
        public Output<bool> AllowSudo { get; private set; } = null!;

        [Output("level")]
        public Output<string> Level { get; private set; } = null!;

        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        [Output("userArn")]
        public Output<string> UserArn { get; private set; } = null!;


        /// <summary>
        /// Create a Permission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permission(string name, PermissionArgs args, CustomResourceOptions? options = null)
            : base("aws:opsworks/permission:Permission", name, args ?? new PermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Permission(string name, Input<string> id, PermissionState? state = null, CustomResourceOptions? options = null)
            : base("aws:opsworks/permission:Permission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Permission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permission Get(string name, Input<string> id, PermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new Permission(name, id, state, options);
        }
    }

    public sealed class PermissionArgs : Pulumi.ResourceArgs
    {
        [Input("allowSsh")]
        public Input<bool>? AllowSsh { get; set; }

        [Input("allowSudo")]
        public Input<bool>? AllowSudo { get; set; }

        [Input("level")]
        public Input<string>? Level { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("userArn", required: true)]
        public Input<string> UserArn { get; set; } = null!;

        public PermissionArgs()
        {
        }
    }

    public sealed class PermissionState : Pulumi.ResourceArgs
    {
        [Input("allowSsh")]
        public Input<bool>? AllowSsh { get; set; }

        [Input("allowSudo")]
        public Input<bool>? AllowSudo { get; set; }

        [Input("level")]
        public Input<string>? Level { get; set; }

        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        [Input("userArn")]
        public Input<string>? UserArn { get; set; }

        public PermissionState()
        {
        }
    }
}
