import Head from 'next/head'
import styles from '@styles/Home.module.css'
import {
  Card,
  PrimaryButton,
  Radio,
} from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { prefectures } from '../Prefecuture/prefectures'
import { setGet } from '@api/api_methods'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router';

interface Props {
    nextModalName?: string,
    nextSetModalName?: Function,
    prevModalName: string,
    prevSetModalName: Function,
}

interface City{
    id: string,
    name: string,
}

function City(props: Props): JSX.Element {
    const [user, setUser] = useRecoilState(userState);
    const [cities, setCities] = useState<City[]>();
    const router = useRouter();
    console.log(user);
  
    // const handler = () => {
    // setUser({ name: 'test user', email: 'test email' })
    // }

    useEffect(() => {
        if(router.isReady){
            const getCitiesUrl = 'https://www.land.mlit.go.jp/webland/api/CitySearch?area=' + user.prefectureId;
            console.log(getCitiesUrl);
            const getCities = async (url: string) => {
                await setGet(url, setCities);
            };
            getCities(getCitiesUrl);
        }
    }, [])


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
                            <PrimaryButton onClick={() => {props.prevSetModalName(props.prevModalName)}}>都道府県の選択に戻る</PrimaryButton>
                        </div>
                        <div className={'text-xl font-bold text-primary-2 my-3'}>
                            <p>{user.prefectureName}-市区郡を選択</p>
                        </div>
                        <div>
                            <p>市区郡にチェックを入れてください</p>
                            <div className={"mb-5 border-2 border-dashed border-primary-1"}></div>
                                {/* <button onClick={() => {console.log(cities)}} >test</button> */}
                            <div>
                                {Object.values(cities).map((city) => {
                                    return(
                                        <p>{city.name}</p>
                                    )
                                })}
                            </div>
                        </div>
                    </Card>
                </div>
            </main>
        </div>
    )
  }
  
  export default City
  