import Head from 'next/head'
import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio } from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { prefectures } from './prefectures'
import { useCallback, useEffect, useState } from 'react'

interface Props {
  nextModalName: string
  setModalName: Function
}

function Prefecture(props: Props): JSX.Element {
  const [user, setUser] = useRecoilState(userState)
  const [select, setSelect] = useState(true)
  // console.log(select);
  console.log(user)
  // console.log(prefectures);

  useEffect(() => {
    setUser({ ...user, ...{ prefectureId: '', prefectureName: '' } })
  }, [])

  const prefectureChangeHandler = (
    prefectureName: string,
    prefectureId: string,
  ) => {
    setSelect(true)
    setUser({
      ...user,
      ...{ prefectureName: prefectureName, prefectureId: prefectureId },
    })
  }

  const submitPrefectureHandler = () => {
    if (user.prefectureId != '') {
      props.setModalName(props.nextModalName)
    } else {
      setSelect(false)
    }
  }

  useEffect(() => {
    return () => {
      document.removeEventListener('click', submitPrefectureHandler)
    }
  }, [submitPrefectureHandler])

  const Caution = <p className={'text-accent-2'}>都道府県を選択してください</p>

  return (
    <div className={styles.container}>
      <Head>
        <title>SumiMatch</title>
        <meta name="description" content="SumiMatch" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <div className="w-full">
          <div className="flex justify-center">
            <div className="w-2/3">
              <h2 className="sr-only">Steps</h2>
              <div className="after:bg-primary-1 relative after:absolute after:inset-x-0 after:top-1/2 after:block after:h-0.5 after:-translate-y-1/2 after:rounded-lg">
                <ol className="text-main-black relative z-10 flex justify-between text-sm font-medium">
                  <li className="bg-background-1 flex items-center p-2">
                    <span className="bg-accent-2 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6 text-white">
                      1
                    </span>
                    <span className="hidden sm:ml-2 sm:block"> 県の選択 </span>
                  </li>
                  <li className="bg-background-1 flex items-center p-2">
                    <span className="bg-primary-1 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6">
                      2
                    </span>
                    <span className="hidden sm:ml-2 sm:block">
                      {' '}
                      市区町村の選択{' '}
                    </span>
                  </li>
                  <li className="bg-background-1 flex items-center p-2">
                    <span className="bg-primary-1 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6">
                      3
                    </span>
                    <span className="hidden sm:ml-2 sm:block">
                      {' '}
                      お部屋条件の選択{' '}
                    </span>
                  </li>
                </ol>
              </div>
            </div>
          </div>
          <Card width={'w-4/5'}>
            <div className={'flex flex-row justify-around py-5'}>
              <Radio checked={true} name={'slectSetting'}>
                エリアから探す
              </Radio>
              <Radio name={'slectSetting'}>郵便番号から探す</Radio>
            </div>
            <div
              className={'border-primary-1 mb-5 border-2 border-dashed'}
            ></div>
            <div>
              {prefectures.areas.map((area) => {
                return (
                  <div>
                    <div className={'py-5'}>
                      <p className={'pb-2 text-lg'}>{area.name}</p>
                      <div className={'flex flex-row flex-wrap gap-5'}>
                        {area.prefectures.map((prefecture) => {
                          return (
                            <Radio
                              name={'prefecture'}
                              onChange={() => {
                                prefectureChangeHandler(
                                  prefecture.name,
                                  prefecture.id,
                                )
                              }}
                            >
                              {prefecture.name}
                            </Radio>
                          )
                        })}
                      </div>
                    </div>
                    <div
                      className={'border-primary-1 border-t border-solid'}
                    ></div>
                  </div>
                )
              })}
            </div>
            <div className={'mt-8 mb-5 flex justify-center'}>
              <PrimaryButton
                onClick={() => {
                  submitPrefectureHandler()
                }}
              >
                市区町村を選択
              </PrimaryButton>
            </div>
            <div className={'flex justify-center'}>{!select && Caution}</div>
          </Card>
        </div>
      </main>
    </div>
  )
}

export default Prefecture
