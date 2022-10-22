import clsx from 'clsx'
import React from 'react'
import s from './PrimaryButton.module.css'

interface Props {
  pointName: string[]
  nowPoint: number
}

function ProgressBar(props: Props): JSX.Element {
  const OrangePoint = (num: number) => (
    <span className="bg-accent-2 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6 text-white">
      {num}
    </span>
  )

  const NormalPoint = (num: number) => (
    <span className="bg-primary-1 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6 text-white">
      {num}
    </span>
  )

  return (
    <div className="flex justify-center">
      <div className="w-2/3">
        <h2 className="sr-only">Steps</h2>
        <div className="after:bg-primary-1 relative after:absolute after:inset-x-0 after:top-1/2 after:block after:h-0.5 after:-translate-y-1/2 after:rounded-lg">
          <ol className="text-main-black relative z-10 flex justify-between text-sm font-medium">
            {props.pointName.map((pointName, index) => {
              return (
                <li
                  className="bg-background-1 flex items-center p-2"
                  key={index}
                >
                  {props.nowPoint == index + 1
                    ? OrangePoint(index + 1)
                    : NormalPoint(index + 1)}
                  <span className="hidden sm:ml-2 sm:block">
                    {' '}
                    {props.pointName[index]}{' '}
                  </span>
                </li>
              )
            })}
          </ol>
        </div>
      </div>
    </div>
  )
}

export default ProgressBar
