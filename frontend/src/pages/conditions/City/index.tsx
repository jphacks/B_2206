import { AnyTxtRecord } from 'dns'
import Head from 'next/head'
import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { useRecoilState } from 'recoil'
import { prefectures } from '../Prefecuture/prefectures'
import { getWithSet } from '@api/api_methods'
import { Card, PrimaryButton, ProgressBar } from '@components/common'
import { userState } from '@components/store/Auth/auth'
import { conditionState } from '@components/store/Condition/condition'

interface Props {
  nextModalName: string
  setModalName: Function
  prevModalName: string
}

interface City {
  id: string
  name: string
  checked: boolean
}

interface CheckCity {
  city: City
  checked: boolean
}

function City(props: Props): JSX.Element {
  const [condition, setCondition] = useRecoilState(conditionState)
  const [cities, setCities] = useState<City[]>()
  const [search, setSearch] = useState('')
  const [searchString, setSearchString] = useState('')
  const [isSearch, setisSearch] = useState(false)
  const [select, setSelect] = useState(true)
  const router = useRouter()

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

  const [filterDistrict, setFilterDistrict] = useState<City[]>([])
  const [filterCity, setFilterCity] = useState<City[]>([])
  const [filterTown, setFilterTown] = useState<City[]>([])
  const [filterVillage, setFilterVillage] = useState<City[]>([])

  useEffect(() => {
    let districtTmp = [] as City[]
    let cityTmp = [] as City[]
    let townTmp = [] as City[]
    let villageTmp = [] as City[]

    cities?.map((city) => {
      if (city.name.includes('区')) {
        districtTmp.push(city)
      } else if (city.name.includes('市')) {
        cityTmp.push(city)
      } else if (city.name.includes('町')) {
        townTmp.push(city)
      } else if (city.name.includes('村')) {
        villageTmp.push(city)
      }
    })
    setFilterDistrict(districtTmp)
    setFilterCity(cityTmp)
    setFilterTown(townTmp)
    setFilterVillage(villageTmp)
  }, [cities])

  const allCityCheckHandler = (bool: boolean) => {
    setFilterCity((cities) =>
      cities.map((city) => {
        return { id: city.id, name: city.name, checked: bool }
      }),
    )
  }

  const allDistrictCheckHandler = (bool: boolean) => {
    setFilterDistrict((cities) =>
      cities.map((city) => {
        return { id: city.id, name: city.name, checked: bool }
      }),
    )
  }

  const allTownCheckHandler = (bool: boolean) => {
    setFilterTown((cities) =>
      cities.map((city) => {
        return { id: city.id, name: city.name, checked: bool }
      }),
    )
  }

  const allVillageCheckHandler = (bool: boolean) => {
    setFilterVillage((cities) =>
      cities.map((city) => {
        return { id: city.id, name: city.name, checked: bool }
      }),
    )
  }

  useEffect(() => {}, [filterCity, filterDistrict, filterTown, filterVillage])

  function ShowDistrict() {
    if (filterDistrict?.length != 0) {
      return (
        <div>
          <div className="my-5 flex flex-row items-center gap-5">
            <p className={'my-3 text-lg'}>
              {cities?.at(0)?.name == '千代田区'
                ? '東京都'
                : cities?.at(0)?.name}
            </p>
            <div className="flex h-3/5 flex-row gap-5 text-xs">
              <PrimaryButton onClick={() => allDistrictCheckHandler(true)}>
                全てチェック
              </PrimaryButton>
              <PrimaryButton onClick={() => allDistrictCheckHandler(false)}>
                全てはずす
              </PrimaryButton>
            </div>
          </div>

          <div className={'flex flex-row flex-wrap gap-3'}>
            {filterDistrict?.map((district) => {
              if (searchString == '' || district.name.includes(searchString)) {
                let conditionDistrict = condition.cityNames?.includes(
                  district.name,
                )
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={() => {
                        setFilterDistrict((filterDistricts) =>
                          filterDistricts.map((setCity) => {
                            if (setCity.id == district.id) {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: !setCity.checked,
                              }
                            } else {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: setCity.checked,
                              }
                            }
                          }),
                        )
                      }}
                      checked={district.checked}
                      defaultChecked={conditionDistrict}
                      type="checkbox"
                      name="city"
                      id={`city_${district.id}`}
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label htmlFor={`city_${district.id}`}>
                      {district.name}
                    </label>
                  </div>
                )
              }
            })}
          </div>
        </div>
      )
    } else {
      return <></>
    }
  }

  function ShowCity() {
    return (
      <div>
        <div className="my-5 flex flex-row items-center gap-5">
          <p className={'my-3 text-lg'}>市一覧</p>
          <div className="flex h-3/5 flex-row gap-5 text-xs">
            <PrimaryButton onClick={() => allCityCheckHandler(true)}>
              全てチェック
            </PrimaryButton>
            <PrimaryButton onClick={() => allCityCheckHandler(false)}>
              全てはずす
            </PrimaryButton>
          </div>
        </div>
        <div className={'flex flex-row flex-wrap gap-3'}>
          {filterCity?.map((city, index) => {
            if (
              (!cities?.at(1)?.name.includes('区') && index == 0) ||
              index > 0
            ) {
              if (searchString == '' || city.name.includes(searchString)) {
                let conditionCity = condition.cityNames?.includes(city.name)
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={() => {
                        setFilterCity((filterCities) =>
                          filterCities.map((setCity) => {
                            if (setCity.id == city.id) {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: !setCity.checked,
                              }
                            } else {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: setCity.checked,
                              }
                            }
                          }),
                        )
                      }}
                      checked={city.checked}
                      defaultChecked={conditionCity}
                      type="checkbox"
                      name="city"
                      id={`city_${city.id}`}
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label htmlFor={`city_${city.id}`}>{city.name}</label>
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
        <div className="my-5 flex flex-row items-center gap-5">
          <p className={'my-3 text-lg'}>町一覧</p>
          <div className="flex h-3/5 flex-row gap-5 text-xs">
            <PrimaryButton onClick={() => allTownCheckHandler(true)}>
              全てチェック
            </PrimaryButton>
            <PrimaryButton onClick={() => allTownCheckHandler(false)}>
              全てはずす
            </PrimaryButton>
          </div>
        </div>
        <div className={'flex flex-row flex-wrap gap-3'}>
          {filterTown?.map((town) => {
            if (searchString == '' || town.name.includes(searchString)) {
              let conditionTown = condition.cityNames?.includes(town.name)
              return (
                <div className={'mb-4 flex items-center gap-1'}>
                  <input
                    onChange={() => {
                      setFilterTown((filterTowns) =>
                        filterTowns.map((setCity) => {
                          if (setCity.id == town.id) {
                            return {
                              id: setCity.id,
                              name: setCity.name,
                              checked: !setCity.checked,
                            }
                          } else {
                            return {
                              id: setCity.id,
                              name: setCity.name,
                              checked: setCity.checked,
                            }
                          }
                        }),
                      )
                    }}
                    checked={town.checked}
                    defaultChecked={conditionTown}
                    type="checkbox"
                    name="city"
                    id={`city_${town.id}`}
                    className={
                      'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                    }
                  />
                  <label htmlFor={`city_${town.id}`}>{town.name}</label>
                </div>
              )
            }
          })}
        </div>
      </div>
    )
  }

  function ShowVillage() {
    if (filterVillage?.length != 0) {
      return (
        <div>
          <div className="my-5 flex flex-row items-center gap-5">
            <p className={'my-3 text-lg'}>村一覧</p>
            <div className="flex h-3/5 flex-row gap-5 text-xs">
              <PrimaryButton onClick={() => allVillageCheckHandler(true)}>
                全てチェック
              </PrimaryButton>
              <PrimaryButton onClick={() => allVillageCheckHandler(false)}>
                全てはずす
              </PrimaryButton>
            </div>
          </div>
          <div className={'flex flex-row flex-wrap gap-3'}>
            {filterVillage?.map((village) => {
              if (searchString == '' || village.name.includes(searchString)) {
                let conditionVillage = condition.cityNames?.includes(
                  village.name,
                )
                return (
                  <div className={'mb-4 flex items-center gap-1'}>
                    <input
                      onChange={() => {
                        setFilterVillage((filterVillages) =>
                          filterVillages.map((setCity) => {
                            if (setCity.id == village.id) {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: !setCity.checked,
                              }
                            } else {
                              return {
                                id: setCity.id,
                                name: setCity.name,
                                checked: setCity.checked,
                              }
                            }
                          }),
                        )
                      }}
                      checked={village.checked}
                      defaultChecked={conditionVillage}
                      type="checkbox"
                      name="city"
                      id={`city_${village.id}`}
                      className={
                        'h-4 w-4 rounded border-gray-300 bg-gray-100 pr-3 text-blue-600 focus:ring-2 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-blue-600'
                      }
                    />
                    <label htmlFor={`city_${village.id}`}>{village.name}</label>
                  </div>
                )
              }
            })}
          </div>
        </div>
      )
    } else {
      return <></>
    }
  }

  const Caution = <p className={'text-accent-2'}>市区町村を選択してください</p>

  const submitCityHandler = () => {
    let selectCities: string[] = []
    filterCity?.map((city) => {
      if (city.checked) {
        selectCities.push(city.name)
      }
    })

    filterDistrict?.map((city) => {
      if (city.checked) {
        selectCities.push(city.name)
      }
    })

    filterTown.map((city) => {
      if (city.checked) {
        selectCities.push(city.name)
      }
    })

    filterVillage?.map((city) => {
      if (city.checked) {
        selectCities.push(city.name)
      }
    })

    if (selectCities.length != 0 || condition.cityNames.length != 0) {
      if (selectCities.length != 0) {
        setCondition({ ...condition, ...{ cityNames: selectCities } })
      }
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

  const pointName = ['都道府県の選択', '市区町村の選択', 'お部屋条件の選択']

  return (
    <div className="w-full">
      <ProgressBar pointName={pointName} nowPoint={2} />
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
          <p>{condition.prefectureName}-市区町村を選択</p>
        </div>
        <div>
          <p className={'border-primary-1 mb-5 border-b-2 border-dashed'}>
            市区町村にチェックを入れてください
          </p>
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
                    'focus:border-primary-2 focus:ring-prymary-2 dark:focus:border-primary-1 dark:focus:ring-primary-1 block w-full rounded-lg border border-gray-300 bg-gray-50 p-4 pl-10 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400'
                  }
                  placeholder="市区町村を選択"
                />
                <button
                  onClick={(e) => {
                    changeSearchHandler(e)
                  }}
                  type="submit"
                  className={
                    'bg-primary-1 hover:bg-primary-2 absolute right-2.5 bottom-2.5 rounded-lg px-4 py-2 text-sm font-medium text-white focus:outline-none focus:ring-4 focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800'
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
              お部屋の条件選択へ
            </PrimaryButton>
          </div>
          <div className={'flex justify-center'}>{!select && Caution}</div>
        </div>
      </Card>
    </div>
  )
}

export default City
