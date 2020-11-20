## Installation

``` shell
cd ssl && make cert
```

``` shell
cd.. && helm install helm
````

## Usage

Add annotation to pod and injector will inject the tolerations. For eg
``` yaml
tolerations-injector/inject-toleration-1: "{\"key\": \"seqster.com/group\",\"effect\": \"NoExecute\",\"operator\": \"Equal\",\"value\": \"ops\"}"
```
