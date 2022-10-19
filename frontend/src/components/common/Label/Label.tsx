import clsx from 'clsx'
import React from 'react'
import s from './Label.module.css'

interface Props {
  className?: string
  onClick?: () => void
  children?: React.ReactNode
  width?: string
  name: string
}

function PrimaryButton(props: Props): JSX.Element {
  const className =
    'text-white font-bold text-md rounded-full bg-primary-2 shadow-lg p-1 ' +
    (props.className ? ` ${props.className}` : '') +
    (props.width ? ` ${props.width}` : '')
  return (
    <div className={clsx(className, s.label)} onClick={props.onClick}>
      <div className={clsx('flex items-center px-5 py-1')}>{props.name}</div>
    </div>
  )
}

export default PrimaryButton
