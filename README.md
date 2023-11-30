# Semantic Versioning for Cloud Resources


## Introduction

In the context of cloud deployments, there exists an identity challenge. More and more deployments you are doing, the more likely you start to ask yourself how to distinguish them.

Even a small system, doing a few testing, staging and production deployments quickly become a nightmare. If your stack spawns many cloud resources, you are in danger of making a mistake by destroying production environments. If identity schema are loosely defined, the automation becomes harder. 

As a solution to this problem, we propose a simple set of rules and requirements that dictate how naming to cloud resources is assigned. These rules derived from widespread common practices on building automation over git-based version control system. For this approach to work, you first need to adapt staging deployment practices. Consider a sandbox deployment to TEST, than integration of feature to MAIN snapshot and finally integration into LIVE environments.


## Semantic Versioning Specification (TagVer)

The key words “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in [RFC 2119](https://tools.ietf.org/html/rfc2119).

1. Software using Semantic Versioning MUST use version control system, accept contribution via change proposals (e.g. pull requests) and use tags ("releases") as part of lifecycle management. 

2. Software using Semantic Versioning MUST deploy artifacts to cloud, using the naming sematic defined by the cloud provider (e.g. "A stack name can contain only alphanumeric characters (case-sensitive) and hyphens"). A normal identity of cloud resource MUST be denoted by appending a hyphen and TagVer identifier (e.g. example-tag).

4. A TagVer identifier MUST be either identifying a deployment for a change proposal (TEST) or identifying a deployment of latest snapshot (MAIN) or identifying a production deployment from tag/release (LIVE).    

5. The deployment identifier of change proposal (TEST) MUST take a form `prX`, where `pr` is constant literal standing for "preview" abbreviation. `X` is non-negative integers, which MUST increase numerically. For instance: pr1 -> pr2 -> ... -> pr999. Software MAY derive `X` from identity of change proposal, which is managed through version control system.

6. The deployment identifier of latest snapshot (MAIN) MUST be associated with the branch name in the version control system where change proposals are integrated. Typically, microservices uses `main` branch for this purpose. 

7. Once a change proposal is integrated into the main branch, the corresponding TEST deployment MUST be destroyed and MAIN deployment updated accordingly.   

8. The deployment identifier of production release (LIVE) MUST be associated with the tag name in the version control system. Software system MAY use [SemVer](https://semver.org) specification to identify production releases. However, Software MAY benefit from usage of simple form `vX`, where `X` is non-negative integers, which MUST increase numerically.

9. Software MAY follow a "immutable" identity pattern where new cloud resources are created for each deployment (nothing is shared between them). The pattern supports parallel deployment of cloud resource and high robustness to failures, the rollback is a traffic switch from resource v2 back to v1. 
```
            v1             v2
main ─┬───┬─┴── main ─┬───┬─┴── ...
     prX──┘          prX──┘
```

10. Software MAY follow a "continual" identity pattern where new cloud resources are only created for TEST deployment but MAIN and LIVE are mutable in between releases. The pattern supports parallel preview of configurations but main and live environments are mutable. The pattern simplify the migration process but trade-offs for roll forward strategy to deal with failures.
```
v1 ─────────┬── v1 ─────────┬── ...
            │               │              
main ─┬───┬─┴── main ─┬───┬─┴── ...
     prX──┘          prX──┘
```

11. Software MAY follow a "permanent" identity pattern where mutable  cloud resources are only created only for MAIN and LIVE. Software MAY use the pattern with deployment of very expensive and time consuming deployments.

```
v1 ─────────┬── v1 ─────────┬── ...
            │               │              
main ───────┴── main ───────┴── ...
      prX             prX
```

12. Software MAY establish dependencies between stacks (e.g. api layer depends on the database). Software MUST prepend TagVer with a name of layer following `-` and use `;` to concatenate versions (e.g. `api-pr1;db-main`, `api-v1;db-v1`)


## License

[![See LICENSE](https://img.shields.io/github/license/fogfish/tagver.svg?style=for-the-badge)](LICENSE)
