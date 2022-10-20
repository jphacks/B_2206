import clsx from 'clsx'
import React, { useState } from 'react'
import s from './LabelButton.module.css'

interface Props {
  className?: string
  onClick?: () => void
  children?: React.ReactNode
  width?: string
  name: string
}

function LabelButton(props: Props): JSX.Element {
  const [className, setClassName] = useState<string>(
    'text-primary-2 font-bold text-md rounded-full shadow-lg p-1 border-t border-primary-1 ' +
      (props.className ? ` ${props.className}` : '') +
      (props.width ? ` ${props.width}` : ''),
  )
  const [isClick, setIsClick] = useState<Boolean>(false)

  function colorChangeHandler(className: string) {
    if (!isClick) {
      setClassName(
        'text-white font-bold text-md rounded-full bg-primary-2 shadow-lg p-1 ' +
          (props.className ? ` ${props.className}` : '') +
          (props.width ? ` ${props.width}` : ''),
      )
    } else {
      setClassName(
        'text-primary-2 font-bold text-md rounded-full shadow-lg p-1 border-t border-primary-1 ' +
          (props.className ? ` ${props.className}` : '') +
          (props.width ? ` ${props.width}` : ''),
      )
    }
    setIsClick(!isClick);
  }

  return (
    <button
      className={clsx(className, s.label)}
      onClick={() => {
        colorChangeHandler(className)
        props.onClick
      }}
    >
      <div className={clsx('flex items-center justify-center w-1/1 px-5 py-1')}>{props.name}</div>
    </button>
  )
}

export default LabelButton
