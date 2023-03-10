import React, { useCallback, useState } from 'react';
import { useDispatch } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { Dropdown, DropdownItem, DropdownMenu, DropdownToggle } from 'reactstrap';

import { dispatchSignOutAction } from 'src/domain/actions/sessions';
import language from 'src/common/language/index';
import { RoutePath } from 'src/app/types/routes';
import Icon from 'src/common/components/Icon/Icon';
import profileIcon from 'src/common/assets/icons/nav/profile_icon.png';

const welcomeLabel = language('navbar.dropdown.header');
const profileLabel = language('navbar.dropdown.profile');
const signOutLabel = language('navbar.dropdown.signout');

interface ProfileDropdownProps {}

const dropdownToggleStyle = {backgroundColor: 'transparent', borderWidth: 0, borderRadius: 50};

export const ProfileDropdown:React.FC<ProfileDropdownProps>  = () => {
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const [dropdownOpen, setDropdownOpen] = useState(false);

    const toggle = () => setDropdownOpen((prevState) => !prevState);

    const onSignOutClick = useCallback(() => {
        dispatch(dispatchSignOutAction())
        navigate(RoutePath.HOME);
    }, [dispatch, navigate])

    return (
        <div className="d-flex p-1">
            <Dropdown tag={'div'} isOpen={dropdownOpen} toggle={toggle} direction={'down'}>
                <DropdownToggle style={dropdownToggleStyle}>
                    <Icon src={profileIcon} />
                </DropdownToggle>
                <DropdownMenu container="body">
                    <DropdownItem header>{welcomeLabel}</DropdownItem>
                    <DropdownItem>{profileLabel}</DropdownItem>
                    <DropdownItem divider />
                    <DropdownItem onClick={onSignOutClick}>{signOutLabel}</DropdownItem>
                </DropdownMenu>
            </Dropdown>
        </div>
    )
}