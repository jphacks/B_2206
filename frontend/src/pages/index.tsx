import type { NextPage } from 'next'
import Link from 'next/link'
import Head from 'next/head'
import Image from 'next/image'
import styles from '@styles/Home.module.css'
import {
  Card,
  FillCard,
  Input,
  Select,
  Textarea,
  PrimaryButton,
} from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import clsx from 'clsx'

const Home: NextPage = () => {
  const [user, setUser] = useRecoilState(userState)
  // console.log(user)

  // const handler = () => {
  // setUser({ name: 'test user', email: 'test email' })
  // }

  return (
    <div className={styles.container}>
      <Head>
        <title>SumiMatch</title>
        <meta name="description" content="SumiMatch" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className="grid w-full grid-cols-2">
          <FillCard width={'w-4/5'}>
            <Link href="/rent">
              <div className="py-20 text-center">
                <p className="text-center text-8xl text-fuchsia-500">借りる</p>
                <p className="text-center text-6xl text-rose-500">Rent</p>
              </div>
            </Link>
          </FillCard>

          <FillCard width={'w-4/5'}>
            <Link href="/lend">
              <div className="py-20 text-center">
                <p className="text-center text-8xl ">貸す</p>
                <p className="text-center text-6xl ">Lend</p>
              </div>
            </Link>
          </FillCard>
          <FillCard width={'w-4/5'}>
            <Link href="/buy">
              <div className="py-20 text-center">
                <p className="text-center text-8xl ">買う</p>
                <p className="text-center text-6xl ">Buy</p>
              </div>
            </Link>
          </FillCard>
          <FillCard width={'w-4/5'}>
            <Link href="/sell">
              <div className="py-20 text-center">
                <p className="text-center text-8xl ">売る</p>{' '}
                <p className="text-center text-6xl ">Sell</p>
              </div>
            </Link>
          </FillCard>
        </div>
      </main>
    </div>
  )
}

export default Home
