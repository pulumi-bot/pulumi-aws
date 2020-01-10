// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// [IPv6 only] Creates an egress-only Internet gateway for your VPC.
    /// An egress-only Internet gateway is used to enable outbound communication
    /// over IPv6 from instances in your VPC to the Internet, and prevents hosts
    /// outside of your VPC from initiating an IPv6 connection with your instance.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/egress_only_internet_gateway.html.markdown.
    /// </summary>
    public partial class EgressOnlyInternetGateway : Pulumi.CustomResource
    {
        /// <summary>
        /// The VPC ID to create in.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;


        /// <summary>
        /// Create a EgressOnlyInternetGateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EgressOnlyInternetGateway(string name, EgressOnlyInternetGatewayArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private EgressOnlyInternetGateway(string name, Input<string> id, EgressOnlyInternetGatewayState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/egressOnlyInternetGateway:EgressOnlyInternetGateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EgressOnlyInternetGateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EgressOnlyInternetGateway Get(string name, Input<string> id, EgressOnlyInternetGatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new EgressOnlyInternetGateway(name, id, state, options);
        }
    }

    public sealed class EgressOnlyInternetGatewayArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VPC ID to create in.
        /// </summary>
        [Input("vpcId", required: true)]
        public Input<string> VpcId { get; set; } = null!;

        public EgressOnlyInternetGatewayArgs()
        {
        }
    }

    public sealed class EgressOnlyInternetGatewayState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The VPC ID to create in.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public EgressOnlyInternetGatewayState()
        {
        }
    }
}
