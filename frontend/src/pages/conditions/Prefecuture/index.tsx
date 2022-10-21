import styles from '@styles/Home.module.css'
import { Card, PrimaryButton, Radio, ProgressBar } from '@components/common'
import { useRecoilState } from 'recoil'
import { conditionState } from '@components/store/Condition/condition'
import { prefectures } from './prefectures'
import { useEffect, useState } from 'react'
import { userAgent } from 'next/server'

interface Props {
  nextModalName: string
  setModalName: Function
}

const Modals = {
  search: 'search',
  postCode: 'postcode',
}

function Prefecture(props: Props): JSX.Element {
  const [condition, setCondition] = useRecoilState(conditionState)
  const [modalName, setModalName] = useState(Modals.search)
  const [select, setSelect] = useState(true)

  useEffect(() => {
    setCondition({ ...condition, ...{ cityNames: [] } })
  }, [])

  const prefectureChangeHandler = (
    prefectureName: string,
    prefectureId: string,
  ) => {
    setSelect(true)
    setCondition({
      ...condition,
      ...{ prefectureName: prefectureName, prefectureId: prefectureId },
    })
  }

  const submitPrefectureHandler = () => {
    if (condition.prefectureId != '') {
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

  const pointName = ['県の選択', '市区町村の選択', 'お部屋条件の選択']
  const pointNamePostCode = ['県と市区町村の選択', 'お部屋条件の選択']

  const changeModalHandler = (mode: string) => {
    if (mode == 'search') {
      setModalName(Modals.search)
    } else {
      setModalName(Modals.postCode)
    }
  }

  useEffect(() => {}, [changeModalHandler])

  const SearchModal = (
    <>
      <div>
        {prefectures.areas.map((area) => {
          return (
            <div>
              <div className={'py-5'}>
                <p className={'pb-2 text-lg'}>{area.name}</p>
                <div className={'flex flex-row flex-wrap gap-5'}>
                  {area.prefectures.map((prefecture) => {
                    let conditionPrefecture = false
                    if (condition.prefectureId == prefecture.id) {
                      conditionPrefecture = true
                    }
                    return (
                      <Radio
                        name={'prefecture'}
                        onChange={() => {
                          prefectureChangeHandler(
                            prefecture.name,
                            prefecture.id,
                          )
                        }}
                        checked={conditionPrefecture}
                        id={`prefecture_${prefecture.id}`}
                      >
                        {prefecture.name}
                      </Radio>
                    )
                  })}
                </div>
              </div>
              <div className={'border-primary-1 border-t border-solid'}></div>
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
    </>
  )

  const [search, setSearch] = useState('')
  const [searchString, setSearchString] = useState('')
  const [isSearch, setisSearch] = useState(false)

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

  useEffect(() => {}, [searchString])

  const PostCodeModal = (
    <>
      <div className="flex flex-col gap-5">
        <p className="text-primary-2 text-xl">郵便番号</p>
        <form
          onSubmit={(e) => {
            changeSearchHandler(e)
          }}
          className={'mr-auto w-3/5'}
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
              placeholder="例:123-4567"
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
      <div></div>

      <div className={'mt-8 mb-5 flex justify-center'}>
        <PrimaryButton
          onClick={() => {
            submitPrefectureHandler()
          }}
        >
          お部屋の条件選択へ
        </PrimaryButton>
      </div>
      <div className={'flex justify-center'}>{!select && Caution}</div>
    </>
  )

  return (
    <div className={styles.container}>
      <main className={styles.main}>
        <div className="w-full">
          {modalName === Modals.search && (
            <ProgressBar pointName={pointName} nowPoint={1} />
          )}
          {modalName === Modals.postCode && (
            <ProgressBar pointName={pointNamePostCode} nowPoint={1} />
          )}

          <Card width={'w-4/5'}>
            <div className={'flex flex-row justify-around py-5'}>
              <Radio
                name={'slectSetting'}
                onClick={() => changeModalHandler('search')}
              >
                エリアから探す
              </Radio>
              <Radio
                name={'slectSetting'}
                onClick={() => changeModalHandler('postCode')}
              >
                郵便番号から探す
              </Radio>
            </div>
            <div
              className={'border-primary-1 mb-5 border-2 border-dashed'}
            ></div>
            <div>
              {modalName === Modals.search && SearchModal}
              {modalName === Modals.postCode && PostCodeModal}
            </div>
          </Card>
        </div>
      </main>
    </div>
  )
}

export default Prefecture
