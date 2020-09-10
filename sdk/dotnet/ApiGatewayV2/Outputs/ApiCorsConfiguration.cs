// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ApiGatewayV2.Outputs
{

    [OutputType]
    public sealed class ApiCorsConfiguration
    {
        public readonly bool? AllowCredentials;
        public readonly ImmutableArray<string> AllowHeaders;
        public readonly ImmutableArray<string> AllowMethods;
        public readonly ImmutableArray<string> AllowOrigins;
        public readonly ImmutableArray<string> ExposeHeaders;
        public readonly int? MaxAge;

        [OutputConstructor]
        private ApiCorsConfiguration(
            bool? allowCredentials,

            ImmutableArray<string> allowHeaders,

            ImmutableArray<string> allowMethods,

            ImmutableArray<string> allowOrigins,

            ImmutableArray<string> exposeHeaders,

            int? maxAge)
        {
            AllowCredentials = allowCredentials;
            AllowHeaders = allowHeaders;
            AllowMethods = allowMethods;
            AllowOrigins = allowOrigins;
            ExposeHeaders = exposeHeaders;
            MaxAge = maxAge;
        }
    }
}
