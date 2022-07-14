import * as pulumi from "@pulumi/pulumi";
export interface ControlPlaneConfigArgs {
    endpoint?: pulumi.Input<string>;
    localAPIServerPort?: pulumi.Input<number>;
}
