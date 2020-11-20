package mutate

import (
	v1beta1 "k8s.io/api/admission/v1beta1"
)

type TMutator func(v1beta1.AdmissionReview) ([]byte, error)

func GetAvailableMutators() map[string]TMutator {
	return map[string]TMutator{
		"Pod":            MutatePod,
	}
}
