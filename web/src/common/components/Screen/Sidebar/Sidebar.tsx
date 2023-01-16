import React from 'react';
import {
  Offcanvas,
  OffcanvasHeader,
  OffcanvasBody,
  ListGroup,
  ListGroupItem,
} from 'reactstrap';
import Icon from '../../Icon/Icon';

import Logo from '../../Logo/Logo';
import './styles.scss';

export interface SidebarOptions {
  label: string;
  path: string;
  icon: string;
}

export interface SidebarProps {
  testID?: string;
  isOpen?: boolean;
  onClose?: () => void;
  onOptionSelection: (path?: string) => void;
  options?: SidebarOptions[];
}

const Sidebar: React.FC<SidebarProps> = ({
  testID = 'global-layout-sidebar',
  isOpen,
  onClose,
  onOptionSelection,
  options = [],
}) => {
  const renderSidebarOptions = (options?: SidebarOptions[]) => {
    if (!options) return;
    return options.map((option: SidebarOptions, index: number) => (
      <ListGroupItem
        className={'sidebar-list-item'}
        tag={'a'}
        onClick={() => onOptionSelection(option.path)}
        key={index}
      >
        <Icon src={option.icon} />
        <div className='sidebar-list-item-label'>{option.label}</div>
      </ListGroupItem>
    ));
  };

  return (
    <Offcanvas
      id={`${testID}:container`}
      className={'sidebar-container'}
      direction='start'
      isOpen={isOpen}
    >
      <OffcanvasHeader
        id={`${testID}:header`}
        className={'sidebar-header'}
        toggle={onClose}
      >
        <Logo
          testID={`${testID}:logo`}
          onClick={() => onOptionSelection('/')}
        />
      </OffcanvasHeader>
      <OffcanvasBody id={`${testID}:body`} className={'sidebar-body'}>
        <ListGroup className={'sidebar-list-container'} flush>
          {renderSidebarOptions(options)}
        </ListGroup>
      </OffcanvasBody>
    </Offcanvas>
  );
};

export default Sidebar;
