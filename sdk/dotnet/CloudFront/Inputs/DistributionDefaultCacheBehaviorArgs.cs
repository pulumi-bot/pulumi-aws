// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.CloudFront.Inputs
{

    public sealed class DistributionDefaultCacheBehaviorArgs : Pulumi.ResourceArgs
    {
        [Input("allowedMethods", required: true)]
        private InputList<string>? _allowedMethods;
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        [Input("cachedMethods", required: true)]
        private InputList<string>? _cachedMethods;
        public InputList<string> CachedMethods
        {
            get => _cachedMethods ?? (_cachedMethods = new InputList<string>());
            set => _cachedMethods = value;
        }

        [Input("compress")]
        public Input<bool>? Compress { get; set; }

        [Input("defaultTtl")]
        public Input<int>? DefaultTtl { get; set; }

        [Input("fieldLevelEncryptionId")]
        public Input<string>? FieldLevelEncryptionId { get; set; }

        [Input("forwardedValues", required: true)]
        public Input<Inputs.DistributionDefaultCacheBehaviorForwardedValuesArgs> ForwardedValues { get; set; } = null!;

        [Input("lambdaFunctionAssociations")]
        private InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationArgs>? _lambdaFunctionAssociations;
        public InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationArgs> LambdaFunctionAssociations
        {
            get => _lambdaFunctionAssociations ?? (_lambdaFunctionAssociations = new InputList<Inputs.DistributionDefaultCacheBehaviorLambdaFunctionAssociationArgs>());
            set => _lambdaFunctionAssociations = value;
        }

        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        [Input("minTtl")]
        public Input<int>? MinTtl { get; set; }

        [Input("smoothStreaming")]
        public Input<bool>? SmoothStreaming { get; set; }

        [Input("targetOriginId", required: true)]
        public Input<string> TargetOriginId { get; set; } = null!;

        [Input("trustedSigners")]
        private InputList<string>? _trustedSigners;
        public InputList<string> TrustedSigners
        {
            get => _trustedSigners ?? (_trustedSigners = new InputList<string>());
            set => _trustedSigners = value;
        }

        [Input("viewerProtocolPolicy", required: true)]
        public Input<string> ViewerProtocolPolicy { get; set; } = null!;

        public DistributionDefaultCacheBehaviorArgs()
        {
        }
    }
}
