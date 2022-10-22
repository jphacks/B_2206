import clsx from 'clsx'
import type { NextPage } from 'next'
import { MODERN_BROWSERSLIST_TARGET } from 'next/dist/shared/lib/constants'
import Head from 'next/head'
import Image from 'next/image'
import Link from 'next/link'
import { useEffect } from 'react'
import { useRecoilState } from 'recoil'
import {
  Card,
  FillCard,
  Input,
  Select,
  Textarea,
  PrimaryButton,
} from '@components/common'
import { userState } from '@components/store/Auth/auth'
import { conditionState } from '@components/store/Condition/condition'
import styles from '@styles/Home.module.css'

const recoilDefaultValue = {
  id: 1,
  name: 'gidai yuuki',
  email: 'gidai@email',
}

const recoilConditionDefaultValue = {
  id: 1,
  mode: '',
  prefectureId: '',
  prefectureName: '',
  cityNames: [],
  conditions: {},
  isPostCode: false,
}

const Home: NextPage = () => {
  const [user, setUser] = useRecoilState(userState)
  const [condition, setCondition] = useRecoilState(conditionState)

  // 初期化用として使いたいのでコメントアウトして残します
  useEffect(() => {
    // setUser(recoilDefaultValue)
    setCondition(recoilConditionDefaultValue)
  }, [])

  const setModeHandler = (modeString: string) => {
    setCondition({ ...condition, ...{ mode: modeString } })
  }

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
            <Link href="/conditions">
              <div
                className="py-20 text-center"
                onClick={() => setModeHandler('rent')}
              >
                <p className="text-center text-8xl ">借りる</p>
                <p className="text-center text-6xl ">Rent</p>
              </div>
            </Link>
          </FillCard>
          <FillCard width={'w-4/5'}>
            <Link href="/conditions">
              <div
                className="py-20 text-center"
                onClick={() => setModeHandler('lend')}
              >
                <p className="text-center text-8xl ">貸す</p>
                <p className="text-center text-6xl ">Lend</p>
              </div>
            </Link>
          </FillCard>
          <FillCard width={'w-4/5'}>
            <Link href="/conditions">
              <div
                className="py-20 text-center"
                onClick={() => setModeHandler('buy')}
              >
                <p className="text-center text-8xl ">買う</p>
                <p className="text-center text-6xl ">Buy</p>
              </div>
            </Link>
          </FillCard>
          <FillCard width={'w-4/5'}>
            <Link href="/conditions">
              <div
                className="py-20 text-center"
                onClick={() => setModeHandler('sell')}
              >
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
