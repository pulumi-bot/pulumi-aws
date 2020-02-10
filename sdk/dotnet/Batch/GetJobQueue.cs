// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Batch
{
    public static partial class Invokes
    {
        /// <summary>
        /// The Batch Job Queue data source allows access to details of a specific
        /// job queue within AWS Batch.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-aws/blob/master/website/docs/d/batch_job_queue.html.markdown.
        /// </summary>
        public static Task<GetJobQueueResult> GetJobQueue(GetJobQueueArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetJobQueueResult>("aws:batch/getJobQueue:getJobQueue", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetJobQueueArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the job queue.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetJobQueueArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetJobQueueResult
    {
        /// <summary>
        /// The ARN of the job queue.
        /// </summary>
        public readonly string Arn;
        /// <summary>
        /// The compute environments that are attached to the job queue and the order in
        /// which job placement is preferred. Compute environments are selected for job placement in ascending order.
        /// * `compute_environment_order.#.order` - The order of the compute environment.
        /// * `compute_environment_order.#.compute_environment` - The ARN of the compute environment.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetJobQueueComputeEnvironmentOrdersResult> ComputeEnvironmentOrders;
        public readonly string Name;
        /// <summary>
        /// The priority of the job queue. Job queues with a higher priority are evaluated first when
        /// associated with the same compute environment.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Describes the ability of the queue to accept new jobs (for example, `ENABLED` or `DISABLED`).
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The current status of the job queue (for example, `CREATING` or `VALID`).
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A short, human-readable string to provide additional details about the current status
        /// of the job queue.
        /// </summary>
        public readonly string StatusReason;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetJobQueueResult(
            string arn,
            ImmutableArray<Outputs.GetJobQueueComputeEnvironmentOrdersResult> computeEnvironmentOrders,
            string name,
            int priority,
            string state,
            string status,
            string statusReason,
            string id)
        {
            Arn = arn;
            ComputeEnvironmentOrders = computeEnvironmentOrders;
            Name = name;
            Priority = priority;
            State = state;
            Status = status;
            StatusReason = statusReason;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetJobQueueComputeEnvironmentOrdersResult
    {
        public readonly string ComputeEnvironment;
        public readonly int Order;

        [OutputConstructor]
        private GetJobQueueComputeEnvironmentOrdersResult(
            string computeEnvironment,
            int order)
        {
            ComputeEnvironment = computeEnvironment;
            Order = order;
        }
    }
    }
}
