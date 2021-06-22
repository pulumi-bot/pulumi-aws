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
    /// Gives an external source (like a CloudWatch Event Rule, SNS, or S3) permission to access the Lambda function.
    /// 
    /// ## Example Usage
    /// ### Basic Example
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var iamForLambda = new Aws.Iam.Role("iamForLambda", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2012-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "sts:AssumeRole" },
    ///                             { "Effect", "Allow" },
    ///                             { "Sid", "" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "lambda.amazonaws.com" },
    ///                             } },
    ///                         },
    ///                     }
    ///                  },
    ///             }),
    ///         });
    ///         var testLambda = new Aws.Lambda.Function("testLambda", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("lambdatest.zip"),
    ///             Role = iamForLambda.Arn,
    ///             Handler = "exports.handler",
    ///             Runtime = "nodejs12.x",
    ///         });
    ///         var testAlias = new Aws.Lambda.Alias("testAlias", new Aws.Lambda.AliasArgs
    ///         {
    ///             Description = "a sample description",
    ///             FunctionName = testLambda.Name,
    ///             FunctionVersion = "$LATEST",
    ///         });
    ///         var allowCloudwatch = new Aws.Lambda.Permission("allowCloudwatch", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = testLambda.Name,
    ///             Principal = "events.amazonaws.com",
    ///             SourceArn = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
    ///             Qualifier = testAlias.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Usage with SNS
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultTopic = new Aws.Sns.Topic("defaultTopic", new Aws.Sns.TopicArgs
    ///         {
    ///         });
    ///         var defaultRole = new Aws.Iam.Role("defaultRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2012-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Action", "sts:AssumeRole" },
    ///                             { "Effect", "Allow" },
    ///                             { "Sid", "" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "Service", "lambda.amazonaws.com" },
    ///                             } },
    ///                         },
    ///                     }
    ///                  },
    ///             }),
    ///         });
    ///         var func = new Aws.Lambda.Function("func", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("lambdatest.zip"),
    ///             Role = defaultRole.Arn,
    ///             Handler = "exports.handler",
    ///             Runtime = "python2.7",
    ///         });
    ///         var withSns = new Aws.Lambda.Permission("withSns", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = func.Name,
    ///             Principal = "sns.amazonaws.com",
    ///             SourceArn = defaultTopic.Arn,
    ///         });
    ///         var lambda = new Aws.Sns.TopicSubscription("lambda", new Aws.Sns.TopicSubscriptionArgs
    ///         {
    ///             Topic = defaultTopic.Arn,
    ///             Protocol = "lambda",
    ///             Endpoint = func.Arn,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Specify Lambda permissions for API Gateway REST API
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new Aws.ApiGateway.RestApiArgs
    ///         {
    ///             Description = "This is my API for demonstration purposes",
    ///         });
    ///         var lambdaPermission = new Aws.Lambda.Permission("lambdaPermission", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = "MyDemoFunction",
    ///             Principal = "apigateway.amazonaws.com",
    ///             SourceArn = myDemoAPI.ExecutionArn.Apply(executionArn =&gt; $"{executionArn}/*/*/*"),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ## Usage with CloudWatch log group
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var defaultLogGroup = new Aws.CloudWatch.LogGroup("defaultLogGroup", new Aws.CloudWatch.LogGroupArgs
    ///         {
    ///         });
    ///         var defaultRole = new Aws.Iam.Role("defaultRole", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""lambda.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow"",
    ///       ""Sid"": """"
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var loggingFunction = new Aws.Lambda.Function("loggingFunction", new Aws.Lambda.FunctionArgs
    ///         {
    ///             Code = new FileArchive("lamba_logging.zip"),
    ///             Handler = "exports.handler",
    ///             Role = defaultRole.Arn,
    ///             Runtime = "python2.7",
    ///         });
    ///         var loggingPermission = new Aws.Lambda.Permission("loggingPermission", new Aws.Lambda.PermissionArgs
    ///         {
    ///             Action = "lambda:InvokeFunction",
    ///             Function = loggingFunction.Name,
    ///             Principal = "logs.eu-west-1.amazonaws.com",
    ///             SourceArn = defaultLogGroup.Arn.Apply(arn =&gt; $"{arn}:*"),
    ///         });
    ///         var loggingLogSubscriptionFilter = new Aws.CloudWatch.LogSubscriptionFilter("loggingLogSubscriptionFilter", new Aws.CloudWatch.LogSubscriptionFilterArgs
    ///         {
    ///             DestinationArn = loggingFunction.Arn,
    ///             FilterPattern = "",
    ///             LogGroup = defaultLogGroup.Name,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 loggingPermission,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Lambda permission statements can be imported using function_name/statement_id, with an optional qualifier, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function/AllowExecutionFromCloudWatch
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function:qualifier_name/AllowExecutionFromCloudWatch
    /// ```
    /// </summary>
    [AwsResourceType("aws:lambda/permission:Permission")]
    public partial class Permission : Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        /// </summary>
        [Output("eventSourceToken")]
        public Output<string?> EventSourceToken { get; private set; } = null!;

        /// <summary>
        /// Name of the Lambda function whose resource policy you are updating
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// The principal who is getting this permission. e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARN. e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        /// </summary>
        [Output("qualifier")]
        public Output<string?> Qualifier { get; private set; } = null!;

        /// <summary>
        /// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        /// </summary>
        [Output("sourceAccount")]
        public Output<string?> SourceAccount { get; private set; } = null!;

        /// <summary>
        /// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
        /// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
        /// For S3, this should be the ARN of the S3 Bucket.
        /// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
        /// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        /// </summary>
        [Output("sourceArn")]
        public Output<string?> SourceArn { get; private set; } = null!;

        /// <summary>
        /// A unique statement identifier. By default generated by this provider.
        /// </summary>
        [Output("statementId")]
        public Output<string> StatementId { get; private set; } = null!;

        /// <summary>
        /// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
        /// </summary>
        [Output("statementIdPrefix")]
        public Output<string?> StatementIdPrefix { get; private set; } = null!;


        /// <summary>
        /// Create a Permission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permission(string name, PermissionArgs args, CustomResourceOptions? options = null)
            : base("aws:lambda/permission:Permission", name, args ?? new PermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Permission(string name, Input<string> id, PermissionState? state = null, CustomResourceOptions? options = null)
            : base("aws:lambda/permission:Permission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Permission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permission Get(string name, Input<string> id, PermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new Permission(name, id, state, options);
        }
    }

    public sealed class PermissionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        /// </summary>
        [Input("eventSourceToken")]
        public Input<string>? EventSourceToken { get; set; }

        /// <summary>
        /// Name of the Lambda function whose resource policy you are updating
        /// </summary>
        [Input("function", required: true)]
        public string Function { get; set; } = null!;

        /// <summary>
        /// The principal who is getting this permission. e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARN. e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        /// </summary>
        [Input("sourceAccount")]
        public Input<string>? SourceAccount { get; set; }

        /// <summary>
        /// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
        /// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
        /// For S3, this should be the ARN of the S3 Bucket.
        /// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
        /// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// A unique statement identifier. By default generated by this provider.
        /// </summary>
        [Input("statementId")]
        public Input<string>? StatementId { get; set; }

        /// <summary>
        /// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
        /// </summary>
        [Input("statementIdPrefix")]
        public Input<string>? StatementIdPrefix { get; set; }

        public PermissionArgs()
        {
        }
    }

    public sealed class PermissionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The AWS Lambda action you want to allow in this statement. (e.g. `lambda:InvokeFunction`)
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// The Event Source Token to validate.  Used with [Alexa Skills](https://developer.amazon.com/docs/custom-skills/host-a-custom-skill-as-an-aws-lambda-function.html#use-aws-cli).
        /// </summary>
        [Input("eventSourceToken")]
        public Input<string>? EventSourceToken { get; set; }

        /// <summary>
        /// Name of the Lambda function whose resource policy you are updating
        /// </summary>
        [Input("function")]
        public string? Function { get; set; }

        /// <summary>
        /// The principal who is getting this permission. e.g. `s3.amazonaws.com`, an AWS account ID, or any valid AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// Query parameter to specify function version or alias name. The permission will then apply to the specific qualified ARN. e.g. `arn:aws:lambda:aws-region:acct-id:function:function-name:2`
        /// </summary>
        [Input("qualifier")]
        public Input<string>? Qualifier { get; set; }

        /// <summary>
        /// This parameter is used for S3 and SES. The AWS account ID (without a hyphen) of the source owner.
        /// </summary>
        [Input("sourceAccount")]
        public Input<string>? SourceAccount { get; set; }

        /// <summary>
        /// When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
        /// Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
        /// For S3, this should be the ARN of the S3 Bucket.
        /// For CloudWatch Events, this should be the ARN of the CloudWatch Events Rule.
        /// For API Gateway, this should be the ARN of the API, as described [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
        /// </summary>
        [Input("sourceArn")]
        public Input<string>? SourceArn { get; set; }

        /// <summary>
        /// A unique statement identifier. By default generated by this provider.
        /// </summary>
        [Input("statementId")]
        public Input<string>? StatementId { get; set; }

        /// <summary>
        /// A statement identifier prefix. This provider will generate a unique suffix. Conflicts with `statement_id`.
        /// </summary>
        [Input("statementIdPrefix")]
        public Input<string>? StatementIdPrefix { get; set; }

        public PermissionState()
        {
        }
    }
}
