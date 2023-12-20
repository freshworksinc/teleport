---
authors: Hugo Hervieux (hugo.hervieux@goteleport.com)
state: draft
---

# RFD 0160 - Kubernetes Operator Resource Versioning

## Required Approvers

* Engineering @r0mant && (@tigrato || @marcoandredinis)
* Security: (@reedloden || @jentfoo)
* Product: (@xinding33 || @klizhentas )

## What

This RFD discusses how we can implement multiple version support for the same
resource in the Teleport Kubernetes Operator. For example, the operator currently
only supports `RoleV5` while we have released `RoleV6` and `RoleV7` resources
with new capabilities.

## Why

Users want to manage their Teleport resources via the Teleport Kubernetes Operator.
However, in its current state, the operator does not support when we introduce a
newer version of a resource. This blocks users from leveraging new Teleport
features such as granular Kubernetes Access Control with RoleV7.

## Details

### Current state

We currently version the Teleport CRs like the Teleport resources. While this
makes sense from a user point of view, this is not compatible with the way
Teleport manages versions. Teleport does not support exposing resources through
separate per-version APIs. The version conversion happens on the Teleport side,
when establishing new defaults.

This means we end up with two sources of truth:
- the resource stored in Kubernetes.
  The Kubernetes storage is version-aware but does not convert between versions
- the resource stored in Teleport.
  The Teleport storage is not entirely version-aware as it converts all resources
  to the latest version on startup and during `CheckAndSetDefaults()`.

Both storages are treating versions differently and don't agree as to how to
represent a resource in a given version.

### Suggested approach

We break the relation between the CRD version and the Teleport resource, and
specify the version in the CR spec. This means users would use a single version
of `resources.teleport.dev` for all their resources.

Before, an admin would create a RoleV5 by creating a `TeleportRole` through via the API
`resources.teleport.dev/v5` and a UserV2 through the api `resources.teleport.dev/v2`.

```yaml
apiVersion: resources.teleport.dev/v5
kind: TeleportRole
metadata:
  name: myrole
spec:
  allow:
    rules:
      - resources: ['user', 'role']
        verbs: ['list','create','read','update','delete']
---
apiVersion: resources.teleport.dev/v2
kind: TeleportUser
metadata:
  name: myuser
spec:
  roles: ['myrole']
```

After this RFD, both `TeleportRole` and `TeleportUser` resources would be created
through the `resources.teleport.dev/vX` API. The Teleport resource version would
be specified in a separate field: `teleportResourceVersion`.

```yaml
apiVersion: resources.teleport.dev/vX
teleportResourceVersion: "v5"
kind: TeleportRole
metadata:
  name: myrole
spec:
  allow:
    rules:
      - resources: ['user', 'role']
        verbs: ['list','create','read','update','delete']
---
apiVersion: resources.teleport.dev/vX
teleportResourceVersion: "v2"
kind: TeleportUser
metadata:
  name: myuser
spec:
  roles: ['myrole']
```

The version `vX` in `resources.teleport.dev/vX` needs to be higher than the highest
current version (TeleportRole is served under `v6`). We can set it to the current
Teleport version (`v15` or `v16` depending on the timing).

To avoid CRD inflation and non-semantic updates, this version should not be updated
with every Teleport major. Updating this version should only happen if we introduce
a breaking change in the API. See the [API Evolution](#api-evolution) section for
more details.

### Pros

- This is backward compatible as we can default to the existing versions when
  the `teleportResourceVersion` field is not set. We can also continue to support the
  existing CRD APIs.
- The implementation cost of this solution is several orders of magnitude lower
  than the alternatives. See [the Alternatives section](#alternatives).
- This is the same versioning approach we are already using for the Terraform Provider.
- This design doesn't add new failure modes, unlike solutions implying conversion webhooks.
- Upgrading the resource version is easy for a Teleport user: they change both the
  `teleportResourceVersion` and the affected `spec` fields simultaneously.

### Cons

- This design doesn't follow usual Kubernetes resource versioning.
- Settling on `vX` can be confusing. See [the API evolution section](#api-evolution).
- This design relies on the fact we don't radically change the resource structure.
  See [the API evolution section](#api-evolution).

### Security

This design adds little changes to the current Teleport Kubernetes Operator
security model. The only risk is users forgetting to set the `teleportResourceVersion`
field, in this case we would default to the current version. In a future operator
version, we can enforce the presence of this field if deemed necessary.

### UX

- This change is backward compatible and offers a straightforward migration
  path when we introduce a new resource version.
- The atypical versioning might be confusing for users as it doesn't follow
  usual Kubernetes patterns.

### API evolution

This design unblocks us today, but does not cover cases in which we introduce a
new resource version with a completely different structure with fields with
different types. It is unclear if this can or will happen as this would very
likely cause other issues in Teleport due to the way we handle resource versioning.

If this were to happen, we would need to revise the solution from this RFD and
introduce a new resource, or implement conversion hooks. We would also need to
change the `resources.teleport.dev/vX` version to the current Teleport version.

For example:
- We implement this RFD in Teleport 15: the API is `resources.teleport.dev/v15`
- We don't do any breaking changes before v20: the API is still `resources.teleport.dev/v15` for Teleport version 16, 17, 18 and 19.
- In Teleport v20, we introduce a breaking change in Role. We introduce a new API `resoucres.teleport.dev/v20`.

### Alternatives

When writing this RFD, two alternatives were considered:

#### Conversion hooks

The Kubernetes-friendly approach would be to make the operator aware of how the
resource is stored in Kubernetes, and do the conversions for every Kubernetes CR
API call via webhooks.

This alternative has been rejected because of the following reasons:
- We currently handle resource conversion only from the old version to the newer versions,
  this would require implementing bidirectional conversion between each resource version
  (or use a Hub & Spoke approach by converting every version to a common version). This also 
  raises questions about how to represent non-supported fields during downgrades (e.g. a
  user gets a Role v7 through a Role v5 API). This would increase the cost of adding new
  resources and increase the chances of addings bugs in the conversions due to the increasing
  conversion matrix.
- We are currently working on hiding `CheckAndSetDefaults()` from the client
  and consolidating defaults injection and resource conversion server-side
  This approach is not compatible with the conversion hooks pattern as we'd
  need to run `CheckAndSetDefaults` in the operator and duplicate the logic.
- This requires relying on conversion webhooks:
  - This adds new failure modes: the operator is not healthy, kubernetes cannot talk to the operator
  - Not all Kubernetes distributions support kube control plane to workload communication by default
    (this is blocked on private GKE clusters, and most security teams also block this on their EKS clusters).
    This would slow adoption and worsen UX.
  - This requires the webhook server to be trusted by Kubernetes, which implies creating a CA, signing certs,
    and putting them in the CRD resource. This complexifies the deployment process a lot.
- The cost of implementing this solution is way higher than the other alternatives.

#### Introducing a new API

This is a variant of the suggested solution, but instead of using `resources.teleport.dev/vX`
with `vX` being the Teleport version when this was implemented, we introduce a new `v1` API.

For example: `operator.teleport.dev/v1`.

This is cleaner and semantically easier to understand than using `vX`. However, this does
not provide easy upgrade paths from existing resources under `resources.teleport.dev` to
the new API.

One workaround would be to write lightweight controllers reconciling resources
`resources.teleport.dev` with `operator.teleport.dev` but this would add complexity
that might not be compensated by the benefits which are mostly cosmetic.

### Test Plan

The test plan is the following:
- validate that existing CRs are reconciled the same way by the operator
- validate that the operator uses the `teleportResourceVersion` field when specified
- validate that users can create RoleV6 and RoleV7 once this is implemented
