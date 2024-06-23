

### Secrets management

Secrets management is sketchy AF.

First git clone this very private repository. The repository should be placed beside the main `salz` repository.

```
git clone triwizard.tail7ecee.ts.net:/home/mosiman/salz-secret
```

This repository contains scripts for decrypting secrets. The secrets are encrypted with the public keys of the salz maintainers as listed on Github.

To decrypt, `../salz-secret/decrypt.sh ../salz-secret/dev-secrets.json.age | jq '.'`
