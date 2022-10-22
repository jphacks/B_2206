import clsx from 'clsx'
import React from 'react'

interface Props {
  className?: string
  children?: React.ReactNode
  width?: string
  height?: string
}

export default function Modal(props: Props) {
  const className =
    'relative my-6 p-10 mx-auto w-3xl bg-white rounded-lg ' +
    (props.className ? ` ${props.className}` : ' ') +
    (props.width ? ` ${props.width}` : ' w-4/5 ') +
    (props.height ? ` ${props.height}` : ' h-auto ')

  return (
    <>
      <div className="fixed inset-0 z-50 flex items-center justify-center overflow-y-auto overflow-x-hidden bg-black/50 outline-none focus:outline-none">
        <div className={clsx(className)}>{props.children}</div>
      </div>
    </>
  )
}
