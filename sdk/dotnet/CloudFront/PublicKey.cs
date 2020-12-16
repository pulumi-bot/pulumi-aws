// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// The following example below creates a CloudFront public key.
    /// 
    /// ```csharp
    /// using System.IO;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.CloudFront.PublicKey("example", new Aws.CloudFront.PublicKeyArgs
    ///         {
    ///             Comment = "test public key",
    ///             EncodedKey = File.ReadAllText("public_key.pem"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:cloudfront/publicKey:PublicKey")]
    public partial class PublicKey : Pulumi.CustomResource
    {
        /// <summary>
        /// Internal value used by CloudFront to allow future updates to the public key configuration.
        /// </summary>
        [Output("callerReference")]
        public Output<string> CallerReference { get; private set; } = null!;

        /// <summary>
        /// An optional comment about the public key.
        /// </summary>
        [Output("comment")]
        public Output<string?> Comment { get; private set; } = null!;

        /// <summary>
        /// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
        /// </summary>
        [Output("encodedKey")]
        public Output<string> EncodedKey { get; private set; } = null!;

        /// <summary>
        /// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name for the public key. By default generated by this provider.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name for the public key. Conflicts with `name`.
        /// </summary>
        [Output("namePrefix")]
        public Output<string> NamePrefix { get; private set; } = null!;


        /// <summary>
        /// Create a PublicKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PublicKey(string name, PublicKeyArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudfront/publicKey:PublicKey", name, args ?? new PublicKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PublicKey(string name, Input<string> id, PublicKeyState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudfront/publicKey:PublicKey", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PublicKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PublicKey Get(string name, Input<string> id, PublicKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new PublicKey(name, id, state, options);
        }
    }

    public sealed class PublicKeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional comment about the public key.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
        /// </summary>
        [Input("encodedKey", required: true)]
        public Input<string> EncodedKey { get; set; } = null!;

        /// <summary>
        /// The name for the public key. By default generated by this provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name for the public key. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        public PublicKeyArgs()
        {
        }
    }

    public sealed class PublicKeyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Internal value used by CloudFront to allow future updates to the public key configuration.
        /// </summary>
        [Input("callerReference")]
        public Input<string>? CallerReference { get; set; }

        /// <summary>
        /// An optional comment about the public key.
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// The encoded public key that you want to add to CloudFront to use with features like field-level encryption.
        /// </summary>
        [Input("encodedKey")]
        public Input<string>? EncodedKey { get; set; }

        /// <summary>
        /// The current version of the public key. For example: `E2QWRUHAPOMQZL`.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name for the public key. By default generated by this provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name for the public key. Conflicts with `name`.
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        public PublicKeyState()
        {
        }
    }
}
