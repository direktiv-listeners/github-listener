# Github Listener

The Github listener accepts webhook POST requests from Github and sends them as cloudevents to Direktiv. The webhook content is in the `data` section of the cloud event. The content is based on the event posted from Github. The event type is the lower case Github event. e.g. `push`. It can be installed in plain mode without Knative or as a Knative source. The following YAML files are installing the listener but the value for `ingressClassName` might need to change.

## Plain Mode

[plain.yaml](https://github.com/direktiv-listeners/github-listener/blob/main/kubernetes/plain.yaml)

## Knative Mode

[knative.yaml](https://github.com/direktiv-listeners/github-listener/blob/main/kubernetes/knative.yaml)

## Configuration

| Environment Variable      | Description |
| ----------- | ----------- |
| DIREKTIV_GITHUB_DEBUG      | Enable debug mode      |
| DIREKTIV_GITHUB_SECRET      | Secret value for Github webhook     |
| DIREKTIV_GITHUB_TOKEN      | Direktiv token if /broadcast API is used    |
| DIREKTIV_GITHUB_ENDPOINT      | Direktiv endpoint    |
| DIREKTIV_GITHUB_INSECURE_TLS      | Skip verifying certificates    |
| DIREKTIV_GITHUB_PATH      | Path to serve the webhook destination   |
