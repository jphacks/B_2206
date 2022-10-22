import clsx from 'clsx'
import React from 'react'

import s from './FillCard.module.css'

interface Props {
  children?: React.ReactNode
  width?: string
  height?: string
  m?: string
}

function Card(props: Props): JSX.Element {
  const m = props.m ? props.m : ' m-10 '
  return (
    <div className={clsx('mx-auto', props.width ? props.width : 'w-4/5')}>
      <div
        className={clsx(
          'border-primary-2 border-opacity-1 shadow-primary-2 hover:border-accent-2 rounded-lg border shadow-md hover:border',
          s.card,
          props.height ? props.height : 'h-auto',
          m,
        )}
      >
        {props.children}
      </div>
    </div>
  )
}

export default Card
