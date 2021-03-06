## Sign Policy

### Define signer for each namespaces

SignPolicy is a custom resource to define who can be a valid signer for resources in a namespace or for cluster scope resources.
Only a SignPolicy resource is defined in IE namespace (`integrity-enforcer-ns` in this documentation)and initial SignPolicy resource is created during IE installation. You can access it by
```
$ oc get signpolicies.research.ibm.com signer-policy -n integrity-enforcer-ns -o yaml > /tmp/sign-policy.yaml
```

You can configure the policy by adding the following snipet to `/tmp/sign-policy.yaml`

Example below is to define
- signer `signer-a` is identified when email of subject of signature is `signer@enterprise.com`
- signer `signer-a` is approved signer for the resources to be created in namespace `secure-ns`.

```yaml
spec:
  policy:
    policies:
    - namespaces:
      - secure-ns
      signers:
      - signer-a
    signers:
    - name: signer-a
      subjects:
      - email: signer@enterprise.com
```

For matching signer, you can use the following attributes: `email`, `uid`, `country`, `organization`, `organizationalUnit`, `locality`, `province`, `streetAddress`, `postalCode`, `commonName` and `serialNumber`.

Then, this policy is applied back to a cluster by:

```
$ oc apply -f /tmp/sign-policy.yaml -n integrity-enforcer-ns signpolicy.research.ibm.com/signer-policy configured
```

You can define namespace matcher by using `excludeNamespaces`. For example below, signer `signer-a` can sign resource in `secure-ns` namespace, and another signer `signer-b` can sign resource in all other namespaces except `secure-ns`.

```yaml
policies:
- namespaces:
  - secure-ns
  signers:
  - signer-a
- namespaces:
  - '*'
  excludeNamespaces
  - secure-ns
  signers:
  - signer-b
- scope: Cluster
  signers:
  - signer-a
  - signer-b
```

### Define Signer for cluster-scope resources
You can define a signer for cluster-scope resources similarily. Signer `signer-a` and `signer-b` can sign cluster-scope resources in the example below.

```yaml
policies:
- scope: Cluster
  signers:
  - signer-a
  - signer-b
```

### Break Glass
When you need to disable blocking by signature verification in a certain namespace, you can enable break glass mode, which means the request to the namespace without valid signature is allowed during the break glass on. For example, break glass on `secure-ns` namespace can be set on by

```
spec:
  policy:
    breakGlass:
      - namespaces:
        - secure-ns
```
Break glass on cluster-scope resources can be set on by
```
spec:
  policy:
    breakGlass:
      - scope: Cluster
```

During break glass mode on, the request without signature is allowed but it is marked by `integrityUnverified` label.


### Example of Sign Policy

```yaml
spec:
  policy:
    policies:
    - namespaces:
      - secure-ns
      signers:
      - signer-a
    - namespaces:
      - '*'
      excludeNamespaces
      - secure-ns
      signers:
      - signer-b
    - scope: Cluster
      signers:
      - signer-a
      - signer-b
    signers:
    - name: signer-a
      subjects:
      - email: secure-ns-signer@enterprise.com
    - name: signer-b
      subjects:
      - email: default-signer@enterprise.com
```

