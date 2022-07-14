import * as pulumi from "@pulumi/pulumi";
import { input as inputs } from "./types";
export declare class ClusterConfig extends pulumi.ComponentResource {
    /**
     * Returns true if the given object is an instance of ClusterConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    static isInstance(obj: any): obj is ClusterConfig;
    /**
     * Provides control plane specific configuration options.
     */
    readonly configYAML: pulumi.Output<string>;
    /**
     * Create a ClusterConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterConfigArgs, opts?: pulumi.ComponentResourceOptions);
}
/**
 * The set of arguments for constructing a ClusterConfig resource.
 */
export interface ClusterConfigArgs {
    /**
     * Configures the cluster's name.
     */
    clusterName: pulumi.Input<string>;
    /**
     * Provides control plane specific configuration options.
     */
    controlPlaneConfig: pulumi.Input<inputs.ControlPlaneConfigArgs>;
}
