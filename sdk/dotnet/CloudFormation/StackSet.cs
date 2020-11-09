// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFormation
{
    /// <summary>
    /// Manages a CloudFormation StackSet. StackSets allow CloudFormation templates to be easily deployed across multiple accounts and regions via StackSet Instances (`aws.cloudformation.StackSetInstance` resource). Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).
    /// 
    /// &gt; **NOTE:** All template parameters, including those with a `Default`, must be configured or ignored with the `lifecycle` configuration block `ignore_changes` argument.
    /// 
    /// &gt; **NOTE:** All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
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
    ///         var aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = Output.Create(Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "sts:AssumeRole",
    ///                     },
    ///                     Effect = "Allow",
    ///                     Principals = 
    ///                     {
    ///                         new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalArgs
    ///                         {
    ///                             Identifiers = 
    ///                             {
    ///                                 "cloudformation.amazonaws.com",
    ///                             },
    ///                             Type = "Service",
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var aWSCloudFormationStackSetAdministrationRole = new Aws.Iam.Role("aWSCloudFormationStackSetAdministrationRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Apply(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy =&gt; aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Json),
    ///         });
    ///         var example = new Aws.CloudFormation.StackSet("example", new Aws.CloudFormation.StackSetArgs
    ///         {
    ///             AdministrationRoleArn = aWSCloudFormationStackSetAdministrationRole.Arn,
    ///             Parameters = 
    ///             {
    ///                 { "VPCCidr", "10.0.0.0/16" },
    ///             },
    ///             TemplateBody = @"{
    ///   ""Parameters"" : {
    ///     ""VPCCidr"" : {
    ///       ""Type"" : ""String"",
    ///       ""Default"" : ""10.0.0.0/16"",
    ///       ""Description"" : ""Enter the CIDR block for the VPC. Default is 10.0.0.0/16.""
    ///     }
    ///   },
    ///   ""Resources"" : {
    ///     ""myVpc"": {
    ///       ""Type"" : ""AWS::EC2::VPC"",
    ///       ""Properties"" : {
    ///         ""CidrBlock"" : { ""Ref"" : ""VPCCidr"" },
    ///         ""Tags"" : [
    ///           {""Key"": ""Name"", ""Value"": ""Primary_CF_VPC""}
    ///         ]
    ///       }
    ///     }
    ///   }
    /// }
    /// ",
    ///         });
    ///         var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = example.ExecutionRoleName.Apply(executionRoleName =&gt; Aws.Iam.GetPolicyDocument.InvokeAsync(new Aws.Iam.GetPolicyDocumentArgs
    ///         {
    ///             Statements = 
    ///             {
    ///                 new Aws.Iam.Inputs.GetPolicyDocumentStatementArgs
    ///                 {
    ///                     Actions = 
    ///                     {
    ///                         "sts:AssumeRole",
    ///                     },
    ///                     Effect = "Allow",
    ///                     Resources = 
    ///                     {
    ///                         $"arn:aws:iam::*:role/{executionRoleName}",
    ///                     },
    ///                 },
    ///             },
    ///         }));
    ///         var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new Aws.Iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", new Aws.Iam.RolePolicyArgs
    ///         {
    ///             Policy = aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Apply(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument =&gt; aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Json),
    ///             Role = aWSCloudFormationStackSetAdministrationRole.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CloudFormation StackSets can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:cloudformation/stackSet:StackSet example example
    /// ```
    /// </summary>
    public partial class StackSet : Pulumi.CustomResource
    {
        /// <summary>
        /// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
        /// </summary>
        [Output("administrationRoleArn")]
        public Output<string> AdministrationRoleArn { get; private set; } = null!;

        /// <summary>
        /// Amazon Resource Name (ARN) of the StackSet.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
        /// </summary>
        [Output("capabilities")]
        public Output<ImmutableArray<string>> Capabilities { get; private set; } = null!;

        /// <summary>
        /// Description of the StackSet.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
        /// </summary>
        [Output("executionRoleName")]
        public Output<string?> ExecutionRoleName { get; private set; } = null!;

        /// <summary>
        /// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the StackSet.
        /// </summary>
        [Output("stackSetId")]
        public Output<string> StackSetId { get; private set; } = null!;

        /// <summary>
        /// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
        /// </summary>
        [Output("templateBody")]
        public Output<string> TemplateBody { get; private set; } = null!;

        /// <summary>
        /// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
        /// </summary>
        [Output("templateUrl")]
        public Output<string?> TemplateUrl { get; private set; } = null!;


        /// <summary>
        /// Create a StackSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StackSet(string name, StackSetArgs args, CustomResourceOptions? options = null)
            : base("aws:cloudformation/stackSet:StackSet", name, args ?? new StackSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StackSet(string name, Input<string> id, StackSetState? state = null, CustomResourceOptions? options = null)
            : base("aws:cloudformation/stackSet:StackSet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing StackSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StackSet Get(string name, Input<string> id, StackSetState? state = null, CustomResourceOptions? options = null)
        {
            return new StackSet(name, id, state, options);
        }
    }

    public sealed class StackSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
        /// </summary>
        [Input("administrationRoleArn", required: true)]
        public Input<string> AdministrationRoleArn { get; set; } = null!;

        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// Description of the StackSet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
        /// </summary>
        [Input("executionRoleName")]
        public Input<string>? ExecutionRoleName { get; set; }

        /// <summary>
        /// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
        /// </summary>
        [Input("templateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
        /// </summary>
        [Input("templateUrl")]
        public Input<string>? TemplateUrl { get; set; }

        public StackSetArgs()
        {
        }
    }

    public sealed class StackSetState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Amazon Resource Number (ARN) of the IAM Role in the administrator account.
        /// </summary>
        [Input("administrationRoleArn")]
        public Input<string>? AdministrationRoleArn { get; set; }

        /// <summary>
        /// Amazon Resource Name (ARN) of the StackSet.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// A list of capabilities. Valid values: `CAPABILITY_IAM`, `CAPABILITY_NAMED_IAM`, `CAPABILITY_AUTO_EXPAND`.
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// Description of the StackSet.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the IAM Role in all target accounts for StackSet operations. Defaults to `AWSCloudFormationStackSetExecutionRole`.
        /// </summary>
        [Input("executionRoleName")]
        public Input<string>? ExecutionRoleName { get; set; }

        /// <summary>
        /// Name of the StackSet. The name must be unique in the region where you create your StackSet. The name can contain only alphanumeric characters (case-sensitive) and hyphens. It must start with an alphabetic character and cannot be longer than 128 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Key-value map of input parameters for the StackSet template. All template parameters, including those with a `Default`, must be configured or ignored with `lifecycle` configuration block `ignore_changes` argument. All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Unique identifier of the StackSet.
        /// </summary>
        [Input("stackSetId")]
        public Input<string>? StackSetId { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value map of tags to associate with this StackSet and the Stacks created from it. AWS CloudFormation also propagates these tags to supported resources that are created in the Stacks. A maximum number of 50 tags can be specified.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// String containing the CloudFormation template body. Maximum size: 51,200 bytes. Conflicts with `template_url`.
        /// </summary>
        [Input("templateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// String containing the location of a file containing the CloudFormation template body. The URL must point to a template that is located in an Amazon S3 bucket. Maximum location file size: 460,800 bytes. Conflicts with `template_body`.
        /// </summary>
        [Input("templateUrl")]
        public Input<string>? TemplateUrl { get; set; }

        public StackSetState()
        {
        }
    }
}
