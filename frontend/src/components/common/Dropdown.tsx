import clsx from 'clsx'
import React, { useRef, useState, useEffect } from 'react'
import { Menu, Close } from '@components/icons'

interface Props {
  children: React.ReactNode
}

const Dropdown = (props: Props) => {
  const [isOpen, setIsOpen] = useState(false)
  const dropdownRef = useRef<any>(null)

  useEffect(() => {
    document.addEventListener('mousedown', handleOutsideClick)
    return () => {
      document.removeEventListener('mousedown', handleOutsideClick)
    }
  }, [])

  const handleOutsideClick = (e: any) => {
    if (dropdownRef.current && !dropdownRef.current.contains(e.target)) {
      setIsOpen(false)
    }
  }

  return (
    <>
      <div ref={dropdownRef} className="relative inline-block text-left">
        <span>
          <button
            onClick={() => setIsOpen(!isOpen)}
            type="button"
            className="inline-flex w-full justify-center px-4 text-xl transition duration-150 ease-in-out focus:outline-none"
            id="options-menu"
            aria-haspopup="true"
            aria-expanded={isOpen}
          >
            <div className={clsx('flex items-center')}>
              {isOpen ? (
                <Close color={'#fff'} />
              ) : (
                <Menu color={'#fff'} className={'lg:hidden'} />
              )}
            </div>
          </button>
        </span>

        {isOpen && (
          <>
            <div className="absolute right-0 z-50 mt-2 w-56 origin-top-left rounded-md shadow-lg">
              <div className="shadow-xs rounded-md bg-white/70 backdrop-blur">
                <div
                  className="py-1"
                  role="menu"
                  aria-orientation="vertical"
                  aria-labelledby="options-menu"
                >
                  {props.children}
                </div>
              </div>
            </div>
          </>
        )}
      </div>
    </>
  )
}

export default Dropdown
