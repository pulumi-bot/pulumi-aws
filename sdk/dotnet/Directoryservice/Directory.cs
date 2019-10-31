// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Directoryservice
{
    /// <summary>
    /// Provides a Simple or Managed Microsoft directory in AWS Directory Service.
    /// 
    /// &gt; **Note:** All arguments including the password and customer username will be stored in the raw state as plain-text.
    /// [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/directory_service_directory.html.markdown.
    /// </summary>
    public partial class Directory : Pulumi.CustomResource
    {
        /// <summary>
        /// The access URL for the directory, such as `http://alias.awsapps.com`.
        /// </summary>
        [Output("accessUrl")]
        public Output<string> AccessUrl { get; private set; } = null!;

        /// <summary>
        /// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// Connector related information about the directory. Fields documented below.
        /// </summary>
        [Output("connectSettings")]
        public Output<Outputs.DirectoryConnectSettings?> ConnectSettings { get; private set; } = null!;

        /// <summary>
        /// A textual description for the directory.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A list of IP addresses of the DNS servers for the directory or connector.
        /// </summary>
        [Output("dnsIpAddresses")]
        public Output<ImmutableArray<string>> DnsIpAddresses { get; private set; } = null!;

        /// <summary>
        /// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        /// </summary>
        [Output("enableSso")]
        public Output<bool?> EnableSso { get; private set; } = null!;

        /// <summary>
        /// The fully qualified name for the directory, such as `corp.example.com`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password for the directory administrator or connector user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The ID of the security group created by the directory.
        /// </summary>
        [Output("securityGroupId")]
        public Output<string> SecurityGroupId { get; private set; } = null!;

        /// <summary>
        /// The short name of the directory, such as `CORP`.
        /// </summary>
        [Output("shortName")]
        public Output<string> ShortName { get; private set; } = null!;

        /// <summary>
        /// The size of the directory (`Small` or `Large` are accepted values).
        /// </summary>
        [Output("size")]
        public Output<string> Size { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;

        /// <summary>
        /// VPC related information about the directory. Fields documented below.
        /// </summary>
        [Output("vpcSettings")]
        public Output<Outputs.DirectoryVpcSettings?> VpcSettings { get; private set; } = null!;


        /// <summary>
        /// Create a Directory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Directory(string name, DirectoryArgs args, CustomResourceOptions? options = null)
            : base("aws:directoryservice/directory:Directory", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Directory(string name, Input<string> id, DirectoryState? state = null, CustomResourceOptions? options = null)
            : base("aws:directoryservice/directory:Directory", name, state, MakeResourceOptions(options, id))
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
        /// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Connector related information about the directory. Fields documented below.
        /// </summary>
        [Input("connectSettings")]
        public Input<Inputs.DirectoryConnectSettingsArgs>? ConnectSettings { get; set; }

        /// <summary>
        /// A textual description for the directory.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        /// </summary>
        [Input("enableSso")]
        public Input<bool>? EnableSso { get; set; }

        /// <summary>
        /// The fully qualified name for the directory, such as `corp.example.com`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password for the directory administrator or connector user.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The short name of the directory, such as `CORP`.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// The size of the directory (`Small` or `Large` are accepted values).
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

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

        /// <summary>
        /// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// VPC related information about the directory. Fields documented below.
        /// </summary>
        [Input("vpcSettings")]
        public Input<Inputs.DirectoryVpcSettingsArgs>? VpcSettings { get; set; }

        public DirectoryArgs()
        {
        }
    }

    public sealed class DirectoryState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access URL for the directory, such as `http://alias.awsapps.com`.
        /// </summary>
        [Input("accessUrl")]
        public Input<string>? AccessUrl { get; set; }

        /// <summary>
        /// The alias for the directory (must be unique amongst all aliases in AWS). Required for `enable_sso`.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// Connector related information about the directory. Fields documented below.
        /// </summary>
        [Input("connectSettings")]
        public Input<Inputs.DirectoryConnectSettingsGetArgs>? ConnectSettings { get; set; }

        /// <summary>
        /// A textual description for the directory.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dnsIpAddresses")]
        private InputList<string>? _dnsIpAddresses;

        /// <summary>
        /// A list of IP addresses of the DNS servers for the directory or connector.
        /// </summary>
        public InputList<string> DnsIpAddresses
        {
            get => _dnsIpAddresses ?? (_dnsIpAddresses = new InputList<string>());
            set => _dnsIpAddresses = value;
        }

        /// <summary>
        /// The MicrosoftAD edition (`Standard` or `Enterprise`). Defaults to `Enterprise` (applies to MicrosoftAD type only).
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Whether to enable single-sign on for the directory. Requires `alias`. Defaults to `false`.
        /// </summary>
        [Input("enableSso")]
        public Input<bool>? EnableSso { get; set; }

        /// <summary>
        /// The fully qualified name for the directory, such as `corp.example.com`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password for the directory administrator or connector user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The ID of the security group created by the directory.
        /// </summary>
        [Input("securityGroupId")]
        public Input<string>? SecurityGroupId { get; set; }

        /// <summary>
        /// The short name of the directory, such as `CORP`.
        /// </summary>
        [Input("shortName")]
        public Input<string>? ShortName { get; set; }

        /// <summary>
        /// The size of the directory (`Small` or `Large` are accepted values).
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

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

        /// <summary>
        /// The directory type (`SimpleAD`, `ADConnector` or `MicrosoftAD` are accepted values). Defaults to `SimpleAD`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// VPC related information about the directory. Fields documented below.
        /// </summary>
        [Input("vpcSettings")]
        public Input<Inputs.DirectoryVpcSettingsGetArgs>? VpcSettings { get; set; }

        public DirectoryState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class DirectoryConnectSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("customerDnsIps", required: true)]
        private InputList<string>? _customerDnsIps;

        /// <summary>
        /// The DNS IP addresses of the domain to connect to.
        /// </summary>
        public InputList<string> CustomerDnsIps
        {
            get => _customerDnsIps ?? (_customerDnsIps = new InputList<string>());
            set => _customerDnsIps = value;
        }

        /// <summary>
        /// The username corresponding to the password provided.
        /// </summary>
        [Input("customerUsername", required: true)]
        public Input<string> CustomerUsername { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DirectoryConnectSettingsArgs()
        {
        }
    }

    public sealed class DirectoryConnectSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("customerDnsIps", required: true)]
        private InputList<string>? _customerDnsIps;

        /// <summary>
        /// The DNS IP addresses of the domain to connect to.
        /// </summary>
        public InputList<string> CustomerDnsIps
        {
            get => _customerDnsIps ?? (_customerDnsIps = new InputList<string>());
            set => _customerDnsIps = value;
        }

        /// <summary>
        /// The username corresponding to the password provided.
        /// </summary>
        [Input("customerUsername", required: true)]
        public Input<string> CustomerUsername { get; set; } = null!;

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DirectoryConnectSettingsGetArgs()
        {
        }
    }

    public sealed class DirectoryVpcSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DirectoryVpcSettingsArgs()
        {
        }
    }

    public sealed class DirectoryVpcSettingsGetArgs : Pulumi.ResourceArgs
    {
        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public DirectoryVpcSettingsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class DirectoryConnectSettings
    {
        /// <summary>
        /// The DNS IP addresses of the domain to connect to.
        /// </summary>
        public readonly ImmutableArray<string> CustomerDnsIps;
        /// <summary>
        /// The username corresponding to the password provided.
        /// </summary>
        public readonly string CustomerUsername;
        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private DirectoryConnectSettings(
            ImmutableArray<string> customerDnsIps,
            string customerUsername,
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            CustomerDnsIps = customerDnsIps;
            CustomerUsername = customerUsername;
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }

    [OutputType]
    public sealed class DirectoryVpcSettings
    {
        /// <summary>
        /// The identifiers of the subnets for the directory servers (2 subnets in 2 different AZs).
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The identifier of the VPC that the directory is in.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private DirectoryVpcSettings(
            ImmutableArray<string> subnetIds,
            string vpcId)
        {
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
    }
}
