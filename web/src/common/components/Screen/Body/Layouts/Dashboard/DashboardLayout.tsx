import React from 'react';
import {Container, Row, Col} from 'reactstrap';
import ChartCardPlaceholder from '../../../../Placeholder/ChartCardPlaceholder/ChartCardPlaceholder';
import TableCardPlaceHolder from '../../../../Placeholder/TableCardPlaceHolder/TableCardPlaceHolder';

import './styles.scss';

export interface DashboardLayoutProps {
  testID?: string;
  topContent?: JSX.Element | JSX.Element[];
  bottomContent?: JSX.Element | JSX.Element[];
}

const DashboardLayout: React.FC<DashboardLayoutProps> = ({
  testID = 'dashboard-layout',
  topContent,
  bottomContent,
}) => {
  return (
    <Container
      id={`${testID}:container`}
      className={'dash-layout-container'}
      style={{display: 'flex', flexDirection: 'column'}}
    >
      <Row id={`${testID}:row:top`} className={'dash-top-content'}>
        <Col className={'dash-top-item'}>
          <ChartCardPlaceholder />
        </Col>
        <Col className={'dash-top-item'}>
          <ChartCardPlaceholder />
        </Col>
        <Col className={'dash-top-item'}>
          <ChartCardPlaceholder />
        </Col>
        <Col className={'dash-top-item'}>
          <ChartCardPlaceholder />
        </Col>
      </Row>
      <Row id={`${testID}:row:bottom`} className={'dash-bottom-content'}>
        <Col className=''>
          <TableCardPlaceHolder />
        </Col>
      </Row>
    </Container>
  );
};

export default DashboardLayout;
