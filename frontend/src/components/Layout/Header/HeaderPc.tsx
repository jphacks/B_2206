import clsx from 'clsx'
import { Account, WhiteLogo } from '@components/icons'
import { Dropdown } from '@components/common'

const HeaderPc = () => {
  return (
    <>
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
      </div>
    </>
  )
}

export default HeaderPc
