import Head from 'next/head'
import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio, Select } from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { useCallback, useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import {get, setGet} from '@api/api_methods'

interface Props {
  prevModalName: string
  nextModalName: string
  setModalName: Function
}

interface City {
  id: string
  name: string
}

export function SelectOptions(options: string[]): JSX.Element {
    options.map((option) => {
        return(
            <option>{option}</option>
        )
    })
}

function Other(props: Props): JSX.Element {
  const [user, setUser] = useRecoilState(userState);
  const [cities, setCities] = useState<City[]>()
  const router = useRouter();
  
  // console.log(select);
  console.log(user);
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
        "1万","1.5万","2万","2.5万","3万","3.5万","4万","4.5万","5万","5.5万","6万","6.5万","7万","7.5万","8万","8.5万","9万","9.5万","10万",
        "11万", "12万", "13万", "14万", "15万", "16万", "17万", "18万", "19万", "20万", 
        "25万", "30万", "35万", "40万", "45万", "50万", "100万" 
    ]

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
            <p className={'text-xl text-primary-2'}>エリア</p>
            <div className={'flex flex-row gap-3 text-lg mb-10'}>
              <p className={'text-lg'}>{user.prefectureName}</p>
              <p>＞</p>
              {user.cityNames.map((city: any, index: any) => {
                if(index == 0 && city.includes("区")){
                  return(
                    <div className={'flex flex-row gap-3'}>
                      <p>{cities?.at(0)?.name}</p>
                      <p>{city}</p>
                    </div>
                  )
                }else{
                  return(
                    <p>{city}</p>
                  )
                }
              })}
            </div>
            <div>
              <p className={'text-xl text-primary-2'}>絞り込み条件</p>
              <div>
                <p className={'text-lg my-3'}>賃料</p>
                <div className={'flex flex-row items-center gap-5 w-3/5 my-3'}>
                  <Select placeholder={'下限'}>
                    <option>下限無し</option>
                    {SelectOptions(priceOptions)}
                  </Select>
                  <p className={'text-xl'}>~</p>
                  <Select placeholder={'下限'}></Select>
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
