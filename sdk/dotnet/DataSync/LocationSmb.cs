// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.DataSync
{
    /// <summary>
    /// Manages a SMB Location within AWS DataSync.
    /// 
    /// &gt; **NOTE:** The DataSync Agents must be available before creating this resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.DataSync.LocationSmb("example", new Aws.DataSync.LocationSmbArgs
    ///         {
    ///             AgentArns = 
    ///             {
    ///                 aws_datasync_agent.Example.Arn,
    ///             },
    ///             Password = "ANotGreatPassword",
    ///             ServerHostname = "smb.example.com",
    ///             Subdirectory = "/exported/path",
    ///             User = "Guest",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class LocationSmb : Pulumi.CustomResource
    {
        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        [Output("agentArns")]
        public Output<ImmutableArray<string>> AgentArns { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the DataSync Location.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The name of the Windows domain the SMB server belongs to.
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        /// </summary>
        [Output("mountOptions")]
        public Output<Outputs.LocationSmbMountOptions?> MountOptions { get; private set; } = null!;

        /// <summary>
        /// The password of the user who can mount the share and has file permissions in the SMB.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        /// </summary>
        [Output("serverHostname")]
        public Output<string> ServerHostname { get; private set; } = null!;

        /// <summary>
        /// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        /// </summary>
        [Output("subdirectory")]
        public Output<string> Subdirectory { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        [Output("uri")]
        public Output<string> Uri { get; private set; } = null!;

        /// <summary>
        /// The user who can mount the share and has file and folder permissions in the SMB share.
        /// </summary>
        [Output("user")]
        public Output<string> User { get; private set; } = null!;


        /// <summary>
        /// Create a LocationSmb resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LocationSmb(string name, LocationSmbArgs args, CustomResourceOptions? options = null)
            : base("aws:datasync/locationSmb:LocationSmb", name, args ?? new LocationSmbArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LocationSmb(string name, Input<string> id, LocationSmbState? state = null, CustomResourceOptions? options = null)
            : base("aws:datasync/locationSmb:LocationSmb", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LocationSmb resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LocationSmb Get(string name, Input<string> id, LocationSmbState? state = null, CustomResourceOptions? options = null)
        {
            return new LocationSmb(name, id, state, options);
        }
    }

    public sealed class LocationSmbArgs : Pulumi.ResourceArgs
    {
        [Input("agentArns", required: true)]
        private InputList<string>? _agentArns;

        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        public InputList<string> AgentArns
        {
            get => _agentArns ?? (_agentArns = new InputList<string>());
            set => _agentArns = value;
        }

        /// <summary>
        /// The name of the Windows domain the SMB server belongs to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        /// </summary>
        [Input("mountOptions")]
        public Input<Inputs.LocationSmbMountOptionsArgs>? MountOptions { get; set; }

        /// <summary>
        /// The password of the user who can mount the share and has file permissions in the SMB.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        /// </summary>
        [Input("serverHostname", required: true)]
        public Input<string> ServerHostname { get; set; } = null!;

        /// <summary>
        /// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        /// </summary>
        [Input("subdirectory", required: true)]
        public Input<string> Subdirectory { get; set; } = null!;

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The user who can mount the share and has file and folder permissions in the SMB share.
        /// </summary>
        [Input("user", required: true)]
        public Input<string> User { get; set; } = null!;

        public LocationSmbArgs()
        {
        }
    }

    public sealed class LocationSmbState : Pulumi.ResourceArgs
    {
        [Input("agentArns")]
        private InputList<string>? _agentArns;

        /// <summary>
        /// A list of DataSync Agent ARNs with which this location will be associated.
        /// </summary>
        public InputList<string> AgentArns
        {
            get => _agentArns ?? (_agentArns = new InputList<string>());
            set => _agentArns = value;
        }

        /// <summary>
        /// Amazon Resource Name (ARN) of the DataSync Location.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The name of the Windows domain the SMB server belongs to.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Configuration block containing mount options used by DataSync to access the SMB Server. Can be `AUTOMATIC`, `SMB2`, or `SMB3`.
        /// </summary>
        [Input("mountOptions")]
        public Input<Inputs.LocationSmbMountOptionsGetArgs>? MountOptions { get; set; }

        /// <summary>
        /// The password of the user who can mount the share and has file permissions in the SMB.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Specifies the IP address or DNS name of the SMB server. The DataSync Agent(s) use this to mount the SMB share.
        /// </summary>
        [Input("serverHostname")]
        public Input<string>? ServerHostname { get; set; }

        /// <summary>
        /// Subdirectory to perform actions as source or destination. Should be exported by the NFS server.
        /// </summary>
        [Input("subdirectory")]
        public Input<string>? Subdirectory { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value pairs of resource tags to assign to the DataSync Location.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        [Input("uri")]
        public Input<string>? Uri { get; set; }

        /// <summary>
        /// The user who can mount the share and has file and folder permissions in the SMB share.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public LocationSmbState()
        {
        }
    }
}
