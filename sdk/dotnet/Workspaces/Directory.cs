// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Workspaces
{
    /// <summary>
    /// Provides a directory registration in AWS WorkSpaces Service
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var mainVpc = new Aws.Ec2.Vpc("mainVpc", new Aws.Ec2.VpcArgs
    ///         {
    ///             CidrBlock = "10.0.0.0/16",
    ///         });
    ///         var private-a = new Aws.Ec2.Subnet("private-a", new Aws.Ec2.SubnetArgs
    ///         {
    ///             AvailabilityZone = "us-east-1a",
    ///             CidrBlock = "10.0.0.0/24",
    ///             VpcId = mainVpc.Id,
    ///         });
    ///         var private-b = new Aws.Ec2.Subnet("private-b", new Aws.Ec2.SubnetArgs
    ///         {
    ///             AvailabilityZone = "us-east-1b",
    ///             CidrBlock = "10.0.1.0/24",
    ///             VpcId = mainVpc.Id,
    ///         });
    ///         var mainDirectory = new Aws.DirectoryService.Directory("mainDirectory", new Aws.DirectoryService.DirectoryArgs
    ///         {
    ///             Password = "#S1ncerely",
    ///             Size = "Small",
    ///             VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
    ///             {
    ///                 SubnetIds = 
    ///                 {
    ///                     private-a.Id,
    ///                     private-b.Id,
    ///                 },
    ///                 VpcId = mainVpc.Id,
    ///             },
    ///         });
    ///         var mainWorkspaces/directoryDirectory = new Aws.Workspaces.Directory("mainWorkspaces/directoryDirectory", new Aws.Workspaces.DirectoryArgs
    ///         {
    ///             DirectoryId = mainDirectory.Id,
    ///             SelfServicePermissions = new Aws.Workspaces.Inputs.DirectorySelfServicePermissionsArgs
    ///             {
    ///                 IncreaseVolumeSize = true,
    ///                 RebuildWorkspace = true,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class Directory : Pulumi.CustomResource
    {
        /// <summary>
        /// The directory alias.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// The user name for the service account.
        /// </summary>
        [Output("customerUserName")]
        public Output<string> CustomerUserName { get; private set; } = null!;

        /// <summary>
        /// The directory identifier for registration in WorkSpaces service.
        /// </summary>
        [Output("directoryId")]
        public Output<string> DirectoryId { get; private set; } = null!;

        /// <summary>
        /// The name of the directory.
        /// </summary>
        [Output("directoryName")]
        public Output<string> DirectoryName { get; private set; } = null!;

        /// <summary>
        /// The directory type.
        /// </summary>
        [Output("directoryType")]
        public Output<string> DirectoryType { get; private set; } = null!;

        /// <summary>
        /// The IP addresses of the DNS servers for the directory.
        /// </summary>
        [Output("dnsIpAddresses")]
        public Output<ImmutableArray<string>> DnsIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces to make calls to other services, such as Amazon EC2, on your behalf.
        /// </summary>
        [Output("iamRoleId")]
        public Output<string> IamRoleId { get; private set; } = null!;

        /// <summary>
        /// The identifiers of the IP access control groups associated with the directory.
        /// </summary>
        [Output("ipGroupIds")]
        public Output<ImmutableArray<string>> IpGroupIds { get; private set; } = null!;

        /// <summary>
        /// The registration code for the directory. This is the code that users enter in their Amazon WorkSpaces client application to connect to the directory.
        /// </summary>
        [Output("registrationCode")]
        public Output<string> RegistrationCode { get; private set; } = null!;

        /// <summary>
        /// The permissions to enable or disable self-service capabilities.
        /// </summary>
        [Output("selfServicePermissions")]
        public Output<Outputs.DirectorySelfServicePermissions> SelfServicePermissions { get; private set; } = null!;

        /// <summary>
        /// The identifiers of the subnets where the directory resides.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// A map of tags assigned to the WorkSpaces directory.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The identifier of the security group that is assigned to new WorkSpaces.
        /// </summary>
        [Output("workspaceSecurityGroupId")]
        public Output<string> WorkspaceSecurityGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a Directory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Directory(string name, DirectoryArgs args, CustomResourceOptions? options = null)
            : base("aws:workspaces/directory:Directory", name, args ?? new DirectoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Directory(string name, Input<string> id, DirectoryState? state = null, CustomResourceOptions? options = null)
            : base("aws:workspaces/directory:Directory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Directory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Directory Get(string name, Input<string> id, DirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Directory(name, id, state, options);
        }
    }

    public sealed class DirectoryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The directory identifier for registration in WorkSpaces service.
        /// </summary>
        [Input("directoryId", required: true)]
        public Input<string> DirectoryId { get; set; } = null!;

        /// <summary>
        /// The permissions to enable or disable self-service capabilities.
        /// </summary>
        [Input("selfServicePermissions")]
        public Input<Inputs.DirectorySelfServicePermissionsArgs>? SelfServicePermissions { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets where the directory resides.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the WorkSpaces directory.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public DirectoryArgs()
        {
        }
    }

    public sealed class DirectoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The directory alias.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The user name for the service account.
        /// </summary>
        [Input("customerUserName")]
        public Input<string>? CustomerUserName { get; set; }

        /// <summary>
        /// The directory identifier for registration in WorkSpaces service.
        /// </summary>
        [Input("directoryId")]
        public Input<string>? DirectoryId { get; set; }

        /// <summary>
        /// The name of the directory.
        /// </summary>
        [Input("directoryName")]
        public Input<string>? DirectoryName { get; set; }

        /// <summary>
        /// The directory type.
        /// </summary>
        [Input("directoryType")]
        public Input<string>? DirectoryType { get; set; }

        [Input("dnsIpAddresses")]
        private InputList<string>? _dnsIpAddresses;

        /// <summary>
        /// The IP addresses of the DNS servers for the directory.
        /// </summary>
        public InputList<string> DnsIpAddresses
        {
            get => _dnsIpAddresses ?? (_dnsIpAddresses = new InputList<string>());
            set => _dnsIpAddresses = value;
        }

        /// <summary>
        /// The identifier of the IAM role. This is the role that allows Amazon WorkSpaces to make calls to other services, such as Amazon EC2, on your behalf.
        /// </summary>
        [Input("iamRoleId")]
        public Input<string>? IamRoleId { get; set; }

        [Input("ipGroupIds")]
        private InputList<string>? _ipGroupIds;

        /// <summary>
        /// The identifiers of the IP access control groups associated with the directory.
        /// </summary>
        public InputList<string> IpGroupIds
        {
            get => _ipGroupIds ?? (_ipGroupIds = new InputList<string>());
            set => _ipGroupIds = value;
        }

        /// <summary>
        /// The registration code for the directory. This is the code that users enter in their Amazon WorkSpaces client application to connect to the directory.
        /// </summary>
        [Input("registrationCode")]
        public Input<string>? RegistrationCode { get; set; }

        /// <summary>
        /// The permissions to enable or disable self-service capabilities.
        /// </summary>
        [Input("selfServicePermissions")]
        public Input<Inputs.DirectorySelfServicePermissionsGetArgs>? SelfServicePermissions { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets where the directory resides.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A map of tags assigned to the WorkSpaces directory.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// The identifier of the security group that is assigned to new WorkSpaces.
        /// </summary>
        [Input("workspaceSecurityGroupId")]
        public Input<string>? WorkspaceSecurityGroupId { get; set; }

        public DirectoryState()
        {
        }
    }
}
