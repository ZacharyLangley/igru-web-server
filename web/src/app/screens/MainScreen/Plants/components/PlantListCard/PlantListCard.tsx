import React from 'react';
import Card from '../../../../../../common/components/Card/Card';

import PlantListCardHeader from '../PlantListCardHeader/PlantListCardHeader';

interface PlantListCardProps {
  testID?: string;
  onAddPlant?: () => void;
}

const PlantListCard: React.FC<PlantListCardProps> = ({
  testID = 'Plant-list-card',
  onAddPlant,
}) => {
  return (
    <Card testID={testID}>
      <PlantListCardHeader onAddPlant={onAddPlant}/>
      <div
        style={{
          flex: 1,
          minHeight: 0,
          height: 400,
          minWidth: 0,
          backgroundColor: 'grey',
          borderRadius: 5,
        }}
      />
    </Card>
  );
};

export default PlantListCard;