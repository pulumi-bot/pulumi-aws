// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AppSync.Outputs
{

    [OutputType]
    public sealed class DataSourceDynamodbConfig
    {
        /// <summary>
        /// AWS region of the DynamoDB table. Defaults to current region.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// Name of the DynamoDB table.
        /// </summary>
        public readonly string TableName;
        /// <summary>
        /// Set to `true` to use Amazon Cognito credentials with this data source.
        /// </summary>
        public readonly bool? UseCallerCredentials;

        [OutputConstructor]
        private DataSourceDynamodbConfig(
            string? region,

            string tableName,

            bool? useCallerCredentials)
        {
            Region = region;
            TableName = tableName;
            UseCallerCredentials = useCallerCredentials;
        }
    }
}
