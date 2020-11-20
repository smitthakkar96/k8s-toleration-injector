package mutate

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	v1beta1 "k8s.io/api/admission/v1beta1"
)

func TestMutatePod(t *testing.T) {
	rawJSON := `{
		"kind": "AdmissionReview",
		"apiVersion": "admission.k8s.io/v1beta1",
		"request": {
			"uid": "7f0b2891-916f-4ed6-b7cd-27bff1815a8c",
			"kind": {
				"group": "",
				"version": "v1",
				"kind": "Pod"
			},
			"resource": {
				"group": "",
				"version": "v1",
				"resource": "pods"
			},
			"requestKind": {
				"group": "",
				"version": "v1",
				"kind": "Pod"
			},
			"requestResource": {
				"group": "",
				"version": "v1",
				"resource": "pods"
			},
			"namespace": "yolo",
			"operation": "CREATE",
			"userInfo": {
				"username": "kubernetes-admin",
				"groups": [
					"system:masters",
					"system:authenticated"
				]
			},
			"object": {
				"kind": "Pod",
				"apiVersion": "v1",
				"metadata": {
					"name": "c7m",
					"namespace": "yolo",
					"creationTimestamp": null,
					"labels": {
						"name": "c7m"
					},
					"annotations": {
						"tolerations-injector/inject-toleration-0": "{\"key\": \"seqster.com/group\",\"effect\": \"NoExecute\",\"operator\": \"Equal\",\"value\": \"datastore\"}",
						"tolerations-injector/inject-toleration-1": "{\"key\": \"seqster.com/group\",\"effect\": \"NoExecute\",\"operator\": \"Equal\",\"value\": \"ops\"}"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "vault-agent",
							"image": "centos:7",
							"command": [
								"/bin/bash"
							],
							"args": [
								"-c",
								"trap \"killall sleep\" TERM; trap \"kill -9 sleep\" KILL; sleep infinity"
							],
							"resources": {},
							"volumeMounts": [
								{
									"name": "default-token-5z7xl",
									"readOnly": true,
									"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
								}
							],
							"terminationMessagePath": "/dev/termination-log",
							"terminationMessagePolicy": "File",
							"imagePullPolicy": "IfNotPresent",
							"env": [
								{
									"name": "VAULT_CONFIG",
									"value": "eyJhdXRvX2F1dGgiOnsibWV0aG9kIjp7InR5cGUiOiJrdWJlcm5ldGVzIiwiY29uZmlnIjp7InJvbGUiOiJla3MtZGV2LXNwYWNlIn19LCJzaW5rIjpbeyJ0eXBlIjoiZmlsZSIsImNvbmZpZyI6eyJwYXRoIjoiL2hvbWUvdmF1bHQvLnRva2VuIn19XX0sImV4aXRfYWZ0ZXJfYXV0aCI6dHJ1ZSwicGlkX2ZpbGUiOiIvaG9tZS92YXVsdC8ucGlkIiwidmF1bHQiOnsiYWRkcmVzcyI6Imh0dHA6Ly9pbnRlcm5hbC1kZXYtdmF1bHQtODg3NTQ4OTAzLmV1LXdlc3QtMS5lbGIuYW1hem9uYXdzLmNvbTo4MjAwIn0sInRlbXBsYXRlIjpbeyJkZXN0aW5hdGlvbiI6Ii92YXVsdC9zZWNyZXRzL2RiLWNyZWRzIiwiY29udGVudHMiOiJ7eyB3aXRoIHNlY3JldCBcInNlY3JldC9kZXYvYXV0aGVudGljYXRvci1mcm9udGVuZC9hdXRoZW50aWNhdG9yLWZyb250ZW5kLWRldi1zcGFjZS9BTEVYQV9WRVJJRllfSURcIiB9fXt7IHJhbmdlICRrLCAkdiA6PSAuRGF0YSB9fXt7ICRrIH19OiB7eyAkdiB9fVxue3sgZW5kIH19e3sgZW5kIH19IiwibGVmdF9kZWxpbWl0ZXIiOiJ7eyIsInJpZ2h0X2RlbGltaXRlciI6In19In1dfQ=="
								}
							]
						}
					],
					"initContainers": [
						{
							"name": "vault-agent-init",
							"image": "centos:7",
							"command": [
								"/bin/bash"
							],
							"args": [
								"-c",
								"trap \"killall sleep\" TERM; trap \"kill -9 sleep\" KILL; sleep infinity"
							],
							"resources": {},
							"volumeMounts": [
								{
									"name": "default-token-5z7xl",
									"readOnly": true,
									"mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
								}
							],
							"terminationMessagePath": "/dev/termination-log",
							"terminationMessagePolicy": "File",
							"imagePullPolicy": "IfNotPresent",
							"env": [
								{
									"name": "VAULT_CONFIG",
									"value": "eyJhdXRvX2F1dGgiOnsibWV0aG9kIjp7InR5cGUiOiJrdWJlcm5ldGVzIiwiY29uZmlnIjp7InJvbGUiOiJla3MtZGV2LXNwYWNlIn19LCJzaW5rIjpbeyJ0eXBlIjoiZmlsZSIsImNvbmZpZyI6eyJwYXRoIjoiL2hvbWUvdmF1bHQvLnRva2VuIn19XX0sImV4aXRfYWZ0ZXJfYXV0aCI6dHJ1ZSwicGlkX2ZpbGUiOiIvaG9tZS92YXVsdC8ucGlkIiwidmF1bHQiOnsiYWRkcmVzcyI6Imh0dHA6Ly9pbnRlcm5hbC1kZXYtdmF1bHQtODg3NTQ4OTAzLmV1LXdlc3QtMS5lbGIuYW1hem9uYXdzLmNvbTo4MjAwIn0sInRlbXBsYXRlIjpbeyJkZXN0aW5hdGlvbiI6Ii92YXVsdC9zZWNyZXRzL2RiLWNyZWRzIiwiY29udGVudHMiOiJ7eyB3aXRoIHNlY3JldCBcInNlY3JldC9kZXYvYXV0aGVudGljYXRvci1mcm9udGVuZC9hdXRoZW50aWNhdG9yLWZyb250ZW5kLWRldi1zcGFjZS9BTEVYQV9WRVJJRllfSURcIiB9fXt7IHJhbmdlICRrLCAkdiA6PSAuRGF0YSB9fXt7ICRrIH19OiB7eyAkdiB9fVxue3sgZW5kIH19e3sgZW5kIH19IiwibGVmdF9kZWxpbWl0ZXIiOiJ7eyIsInJpZ2h0X2RlbGltaXRlciI6In19In1dfQ=="
								}
							]
						}
					],
					"restartPolicy": "Always",
					"terminationGracePeriodSeconds": 30,
					"dnsPolicy": "ClusterFirst",
					"serviceAccountName": "default",
					"serviceAccount": "default",
					"securityContext": {},
					"schedulerName": "default-scheduler",
					"tolerations": [
						{
							"key": "node.kubernetes.io/not-ready",
							"operator": "Exists",
							"effect": "NoExecute",
							"tolerationSeconds": 300
						},
						{
							"key": "node.kubernetes.io/unreachable",
							"operator": "Exists",
							"effect": "NoExecute",
							"tolerationSeconds": 300
						}
					],
					"priority": 0,
					"enableServiceLinks": true
				},
				"status": {}
			},
			"oldObject": null,
			"dryRun": false,
			"options": {
				"kind": "CreateOptions",
				"apiVersion": "meta.k8s.io/v1"
			}
		}
	}`
	admReview := v1beta1.AdmissionReview{}
	err := json.Unmarshal([]byte(rawJSON), &admReview)
	response, err := MutatePod(admReview)
	if err != nil {
		t.Errorf("failed to mutate AdmissionRequest %s with error %s", string(response), err)
	}
	admReview = v1beta1.AdmissionReview{}
	err = json.Unmarshal(response, &admReview)
	assert.NoError(t, err, "failed to unmarshal with error %s", err)

	admReviewResponse := admReview.Response
	parsedPatch := new([]map[string]interface{})
	json.Unmarshal(admReviewResponse.Patch, &parsedPatch)
	assert.Equal(t, 2, len(*parsedPatch))
}
