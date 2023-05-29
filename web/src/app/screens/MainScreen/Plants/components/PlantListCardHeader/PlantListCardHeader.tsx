import React from 'react';
import './styles.scss';
import Button from '../../../../../../common/components/Button/Button';

interface PlantListCardHeaderProps {
  onAddPlant?: () => void;
}

const PlantListCardHeader: React.FC<PlantListCardHeaderProps> = ({ onAddPlant }) => {
  return (
    <div className={'plant-list-card-header'}>
        <div className={'header-title'}>
            {'Plants List'}
        </div>
        <div className={'header-actions'}>
            <Button title={'+ Add Plant'} onClick={onAddPlant}/> 
        </div>
    </div>
  )
};

export default PlantListCardHeader