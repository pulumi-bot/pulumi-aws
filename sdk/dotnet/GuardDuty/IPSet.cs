// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.GuardDuty
{
    /// <summary>
    /// Provides a resource to manage a GuardDuty IPSet.
    /// 
    /// &gt; **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the master account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)
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
    ///         var master = new Aws.GuardDuty.Detector("master", new Aws.GuardDuty.DetectorArgs
    ///         {
    ///             Enable = true,
    ///         });
    ///         var bucket = new Aws.S3.Bucket("bucket", new Aws.S3.BucketArgs
    ///         {
    ///             Acl = "private",
    ///         });
    ///         var myIPSetBucketObject = new Aws.S3.BucketObject("myIPSetBucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Acl = "public-read",
    ///             Bucket = bucket.Id,
    ///             Content = @"10.0.0.0/8
    /// 
    /// ",
    ///             Key = "MyIPSet",
    ///         });
    ///         var myIPSetIPSet = new Aws.GuardDuty.IPSet("myIPSetIPSet", new Aws.GuardDuty.IPSetArgs
    ///         {
    ///             Activate = true,
    ///             DetectorId = master.Id,
    ///             Format = "TXT",
    ///             Location = Output.Tuple(myIPSetBucketObject.Bucket, myIPSetBucketObject.Key).Apply(values =&gt;
    ///             {
    ///                 var bucket = values.Item1;
    ///                 var key = values.Item2;
    ///                 return $"https://s3.amazonaws.com/{bucket}/{key}";
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class IPSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies whether GuardDuty is to start using the uploaded IPSet.
        /// </summary>
        [Output("activate")]
        public Output<bool> Activate { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the GuardDuty IPSet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Output("detectorId")]
        public Output<string> DetectorId { get; private set; } = null!;

        /// <summary>
        /// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
        /// </summary>
        [Output("format")]
        public Output<string> Format { get; private set; } = null!;

        /// <summary>
        /// The URI of the file that contains the IPSet.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The friendly name to identify the IPSet.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a IPSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IPSet(string name, IPSetArgs args, CustomResourceOptions? options = null)
            : base("aws:guardduty/iPSet:IPSet", name, args ?? new IPSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IPSet(string name, Input<string> id, IPSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:guardduty/iPSet:IPSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IPSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IPSet Get(string name, Input<string> id, IPSetState? state = null, CustomResourceOptions? options = null)
        {
            return new IPSet(name, id, state, options);
        }
    }

    public sealed class IPSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether GuardDuty is to start using the uploaded IPSet.
        /// </summary>
        [Input("activate", required: true)]
        public Input<bool> Activate { get; set; } = null!;

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Input("detectorId", required: true)]
        public Input<string> DetectorId { get; set; } = null!;

        /// <summary>
        /// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// The URI of the file that contains the IPSet.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The friendly name to identify the IPSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IPSetArgs()
        {
        }
    }

    public sealed class IPSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether GuardDuty is to start using the uploaded IPSet.
        /// </summary>
        [Input("activate")]
        public Input<bool>? Activate { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the GuardDuty IPSet.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The detector ID of the GuardDuty.
        /// </summary>
        [Input("detectorId")]
        public Input<string>? DetectorId { get; set; }

        /// <summary>
        /// The format of the file that contains the IPSet. Valid values: `TXT` | `STIX` | `OTX_CSV` | `ALIEN_VAULT` | `PROOF_POINT` | `FIRE_EYE`
        /// </summary>
        [Input("format")]
        public Input<string>? Format { get; set; }

        /// <summary>
        /// The URI of the file that contains the IPSet.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The friendly name to identify the IPSet.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public IPSetState()
        {
        }
    }
}
