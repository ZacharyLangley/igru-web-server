import React from 'react';
import {Navbar, NavbarBrand, Button} from 'reactstrap';

import Logo from '../../Logo/Logo';
import { ProfileDropdown } from 'src/common/components/ProfileDropdown/ProfileDropdown';
import './styles.scss';

export interface HeaderProps {
  testID?: string;
  onClick?: () => void;
}

const menubarUnicode = '\u2261'

const Header: React.FC<HeaderProps> = ({
  testID = 'global-layout-header',
  onClick,
}) => {
  return (
    <Navbar
      id={`${testID}:container`}
      color={'faded'}
      className={'header-container'}
    >
      <Button
        id={`${testID}:menu-btn`}
        className={'menu-btn'}
        outline
        onClick={onClick}
      >
        {menubarUnicode}
      </Button>
      <NavbarBrand id={`${testID}:logo:container`} href='/' className='me-auto'>
        <Logo testID={`${testID}:logo`} />
      </NavbarBrand>
      <ProfileDropdown />
    </Navbar>
  );
};

export default Header;
