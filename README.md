# r53-operator

**NOTE**: This is an alpha-status project. We do regular tests on the code and functionality, but we can not assure a production-ready stability.

route53 operator create/manage/delete route53 records atop Kubernetes.

## Requirements

route53 operator is meant to be run on Kubernetes 1.8+.
All dependecies have been vendored, so there's no need to any additional download.

## Testing

- Run `kubectl proxy` to get local access to kubernetes API server
- Run `go run cmd/*` to start controller
- Create sample resource with `kubectl apply -f manifests/`
- Delete sample resource with `kubectl delete -f manifests/`
