import React, { useEffect, useMemo } from 'react';
import { Select, ComboboxItem, Group, Text } from '@mantine/core';

import useGroup from 'src/store/useGroup/useGroup';
import useUser from 'src/store/useUser/useUser';
import language from '../../../../../common/language/index';

const lang = language();

export const GroupSelect = React.memo(() => {

    const {user} = useUser();
    const {activeUserGroup, groupsBySelectedUser, getAllGroupsByUser, setActiveGroup} = useGroup();

    const defaultGroup = useMemo(() => !activeUserGroup ? ({ value: groupsBySelectedUser?.[0]?.id, label: groupsBySelectedUser?.[0]?.name }) : ({ value: activeUserGroup.id, label: activeUserGroup.name }), [activeUserGroup, groupsBySelectedUser])
    const options = useMemo(() => groupsBySelectedUser?.map((group: any) => ({ value: group.id, label: group.name })), [groupsBySelectedUser])

    useEffect(() => {
        if (!activeUserGroup) {
            setActiveGroup(groupsBySelectedUser?.[0])
            return;
        } else return;
    }, [activeUserGroup, groupsBySelectedUser, setActiveGroup])

    const handleGroupSelect = (value: string | null, option?: ComboboxItem) => {
        if (!value || value === null) return;
        const group = groupsBySelectedUser?.find((group) => group.id === value )
        setActiveGroup(group);
    }
  
    useEffect(() => {
      if (user?.id) getAllGroupsByUser(user?.id);
    }, [user, getAllGroupsByUser]);

    return (
        <Group flex={2} align="center" justify='flex-end'>
            <Text style={{textAlign: 'right'}}>{lang.inputs.activeGroupSelect.label}</Text>
            <Select
                data={options}
                value={defaultGroup?.value}
                placeholder={defaultGroup?.value}
                onChange={handleGroupSelect}
                styles={{ dropdown: { maxHeight: 300, overflowY: 'auto' } }}
                searchable
                nothingFoundMessage={lang.inputs.activeGroupSelect.nothingFound}
            />
        </Group>
    )
})

export default GroupSelect;