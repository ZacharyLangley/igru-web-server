import React, { useCallback, useMemo, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Dropdown, DropdownItem, DropdownMenu, DropdownToggle } from 'reactstrap';

import language from 'src/common/language/index';
import { RoutePath } from 'src/app/types/routes';
import Icon from 'src/common/components/Icon/Icon';
import profileIcon from 'src/common/assets/icons/nav/profile_icon.png';
import useSession from 'src/store/useSession/useSession';
import useUser from 'src/store/useUser/useUser';

const welcomeLabel = language('navbar.dropdown.header');
const profileLabel = language('navbar.dropdown.profile');
const signOutLabel = language('navbar.dropdown.signout');

interface ProfileDropdownProps {}

const dropdownToggleStyle = {backgroundColor: 'transparent', borderWidth: 0, borderRadius: 50};

const formatUserName = (name?: string) => {
    if (!name) return;
    return name?.charAt(0).toUpperCase() + name?.slice(1)
}

export const ProfileDropdown:React.FC<ProfileDropdownProps>  = () => {
    const {signOut} = useSession();
    const {user} = useUser();
    const navigate = useNavigate();

    const [dropdownOpen, setDropdownOpen] = useState(false);

    const toggle = () => setDropdownOpen((prevState) => !prevState);
    const name = useMemo(() => formatUserName(user?.fullName || user?.email), [user])
    const onSignOutClick = useCallback(() => {
        signOut();
        navigate(RoutePath.HOME);
    }, [signOut, navigate])

    return (
        <div className="d-flex p-1">
            <Dropdown tag={'div'} isOpen={dropdownOpen} toggle={toggle} direction={'down'}>
                <DropdownToggle style={dropdownToggleStyle}>
                    <Icon src={profileIcon} />
                </DropdownToggle>
                <DropdownMenu container="body">
                    <DropdownItem header>{`${welcomeLabel}, ${name}`}</DropdownItem>
                    <DropdownItem>{profileLabel}</DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem onClick={onSignOutClick}>{signOutLabel}</DropdownItem>
                </DropdownMenu>
            </Dropdown>
        </div>
    )
}