import styles from '@styles/Home.module.css'
import Link from 'next/link'
import {
  Card,
  PrimaryButton,
  Select,
  LabelButton,
  Modal,
  ProgressBar,
} from '@components/common'
import { Close } from '@components/icons'
import { useRecoilState } from 'recoil'
import { conditionState } from '@components/store/Condition/condition'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { getWithSet } from '@api/api_methods'

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
  const [condition, setCondition] = useRecoilState(conditionState)
  const [cities, setCities] = useState<City[]>()
  const router = useRouter()

  useEffect(() => {
    setCondition({ ...condition, ...{ conditions: {} } })
  }, [])

  useEffect(() => {
    scrollTo(0, 0)

    if (router.isReady) {
      const getCitiesUrl =
        'https://www.land.mlit.go.jp/webland/api/CitySearch?area=' +
        condition.prefectureId
      const getCities = async (url: string) => {
        await getWithSet(url, setCities)
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

  const morePriceOption = ['管理費・共益費込み', '礼金なし', '敷金・保証金なし']

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

  const otherConditions = [
    '駐車場あり',
    'バス・トイレ別',
    'ペット相談',
    '2階以上住戸',
    '室内洗濯機付き',
    'エアコン付き',
    'オートロック',
    'フローリング',
    '間取り図付き',
    '物件動画付き',
    'パノラマ付き',
    '定期借家を含まない',
  ]

  const buildTypes = ['マンション', 'アパート', '一戸建て・その他']

  const [isOpen, setIsOpen] = useState<boolean>(false)
  const onOpen = () => setIsOpen(true)
  const onClose = () => setIsOpen(false)

  const [priceMax, setPriceMax] = useState<string>('上限なし')
  const [priceMin, setPriceMin] = useState<string>('下限なし')
  const [priceOption, setPriceOption] = useState<string[]>([])
  const [selectRoomType, setSelectRoomType] = useState<string[]>([])
  const [selectBuildType, setSelectBuildType] = useState<string[]>([])
  const [selectTime, setSelectTime] = useState<string>('指定なし')
  const [areaMax, setAreaMax] = useState<string>('上限なし')
  const [areaMin, setAreaMin] = useState<string>('下限なし')
  const [buildAge, setBuildAge] = useState<string>('指定なし')
  const [otherCondition, setOtherCondition] = useState<string[]>([])
  const [isFinish, setIsFinish] = useState<boolean>(false)

  function optionShow(options: string[]) {
    if (options.length == 0) {
      return <p>指定なし</p>
    } else {
      return options.map((option) => {
        return <p>{option}</p>
      })
    }
  }

  function selectArrayHandler(str: string, arr: string[], setArr: Function) {
    if (arr.includes(str)) {
      setArr(arr.filter((arrNode) => arrNode !== str))
    } else {
      setArr([...arr, str])
    }
  }

  useEffect(() => {
    if (isFinish) {
      if (priceOption.length == 0) setPriceOption([...priceOption, '指定なし'])
      if (selectRoomType.length == 0)
        setSelectRoomType([...selectRoomType, '指定なし'])
      if (selectBuildType.length == 0)
        setSelectBuildType([...selectBuildType, '指定なし'])
      if (otherCondition.length == 0)
        setOtherCondition([...otherCondition, '指定なし'])
    }
  }, [clickHandler])

  function clickHandler() {
    setIsFinish(true)
    const conditions = {
      priceRange: `${priceMin} ~ ${priceMax}`,
      priceOptions: priceOption,
      roomTypes: selectRoomType,
      buildTypes: selectBuildType,
      time: selectTime,
      areaRange: `${areaMin} ~ ${areaMax}`,
      buildAge: buildAge,
      otherConditions: otherCondition,
    }
    setCondition({ ...condition, ...{ conditions: conditions } })
    props.setModalName(props.nextModalName)
  }

  const pointName = ['県の選択', '市区町村の選択', 'お部屋条件の選択']

  const buttonByMode = (mode: string) => {
    if (mode == 'lend') {
      return (
        <Link href="/rentlist">
          <PrimaryButton>この条件で借り手を探す</PrimaryButton>
        </Link>
      )
    } else if (mode == 'sell') {
      return (
        <Link href="/buylist">
          <PrimaryButton>この条件で買い手を探す</PrimaryButton>
        </Link>
      )
    } else {
      return (
        <PrimaryButton
          onClick={() => {
            clickHandler()
          }}
        >
          条件を確定して登録する
        </PrimaryButton>
      )
    }
  }

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <div className="w-full">
          <ProgressBar pointName={pointName} nowPoint={3} />
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
            <div className={'mb-10 flex flex-row flex-wrap gap-3 text-lg'}>
              <p className={'text-lg'}>{condition.prefectureName}</p>
              <p>＞</p>
              {condition.cityNames.map((city: any, index: any) => {
                if (index == 0 && city.includes('区')) {
                  return (
                    <div className={'flex flex-row gap-3'}>
                      <p>{cities?.at(0)?.name}</p>
                      <p>{city}</p>
                    </div>
                  )
                } else if (index == 5) {
                  return <p>ほか {condition.cityNames.length - 5} 市区町村</p>
                } else if (index >= 6) {
                  return
                } else {
                  return <p>{city}</p>
                }
              })}
            </div>
            <div>
              <p className={'text-primary-2 text-xl'}>絞り込み条件</p>
              <div className={'flex flex-col gap-5'}>
                <div className="w-4/5">
                  <p className={'my-3 text-lg'}>賃料</p>
                  <div className={'my-3 flex flex-row items-center gap-5'}>
                    <Select
                      onChange={(e: any) => {
                        setPriceMin(e.target.value)
                      }}
                    >
                      <option>下限無し</option>
                      {priceOptions.map((price) => {
                        return <option>{price}</option>
                      })}
                    </Select>
                    <p className={'text-xl'}>~</p>
                    <Select
                      onChange={(e: any) => {
                        setPriceMax(e.target.value)
                      }}
                    >
                      <option>上限無し</option>
                      {priceOptions.map((price) => {
                        return <option>{price}</option>
                      })}
                    </Select>
                  </div>
                  <div
                    className={'w-1/1 flex flex-row justify-start gap-5 pt-5'}
                  >
                    {morePriceOption.map((option) => {
                      return (
                        <LabelButton
                          name={option}
                          onClick={() => {
                            selectArrayHandler(
                              option,
                              priceOption,
                              setPriceOption,
                            )
                          }}
                        />
                      )
                    })}
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>間取りタイプ</p>
                  <div className={'grid auto-rows-auto grid-cols-4 gap-5'}>
                    {roomTypes.map((roomType) => {
                      return (
                        <LabelButton
                          name={roomType}
                          onClick={() => {
                            selectArrayHandler(
                              roomType,
                              selectRoomType,
                              setSelectRoomType,
                            )
                          }}
                        />
                      )
                    })}
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>建物の種類</p>
                  <div className={'flex w-2/3 flex-row flex-wrap gap-5'}>
                    {buildTypes.map((buildType) => {
                      return (
                        <LabelButton
                          name={buildType}
                          onClick={() => {
                            selectArrayHandler(
                              buildType,
                              selectBuildType,
                              setSelectBuildType,
                            )
                          }}
                        />
                      )
                    })}
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>駅徒歩</p>
                  <div className={'w-1/3'}>
                    <Select
                      onChange={(e: any) => {
                        setSelectTime(e.target.value)
                      }}
                    >
                      {times.map((time) => {
                        return <option>{time}</option>
                      })}
                    </Select>
                  </div>
                </div>
                <div className={'w-2/3'}>
                  <p className={'my-3 text-lg'}>占有面積</p>
                  <div className={'my-3 flex flex-row items-center gap-5'}>
                    <Select
                      onChange={(e: any) => {
                        setAreaMin(e.target.value)
                      }}
                    >
                      <option>下限無し</option>
                      {roomArea.map((area) => {
                        return <option>{area}</option>
                      })}
                    </Select>
                    <p className={'text-xl'}>~</p>
                    <Select
                      onChange={(e: any) => {
                        setAreaMax(e.target.value)
                      }}
                    >
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
                    <Select
                      onChange={(e: any) => {
                        setBuildAge(e.target.value)
                      }}
                    >
                      {roomAge.map((age) => {
                        return <option>{age}</option>
                      })}
                    </Select>
                  </div>
                </div>
                <div>
                  <p className={'my-3 text-lg'}>こだわり条件</p>
                  <div className={'grid grid-cols-3 gap-5'}>
                    {otherConditions.map((other) => {
                      return (
                        <LabelButton
                          name={other}
                          onClick={() => {
                            selectArrayHandler(
                              other,
                              otherCondition,
                              setOtherCondition,
                            )
                          }}
                        />
                      )
                    })}
                  </div>
                </div>
              </div>
              <div className={'mt-10 mb-5 flex justify-center'}>
                <PrimaryButton onClick={onOpen}>お部屋条件の確認</PrimaryButton>
              </div>
            </div>
          </Card>
        </div>
        <div>
          {isOpen && (
            <Modal>
              <div
                className={
                  'justify-left ml-4 mb-4 w-fit cursor-pointer items-center'
                }
              >
                <Close width="32" height="32" onClick={onClose} />
              </div>
              <div className={'mx-5 flex flex-col gap-5'}>
                <div
                  className={
                    'text-accent-2 mb-10 flex flex-row flex-wrap gap-3 text-lg'
                  }
                >
                  <p className={'text-lg'}>{condition.prefectureName}</p>
                  <p>＞</p>
                  {condition.cityNames.map((city: any, index: any) => {
                    if (index == 0 && city.includes('区')) {
                      return (
                        <div className={'flex flex-row gap-3'}>
                          <p>{cities?.at(0)?.name}</p>
                          <p>{city}</p>
                        </div>
                      )
                    } else if (index == 5) {
                      return (
                        <p>ほか {condition.cityNames.length - 5} 市区町村</p>
                      )
                    } else if (index >= 6) {
                      return
                    } else {
                      return <p>{city}</p>
                    }
                  })}
                </div>
              </div>
              <div>
                <div>
                  <p className={'text-primary-2 text-xl'}>賃料</p>
                  <div className={'mx-3 flex flex-row gap-5'}>
                    <p>
                      {priceMin} ~ {priceMax}
                    </p>
                  </div>
                  <div className={'mx-3 flex flex-row gap-4'}>
                    {priceOption.map((option) => {
                      return <p>{option}</p>
                    })}
                  </div>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>間取りタイプ</p>
                  <div className={'mx-3 flex flex-row gap-4'}>
                    {optionShow(selectRoomType)}
                  </div>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>建物の種類</p>
                  <div className={'mx-3 flex flex-row gap-4'}>
                    {optionShow(selectBuildType)}
                  </div>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>駅徒歩</p>
                  <p className={'mx-3'}>{selectTime}</p>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>占有面積</p>
                  <div className={'mx-3 flex flex-row gap-4'}>
                    <p>
                      {areaMin} ~ {areaMax}
                    </p>
                  </div>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>築年数</p>
                  <p className={'mx-3'}>{buildAge}</p>
                </div>
                <div>
                  <p className={'text-primary-2 text-xl'}>こだわり条件</p>
                  <div className={'mx-3 flex flex-row gap-4'}>
                    {optionShow(otherCondition)}
                  </div>
                </div>
              </div>
              <div className={'mt-10 mb-5 flex justify-center'}>
                {buttonByMode(condition.mode)}
              </div>
            </Modal>
          )}
        </div>
      </main>
    </div>
  )
}

export default Other
