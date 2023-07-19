import React, { useState } from 'react';
import { Modal, ModalBody, ModalFooter, ModalHeader } from 'reactstrap';
import Button from '../../../../../../common/components/Button/Button';
import GardenForm from '../GardenForm/GardenForm';
import './styles.scss';
import useGarden from 'src/store/useGarden/useGarden';
import useGroup from 'src/store/useGroup/useGroup';

interface GardenFormModalProps {
    isOpen?: boolean;
    toggle?: () => void;
}

// This is temporary as to quickly get rid of once the garden form is complete
const defaultGardenInfo = {
    name:          undefined,
    groupId:       undefined,
    comment:       'Garden Mock Comment',
    location:      'OUTSIDE',
    growType:      'SOILLESS',
    growSize:      '{\'value\':\'7.15\',\'metric\':\'sq. ft.\'}',
    growStyle:     'HYDROPONIC',
    containerSize: '{\'value\':\'7.15\',\'metric\':\'sq. ft.\'}',
    tags:          '[\'Tag A\', \'Tag B\']',
}


const GardenFormModal: React.FC<GardenFormModalProps> = ({isOpen = false, toggle}) => {
    const {createGarden} = useGarden();
    const {activeUserGroup} = useGroup();
    const [formData, setFormData] = useState(defaultGardenInfo);

    const updateFormData = (key: string, value: number | string) => {
      setFormData({
        ...formData,
        [key]: value,
      });
    };
  
    const handleCreate = () => {
        if (!formData.name) return;
        const garden = {...formData, groupId: activeUserGroup}
        // @ts-ignore
        createGarden(garden);
    }

    const handleCancel = () => {
        setFormData(defaultGardenInfo);
        toggle?.();
    }

    return (
        <Modal isOpen={isOpen} toggle={toggle}>
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