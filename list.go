package certificatetpr

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// List represents a list of CustomObject resources.
type List struct {
	metav1.ObjectMeta `json:"metadata"`
	metav1.TypeMeta   `json:",inline"`

	Items []CustomObject `json:"items" yaml:"items"`
}
