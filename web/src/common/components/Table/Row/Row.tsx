import React from 'react';

export interface RowProps {
  index: number;
}

const Row: React.FC<RowProps> = ({index}) => {
  return (
    <tr>
      {/** Make sure your table rows are the same height as what you passed into the list... */}
      <td style={{height: '36px'}}>Row {index}</td>
      <td>Col 2</td>
      <td>Col 3</td>
      <td>Col 4</td>
    </tr>
  );
};

export default Row;
