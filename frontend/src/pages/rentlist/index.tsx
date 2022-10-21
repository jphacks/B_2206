import clsx from 'clsx'
import type { NextPage } from 'next'
import { useState, useEffect, memo, useRef } from 'react'
import { useRecoilState } from 'recoil'
import { Card, Label, Modal, PrimaryButton, Textarea } from '@components/common'
import { Close } from '@components/icons'
import { allDataState } from '@components/store/data/allData'
import { Area, Tag, Limit, PersonalInfo, CompanyInfo } from '@type/common'


interface Props {
  children?: React.ReactNode
  data: any
}

const Rentlist: NextPage = () => {
  const personalInfo: PersonalInfo = {
    id: 1,
    familyName: '藤崎',
    firstName: '竜成',
    familyNameKana: 'フジサキ',
    firstNameKana: 'リュウセイ',
    birthDay: '19990202',
    phoneNumber: '09012345678',
    userId: 1,
    created_at: '',
    updated_at: '',
  }
  const companyInfo = {
    id: 1,
    name: '長岡金型',
    phoneNumber: '09012345678',
    postCode: '9456666',
    prefecture: '新潟県',
    city: '長岡市西陵町',
    addressNumber: '2674-31',
    buildingName: '長岡金型ビル',
    website: 'https://nagaoka-kanagata.co.jp',
    userId: 1,
    created_at: '',
    updated_at: '',
  }
  const user = {
    id: 1,
    name: 'test',
    email: 'test',
    password: 'test',
    personalInfo: personalInfo,
    companyInfo: null,
    created_at: '',
    updated_at: '',
  }

  const area: Area = {
    id: 1,
    postCode: 1110000,
    prefecture: '北海道',
    city: '札幌市',
    addressNumber: '1-1-1',
    buildingName: 'test',
    created_at: '',
    updated_at: '',
  }
  const [allData, setAllData] = useRecoilState(allDataState)

  // カードの中に表示するデータ
  const CardContent: React.FC<Props> = memo((props) => {
    // 表示する条件を切り替えるためのref
    const defaultContentRef = useRef<HTMLDivElement>(null)
    const subContentRef = useRef<HTMLDivElement>(null)
    const modalRef = useRef<HTMLDivElement>(null)

    useEffect(() => {
      // if文を使って、inputRefのcurrentプロパティが存在するか確認
      if (defaultContentRef.current && subContentRef.current) {
        subContentRef.current.focus()
        defaultContentRef.current.focus()
        // subContent を非表示にする
        subContentRef.current.style.opacity = '0'
        subContentRef.current.style.visibility = 'hidden'
        subContentRef.current.style.height = '0px'
        subContentRef.current.style.width = '0px'
      }
    }, [])

    // マウスが乗ったら「その他の条件」を表示
    const handleMouseEnter = () => {
      if (!defaultContentRef.current && modalRef.current) return
      else if (!defaultContentRef.current || !subContentRef.current) return
      defaultContentRef.current.style.opacity = '0'
      defaultContentRef.current.style.visibility = 'hidden'
      defaultContentRef.current.style.height = '0px'
      defaultContentRef.current.style.width = '0px'
      subContentRef.current.style.opacity = '1'
      subContentRef.current.style.visibility = 'visible'
      subContentRef.current.style.height = 'auto'
      subContentRef.current.style.width = 'auto'
    }
    // マウスが離れたら「その他の条件」を非表示
    const handleMouseLeave = () => {
      if (!defaultContentRef.current || !subContentRef.current) return
      defaultContentRef.current.style.opacity = '1'
      defaultContentRef.current.style.visibility = 'visible'
      defaultContentRef.current.style.height = 'auto'
      defaultContentRef.current.style.width = 'auto'
      subContentRef.current.style.opacity = '0'
      subContentRef.current.style.visibility = 'hidden'
      subContentRef.current.style.height = '0px'
      subContentRef.current.style.width = '0px'
    }

    const [isOpen, setIsOpen] = useState<boolean>(false)
    const onOpen = () => setIsOpen(true)
    const onClose = () => setIsOpen(false)

    // 希望する条件を表示
    const defaultContent = (
      <>
        <div>
          <span className="border-primary-1 border-l-4 pl-1 text-left text-xl font-bold">
            希望する条件
          </span>
        </div>
        <div className="grid grid-cols-2 gap-2 pt-2 text-left">
          {props.data.user.request.detail.detailLimits.map(
            (detailLimits: {
              id: number
              classificationName: string
              limits: Limit[]
            }) => (
              <div className="col-span-1">
                <span className="text-md">
                  {detailLimits.limits.map((limit: any) => (
                    <>
                      {limit.name === '賃料' && (
                        <div>
                          <span className="border-primary-1 border-b-2  pt-1 text-lg font-bold">
                            {limit.name}
                          </span>
                          <br />
                          <span>
                            {limit.minValue.name}〜{limit.maxValue.name}
                          </span>
                        </div>
                      )}
                    </>
                  ))}
                </span>
              </div>
            ),
          )}
          {props.data.user.request.detail.detailValues.map(
            (detailValues: any) => (
              <div className="col-span-1">
                <span className="text-md">
                  {detailValues.values.map((value: any) => (
                    <>
                      <span className="border-primary-1 mt-2  border-b-2 text-lg font-bold">
                        {value.name}
                      </span>
                      <br />
                      <span>{value.value?.name}</span>
                    </>
                  ))}
                </span>
              </div>
            ),
          )}
          {props.data.user.request.detail.detailTags.map((detailTags: any) => (
            <>
              {detailTags.classificationName === '間取り' && (
                <div className="col-span-1">
                  <span className="border-primary-1 border-b-2  pt-1 text-lg font-bold">
                    {detailTags.classificationName}
                  </span>
                  <br />
                  <span className="text-md">
                    {detailTags.tags.map((tag: Tag, index: number) => (
                      <>
                        <span>{tag.name}</span>
                        {index !== detailTags.tags.length - 1 && (
                          <span>, </span>
                        )}
                      </>
                    ))}
                  </span>
                </div>
              )}
              {detailTags.tags.map((tag: any) => (
                <>
                  {tag.name === '間取り' && (
                    <div>
                      <div className="pt-1 text-lg font-bold">{tag.name}</div>
                    </div>
                  )}
                </>
              ))}
              {detailTags.name === '間取り' && (
                <div className="py-2 text-lg font-bold">
                  {detailTags.name}
                  <br />
                  <div className="w-1/1 flex flex-wrap text-sm">
                    {detailTags.tags.map((tag: Tag) => (
                      <Label
                        name={tag.name}
                        width={'w-auto'}
                        className={'mx-1 mt-2 '}
                      />
                    ))}
                  </div>
                </div>
              )}
            </>
          ))}
        </div>
      </>
    )

    // その他の条件を表示
    const subContent = (
      <>
        <div>
          <span className="border-primary-2 border-l-4 pl-1 text-left text-xl font-bold">
            その他の条件
          </span>
        </div>
        <div className="grid grid-cols-2 gap-2 pt-2 text-left">
          {props.data.user.request.detail.detailLimits.map(
            (detailLimits: {
              id: number
              classificationName: string
              limits: Limit[]
            }) => (
              <div className="col-span-1">
                <span className="text-md">
                  {detailLimits.limits.map((limit: any) => (
                    <>
                      {limit.name != '賃料' && (
                        <div>
                          <span className="border-primary-2 border-b-2  pt-1 text-lg font-bold">
                            {limit.name}
                          </span>
                          <br />
                          <span>
                            {limit.minValue.name}〜{limit.maxValue.name}
                          </span>
                        </div>
                      )}
                    </>
                  ))}
                </span>
              </div>
            ),
          )}
          {props.data.user.request.detail.detailTags.map((detailTags: any) => (
            <>
              {detailTags.classificationName != '間取り' && (
                <div className="col-span-1">
                  <span className="border-primary-2 border-b-2  pt-1 text-lg font-bold">
                    {detailTags.classificationName}
                  </span>
                  <br />
                  <span className="text-md">
                    {detailTags.tags.map((tag: Tag, index: number) => (
                      <>
                        <span>{tag.name}</span>
                        {index !== detailTags.tags.length - 1 && (
                          <span>, </span>
                        )}
                      </>
                    ))}
                  </span>
                </div>
              )}
            </>
          ))}
        </div>
        <div className="flex justify-center pt-4 pb-2 text-xl font-bold">
          <PrimaryButton onClick={onOpen}>オファーを送る</PrimaryButton>
        </div>
      </>
    )

    return (
      <div>
        <div
          ref={defaultContentRef}
          onMouseEnter={handleMouseEnter}
          onMouseLeave={handleMouseLeave}
        >
          {defaultContent}
        </div>

        <div
          ref={subContentRef}
          onMouseEnter={handleMouseEnter}
          onMouseLeave={handleMouseLeave}
        >
          {subContent}
        </div>
        <div ref={modalRef}>
          {isOpen && (
            <Modal>
              <div
                className={clsx(
                  'justify-left ml-4 mb-4 w-fit cursor-pointer items-center',
                )}
              >
                <Close width="32" height="32" onClick={onClose} />
              </div>
              <div className="border-primary-1 mb-4 border-l-8 pl-2 text-left text-3xl font-bold">
                オファーを送る
              </div>
              <div>
                <span className="border-primary-1 mb-2 border-b-2 text-left text-xl font-bold">
                  希望する条件
                </span>
              </div>

              <div className="grid grid-cols-4 gap-2 pt-2 pb-4 text-left">
                {props.data.user.request.detail.detailLimits.map(
                  (detailLimits: {
                    id: number
                    classificationName: string
                    limits: Limit[]
                  }) => (
                    <div className="col-span-1">
                      <span className="text-md">
                        {detailLimits.limits.map((limit: any) => (
                          <>
                            {limit.name === '賃料' && (
                              <div>
                                <span className=" pt-1 text-lg font-bold">
                                  {limit.name}
                                </span>
                                <br />
                                <span>
                                  {limit.minValue.name}〜{limit.maxValue.name}
                                </span>
                              </div>
                            )}
                          </>
                        ))}
                      </span>
                    </div>
                  ),
                )}
                {props.data.user.request.detail.detailValues.map(
                  (detailValues: any) => (
                    <div className="col-span-1">
                      <span className="text-md">
                        {detailValues.values.map((value: any) => (
                          <>
                            <span className=" mt-2 text-lg font-bold">
                              {value.name}
                            </span>
                            <br />
                            <span>{value.value?.name}</span>
                          </>
                        ))}
                      </span>
                    </div>
                  ),
                )}
                {props.data.user.request.detail.detailTags.map(
                  (detailTags: any) => (
                    <>
                      {detailTags.classificationName === '間取り' && (
                        <div className="col-span-1">
                          <span className=" pt-1 text-lg font-bold">
                            {detailTags.classificationName}
                          </span>
                          <br />
                          <span className="text-md">
                            {detailTags.tags.map((tag: Tag, index: number) => (
                              <>
                                <span>{tag.name}</span>
                                {index !== detailTags.tags.length - 1 && (
                                  <span>, </span>
                                )}
                              </>
                            ))}
                          </span>
                        </div>
                      )}
                      {detailTags.tags.map((tag: any) => (
                        <>
                          {tag.name === '間取り' && (
                            <div>
                              <div className="pt-1 text-lg font-bold">
                                {tag.name}
                              </div>
                            </div>
                          )}
                        </>
                      ))}
                      {detailTags.name === '間取り' && (
                        <div className="py-2 text-lg font-bold">
                          {detailTags.name}
                          <br />
                          <div className="w-1/1 flex flex-wrap text-sm">
                            {detailTags.tags.map((tag: Tag) => (
                              <Label
                                name={tag.name}
                                width={'w-auto'}
                                className={'mx-1 mt-2 '}
                              />
                            ))}
                          </div>
                        </div>
                      )}
                    </>
                  ),
                )}
              </div>

              <div>
                <span className="border-primary-1 mb-2 border-b-2 text-left text-xl font-bold">
                  オファー内容
                </span>
              </div>
              <Textarea className=" mt-3" />
              <div className="border-primary-1 mt-2 flex justify-center pt-4 pb-2 text-2xl font-bold">
                <PrimaryButton onClick={onClose}>オファーを送る</PrimaryButton>
              </div>
            </Modal>
          )}
        </div>
      </div>
    )
  })

  return (
    <>
      <div className="px-10 pt-10 pb-3">
        <span className="text-start border-primary-1 border-l-8 pl-2 text-4xl">
          {/* {console.log(allData)} */}
          {area.prefecture} {area.city}
          で物件を探している人たち
        </span>
      </div>
      <div className="grid w-full sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
        {allData.map((data: any) => (
          <Card width="w-1/1">
            <CardContent data={data}></CardContent>
          </Card>
        ))}
      </div>
    </>
  )
}

export default Rentlist
