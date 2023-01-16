import React from 'react';
import Card from '../../Card/Card';

export interface CardPlaceHolderProps {
  testID?: string;
}

const ChartCardPlaceholder: React.FC<CardPlaceHolderProps> = ({
  testID = 'chart-card-placeholder',
}) => {
  return (
    <Card testID={testID} style={{height: 206, width: 216}}>
      <div
        style={{
          height: 172,
          width: 182,
          backgroundColor: 'grey',
          borderRadius: 5,
        }}
      />
    </Card>
  );
};

export default ChartCardPlaceholder;
