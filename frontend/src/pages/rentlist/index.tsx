import clsx from 'clsx'
import type { NextPage } from 'next'
import { useState, useEffect, memo, useRef } from 'react'
import { Card, Label, Modal, PrimaryButton, Textarea } from '@components/common'
import { Close } from '@components/icons'
import { User, Area, Value, Tag, PersonalInfo, CompanyInfo } from '@type/common'

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
  const companyInfo: CompanyInfo = {
    id: 1,
    name: '長岡金型',
    phoneNumber: '09012345678',
    postCode: '9456666',
    prefecture: '新潟県',
    city: '長岡市西陵町',
    addressNumber: '2674-31',
    buildingName: '長岡金型ビル',
    website: 'https://nagaoka-kanagata.co.jp',
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

  const tagList: Tag[] = [
    { id: 1, name: 'IHコンロ', created_at: '', updated_at: '' },
    { id: 2, name: 'バス・トイレ別', created_at: '', updated_at: '' },
    { id: 3, name: '1R', created_at: '', updated_at: '' },
    { id: 4, name: '1K', created_at: '', updated_at: '' },
  ]
  const tags = [
    {
      name: '条件',
      tags: [tagList[0], tagList[1]],
    },
    {
      name: '間取り',
      tags: [tagList[2], tagList[3]],
    },
  ]
  const valueList: Value[] = [
    {
      id: 1,
      name: '20㎡',
      value: 20,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 2,
      name: '30㎡',
      value: 30,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 3,
      name: '3万円',
      value: 3,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 4,
      name: '4万円',
      value: 4,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
    {
      id: 5,
      name: '新築',
      value: 0,
      rangeId: 1,
      created_at: '',
      updated_at: '',
    },
  ]
  const values = [
    {
      name: '築年数',
      values: [valueList[4]],
    },
  ]
  const ranges = [
    {
      name: '面積',
      values: [valueList[0], valueList[1]],
    },
    {
      name: '賃料',
      values: [valueList[2], valueList[3]],
    },
  ]
  const classifications = {
    value: values,
    range: ranges,
    tag: tags,
  }
  const details = {
    area: area,
    classification: classifications,
  }
  const requests = {
    detail: details,
  }
  const data = {
    user: user,
    request: requests,
  }
  const dataList = [data, data, data, data]
  // const displayData = [{ user: user, request: request, detail: detail }]

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
          <span className="border-primary-1 border-b-2 text-left text-xl font-bold">
            希望する条件
          </span>
        </div>
        <div className="grid grid-cols-2 text-left">
          {props.data.request.detail.classification.range.map((range: any) => (
            <>
              {range.name === '賃料' && (
                <span className="pt-2 pb-1 text-lg font-bold">
                  {range.name}
                  <br />
                  <span className="text-sm">
                    {range.values.map((value: Value, index: number) => (
                      <span>
                        {value.name}
                        {index < range.values.length - 1 && '〜'}
                      </span>
                    ))}
                  </span>
                </span>
              )}
            </>
          ))}
          {props.data.request.detail.classification.value.map(
            (valuelist: any) => (
              <>
                {valuelist.name === '築年数' && (
                  <span className="pt-2 pb-1 text-lg font-bold">
                    {valuelist.name}
                    <br />
                    <span className="text-sm">
                      {valuelist.values.map((value: Value) => (
                        <span>{value.name}</span>
                      ))}
                    </span>
                  </span>
                )}
              </>
            ),
          )}
        </div>
        <div className="text-left">
          {props.data.request.detail.classification.tag.map((taglist: any) => (
            <>
              {taglist.name === '間取り' && (
                <div className="py-2 text-lg font-bold">
                  {taglist.name}
                  <br />
                  <div className="w-1/1 flex flex-wrap text-sm">
                    {taglist.tags.map((tag: Tag) => (
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
          <span className="border-primary-1 border-b-2 text-left text-xl font-bold">
            その他の条件
          </span>
        </div>
        <div className="grid grid-cols-2 text-left">
          {data.request.detail.classification.range.map((range) => (
            <>
              {range.name != '賃料' && (
                <span className="pt-2 pb-1 text-lg font-bold">
                  {range.name}
                  <br />
                  <span className="text-sm">
                    {range.values.map((value, index) => (
                      <span>
                        {value.name}
                        {index < range.values.length - 1 && '〜'}
                      </span>
                    ))}
                  </span>
                </span>
              )}
            </>
          ))}
          {data.request.detail.classification.value.map((valuelist) => (
            <>
              {valuelist.name != '築年数' && (
                <span className="pt-2 pb-1 text-lg font-bold">
                  {valuelist.name}
                  <br />
                  <span className="text-sm">
                    {valuelist.values.map((value) => (
                      <span>{value.name}</span>
                    ))}
                  </span>
                </span>
              )}
            </>
          ))}
        </div>
        <div className="py-2">
          <span className="border-primary-1 border-b-2 text-left text-xl font-bold">
            より詳しい条件
          </span>
        </div>
        <div className="text-left">
          {data.request.detail.classification.tag.map((taglist) => (
            <>
              {taglist.name != '間取り' && (
                <div className="pb-2 text-lg font-bold">
                  {taglist.name}
                  <br />

                  <div className="w-1/1 flex flex-wrap text-sm">
                    {taglist.tags.map((tag) => (
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
          <div className="border-primary-1 mt-2 flex justify-center border-t-2 border-dotted pt-4 pb-2 text-xl font-bold">
            <PrimaryButton onClick={onOpen}>オファーを送る</PrimaryButton>
          </div>
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
                <span className="border-primary-1 border-b-2 text-left text-xl font-bold">
                  希望する条件
                </span>
              </div>
              <div className="grid grid-cols-4 text-left">
                {props.data.request.detail.classification.range.map(
                  (range: any) => (
                    <span className="pt-2 pb-1 text-lg font-bold">
                      {range.name}
                      <br />
                      <span className="text-sm">
                        {range.values.map((value: Value, index: number) => (
                          <span>
                            {value.name}
                            {index < range.values.length - 1 && '〜'}
                          </span>
                        ))}
                      </span>
                    </span>
                  ),
                )}
                {props.data.request.detail.classification.value.map(
                  (valuelist: any) => (
                    <>
                      {valuelist.name === '築年数' && (
                        <span className="pt-2 pb-1 text-lg font-bold">
                          {valuelist.name}
                          <br />
                          <span className="text-sm">
                            {valuelist.values.map((value: Value) => (
                              <span>{value.name}</span>
                            ))}
                          </span>
                        </span>
                      )}
                    </>
                  ),
                )}
              </div>
              <div className="text-left">
                {props.data.request.detail.classification.tag.map(
                  (taglist: any) => (
                    <div className="py-1 text-lg font-bold">
                      {taglist.name}
                      <br />
                      <div className="w-1/1 flex flex-wrap text-sm">
                        {taglist.tags.map((tag: Tag) => (
                          <Label
                            name={tag.name}
                            width={'w-auto'}
                            className={'m-1'}
                          />
                        ))}
                      </div>
                    </div>
                  ),
                )}
              </div>
              <div className="border-primary-1 mb-2 border-b-2 text-left text-xl font-bold">
                オファー内容
              </div>
              <Textarea className=" mt-3" />
              <div className="border-primary-1 mt-2 flex justify-center pt-4 pb-2 text-2xl font-bold">
                <PrimaryButton onClick={onOpen}>オファーを送る</PrimaryButton>
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
          {data.request.detail.area.prefecture} {data.request.detail.area.city}
          で物件を探している人たち
        </span>
      </div>
      <div className="grid w-full sm:grid-cols-1 md:grid-cols-2 lg:grid-cols-4">
        {dataList.map((data) => (
          <Card width="w-1/1">
            <CardContent data={data}></CardContent>
          </Card>
        ))}
      </div>
    </>
  )
}

export default Rentlist
