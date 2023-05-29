import React, { useCallback, useState } from 'react';

import DashboardLayout from '../../../../common/components/Screen/Body/Layouts/Dashboard/DashboardLayout';
import PlantListCard from './components/PlantListCard/PlantListCard';
import PlantFormModal from './components/PlantFormModal/PlantFormModal';

export interface PlantsProps {}

const Plants: React.FC<PlantsProps> = () => {
  const [isModalOpen, setModalOpen] = useState<boolean>(false)

  const toggleModal = useCallback(() => setModalOpen(!isModalOpen), [isModalOpen])

  return (
    <>
      <DashboardLayout
        bottomContent={<PlantListCard onAddPlant={toggleModal} />}
      />
      <PlantFormModal isOpen={isModalOpen} toggle={toggleModal}/>
    </>
  )
};

export default Plants;
