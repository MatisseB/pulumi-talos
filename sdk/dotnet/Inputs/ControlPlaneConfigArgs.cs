// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Talos.Inputs
{

    public sealed class ControlPlaneConfigArgs : Pulumi.ResourceArgs
    {
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("localAPIServerPort")]
        public Input<int>? LocalAPIServerPort { get; set; }

        public ControlPlaneConfigArgs()
        {
        }
    }
}
