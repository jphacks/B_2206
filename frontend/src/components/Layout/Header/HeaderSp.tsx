import clsx from 'clsx'
import {
  Apartment,
  Exchange,
  FormatList,
  House,
  Swap,
  WhiteLogo,
} from '@components/icons'
import { Dropdown } from '@components/common'

const Header = () => {
  return (
    <>
      <div className={clsx('justify-left h-1/1 col-span-1 sm:pl-2')}>
        <a href={process.env.CLIENT_URL}>
          <WhiteLogo />
        </a>
      </div>
      <div className={clsx('col-span-3')}></div>
      <div
        className={clsx(
          'col-span-1 flex h-full w-full items-center justify-end pr-4',
        )}
      >
        <Dropdown>
          <a
            className={clsx(
              'text-md w-1/1 mt-1 grid grid-cols-4 border-b-2 border-gray-300 pl-2 pb-1',
            )}
            href={'/rent'}
          >
            <div className="col-span-1 flex justify-center">
              <Apartment />
              <Swap />
            </div>
            <span className="border-primary-1 col-span-2 ml-1 border-l-2 border-dashed pl-1">
              借りる
            </span>
          </a>
          <a
            className={clsx(
              'text-md w-1/1 mt-1 grid grid-cols-4 border-b-2 border-gray-300 pl-2 pb-1',
            )}
            href={'/buy'}
          >
            <div className="col-span-1 flex justify-center">
              <House />
              <Swap />
            </div>
            <span className="border-primary-1 col-span-2 ml-1 border-l-2 border-dashed pl-1">
              買う
            </span>
          </a>
          <a
            className={clsx(
              'text-md w-1/1 mt-1 grid grid-cols-4 border-b-2 border-gray-300 pl-2 pb-1',
            )}
            href={'/lend'}
          >
            <div className="col-span-1 flex justify-center">
              <Apartment />
              <Exchange />
            </div>
            <span className="border-primary-1 col-span-2 ml-1 border-l-2 border-dashed pl-1">
              貸す
            </span>
          </a>
          <a
            className={clsx(
              'text-md w-1/1  my-1 grid grid-cols-4 border-b-2 border-gray-300 pb-1  pl-2',
            )}
            href={'/sell'}
          >
            <div className="col-span-1 flex  justify-center">
              <House />
              <Exchange />
            </div>
            <span className="border-primary-1 col-span-2 ml-1 border-l-2 border-dashed pl-1">
              売る
            </span>
          </a>
          <a
            className={clsx('text-md w-1/1 my-1 grid grid-cols-4 pl-2')}
            href={'/offers'}
          >
            <div className="col-span-1 flex justify-center">
              <FormatList />
            </div>
            <span className="border-primary-1 col-span-2 ml-1 border-l-2 border-dashed pl-1">
              オファー一覧
            </span>
          </a>
        </Dropdown>
      </div>
    </>
  )
}

export default Header
