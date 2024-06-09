package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Define the MailgunConfig struct
type MailgunConfig struct {
	Domain          string `json:"domain,omitempty"`
	APIKeySecretRef string `json:"apiKeySecretRef,omitempty"`
}

// Define the EmailSenderConfigSpec struct with the new field for Mailgun configuration
type EmailSenderConfigSpec struct {
	ApiTokenSecretRef string        `json:"apiTokenSecretRef"`
	SenderEmail       string        `json:"senderEmail"`
	MailgunConfig     MailgunConfig `json:"mailgunConfig,omitempty"`
}

// Define the EmailSenderConfigStatus struct
type EmailSenderConfigStatus struct {
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
}

// Define the EmailSenderConfig custom resource
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type EmailSenderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EmailSenderConfigSpec   `json:"spec,omitempty"`
	Status EmailSenderConfigStatus `json:"status,omitempty"`
}

// Define the EmailSenderConfigList struct
// +kubebuilder:object:root=true
type EmailSenderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmailSenderConfig `json:"items"`
}

// Register the EmailSenderConfig and EmailSenderConfigList with the scheme
func init() {
	SchemeBuilder.Register(&EmailSenderConfig{}, &EmailSenderConfigList{})
}
