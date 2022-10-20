import clsx from 'clsx'
import HeaderPc from './HeaderPc'
import HeaderSp from './HeaderSp'

const Header = () => {
  return (
    <div
      className={clsx(
        'bg-primary-1 align-center grid h-16 w-full grid-cols-5 items-center',
      )}
    >
      <HeaderSp />
      <HeaderPc />
    </div>
  )
}

export default Header
