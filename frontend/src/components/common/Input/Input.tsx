import clsx from 'clsx'
import React from 'react'

import s from './Input.module.css'

interface Props {
  className?: string
  placeholder?: string
  id?: string
  value?: string | number
  onChange?: any
  children?: React.ReactNode
}

function Input(props: Props): JSX.Element {
  const className =
    'rounded-full border-2 border-primary-2 py-2 px-4' +
    (props.className ? ` ${props.className}` : '')
  return (
    <input
      className={clsx(s.input, className)}
      placeholder={props.placeholder}
      id={props.id}
      value={props.value}
      onChange={props.onChange}
    >
      {props.children}
    </input>
  )
}

export default Input
