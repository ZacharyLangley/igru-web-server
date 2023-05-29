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
                <FormGroup floating>
                    <Input
                        id={'growType'}
                        name={'growType'}
                        type={'text'}
                        value={formData.growType}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'growType'}>{'Grow Type'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'growType'}
                        name={'growType'}
                        type={'text'}
                        value={formData.growType}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'growType'}>{'Grow Type'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'growSize'}
                        name={'growSize'}
                        type={'text'}
                        value={formData.growSize}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'growSize'}>{'Grow Size'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'containerSize'}
                        name={'containerSize'}
                        type={'text'}
                        value={formData.containerSize}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'containerSize'}>{'Container Size'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'tags'}
                        name={'tags'}
                        type={'text'}
                        value={formData.tags}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'tags'}>{'Tags'}</Label>
                </FormGroup>
            </Form>
        </div>
    );
}

export default GardenForm