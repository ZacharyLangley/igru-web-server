import React, { useEffect, useMemo } from 'react';
import {Navbar, NavbarBrand, Button} from 'reactstrap';

import Logo from '../../Logo/Logo';
import { ProfileDropdown } from 'src/common/components/ProfileDropdown/ProfileDropdown';
import SelectDropdown from 'src/common/components/Dropdown/Dropdown';

import './styles.scss';
import useGroup from 'src/store/useGroup/useGroup';
import useUser from 'src/store/useUser/useUser';

export interface HeaderProps {
  testID?: string;
  onClick?: () => void;
}

const menubarUnicode = '\u2261'

const mockOptions = [
  {
    name: 'Group A',
    value: 'Group A',
    groupId: '1234-1234-1234-1234',
  },
  {
    name: 'Group B',
    value: 'Group B',
    groupId: '1234-1234-1234-1234',
  },
  {
    name: 'Group C',
    value: 'Group C',
    groupId: '1234-1234-1234-1234',
  }
]

const Header: React.FC<HeaderProps> = ({
  testID = 'global-layout-header',
  onClick,
}) => {
  const {user} = useUser();
  const {groupsBySelectedUser, getAllGroupsByUser, activeUserGroup, setActiveGroup} = useGroup();

  const handleGroupSelect = (group: any) => {
    setActiveGroup(group);
  }

  useEffect(() => {
    if (user?.id) getAllGroupsByUser(user?.id);
  }, [user, getAllGroupsByUser]);

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
      <SelectDropdown label={'Active Group: '} options={groupsBySelectedUser} onSelect={handleGroupSelect}/>
      <ProfileDropdown />
    </Navbar>
  );
};

export default Header;
