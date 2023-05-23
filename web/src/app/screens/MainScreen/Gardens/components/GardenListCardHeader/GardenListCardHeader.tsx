import React from 'react';
import './styles.scss';
import Button from '../../../../../../common/components/Button/Button';

interface GardenListCardHeaderProps {
  onAddGarden?: () => void;
}

const GardenListCardHeader: React.FC<GardenListCardHeaderProps> = ({ onAddGarden }) => {
  return (
    <div className={'garden-list-card-header'}>
        <div className={'header-title'}>
            {'Gardens List'}
        </div>
        <div className={'header-actions'}>
            <Button title={'+ Add Garden'} onClick={onAddGarden}/> 
        </div>
    </div>
  )
};

export default GardenListCardHeader