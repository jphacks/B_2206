import clsx from 'clsx'
import React from 'react'

interface Props {
  className?: string
  value?: string | number
  name?: string
  checked?: boolean
  onClick?: any
  onChange?: any
  children?: React.ReactNode
  id?: string
  defaultChecked?: boolean
}

function Radio(props: Props): JSX.Element {
  const className = '' + (props.className ? ` ${props.className}` : '')
  return (
    <>
      <div className="flex flex-row gap-1">
        <input
          type="radio"
          name={props.name}
          defaultChecked={props.defaultChecked}
          checked={props.checked}
          className={clsx(className)}
          value={props.value}
          onClick={props.onClick}
          onChange={props.onChange}
          id={props.id}
        />
        <label htmlFor={props.id}>{props.children}</label>
      </div>
    </>
  )
}

export default Radio
