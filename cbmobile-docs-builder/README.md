
history: https://gist.github.com/tleyden/c90c2708e1bf4894195f

## Configure git with credentials

## Git clone

And submodules ..

## Generate docs

## Configure aws cli

## Run aws cli

```
cd /root/couchbase-mobile-portal/site/gen
aws s3 sync --acl public-read --exclude *.svg . s3://tleyden-couchbase/mobile-docs/master/
aws s3 sync --acl public-read --include *.svg --content-type "image/svg+xml" . s3://tleyden-couchbase/mobile-docs/master/
```

## References

* https://github.com/s3tools/s3cmd/issues/198
* 