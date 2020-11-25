# GCX SRE STUB IMAGE

## Purpose
This is a simple stub image which can be used to verify a working infrastructure. In contains following feature:

- `/` root endpoint for verifying response - request use cases
- `/metrics` metrics endpoint for exposing Prometheus metrics

## Configuration
The listening Port can be set with the environment variable `SRE_STUB_PORT` or defaults to 8080.
