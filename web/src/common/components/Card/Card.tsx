import React from 'react';
import {
  Card as RSCard,
  CardProps as RSCardProps,
  CardHeader as RSCardHeader,
  CardFooter as RSCardFooter,
  CardTitle as RSCardTitle,
  CardSubtitle as RSCardSubtitle,
  CardBody as RSCardBody,
} from 'reactstrap';

export interface CardProps extends RSCardProps {
  testID?: string;
  header?: string;
  title?: string;
  subtitle?: string;
  footer?: string;
}

const Card: React.FC<CardProps> = ({
  testID = 'card',
  header,
  footer,
  title,
  subtitle,
  children,
  ...props
}) => {
  return (
    <RSCard {...props}>
      {header && <RSCardHeader>{header}</RSCardHeader>}
      <RSCardBody>
        {title && <RSCardTitle tag={'h5'}>{title}</RSCardTitle>}
        {subtitle && (
          <RSCardSubtitle tag={'h6'} className={'mb-2 text-muted'}>
            {subtitle}
          </RSCardSubtitle>
        )}
        {children}
      </RSCardBody>
      {footer && <RSCardFooter>{footer}</RSCardFooter>}
    </RSCard>
  );
};

export default Card;
