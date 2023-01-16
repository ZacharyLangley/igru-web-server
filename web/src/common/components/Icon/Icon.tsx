import React from 'react';

export interface IconProps {
  testID?: string;
  src?: string;
  alt?: string;
  height?: number;
  width?: number;
}

const Icon: React.FC<IconProps> = ({
  testID = 'icon',
  src,
  alt = 'icon',
  height = 35,
  width = 35,
}) => {
  return (
    <div id={`${testID}:container`} className={'icon-container'}>
      <img
        id={`${testID}:img`}
        className={'icon-img'}
        src={src}
        alt={alt}
        height={height}
        width={width}
      />
    </div>
  );
};

export default Icon;
