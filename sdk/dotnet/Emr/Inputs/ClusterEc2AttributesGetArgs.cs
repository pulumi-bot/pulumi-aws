// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Emr.Inputs
{

    public sealed class ClusterEc2AttributesGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// String containing a comma separated list of additional Amazon EC2 security group IDs for the master node
        /// </summary>
        [Input("additionalMasterSecurityGroups")]
        public Input<string>? AdditionalMasterSecurityGroups { get; set; }

        /// <summary>
        /// String containing a comma separated list of additional Amazon EC2 security group IDs for the slave nodes as a comma separated string
        /// </summary>
        [Input("additionalSlaveSecurityGroups")]
        public Input<string>? AdditionalSlaveSecurityGroups { get; set; }

        /// <summary>
        /// Identifier of the Amazon EC2 EMR-Managed security group for the master node
        /// </summary>
        [Input("emrManagedMasterSecurityGroup")]
        public Input<string>? EmrManagedMasterSecurityGroup { get; set; }

        /// <summary>
        /// Identifier of the Amazon EC2 EMR-Managed security group for the slave nodes
        /// </summary>
        [Input("emrManagedSlaveSecurityGroup")]
        public Input<string>? EmrManagedSlaveSecurityGroup { get; set; }

        /// <summary>
        /// Instance Profile for EC2 instances of the cluster assume this role
        /// </summary>
        [Input("instanceProfile", required: true)]
        public Input<string> InstanceProfile { get; set; } = null!;

        /// <summary>
        /// Amazon EC2 key pair that can be used to ssh to the master node as the user called `hadoop`
        /// </summary>
        [Input("keyName")]
        public Input<string>? KeyName { get; set; }

        /// <summary>
        /// Identifier of the Amazon EC2 service-access security group - required when the cluster runs on a private subnet
        /// </summary>
        [Input("serviceAccessSecurityGroup")]
        public Input<string>? ServiceAccessSecurityGroup { get; set; }

        /// <summary>
        /// VPC subnet id where you want the job flow to launch. Cannot specify the `cc1.4xlarge` instance type for nodes of a job flow launched in a Amazon VPC
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public ClusterEc2AttributesGetArgs()
        {
        }
    }
}
