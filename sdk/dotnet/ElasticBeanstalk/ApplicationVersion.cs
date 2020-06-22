// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticBeanstalk
{
    /// <summary>
    /// Provides an Elastic Beanstalk Application Version Resource. Elastic Beanstalk allows
    /// you to deploy and manage applications in the AWS cloud without worrying about
    /// the infrastructure that runs those applications.
    /// 
    /// This resource creates a Beanstalk Application Version that can be deployed to a Beanstalk
    /// Environment.
    /// 
    /// &gt; **NOTE on Application Version Resource:**  When using the Application Version resource with multiple
    /// Elastic Beanstalk Environments it is possible that an error may be returned
    /// when attempting to delete an Application Version while it is still in use by a different environment.
    /// To work around this you can either create each environment in a separate AWS account or create your `aws.elasticbeanstalk.ApplicationVersion` resources with a unique names in your Elastic Beanstalk Application. For example &amp;lt;revision&amp;gt;-&amp;lt;environment&amp;gt;.
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
    ///         var defaultBucket = new Aws.S3.Bucket("defaultBucket", new Aws.S3.BucketArgs
    ///         {
    ///         });
    ///         var defaultBucketObject = new Aws.S3.BucketObject("defaultBucketObject", new Aws.S3.BucketObjectArgs
    ///         {
    ///             Bucket = defaultBucket.Id,
    ///             Key = "beanstalk/go-v1.zip",
    ///             Source = new FileAsset("go-v1.zip"),
    ///         });
    ///         var defaultApplication = new Aws.ElasticBeanstalk.Application("defaultApplication", new Aws.ElasticBeanstalk.ApplicationArgs
    ///         {
    ///             Description = "tf-test-desc",
    ///         });
    ///         var defaultApplicationVersion = new Aws.ElasticBeanstalk.ApplicationVersion("defaultApplicationVersion", new Aws.ElasticBeanstalk.ApplicationVersionArgs
    ///         {
    ///             Application = "tf-test-name",
    ///             Bucket = defaultBucket.Id,
    ///             Description = "application version",
    ///             Key = defaultBucketObject.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class ApplicationVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the Beanstalk Application the version is associated with.
        /// </summary>
        [Output("application")]
        public Output<string> Application { get; private set; } = null!;

        /// <summary>
        /// The ARN assigned by AWS for this Elastic Beanstalk Application.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// S3 bucket that contains the Application Version source bundle.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Short description of the Application Version.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// On delete, force an Application Version to be deleted when it may be in use
        /// by multiple Elastic Beanstalk Environments.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// S3 object that is the Application Version source bundle.
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// A unique name for the this Application Version.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of tags for the Elastic Beanstalk Application Version.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationVersion(string name, ApplicationVersionArgs args, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, args ?? new ApplicationVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationVersion(string name, Input<string> id, ApplicationVersionState? state = null, CustomResourceOptions? options = null)
            : base("aws:elasticbeanstalk/applicationVersion:ApplicationVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ApplicationVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationVersion Get(string name, Input<string> id, ApplicationVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationVersion(name, id, state, options);
        }
    }

    public sealed class ApplicationVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Beanstalk Application the version is associated with.
        /// </summary>
        [Input("application", required: true)]
        public Input<string> Application { get; set; } = null!;

        /// <summary>
        /// S3 bucket that contains the Application Version source bundle.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Short description of the Application Version.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// On delete, force an Application Version to be deleted when it may be in use
        /// by multiple Elastic Beanstalk Environments.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// S3 object that is the Application Version source bundle.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// A unique name for the this Application Version.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value map of tags for the Elastic Beanstalk Application Version.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ApplicationVersionArgs()
        {
        }
    }

    public sealed class ApplicationVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Beanstalk Application the version is associated with.
        /// </summary>
        [Input("application")]
        public Input<string>? Application { get; set; }

        /// <summary>
        /// The ARN assigned by AWS for this Elastic Beanstalk Application.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// S3 bucket that contains the Application Version source bundle.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Short description of the Application Version.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// On delete, force an Application Version to be deleted when it may be in use
        /// by multiple Elastic Beanstalk Environments.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// S3 object that is the Application Version source bundle.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// A unique name for the this Application Version.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Key-value map of tags for the Elastic Beanstalk Application Version.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public ApplicationVersionState()
        {
        }
    }
}
