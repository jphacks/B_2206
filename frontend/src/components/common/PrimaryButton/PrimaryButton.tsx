import clsx from 'clsx'
import React from 'react'
import s from './PrimaryButton.module.css'

interface Props {
  className?: string
  onClick?: () => void
  children?: React.ReactNode
  width?: string
}

function PrimaryButton(props: Props): JSX.Element {
  const className =
    'text-white faont-bold text-md rounded-full bg-primary-1 shadow-lg ' +
    (props.className ? ` ${props.className}` : '') +
    (props.width ? ` ${props.width}` : '')
  return (
    <button className={clsx(className, s.btn)} onClick={props.onClick}>
      <div className={clsx('flex items-center px-5 py-1')}>
        {props.children}
      </div>
    </button>
  )
}

export default PrimaryButton
