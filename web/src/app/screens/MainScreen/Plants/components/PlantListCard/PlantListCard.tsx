import React, { useEffect } from 'react';
import Button from '../../../../../../common/components/Button/Button';
import Card from '../../../../../../common/components/Card/Card';

import PlantListCardHeader from '../PlantListCardHeader/PlantListCardHeader';
import useGroup from 'src/store/useGroup/useGroup';
import usePlant from 'src/store/usePlant/usePlant';
import { Plant } from 'client/garden/v1/schema_pb';


interface PlantListCardProps {
  testID?: string;
  onAddPlant?: () => void;
}

const PlantListCard: React.FC<PlantListCardProps> = ({
  testID = 'plant-list-card',
  onAddPlant,
}) => {

  const {activeUserGroup} = useGroup();
  const {plants, getAllPlants, deletePlant} = usePlant();

  useEffect(() => {
    getAllPlants(activeUserGroup?.id);
  }, [activeUserGroup, getAllPlants])

  const onDelete = (id: string, groupId: string) => {
    if (id && groupId) {
      deletePlant(id, groupId)
    } else return;
  }

  return (
    <Card testID={testID}>
      <PlantListCardHeader onAddPlant={onAddPlant}/>
      {!plants && <div>{'No Data Available'}</div>}
      {plants && plants?.length === 0 && <div>{'No Gardens Available. Please Add a Garden'}</div>}
      {plants && plants?.length > 0 && plants?.map((plant: Plant) => {
        return (
          <div className={'garden-list-item-container'}>
            <div className={'item'}>
              <span className={'label'}>{'Name: '}</span>
              <span className={'value'}>{plant?.name}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'ID: '}</span>
              <span className={'value'}>{plant?.id}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Group ID: '}</span>
              <span className={'value'}>{plant?.groupId}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Comment: '}</span>
              <span className={'value'}>{plant?.comment}</span>
            </div>
            <div className={'item'}>
              <span className={'label'}>{'Create At: '}</span>
              <span className={'value'}>{plant?.createdAt?.toJsonString()}</span>
            </div>
            <div className='actions'><Button title={'Delete'} color={'danger'} onClick={() => onDelete(plant.id, plant.groupId)}/></div>
          </div>
        )
      })}
    </Card>
  );
};

export default PlantListCard;