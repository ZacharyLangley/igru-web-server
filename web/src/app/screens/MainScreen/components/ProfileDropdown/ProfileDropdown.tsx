import React, { useCallback, useMemo } from 'react';
import { useState } from 'react';
import { UnstyledButton, Popover, Avatar, Stack, Divider, Title } from '@mantine/core';
import language from '../../../../../common/language/index';
import useSession from 'src/store/useSession/useSession';
import { useNavigate } from 'react-router-dom';
import { RoutePath } from 'src/app/types/routes';
import useUser from 'src/store/useUser/useUser';

const lang = language();

export const ProfileDropdown = () => {
    const [opened, setOpened] = useState(false);

    const {user} = useUser();
    const {signOut} = useSession();
    const navigate = useNavigate();

    const toggleDropdown = useCallback(() => {
        setOpened((o) => !o);
    }, [setOpened]);

    const handleLogout = useCallback(() => {
        setOpened((o) => !o);
        signOut();
        navigate(RoutePath.HOME);
    }, [signOut, setOpened, navigate])

    const handleProfile = useCallback(() => {
        setOpened((o) => !o);
    }, [setOpened]);

    const name = useMemo(() => {
        const setName = user?.fullName || user?.email
        return `${setName?.charAt(0)?.toLocaleUpperCase()}${setName?.slice(1)}`;
    }, [user])

    return (
      <Popover opened={opened} onChange={setOpened}>
        <Popover.Target>
          <UnstyledButton onClick={toggleDropdown}>
            <Avatar radius='xl'/>            
          </UnstyledButton>
        </Popover.Target>
  
        <Popover.Dropdown>
            <Stack>
                <Title order={5}>{lang.navbar.profileDropdown.welcome} {name}</Title>
                <Divider my={"xs"} />
                <UnstyledButton onClick={handleProfile}>
                    {lang.navbar.profileDropdown.profile}
                </UnstyledButton>
                <UnstyledButton onClick={handleLogout}>
                    {lang.navbar.profileDropdown.logout}
                </UnstyledButton>
            </Stack>
        </Popover.Dropdown>
      </Popover>
    )
};

export default ProfileDropdown;