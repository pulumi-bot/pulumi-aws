// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dms
{
    /// <summary>
    /// Provides a DMS (Data Migration Service) endpoint resource. DMS endpoints can be created, updated, deleted, and imported.
    /// 
    /// &gt; **Note:** All arguments including the password will be stored in the raw state as plain-text.
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
    ///         // Create a new endpoint
    ///         var test = new Aws.Dms.Endpoint("test", new Aws.Dms.EndpointArgs
    ///         {
    ///             CertificateArn = "arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012",
    ///             DatabaseName = "test",
    ///             EndpointId = "test-dms-endpoint-tf",
    ///             EndpointType = "source",
    ///             EngineName = "aurora",
    ///             ExtraConnectionAttributes = "",
    ///             KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    ///             Password = "test",
    ///             Port = 3306,
    ///             ServerName = "test",
    ///             SslMode = "none",
    ///             Tags = 
    ///             {
    ///                 { "Name", "test" },
    ///             },
    ///             Username = "test",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Endpoints can be imported using the `endpoint_id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:dms/endpoint:Endpoint test test-dms-endpoint-tf
    /// ```
    /// </summary>
    public partial class Endpoint : Pulumi.CustomResource
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the certificate.
        /// </summary>
        [Output("certificateArn")]
        public Output<string> CertificateArn { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint database.
        /// </summary>
        [Output("databaseName")]
        public Output<string?> DatabaseName { get; private set; } = null!;

        /// <summary>
        /// Configuration block with Elasticsearch settings. Detailed below.
        /// </summary>
        [Output("elasticsearchSettings")]
        public Output<Outputs.EndpointElasticsearchSettings?> ElasticsearchSettings { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the endpoint.
        /// </summary>
        [Output("endpointArn")]
        public Output<string> EndpointArn { get; private set; } = null!;

        /// <summary>
        /// The database endpoint identifier.
        /// </summary>
        [Output("endpointId")]
        public Output<string> EndpointId { get; private set; } = null!;

        /// <summary>
        /// The type of endpoint. Can be one of `source | target`.
        /// </summary>
        [Output("endpointType")]
        public Output<string> EndpointType { get; private set; } = null!;

        /// <summary>
        /// The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        /// </summary>
        [Output("engineName")]
        public Output<string> EngineName { get; private set; } = null!;

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        /// </summary>
        [Output("extraConnectionAttributes")]
        public Output<string> ExtraConnectionAttributes { get; private set; } = null!;

        /// <summary>
        /// Configuration block with Kafka settings. Detailed below.
        /// </summary>
        [Output("kafkaSettings")]
        public Output<Outputs.EndpointKafkaSettings?> KafkaSettings { get; private set; } = null!;

        /// <summary>
        /// Configuration block with Kinesis settings. Detailed below.
        /// </summary>
        [Output("kinesisSettings")]
        public Output<Outputs.EndpointKinesisSettings?> KinesisSettings { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        /// </summary>
        [Output("kmsKeyArn")]
        public Output<string> KmsKeyArn { get; private set; } = null!;

        /// <summary>
        /// Configuration block with MongoDB settings. Detailed below.
        /// </summary>
        [Output("mongodbSettings")]
        public Output<Outputs.EndpointMongodbSettings?> MongodbSettings { get; private set; } = null!;

        /// <summary>
        /// The password to be used to login to the endpoint database.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// The port used by the endpoint database.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// Configuration block with S3 settings. Detailed below.
        /// </summary>
        [Output("s3Settings")]
        public Output<Outputs.EndpointS3Settings?> S3Settings { get; private set; } = null!;

        /// <summary>
        /// The host name of the server.
        /// </summary>
        [Output("serverName")]
        public Output<string?> ServerName { get; private set; } = null!;

        /// <summary>
        /// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Output("serviceAccessRole")]
        public Output<string?> ServiceAccessRole { get; private set; } = null!;

        /// <summary>
        /// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        /// </summary>
        [Output("sslMode")]
        public Output<string> SslMode { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The user name to be used to login to the endpoint database.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;


        /// <summary>
        /// Create a Endpoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Endpoint(string name, EndpointArgs args, CustomResourceOptions? options = null)
            : base("aws:dms/endpoint:Endpoint", name, args ?? new EndpointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Endpoint(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
            : base("aws:dms/endpoint:Endpoint", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Endpoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Endpoint Get(string name, Input<string> id, EndpointState? state = null, CustomResourceOptions? options = null)
        {
            return new Endpoint(name, id, state, options);
        }
    }

    public sealed class EndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The name of the endpoint database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Configuration block with Elasticsearch settings. Detailed below.
        /// </summary>
        [Input("elasticsearchSettings")]
        public Input<Inputs.EndpointElasticsearchSettingsArgs>? ElasticsearchSettings { get; set; }

        /// <summary>
        /// The database endpoint identifier.
        /// </summary>
        [Input("endpointId", required: true)]
        public Input<string> EndpointId { get; set; } = null!;

        /// <summary>
        /// The type of endpoint. Can be one of `source | target`.
        /// </summary>
        [Input("endpointType", required: true)]
        public Input<string> EndpointType { get; set; } = null!;

        /// <summary>
        /// The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        /// </summary>
        [Input("engineName", required: true)]
        public Input<string> EngineName { get; set; } = null!;

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        /// </summary>
        [Input("extraConnectionAttributes")]
        public Input<string>? ExtraConnectionAttributes { get; set; }

        /// <summary>
        /// Configuration block with Kafka settings. Detailed below.
        /// </summary>
        [Input("kafkaSettings")]
        public Input<Inputs.EndpointKafkaSettingsArgs>? KafkaSettings { get; set; }

        /// <summary>
        /// Configuration block with Kinesis settings. Detailed below.
        /// </summary>
        [Input("kinesisSettings")]
        public Input<Inputs.EndpointKinesisSettingsArgs>? KinesisSettings { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Configuration block with MongoDB settings. Detailed below.
        /// </summary>
        [Input("mongodbSettings")]
        public Input<Inputs.EndpointMongodbSettingsArgs>? MongodbSettings { get; set; }

        /// <summary>
        /// The password to be used to login to the endpoint database.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The port used by the endpoint database.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Configuration block with S3 settings. Detailed below.
        /// </summary>
        [Input("s3Settings")]
        public Input<Inputs.EndpointS3SettingsArgs>? S3Settings { get; set; }

        /// <summary>
        /// The host name of the server.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Input("serviceAccessRole")]
        public Input<string>? ServiceAccessRole { get; set; }

        /// <summary>
        /// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The user name to be used to login to the endpoint database.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public EndpointArgs()
        {
        }
    }

    public sealed class EndpointState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Amazon Resource Name (ARN) for the certificate.
        /// </summary>
        [Input("certificateArn")]
        public Input<string>? CertificateArn { get; set; }

        /// <summary>
        /// The name of the endpoint database.
        /// </summary>
        [Input("databaseName")]
        public Input<string>? DatabaseName { get; set; }

        /// <summary>
        /// Configuration block with Elasticsearch settings. Detailed below.
        /// </summary>
        [Input("elasticsearchSettings")]
        public Input<Inputs.EndpointElasticsearchSettingsGetArgs>? ElasticsearchSettings { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the endpoint.
        /// </summary>
        [Input("endpointArn")]
        public Input<string>? EndpointArn { get; set; }

        /// <summary>
        /// The database endpoint identifier.
        /// </summary>
        [Input("endpointId")]
        public Input<string>? EndpointId { get; set; }

        /// <summary>
        /// The type of endpoint. Can be one of `source | target`.
        /// </summary>
        [Input("endpointType")]
        public Input<string>? EndpointType { get; set; }

        /// <summary>
        /// The type of engine for the endpoint. Can be one of `aurora | aurora-postgresql| azuredb | db2 | docdb | dynamodb | elasticsearch | kafka | kinesis | mariadb | mongodb | mysql | oracle | postgres | redshift | s3 | sqlserver | sybase`.
        /// </summary>
        [Input("engineName")]
        public Input<string>? EngineName { get; set; }

        /// <summary>
        /// Additional attributes associated with the connection. For available attributes see [Using Extra Connection Attributes with AWS Database Migration Service](http://docs.aws.amazon.com/dms/latest/userguide/CHAP_Introduction.ConnectionAttributes.html).
        /// </summary>
        [Input("extraConnectionAttributes")]
        public Input<string>? ExtraConnectionAttributes { get; set; }

        /// <summary>
        /// Configuration block with Kafka settings. Detailed below.
        /// </summary>
        [Input("kafkaSettings")]
        public Input<Inputs.EndpointKafkaSettingsGetArgs>? KafkaSettings { get; set; }

        /// <summary>
        /// Configuration block with Kinesis settings. Detailed below.
        /// </summary>
        [Input("kinesisSettings")]
        public Input<Inputs.EndpointKinesisSettingsGetArgs>? KinesisSettings { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for `kms_key_arn`, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
        /// </summary>
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        /// <summary>
        /// Configuration block with MongoDB settings. Detailed below.
        /// </summary>
        [Input("mongodbSettings")]
        public Input<Inputs.EndpointMongodbSettingsGetArgs>? MongodbSettings { get; set; }

        /// <summary>
        /// The password to be used to login to the endpoint database.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The port used by the endpoint database.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Configuration block with S3 settings. Detailed below.
        /// </summary>
        [Input("s3Settings")]
        public Input<Inputs.EndpointS3SettingsGetArgs>? S3Settings { get; set; }

        /// <summary>
        /// The host name of the server.
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// The Amazon Resource Name (ARN) used by the service access IAM role for dynamodb endpoints.
        /// </summary>
        [Input("serviceAccessRole")]
        public Input<string>? ServiceAccessRole { get; set; }

        /// <summary>
        /// The SSL mode to use for the connection. Can be one of `none | require | verify-ca | verify-full`
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The user name to be used to login to the endpoint database.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public EndpointState()
        {
        }
    }
}
