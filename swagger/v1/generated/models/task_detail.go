// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskDetail task detail
// swagger:model TaskDetail
type TaskDetail struct {

	// cluster a r n
	// Required: true
	ClusterARN *string `json:"clusterARN"`

	// container instance a r n
	// Required: true
	ContainerInstanceARN *string `json:"containerInstanceARN"`

	// containers
	// Required: true
	Containers []*TaskContainer `json:"containers"`

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// desired status
	// Required: true
	DesiredStatus *string `json:"desiredStatus"`

	// last status
	// Required: true
	LastStatus *string `json:"lastStatus"`

	// overrides
	// Required: true
	Overrides *TaskOverride `json:"overrides"`

	// started at
	StartedAt string `json:"startedAt,omitempty"`

	// started by
	StartedBy string `json:"startedBy,omitempty"`

	// stopped at
	StoppedAt string `json:"stoppedAt,omitempty"`

	// stopped reason
	StoppedReason string `json:"stoppedReason,omitempty"`

	// task a r n
	// Required: true
	TaskARN *string `json:"taskARN"`

	// task definition a r n
	// Required: true
	TaskDefinitionARN *string `json:"taskDefinitionARN"`
}

// Validate validates this task detail
func (m *TaskDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainerInstanceARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateContainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDesiredStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOverrides(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaskDefinitionARN(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskDetail) validateClusterARN(formats strfmt.Registry) error {

	if err := validate.Required("clusterARN", "body", m.ClusterARN); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateContainerInstanceARN(formats strfmt.Registry) error {

	if err := validate.Required("containerInstanceARN", "body", m.ContainerInstanceARN); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateContainers(formats strfmt.Registry) error {

	if err := validate.Required("containers", "body", m.Containers); err != nil {
		return err
	}

	for i := 0; i < len(m.Containers); i++ {

		if swag.IsZero(m.Containers[i]) { // not required
			continue
		}

		if m.Containers[i] != nil {

			if err := m.Containers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("containers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TaskDetail) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateDesiredStatus(formats strfmt.Registry) error {

	if err := validate.Required("desiredStatus", "body", m.DesiredStatus); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateLastStatus(formats strfmt.Registry) error {

	if err := validate.Required("lastStatus", "body", m.LastStatus); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateOverrides(formats strfmt.Registry) error {

	if err := validate.Required("overrides", "body", m.Overrides); err != nil {
		return err
	}

	if m.Overrides != nil {

		if err := m.Overrides.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("overrides")
			}
			return err
		}
	}

	return nil
}

func (m *TaskDetail) validateTaskARN(formats strfmt.Registry) error {

	if err := validate.Required("taskARN", "body", m.TaskARN); err != nil {
		return err
	}

	return nil
}

func (m *TaskDetail) validateTaskDefinitionARN(formats strfmt.Registry) error {

	if err := validate.Required("taskDefinitionARN", "body", m.TaskDefinitionARN); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskDetail) UnmarshalBinary(b []byte) error {
	var res TaskDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
