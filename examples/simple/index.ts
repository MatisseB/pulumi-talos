import * as talos from "@pulumi/talos";

const page = new talos.ClusterConfig("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
});

export const bucket = page.bucket;
export const url = page.websiteUrl;
