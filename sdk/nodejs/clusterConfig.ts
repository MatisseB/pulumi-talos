// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class ClusterConfig extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'talos:index:ClusterConfig';

    /**
     * Returns true if the given object is an instance of ClusterConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterConfig.__pulumiType;
    }

    /**
     * Provides control plane specific configuration options.
     */
    public /*out*/ readonly configYAML!: pulumi.Output<string>;

    /**
     * Create a ClusterConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterConfigArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.controlPlaneConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controlPlaneConfig'");
            }
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["controlPlaneConfig"] = args ? args.controlPlaneConfig : undefined;
            resourceInputs["configYAML"] = undefined /*out*/;
        } else {
            resourceInputs["configYAML"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterConfig.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
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
