// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    public partial class VpcEndpointRouteTableAssociation : Pulumi.CustomResource
    {
        [Output("routeTableId")]
        public Output<string> RouteTableId { get; private set; } = null!;

        [Output("vpcEndpointId")]
        public Output<string> VpcEndpointId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcEndpointRouteTableAssociation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcEndpointRouteTableAssociation(string name, VpcEndpointRouteTableAssociationArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, args ?? new VpcEndpointRouteTableAssociationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcEndpointRouteTableAssociation(string name, Input<string> id, VpcEndpointRouteTableAssociationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/vpcEndpointRouteTableAssociation:VpcEndpointRouteTableAssociation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcEndpointRouteTableAssociation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcEndpointRouteTableAssociation Get(string name, Input<string> id, VpcEndpointRouteTableAssociationState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcEndpointRouteTableAssociation(name, id, state, options);
        }
    }

    public sealed class VpcEndpointRouteTableAssociationArgs : Pulumi.ResourceArgs
    {
        [Input("routeTableId", required: true)]
        public Input<string> RouteTableId { get; set; } = null!;

        [Input("vpcEndpointId", required: true)]
        public Input<string> VpcEndpointId { get; set; } = null!;

        public VpcEndpointRouteTableAssociationArgs()
        {
        }
    }

    public sealed class VpcEndpointRouteTableAssociationState : Pulumi.ResourceArgs
    {
        [Input("routeTableId")]
        public Input<string>? RouteTableId { get; set; }

        [Input("vpcEndpointId")]
        public Input<string>? VpcEndpointId { get; set; }

        public VpcEndpointRouteTableAssociationState()
        {
        }
    }
}
