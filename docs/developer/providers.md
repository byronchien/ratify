# Store/Verifier Providers
The framework uses a provider model for extensibility to support different types of referrer stores and verifiers. It supports two types of providers. 

Built-In/Internal providers are available in the source of the framework and are registered during startup using the ```init``` function. Referrer store using [ORAS](https://github.com/oras-project/oras) and signature verification using [notaryv2](https://github.com/notaryproject/notation) are currently available as the built-in providers in the framework and are managed by the framework.

External/Plugin providers are external to the source and process of the framework. They are registered as binaries that implement the plugin specification for the corresponding provider. The framework will locate these binaries in the configured paths and invoke them by passing the required parameters as per the specification. [SBOM](..\plugins\verifier\sbom) verification and signature verification using [cosign](..\plugins\verifier\cosign) libraries are examples of the providers that are implemented as plugins.

# Policy Providers

Ratify implements an extensible policy provider interface allowing for different policy providers to be created and registered. The policy provider to be used is determined by the policy plugin specified in the `policy` section of the configuration. 

Currently, Ratify supports a Configuration based Policy Provider named `configPolicy`.

## How is the Policy Provider used in Ratify execution?

The executor is the "glue" that links all Ratify plugin-based components such as the verifiers, referrer stores, and policy providers. The policy provider is primarily invoked for each reference artifact discovered AFTER a configured verifier is found that can be used to verify that particular artifacy type. If a reference artifact is found that does not have a corresponding verifier configured that can perform verification, the executor will ignore this artifact. As a result, this unverifiable artifact will NOT influence the final overall verification success determined by the policy provider.

## Config Policy Provider

```
...
"policy": {
    "version": "1.0.0",
    "plugin": {
        "name": "configPolicy",
        "artifactVerificationPolicies": {
            "application/vnd.cncf.notary.signature": "any"
        }
    }
},
...
```

- The `name` field is REQUIRED and MUST match the name of the registered policy provider
- `artifactVerificationPolicies`: map of artifact type to policy; each entry in the map's policy must be satisfied for Ratify to return true.
    - `any`: policy that REQUIRES at least one artifact of specified type to verify to `true` 
    - `all`: policy that REQUIRES all artifacts of specified type to verify to `true``
- Default policy:
    - The `default` policy applies to unspecified artifact types. The `default` policy is set to `all`. Thus, all unspecified artifact types must have all successful verification results for an overall success result.
    - The `default` policy can be overridden to `any` in the map:
        
        ```
        ...
        "policy": {
            "version": "1.0.0",
            "plugin": {
                "name": "configPolicy",
                "artifactVerificationPolicies": {
                    "default": "any"
                }
            }
        },
        ...
        ```
### Examples:

- Require all reference artifacts associated with subject image to be verify successfully:
    ```
    ...
    "policy": {
        "version": "1.0.0",
        "plugin": {
            "name": "configPolicy"
        }
    },
    ...
    ```
- Require at least one reference artifact of the same type to verify succesfully. (relaxes the default policy to 'any'):
    ```
    ...
    "policy": {
        "version": "1.0.0",
        "plugin": {
            "name": "configPolicy",
            "artifactVerificationPolicies": {
                "default": "any"
            }
        }
    },
    ...
    ```
- For a specific artifact type, relax requirement so only one success is needed for artifacts of that type:
    ```
    ...
    "policy": {
        "version": "1.0.0",
        "plugin": {
            "name": "configPolicy",
            "artifactVerificationPolicies": {
                "application/vnd.cncf.notary.signature": "any"
            }
        }
    },
    ...
    ```

## Notational Conventions

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" are to be interpreted as described in [RFC 2119](http://tools.ietf.org/html/rfc2119).

The key words "unspecified", "undefined", and "implementation-defined" are to be interpreted as described in the [rationale for the C99 standard](http://www.open-std.org/jtc1/sc22/wg14/www/C99RationaleV5.10.pdf#page=18).