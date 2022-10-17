import clsx from 'clsx'
import React from 'react'

import s from './FillCard.module.css'

interface Props {
  children?: React.ReactNode
}

function Card(props: Props): JSX.Element {
  return (
    <div className={clsx('mx-auto w-4/5')}>
      <div
        className={clsx(
          'border-primary-2 border-opacity-1 shadow-primary-2 hover:border-accent-2 m-10 rounded-lg border px-10 py-5 shadow-md hover:border',
          s.card,
        )}
      >
        {props.children}
      </div>
    </div>
  )
}

export default Card
