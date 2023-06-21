import React, { useCallback, useState } from 'react';

import DashboardLayout from '../../../../common/components/Screen/Body/Layouts/Dashboard/DashboardLayout'

export interface AdminProps {}

const Admin: React.FC<AdminProps> = () => {
  const [isModalOpen, setModalOpen] = useState<boolean>(false)

  const toggleModal = useCallback(() => setModalOpen(!isModalOpen), [isModalOpen])

  return (
    <>
      <DashboardLayout />
    </>
  );
};

export default Admin;
