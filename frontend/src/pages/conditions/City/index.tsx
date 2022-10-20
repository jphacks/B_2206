import Head from 'next/head'
import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio } from '@components/common'
import { useRecoilState } from 'recoil'
import { userState } from '@components/store/Auth/auth'
import { prefectures } from '../Prefecuture/prefectures'
import { setGet } from '@api/api_methods'
import { useEffect, useState } from 'react'
import { useRouter } from 'next/router'
import { AnyTxtRecord } from 'dns'

interface Props {
  nextModalName: string
  setModalName: Function
  prevModalName: string
}

interface City {
  id: string
  name: string
}

function City(props: Props): JSX.Element {
  const [user, setUser] = useRecoilState(userState)
  const [cities, setCities] = useState<City[]>()
  const [search, setSearch] = useState('')
  const [searchString, setSearchString] = useState('')
  const [isSearch, setisSearch] = useState(false)
  const [select, setSelect] = useState(true)
  const [selectCities, setSelectCities] = useState<string[]>([])
  const router = useRouter()
  console.log(user)

  // const handler = () => {
  // setUser({ name: 'test user', email: 'test email' })
  // }

  useEffect(() => {
    setUser({ ...user, ...{ cityNames: [] } })
  }, [])

  const SearchString = <p className={'text-xl'}>{searchString} の検索結果</p>

  function changeSearchHandler(e: any) {
    e.preventDefault()
    if (search != '') {
      setisSearch(true)
    } else {
      setisSearch(false)
    }
    setSearchString(search)
  }

  function fixSearchStringHandler(e: any) {
    e.preventDefault()
    setSearch(e.target.value)
  }

  useEffect(() => {
    return () => {
      document.removeEventListener('click', changeSearchHandler)
    }
  }, [changeSearchHandler])

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

  const filter_district = cities?.filter((city) => {
    return city.name.includes('区')
  })

  const filter_town = cities?.filter((city) => {
    return city.name.includes('町')
  })

  const filter_city = cities?.filter((city) => {
    return city.name.includes('市')
  })

  const filter_village = cities?.filter((city) => {
    return city.name.includes('村')
  })

  function ShowDistrict() {
    if (filter_district?.length != 0) {
      return (
        <div>
          <p className={'my-3 text-lg'}>{(cities?.at(0)?.name == '千代田区')?"東京都":cities?.at(0)?.name}</p>
          <div className={'flex flex-row flex-wrap gap-3'}>
            {filter_district?.map((district) => {
              if (searchString == '' || district.name.includes(searchString)) {
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={(e) => {
                        citySelectHandler(e, district.name)
                      }}
                      checked={selectCities.includes(district.name)}
                      type="checkbox"
                      name="city"
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label>{district.name}</label>
                  </div>
                )
              }
            })}
          </div>
        </div>
      )
    } else {
      return <div></div>
    }
  }

  function ShowCity() {
    return (
      <div>
        <p className={'my-3 text-lg'}>市</p>
        <div className={'flex flex-row flex-wrap gap-3'}>
          {filter_city?.map((city, index) => {
            if (
              (!cities?.at(1)?.name.includes('区') && index == 0) ||
              index > 0
            ) {
              if (searchString == '' || city.name.includes(searchString)) {
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={(e) => {
                        citySelectHandler(e, city.name)
                      }}
                      checked={selectCities.includes(city.name)}
                      type="checkbox"
                      name="city"
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label>{city.name}</label>
                  </div>
                )
              }
            }
          })}
        </div>
      </div>
    )
  }

  function ShowTown() {
    return (
      <div>
        <p className={'my-3 text-lg'}>町</p>
        <div className={'flex flex-row flex-wrap gap-3'}>
          {filter_town?.map((town) => {
            if (searchString == '' || town.name.includes(searchString)) {
              return (
                <div className={'mb-4 flex items-center gap-1'}>
                  <input
                    onChange={(e) => {
                      citySelectHandler(e, town.name)
                    }}
                    checked={selectCities.includes(town.name)}
                    type="checkbox"
                    name="city"
                    className={
                      'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                    }
                  />
                  <label>{town.name}</label>
                </div>
              )
            }
          })}
        </div>
      </div>
    )
  }

  function ShowVillage() {
    if (filter_village?.length != 0) {
      return (
        <div>
          <p className={'my-3 text-lg'}>村</p>
          <div className={'flex flex-row flex-wrap gap-3'}>
            {filter_village?.map((village) => {
              if (searchString == '' || village.name.includes(searchString)) {
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={(e) => {
                        citySelectHandler(e, village.name)
                      }}
                      checked={selectCities.includes(village.name)}
                      type="checkbox"
                      name="city"
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label>{village.name}</label>
                  </div>
                )
              }
            })}
          </div>
        </div>
      )
    } else {
      return <div></div>
    }
  }

  const Caution = <p className={'text-accent-2'}>市区町村を選択してください</p>

  const citySelectHandler = (e: any, cityName: string) => {
    if (e.target.checked) {
      //   console.log('true')
      setSelect(true)
      setSelectCities([...selectCities, cityName])
    } else {
      //   console.log('else')
      setSelectCities(selectCities.filter((city) => city !== cityName))

      if (selectCities.length == 0) {
        setSelect(false)
      }
    }
    // console.log(selectCities)
  }

  const submitCityHandler = () => {
    console.log(user.cityNames)
    if (selectCities.length != 0) {
      setUser({ ...user, ...{ cityNames: selectCities } })
      props.setModalName(props.nextModalName)
    } else {
      setSelect(false)
    }
  }

  useEffect(() => {
    return () => {
      document.removeEventListener('click', submitCityHandler)
    }
  }, [submitCityHandler])

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
                    <span className="bg-primary-1 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6">
                      1
                    </span>
                    <span className="hidden sm:ml-2 sm:block"> 県の選択 </span>
                  </li>
                  <li className="bg-background-1 flex items-center p-2">
                    <span className="bg-accent-2 h-6 w-6 rounded-full text-center text-[10px] font-bold leading-6 text-white">
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
            <div className={'flex justify-start pb-8'}>
              <PrimaryButton
                onClick={() => {
                  props.setModalName(props.prevModalName)
                }}
              >
                都道府県の選択に戻る
              </PrimaryButton>
            </div>
            <div className={'text-primary-2 my-3 text-xl font-bold'}>
              <p>{user.prefectureName}-市区町村を選択</p>
            </div>
            <div>
              <p>市区町村にチェックを入れてください</p>
              <div
                className={'border-primary-1 mb-5 border-2 border-dashed'}
              ></div>
              {/* <button onClick={() => {console.log(cities?.at(0)?.name)}} >test</button> */}
              <div>
                <form
                  onSubmit={(e) => {
                    changeSearchHandler(e)
                  }}
                  className={'ml-auto w-3/5'}
                >
                  <label
                    className={
                      'sr-only mb-2 text-sm font-medium text-gray-900 dark:text-gray-300'
                    }
                  >
                    Search
                  </label>
                  <div className={'relative'}>
                    <div
                      className={
                        'pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3'
                      }
                    >
                      <svg
                        aria-hidden="true"
                        className={'h-5 w-5 text-gray-500 dark:text-gray-400'}
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                      >
                        <path
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          stroke-width="2"
                          d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                        ></path>
                      </svg>
                    </div>
                    <input
                      value={search}
                      onChange={(e) => {
                        fixSearchStringHandler(e)
                      }}
                      onSubmit={(e) => {
                        fixSearchStringHandler(e)
                      }}
                      type="search"
                      id="default-search"
                      className={
                        'block w-full rounded-lg border border-gray-300 bg-gray-50 p-4 pl-10 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500'
                      }
                      placeholder="市区町村を選択"
                    />
                    <button
                      onClick={(e) => {
                        changeSearchHandler(e)
                      }}
                      type="submit"
                      className={
                        'absolute right-2.5 bottom-2.5 rounded-lg bg-blue-700 px-4 py-2 text-sm font-medium text-white hover:bg-blue-800 focus:outline-none focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800'
                      }
                    >
                      Search
                    </button>
                  </div>
                </form>
              </div>
              <div>{isSearch && SearchString}</div>
              <div>
                <ShowDistrict />
                <ShowCity />
                <ShowTown />
                <ShowVillage />
              </div>
              <div className={'mt-8 mb-5 flex justify-center'}>
                <PrimaryButton
                  onClick={() => {
                    submitCityHandler()
                  }}
                >
                  お部屋の条件へ
                </PrimaryButton>
              </div>
              <div className={'flex justify-center'}>{!select && Caution}</div>
            </div>
          </Card>
        </div>
      </main>
    </div>
  )
}

export default City
