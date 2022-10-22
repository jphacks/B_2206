import clsx from 'clsx'
import React from 'react'

import s from './Textarea.module.css'

interface Props {
  className?: string
  placeholder?: string
  value?: string | number
  onChange?: any
  children?: React.ReactNode
  width?: string
  height?: string
}

function Textarea(props: Props): JSX.Element {
  const className =
    'rounded-3xl border border-primary-2 py-2 px-4 ' +
    (props.className ? ` ${props.className}` : ' ') +
    (props.width ? ` ${props.width}` : ' w-1/1 ') +
    (props.height ? ` ${props.height}` : ' h-64 ')
  return (
    <textarea
      className={clsx(s.textarea, className)}
      placeholder={props.placeholder}
      value={props.value}
      onChange={props.onChange}
    >
      {props.children}
    </textarea>
  )
}

export default Textarea
