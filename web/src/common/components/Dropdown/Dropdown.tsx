import React, { useCallback } from 'react';
import {Form, Input} from 'reactstrap';
import './styles.scss';

interface SelectDropdownProps {
    options?: any;
    onSelect?: (value: any) => void
}

const SelectDropdown: React.FC<SelectDropdownProps> = ({options, onSelect, ...props}) => {

    const renderOptions = useCallback(() => {
        return options?.map((option: any) => (
            <option className='select-dropdown-container' onClick={() => onSelect?.(option)}>
                {option.name}
            </option>
        ))
    }, [options, onSelect])

    return (
        <Form className='select-dropdown-container'>
            <Input className="mb-3" type="select">
                {options && renderOptions()}
            </Input>
        </Form>
    )
}

export default SelectDropdown;