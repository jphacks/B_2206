import Head from 'next/head'
import styles from '@styles/Home.module.css'
import {
  Card,
  PrimaryButton,
  Radio,
  Select,
  Label,
  LabelButton,
} from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { useCallback, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { get, setGet } from '@api/api_methods'

interface Props {
  prevModalName: string
  nextModalName: string
  setModalName: Function
}

interface City {
  id: string
  name: string
}

function Other(props: Props): JSX.Element {
  const [user, setUser] = useRecoilState(userState)
  const [cities, setCities] = useState<City[]>()
  const router = useRouter()

  // console.log(select);
  console.log(user)
  // console.log(prefectures);

  useEffect(() => {
    if (router.isReady) {
      const getCitiesUrl =
        'https://www.land.mlit.go.jp/webland/api/CitySearch?area=' +
        user.prefectureId
      console.log(getCitiesUrl)
      const getCities = async (url: string) => {
        await setGet(url, setCities)
      }
      getCities(getCitiesUrl)
    }
  }, [])

  const priceOptions = [
    '1万',
    '1.5万',
    '2万',
    '2.5万',
    '3万',
    '3.5万',
    '4万',
    '4.5万',
    '5万',
    '5.5万',
    '6万',
    '6.5万',
    '7万',
    '7.5万',
    '8万',
    '8.5万',
    '9万',
    '9.5万',
    '10万',
    '11万',
    '12万',
    '13万',
    '14万',
    '15万',
    '16万',
    '17万',
    '18万',
    '19万',
    '20万',
    '25万',
    '30万',
    '35万',
    '40万',
    '45万',
    '50万',
    '100万',
  ]

  const roomTypes = [
    'ワンルーム',
    '1K',
    '1DK',
    '1LDK',
    '2K',
    '2DK',
    '2LDK',
    '3K',
    '3DK',
    '3LDK',
    '4K',
    '4DK',
    '4LDK',
    '5K以上',
  ]

  const times = [
    '指定しない',
    '1分以内',
    '3分以内',
    '5分以内',
    '7分以内',
    '10分以内',
    '15分以内',
    '20分以内',
  ]

  const roomArea = [
    '20m2',
    '25m2',
    '30m2',
    '35m2',
    '40m2',
    '45m2',
    '50m2',
    '55m2',
    '60m2',
    '65m2',
    '70m2',
    '80m2',
    '90m2',
    '100m2',
  ]

  const roomAge = [
    '指定しない',
    '新築',
    '1年以内',
    '3年以内',
    '5年以内',
    '7年以内',
    '10年以内',
    '15年以内',
    '20年以内',
    '25年以内',
    '30年以内',
  ]

  const buildTypes = ['マンション', 'アパート', '一戸建て・その他']

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
            <div className={'flex justify-start pb-8'}>
              <PrimaryButton
                onClick={() => {
                  props.setModalName(props.prevModalName)
                }}
              >
                市区町村の選択に戻る
              </PrimaryButton>
            </div>
            <p className={'text-primary-2 text-xl'}>エリア</p>
            <div className={'mb-10 flex flex-row gap-3 text-lg'}>
              <p className={'text-lg'}>{user.prefectureName}</p>
              <p>＞</p>
              {user.cityNames.map((city: any, index: any) => {
                if (index == 0 && city.includes('区')) {
                  return (
                    <div className={'flex flex-row gap-3'}>
                      <p>{cities?.at(0)?.name}</p>
                      <p>{city}</p>
                    </div>
                  )
                } else {
                  return <p>{city}</p>
                }
              })}
            </div>
            <div>
              <p className={'text-primary-2 text-xl'}>絞り込み条件</p>
              <div className={'flex flex-col gap-5'}>
                <div className="w-2/3">
                  <p className={'my-3 text-lg'}>賃料</p>
                  <div className={'my-3 flex flex-row items-center gap-5'}>
                    <Select>
                      <option>下限無し</option>
                      {priceOptions.map((price) => {
                        return <option>{price}</option>
                      })}
                    </Select>
                    <p className={'text-xl'}>~</p>
                    <Select>
                      <option>上限無し</option>
                      {priceOptions.map((price) => {
                        return <option>{price}</option>
                      })}
                    </Select>
                  </div>
                  <div
                    className={'w-1/1 flex flex-row justify-center gap-5 pt-5'}
                  >
                    <LabelButton name={'管理費・共益費込み'} />
                    <LabelButton name={'礼金なし'} />
                    <LabelButton name={'敷金・保証金なし'} />
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>間取りタイプ</p>
                  <div className={'flex w-2/3 flex-row flex-wrap gap-5'}>
                    {roomTypes.map((roomType) => {
                      return <LabelButton name={roomType} />
                    })}
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>建物の種類</p>
                  <div className={'flex w-2/3 flex-row flex-wrap gap-5'}>
                    {buildTypes.map((buildType) => {
                      return <LabelButton name={buildType} />
                    })}
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>駅徒歩</p>
                  <div className={'w-1/3'}>
                    <Select>
                      {times.map((time) => {
                        return <option>{time}</option>
                      })}
                    </Select>
                  </div>
                </div>
                <div className={'w-2/3'}>
                  <p className={'my-3 text-lg'}>占有面積</p>
                  <div className={'my-3 flex flex-row items-center gap-5'}>
                    <Select>
                      <option>下限無し</option>
                      {roomArea.map((area) => {
                        return <option>{area}</option>
                      })}
                    </Select>
                    <p className={'text-xl'}>~</p>
                    <Select>
                      <option>上限無し</option>
                      {roomArea.map((area) => {
                        return <option>{area}</option>
                      })}
                    </Select>
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>築年数</p>
                  <div className={'w-1/3'}>
                    <Select>
                      {roomAge.map((age) => {
                        return <option>{age}</option>
                      })}
                    </Select>
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>こだわり条件</p>
                </div>
              </div>
            </div>
          </Card>
        </div>
      </main>
    </div>
  )
}

export default Other
