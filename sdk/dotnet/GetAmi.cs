// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws
{
    public static class GetAmi
    {
        /// <summary>
        /// Use this data source to get the ID of a registered AMI for use in other
        /// resources.
        /// </summary>
        public static Task<GetAmiResult> InvokeAsync(GetAmiArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAmiResult>("aws:index/getAmi:getAmi", args ?? new GetAmiArgs(), options.WithVersion());
    }


    public sealed class GetAmiArgs : Pulumi.InvokeArgs
    {
        [Input("executableUsers")]
        private List<string>? _executableUsers;

        /// <summary>
        /// Limit search to users with *explicit* launch permission on
        /// the image. Valid items are the numeric account ID or `self`.
        /// </summary>
        public List<string> ExecutableUsers
        {
            get => _executableUsers ?? (_executableUsers = new List<string>());
            set => _executableUsers = value;
        }

        [Input("filters")]
        private List<Inputs.GetAmiFilterArgs>? _filters;

        /// <summary>
        /// One or more name/value pairs to filter off of. There are
        /// several valid keys, for a full reference, check out
        /// [describe-images in the AWS CLI reference][1].
        /// </summary>
        public List<Inputs.GetAmiFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetAmiFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// If more than one result is returned, use the most
        /// recent AMI.
        /// </summary>
        [Input("mostRecent")]
        public bool? MostRecent { get; set; }

        /// <summary>
        /// A regex string to apply to the AMI list returned
        /// by AWS. This allows more advanced filtering not supported from the AWS API. This
        /// filtering is done locally on what AWS returns, and could have a performance
        /// impact if the result is large. It is recommended to combine this with other
        /// options to narrow down the list AWS returns.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        [Input("owners", required: true)]
        private List<string>? _owners;

        /// <summary>
        /// List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g. `amazon`, `aws-marketplace`, `microsoft`).
        /// </summary>
        public List<string> Owners
        {
            get => _owners ?? (_owners = new List<string>());
            set => _owners = value;
        }

        [Input("tags")]
        private Dictionary<string, string>? _tags;

        /// <summary>
        /// Any tags assigned to the image.
        /// * `tags.#.key` - The key name of the tag.
        /// * `tags.#.value` - The value of the tag.
        /// </summary>
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetAmiArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAmiResult
    {
        /// <summary>
        /// The OS architecture of the AMI (ie: `i386` or `x86_64`).
        /// </summary>
        public readonly string Architecture;
        /// <summary>
        /// The ARN of the AMI.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The block device mappings of the AMI.
        /// * `block_device_mappings.#.device_name` - The physical name of the device.
        /// * `block_device_mappings.#.ebs.delete_on_termination` - `true` if the EBS volume
        /// will be deleted on termination.
        /// * `block_device_mappings.#.ebs.encrypted` - `true` if the EBS volume
        /// is encrypted.
        /// * `block_device_mappings.#.ebs.iops` - `0` if the EBS volume is
        /// not a provisioned IOPS image, otherwise the supported IOPS count.
        /// * `block_device_mappings.#.ebs.snapshot_id` - The ID of the snapshot.
        /// * `block_device_mappings.#.ebs.volume_size` - The size of the volume, in GiB.
        /// * `block_device_mappings.#.ebs.volume_type` - The volume type.
        /// * `block_device_mappings.#.no_device` - Suppresses the specified device
        /// included in the block device mapping of the AMI.
        /// * `block_device_mappings.#.virtual_name` - The virtual device name (for
        /// instance stores).
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAmiBlockDeviceMappingResult> BlockDeviceMappings;
        /// <summary>
        /// The date and time the image was created.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The description of the AMI that was provided during image
        /// creation.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<string> ExecutableUsers;
        public readonly ImmutableArray<Outputs.GetAmiFilterResult> Filters;
        /// <summary>
        /// The hypervisor type of the image.
        /// </summary>
        public readonly string Hypervisor;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the AMI. Should be the same as the resource `id`.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The location of the AMI.
        /// </summary>
        public readonly string ImageLocation;
        /// <summary>
        /// The AWS account alias (for example, `amazon`, `self`) or
        /// the AWS account ID of the AMI owner.
        /// </summary>
        public readonly string ImageOwnerAlias;
        /// <summary>
        /// The type of image.
        /// </summary>
        public readonly string ImageType;
        /// <summary>
        /// The kernel associated with the image, if any. Only applicable
        /// for machine images.
        /// </summary>
        public readonly string KernelId;
        public readonly bool? MostRecent;
        /// <summary>
        /// The name of the AMI that was provided during image creation.
        /// </summary>
        public readonly string Name;
        public readonly string? NameRegex;
        /// <summary>
        /// The AWS account ID of the image owner.
        /// </summary>
        public readonly string OwnerId;
        public readonly ImmutableArray<string> Owners;
        /// <summary>
        /// The value is Windows for `Windows` AMIs; otherwise blank.
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// Any product codes associated with the AMI.
        /// * `product_codes.#.product_code_id` - The product code.
        /// * `product_codes.#.product_code_type` - The type of product code.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAmiProductCodeResult> ProductCodes;
        /// <summary>
        /// `true` if the image has public launch permissions.
        /// </summary>
        public readonly bool Public;
        /// <summary>
        /// The RAM disk associated with the image, if any. Only applicable
        /// for machine images.
        /// </summary>
        public readonly string RamdiskId;
        /// <summary>
        /// The device name of the root device.
        /// </summary>
        public readonly string RootDeviceName;
        /// <summary>
        /// The type of root device (ie: `ebs` or `instance-store`).
        /// </summary>
        public readonly string RootDeviceType;
        /// <summary>
        /// The snapshot id associated with the root device, if any
        /// (only applies to `ebs` root devices).
        /// </summary>
        public readonly string RootSnapshotId;
        /// <summary>
        /// Specifies whether enhanced networking is enabled.
        /// </summary>
        public readonly string SriovNetSupport;
        /// <summary>
        /// The current state of the AMI. If the state is `available`, the image
        /// is successfully registered and can be used to launch an instance.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Describes a state change. Fields are `UNSET` if not available.
        /// * `state_reason.code` - The reason code for the state change.
        /// * `state_reason.message` - The message for the state change.
        /// </summary>
        public readonly ImmutableDictionary<string, string> StateReason;
        /// <summary>
        /// Any tags assigned to the image.
        /// * `tags.#.key` - The key name of the tag.
        /// * `tags.#.value` - The value of the tag.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// The type of virtualization of the AMI (ie: `hvm` or
        /// `paravirtual`).
        /// </summary>
        public readonly string VirtualizationType;

        [OutputConstructor]
        private GetAmiResult(
            string architecture,

            string arn,

            ImmutableArray<Outputs.GetAmiBlockDeviceMappingResult> blockDeviceMappings,

            string creationDate,

            string description,

            ImmutableArray<string> executableUsers,

            ImmutableArray<Outputs.GetAmiFilterResult> filters,

            string hypervisor,

            string id,

            string imageId,

            string imageLocation,

            string imageOwnerAlias,

            string imageType,

            string kernelId,

            bool? mostRecent,

            string name,

            string? nameRegex,

            string ownerId,

            ImmutableArray<string> owners,

            string platform,

            ImmutableArray<Outputs.GetAmiProductCodeResult> productCodes,

            bool @public,

            string ramdiskId,

            string rootDeviceName,

            string rootDeviceType,

            string rootSnapshotId,

            string sriovNetSupport,

            string state,

            ImmutableDictionary<string, string> stateReason,

            ImmutableDictionary<string, string> tags,

            string virtualizationType)
        {
            Architecture = architecture;
            Arn = arn;
            BlockDeviceMappings = blockDeviceMappings;
            CreationDate = creationDate;
            Description = description;
            ExecutableUsers = executableUsers;
            Filters = filters;
            Hypervisor = hypervisor;
            Id = id;
            ImageId = imageId;
            ImageLocation = imageLocation;
            ImageOwnerAlias = imageOwnerAlias;
            ImageType = imageType;
            KernelId = kernelId;
            MostRecent = mostRecent;
            Name = name;
            NameRegex = nameRegex;
            OwnerId = ownerId;
            Owners = owners;
            Platform = platform;
            ProductCodes = productCodes;
            Public = @public;
            RamdiskId = ramdiskId;
            RootDeviceName = rootDeviceName;
            RootDeviceType = rootDeviceType;
            RootSnapshotId = rootSnapshotId;
            SriovNetSupport = sriovNetSupport;
            State = state;
            StateReason = stateReason;
            Tags = tags;
            VirtualizationType = virtualizationType;
        }
    }
}
