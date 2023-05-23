import React, { ChangeEvent } from 'react';
import { Form, FormGroup, Input, Label } from 'reactstrap';

export interface GardenFormProps {
    formData: any;
    onChange: (key: string, value: number | string) => void
}

const GardenForm: React.FC<GardenFormProps> = ({formData, onChange}) => {

    const onFieldChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e === undefined) return;
        e.preventDefault();
        onChange(e.target.name, e.target.value);
      };

    return (
        <div className={'garden-form-container'}>
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
                <FormGroup floating>
                    <Input
                        id={'comment'}
                        name={'comment'}
                        type={'text'}
                        value={formData.comment}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'comment'}>{'Comment'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'location'}
                        name={'location'}
                        type={'text'}
                        value={formData.location}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'location'}>{'Location'}</Label>
                </FormGroup>
            </Form>
        </div>
    );
}

export default GardenForm