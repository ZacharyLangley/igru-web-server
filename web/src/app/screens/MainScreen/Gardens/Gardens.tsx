import React, { useCallback, useState } from 'react';

import DashboardLayout from '../../../../common/components/Screen/Body/Layouts/Dashboard/DashboardLayout';
import GardenListCard from './components/GardenListCard/GardenListCard'
import GardenFormModal from './components/GardenFormModal/GardenFormModal';

export interface GardenProps {}

const Gardens: React.FC<GardenProps> = () => {
  const [isModalOpen, setModalOpen] = useState<boolean>(false)

  const toggleModal = useCallback(() => setModalOpen(!isModalOpen), [isModalOpen])

  return (
    <>
      <DashboardLayout
        bottomContent={<GardenListCard onAddGarden={toggleModal} />}
      />
      <GardenFormModal isOpen={isModalOpen} toggle={toggleModal}/>
    </>
  );
};

export default Gardens;
