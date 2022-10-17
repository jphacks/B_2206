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
  Radio,
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
            <div className="w-full">
            <Card width={'w-4/5'}>
                <div className={'flex flex-row justify-around py-5'}>
                    <Radio checked>エリアから探す</Radio>
                    <Radio>郵便番号から探す</Radio>

                        </div>
                <div className={"mb-5 border-2 border-dashed border-cyan-400"}></div>
                <div>
                    <div className={'py-5'}>
                        <p className={'pb-2 text-lg'}>北海道</p>
                        <div>
                            <Radio>北海道</Radio>
                        </div>
                    </div>
                    <div className={'py-5'}>
                        <p className={'pb-2 text-lg'}>東北</p>
                        <div className={'flex flex-row flex-wrap gap-5'}>
                            <Radio>青森</Radio>
                            <Radio>岩手</Radio>
                            <Radio>宮城</Radio>
                            <Radio>秋田</Radio>
                            <Radio>山形</Radio>
                            <Radio>福島</Radio>
                        </div>
                    </div>
                    <div className={'py-5'}>
                        <p className={'pb-2 text-lg'}>関東</p>
                        <div className={'flex flex-row flex-wrap gap-5'}>
                            <Radio>茨城</Radio>
                            <Radio>栃木</Radio>
                            <Radio>群馬</Radio>
                            <Radio>埼玉</Radio>
                            <Radio>千葉</Radio>
                            <Radio>東京</Radio>
                            <Radio>神奈川</Radio>
                        </div>
                    </div>
                </div>
            </Card>
            </div>
        </main>
      </div>
    )
  }
  
  export default Home
  