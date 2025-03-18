/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	appsv1 "kubeops.dev/newpet/api/v1"
)

// nolint:unused
// log is for logging in this package.
var petsetlog = logf.Log.WithName("petset-resource")

// SetupPetSetWebhookWithManager registers the webhook for PetSet in the manager.
func SetupPetSetWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&appsv1.PetSet{}).
		WithValidator(&PetSetCustomValidator{}).
		WithDefaulter(&PetSetCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-apps-k8s-appscode-com-v1-petset,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.k8s.appscode.com,resources=petsets,verbs=create;update,versions=v1,name=mpetset-v1.kb.io,admissionReviewVersions=v1

// PetSetCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind PetSet when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type PetSetCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &PetSetCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind PetSet.
func (d *PetSetCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	petset, ok := obj.(*appsv1.PetSet)

	if !ok {
		return fmt.Errorf("expected an PetSet object but got %T", obj)
	}
	petsetlog.Info("Defaulting for PetSet", "name", petset.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-apps-k8s-appscode-com-v1-petset,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.k8s.appscode.com,resources=petsets,verbs=create;update,versions=v1,name=vpetset-v1.kb.io,admissionReviewVersions=v1

// PetSetCustomValidator struct is responsible for validating the PetSet resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type PetSetCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &PetSetCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type PetSet.
func (v *PetSetCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	petset, ok := obj.(*appsv1.PetSet)
	if !ok {
		return nil, fmt.Errorf("expected a PetSet object but got %T", obj)
	}
	petsetlog.Info("Validation for PetSet upon creation", "name", petset.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type PetSet.
func (v *PetSetCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	petset, ok := newObj.(*appsv1.PetSet)
	if !ok {
		return nil, fmt.Errorf("expected a PetSet object for the newObj but got %T", newObj)
	}
	petsetlog.Info("Validation for PetSet upon update", "name", petset.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type PetSet.
func (v *PetSetCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	petset, ok := obj.(*appsv1.PetSet)
	if !ok {
		return nil, fmt.Errorf("expected a PetSet object but got %T", obj)
	}
	petsetlog.Info("Validation for PetSet upon deletion", "name", petset.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
