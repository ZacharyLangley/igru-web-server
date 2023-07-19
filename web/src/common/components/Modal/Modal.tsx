import React from 'react';
import {
    Modal as RSModal,
    ModalHeader as RSModalHeader,
    ModalBody as RSModalBody,
    ModalFooter as RSModalFooter
} from 'reactstrap'

interface ModalProps {
    open?: boolean;
    toggle?: () => void; 
    className?: string
    header?: any;
    body?: any;
    footer?: any;
}

export const Modal = ({header, body, footer, open, toggle, className} : ModalProps) => {
    return (
        <RSModal className={className} isOpen={open} toggle={toggle}>
            {
                header &&
                <RSModalHeader toggle={toggle}>
                    {header}
                </RSModalHeader>
            }
            {
                body &&
                <RSModalBody>
                    {body}
                </RSModalBody>
            }
            {
                footer &&
                <RSModalFooter>
                    {footer}
                </RSModalFooter>
            }
        </RSModal>
    )
}