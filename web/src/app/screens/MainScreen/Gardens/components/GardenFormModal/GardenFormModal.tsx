import React, { useState } from 'react';
import { Modal, ModalBody, ModalFooter, ModalHeader } from 'reactstrap';
import Button from '../../../../../../common/components/Button/Button';
import GardenForm from '../GardenForm/GardenForm';
import './styles.scss';

interface GardenFormModalProps {
    isOpen?: boolean;
    toggle?: () => void;
}


const GardenFormModal: React.FC<GardenFormModalProps> = ({isOpen = false, toggle}) => {
    const [formData, setFormData] = useState({});

    const updateFormData = (key: string, value: number | string) => {
      setFormData({
        ...formData,
        [key]: value,
      });
    };
  
    const handleCreate = () => {
        toggle?.();
    }

    const handleCancel = () => {
        toggle?.();
    }

    return (
        <Modal isOpen={true} toggle={toggle}>
            <ModalHeader toggle={toggle}>{'Add Garden'}</ModalHeader>
            <ModalBody>
                <GardenForm formData={formData} onChange={updateFormData}/>
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

export default GardenFormModal;