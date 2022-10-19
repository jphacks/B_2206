import clsx from 'clsx'
import { Account, Logo, WhiteLogo } from '@components/icons'
import Link from 'next/link'

function Header(): JSX.Element {
  return (
    <div
      className={clsx(
        'bg-primary-1 align-center grid h-16 w-full grid-cols-5 items-center',
      )}
    >
      <div className={clsx('justify-left h-1/1 col-span-1 items-center pl-4')}>
        <Link href="/">
          <WhiteLogo />
        </Link>
      </div>
      <div className={clsx('col-span-3 mx-auto')}>sumimatch</div>
      <div
        className={clsx(
          'col-span-1 flex h-full w-full items-center justify-end pr-4',
        )}
      >
        <Account color={'#fff'} />
      </div>
    </div>
  )
}

export default Header
