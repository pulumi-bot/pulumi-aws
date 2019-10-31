// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lightsail
{
    /// <summary>
    /// Provides a Lightsail Instance. Amazon Lightsail is a service to provide easy virtual private servers
    /// with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
    /// for more information.
    /// 
    /// &gt; **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details
    /// 
    /// ## Availability Zones
    /// 
    /// Lightsail currently supports the following Availability Zones (e.g. `us-east-1a`):
    /// 
    /// - `ap-northeast-1{a,c,d}`
    /// - `ap-northeast-2{a,c}`
    /// - `ap-south-1{a,b}`
    /// - `ap-southeast-1{a,b,c}`
    /// - `ap-southeast-2{a,b,c}`
    /// - `ca-central-1{a,b}`
    /// - `eu-central-1{a,b,c}`
    /// - `eu-west-1{a,b,c}`
    /// - `eu-west-2{a,b,c}`
    /// - `eu-west-3{a,b,c}`
    /// - `us-east-1{a,b,c,d,e,f}`
    /// - `us-east-2{a,b,c}`
    /// - `us-west-2{a,b,c}`
    /// 
    /// ## Blueprints
    /// 
    /// Lightsail currently supports the following Blueprint IDs:
    /// 
    /// ### OS Only
    /// 
    /// - `amazon_linux_2018_03_0_2`
    /// - `centos_7_1901_01`
    /// - `debian_8_7`
    /// - `debian_9_5`
    /// - `freebsd_11_1`
    /// - `opensuse_42_2`
    /// - `ubuntu_16_04_2`
    /// - `ubuntu_18_04`
    /// 
    /// ### Apps and OS
    /// 
    /// - `drupal_8_5_6`
    /// - `gitlab_11_1_4_1`
    /// - `joomla_3_8_11`
    /// - `lamp_5_6_37_2`
    /// - `lamp_7_1_20_1`
    /// - `magento_2_2_5`
    /// - `mean_4_0_1`
    /// - `nginx_1_14_0_1`
    /// - `nodejs_10_8_0`
    /// - `plesk_ubuntu_17_8_11_1`
    /// - `redmine_3_4_6`
    /// - `wordpress_4_9_8`
    /// - `wordpress_multisite_4_9_8`
    /// 
    /// ## Bundles
    /// 
    /// Lightsail currently supports the following Bundle IDs (e.g. an instance in `ap-northeast-1` would use `small_2_0`):
    /// 
    /// ### Prefix
    /// 
    /// A Bundle ID starts with one of the below size prefixes:
    /// 
    /// - `nano_`
    /// - `micro_`
    /// - `small_`
    /// - `medium_`
    /// - `large_`
    /// - `xlarge_`
    /// - `2xlarge_`
    /// 
    /// ### Suffix
    /// 
    /// A Bundle ID ends with one of the following suffixes depending on Availability Zone:
    /// 
    /// - ap-northeast-1: `2_0`
    /// - ap-northeast-2: `2_0`
    /// - ap-south-1: `2_1`
    /// - ap-southeast-1: `2_0`
    /// - ap-southeast-2: `2_2`
    /// - ca-central-1: `2_0`
    /// - eu-central-1: `2_0`
    /// - eu-west-1: `2_0`
    /// - eu-west-2: `2_0`
    /// - eu-west-3: `2_0`
    /// - us-east-1: `2_0`
    /// - us-east-2: `2_0`
    /// - us-west-2: `2_0`
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/lightsail_instance.html.markdown.
    /// </summary>
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// The ARN of the Lightsail instance (matches `id`).
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// The ID for a virtual private server image
        /// (see list below)
        /// </summary>
        [Output("blueprintId")]
        public Output<string> BlueprintId { get; private set; } = null!;

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Output("bundleId")]
        public Output<string> BundleId { get; private set; } = null!;

        [Output("cpuCount")]
        public Output<int> CpuCount { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the instance was created.
        /// * `availability_zone`
        /// * `blueprint_id`
        /// * `bundle_id`
        /// * `key_pair_name`
        /// * `user_data`
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("ipv6Address")]
        public Output<string> Ipv6Address { get; private set; } = null!;

        [Output("isStaticIp")]
        public Output<bool> IsStaticIp { get; private set; } = null!;

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Output("keyPairName")]
        public Output<string?> KeyPairName { get; private set; } = null!;

        /// <summary>
        /// The name of the Lightsail Instance
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;

        [Output("publicIpAddress")]
        public Output<string> PublicIpAddress { get; private set; } = null!;

        [Output("ramSize")]
        public Output<int> RamSize { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// launch script to configure server with additional user data
        /// </summary>
        [Output("userData")]
        public Output<string?> UserData { get; private set; } = null!;

        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("aws:lightsail/instance:Instance", name, args, MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("aws:lightsail/instance:Instance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Input("availabilityZone", required: true)]
        public Input<string> AvailabilityZone { get; set; } = null!;

        /// <summary>
        /// The ID for a virtual private server image
        /// (see list below)
        /// </summary>
        [Input("blueprintId", required: true)]
        public Input<string> BlueprintId { get; set; } = null!;

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Input("bundleId", required: true)]
        public Input<string> BundleId { get; set; } = null!;

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// The name of the Lightsail Instance
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// launch script to configure server with additional user data
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        public InstanceArgs()
        {
        }
    }

    public sealed class InstanceState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Lightsail instance (matches `id`).
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The Availability Zone in which to create your
        /// instance (see list below)
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// The ID for a virtual private server image
        /// (see list below)
        /// </summary>
        [Input("blueprintId")]
        public Input<string>? BlueprintId { get; set; }

        /// <summary>
        /// The bundle of specification information (see list below)
        /// </summary>
        [Input("bundleId")]
        public Input<string>? BundleId { get; set; }

        [Input("cpuCount")]
        public Input<int>? CpuCount { get; set; }

        /// <summary>
        /// The timestamp when the instance was created.
        /// * `availability_zone`
        /// * `blueprint_id`
        /// * `bundle_id`
        /// * `key_pair_name`
        /// * `user_data`
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("ipv6Address")]
        public Input<string>? Ipv6Address { get; set; }

        [Input("isStaticIp")]
        public Input<bool>? IsStaticIp { get; set; }

        /// <summary>
        /// The name of your key pair. Created in the
        /// Lightsail console (cannot use `aws.ec2.KeyPair` at this time)
        /// </summary>
        [Input("keyPairName")]
        public Input<string>? KeyPairName { get; set; }

        /// <summary>
        /// The name of the Lightsail Instance
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        [Input("publicIpAddress")]
        public Input<string>? PublicIpAddress { get; set; }

        [Input("ramSize")]
        public Input<int>? RamSize { get; set; }

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
        /// launch script to configure server with additional user data
        /// </summary>
        [Input("userData")]
        public Input<string>? UserData { get; set; }

        [Input("username")]
        public Input<string>? Username { get; set; }

        public InstanceState()
        {
        }
    }
}
