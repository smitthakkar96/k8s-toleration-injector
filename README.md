# k8s-mutate-webhook

A playground to try build a crude k8s mutating webhook; the goal is to mutate a Pod CREATE request to _always_ use a debian image and by doing this, learning more about
the k8s api, objects, etc. - eventually figure out how scalable this is (could be made) if one had 1000 pods to schedule (concurrently) 

## build 

```
make
```

## test

```
make test
```

## ssl/tls

the `ssl/` dir contains a script to create a self-signed certificate, not sure this will even work when running in k8s but that's part of figuring this out I guess

_NOTE: the app expects the cert/key to be in `ssl/` dir relative to where the app is running/started and currently is hardcoded to `mutateme.{key,pem}`_

```
cd ssl/ 
make 
```

## docker

to create a docker image .. 

```
make docker
```

it'll be tagged with the current git commit (short `ref`) and `:latest`

don't forget to update `IMAGE_PREFIX` in the Makefile or set it when running `make`

### images

[`alexleonhardt/k8s-mutate-webhook`](https://cloud.docker.com/repository/docker/alexleonhardt/k8s-mutate-webhook)



## watcher

useful during devving ... 

```
watcher -watch github.com/alex-leonhardt/k8s-mutate-webhook -run github.com/alex-leonhardt/k8s-mutate-webhook/cmd/
```

## kudos

- [https://github.com/morvencao/kube-mutating-webhook-tutorial](https://github.com/morvencao/kube-mutating-webhook-tutorial)
