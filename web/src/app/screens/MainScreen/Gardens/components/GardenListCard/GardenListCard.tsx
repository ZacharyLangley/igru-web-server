import React, { useEffect } from 'react';

import Button from '../../../../../../common/components/Button/Button';
import Card from '../../../../../../common/components/Card/Card';

import GardenListCardHeader from '../GardenListCardHeader/GardenListCardHeader';
import useGarden from 'src/store/useGarden/useGarden';
import useGroup from 'src/store/useGroup/useGroup';
import { Garden } from '../../../../../../client/garden/v1/schema_pb';

import './styles.scss';

interface GardenListCardProps {
  testID?: string;
  onAddGarden?: () => void;
}

const GardenListCard: React.FC<GardenListCardProps> = ({
  testID = 'garden-list-card',
  onAddGarden,
}) => {
  const {activeUserGroup} = useGroup();
  const {gardens, getAllGardens, deleteGarden} = useGarden();

  useEffect(() => {
    getAllGardens(activeUserGroup?.id);
  }, [activeUserGroup, getAllGardens])


  const onDelete = (id: string) => {
    if (id) {
      deleteGarden(id)
    } else return;
}

  return (
    <Card testID={testID}>
      <GardenListCardHeader onAddGarden={onAddGarden}/>
      {!gardens && <div>{'No Data Available'}</div>}
      {gardens && gardens?.length === 0 && <div>{'No Gardens Available. Please Add a Garden'}</div>}
      {gardens && gardens?.length > 0 && gardens?.map((garden: Garden) => {
        return (
          <div className={'garden-list-item-container'}>
            <div className={'item'}>
              <span className={'label'}>{'Name: '}</span>
              <span className={'value'}>{garden?.name}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'ID: '}</span>
              <span className={'value'}>{garden?.id}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Group ID: '}</span>
              <span className={'value'}>{garden?.groupId}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Comment: '}</span>
              <span className={'value'}>{garden?.comment}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Create At: '}</span>
              <span className={'value'}>{garden?.createdAt?.toJsonString()}</span>
            </div>
            <div className='actions'><Button title={'Delete'} color={'danger'} onClick={() => onDelete(garden.id)}/></div>
          </div>
        )
      })}
    </Card>
  );
};

export default GardenListCard;