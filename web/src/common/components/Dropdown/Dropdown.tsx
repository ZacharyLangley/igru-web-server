import React, { useCallback } from 'react';
import {Form, Input, Label} from 'reactstrap';
import './styles.scss';

interface SelectDropdownProps {
    options?: any;
    onSelect?: (value: any) => void
    label?: string;
}

const SelectDropdown: React.FC<SelectDropdownProps> = ({label, options, onSelect, ...props}) => {

    const renderOptions = useCallback(() => {
        if (!options || options?.length === 0) return <option className='select-dropdown-item'>{'--- No Groups ---'}</option>
        else return options?.map((option: any) => (<option className='select-dropdown-item' value={option.id}>{option.name}</option>))
    }, [options])

    const handleSelect = (event: any) => {
        onSelect?.(event?.target?.value);
    }

    return (
        <Form className='select-dropdown-container'>
                {label && <Label className={"label"} for="" sm={2}>{label}</Label>}
                <Input className="mb-3" name={'select'} type="select" onChange={handleSelect}>
                    {renderOptions()}
                </Input>
        </Form>
    )
}

export default SelectDropdown;