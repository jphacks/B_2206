import Head from 'next/head'
import styles from '@styles/Home.module.css'
import {
  Card,
  PrimaryButton,
  Radio,
} from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { prefectures } from './prefectures'
import { useCallback, useEffect, useState } from 'react'

interface Props {
    nextModalName: string
    nextSetModalName: Function
}

function Prefecture(props: Props): JSX.Element {
    const [user, setUser] = useRecoilState(userState);
    const [select, setSelect] = useState(true);
    // console.log(select);
    // console.log(user);
    // console.log(prefectures);

    useEffect(()=> {
        setUser({...user, ...{prefecture: ''}});
    },[])
  
    const prefectureChangeHandler = ( prefectureName: string, prefectureId: string ) => {
        setSelect(true);
        setUser({...user, ...{prefectureName: prefectureName, prefectureId: prefectureId}});
    }
    
    const submitPrefectureHandler = () => {
        if(user.prefectureId != ''){
            props.nextSetModalName(props.nextModalName);
        }else{
            setSelect(false);
        }
    }

    useEffect(() => {
        return () => {
            document.removeEventListener('click', submitPrefectureHandler);
        }
    }, [submitPrefectureHandler])

    const Caution = (
        <p className={'text-accent-2'}>都道府県を選択してください</p>
    )

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
                        <Radio checked={true} name={'slectSetting'}>エリアから探す</Radio>
                        <Radio name={'slectSetting'} >郵便番号から探す</Radio>
                    </div>
                    <div className={"mb-5 border-2 border-dashed border-primary-1"}></div>
                    <div>
                        {prefectures.areas.map((area) => {
                            return(
                                <div>
                                    <div className={'py-5'}>
                                        <p className={'pb-2 text-lg'}>{area.name}</p>
                                        <div className={'flex flex-row flex-wrap gap-5'}>
                                            {area.prefectures.map((prefecture) => {
                                                return(
                                                    <Radio name={'prefecture'} onChange={() => {prefectureChangeHandler(prefecture.name, prefecture.id)}}>{prefecture.name}</Radio>
                                                )
                                            })}
                                        </div>
                                    </div>
                                    <div className={"border-t border-solid border-primary-1"}></div>
                                </div>
                            )
                        })}
                    </div>
                    <div className={'flex justify-center mt-8 mb-5'}>
                        <PrimaryButton onClick={() => {submitPrefectureHandler()}}>市区郡を検索へ</PrimaryButton>
                    </div>
                    <div className={'flex justify-center'}>
                        {!select && Caution}
                    </div>
                </Card>
            </div>
        </main>
      </div>
    )
  }
  
  export default Prefecture
  