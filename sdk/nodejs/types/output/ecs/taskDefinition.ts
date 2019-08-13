// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputApi from "../../../types/input";
import * as outputApi from "../../../types/output";
import * as utilities from "../../../utilities";

export interface TaskDefinitionPlacementConstraint {
    /**
     * Cluster Query Language expression to apply to the constraint.
     * For more information, see [Cluster Query Language in the Amazon EC2 Container
     * Service Developer
     * Guide](http://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-query-language.html).
     */
    expression?: string;
    /**
     * The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
     */
    type: string;
}

export interface TaskDefinitionProxyConfiguration {
    /**
     * The name of the container that will serve as the App Mesh proxy.
     */
    containerName: string;
    /**
     * The set of network configuration parameters to provide the Container Network Interface (CNI) plugin, specified a key-value mapping.
     */
    properties?: {[key: string]: string};
    /**
     * The proxy type. The default value is `APPMESH`. The only supported value is `APPMESH`.
     */
    type?: string;
}

export interface TaskDefinitionVolume {
    /**
     * Used to configure a docker volume
     */
    dockerVolumeConfiguration?: outputApi.ecs.TaskDefinitionVolumeDockerVolumeConfiguration;
    /**
     * The path on the host container instance that is presented to the container. If not set, ECS will create a nonpersistent data volume that starts empty and is deleted after the task has finished.
     */
    hostPath?: string;
    /**
     * The name of the volume. This name is referenced in the `sourceVolume`
     * parameter of container definition in the `mountPoints` section.
     */
    name: string;
}

export interface TaskDefinitionVolumeDockerVolumeConfiguration {
    /**
     * If this value is `true`, the Docker volume is created if it does not already exist. *Note*: This field is only used if the scope is `shared`.
     */
    autoprovision?: boolean;
    /**
     * The Docker volume driver to use. The driver value must match the driver name provided by Docker because it is used for task placement.
     */
    driver?: string;
    /**
     * A map of Docker driver specific options.
     */
    driverOpts?: {[key: string]: string};
    /**
     * A map of custom metadata to add to your Docker volume.
     */
    labels?: {[key: string]: string};
    /**
     * The scope for the Docker volume, which determines its lifecycle, either `task` or `shared`.  Docker volumes that are scoped to a `task` are automatically provisioned when the task starts and destroyed when the task stops. Docker volumes that are `scoped` as shared persist after the task stops.
     */
    scope: string;
}
