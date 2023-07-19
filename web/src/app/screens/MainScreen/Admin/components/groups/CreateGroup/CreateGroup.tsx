import React, { ChangeEvent, useState } from 'react';
import { Form, FormGroup, Input, Label } from 'reactstrap';

import Button from '../../../../../../../common/components/Button/Button';
import Card from '../../../../../../../common/components/Card/Card';

import './styles.scss';
import useGroup from 'src/store/useGroup/useGroup';

interface CreateGroupProps {}

const CreateGroup: React.FC<CreateGroupProps> = () => {

    const [formData, setFormData] = useState({name: undefined});

    const {createGroup} = useGroup();

    const updateFormData = (key: string, value: number | string) => {
        setFormData({
          ...formData,
          [key]: value,
        });
    };

    const onFieldChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e === undefined) return;
        e.preventDefault();
        updateFormData(e.target.name, e.target.value);
    };

    const onCreateGroup = () => {
        if (!formData.name) return;
        createGroup(formData.name);
    }

    return (
        <div className="create-group-container">
            <Card
                header='Create Group'
                footer={
                    <div className='create-group-actions'>
                        <Button title={'Create Group'} onClick={onCreateGroup} className={'create-btn'}/>
                    </div>
                }>
                    <Form>
                        <FormGroup floating>
                        <Input
                            id={'name'}
                            name={'name'}
                            type={'text'}
                            value={formData.name}
                            placeholder={''}
                            onChange={onFieldChange}
                            />
                            <Label for={'name'}>{'Name'}</Label>
                        </FormGroup>
                    </Form>
            </Card>
        </div>
    )
}

export default CreateGroup;