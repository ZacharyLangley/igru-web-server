import React from 'react';
import Card from '../../../../../../common/components/Card/Card';

import GardenListCardHeader from '../GardenListCardHeader/GardenListCardHeader';

interface GardenListCardProps {
  testID?: string;
  onAddGarden?: () => void;
}

const GardenListCard: React.FC<GardenListCardProps> = ({
  testID = 'garden-list-card',
  onAddGarden,
}) => {
  return (
    <Card testID={testID}>
      <GardenListCardHeader onAddGarden={onAddGarden}/>
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

export default GardenListCard;