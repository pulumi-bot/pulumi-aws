// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.IdentityStore
{
    public static class GetUser
    {
        /// <summary>
        /// Use this data source to get an Identity Store User.
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("aws:identitystore/getUser:getUser", args ?? new GetUserArgs(), options.WithVersion());

        public static Output<GetUserResult> Invoke(GetUserOutputArgs args, InvokeOptions? options = null)
        {
            return Pulumi.Output.All(
                args.Filters.ToList().Box(),
                args.IdentityStoreId.Box(),
                args.UserId.Box()
            ).Apply(a => {
                    var args = new GetUserArgs();
                    a[0].Set(args, nameof(args.Filters));
                    a[1].Set(args, nameof(args.IdentityStoreId));
                    a[2].Set(args, nameof(args.UserId));
                    return InvokeAsync(args, options);
            });
        }
    }


    public sealed class GetUserArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private List<Inputs.GetUserFilterArgs>? _filters;

        /// <summary>
        /// Configuration block(s) for filtering. Currently, the AWS Identity Store API supports only 1 filter. Detailed below.
        /// </summary>
        public List<Inputs.GetUserFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetUserFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The Identity Store ID associated with the Single Sign-On Instance.
        /// </summary>
        [Input("identityStoreId", required: true)]
        public string IdentityStoreId { get; set; } = null!;

        /// <summary>
        /// The identifier for a user in the Identity Store.
        /// </summary>
        [Input("userId")]
        public string? UserId { get; set; }

        public GetUserArgs()
        {
        }
    }

    public sealed class GetUserOutputArgs
    {
        [Input("filters", required: true)]
        private InputList<Inputs.GetUserFilterArgs>? _filters;

        /// <summary>
        /// Configuration block(s) for filtering. Currently, the AWS Identity Store API supports only 1 filter. Detailed below.
        /// </summary>
        public InputList<Inputs.GetUserFilterArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetUserFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The Identity Store ID associated with the Single Sign-On Instance.
        /// </summary>
        [Input("identityStoreId", required: true)]
        public Input<string> IdentityStoreId { get; set; } = null!;

        /// <summary>
        /// The identifier for a user in the Identity Store.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GetUserOutputArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserResult
    {
        public readonly ImmutableArray<Outputs.GetUserFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IdentityStoreId;
        public readonly string UserId;
        /// <summary>
        /// The user's user name value.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetUserResult(
            ImmutableArray<Outputs.GetUserFilterResult> filters,

            string id,

            string identityStoreId,

            string userId,

            string userName)
        {
            Filters = filters;
            Id = id;
            IdentityStoreId = identityStoreId;
            UserId = userId;
            UserName = userName;
        }
    }
}
