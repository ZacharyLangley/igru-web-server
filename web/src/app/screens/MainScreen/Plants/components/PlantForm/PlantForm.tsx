import React, { ChangeEvent } from 'react';
import { Form, FormGroup, Input, Label } from 'reactstrap';

export interface PlantFormProps {
    formData: any;
    onChange: (key: string, value: number | string) => void
}

const PlantForm: React.FC<PlantFormProps> = ({formData, onChange}) => {

    const onFieldChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (e === undefined) return;
        e.preventDefault();
        onChange(e.target.name, e.target.value);
      };

    return (
        <div className={'plant-form-container'}>
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
                        id={'notes'}
                        name={'notes'}
                        type={'text'}
                        value={formData.notes}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'notes'}>{'Notes'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'growCycleLength'}
                        name={'growCycleLength'}
                        type={'text'}
                        value={formData.growCycleLength}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'growCycleLength'}>{'Grow Cycle Length'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'parentage'}
                        name={'parentage'}
                        type={'text'}
                        value={formData.parentage}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'parentage'}>{'Parentage'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'origin'}
                        name={'origin'}
                        type={'text'}
                        value={formData.origin}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'origin'}>{'Origin'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'gender'}
                        name={'gender'}
                        type={'text'}
                        value={formData.gender}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'gender'}>{'Gender'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'daysFlowering'}
                        name={'daysFlowering'}
                        type={'text'}
                        value={formData.daysFlowering}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'daysFlowering'}>{'Days Flowering'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'daysCured'}
                        name={'daysCured'}
                        type={'text'}
                        value={formData.daysCured}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'daysCured'}>{'Days Cured'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'harvestedWeight'}
                        name={'harvestedWeight'}
                        type={'text'}
                        value={formData.harvestedWeight}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'harvestedWeight'}>{'Harvested Weight'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'budDensity'}
                        name={'budDensity'}
                        type={'text'}
                        value={formData.budDensity}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'budDensity'}>{'Bud Density'}</Label>
                </FormGroup>
                <FormGroup floating>
                    <Input
                        id={'budPistils'}
                        name={'budPistils'}
                        type={'text'}
                        value={formData.budPistils}
                        placeholder={''}
                        onChange={onFieldChange}
                        />
                        <Label for={'budPistils'}>{'Bud Pistils'}</Label>
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

export default PlantForm