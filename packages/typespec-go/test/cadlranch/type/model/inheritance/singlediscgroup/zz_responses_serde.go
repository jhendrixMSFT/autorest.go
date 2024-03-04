// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package singlediscgroup

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleDiscriminatorClientGetLegacyModelResponse.
func (s *SingleDiscriminatorClientGetLegacyModelResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalDinosaurClassification(data)
	if err != nil {
		return err
	}
	s.DinosaurClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleDiscriminatorClientGetMissingDiscriminatorResponse.
func (s *SingleDiscriminatorClientGetMissingDiscriminatorResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalBirdClassification(data)
	if err != nil {
		return err
	}
	s.BirdClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleDiscriminatorClientGetModelResponse.
func (s *SingleDiscriminatorClientGetModelResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalBirdClassification(data)
	if err != nil {
		return err
	}
	s.BirdClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleDiscriminatorClientGetRecursiveModelResponse.
func (s *SingleDiscriminatorClientGetRecursiveModelResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalBirdClassification(data)
	if err != nil {
		return err
	}
	s.BirdClassification = res
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SingleDiscriminatorClientGetWrongDiscriminatorResponse.
func (s *SingleDiscriminatorClientGetWrongDiscriminatorResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalBirdClassification(data)
	if err != nil {
		return err
	}
	s.BirdClassification = res
	return nil
}
