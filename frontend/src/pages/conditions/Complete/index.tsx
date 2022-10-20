import Head from 'next/head'
import styles from '@styles/Home.module.css'
import Link from 'next/link'
import {
  Card,
  PrimaryButton,
  Radio,
  Select,
  Label,
  LabelButton,
  Modal,
  Textarea,
} from '@components/common'
import { Close } from '@components/icons'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { useCallback, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { get, setGet } from '@api/api_methods'

interface City {
  id: string
  name: string
}

function Complete(): JSX.Element {
  const [user, setUser] = useRecoilState(userState)

  // console.log(select);
  console.log(user)
  // console.log(prefectures);

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
            <div className={'text-primary-2 flex justify-center text-xl'}>
              <p>登録が完了しました!</p>
            </div>
            <div className="inline-flex w-full items-center justify-center">
              <hr className="my-8 h-px w-9/10 border-0 bg-primary-1" />
              <span className="absolute left-1/2 -translate-x-1/2 bg-white px-3 font-medium text-gray-900 dark:bg-gray-900 dark:text-white">
                詳細
              </span>
            </div>
            <div className='w-9/10 mx-auto'>
              <div className='flex flex-col gap-3'>
                <div className='flex gap-3'>
                  <p>賃料</p>
                  <p>{user.conditions.priceRange}</p>
                </div>
                <div className='flex flex-row gap-3 mx-3'>
                  <p>賃料の条件</p>
                  {user.conditions.priceOptions.map((option: string) => {
                    return(
                      <p>{option}</p>
                    )
                  })}
                </div>
                <div className='flex flex-row gap-3'>
                  <p>間取りタイプ</p>
                  {user.conditions.roomTypes.map((roomType: string) => {
                    return(
                      <p>{roomType}</p>
                    )
                  })}
                </div>
                <div className='flex flex-row gap-3'>
                  <p>建物の種類</p>
                  {user.conditions.buildTypes.map((buildType: string) => {
                    return(
                      <p>{buildType}</p>
                    )
                  })}
                </div>
                <div className='flex flex-row gap-3'>
                  <p>駅徒歩</p>
                  <p>{user.conditions.time}</p>
                </div>
                <div className='flex flex-row gap-3'>
                  <p>占有面積</p>
                  <p>{user.conditions.areaRange}</p>
                </div>
                <div className='flex flex-row gap-3'>
                  <p>築年数</p>
                  <p>{user.conditions.buildAge}</p>
                </div>
                <div className='flex flex-row gap-3'>
                  <p>建物の種類</p>
                  {user.conditions.otherConditions.map((other: string) => {
                    return(
                      <p>{other}</p>
                    )
                  })}
                </div>
              </div>
            </div>
            <div className="inline-flex w-full items-center justify-center">
              <hr className="my-8 h-px w-9/10 border-0 bg-primary-1" />
            </div>
            <div className={'mt-10 mb-5 flex justify-center'}>
              <Link href={'/'}>
                <PrimaryButton>トップページに戻る</PrimaryButton>
              </Link>
            </div>
          </Card>
        </div>
      </main>
    </div>
  )
}

export default Complete
