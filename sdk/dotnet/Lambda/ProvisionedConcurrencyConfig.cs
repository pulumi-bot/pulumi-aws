// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Lambda
{
    /// <summary>
    /// Manages a Lambda Provisioned Concurrency Configuration.
    /// 
    /// ## Example Usage
    /// ### Alias Name
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.ProvisionedConcurrencyConfig("example", new Aws.Lambda.ProvisionedConcurrencyConfigArgs
    ///         {
    ///             FunctionName = aws_lambda_alias.Example.Function_name,
    ///             ProvisionedConcurrentExecutions = 1,
    ///             Qualifier = aws_lambda_alias.Example.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Function Version
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.Lambda.ProvisionedConcurrencyConfig("example", new Aws.Lambda.ProvisionedConcurrencyConfigArgs
    ///         {
    ///             FunctionName = aws_lambda_function.Example.Function_name,
    ///             ProvisionedConcurrentExecutions = 1,
    ///             Qualifier = aws_lambda_function.Example.Version,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Lambda Provisioned Concurrency Configs can be imported using the `function_name` and `qualifier` separated by a colon (`:`), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig example my_function:production
    /// ```
    /// </summary>
    [AwsResourceType("aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig")]
    public partial class ProvisionedConcurrencyConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the Lambda Function.
        /// </summary>
        [Output("functionName")]
        public Output<string> FunctionName { get; private set; } = null!;

        /// <summary>
        /// Amount of capacity to allocate. Must be greater than or equal to `1`.
        /// </summary>
        [Output("provisionedConcurrentExecutions")]
        public Output<int> ProvisionedConcurrentExecutions { get; private set; } = null!;

        /// <summary>
        /// Lambda Function version or Lambda Alias name.
        /// </summary>
        [Output("qualifier")]
        public Output<string> Qualifier { get; private set; } = null!;


        /// <summary>
        /// Create a ProvisionedConcurrencyConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProvisionedConcurrencyConfig(string name, ProvisionedConcurrencyConfigArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig", name, args ?? new ProvisionedConcurrencyConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProvisionedConcurrencyConfig(string name, Input<string> id, ProvisionedConcurrencyConfigState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/provisionedConcurrencyConfig:ProvisionedConcurrencyConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProvisionedConcurrencyConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProvisionedConcurrencyConfig Get(string name, Input<string> id, ProvisionedConcurrencyConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new ProvisionedConcurrencyConfig(name, id, state, options);
        }
    }

    public sealed class ProvisionedConcurrencyConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the Lambda Function.
        /// </summary>
        [Input("functionName", required: true)]
        public Input<string> FunctionName { get; set; } = null!;

        /// <summary>
        /// Amount of capacity to allocate. Must be greater than or equal to `1`.
        /// </summary>
        [Input("provisionedConcurrentExecutions", required: true)]
        public Input<int> ProvisionedConcurrentExecutions { get; set; } = null!;

        /// <summary>
        /// Lambda Function version or Lambda Alias name.
        /// </summary>
        [Input("qualifier", required: true)]
        public Input<string> Qualifier { get; set; } = null!;

        public ProvisionedConcurrencyConfigArgs()
        {
        }
    }

    public sealed class ProvisionedConcurrencyConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name or Amazon Resource Name (ARN) of the Lambda Function.
        /// </summary>
        [Input("functionName")]
        public Input<string>? FunctionName { get; set; }

        /// <summary>
        /// Amount of capacity to allocate. Must be greater than or equal to `1`.
        /// </summary>
        [Input("provisionedConcurrentExecutions")]
        public Input<int>? ProvisionedConcurrentExecutions { get; set; }

        /// <summary>
        /// Lambda Function version or Lambda Alias name.
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        public ProvisionedConcurrencyConfigState()
        {
        }
    }
}
