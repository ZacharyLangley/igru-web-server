import React, { useState } from 'react';
import { Modal, ModalBody, ModalFooter, ModalHeader } from 'reactstrap';
import Button from '../../../../../../common/components/Button/Button';
import PlantForm from '../PlantForm/PlantForm';
import './styles.scss';
import usePlant from 'src/store/usePlant/usePlant';

interface PlantFormModalProps {
    isOpen?: boolean;
    toggle?: () => void;
}

const defaultPlantInfo = {
    name: undefined,
    comment: 'Mock Plant Comment',
    notes: 'Mock Plant Notes',
    growCycleLength: '{\'value\':\'28\',\'metric\':\'days\'}',
    parentage: 'Mock Parent Strain',
    origin: 'Clone',
    gender: 'Feminized',
    daysFlowering: 2.4,
    daysCured: 1.2,
    harvestedWeight: '{\'value\':\'1.05\',\'metric\':\'lbs.\'}',
    budDensity: 0.7,
    budPistils: false,
    tags: '[\'Tag A\', \'Tag B\']',
}

const PlantFormModal: React.FC<PlantFormModalProps> = ({isOpen = false, toggle}) => {
    const {createPlant} = usePlant()
    const [formData, setFormData] = useState(defaultPlantInfo);

    const updateFormData = (key: string, value: number | string) => {
      setFormData({
        ...formData,
        [key]: value,
      });
    };
  
    const handleCreate = () => {
        if (!formData.name) return;
        createPlant(formData);
        toggle?.();
    }

    const handleCancel = () => {
        setFormData(defaultPlantInfo);
        toggle?.();
    }

    return (
        <Modal isOpen={isOpen} toggle={toggle}>
            <ModalHeader toggle={toggle}>{'Add Plant'}</ModalHeader>
            <ModalBody>
                <PlantForm formData={formData} onChange={updateFormData}/>
            </ModalBody>
            <ModalFooter>
                <div className='modal-actions'>
                    <Button title={'Create'} onClick={handleCreate}/>
                    <div style={{width: '32px'}}/>
                    <Button title={'Cancel'} color='secondary' onClick={handleCancel}/>
                </div>
            </ModalFooter>
        </Modal>
    );
}

export default PlantFormModal;