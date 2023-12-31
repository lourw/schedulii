import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { EventForm } from '@schedulii/schedulii-components';

interface FormData {
  eventName: string;
  startDate: string;
  endDate: string;
  earliestTime: string;
  latestTime: string;
  location: string;
}

const AddEvent: React.FC = () => {
  const API_URL = import.meta.env.VITE_API_URL;
  const [formData, setFormData] = useState<FormData>({
    eventName: '',
    startDate: '',
    endDate: '',
    earliestTime: '',
    latestTime: '',
    location: '',
  });
  const [isFormValid, setIsFormValid] = useState(false);
  const navigate = useNavigate();

  useEffect(() => {
    setIsFormValid(validateForm());
  }, [formData]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (isFormValid) {
      axios
        .post(`${API_URL}/events/add`, {
          event_name: formData.eventName,
          start_time: new Date(formData.startDate),
          end_time: new Date(formData.endDate),
          location: formData.location,
        })
        .then(() => {
          navigate('/home');
        })
        .catch((err) => {
          console.log(err);
        });
    }
  };

  const validateForm = () => {
    const isEmptyField = Object.values(formData).some(
      (field) => field.trim().length === 0
    );
    const invalidDateRange = formData.endDate < formData.startDate;

    return !(isEmptyField || invalidDateRange);
  };

  const handleInputChange = (field: string, value: string) => {
    setFormData({ ...formData, [field]: value });
  };

  return (
    <EventForm
      formData={formData}
      isFormValid={isFormValid}
      onInputChange={handleInputChange}
      onSubmit={handleSubmit}
    />
  );
};

export default AddEvent;
