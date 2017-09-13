package v1alpha1

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime"
)

type Alert interface {
	GetName() string
	GetNamespace() string
	Command() string
	GetCheckInterval() time.Duration
	GetAlertInterval() time.Duration
	IsValid() (bool, error)
	GetNotifierSecretName() string
	GetReceivers() []Receiver
	ObjectReference() runtime.Object
}
