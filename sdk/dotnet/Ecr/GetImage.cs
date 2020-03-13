// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ecr
{
    public static partial class Invokes
    {
        /// <summary>
        /// The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecr_image.html.markdown.
        /// </summary>
        [Obsolete("Use GetImage.InvokeAsync() instead")]
        public static Task<GetImageResult> GetImage(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("aws:ecr/getImage:getImage", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetImage
    {
        /// <summary>
        /// The ECR Image data source allows the details of an image with a particular tag or digest to be retrieved.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/ecr_image.html.markdown.
        /// </summary>
        public static Task<GetImageResult> InvokeAsync(GetImageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetImageResult>("aws:ecr/getImage:getImage", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetImageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The sha256 digest of the image manifest. At least one of `image_digest` or `image_tag` must be specified.
        /// </summary>
        [Input("imageDigest")]
        public string? ImageDigest { get; set; }

        /// <summary>
        /// The tag associated with this image. At least one of `image_digest` or `image_tag` must be specified.
        /// </summary>
        [Input("imageTag")]
        public string? ImageTag { get; set; }

        /// <summary>
        /// The ID of the Registry where the repository resides.
        /// </summary>
        [Input("registryId")]
        public string? RegistryId { get; set; }

        /// <summary>
        /// The name of the ECR Repository.
        /// </summary>
        [Input("repositoryName", required: true)]
        public string RepositoryName { get; set; } = null!;

        public GetImageArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetImageResult
    {
        public readonly string ImageDigest;
        /// <summary>
        /// The date and time, expressed as a unix timestamp, at which the current image was pushed to the repository.
        /// </summary>
        public readonly int ImagePushedAt;
        /// <summary>
        /// The size, in bytes, of the image in the repository.
        /// </summary>
        public readonly int ImageSizeInBytes;
        public readonly string? ImageTag;
        /// <summary>
        /// The list of tags associated with this image.
        /// </summary>
        public readonly ImmutableArray<string> ImageTags;
        public readonly string RegistryId;
        public readonly string RepositoryName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetImageResult(
            string imageDigest,
            int imagePushedAt,
            int imageSizeInBytes,
            string? imageTag,
            ImmutableArray<string> imageTags,
            string registryId,
            string repositoryName,
            string id)
        {
            ImageDigest = imageDigest;
            ImagePushedAt = imagePushedAt;
            ImageSizeInBytes = imageSizeInBytes;
            ImageTag = imageTag;
            ImageTags = imageTags;
            RegistryId = registryId;
            RepositoryName = repositoryName;
            Id = id;
        }
    }
}
