// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ssm
{
    /// <summary>
    /// Registers an on-premises server or virtual machine with Amazon EC2 so that it can be managed using Run Command.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/r/ssm_activation.html.markdown.
    /// </summary>
    public partial class Activation : Pulumi.CustomResource
    {
        /// <summary>
        /// The code the system generates when it processes the activation.
        /// </summary>
        [Output("activationCode")]
        public Output<string> ActivationCode { get; private set; } = null!;

        /// <summary>
        /// The description of the resource that you want to register.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time.
        /// </summary>
        [Output("expirationDate")]
        public Output<string?> ExpirationDate { get; private set; } = null!;

        /// <summary>
        /// If the current activation has expired.
        /// </summary>
        [Output("expired")]
        public Output<string> Expired { get; private set; } = null!;

        /// <summary>
        /// The IAM Role to attach to the managed instance.
        /// </summary>
        [Output("iamRole")]
        public Output<string> IamRole { get; private set; } = null!;

        /// <summary>
        /// The default name of the registered managed instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of managed instances that are currently registered using this activation.
        /// </summary>
        [Output("registrationCount")]
        public Output<int> RegistrationCount { get; private set; } = null!;

        /// <summary>
        /// The maximum number of managed instances you want to register. The default value is 1 instance.
        /// </summary>
        [Output("registrationLimit")]
        public Output<int?> RegistrationLimit { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Activation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Activation(string name, ActivationArgs args, CustomResourceOptions? options = null)
            : base("aws:ssm/activation:Activation", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Activation(string name, Input<string> id, ActivationState? state = null, CustomResourceOptions? options = null)
            : base("aws:ssm/activation:Activation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Activation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Activation Get(string name, Input<string> id, ActivationState? state = null, CustomResourceOptions? options = null)
        {
            return new Activation(name, id, state, options);
        }
    }

    public sealed class ActivationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the resource that you want to register.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// The IAM Role to attach to the managed instance.
        /// </summary>
        [Input("iamRole", required: true)]
        public Input<string> IamRole { get; set; } = null!;

        /// <summary>
        /// The default name of the registered managed instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The maximum number of managed instances you want to register. The default value is 1 instance.
        /// </summary>
        [Input("registrationLimit")]
        public Input<int>? RegistrationLimit { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ActivationArgs()
        {
        }
    }

    public sealed class ActivationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The code the system generates when it processes the activation.
        /// </summary>
        [Input("activationCode")]
        public Input<string>? ActivationCode { get; set; }

        /// <summary>
        /// The description of the resource that you want to register.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) by which this activation request should expire. The default value is 24 hours from resource creation time.
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// If the current activation has expired.
        /// </summary>
        [Input("expired")]
        public Input<string>? Expired { get; set; }

        /// <summary>
        /// The IAM Role to attach to the managed instance.
        /// </summary>
        [Input("iamRole")]
        public Input<string>? IamRole { get; set; }

        /// <summary>
        /// The default name of the registered managed instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The number of managed instances that are currently registered using this activation.
        /// </summary>
        [Input("registrationCount")]
        public Input<int>? RegistrationCount { get; set; }

        /// <summary>
        /// The maximum number of managed instances you want to register. The default value is 1 instance.
        /// </summary>
        [Input("registrationLimit")]
        public Input<int>? RegistrationLimit { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the object.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ActivationState()
        {
        }
    }
}
