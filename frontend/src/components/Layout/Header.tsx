import clsx from 'clsx'
import { Account, WhiteLogo, Menu } from '@components/icons'
import Link from 'next/link'
import { Dropdown } from '@components/common'

const Header = () => {
  return (
    <div
      className={clsx(
        'bg-primary-1 align-center grid h-16 w-full grid-cols-5 items-center',
      )}
    >
      <div
        className={clsx(
          'justify-left h-1/1 col-span-1 items-center sm:pl-2 lg:pl-4',
        )}
      >
        <a href={process.env.CLIENT_URL}>
          <WhiteLogo />
        </a>
      </div>
      <div className={clsx('col-span-3 mx-auto')}>
        <div className={clsx('hidden text-xl font-bold lg:flex')}>
          <a
            className={clsx('hover:text-white hover:underline')}
            href={'/rent'}
          >
            借りる
          </a>{' '}
          /{' '}
          <a className={clsx('hover:text-white hover:underline')} href={'/buy'}>
            買う
          </a>{' '}
          /{' '}
          <a
            className={clsx('hover:text-white hover:underline')}
            href={'/lend'}
          >
            貸す
          </a>{' '}
          /{' '}
          <a
            className={clsx('hover:text-white hover:underline')}
            href={'/sell'}
          >
            売る
          </a>{' '}
          /{' '}
          <a
            className={clsx('hover:text-white hover:underline')}
            href={'/offers'}
          >
            オファー一覧
          </a>
        </div>
      </div>
      <div
        className={clsx(
          'col-span-1 flex h-full w-full items-center justify-end pr-4',
        )}
      >
        <Account color={'#fff'} className={'sm:hidden lg:flex'} />
        <Dropdown>
          <div
            className={clsx(
              'text-md w-1/1 mt-1 border-b-2 border-gray-300 pl-2 pb-1',
            )}
          >
            <a href={'/rent'}>借りる</a>
          </div>
          <div
            className={clsx(
              'text-md w-1/1 mt-1 border-b-2 border-gray-300 pl-2 pb-1',
            )}
          >
            <a href={'/buy'}>買う</a>
          </div>
          <div
            className={clsx(
              'text-md w-1/1 mt-1 border-b-2 border-gray-300 pl-2 pb-1',
            )}
          >
            <a href={'/lend'}>貸す</a>
          </div>
          <div
            className={clsx(
              'text-md w-1/1  my-1 border-b-2 border-gray-300 pb-1 pl-2 ',
            )}
          >
            <a href={'/sell'}>売る</a>
          </div>
          <div className={clsx('text-md w-1/1 my-1 pl-2 ')}>
            <a href={'/offers'}>オファー一覧</a>
          </div>
        </Dropdown>
      </div>
    </div>
  )
}

export default Header
